//will contain <StudentFormContainer> <TableContainer>
var React = require("react");
var Home = require("../components/Home");
var PropTypes = React.PropTypes;

var HomeContainer = React.createClass({
    propTypes: {
      isLoggedIn: PropTypes.bool.isRequired,
      user: PropTypes.string
    },

    render: function(){
        return (
            <Home isLoggedIn={this.props.isLoggedIn} user={this.props.user}/>
        )
    }
});

module.exports = HomeContainer;