import { alertInfo } from "@/lib/alerts"
import { $alerts, $allSystemsById } from "@/lib/stores"
import type { AlertRecord } from "@/types"
import { Plural, Trans } from "@lingui/react/macro"
import { useStore } from "@nanostores/react"
import { getPagePath } from "@nanostores/router"
import { useMemo } from "react"
import { $router, Link } from "./router"
import { Alert, AlertTitle, AlertDescription } from "./ui/alert"
import { Card, CardHeader, CardTitle, CardContent } from "./ui/card"

export const ActiveAlerts = () => {
	const alerts = useStore($alerts)
	const systems = useStore($allSystemsById)

	const activeAlerts = useMemo(() => {
		const activeAlerts: AlertRecord[] = []

		for (const systemId of Object.keys(alerts)) {
			for (const alert of alerts[systemId].values()) {
				if (alert.triggered && alert.name in alertInfo) {
					activeAlerts.push(alert)
				}
			}
		}

		return activeAlerts
	}, [alerts])

	if (activeAlerts.length === 0) {
		return null
	}

	return (
		<Card className="border-red-500/25 bg-red-500/[0.035] shadow-none">
			<CardHeader className="px-4 py-3 border-b border-red-500/15">
				<div>
					<CardTitle className="text-base">
						<Trans>Active Alerts</Trans>
					</CardTitle>
				</div>
			</CardHeader>
			<CardContent className="p-3">
				{activeAlerts.length > 0 && (
					<div className="grid sm:grid-cols-2 lg:grid-cols-3 2xl:grid-cols-4 gap-2">
						{activeAlerts.map((alert) => {
							const info = alertInfo[alert.name as keyof typeof alertInfo]
							return (
								<Alert key={alert.id} className="duration-150 bg-card border-red-500/15 hover:border-red-500/35">
									<info.icon className="h-4 w-4" />
									<AlertTitle>
										{systems[alert.system]?.name} {info.name()}
									</AlertTitle>
									<AlertDescription>
										{alert.name === "Status" ? (
											<Trans>Connection is down</Trans>
										) : info.invert ? (
											<Trans>
												Below {alert.value}
												{info.unit} in last <Plural value={alert.min} one="# minute" other="# minutes" />
											</Trans>
										) : (
											<Trans>
												Exceeds {alert.value}
												{info.unit} in last <Plural value={alert.min} one="# minute" other="# minutes" />
											</Trans>
										)}
									</AlertDescription>
									<Link
										href={getPagePath($router, "system", { id: systems[alert.system]?.id })}
										className="absolute inset-0 w-full h-full"
										aria-label="View system"
									></Link>
								</Alert>
							)
						})}
					</div>
				)}
			</CardContent>
		</Card>
	)
}
