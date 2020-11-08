
import React, { useState, ChangeEvent, FormEvent } from 'react'
import { Link } from 'react-router-dom'
import { Button, Form, Message, Segment } from 'semantic-ui-react'
import { useDispatch } from 'react-redux'
import { Dispatch } from 'redux'
import { History } from 'history';

import { PageSingleWrapper } from '@/components/PageSingleWrapper/PageSingleWrapper'
import { signIn } from '@/redux/actions/customerAction'

type Props = {
    history: History
}

const SignIn: React.FC<Props> = ({ history }) => {
    const [username, setUsername] = useState("")
    const [password, setPassword] = useState("")
    const [loading, setLoading] = useState(false)
    const dispatch: Dispatch<any> = useDispatch()

    const onChangeUsername = (event: ChangeEvent<HTMLInputElement>) => {
        const username = event.target.value
        setUsername(username)
    }

    const onChangePassword = (event: ChangeEvent<HTMLInputElement>) => {
        const password = event.target.value
        setPassword(password)
    }

    const handleSignIn = (event: FormEvent<HTMLFormElement>) => {
        event.preventDefault();

        setLoading(true);
        dispatch(signIn({ username, password }))
        //     .then(() => {
        //         history.push("/");
        //     })
        // .finally(() => setLoading(false))

    }

    return (
        <PageSingleWrapper title='Авторизация'>
            <Form size='large' onSubmit={handleSignIn}>
                <Segment stacked>
                    <Form.Input fluid icon='user' iconPosition='left' placeholder='Электронная почта' onChange={onChangeUsername} />
                    <Form.Input
                        fluid
                        icon='lock'
                        iconPosition='left'
                        placeholder='Пароль'
                        type='password'
                        onChange={onChangePassword}
                    />
                    <Button color='teal' fluid size='large' type='submit'>
                        Войти
                </Button>
                </Segment>
            </Form>
            <Message>
                Если вы первые у нас советуем пройти <Link to='/signup'>Регистрацию</Link>
            </Message>
        </PageSingleWrapper>
    )
}


export default SignIn