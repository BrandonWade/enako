export const ValidateNewEmail = (currentEmail, newEmail) => currentEmail !== newEmail;

export const ValidateEmailFormat = email => /^[^@]+@[^.@]+\..{2,}$/.test(email);
