import React from 'react';
import PasswordField from '../../molecules/PasswordField';
import './PasswordChangeForm.scss';

const PasswordChangeForm = ({
    changePassword = false,
    password = '',
    confirmPassword = '',
    children = [],
    setPassword = () => {},
    setConfirmPassword = () => {},
}) => {
    return (
        <div className='PasswordChangeForm-grid'>
            <PasswordField
                type='password'
                label={changePassword ? 'New Password' : 'Password'}
                value={password}
                autoComplete='current-password'
                onChange={e => setPassword(e.target.value)}
            />
            <div className='PasswordChangeForm-validationRules'>{children}</div>
            <PasswordField
                type='password'
                label='Confirm Password'
                value={confirmPassword}
                autoComplete='new-password'
                onChange={e => setConfirmPassword(e.target.value)}
            />
        </div>
    );
};

export default PasswordChangeForm;
