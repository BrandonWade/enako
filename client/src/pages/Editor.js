import React, { Component } from 'react';
import '../css/Editor.css';

class Editor extends Component {
    render() {
        return (
            <div className='Editor'>
                <div className='Editor-content'>
                    <div>
                        <label>Date</label>
                        <input type='text' />
                    </div>
                    <div>
                        <label>Category</label>
                        <select>
                            <option>item 1</option>
                            <option>item 2</option>
                            <option>item 3</option>
                        </select>
                    </div>
                    <div>
                        <label>Description</label>
                        <input type='text' />
                    </div>
                    <div>
                        <label>Amount</label>
                        <input type='text' />
                    </div>
                    <div>
                        <button>Cancel</button>
                        <button>Save</button>
                    </div>
                </div>
            </div>
        );
    }
}

export default Editor;
