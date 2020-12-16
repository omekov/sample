
import React, { useState, ChangeEvent, FormEvent } from 'react'
import { Link } from 'react-router-dom'
import { Button, Form, Message, Segment } from 'semantic-ui-react'
import { useDispatch } from 'react-redux'
import { Dispatch } from 'redux'
import { History } from 'history';
import Axios from 'axios'
import SignService from '@/services/sign.service'
import { PageSingleWrapper } from '@/components/PageSingleWrapper/PageSingleWrapper'
import { signIn } from '@/redux/actions/customerAction'
import { SIGNIN_SUCCESS, SIGNIN_FAIL } from '@/redux/types'
import { handlerError } from '@/services/request.service'
export type Props = {
    history: History
}
const SignIn: React.FC<Props> = ({ history }) => {
    const [username, setUsername] = useState('example@gmail.com')
    const [password, setPassword] = useState('121212')
    const [loading, setLoading] = useState(false)
    const dispatch: Dispatch<any> = useDispatch()
    const onChangeUsername = (event: ChangeEvent<HTMLInputElement>) => setUsername(event.target.value)
    const onChangePassword = (event: ChangeEvent<HTMLInputElement>) => setPassword(event.target.value)

    const handleSignIn = (event: FormEvent<HTMLFormElement>) => {
        event.preventDefault();
        setLoading(true);
        dispatch(signIn({ username, password }, history))
        // SignService.signIn({ username, password })
        //     .then((response) => {
        //         if (response.status == 200) {
        //             dispatch({
        //                 type: SIGNIN_SUCCESS,
        //                 payload: response.data
        //             })
        //             const headerValue = `Bearer ${response.data.accessToken}`
        //             Axios.defaults.headers.common['Authorization'] = headerValue
        //             localStorage.setItem('access_token', JSON.stringify(response.data.accessToken));
        //             localStorage.setItem('refresh_token', JSON.stringify(response.data.refreshToken));
        //             return
        //         }
        //     }, (error) => handlerError(error, SIGNIN_FAIL, dispatch))
        //     .finally(() => setLoading(false))

    }

    return (
        <PageSingleWrapper title='Авторизация' data-test='SignInComponent'>
            <Form size='large' onSubmit={handleSignIn}>
                <Segment stacked>
                    <Form.Input
                        fluid
                        icon='user'
                        iconPosition='left'
                        placeholder='Электронная почта'
                        type='email'
                        value={username}
                        onChange={onChangeUsername}
                        required

                    />
                    <Form.Input
                        fluid
                        icon='lock'
                        iconPosition='left'
                        placeholder='Пароль'
                        type='password'
                        value={password}
                        onChange={onChangePassword}
                        required
                    />
                    <Button color='teal' fluid size='large' type='submit' disabled={loading}>
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

export default SignIn;