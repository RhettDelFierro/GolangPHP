var React = require("react");
var LoginForm = require("../components/LoginForm");

var LoginFormContainer = React.createClass({
    render: function() {
        return (
            <LoginForm isLoggedIn={this.props.isLoggedIn}
                       user={this.props.user}
                       onUpdateLogin={this.props.onUpdateLogin}/>
        )
    }
});

module.exports = LoginFormContainer;