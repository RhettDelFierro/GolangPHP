var React = require("react");
var RegisterForm = require("../components/RegisterForm");
var userFunctions = require("../utils/userFunctions");

var RegisterFormContainer = React.createClass({
    //make the ajax calls from here.
    getInitialState: function () {
        //bet we get the values of the input texts and make an object out of it
        return {
            //isLoggedIn maybe will pass into /teachers/:users to re-render "/"
            isLoggedIn: false,
            user: "",
            email: "",
            password: "",
            //maybe have all the user info here: current grades, class average, etc.
            userInfo: []
        }
    },
    handleUpdateUser: function (e) {
        this.setState({
            user: e.target.value,
        });

        if(this.state.user.length >= 5){
            userFunctions.verifyName(this.state.user).then();
        } else {

        }
    },
    handleUpdateEmail: function (e) {
        this.setState({
            email: e.target.value,
        });
    },
    handleUpdatePassword: function (e) {
        this.setState({
            password: e.target.value,
        });
    },
    handleSubmitUser: function (e) {
        e.preventDefault();
        var username = this.state.username;
        //in case registration fails:
        this.setState({
            user: "",
            email: "",
            password: "",
            //maybe have all the user info here: current grades, class average, etc.
            userInfo: []
        });

        if (this.props.routeParams.playerOne) {
            //going to go to "/teachers/:users" with the user information to re-render.
            this.context.router.push({
                pathname: "/teachers/"+ this.state.username,
                query: {
                    playerOne: this.props.routeParams.playerOne,
                    playerTwo: this.state.username
                }
            })
        } else {
            //go to /playerTwo.
            this.context.router.push("/playerTwo/" + this.state.username);
        }
    },
    render: function () {
        return (
            <RegisterForm onupdateUser={this.handleUpdateUser}
                          onupdateEmail={this.handleUpdateEmail}
                          onupdatePassword={this.handleUpdatePassword}
                          onSubmitUser={this.handleSubmitUser}
                          user={this.state.user}/>
        )
    }
});

module.exports = RegisterFormContainer;