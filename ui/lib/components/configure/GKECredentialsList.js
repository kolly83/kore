import PropTypes from 'prop-types'
import { Typography, Card, List, Button, Drawer, message, Icon } from 'antd'
const { Text, Title } = Typography

import { kore } from '../../../config'
import Credentials from '../team/Credentials'
import ResourceList from '../configure/ResourceList'
import GKECredentialsForm from '../forms/GKECredentialsForm'
import apiRequest from '../../utils/api-request'
import apiPaths from '../../utils/api-paths'

class GKECredentialsList extends ResourceList {

  static propTypes = {
    style: PropTypes.object
  }

  createdMessage = 'GKE credentials created successfully'
  updatedMessage = 'GKE credentials updated successfully'

  async fetchComponentData() {
    const [ allTeams, gkeCredentials, allAllocations ] = await Promise.all([
      apiRequest(null, 'get', apiPaths.teams),
      apiRequest(null, 'get', apiPaths.team(kore.koreAdminTeamName).gkeCredentials),
      apiRequest(null, 'get', apiPaths.team(kore.koreAdminTeamName).allocations)
    ])
    allTeams.items = allTeams.items.filter(t => !kore.ignoreTeams.includes(t.metadata.name))
    gkeCredentials.items.forEach(gke => {
      gke.allocation = (allAllocations.items || []).find(alloc => alloc.metadata.name === gke.metadata.name)
    })
    return { resources: gkeCredentials, allTeams }
  }

  render() {
    const { resources, allTeams, edit, add } = this.state

    return (
      <Card style={this.props.style} title="GKE credentials" extra={<Button type="primary" onClick={this.add(true)}>+ New</Button>}>
        {!resources ? <Icon type="loading" /> : (
          <>
            <List
              dataSource={resources.items}
              renderItem={gke =>
                <Credentials
                  gke={gke}
                  allTeams={allTeams.items}
                  editGKECredential={this.edit}
                  handleUpdate={this.handleStatusUpdated}
                  refreshMs={2000}
                  stateResourceDataKey="gke"
                  resourceApiPath={`/teams/${kore.koreAdminTeamName}/gkecredentials/${gke.metadata.name}`}
                />
              }
            >
            </List>
            {edit ? (
              <Drawer
                title={
                  edit.allocation ? (
                    <div>
                      <Title level={4}>{edit.allocation.spec.name}</Title>
                      <Text>{edit.allocation.spec.summary}</Text>
                    </div>
                  ) : (
                    <Title level={4}>{edit.metadata.name}</Title>
                  )
                }
                visible={!!edit}
                onClose={this.edit(false)}
                width={700}
              >
                <GKECredentialsForm
                  team={kore.koreAdminTeamName}
                  allTeams={allTeams}
                  data={edit}
                  handleSubmit={this.handleEditSave}
                />
              </Drawer>
            ) : null}
            {add ? (
              <Drawer
                title={<Title level={4}>New GKE credentials</Title>}
                visible={add}
                onClose={this.add(false)}
                width={700}
              >
                <GKECredentialsForm
                  team={kore.koreAdminTeamName}
                  allTeams={allTeams}
                  handleSubmit={this.handleAddSave}
                />
              </Drawer>
            ) : null}
          </>
        )}
      </Card>
    )
  }
}

export default GKECredentialsList
