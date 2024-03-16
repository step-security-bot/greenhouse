/*
 * SPDX-FileCopyrightText: 2024-2026 SAP SE or an SAP affiliate company and Greenhouse contributors
 * SPDX-License-Identifier: Apache-2.0
 */

import {
  Container,
  DataGrid,
  DataGridCell,
  DataGridHeadCell,
  DataGridRow,
  Panel,
  PanelBody,
  Pill,
  Stack,
} from "juno-ui-components"
import React, { useMemo } from "react"
import humanizedTimePeriodToNow from "../lib/utils/humanizedTimePeriodToNow"
import useStore from "../store"
import ClusterPluginConfigList from "./ClusterPluginConfigList"
import ConditionList from "./ConditionList"
import NodeList from "./NodeList"
import ResourceStatusIcon from "./ResourceStatusIcon"

const ClusterDetail: React.FC<any> = () => {
  const clusterDetails = useStore((state) => state.clusterDetails)
  const showClusterDetails = useStore((state) => state.showClusterDetails)
  const setShowClusterDetails = useStore((state) => state.setShowClusterDetails)

  const clusterAge = useMemo(() => {
    if (clusterDetails.cluster?.metadata?.creationTimestamp) {
      return humanizedTimePeriodToNow(
        clusterDetails.cluster?.metadata?.creationTimestamp
      )
    }
    return "unknown"
  }, [clusterDetails.cluster?.metadata?.creationTimestamp])

  const onPanelClose = () => {
    setShowClusterDetails(false)
  }

  return (
    <Panel
      heading={
        <Stack gap="2">
          <ResourceStatusIcon status={clusterDetails.clusterStatus!} />
          <span>{clusterDetails.cluster?.metadata?.name! || "Not found"}</span>
        </Stack>
      }
      opened={!!showClusterDetails}
      onClose={onPanelClose}
      size="large"
    >
      <PanelBody>
        <Container px={false} py>
          <>
            <DataGrid columns={2}>
              <DataGridRow>
                <DataGridHeadCell>State</DataGridHeadCell>
                <DataGridCell>
                  {clusterDetails.clusterStatus?.state}
                </DataGridCell>
              </DataGridRow>
              <DataGridRow>
                <DataGridHeadCell>Age</DataGridHeadCell>
                <DataGridCell>{clusterAge}</DataGridCell>
              </DataGridRow>
              <DataGridRow>
                <DataGridHeadCell>Kubernetes Version</DataGridHeadCell>
                <DataGridCell>
                  {clusterDetails.cluster?.status?.kubernetesVersion ??
                    "unknown"}
                </DataGridCell>
              </DataGridRow>
              <DataGridRow>
                <DataGridHeadCell>Access Mode</DataGridHeadCell>
                <DataGridCell>
                  {clusterDetails.cluster?.spec?.accessMode ?? "unknown"}
                </DataGridCell>
              </DataGridRow>
              {clusterDetails.clusterStatus?.message && (
                <DataGridRow>
                  <DataGridHeadCell>Message</DataGridHeadCell>
                  <DataGridCell>
                    {clusterDetails.clusterStatus?.message}
                  </DataGridCell>
                </DataGridRow>
              )}
              {clusterDetails.cluster?.metadata?.labels && (
                <DataGridRow>
                  <DataGridHeadCell>Labels</DataGridHeadCell>
                  <DataGridCell>
                    <Stack gap="2" alignment="start" wrap={true}>
                      {Object.keys(
                        clusterDetails.cluster?.metadata?.labels!
                      ).map((labelKey) => {
                        const labelValue =
                          clusterDetails.cluster?.metadata?.labels![labelKey]
                        return (
                          <Pill
                            pillKeyLabel={labelKey}
                            pillKey={labelKey}
                            pillKeyValue={labelValue}
                            pillValue={labelValue}
                          />
                        )
                      })}
                    </Stack>
                  </DataGridCell>
                </DataGridRow>
              )}
              {clusterDetails.cluster?.status?.statusConditions?.conditions && (
                <DataGridRow>
                  <DataGridHeadCell>Conditions</DataGridHeadCell>
                  <DataGridCell>
                    <ConditionList
                      conditionArray={
                        clusterDetails.cluster?.status?.statusConditions
                          ?.conditions
                      }
                    />
                  </DataGridCell>
                </DataGridRow>
              )}
              {clusterDetails.pluginConfigs && (
                <ClusterPluginConfigList
                  pluginConfigs={clusterDetails.pluginConfigs}
                />
              )}
            </DataGrid>

            {clusterDetails.cluster?.status?.nodes &&
              Object.keys(clusterDetails.cluster?.status?.nodes).length > 0 && (
                <Container px={false} py>
                  <h2 className="text-xl font-bold mb-2 mt-8">Nodes</h2>
                  <NodeList cluster={clusterDetails.cluster} />
                </Container>
              )}
          </>
        </Container>
      </PanelBody>
    </Panel>
  )
}

export default ClusterDetail
