var React = require("react");
var LoginForm = require("../components/LoginForm");
var userFunction = require("..utils/userFunctions");

var LoginFormContainer = React.createClass({
    //don't forget to go to the logged in route.
    contextTypes: {
      router: React.PropTypes.object.isRequired
    },
    getInitialState: function(){
      return {
          isLoggedIn: false,
          user: "",
          password: "",
          login: ""
      }

    },
    handleUpdateUser: function(e){
        this.setState({
            login: e.target.value
        })
    },
    handleUpdatePassword: function(e){
        this.setState({
            password: e.target.value
        })
    },
    handleSubmitUser: function(){

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