import React from 'react';
import ReactDOM from 'react-dom';
import Api from "./src/services/Api";

class AnimalListing extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            animals: []
        };
        Api.fetch("/api/v1/pets").then(r => this.setState({animals: r}));
    }


    render() {
        return (
            <div className="row">
                {
                    this.state.animals.map(i => {
                        return (
                            <article className="col-md-4" key={i.id}>
                                <b>{i.name} (id: {i.id})</b>
                                <p>
                                    {i.age} years old
                                </p>
                            </article>
                        )
                    })
                }
            </div>
        )
    }
}


ReactDOM.render(<AnimalListing/>, document.getElementById("react-container"));