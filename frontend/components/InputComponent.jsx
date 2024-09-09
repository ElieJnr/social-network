import React from "react";
import { useToast } from "./ui/use-toast";
const InputComponent = ({ text }) => {
  const { toast } = useToast();
  if (!text.trim()) {
    toast({
      title: "Invalide Input",
      description: "Empty value",
    });
    return
  }
  if(text.trim().lenght() > 20) {
    toast({
      title: "Invalide Input",
      description: "Text too much long",
    });
    return
  }
 };

export default InputComponent;
