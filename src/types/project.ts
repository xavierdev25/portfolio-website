export interface Project {
    id: number;
    title: string;
    category: string;
    description: string;
    images: string[];
    technologies: string[];
    repositoryUrl?: string;
    liveUrl?: string;
}
