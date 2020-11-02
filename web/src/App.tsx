import React, { lazy, Suspense } from 'react'
import { Route, Switch } from 'react-router-dom';
import { HashRouter as Router } from 'react-router-dom';
import { Loader } from 'semantic-ui-react';
const pageSignUp = lazy(() => import('@/pages/SignUp'))
const pageSignIn = lazy(() => import('@/pages/SignIn'))
const pageLanding = lazy(() => import('@/pages/Landing'))

const App = () => (
    <Router>
        <Suspense fallback={<Loader active inline='centered' />}>
            <Switch>
                <Route exact path='/' component={pageLanding} />
                <Route path='/signin' component={pageSignIn} />
                <Route path='/signup' component={pageSignUp} />
                {/* <Route path='/signup' component={SignUp} /> */}
            </Switch>
        </Suspense>
    </Router>
)
export default App