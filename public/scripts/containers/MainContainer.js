var React = require('react');
var HomeContainer = require("./HomeContainer");
var NavbarContainer = require("./NavbarContainer");
var Header = require("../components/Header");

var MainContainer = React.createClass({
    getInitialState: function () {
        return {
            isLoggedIn: true,
            username: "Somebody"
        }
    },
    handleUpdateLogin: function (loggedIn, username) {
        this.setState({
            isLoggedIn: loggedIn, //!(this.state.isLoggedIn),
            username: username
        })
    },
    render: function () {
        //<HomeContainer> is a child.
        console.log("username: ", this.state.username);
        return (
            <div className="container-fluid">
                <NavbarContainer onUpdateLogin={this.handleUpdateLogin}
                                 isLoggedIn={this.state.isLoggedIn}
                                 user={this.state.username} />

                {React.cloneElement(this.props.children, {
                    isLoggedIn: this.state.isLoggedIn,
                    user: this.state.username
                })}
            </div>
        )
    }
});

module.exports = MainContainer;