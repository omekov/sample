  
import React from 'react'
import { Button, Form, Grid, Header, Message, Segment } from 'semantic-ui-react'

export const SignUp = () => (
  <Grid textAlign='center' style={{ height: '100vh' }} verticalAlign='middle'>
    <Grid.Column style={{ maxWidth: 450 }}>
      <Header as='h2' color='teal' textAlign='center'>
        {/* <Image src='/logo.png' /> */}
        Регистрация
      </Header>
      <Form size='large'>
        <Segment stacked>
          <Form.Input fluid icon='user' iconPosition='left' placeholder='Имя' />
          <Form.Input fluid icon='user' iconPosition='left' placeholder='Электронная почта' />
          <Form.Input
            fluid
            icon='lock'
            iconPosition='left'
            placeholder='Пароль'
            type='password'
          />
          <Form.Input
            fluid
            icon='lock'
            iconPosition='left'
            placeholder='Повторите пароль'
            type='password'
          />
          <Button color='teal' fluid size='large'>
            Зарегистрироваться 
          </Button>
        </Segment>
      </Form>
      <Message>
        Если раньше были у нас можете <a href='#'>Авторизоваться</a>
      </Message>
    </Grid.Column>
  </Grid>
)