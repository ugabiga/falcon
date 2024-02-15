"use client";

import React from "react";

export default function Privacy() {
    return (
        <main className="min-h-screen mt-12 pr-4 pl-4 md:max-w-[1200px] overflow-auto w-full mx-auto">
            <h1 className="scroll-m-20 text-4xl font-extrabold tracking-tight lg:text-5xl">
                개인 정보 처리 방침
            </h1>
            <h2 className="mt-10 scroll-m-20 border-b pb-2 text-3xl font-semibold tracking-tight transition-colors first:mt-0">
                1. 수집하는 개인정보 항목
            </h2>
            <p className="leading-7 [&:not(:first-child)]:mt-6">
                <p>
                    - Falcon은 서비스 이용 과정에서 다음과 같은 개인정보를 수집할 수 있습니다:
                </p>
                <p>
                    - 이용자의 식별 정보: 이름, 이메일 주소 등
                </p>
                <p>
                    - 거래 관련 정보: 건당 거래 금액, 거래소 연동 정보
                </p>
                <p>
                    - 기기 정보: IP 주소, 브라우저 종류 및 버전, 운영 체제 등
                </p>
            </p>

            <h2 className="mt-10 scroll-m-20 border-b pb-2 text-3xl font-semibold tracking-tight transition-colors first:mt-0">
                2. 개인정보의 수집 및 이용 목적
            </h2>

            <p className="leading-7 [&:not(:first-child)]:mt-6">
                <p>
                    - Falcon은 다음과 같은 목적으로 개인정보를 수집 및 이용합니다:
                </p>
                <p>
                    - 서비스 제공 및 운영
                </p>
                <p>
                    - 서비스 개선 및 유지보수
                </p>
                <p>
                    - 보안 및 법적 요구 준수
                </p>
            </p>

            <h2 className="mt-10 scroll-m-20 border-b pb-2 text-3xl font-semibold tracking-tight transition-colors first:mt-0">
                3. 개인정보의 보유 및 이용 기간
            </h2>

            <p className="leading-7 [&:not(:first-child)]:mt-6">
                - Falcon은 이용자의 개인정보를 서비스 제공 및 운영 목적으로 필요한 기간 동안 보유합니다. 법령에 따라 보존할 필요가 있는 경우 해당 법령에서 정한 기간 동안 보존될
                수 있습니다.
            </p>

            <h2 className="mt-10 scroll-m-20 border-b pb-2 text-3xl font-semibold tracking-tight transition-colors first:mt-0">
                4. 개인정보의 제공 및 공유
            </h2>

            <p className="leading-7 [&:not(:first-child)]:mt-6">
                <p>
                    - Falcon은 이용자의 동의 없이는 개인정보를 타 기업, 기관 또는 개인과 공유하지 않습니다. 단, 다음의 경우에는 예외로 합니다:
                </p>
                <p>
                    - 관련 법령에 따른 요청이 있는 경우
                </p>
                <p>
                    - 서비스 제공에 필요한 경우
                </p>
            </p>

            <h2 className="mt-10 scroll-m-20 border-b pb-2 text-3xl font-semibold tracking-tight transition-colors first:mt-0">
                5. 개인정보의 보안 조치
            </h2>
            <p className="leading-7 [&:not(:first-child)]:mt-6">
                <p>
                    - Falcon은 이용자의 개인정보 보호를 위해 다음과 같은 보안 조치를 시행합니다:
                </p>
                <p>
                    - 암호화된 통신을 통한 정보 전송
                </p>
                <p>
                    - 접근 제어를 위한 시스템 구축
                </p>
                <p>
                    - 개인정보 처리 직원의 교육 및 감독
                </p>
            </p>

            <h2 className="mt-10 scroll-m-20 border-b pb-2 text-3xl font-semibold tracking-tight transition-colors first:mt-0">
                6. 개인정보 처리 방침의 변경
            </h2>
            <p className="leading-7 [&:not(:first-child)]:mt-6">
                - Falcon은 법령의 개정이나 서비스의 변경 등에 따라 개인정보 처리 방침을 변경할 수 있습니다. 변경 사항이 있는 경우 변경 이전에 이용자에게 알립니다.
            </p>

            <h2 className="mt-10 scroll-m-20 border-b pb-2 text-3xl font-semibold tracking-tight transition-colors first:mt-0">
                7. 문의 및 불만 처리
            </h2>
            <p className="leading-7 [&:not(:first-child)]:mt-6">
                <p>
                    - 이용자는 개인정보 처리와 관련된 문의 및 불만을 위하여 다음의 연락처로 문의할 수 있습니다:
                </p>
                <p>
                    - 이메일: <a href="mailto:vultor.xyz@gmail.com" className="text-blue-500">
                    vultor.xyz@gmail.com
                </a>
                </p>
            </p>
            <p className="leading-7 [&:not(:first-child)]:mt-6 mb-6">
                본 개인정보 처리 방침은 2024/02/15 에 마지막으로 업데이트되었습니다.
            </p>
        </main>
    )
}
