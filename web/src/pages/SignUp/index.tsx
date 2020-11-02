
import React from 'react'
import { Link } from 'react-router-dom'
import { Button, Form, Message, Segment } from 'semantic-ui-react'
import { PageSingleWrapper } from '@/components/PageSingleWrapper/PageSingleWrapper'

const SignUp = () => (
    <PageSingleWrapper title='Регистрация'>
        <Form size='large'>
            <Segment stacked>
                <Form.Input fluid icon='user' iconPosition='left' placeholder='Имя' />
                <Form.Input fluid icon='user plus' iconPosition='left' placeholder='Электронная почта' />
                <Form.Input
                    fluid
                    icon='lock'
                    iconPosition='left'
                    placeholder='Пароль'
                    type='password'
                />
                <Form.Input
                    fluid
                    icon='repeat'
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
            Если раньше были у нас можете <Link to='/signin'>Авторизоваться</Link>
        </Message>
    </PageSingleWrapper>
)

export default SignUp