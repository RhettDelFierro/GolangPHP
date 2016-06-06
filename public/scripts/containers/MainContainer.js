var React = require('react');
var HomeContainer = require("./HomeContainer");

var MainContainer = React.createClass({
    getInitialState: function () {
        return {
            isLoggedIn: false
        }
    },
    handleUpdateLogin: function () {
        this.setState({
            isLoggedIn: true
        })
    },
    render: function () {
        //<HomeContainer> is a child.
        return (
            <div className="container-fluid">
                <NavBarContainer onUpdateLogin={this.handleUpdateLogin} isLoggedIn={this.state.isLoggedIn}/>

                {React.cloneElement(this.props.children, {isLoggedIn: this.state.isLoggedIn})}
            </div>
        )
    }
});

module.exports = MainContainer;