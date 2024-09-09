export const ValidateInput = (text) => {
  if (!text.trim()) {
    return {
      isValid: false,
      message: "Le texte est vide.",
    };
  }
  if (text.trim().length > 30) {
    return {
      isValid: false,
      message: "Le texte dépasse 20 caractères.",
    };
  }
  return {
    isValid: true,
    message: "Le texte est valide.",
  };
};
