export const ValidateUsernameLength = uname => uname.length >= 5 && uname.length <= 32;

export const ValidateUsernameCharacters = uname => /^[^\W_]+$/.test(uname);
