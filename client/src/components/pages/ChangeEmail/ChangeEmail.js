import React, { useState, useContext } from 'react';
import { Link, useHistory } from 'react-router-dom';
import MessageContext from '../../../contexts/MessageContext';
import changeEmail from '../../../effects/changeEmail';
import Card from '../../atoms/Card';
import Button from '../../atoms/Button';
import ValidationRow from '../../atoms/ValidationRow';
import InputField from '../../molecules/InputField';
import EmailChangeForm from '../../organisms/EmailChangeForm';
import { ValidateNewEmail, ValidateEmailFormat } from '../../../validators/email';
import './ChangeEmail.scss';

const ChangeEmail = () => {
    const history = useHistory();
    const { setMessages } = useContext(MessageContext);
    const [currentEmail, setCurrentEmail] = useState('foo@bar.net');
    const [password, setPassword] = useState('testpassword123');
    const [newEmail, setNewEmail] = useState('foo@bar.net');

    const validNewEmailFormat = ValidateEmailFormat(newEmail);
    const validEmailsDoNotMatch = ValidateNewEmail(currentEmail, newEmail);
    const isEmailValid = validNewEmailFormat && validEmailsDoNotMatch;

    const onSubmit = async () => {
        const data = {
            current_email: currentEmail,
            password: password,
            new_email: newEmail,
        };

        const response = await changeEmail(data);
        if (response?.messages?.length > 0) {
            setMessages(response.messages);
            return;
        }

        history.push('/logout');
    };

    return (
        <div className='ChangeEmail'>
            <Card className='ChangeEmail-content' heading='Change Email'>
                <div className='ChangeEmail-grid'>
                    <InputField
                        type='text'
                        label='Current Email'
                        value={currentEmail}
                        autoComplete='email'
                        onChange={e => setCurrentEmail(e.target.value)}
                    />
                    <InputField
                        type='password'
                        label='Password'
                        value={password}
                        autoComplete='current-password'
                        formClassName='ChangeEmail-password'
                        onChange={e => setPassword(e.target.value)}
                    />
                </div>
                <EmailChangeForm newEmail={newEmail} setNewEmail={setNewEmail}>
                    <ValidationRow valid={validNewEmailFormat} description='Email is well formatted' />
                    <ValidationRow valid={validEmailsDoNotMatch} description='Current and new emails do not match' />
                </EmailChangeForm>
                <div className='ChangeEmail-buttons'>
                    <Link to='/account'>
                        <Button text='Cancel' />
                    </Link>
                    <Button color='orange' text='Submit' onClick={onSubmit} disabled={!(currentEmail.length > 0 && isEmailValid)} />
                </div>
            </Card>
        </div>
    );
};

export default ChangeEmail;
