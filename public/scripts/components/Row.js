var React = require("react");
var RowContainer = require("../containers/RowContainer");


//these components may need a wrapper. May.
//however,
function Course(){

}

function Grade(){

}

function Student(props){

}

function Button(){

}

function Row(props) {
    return (
        <tr>
            <td>
                <Student />
                <Course />
                <Grade />
                <Button />
            </td>
        </tr>
    )
}

module.exports = Row;