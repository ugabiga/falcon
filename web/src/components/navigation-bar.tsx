import {NavigationMobileBar} from "@/components/navigation-mobile-bar";
import {NavigationDesktopBar} from "@/components/navigation-desktop-bar";

export default async function NavigationBar() {
    return (
        <header
            className="sticky top-0 z-50 border-b border-border/40 bg-background/95 backdrop-blur supports-[backdrop-filter]:bg-background/60">
            <div className="container flex h-14 max-w-screen-2xl items-center">
                <NavigationMobileBar/>
                <NavigationDesktopBar/>
            </div>
        </header>
    )
}








