var React = require("react");

function Header(props) {
    return (
        <div className="page-header">
            <h1 className="hidden-xs hidden-sm">Student Grade Table
                <small className="col-md-offset-6 text-right">Grade Average : <span className="avgGrade"></span></small>
            </h1>
            <h3 className="hidden-md hidden-lg">Student Grade Table
                <small className=" col-xs-offset-6 col-xs-6 text-right">Grade Average : <span className="avgGrade"></span>
                </small>
            </h3>
        </div>
    )
}

module.exports = Header;