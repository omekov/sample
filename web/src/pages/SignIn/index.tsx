
import React from 'react'
import { Link } from 'react-router-dom'
import { Button, Form, Message, Segment } from 'semantic-ui-react'
import { PageSingleWrapper } from '@components/PageSingleWrapper/PageSingleWrapper'
const SignIn = () => (
    <PageSingleWrapper title='Авторизация'>
        <Form size='large'>
            <Segment stacked>
                <Form.Input fluid icon='user' iconPosition='left' placeholder='Электронная почта' />
                <Form.Input
                    fluid
                    icon='lock'
                    iconPosition='left'
                    placeholder='Пароль'
                    type='password'
                />
                <Button color='teal' fluid size='large'>
                    Войти
          </Button>
            </Segment>
        </Form>
        <Message>
            Если вы первые у нас советуем пройти <Link to="/signup">Регистрацию</Link>
        </Message>
    </PageSingleWrapper>
)

export default SignIn