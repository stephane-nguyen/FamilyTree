export type Member = {
  id: number;
  firstname: string;
  middlename?: string;
  lastname: string;
  birthdate: string;
  gender: string;
  country: string;
  city: string;
  description?: string;
  photo?: string;
  [key: string]: any; // Adding index signature to allow any other properties
};
