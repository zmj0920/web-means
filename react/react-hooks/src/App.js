import React from 'react';
import { BrowserRouter as Router, Route, Link } from "react-router-dom"
import Index from './components/Index'
import List from './components/List'
import Example from './components/Example'
import Example2 from './components/Example2'
import Example3 from './Example/Example3'
import Example4 from './components/Example4'
import Example5 from './components/Example5'
function App() {
    return (
        <div>
            <Router>
                <ul>
                    <li> <Link to="/">useState</Link> </li>
                    <li><Link to="/index">useEffect</Link> </li>
                    <li><Link to="/list">useEffect1</Link> </li>
                    <li><Link to="/Example2">useContext&useReducer</Link> </li>
                    <li><Link to="/Example3">useReducer</Link> </li>
                    <li><Link to="/Example4">useMemo</Link> </li>
                    <li><Link to="/Example5">useRef</Link> </li>
                </ul>
                <Route path="/" exact component={Example} />
                <Route path="/index" component={Index} />
                <Route path="/list" component={List} />
                <Route path="/Example2" component={Example2} />
                <Route path="/Example3" component={Example3} />
                <Route path="/Example4" component={Example4} />
                <Route path="/Example5" component={Example5} />
            </Router>
        </div>
    )
}

export default App