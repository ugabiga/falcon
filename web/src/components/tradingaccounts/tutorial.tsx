"use client";

import {useAppDispatch, useAppSelector} from "@/store";
import {useTranslation} from "@/lib/i18n";
import React, {useEffect, useState} from "react";
import {setTradingAccountTutorial} from "@/store/tutorialSlice";
import {Dialog} from "@radix-ui/react-dialog";
import {DialogContent, DialogDescription, DialogHeader, DialogTitle} from "@/components/ui/dialog";


export function TradingAccountTutorial() {
    const dispatch = useAppDispatch()
    const tutorial = useAppSelector((state) => state.tutorial);
    const {t} = useTranslation()
    const [open, setOpen] = useState(false)

    useEffect(() => {
        const timer = setTimeout(() => {
            if (tutorial?.tradingAccountTutorial === false) {
                setOpen(true)
            }
        }, 500);

        return () => {
            clearTimeout(timer);
        };
    }, [tutorial]);

    useEffect(() => {
        if (!open) {
            dispatch(setTradingAccountTutorial(true))
        }
    }, [dispatch, open])

    const desc = "This page is where you can view and manage your trading accounts. You can add a new trading account by clicking the \"Add Trading Account\" button. You can also view your trading accounts in a table or in a card view. You can switch between the two views by clicking the \"Table\" or \"Cards\" button."

    return (
        <Dialog open={open} onOpenChange={setOpen}>
            <DialogContent className={"sm:max-w-[425px] overflow-y-scroll h-[calc(100dvh)] sm:h-auto"}>
                <DialogHeader>
                    <DialogTitle>
                        Trading Account Tutorial
                    </DialogTitle>
                    <DialogDescription>
                        <div className="pl-12 pr-12 pt-8 pb-8">
                            <p>
                                이 페이지는 작업을 수행하기 위한 거래소 정보를 등록, 관리하는 페이지입니다.
                            </p>
                            <p>
                                {"\"추가\" 버튼을 클릭하여 새로운 거래소 정보를 등록할 수 있습니다."}
                            </p>
                            <p>
                                {"\"계정이름\"은 등록할 거래소 정보에 붙일 닉네임과 같은것으로 임의로 지정할 수 있습니다."}
                            </p>
                            <p>
                                {"\"거래소\"는 현재 지원하는 거래소 목록입니다. 본인이 사용하고 있는 거래소를 선택해주세요."}
                            </p>
                            <p>
                                {"\"API Key\"와 \"Secret Key\"는 거래소에서 발급받은 API 정보를 입력해주세요."}
                            </p>
                            <p>
                                {"\"저장\"을 누르면 거래소 정보가 등록됩니다."}
                            </p>
                        </div>
                    </DialogDescription>
                </DialogHeader>
            </DialogContent>
        </Dialog>
    )
}
