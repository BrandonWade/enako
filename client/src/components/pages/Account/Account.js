import React from 'react';
import Card from '../../atoms/Card';
import { Link } from 'react-router-dom';
import Button from '../../atoms/Button';
import requestChangeEmail from '../../../effects/requestChangeEmail';
import './Account.scss';

const Account = ({ email = 'foo@bar.net' }) => {
    return (
        <div className='Account'>
            <Card className='Account-content' heading='Account'>
                <Card.Section>
                    <div className='Account-header'>
                        <p className='Account-email'>{email}</p>
                        <Link to='/logout'>
                            <Button color='orange' text='Logout' />
                        </Link>
                    </div>
                </Card.Section>
                <Card.Section heading='Change Password' description='Change your account password.'>
                    <Link to='/account/password'>
                        <Button text='Change Password' />
                    </Link>
                </Card.Section>
                <Card.Section heading='Change Email' description='Change the email address associated with your account.'>
                    <Button text='Change Email' onClick={() => requestChangeEmail()} />
                </Card.Section>
                <Card.Section heading='Download Data' description='Download a copy of your data.'>
                    <Button text='Download Data' />
                </Card.Section>
                <Card.Section
                    heading='Delete Account'
                    description='Permanently delete your account and all data associated with it. This cannot be undone.'
                >
                    <Button color='red' text='Delete Account' />
                </Card.Section>
                <div className='Account-buttons'>
                    <Link to='/'>
                        <Button color='orange' text='OK' />
                    </Link>
                </div>
            </Card>
        </div>
    );
};

export default Account;
