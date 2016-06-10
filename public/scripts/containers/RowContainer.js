var React = require("react");
var Row = require("../components/Row");

var RowContainer = React.createClass({
    componentWillMount: function(){
      //ajax call to get student info.
    },
    render: function () {
        return (
            <Row />
        )
    }
});

module.exports = RowContainer;