
import React from 'react'
import { Grid, Header } from 'semantic-ui-react'

interface PageSingleWrapperProps {
  children: React.ReactNode;
  title: string;
}

export const PageSingleWrapper: React.FC<PageSingleWrapperProps> = ({ children, title }: PageSingleWrapperProps): React.ReactElement => (
  <Grid textAlign='center' style={{ height: '100vh' }} verticalAlign='middle'>
    <Grid.Column style={{ maxWidth: 450 }}>
      <Header as='h2' color='teal' textAlign='center'>
        {/* <Image src='/logo.png' /> */}
        {title}
      </Header>
      {children}
    </Grid.Column>
  </Grid>
)