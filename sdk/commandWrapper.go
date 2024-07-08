package sdk

func BreakpointClearByID() {
	InterpreterEx("bc")
}

func BreakpointDisableByID() {
	InterpreterEx("bd")
}

func BreakpointEnableByID() {
	InterpreterEx("be")
}

func BreakpointList() {
	InterpreterEx("bl")
}

func SetBreakpoint0xcc() {
	InterpreterEx("bp")
}

func CoreOperatingProcessorForShowsAndChanges() {
	InterpreterEx("core")
}

func CpuFeaturesForCollectsReport() {
	InterpreterEx("cpu")
}

func FlushKernelModeBuffers() {
	InterpreterEx("flush")
}

func ContinueDebugger() {
	InterpreterEx("g")
}

func StepExecuteSingleInstructionOut() {
	InterpreterEx("gu")
}

func StepExecuteSingleInstructionIn() {
	InterpreterEx("i")
}

func CallstackOrThreadView() {
	InterpreterEx("k, kd, kq")
}

func ListModules() {
	InterpreterEx("lm")
}

func LoadDriversAndModules() {
	InterpreterEx("load")
}

func OutputEventForwardingInstance() {
	InterpreterEx("output")
}

func StepExecuteSingleInstruction() {
	InterpreterEx("p")
}

func PauseKernelEvents() {
	InterpreterEx("pause")
}

func PreactivateSpecialFunctionality() {
	InterpreterEx("preactivate")
}

func PreallocateBuffer() {
	InterpreterEx("prealloc")
}

func PrintEvaluateExpressions() {
	InterpreterEx("print")
}

func RegistersReadOrModify() {
	InterpreterEx("r")
}

func ReadMsr() {
	InterpreterEx("rdmsr")
}

func SearchMemoryPattern() {
	InterpreterEx("sb, !sb, sd, !sd, sq, !sq")
}

func SettingsManagement() {
	InterpreterEx("settings")
}

func SleepDebugger() {
	InterpreterEx("sleep")
}

func StepInExecute() {
	InterpreterEx("t")
}

func TestHyperDbgFeatures() {
	InterpreterEx("test")
}

func UnloadKernelModules() {
	InterpreterEx("unload")
}

func SearchSymbols() {
	InterpreterEx("x")
}

func CpuidExecutionMonitor() {
	InterpreterEx("!cpuid")
}

func ControlRegisterModificationMonitor() {
	InterpreterEx("!crwrite")
}

func DebugRegistersMonitor() {
	InterpreterEx("!dr")
}

func EptHook() {
	InterpreterEx("!epthook")
}

func EptHook2() {
	InterpreterEx("!epthook2")
}

func IdtEntriesMonitor() {
	InterpreterEx("!exception")
}

func HideHyperDbg() {
	InterpreterEx("!hide")
}

func InterruptExternalMonitor() {
	InterpreterEx("!interrupt")
}

func IoInDetect() {
	InterpreterEx("!ioin")
}

func IoOutDetect() {
	InterpreterEx("!ioout")
}

func MeasureArgumentsForHide() {
	InterpreterEx("!measure")
}

func ModeInstructionsTrap() {
	InterpreterEx("!mode")
}

func MonitorMemoryAccess() {
	InterpreterEx("!monitor")
}

func MsrRead() {
	InterpreterEx("!msrread")
}

func MsrWrite() {
	InterpreterEx("!msrwrite")
}

func Pa2Va() {
	InterpreterEx("!pa2va")
}

func PmcExecutionMonitor() {
	InterpreterEx("!pmc")
}

func Pte() {
	InterpreterEx("!pte")
}

func ReversingMachineModuleUse() {
	InterpreterEx("!rev")
}

func Syscall() {
	InterpreterEx("!syscall")
}

func SysRet() {
	InterpreterEx("!sysret")
}

func TraceExecution() {
	InterpreterEx("!trace")
}

func TrackModeTransitionInstructions() {
	InterpreterEx("!track")
}

func TscInstructionsMonitor() {
	InterpreterEx("!tsc")
}

func UnHide() {
	InterpreterEx("!unhide")
}

func Va2Pa() {
	InterpreterEx("!va2pa")
}

func VmCall() {
	InterpreterEx("!vmcall")
}

func HardwareClockDebugging() {
	InterpreterEx("!hw_clk")
}

func AttachDebugThread() {
	InterpreterEx(".attach")
}

func ClearScreen() {
	InterpreterEx(".cls")
}

func ConnectToMachine() {
	InterpreterEx(".connect")
}

func DebugMachine() {
	InterpreterEx(".debug")
}

func DetachDebugging() {
	InterpreterEx(".detach")
}

func DisconnectSession() {
	InterpreterEx(".disconnect")
}

func DumpMemoryContext() {
	InterpreterEx(".dump")
}

func FormatsDiff() {
	InterpreterEx(".formats")
}

func HelpForCommand() {
	InterpreterEx(".help")
}

func KillProcess() {
	InterpreterEx(".kill")
}

func ListenForClientConnection() {
	InterpreterEx(".listen")
}

func LogClose() {
	InterpreterEx(".logclose")
}

func LogOpen() {
	InterpreterEx(".logopen")
}

func PageAvailableInRam() {
	InterpreterEx(".pagein")
}

func ParseExecutableFiles() {
	InterpreterEx(".pe")
}

func ProcessesView() {
	InterpreterEx(".process")
}

func RestartProcess() {
	InterpreterEx(".restart")
}

func Script() {
	InterpreterEx(".script")
}

func StartProcess() {
	InterpreterEx(".start")
}

func Status() {
	InterpreterEx(".status")
}

func SwitchThread() {
	InterpreterEx(".switch")
}

func Symbol() {
	InterpreterEx(".sym")
}

func SymbolPath() {
	InterpreterEx(".sympath")
}

func Thread() {
	InterpreterEx(".thread")
}
