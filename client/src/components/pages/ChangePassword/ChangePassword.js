import React, { useState } from 'react';
import { Link } from 'react-router-dom';
import Card from '../../atoms/Card';
import InputField from '../../molecules/InputField';
import PasswordChangeForm from '../../organisms/PasswordChangeForm';
import Button from '../../atoms/Button';
import ValidationRow from '../../atoms/ValidationRow';
import {
    ValidateNewPassword,
    ValidatePasswordLength,
    ValidatePasswordCharacters,
    ValidatePasswordsMatch,
} from '../../../validators/password';
import './ChangePassword.scss';

const ChangePassword = () => {
    const [currentPassword, setCurrentPassword] = useState('');
    const [newPassword, setNewPassword] = useState('testpassword123');
    const [confirmPassword, setConfirmPassword] = useState('testpassword123');

    const validNewPassword = ValidateNewPassword(currentPassword, newPassword);
    const validPasswordLength = ValidatePasswordLength(newPassword);
    const validPasswordCharacters = ValidatePasswordCharacters(newPassword);
    const validPasswordsMatch = ValidatePasswordsMatch(newPassword, confirmPassword);
    const isPasswordValid = validNewPassword && validPasswordLength && validPasswordCharacters && validPasswordsMatch;

    const onChangePassword = () => {
        console.log('change password'); // TODO: Implement
    };

    return (
        <div className='ChangePassword'>
            <Card className='ChangePassword-content' heading='Change Password'>
                <div className='ChangePassword-grid'>
                    <InputField
                        type='password'
                        label='Current Password'
                        value={currentPassword}
                        autoComplete='current-password'
                        onChange={e => setCurrentPassword(e.target.value)}
                    />
                </div>
                <PasswordChangeForm
                    changePassword
                    password={newPassword}
                    confirmPassword={confirmPassword}
                    setPassword={setNewPassword}
                    setConfirmPassword={setConfirmPassword}
                >
                    <ValidationRow valid={validNewPassword} description='Current and new passwords do not match' />
                    <ValidationRow valid={validPasswordLength} description='Password is between 15 and 50 characters' />
                    <ValidationRow
                        valid={validPasswordCharacters}
                        description='Password contains only numbers, letters, and valid symbols: ! @ # $ % ^ * _'
                    />
                    <ValidationRow valid={validPasswordsMatch} description='Password and Confirm Password match' />
                </PasswordChangeForm>
                <div className='ChangePassword-buttons'>
                    <Link to='/account'>
                        <Button text='Cancel' />
                    </Link>
                    <Button color='orange' text='Submit' onClick={onChangePassword} disabled={!isPasswordValid} />
                </div>
            </Card>
        </div>
    );
};

export default ChangePassword;
