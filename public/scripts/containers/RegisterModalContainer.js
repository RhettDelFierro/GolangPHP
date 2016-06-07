var React = require("react");
var RegisterModal = require("../components/RegisterModal");

var RegisterModalContainer = React.createClass({
    getInitialState: function () {
        return { showModal: false };
    },
    handleClose: function () {
        this.setState({ showModal: false });
    },
    handleOpen() {
        this.setState({ showModal: true });
        console.log("has been called.");
    },
    render: function () {
        return (
            <RegisterModal onOpen={this.handleOpen} onClose={this.handleClose} modalToggle={this.state.showModal}/>
        )
    }
});

module.exports = RegisterModalContainer;