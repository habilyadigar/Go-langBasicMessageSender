import  React, { Component } from 'react'



export default class Form extends Component {
    constructor(props) {
      super(props);
      this.state = { name: '' };
      this.state = { mail: '' };
      this.state = { message: '' };
    }
    
  
    handleChange = (event) => {
      this.setState({[event.target.name]: event.target.value});
    }
  
    handleSubmit = (event) => {
        const user = {
            name : this.state.name,
            mail : this.state.mail,
            message : this.state.message
          }

      fetch('http://localhost:8080/user', {
          method: 'POST',
          // We convert the React state to JSON and send it as the POST body
          body: JSON.stringify(user)
          
        }).then(function(response) {
          console.log(response)
          
          return response.json();
        });
      console.log(this.state)
      event.preventDefault();
  }
    render() {
      return (
        <form onSubmit={this.handleSubmit}>
          <label>
            Name:
            <input type="text" value={this.state.value} name="name" onChange={this.handleChange} />
          </label>
          <label>
            e-Mail:
            <input type="text" value={this.state.value} name="mail" onChange={this.handleChange} />
          </label>
          <label>
            Message:
            <input type="text" value={this.state.value} name="message" onChange={this.handleChange} />
          </label>
          <input type="submit" value="Submit" />
        </form>
      );
    }
}