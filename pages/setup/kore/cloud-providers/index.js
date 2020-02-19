import React from 'react'
import { Layout, Typography, Card } from 'antd'
const { Footer } = Layout
const { Title, Paragraph } = Typography

import redirect from '../../../../lib/utils/redirect'
import copy from '../../../../lib/utils/object-copy'
import { kore } from '../../../../config'
import GKECredentialsForm from '../../../../lib/components/forms/GKECredentialsForm'
import CloudSelector from '../../../../lib/components/cluster-build/CloudSelector'

class ConfigureCloudProvidersPage extends React.Component {
  static propTypes = {}

  static staticProps = {
    title: 'Configure cluster providers',
    hideSider: true,
    adminOnly: true
  }

  state = {
    selectedCloud: ''
  }

  handleSelectCloud = cloud => {
    if (this.state.selectedCloud !== cloud) {
      const state = copy(this.state)
      state.selectedCloud = cloud
      this.setState(state)
    }
  }

  handleFormSubmit = () => {
    redirect(null, '/setup/kore/complete')
  }

  render() {
    const { selectedCloud } = this.state

    return (
      <div>
        <Title>Configure Cloud Cluster Provider</Title>
        <Paragraph>Choose your first cloud provider for your clusters, more can be configured later.</Paragraph>
        <div style={{ marginTop: '20px', marginBottom: '20px' }}>
          <CloudSelector selectedCloud={selectedCloud} handleSelectCloud={this.handleSelectCloud} />
        </div>

        { selectedCloud === 'GKE' ? (
          <Card title="Enter GKE credentials" style={{ paddingBottom: '0' }}>
            <GKECredentialsForm team={kore.koreAdminTeamName} handleSubmit={this.handleFormSubmit} />
          </Card>
        ) : null }

        <Footer style={{textAlign: 'center', backgroundColor: '#fff'}}>
          <span>
            For more information read the <a href="#">Documentation</a>
          </span>
        </Footer>
      </div>
    )
  }
}

export default ConfigureCloudProvidersPage
