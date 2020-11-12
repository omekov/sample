
import React, { useState, ChangeEvent, FormEvent } from 'react'
import { Link } from 'react-router-dom'
import { Button, Form, Message, Segment } from 'semantic-ui-react'
import { History } from 'history';
import { PageSingleWrapper } from '@/components/PageSingleWrapper/PageSingleWrapper'
import { Dispatch } from 'redux'
import { useDispatch } from 'react-redux'
import signService from '@/services/sign.service'
import { SIGNUP_SUCCESS, SET_MESSAGE, SIGNUP_FAIL } from '@/redux/types'
import { handlerError } from '@/services/request.service'
type Props = {
    history: History
}
const SIGNUP_SUCCESS_TEXT = 'Регистрация прошла успешно'
const SignUp: React.FC<Props> = ({ history }) => {
    const [username, setUsername] = useState("example@gmail.com")
    const [password, setPassword] = useState("123456")
    const [repeatPassword, setRepeatPassword] = useState("123456")
    const [firstname, setFirstname] = useState("Admin")
    const [loading, setLoading] = useState(false)
    const dispatch: Dispatch<any> = useDispatch()
    const onChangeUsername = (event: ChangeEvent<HTMLInputElement>) => setUsername(event.target.value)
    const onChangePassword = (event: ChangeEvent<HTMLInputElement>) => setPassword(event.target.value)
    const onChangeRepeatPassword = (event: ChangeEvent<HTMLInputElement>) => setRepeatPassword(event.target.value)
    const onChangeFirstname = (event: ChangeEvent<HTMLInputElement>) => setFirstname(event.target.value)
    const handleSignUp = (event: FormEvent<HTMLFormElement>) => {
        event.preventDefault();
        setLoading(true)
        signService.signUp({ username, password, repeatPassword, firstname })
            .then((response) => {
                if (response.status == 201) {
                    dispatch({
                        type: SIGNUP_SUCCESS
                    })
                    dispatch({
                        type: SET_MESSAGE,
                        payload: SIGNUP_SUCCESS_TEXT
                    })
                    return
                }
            }, (error) => handlerError(error, SIGNUP_FAIL, dispatch))
    }
    return (
        <PageSingleWrapper title='Регистрация'>
            <Form size='large' onSubmit={handleSignUp}>
                <Segment stacked>
                    <Form.Input
                        fluid
                        icon='user'
                        iconPosition='left'
                        placeholder='Имя'
                        value={firstname}
                        onChange={onChangeFirstname}
                    />
                    <Form.Input
                        fluid
                        icon='user plus'
                        iconPosition='left'
                        placeholder='Электронная почта'
                        type='email'
                        value={username}
                        onChange={onChangeUsername}
                    />
                    <Form.Input
                        fluid
                        icon='lock'
                        iconPosition='left'
                        placeholder='Пароль'
                        type='password'
                        value={password}
                        onChange={onChangePassword}
                    />
                    <Form.Input
                        fluid
                        icon='repeat'
                        iconPosition='left'
                        placeholder='Повторите пароль'
                        type='password'
                        value={repeatPassword}
                        onChange={onChangeRepeatPassword}

                    />
                    <Button color='teal' fluid size='large' disabled={loading}>
                        Зарегистрироваться
                    </Button>
                </Segment>
            </Form>
            <Message>
                Если раньше были у нас, можете <Link to='/signin'>Авторизоваться</Link>
            </Message>
        </PageSingleWrapper>
    )
}

export default SignUp