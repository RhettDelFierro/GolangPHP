var React = require('react');
var ReactRouter = require('react-router');
var Router = ReactRouter.Router;
var Route = ReactRouter.Route;
var hashHistory = ReactRouter.hashHistory;
var IndexRoute = ReactRouter.IndexRoute;
var MainContainer = require('../containers/MainContainer');
var HomeContainer = require('../containers/HomeContainer');
var RegisterFormContainer = require("../containers/RegisterFormContainer");
var LoginFormContainer = require("../containers/LoginFormContainer");

var routes = (
    <Router history={hashHistory}>
        <Route path="/" component={MainContainer}>
            <IndexRoute component={HomeContainer} />
            <Route path="teachers/:user" component={LoginFormContainer} />
        </Route>
    </Router>
);

module.exports = routes;