var React = require("react");

function Header(props) {
    return (
        <div className="page-header" style={{marginTop: "-15px", borderBottom: "hidden"}}>

            <h5 className="col-md-offset-6 text-right">
                Grade Average :
                <span className="avgGrade"></span></h5>

        </div>
    )
}

module.exports = Header;