var React = require("react");
var Table = require("..components/Table");

//methods to add:
//check if the user is the one who made the entry, authenticate if they can edit/delete it.
//populate
//add
//delete?

var TableContainer = React.createClass({

    render: function () {
        return (
            <Table />
        )
    }
});

module.exports = TableContainer;