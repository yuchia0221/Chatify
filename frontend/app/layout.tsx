import { Georama } from "next/font/google";
import "./globals.css";

const georama = Georama({ subsets: ["latin"] });

export const metadata = {
    title: "Home Page",
    description: "This is the home page.",
};

export default function RootLayout({ children }: { children: React.ReactNode }) {
    return (
        <html lang="en">
            <body className={georama.className}>{children}</body>
        </html>
    );
}
