var React = require("react");
var NavigationBar = require("../components/NavigationBar");
var PropTypes = React.PropTypes;

var NavbarContainer = React.createClass({
    propTypes: {
        isLoggedIn: PropTypes.bool.isRequired,
        onUpdateLogin: PropTypes.func.isRequired
    },
    render: function () {
        return (
            <NavigationBar isLoggedIn={this.props.isLoggedIn} user={this.props.user}/>
        )
    }
});

module.exports = NavbarContainer;