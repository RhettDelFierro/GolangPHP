var React = require("react");
var Header = require("./Header");
var PropTypes = React.PropTypes;
var StudentFormContainer = require("../containers/StudentFormContainer");
var TableContainer = require("../containers/TableContainer");


function Home(props) {

    return (
        <div>
            <Header />
            <StudentFormContainer />
            <TableContainer />
        </div>
    )

}

Home.propTypes = {
    isLoggedIn: PropTypes.bool.isRequired,
    user: PropTypes.string
};
module.exports = Home;