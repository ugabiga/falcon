"use client";

import * as React from "react";
import {useEffect} from "react";
import {X} from "lucide-react";

import {Badge} from "@/components/ui/badge";
import {Command, CommandGroup, CommandItem,} from "@/components/ui/command";
import {Command as CommandPrimitive} from "cmdk";
import {useTranslation} from "react-i18next";

type Days = Record<"value" | "label", string>;

const DAYS = [
    {
        value: "1",
        label: "monday"
    },
    {
        value: "2",
        label: "tuesday"
    },
    {
        value: "3",
        label: "wednesday"
    },
    {
        value: "4",
        label: "thursday"
    },
    {
        value: "5",
        label: "friday"
    },
    {
        value: "6",
        label: "saturday"
    },
    {
        value: "0",
        label: "sunday",
    }
] satisfies Days[];

function convertDaysInStringToDaysArray(daysInString?: string): Days[] {
    if (!daysInString) {
        return [];
    }

    if (daysInString === "*") {
        return DAYS;
    }

    const daysInArray = daysInString.split(",");

    return DAYS.filter(day => daysInArray.includes(day.value));
}

export function DaysOfWeekSelector(
    {
        selectedDaysInString,
        placeholder,
        onChange
    }: {
        selectedDaysInString?: string,
        placeholder?: string,
        onChange?: (days: string) => void
    }
) {
    const {t} = useTranslation();
    const inputRef = React.useRef<HTMLInputElement>(null);
    const [open, setOpen] = React.useState(false);
    const [selected, setSelected] = React.useState<Days[]>(convertDaysInStringToDaysArray(selectedDaysInString));
    const [inputValue, setInputValue] = React.useState("");

    useEffect(() => {
        if (onChange) {
            if (selected === DAYS) {
                onChange("*");
                return;
            }
            onChange(selected.map(s => s.value).join(","))
        }
    }, [selected]);

    const handleUnselect = React.useCallback((inputValues: Days) => {
        setSelected(prev => prev.filter(s => s.value !== inputValues.value));
    }, []);

    const handleKeyDown = React.useCallback((e: React.KeyboardEvent<HTMLDivElement>) => {
        const input = inputRef.current
        if (input) {
            if (e.key === "Delete" || e.key === "Backspace") {
                if (input.value === "") {
                    setSelected(prev => {
                        const newSelected = [...prev];
                        newSelected.pop();
                        return newSelected;
                    })
                }
            }
            // This is not a default behaviour of the <input /> field
            if (e.key === "Escape") {
                input.blur();
            }
        }
    }, []);

    const selectables = DAYS.filter(inputValues => !selected.includes(inputValues));

    return (
        <Command onKeyDown={handleKeyDown} className="overflow-visible bg-transparent">
            <div
                className="group border border-input px-3 py-2 text-sm ring-offset-background rounded-md focus-within:ring-2 focus-within:ring-ring focus-within:ring-offset-2"
            >
                <div className="flex gap-1 flex-wrap">
                    {selected.map((inputValue) => {
                        return (
                            <Badge key={inputValue.value} variant="secondary">
                                {t("common.days." + inputValue.label)}
                                <button
                                    className="ml-1 ring-offset-background rounded-full outline-none focus:ring-2 focus:ring-ring focus:ring-offset-2"
                                    onKeyDown={(e) => {
                                        if (e.key === "Enter") {
                                            handleUnselect(inputValue);
                                        }
                                    }}
                                    onMouseDown={(e) => {
                                        e.preventDefault();
                                        e.stopPropagation();
                                    }}
                                    onClick={() => handleUnselect(inputValue)}
                                >
                                    <X className="h-3 w-3 text-muted-foreground hover:text-foreground"/>
                                </button>
                            </Badge>
                        )
                    })}
                    {/* Avoid having the "Search" Icon */}
                    <CommandPrimitive.Input
                        ref={inputRef}
                        value={inputValue}
                        onValueChange={setInputValue}
                        onBlur={() => setOpen(false)}
                        onFocus={() => setOpen(true)}
                        placeholder={placeholder}
                        className="ml-2 bg-transparent outline-none placeholder:text-muted-foreground flex-1"
                    />
                </div>
            </div>
            <div className="relative mt-2">
                {open && selectables.length > 0 ?
                    <div
                        className="absolute w-full z-10 top-0 rounded-md border bg-popover text-popover-foreground shadow-md outline-none animate-in">
                        <CommandGroup className="h-full overflow-auto">
                            {selectables.map((inputValue) => {
                                return (
                                    <CommandItem
                                        key={inputValue.value}
                                        onMouseDown={(e) => {
                                            e.preventDefault();
                                            e.stopPropagation();
                                        }}
                                        onSelect={(value) => {
                                            setInputValue("")
                                            setSelected(prev => [...prev, inputValue])
                                        }}
                                        className={"cursor-pointer"}
                                    >
                                        {t("common.days." + inputValue.label)}
                                    </CommandItem>
                                );
                            })}
                        </CommandGroup>
                    </div>
                    : null}
            </div>
        </Command>
    )
}