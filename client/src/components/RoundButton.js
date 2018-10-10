import React, { Component } from 'react';
import '../css/RoundButton.css';

class RoundButton extends Component {
    render() {
        return (
            <div className='RoundButton'>
                <span className='RoundButton-text'>
                    {this.props.text}
                </span>
            </div>
        );
    }
}

export default RoundButton;
