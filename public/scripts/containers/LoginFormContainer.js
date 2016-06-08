var React = require("react");
var LoginForm = require("../components/LoginForm");

var LoginFormContainer = React.createClass({
    //don't forget to go to the logged in route.
    contextTypes: {
      router: React.PropTypes.object.isRequired
    },
    render: function() {
        return (
            <LoginForm isLoggedIn={this.props.isLoggedIn}
                       user={this.props.user}
                       onUpdateLogin={this.props.onUpdateLogin}/>
        )
    }
});



module.exports = LoginFormContainer;