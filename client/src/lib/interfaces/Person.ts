export interface Person {
    id: number;
    firstname: string;
    middlename?: string;
    lastname: string;
    birthdate: Date;
    gender: string;
    country: string;
    city: string;
    photo?: string;
    description: string;
}