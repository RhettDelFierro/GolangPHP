var React = require("react");
var Navbar = require("../components/Navbar");
var PropTypes = React.PropTypes;

var NavbarContainer = React.createClass({
    propTypes: {
        isLoggedIn: PropTypes.bool.isRequired,
        onUpdateLogin: PropTypes.func.isRequired
    },
    render: function () {
        return (
            <Navbar isLoggedIn={this.props.isLoggedIn}/>
        )
    }
});

module.exports = NavbarContainer;