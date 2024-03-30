import {Separator} from "@/components/ui/separator";

export default function SeparationHeader({name}: { name: string }) {
    return <>
        <Separator/>
        <div>
            <p className="text-lg">
                {name}
            </p>
        </div>
    </>
}