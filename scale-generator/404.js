import React from 'react'

import Layout from '../components/layout'
import SEO from '../components/seo'
import { Link } from 'gatsby'
import { Button, Header, Icon, Segment } from 'semantic-ui-react'


const NotFoundPage = () => (
  <Layout>
    <SEO title="404: Not found" />
    <Segment raised placeholder>
      <Header icon>
        <Icon name='frown outline' />
        Página não encontrada!
      </Header>
      <Link to='/'>
        <Button color='yellow'>Me leve pra casa</Button>
      </Link>
    </Segment>
  </Layout>
)

export default NotFoundPage
