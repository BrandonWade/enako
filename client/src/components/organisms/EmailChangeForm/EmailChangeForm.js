import React from 'react';
import InputField from '../../molecules/InputField';
import './EmailChangeForm.scss';

const EmailChangeForm = ({ newEmail = '', children = [], setNewEmail = () => {} }) => {
    return (
        <div className='EmailChangeForm-grid'>
            <InputField type='text' label='New Email' value={newEmail} autoComplete='email' onChange={e => setNewEmail(e.target.value)} />
            <div className='EmailChangeForm-validationRules'>{children}</div>
        </div>
    );
};

export default EmailChangeForm;
