export const ValidateNewPassword = (currentPword, newPword) => currentPword !== newPword;

export const ValidatePasswordLength = pword => pword.length >= 15 && pword.length <= 50;

export const ValidatePasswordCharacters = pword => /^[\w!@#$%^*]+$/.test(pword);

export const ValidatePasswordsMatch = (pword, confirmPword) => pword === confirmPword;
