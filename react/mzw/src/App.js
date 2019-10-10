import React, { Fragment } from 'react';
import { BrowserRouter as Router, Route } from "react-router-dom"
import Header from './components/Header/Header'
import Index from './pages/Index/Index'
function App() {
    return (
        <Fragment>
            <Header />
            <Router>
                <Route path="/" exact component={Index} />
            </Router>
        </Fragment>
    )
}
export default App