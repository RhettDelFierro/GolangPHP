var React = require('react');
var HomeContainer = require("./HomeContainer");

var MainContainer = React.createClass({
    render: function () {
        //need <NavContainer> <HomeContainer>
        return (
            <HomeContainer />
        )
    }
});

module.exports = MainContainer;