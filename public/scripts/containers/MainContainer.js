var React = require('react');
var HomeContainer = require("./HomeContainer");
var NavbarContainer = require("./NavbarContainer");

var MainContainer = React.createClass({
    getInitialState: function () {
        return {
            isLoggedIn: true,
            username: "Somebody"
        }
    },
    handleUpdateLogin: function () {
        this.setState({
            isLoggedIn: false,
            username: ""
        })
    },
    render: function () {
        //<HomeContainer> is a child.
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