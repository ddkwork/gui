package constants

type NtstatusKind int

const (
	STATUS_SUCCESS                                                     NtstatusKind = 0x00000000
	STATUS_WAIT_1                                                                   = 0x00000001
	STATUS_WAIT_2                                                                   = 0x00000002
	STATUS_WAIT_3                                                                   = 0x00000003
	STATUS_WAIT_63                                                                  = 0x0000003F
	STATUS_ABANDONED                                                                = 0x00000080
	STATUS_ABANDONED_WAIT_63                                                        = 0x000000BF
	STATUS_USER_APC                                                                 = 0x000000C0
	STATUS_KERNEL_APC                                                               = 0x00000100
	STATUS_ALERTED                                                                  = 0x00000101
	STATUS_TIMEOUT                                                                  = 0x00000102
	STATUS_PENDING                                                                  = 0x00000103
	STATUS_REPARSE                                                                  = 0x00000104
	STATUS_MORE_ENTRIES                                                             = 0x00000105
	STATUS_NOT_ALL_ASSIGNED                                                         = 0x00000106
	STATUS_SOME_NOT_MAPPED                                                          = 0x00000107
	STATUS_OPLOCK_BREAK_IN_PROGRESS                                                 = 0x00000108
	STATUS_VOLUME_MOUNTED                                                           = 0x00000109
	STATUS_RXACT_COMMITTED                                                          = 0x0000010A
	STATUS_NOTIFY_CLEANUP                                                           = 0x0000010B
	STATUS_NOTIFY_ENUM_DIR                                                          = 0x0000010C
	STATUS_NO_QUOTAS_FOR_ACCOUNT                                                    = 0x0000010D
	STATUS_PRIMARY_TRANSPORT_CONNECT_FAILED                                         = 0x0000010E
	STATUS_PAGE_FAULT_TRANSITION                                                    = 0x00000110
	STATUS_PAGE_FAULT_DEMAND_ZERO                                                   = 0x00000111
	STATUS_PAGE_FAULT_COPY_ON_WRITE                                                 = 0x00000112
	STATUS_PAGE_FAULT_GUARD_PAGE                                                    = 0x00000113
	STATUS_PAGE_FAULT_PAGING_FILE                                                   = 0x00000114
	STATUS_CACHE_PAGE_LOCKED                                                        = 0x00000115
	STATUS_CRASH_DUMP                                                               = 0x00000116
	STATUS_BUFFER_ALL_ZEROS                                                         = 0x00000117
	STATUS_REPARSE_OBJECT                                                           = 0x00000118
	STATUS_RESOURCE_REQUIREMENTS_CHANGED                                            = 0x00000119
	STATUS_TRANSLATION_COMPLETE                                                     = 0x00000120
	STATUS_DS_MEMBERSHIP_EVALUATED_LOCALLY                                          = 0x00000121
	STATUS_NOTHING_TO_TERMINATE                                                     = 0x00000122
	STATUS_PROCESS_NOT_IN_JOB                                                       = 0x00000123
	STATUS_PROCESS_IN_JOB                                                           = 0x00000124
	STATUS_VOLSNAP_HIBERNATE_READY                                                  = 0x00000125
	STATUS_FSFILTER_OP_COMPLETED_SUCCESSFULLY                                       = 0x00000126
	STATUS_INTERRUPT_VECTOR_ALREADY_CONNECTED                                       = 0x00000127
	STATUS_INTERRUPT_STILL_CONNECTED                                                = 0x00000128
	STATUS_PROCESS_CLONED                                                           = 0x00000129
	STATUS_FILE_LOCKED_WITH_ONLY_READERS                                            = 0x0000012A
	STATUS_FILE_LOCKED_WITH_WRITERS                                                 = 0x0000012B
	STATUS_RESOURCEMANAGER_READ_ONLY                                                = 0x00000202
	STATUS_RING_PREVIOUSLY_EMPTY                                                    = 0x00000210
	STATUS_RING_PREVIOUSLY_FULL                                                     = 0x00000211
	STATUS_RING_PREVIOUSLY_ABOVE_QUOTA                                              = 0x00000212
	STATUS_RING_NEWLY_EMPTY                                                         = 0x00000213
	STATUS_RING_SIGNAL_OPPOSITE_ENDPOINT                                            = 0x00000214
	STATUS_OPLOCK_SWITCHED_TO_NEW_HANDLE                                            = 0x00000215
	STATUS_OPLOCK_HANDLE_CLOSED                                                     = 0x00000216
	STATUS_WAIT_FOR_OPLOCK                                                          = 0x00000367
	DBG_EXCEPTION_HANDLED                                                           = 0x00010001
	DBG_CONTINUE                                                                    = 0x00010002
	STATUS_FLT_IO_COMPLETE                                                          = 0x001C0001
	STATUS_DIS_ATTRIBUTE_BUILT                                                      = 0x003C0001
	STATUS_OBJECT_NAME_EXISTS                                                       = 0x40000000
	STATUS_THREAD_WAS_SUSPENDED                                                     = 0x40000001
	STATUS_WORKING_SET_LIMIT_RANGE                                                  = 0x40000002
	STATUS_IMAGE_NOT_AT_BASE                                                        = 0x40000003
	STATUS_RXACT_STATE_CREATED                                                      = 0x40000004
	STATUS_SEGMENT_NOTIFICATION                                                     = 0x40000005
	STATUS_LOCAL_USER_SESSION_KEY                                                   = 0x40000006
	STATUS_BAD_CURRENT_DIRECTORY                                                    = 0x40000007
	STATUS_SERIAL_MORE_WRITES                                                       = 0x40000008
	STATUS_REGISTRY_RECOVERED                                                       = 0x40000009
	STATUS_FT_READ_RECOVERY_FROM_BACKUP                                             = 0x4000000A
	STATUS_FT_WRITE_RECOVERY                                                        = 0x4000000B
	STATUS_SERIAL_COUNTER_TIMEOUT                                                   = 0x4000000C
	STATUS_NULL_LM_PASSWORD                                                         = 0x4000000D
	STATUS_IMAGE_MACHINE_TYPE_MISMATCH                                              = 0x4000000E
	STATUS_RECEIVE_PARTIAL                                                          = 0x4000000F
	STATUS_RECEIVE_EXPEDITED                                                        = 0x40000010
	STATUS_RECEIVE_PARTIAL_EXPEDITED                                                = 0x40000011
	STATUS_EVENT_DONE                                                               = 0x40000012
	STATUS_EVENT_PENDING                                                            = 0x40000013
	STATUS_CHECKING_FILE_SYSTEM                                                     = 0x40000014
	STATUS_FATAL_APP_EXIT                                                           = 0x40000015
	STATUS_PREDEFINED_HANDLE                                                        = 0x40000016
	STATUS_WAS_UNLOCKED                                                             = 0x40000017
	STATUS_SERVICE_NOTIFICATION                                                     = 0x40000018
	STATUS_WAS_LOCKED                                                               = 0x40000019
	STATUS_LOG_HARD_ERROR                                                           = 0x4000001A
	STATUS_ALREADY_WIN32                                                            = 0x4000001B
	STATUS_WX86_UNSIMULATE                                                          = 0x4000001C
	STATUS_WX86_CONTINUE                                                            = 0x4000001D
	STATUS_WX86_SINGLE_STEP                                                         = 0x4000001E
	STATUS_WX86_BREAKPOINT                                                          = 0x4000001F
	STATUS_WX86_EXCEPTION_CONTINUE                                                  = 0x40000020
	STATUS_WX86_EXCEPTION_LASTCHANCE                                                = 0x40000021
	STATUS_WX86_EXCEPTION_CHAIN                                                     = 0x40000022
	STATUS_IMAGE_MACHINE_TYPE_MISMATCH_EXE                                          = 0x40000023
	STATUS_NO_YIELD_PERFORMED                                                       = 0x40000024
	STATUS_TIMER_RESUME_IGNORED                                                     = 0x40000025
	STATUS_ARBITRATION_UNHANDLED                                                    = 0x40000026
	STATUS_CARDBUS_NOT_SUPPORTED                                                    = 0x40000027
	STATUS_WX86_CREATEWX86TIB                                                       = 0x40000028
	STATUS_MP_PROCESSOR_MISMATCH                                                    = 0x40000029
	STATUS_HIBERNATED                                                               = 0x4000002A
	STATUS_RESUME_HIBERNATION                                                       = 0x4000002B
	STATUS_FIRMWARE_UPDATED                                                         = 0x4000002C
	STATUS_DRIVERS_LEAKING_LOCKED_PAGES                                             = 0x4000002D
	STATUS_MESSAGE_RETRIEVED                                                        = 0x4000002E
	STATUS_SYSTEM_POWERSTATE_TRANSITION                                             = 0x4000002F
	STATUS_ALPC_CHECK_COMPLETION_LIST                                               = 0x40000030
	STATUS_SYSTEM_POWERSTATE_COMPLEX_TRANSITION                                     = 0x40000031
	STATUS_ACCESS_AUDIT_BY_POLICY                                                   = 0x40000032
	STATUS_ABANDON_HIBERFILE                                                        = 0x40000033
	STATUS_BIZRULES_NOT_ENABLED                                                     = 0x40000034
	DBG_REPLY_LATER                                                                 = 0x40010001
	DBG_UNABLE_TO_PROVIDE_HANDLE                                                    = 0x40010002
	DBG_TERMINATE_THREAD                                                            = 0x40010003
	DBG_TERMINATE_PROCESS                                                           = 0x40010004
	DBG_CONTROL_C                                                                   = 0x40010005
	DBG_PRINTEXCEPTION_C                                                            = 0x40010006
	DBG_RIPEXCEPTION                                                                = 0x40010007
	DBG_CONTROL_BREAK                                                               = 0x40010008
	DBG_COMMAND_EXCEPTION                                                           = 0x40010009
	STATUS_HEURISTIC_DAMAGE_POSSIBLE                                                = 0x40190001
	STATUS_GUARD_PAGE_VIOLATION                                                     = 0x80000001
	STATUS_DATATYPE_MISALIGNMENT                                                    = 0x80000002
	STATUS_BREAKPOINT                                                               = 0x80000003
	STATUS_SINGLE_STEP                                                              = 0x80000004
	STATUS_BUFFER_OVERFLOW                                                          = 0x80000005
	STATUS_NO_MORE_FILES                                                            = 0x80000006
	STATUS_WAKE_SYSTEM_DEBUGGER                                                     = 0x80000007
	STATUS_HANDLES_CLOSED                                                           = 0x8000000A
	STATUS_NO_INHERITANCE                                                           = 0x8000000B
	STATUS_GUID_SUBSTITUTION_MADE                                                   = 0x8000000C
	STATUS_PARTIAL_COPY                                                             = 0x8000000D
	STATUS_DEVICE_PAPER_EMPTY                                                       = 0x8000000E
	STATUS_DEVICE_POWERED_OFF                                                       = 0x8000000F
	STATUS_DEVICE_OFF_LINE                                                          = 0x80000010
	STATUS_DEVICE_BUSY                                                              = 0x80000011
	STATUS_NO_MORE_EAS                                                              = 0x80000012
	STATUS_INVALID_EA_NAME                                                          = 0x80000013
	STATUS_EA_LIST_INCONSISTENT                                                     = 0x80000014
	STATUS_INVALID_EA_FLAG                                                          = 0x80000015
	STATUS_VERIFY_REQUIRED                                                          = 0x80000016
	STATUS_EXTRANEOUS_INFORMATION                                                   = 0x80000017
	STATUS_RXACT_COMMIT_NECESSARY                                                   = 0x80000018
	STATUS_NO_MORE_ENTRIES                                                          = 0x8000001A
	STATUS_FILEMARK_DETECTED                                                        = 0x8000001B
	STATUS_MEDIA_CHANGED                                                            = 0x8000001C
	STATUS_BUS_RESET                                                                = 0x8000001D
	STATUS_END_OF_MEDIA                                                             = 0x8000001E
	STATUS_BEGINNING_OF_MEDIA                                                       = 0x8000001F
	STATUS_MEDIA_CHECK                                                              = 0x80000020
	STATUS_SETMARK_DETECTED                                                         = 0x80000021
	STATUS_NO_DATA_DETECTED                                                         = 0x80000022
	STATUS_REDIRECTOR_HAS_OPEN_HANDLES                                              = 0x80000023
	STATUS_SERVER_HAS_OPEN_HANDLES                                                  = 0x80000024
	STATUS_ALREADY_DISCONNECTED                                                     = 0x80000025
	STATUS_LONGJUMP                                                                 = 0x80000026
	STATUS_CLEANER_CARTRIDGE_INSTALLED                                              = 0x80000027
	STATUS_PLUGPLAY_QUERY_VETOED                                                    = 0x80000028
	STATUS_UNWIND_CONSOLIDATE                                                       = 0x80000029
	STATUS_REGISTRY_HIVE_RECOVERED                                                  = 0x8000002A
	STATUS_DLL_MIGHT_BE_INSECURE                                                    = 0x8000002B
	STATUS_DLL_MIGHT_BE_INCOMPATIBLE                                                = 0x8000002C
	STATUS_STOPPED_ON_SYMLINK                                                       = 0x8000002D
	STATUS_CANNOT_GRANT_REQUESTED_OPLOCK                                            = 0x8000002E
	STATUS_NO_ACE_CONDITION                                                         = 0x8000002F
	DBG_EXCEPTION_NOT_HANDLED                                                       = 0x80010001
	STATUS_CLUSTER_NODE_ALREADY_UP                                                  = 0x80130001
	STATUS_CLUSTER_NODE_ALREADY_DOWN                                                = 0x80130002
	STATUS_CLUSTER_NETWORK_ALREADY_ONLINE                                           = 0x80130003
	STATUS_CLUSTER_NETWORK_ALREADY_OFFLINE                                          = 0x80130004
	STATUS_CLUSTER_NODE_ALREADY_MEMBER                                              = 0x80130005
	STATUS_FLT_BUFFER_TOO_SMALL                                                     = 0x801C0001
	STATUS_FVE_PARTIAL_METADATA                                                     = 0x80210001
	STATUS_FVE_TRANSIENT_STATE                                                      = 0x80210002
	STATUS_UNSUCCESSFUL                                                             = 0xC0000001
	STATUS_NOT_IMPLEMENTED                                                          = 0xC0000002
	STATUS_INVALID_INFO_CLASS                                                       = 0xC0000003
	STATUS_INFO_LENGTH_MISMATCH                                                     = 0xC0000004
	STATUS_ACCESS_VIOLATION                                                         = 0xC0000005
	STATUS_IN_PAGE_ERROR                                                            = 0xC0000006
	STATUS_PAGEFILE_QUOTA                                                           = 0xC0000007
	STATUS_INVALID_HANDLE                                                           = 0xC0000008
	STATUS_BAD_INITIAL_STACK                                                        = 0xC0000009
	STATUS_BAD_INITIAL_PC                                                           = 0xC000000A
	STATUS_INVALID_CID                                                              = 0xC000000B
	STATUS_TIMER_NOT_CANCELED                                                       = 0xC000000C
	STATUS_INVALID_PARAMETER                                                        = 0xC000000D
	STATUS_NO_SUCH_DEVICE                                                           = 0xC000000E
	STATUS_NO_SUCH_FILE                                                             = 0xC000000F
	STATUS_INVALID_DEVICE_REQUEST                                                   = 0xC0000010
	STATUS_END_OF_FILE                                                              = 0xC0000011
	STATUS_WRONG_VOLUME                                                             = 0xC0000012
	STATUS_NO_MEDIA_IN_DEVICE                                                       = 0xC0000013
	STATUS_UNRECOGNIZED_MEDIA                                                       = 0xC0000014
	STATUS_NONEXISTENT_SECTOR                                                       = 0xC0000015
	STATUS_MORE_PROCESSING_REQUIRED                                                 = 0xC0000016
	STATUS_NO_MEMORY                                                                = 0xC0000017
	STATUS_CONFLICTING_ADDRESSES                                                    = 0xC0000018
	STATUS_NOT_MAPPED_VIEW                                                          = 0xC0000019
	STATUS_UNABLE_TO_FREE_VM                                                        = 0xC000001A
	STATUS_UNABLE_TO_DELETE_SECTION                                                 = 0xC000001B
	STATUS_INVALID_SYSTEM_SERVICE                                                   = 0xC000001C
	STATUS_ILLEGAL_INSTRUCTION                                                      = 0xC000001D
	STATUS_INVALID_LOCK_SEQUENCE                                                    = 0xC000001E
	STATUS_INVALID_VIEW_SIZE                                                        = 0xC000001F
	STATUS_INVALID_FILE_FOR_SECTION                                                 = 0xC0000020
	STATUS_ALREADY_COMMITTED                                                        = 0xC0000021
	STATUS_ACCESS_DENIED                                                            = 0xC0000022
	STATUS_BUFFER_TOO_SMALL                                                         = 0xC0000023
	STATUS_OBJECT_TYPE_MISMATCH                                                     = 0xC0000024
	STATUS_NONCONTINUABLE_EXCEPTION                                                 = 0xC0000025
	STATUS_INVALID_DISPOSITION                                                      = 0xC0000026
	STATUS_UNWIND                                                                   = 0xC0000027
	STATUS_BAD_STACK                                                                = 0xC0000028
	STATUS_INVALID_UNWIND_TARGET                                                    = 0xC0000029
	STATUS_NOT_LOCKED                                                               = 0xC000002A
	STATUS_PARITY_ERROR                                                             = 0xC000002B
	STATUS_UNABLE_TO_DECOMMIT_VM                                                    = 0xC000002C
	STATUS_NOT_COMMITTED                                                            = 0xC000002D
	STATUS_INVALID_PORT_ATTRIBUTES                                                  = 0xC000002E
	STATUS_PORT_MESSAGE_TOO_LONG                                                    = 0xC000002F
	STATUS_INVALID_PARAMETER_MIX                                                    = 0xC0000030
	STATUS_INVALID_QUOTA_LOWER                                                      = 0xC0000031
	STATUS_DISK_CORRUPT_ERROR                                                       = 0xC0000032
	STATUS_OBJECT_NAME_INVALID                                                      = 0xC0000033
	STATUS_OBJECT_NAME_NOT_FOUND                                                    = 0xC0000034
	STATUS_OBJECT_NAME_COLLISION                                                    = 0xC0000035
	STATUS_PORT_DISCONNECTED                                                        = 0xC0000037
	STATUS_DEVICE_ALREADY_ATTACHED                                                  = 0xC0000038
	STATUS_OBJECT_PATH_INVALID                                                      = 0xC0000039
	STATUS_OBJECT_PATH_NOT_FOUND                                                    = 0xC000003A
	STATUS_OBJECT_PATH_SYNTAX_BAD                                                   = 0xC000003B
	STATUS_DATA_OVERRUN                                                             = 0xC000003C
	STATUS_DATA_LATE_ERROR                                                          = 0xC000003D
	STATUS_DATA_ERROR                                                               = 0xC000003E
	STATUS_CRC_ERROR                                                                = 0xC000003F
	STATUS_SECTION_TOO_BIG                                                          = 0xC0000040
	STATUS_PORT_CONNECTION_REFUSED                                                  = 0xC0000041
	STATUS_INVALID_PORT_HANDLE                                                      = 0xC0000042
	STATUS_SHARING_VIOLATION                                                        = 0xC0000043
	STATUS_QUOTA_EXCEEDED                                                           = 0xC0000044
	STATUS_INVALID_PAGE_PROTECTION                                                  = 0xC0000045
	STATUS_MUTANT_NOT_OWNED                                                         = 0xC0000046
	STATUS_SEMAPHORE_LIMIT_EXCEEDED                                                 = 0xC0000047
	STATUS_PORT_ALREADY_SET                                                         = 0xC0000048
	STATUS_SECTION_NOT_IMAGE                                                        = 0xC0000049
	STATUS_SUSPEND_COUNT_EXCEEDED                                                   = 0xC000004A
	STATUS_THREAD_IS_TERMINATING                                                    = 0xC000004B
	STATUS_BAD_WORKING_SET_LIMIT                                                    = 0xC000004C
	STATUS_INCOMPATIBLE_FILE_MAP                                                    = 0xC000004D
	STATUS_SECTION_PROTECTION                                                       = 0xC000004E
	STATUS_EAS_NOT_SUPPORTED                                                        = 0xC000004F
	STATUS_EA_TOO_LARGE                                                             = 0xC0000050
	STATUS_NONEXISTENT_EA_ENTRY                                                     = 0xC0000051
	STATUS_NO_EAS_ON_FILE                                                           = 0xC0000052
	STATUS_EA_CORRUPT_ERROR                                                         = 0xC0000053
	STATUS_FILE_LOCK_CONFLICT                                                       = 0xC0000054
	STATUS_LOCK_NOT_GRANTED                                                         = 0xC0000055
	STATUS_DELETE_PENDING                                                           = 0xC0000056
	STATUS_CTL_FILE_NOT_SUPPORTED                                                   = 0xC0000057
	STATUS_UNKNOWN_REVISION                                                         = 0xC0000058
	STATUS_REVISION_MISMATCH                                                        = 0xC0000059
	STATUS_INVALID_OWNER                                                            = 0xC000005A
	STATUS_INVALID_PRIMARY_GROUP                                                    = 0xC000005B
	STATUS_NO_IMPERSONATION_TOKEN                                                   = 0xC000005C
	STATUS_CANT_DISABLE_MANDATORY                                                   = 0xC000005D
	STATUS_NO_LOGON_SERVERS                                                         = 0xC000005E
	STATUS_NO_SUCH_LOGON_SESSION                                                    = 0xC000005F
	STATUS_NO_SUCH_PRIVILEGE                                                        = 0xC0000060
	STATUS_PRIVILEGE_NOT_HELD                                                       = 0xC0000061
	STATUS_INVALID_ACCOUNT_NAME                                                     = 0xC0000062
	STATUS_USER_EXISTS                                                              = 0xC0000063
	STATUS_NO_SUCH_USER                                                             = 0xC0000064
	STATUS_GROUP_EXISTS                                                             = 0xC0000065
	STATUS_NO_SUCH_GROUP                                                            = 0xC0000066
	STATUS_MEMBER_IN_GROUP                                                          = 0xC0000067
	STATUS_MEMBER_NOT_IN_GROUP                                                      = 0xC0000068
	STATUS_LAST_ADMIN                                                               = 0xC0000069
	STATUS_WRONG_PASSWORD                                                           = 0xC000006A
	STATUS_ILL_FORMED_PASSWORD                                                      = 0xC000006B
	STATUS_PASSWORD_RESTRICTION                                                     = 0xC000006C
	STATUS_LOGON_FAILURE                                                            = 0xC000006D
	STATUS_ACCOUNT_RESTRICTION                                                      = 0xC000006E
	STATUS_INVALID_LOGON_HOURS                                                      = 0xC000006F
	STATUS_INVALID_WORKSTATION                                                      = 0xC0000070
	STATUS_PASSWORD_EXPIRED                                                         = 0xC0000071
	STATUS_ACCOUNT_DISABLED                                                         = 0xC0000072
	STATUS_NONE_MAPPED                                                              = 0xC0000073
	STATUS_TOO_MANY_LUIDS_REQUESTED                                                 = 0xC0000074
	STATUS_LUIDS_EXHAUSTED                                                          = 0xC0000075
	STATUS_INVALID_SUB_AUTHORITY                                                    = 0xC0000076
	STATUS_INVALID_ACL                                                              = 0xC0000077
	STATUS_INVALID_SID                                                              = 0xC0000078
	STATUS_INVALID_SECURITY_DESCR                                                   = 0xC0000079
	STATUS_PROCEDURE_NOT_FOUND                                                      = 0xC000007A
	STATUS_INVALID_IMAGE_FORMAT                                                     = 0xC000007B
	STATUS_NO_TOKEN                                                                 = 0xC000007C
	STATUS_BAD_INHERITANCE_ACL                                                      = 0xC000007D
	STATUS_RANGE_NOT_LOCKED                                                         = 0xC000007E
	STATUS_DISK_FULL                                                                = 0xC000007F
	STATUS_SERVER_DISABLED                                                          = 0xC0000080
	STATUS_SERVER_NOT_DISABLED                                                      = 0xC0000081
	STATUS_TOO_MANY_GUIDS_REQUESTED                                                 = 0xC0000082
	STATUS_GUIDS_EXHAUSTED                                                          = 0xC0000083
	STATUS_INVALID_ID_AUTHORITY                                                     = 0xC0000084
	STATUS_AGENTS_EXHAUSTED                                                         = 0xC0000085
	STATUS_INVALID_VOLUME_LABEL                                                     = 0xC0000086
	STATUS_SECTION_NOT_EXTENDED                                                     = 0xC0000087
	STATUS_NOT_MAPPED_DATA                                                          = 0xC0000088
	STATUS_RESOURCE_DATA_NOT_FOUND                                                  = 0xC0000089
	STATUS_RESOURCE_TYPE_NOT_FOUND                                                  = 0xC000008A
	STATUS_RESOURCE_NAME_NOT_FOUND                                                  = 0xC000008B
	STATUS_ARRAY_BOUNDS_EXCEEDED                                                    = 0xC000008C
	STATUS_FLOAT_DENORMAL_OPERAND                                                   = 0xC000008D
	STATUS_FLOAT_DIVIDE_BY_ZERO                                                     = 0xC000008E
	STATUS_FLOAT_INEXACT_RESULT                                                     = 0xC000008F
	STATUS_FLOAT_INVALID_OPERATION                                                  = 0xC0000090
	STATUS_FLOAT_OVERFLOW                                                           = 0xC0000091
	STATUS_FLOAT_STACK_CHECK                                                        = 0xC0000092
	STATUS_FLOAT_UNDERFLOW                                                          = 0xC0000093
	STATUS_INTEGER_DIVIDE_BY_ZERO                                                   = 0xC0000094
	STATUS_INTEGER_OVERFLOW                                                         = 0xC0000095
	STATUS_PRIVILEGED_INSTRUCTION                                                   = 0xC0000096
	STATUS_TOO_MANY_PAGING_FILES                                                    = 0xC0000097
	STATUS_FILE_INVALID                                                             = 0xC0000098
	STATUS_ALLOTTED_SPACE_EXCEEDED                                                  = 0xC0000099
	STATUS_INSUFFICIENT_RESOURCES                                                   = 0xC000009A
	STATUS_DFS_EXIT_PATH_FOUND                                                      = 0xC000009B
	STATUS_DEVICE_DATA_ERROR                                                        = 0xC000009C
	STATUS_DEVICE_NOT_CONNECTED                                                     = 0xC000009D
	STATUS_DEVICE_POWER_FAILURE                                                     = 0xC000009E
	STATUS_FREE_VM_NOT_AT_BASE                                                      = 0xC000009F
	STATUS_MEMORY_NOT_ALLOCATED                                                     = 0xC00000A0
	STATUS_WORKING_SET_QUOTA                                                        = 0xC00000A1
	STATUS_MEDIA_WRITE_PROTECTED                                                    = 0xC00000A2
	STATUS_DEVICE_NOT_READY                                                         = 0xC00000A3
	STATUS_INVALID_GROUP_ATTRIBUTES                                                 = 0xC00000A4
	STATUS_BAD_IMPERSONATION_LEVEL                                                  = 0xC00000A5
	STATUS_CANT_OPEN_ANONYMOUS                                                      = 0xC00000A6
	STATUS_BAD_VALIDATION_CLASS                                                     = 0xC00000A7
	STATUS_BAD_TOKEN_TYPE                                                           = 0xC00000A8
	STATUS_BAD_MASTER_BOOT_RECORD                                                   = 0xC00000A9
	STATUS_INSTRUCTION_MISALIGNMENT                                                 = 0xC00000AA
	STATUS_INSTANCE_NOT_AVAILABLE                                                   = 0xC00000AB
	STATUS_PIPE_NOT_AVAILABLE                                                       = 0xC00000AC
	STATUS_INVALID_PIPE_STATE                                                       = 0xC00000AD
	STATUS_PIPE_BUSY                                                                = 0xC00000AE
	STATUS_ILLEGAL_FUNCTION                                                         = 0xC00000AF
	STATUS_PIPE_DISCONNECTED                                                        = 0xC00000B0
	STATUS_PIPE_CLOSING                                                             = 0xC00000B1
	STATUS_PIPE_CONNECTED                                                           = 0xC00000B2
	STATUS_PIPE_LISTENING                                                           = 0xC00000B3
	STATUS_INVALID_READ_MODE                                                        = 0xC00000B4
	STATUS_IO_TIMEOUT                                                               = 0xC00000B5
	STATUS_FILE_FORCED_CLOSED                                                       = 0xC00000B6
	STATUS_PROFILING_NOT_STARTED                                                    = 0xC00000B7
	STATUS_PROFILING_NOT_STOPPED                                                    = 0xC00000B8
	STATUS_COULD_NOT_INTERPRET                                                      = 0xC00000B9
	STATUS_FILE_IS_A_DIRECTORY                                                      = 0xC00000BA
	STATUS_NOT_SUPPORTED                                                            = 0xC00000BB
	STATUS_REMOTE_NOT_LISTENING                                                     = 0xC00000BC
	STATUS_DUPLICATE_NAME                                                           = 0xC00000BD
	STATUS_BAD_NETWORK_PATH                                                         = 0xC00000BE
	STATUS_NETWORK_BUSY                                                             = 0xC00000BF
	STATUS_DEVICE_DOES_NOT_EXIST                                                    = 0xC00000C0
	STATUS_TOO_MANY_COMMANDS                                                        = 0xC00000C1
	STATUS_ADAPTER_HARDWARE_ERROR                                                   = 0xC00000C2
	STATUS_INVALID_NETWORK_RESPONSE                                                 = 0xC00000C3
	STATUS_UNEXPECTED_NETWORK_ERROR                                                 = 0xC00000C4
	STATUS_BAD_REMOTE_ADAPTER                                                       = 0xC00000C5
	STATUS_PRINT_QUEUE_FULL                                                         = 0xC00000C6
	STATUS_NO_SPOOL_SPACE                                                           = 0xC00000C7
	STATUS_PRINT_CANCELLED                                                          = 0xC00000C8
	STATUS_NETWORK_NAME_DELETED                                                     = 0xC00000C9
	STATUS_NETWORK_ACCESS_DENIED                                                    = 0xC00000CA
	STATUS_BAD_DEVICE_TYPE                                                          = 0xC00000CB
	STATUS_BAD_NETWORK_NAME                                                         = 0xC00000CC
	STATUS_TOO_MANY_NAMES                                                           = 0xC00000CD
	STATUS_TOO_MANY_SESSIONS                                                        = 0xC00000CE
	STATUS_SHARING_PAUSED                                                           = 0xC00000CF
	STATUS_REQUEST_NOT_ACCEPTED                                                     = 0xC00000D0
	STATUS_REDIRECTOR_PAUSED                                                        = 0xC00000D1
	STATUS_NET_WRITE_FAULT                                                          = 0xC00000D2
	STATUS_PROFILING_AT_LIMIT                                                       = 0xC00000D3
	STATUS_NOT_SAME_DEVICE                                                          = 0xC00000D4
	STATUS_FILE_RENAMED                                                             = 0xC00000D5
	STATUS_VIRTUAL_CIRCUIT_CLOSED                                                   = 0xC00000D6
	STATUS_NO_SECURITY_ON_OBJECT                                                    = 0xC00000D7
	STATUS_CANT_WAIT                                                                = 0xC00000D8
	STATUS_PIPE_EMPTY                                                               = 0xC00000D9
	STATUS_CANT_ACCESS_DOMAIN_INFO                                                  = 0xC00000DA
	STATUS_CANT_TERMINATE_SELF                                                      = 0xC00000DB
	STATUS_INVALID_SERVER_STATE                                                     = 0xC00000DC
	STATUS_INVALID_DOMAIN_STATE                                                     = 0xC00000DD
	STATUS_INVALID_DOMAIN_ROLE                                                      = 0xC00000DE
	STATUS_NO_SUCH_DOMAIN                                                           = 0xC00000DF
	STATUS_DOMAIN_EXISTS                                                            = 0xC00000E0
	STATUS_DOMAIN_LIMIT_EXCEEDED                                                    = 0xC00000E1
	STATUS_OPLOCK_NOT_GRANTED                                                       = 0xC00000E2
	STATUS_INVALID_OPLOCK_PROTOCOL                                                  = 0xC00000E3
	STATUS_INTERNAL_DB_CORRUPTION                                                   = 0xC00000E4
	STATUS_INTERNAL_ERROR                                                           = 0xC00000E5
	STATUS_GENERIC_NOT_MAPPED                                                       = 0xC00000E6
	STATUS_BAD_DESCRIPTOR_FORMAT                                                    = 0xC00000E7
	STATUS_INVALID_USER_BUFFER                                                      = 0xC00000E8
	STATUS_UNEXPECTED_IO_ERROR                                                      = 0xC00000E9
	STATUS_UNEXPECTED_MM_CREATE_ERR                                                 = 0xC00000EA
	STATUS_UNEXPECTED_MM_MAP_ERROR                                                  = 0xC00000EB
	STATUS_UNEXPECTED_MM_EXTEND_ERR                                                 = 0xC00000EC
	STATUS_NOT_LOGON_PROCESS                                                        = 0xC00000ED
	STATUS_LOGON_SESSION_EXISTS                                                     = 0xC00000EE
	STATUS_INVALID_PARAMETER_1                                                      = 0xC00000EF
	STATUS_INVALID_PARAMETER_2                                                      = 0xC00000F0
	STATUS_INVALID_PARAMETER_3                                                      = 0xC00000F1
	STATUS_INVALID_PARAMETER_4                                                      = 0xC00000F2
	STATUS_INVALID_PARAMETER_5                                                      = 0xC00000F3
	STATUS_INVALID_PARAMETER_6                                                      = 0xC00000F4
	STATUS_INVALID_PARAMETER_7                                                      = 0xC00000F5
	STATUS_INVALID_PARAMETER_8                                                      = 0xC00000F6
	STATUS_INVALID_PARAMETER_9                                                      = 0xC00000F7
	STATUS_INVALID_PARAMETER_10                                                     = 0xC00000F8
	STATUS_INVALID_PARAMETER_11                                                     = 0xC00000F9
	STATUS_INVALID_PARAMETER_12                                                     = 0xC00000FA
	STATUS_REDIRECTOR_NOT_STARTED                                                   = 0xC00000FB
	STATUS_REDIRECTOR_STARTED                                                       = 0xC00000FC
	STATUS_STACK_OVERFLOW                                                           = 0xC00000FD
	STATUS_NO_SUCH_PACKAGE                                                          = 0xC00000FE
	STATUS_BAD_FUNCTION_TABLE                                                       = 0xC00000FF
	STATUS_VARIABLE_NOT_FOUND                                                       = 0xC0000100
	STATUS_DIRECTORY_NOT_EMPTY                                                      = 0xC0000101
	STATUS_FILE_CORRUPT_ERROR                                                       = 0xC0000102
	STATUS_NOT_A_DIRECTORY                                                          = 0xC0000103
	STATUS_BAD_LOGON_SESSION_STATE                                                  = 0xC0000104
	STATUS_LOGON_SESSION_COLLISION                                                  = 0xC0000105
	STATUS_NAME_TOO_LONG                                                            = 0xC0000106
	STATUS_FILES_OPEN                                                               = 0xC0000107
	STATUS_CONNECTION_IN_USE                                                        = 0xC0000108
	STATUS_MESSAGE_NOT_FOUND                                                        = 0xC0000109
	STATUS_PROCESS_IS_TERMINATING                                                   = 0xC000010A
	STATUS_INVALID_LOGON_TYPE                                                       = 0xC000010B
	STATUS_NO_GUID_TRANSLATION                                                      = 0xC000010C
	STATUS_CANNOT_IMPERSONATE                                                       = 0xC000010D
	STATUS_IMAGE_ALREADY_LOADED                                                     = 0xC000010E
	STATUS_ABIOS_NOT_PRESENT                                                        = 0xC000010F
	STATUS_ABIOS_LID_NOT_EXIST                                                      = 0xC0000110
	STATUS_ABIOS_LID_ALREADY_OWNED                                                  = 0xC0000111
	STATUS_ABIOS_NOT_LID_OWNER                                                      = 0xC0000112
	STATUS_ABIOS_INVALID_COMMAND                                                    = 0xC0000113
	STATUS_ABIOS_INVALID_LID                                                        = 0xC0000114
	STATUS_ABIOS_SELECTOR_NOT_AVAILABLE                                             = 0xC0000115
	STATUS_ABIOS_INVALID_SELECTOR                                                   = 0xC0000116
	STATUS_NO_LDT                                                                   = 0xC0000117
	STATUS_INVALID_LDT_SIZE                                                         = 0xC0000118
	STATUS_INVALID_LDT_OFFSET                                                       = 0xC0000119
	STATUS_INVALID_LDT_DESCRIPTOR                                                   = 0xC000011A
	STATUS_INVALID_IMAGE_NE_FORMAT                                                  = 0xC000011B
	STATUS_RXACT_INVALID_STATE                                                      = 0xC000011C
	STATUS_RXACT_COMMIT_FAILURE                                                     = 0xC000011D
	STATUS_MAPPED_FILE_SIZE_ZERO                                                    = 0xC000011E
	STATUS_TOO_MANY_OPENED_FILES                                                    = 0xC000011F
	STATUS_CANCELLED                                                                = 0xC0000120
	STATUS_CANNOT_DELETE                                                            = 0xC0000121
	STATUS_INVALID_COMPUTER_NAME                                                    = 0xC0000122
	STATUS_FILE_DELETED                                                             = 0xC0000123
	STATUS_SPECIAL_ACCOUNT                                                          = 0xC0000124
	STATUS_SPECIAL_GROUP                                                            = 0xC0000125
	STATUS_SPECIAL_USER                                                             = 0xC0000126
	STATUS_MEMBERS_PRIMARY_GROUP                                                    = 0xC0000127
	STATUS_FILE_CLOSED                                                              = 0xC0000128
	STATUS_TOO_MANY_THREADS                                                         = 0xC0000129
	STATUS_THREAD_NOT_IN_PROCESS                                                    = 0xC000012A
	STATUS_TOKEN_ALREADY_IN_USE                                                     = 0xC000012B
	STATUS_PAGEFILE_QUOTA_EXCEEDED                                                  = 0xC000012C
	STATUS_COMMITMENT_LIMIT                                                         = 0xC000012D
	STATUS_INVALID_IMAGE_LE_FORMAT                                                  = 0xC000012E
	STATUS_INVALID_IMAGE_NOT_MZ                                                     = 0xC000012F
	STATUS_INVALID_IMAGE_PROTECT                                                    = 0xC0000130
	STATUS_INVALID_IMAGE_WIN_16                                                     = 0xC0000131
	STATUS_LOGON_SERVER_CONFLICT                                                    = 0xC0000132
	STATUS_TIME_DIFFERENCE_AT_DC                                                    = 0xC0000133
	STATUS_SYNCHRONIZATION_REQUIRED                                                 = 0xC0000134
	STATUS_DLL_NOT_FOUND                                                            = 0xC0000135
	STATUS_OPEN_FAILED                                                              = 0xC0000136
	STATUS_IO_PRIVILEGE_FAILED                                                      = 0xC0000137
	STATUS_ORDINAL_NOT_FOUND                                                        = 0xC0000138
	STATUS_ENTRYPOINT_NOT_FOUND                                                     = 0xC0000139
	STATUS_CONTROL_C_EXIT                                                           = 0xC000013A
	STATUS_LOCAL_DISCONNECT                                                         = 0xC000013B
	STATUS_REMOTE_DISCONNECT                                                        = 0xC000013C
	STATUS_REMOTE_RESOURCES                                                         = 0xC000013D
	STATUS_LINK_FAILED                                                              = 0xC000013E
	STATUS_LINK_TIMEOUT                                                             = 0xC000013F
	STATUS_INVALID_CONNECTION                                                       = 0xC0000140
	STATUS_INVALID_ADDRESS                                                          = 0xC0000141
	STATUS_DLL_INIT_FAILED                                                          = 0xC0000142
	STATUS_MISSING_SYSTEMFILE                                                       = 0xC0000143
	STATUS_UNHANDLED_EXCEPTION                                                      = 0xC0000144
	STATUS_APP_INIT_FAILURE                                                         = 0xC0000145
	STATUS_PAGEFILE_CREATE_FAILED                                                   = 0xC0000146
	STATUS_NO_PAGEFILE                                                              = 0xC0000147
	STATUS_INVALID_LEVEL                                                            = 0xC0000148
	STATUS_WRONG_PASSWORD_CORE                                                      = 0xC0000149
	STATUS_ILLEGAL_FLOAT_CONTEXT                                                    = 0xC000014A
	STATUS_PIPE_BROKEN                                                              = 0xC000014B
	STATUS_REGISTRY_CORRUPT                                                         = 0xC000014C
	STATUS_REGISTRY_IO_FAILED                                                       = 0xC000014D
	STATUS_NO_EVENT_PAIR                                                            = 0xC000014E
	STATUS_UNRECOGNIZED_VOLUME                                                      = 0xC000014F
	STATUS_SERIAL_NO_DEVICE_INITED                                                  = 0xC0000150
	STATUS_NO_SUCH_ALIAS                                                            = 0xC0000151
	STATUS_MEMBER_NOT_IN_ALIAS                                                      = 0xC0000152
	STATUS_MEMBER_IN_ALIAS                                                          = 0xC0000153
	STATUS_ALIAS_EXISTS                                                             = 0xC0000154
	STATUS_LOGON_NOT_GRANTED                                                        = 0xC0000155
	STATUS_TOO_MANY_SECRETS                                                         = 0xC0000156
	STATUS_SECRET_TOO_LONG                                                          = 0xC0000157
	STATUS_INTERNAL_DB_ERROR                                                        = 0xC0000158
	STATUS_FULLSCREEN_MODE                                                          = 0xC0000159
	STATUS_TOO_MANY_CONTEXT_IDS                                                     = 0xC000015A
	STATUS_LOGON_TYPE_NOT_GRANTED                                                   = 0xC000015B
	STATUS_NOT_REGISTRY_FILE                                                        = 0xC000015C
	STATUS_NT_CROSS_ENCRYPTION_REQUIRED                                             = 0xC000015D
	STATUS_DOMAIN_CTRLR_CONFIG_ERROR                                                = 0xC000015E
	STATUS_FT_MISSING_MEMBER                                                        = 0xC000015F
	STATUS_ILL_FORMED_SERVICE_ENTRY                                                 = 0xC0000160
	STATUS_ILLEGAL_CHARACTER                                                        = 0xC0000161
	STATUS_UNMAPPABLE_CHARACTER                                                     = 0xC0000162
	STATUS_UNDEFINED_CHARACTER                                                      = 0xC0000163
	STATUS_FLOPPY_VOLUME                                                            = 0xC0000164
	STATUS_FLOPPY_ID_MARK_NOT_FOUND                                                 = 0xC0000165
	STATUS_FLOPPY_WRONG_CYLINDER                                                    = 0xC0000166
	STATUS_FLOPPY_UNKNOWN_ERROR                                                     = 0xC0000167
	STATUS_FLOPPY_BAD_REGISTERS                                                     = 0xC0000168
	STATUS_DISK_RECALIBRATE_FAILED                                                  = 0xC0000169
	STATUS_DISK_OPERATION_FAILED                                                    = 0xC000016A
	STATUS_DISK_RESET_FAILED                                                        = 0xC000016B
	STATUS_SHARED_IRQ_BUSY                                                          = 0xC000016C
	STATUS_FT_ORPHANING                                                             = 0xC000016D
	STATUS_BIOS_FAILED_TO_CONNECT_INTERRUPT                                         = 0xC000016E
	STATUS_PARTITION_FAILURE                                                        = 0xC0000172
	STATUS_INVALID_BLOCK_LENGTH                                                     = 0xC0000173
	STATUS_DEVICE_NOT_PARTITIONED                                                   = 0xC0000174
	STATUS_UNABLE_TO_LOCK_MEDIA                                                     = 0xC0000175
	STATUS_UNABLE_TO_UNLOAD_MEDIA                                                   = 0xC0000176
	STATUS_EOM_OVERFLOW                                                             = 0xC0000177
	STATUS_NO_MEDIA                                                                 = 0xC0000178
	STATUS_NO_SUCH_MEMBER                                                           = 0xC000017A
	STATUS_INVALID_MEMBER                                                           = 0xC000017B
	STATUS_KEY_DELETED                                                              = 0xC000017C
	STATUS_NO_LOG_SPACE                                                             = 0xC000017D
	STATUS_TOO_MANY_SIDS                                                            = 0xC000017E
	STATUS_LM_CROSS_ENCRYPTION_REQUIRED                                             = 0xC000017F
	STATUS_KEY_HAS_CHILDREN                                                         = 0xC0000180
	STATUS_CHILD_MUST_BE_VOLATILE                                                   = 0xC0000181
	STATUS_DEVICE_CONFIGURATION_ERROR                                               = 0xC0000182
	STATUS_DRIVER_INTERNAL_ERROR                                                    = 0xC0000183
	STATUS_INVALID_DEVICE_STATE                                                     = 0xC0000184
	STATUS_IO_DEVICE_ERROR                                                          = 0xC0000185
	STATUS_DEVICE_PROTOCOL_ERROR                                                    = 0xC0000186
	STATUS_BACKUP_CONTROLLER                                                        = 0xC0000187
	STATUS_LOG_FILE_FULL                                                            = 0xC0000188
	STATUS_TOO_LATE                                                                 = 0xC0000189
	STATUS_NO_TRUST_LSA_SECRET                                                      = 0xC000018A
	STATUS_NO_TRUST_SAM_ACCOUNT                                                     = 0xC000018B
	STATUS_TRUSTED_DOMAIN_FAILURE                                                   = 0xC000018C
	STATUS_TRUSTED_RELATIONSHIP_FAILURE                                             = 0xC000018D
	STATUS_EVENTLOG_FILE_CORRUPT                                                    = 0xC000018E
	STATUS_EVENTLOG_CANT_START                                                      = 0xC000018F
	STATUS_TRUST_FAILURE                                                            = 0xC0000190
	STATUS_MUTANT_LIMIT_EXCEEDED                                                    = 0xC0000191
	STATUS_NETLOGON_NOT_STARTED                                                     = 0xC0000192
	STATUS_ACCOUNT_EXPIRED                                                          = 0xC0000193
	STATUS_POSSIBLE_DEADLOCK                                                        = 0xC0000194
	STATUS_NETWORK_CREDENTIAL_CONFLICT                                              = 0xC0000195
	STATUS_REMOTE_SESSION_LIMIT                                                     = 0xC0000196
	STATUS_EVENTLOG_FILE_CHANGED                                                    = 0xC0000197
	STATUS_NOLOGON_INTERDOMAIN_TRUST_ACCOUNT                                        = 0xC0000198
	STATUS_NOLOGON_WORKSTATION_TRUST_ACCOUNT                                        = 0xC0000199
	STATUS_NOLOGON_SERVER_TRUST_ACCOUNT                                             = 0xC000019A
	STATUS_DOMAIN_TRUST_INCONSISTENT                                                = 0xC000019B
	STATUS_FS_DRIVER_REQUIRED                                                       = 0xC000019C
	STATUS_IMAGE_ALREADY_LOADED_AS_DLL                                              = 0xC000019D
	STATUS_INCOMPATIBLE_WITH_GLOBAL_SHORT_NAME_REGISTRY_SETTING                     = 0xC000019E
	STATUS_SHORT_NAMES_NOT_ENABLED_ON_VOLUME                                        = 0xC000019F
	STATUS_SECURITY_STREAM_IS_INCONSISTENT                                          = 0xC00001A0
	STATUS_INVALID_LOCK_RANGE                                                       = 0xC00001A1
	STATUS_INVALID_ACE_CONDITION                                                    = 0xC00001A2
	STATUS_IMAGE_SUBSYSTEM_NOT_PRESENT                                              = 0xC00001A3
	STATUS_NOTIFICATION_GUID_ALREADY_DEFINED                                        = 0xC00001A4
	STATUS_NETWORK_OPEN_RESTRICTION                                                 = 0xC0000201
	STATUS_NO_USER_SESSION_KEY                                                      = 0xC0000202
	STATUS_USER_SESSION_DELETED                                                     = 0xC0000203
	STATUS_RESOURCE_LANG_NOT_FOUND                                                  = 0xC0000204
	STATUS_INSUFF_SERVER_RESOURCES                                                  = 0xC0000205
	STATUS_INVALID_BUFFER_SIZE                                                      = 0xC0000206
	STATUS_INVALID_ADDRESS_COMPONENT                                                = 0xC0000207
	STATUS_INVALID_ADDRESS_WILDCARD                                                 = 0xC0000208
	STATUS_TOO_MANY_ADDRESSES                                                       = 0xC0000209
	STATUS_ADDRESS_ALREADY_EXISTS                                                   = 0xC000020A
	STATUS_ADDRESS_CLOSED                                                           = 0xC000020B
	STATUS_CONNECTION_DISCONNECTED                                                  = 0xC000020C
	STATUS_CONNECTION_RESET                                                         = 0xC000020D
	STATUS_TOO_MANY_NODES                                                           = 0xC000020E
	STATUS_TRANSACTION_ABORTED                                                      = 0xC000020F
	STATUS_TRANSACTION_TIMED_OUT                                                    = 0xC0000210
	STATUS_TRANSACTION_NO_RELEASE                                                   = 0xC0000211
	STATUS_TRANSACTION_NO_MATCH                                                     = 0xC0000212
	STATUS_TRANSACTION_RESPONDED                                                    = 0xC0000213
	STATUS_TRANSACTION_INVALID_ID                                                   = 0xC0000214
	STATUS_TRANSACTION_INVALID_TYPE                                                 = 0xC0000215
	STATUS_NOT_SERVER_SESSION                                                       = 0xC0000216
	STATUS_NOT_CLIENT_SESSION                                                       = 0xC0000217
	STATUS_CANNOT_LOAD_REGISTRY_FILE                                                = 0xC0000218
	STATUS_DEBUG_ATTACH_FAILED                                                      = 0xC0000219
	STATUS_SYSTEM_PROCESS_TERMINATED                                                = 0xC000021A
	STATUS_DATA_NOT_ACCEPTED                                                        = 0xC000021B
	STATUS_NO_BROWSER_SERVERS_FOUND                                                 = 0xC000021C
	STATUS_VDM_HARD_ERROR                                                           = 0xC000021D
	STATUS_DRIVER_CANCEL_TIMEOUT                                                    = 0xC000021E
	STATUS_REPLY_MESSAGE_MISMATCH                                                   = 0xC000021F
	STATUS_MAPPED_ALIGNMENT                                                         = 0xC0000220
	STATUS_IMAGE_CHECKSUM_MISMATCH                                                  = 0xC0000221
	STATUS_LOST_WRITEBEHIND_DATA                                                    = 0xC0000222
	STATUS_CLIENT_SERVER_PARAMETERS_INVALID                                         = 0xC0000223
	STATUS_PASSWORD_MUST_CHANGE                                                     = 0xC0000224
	STATUS_NOT_FOUND                                                                = 0xC0000225
	STATUS_NOT_TINY_STREAM                                                          = 0xC0000226
	STATUS_RECOVERY_FAILURE                                                         = 0xC0000227
	STATUS_STACK_OVERFLOW_READ                                                      = 0xC0000228
	STATUS_FAIL_CHECK                                                               = 0xC0000229
	STATUS_DUPLICATE_OBJECTID                                                       = 0xC000022A
	STATUS_OBJECTID_EXISTS                                                          = 0xC000022B
	STATUS_CONVERT_TO_LARGE                                                         = 0xC000022C
	STATUS_RETRY                                                                    = 0xC000022D
	STATUS_FOUND_OUT_OF_SCOPE                                                       = 0xC000022E
	STATUS_ALLOCATE_BUCKET                                                          = 0xC000022F
	STATUS_PROPSET_NOT_FOUND                                                        = 0xC0000230
	STATUS_MARSHALL_OVERFLOW                                                        = 0xC0000231
	STATUS_INVALID_VARIANT                                                          = 0xC0000232
	STATUS_DOMAIN_CONTROLLER_NOT_FOUND                                              = 0xC0000233
	STATUS_ACCOUNT_LOCKED_OUT                                                       = 0xC0000234
	STATUS_HANDLE_NOT_CLOSABLE                                                      = 0xC0000235
	STATUS_CONNECTION_REFUSED                                                       = 0xC0000236
	STATUS_GRACEFUL_DISCONNECT                                                      = 0xC0000237
	STATUS_ADDRESS_ALREADY_ASSOCIATED                                               = 0xC0000238
	STATUS_ADDRESS_NOT_ASSOCIATED                                                   = 0xC0000239
	STATUS_CONNECTION_INVALID                                                       = 0xC000023A
	STATUS_CONNECTION_ACTIVE                                                        = 0xC000023B
	STATUS_NETWORK_UNREACHABLE                                                      = 0xC000023C
	STATUS_HOST_UNREACHABLE                                                         = 0xC000023D
	STATUS_PROTOCOL_UNREACHABLE                                                     = 0xC000023E
	STATUS_PORT_UNREACHABLE                                                         = 0xC000023F
	STATUS_REQUEST_ABORTED                                                          = 0xC0000240
	STATUS_CONNECTION_ABORTED                                                       = 0xC0000241
	STATUS_BAD_COMPRESSION_BUFFER                                                   = 0xC0000242
	STATUS_USER_MAPPED_FILE                                                         = 0xC0000243
	STATUS_AUDIT_FAILED                                                             = 0xC0000244
	STATUS_TIMER_RESOLUTION_NOT_SET                                                 = 0xC0000245
	STATUS_CONNECTION_COUNT_LIMIT                                                   = 0xC0000246
	STATUS_LOGIN_TIME_RESTRICTION                                                   = 0xC0000247
	STATUS_LOGIN_WKSTA_RESTRICTION                                                  = 0xC0000248
	STATUS_IMAGE_MP_UP_MISMATCH                                                     = 0xC0000249
	STATUS_INSUFFICIENT_LOGON_INFO                                                  = 0xC0000250
	STATUS_BAD_DLL_ENTRYPOINT                                                       = 0xC0000251
	STATUS_BAD_SERVICE_ENTRYPOINT                                                   = 0xC0000252
	STATUS_LPC_REPLY_LOST                                                           = 0xC0000253
	STATUS_IP_ADDRESS_CONFLICT1                                                     = 0xC0000254
	STATUS_IP_ADDRESS_CONFLICT2                                                     = 0xC0000255
	STATUS_REGISTRY_QUOTA_LIMIT                                                     = 0xC0000256
	STATUS_PATH_NOT_COVERED                                                         = 0xC0000257
	STATUS_NO_CALLBACK_ACTIVE                                                       = 0xC0000258
	STATUS_LICENSE_QUOTA_EXCEEDED                                                   = 0xC0000259
	STATUS_PWD_TOO_SHORT                                                            = 0xC000025A
	STATUS_PWD_TOO_RECENT                                                           = 0xC000025B
	STATUS_PWD_HISTORY_CONFLICT                                                     = 0xC000025C
	STATUS_PLUGPLAY_NO_DEVICE                                                       = 0xC000025E
	STATUS_UNSUPPORTED_COMPRESSION                                                  = 0xC000025F
	STATUS_INVALID_HW_PROFILE                                                       = 0xC0000260
	STATUS_INVALID_PLUGPLAY_DEVICE_PATH                                             = 0xC0000261
	STATUS_DRIVER_ORDINAL_NOT_FOUND                                                 = 0xC0000262
	STATUS_DRIVER_ENTRYPOINT_NOT_FOUND                                              = 0xC0000263
	STATUS_RESOURCE_NOT_OWNED                                                       = 0xC0000264
	STATUS_TOO_MANY_LINKS                                                           = 0xC0000265
	STATUS_QUOTA_LIST_INCONSISTENT                                                  = 0xC0000266
	STATUS_FILE_IS_OFFLINE                                                          = 0xC0000267
	STATUS_EVALUATION_EXPIRATION                                                    = 0xC0000268
	STATUS_ILLEGAL_DLL_RELOCATION                                                   = 0xC0000269
	STATUS_LICENSE_VIOLATION                                                        = 0xC000026A
	STATUS_DLL_INIT_FAILED_LOGOFF                                                   = 0xC000026B
	STATUS_DRIVER_UNABLE_TO_LOAD                                                    = 0xC000026C
	STATUS_DFS_UNAVAILABLE                                                          = 0xC000026D
	STATUS_VOLUME_DISMOUNTED                                                        = 0xC000026E
	STATUS_WX86_INTERNAL_ERROR                                                      = 0xC000026F
	STATUS_WX86_FLOAT_STACK_CHECK                                                   = 0xC0000270
	STATUS_VALIDATE_CONTINUE                                                        = 0xC0000271
	STATUS_NO_MATCH                                                                 = 0xC0000272
	STATUS_NO_MORE_MATCHES                                                          = 0xC0000273
	STATUS_NOT_A_REPARSE_POINT                                                      = 0xC0000275
	STATUS_IO_REPARSE_TAG_INVALID                                                   = 0xC0000276
	STATUS_IO_REPARSE_TAG_MISMATCH                                                  = 0xC0000277
	STATUS_IO_REPARSE_DATA_INVALID                                                  = 0xC0000278
	STATUS_IO_REPARSE_TAG_NOT_HANDLED                                               = 0xC0000279
	STATUS_REPARSE_POINT_NOT_RESOLVED                                               = 0xC0000280
	STATUS_DIRECTORY_IS_A_REPARSE_POINT                                             = 0xC0000281
	STATUS_RANGE_LIST_CONFLICT                                                      = 0xC0000282
	STATUS_SOURCE_ELEMENT_EMPTY                                                     = 0xC0000283
	STATUS_DESTINATION_ELEMENT_FULL                                                 = 0xC0000284
	STATUS_ILLEGAL_ELEMENT_ADDRESS                                                  = 0xC0000285
	STATUS_MAGAZINE_NOT_PRESENT                                                     = 0xC0000286
	STATUS_REINITIALIZATION_NEEDED                                                  = 0xC0000287
	STATUS_DEVICE_REQUIRES_CLEANING                                                 = 0x80000288
	STATUS_DEVICE_DOOR_OPEN                                                         = 0x80000289
	STATUS_ENCRYPTION_FAILED                                                        = 0xC000028A
	STATUS_DECRYPTION_FAILED                                                        = 0xC000028B
	STATUS_RANGE_NOT_FOUND                                                          = 0xC000028C
	STATUS_NO_RECOVERY_POLICY                                                       = 0xC000028D
	STATUS_NO_EFS                                                                   = 0xC000028E
	STATUS_WRONG_EFS                                                                = 0xC000028F
	STATUS_NO_USER_KEYS                                                             = 0xC0000290
	STATUS_FILE_NOT_ENCRYPTED                                                       = 0xC0000291
	STATUS_NOT_EXPORT_FORMAT                                                        = 0xC0000292
	STATUS_FILE_ENCRYPTED                                                           = 0xC0000293
	STATUS_WAKE_SYSTEM                                                              = 0x40000294
	STATUS_WMI_GUID_NOT_FOUND                                                       = 0xC0000295
	STATUS_WMI_INSTANCE_NOT_FOUND                                                   = 0xC0000296
	STATUS_WMI_ITEMID_NOT_FOUND                                                     = 0xC0000297
	STATUS_WMI_TRY_AGAIN                                                            = 0xC0000298
	STATUS_SHARED_POLICY                                                            = 0xC0000299
	STATUS_POLICY_OBJECT_NOT_FOUND                                                  = 0xC000029A
	STATUS_POLICY_ONLY_IN_DS                                                        = 0xC000029B
	STATUS_VOLUME_NOT_UPGRADED                                                      = 0xC000029C
	STATUS_REMOTE_STORAGE_NOT_ACTIVE                                                = 0xC000029D
	STATUS_REMOTE_STORAGE_MEDIA_ERROR                                               = 0xC000029E
	STATUS_NO_TRACKING_SERVICE                                                      = 0xC000029F
	STATUS_SERVER_SID_MISMATCH                                                      = 0xC00002A0
	STATUS_DS_NO_ATTRIBUTE_OR_VALUE                                                 = 0xC00002A1
	STATUS_DS_INVALID_ATTRIBUTE_SYNTAX                                              = 0xC00002A2
	STATUS_DS_ATTRIBUTE_TYPE_UNDEFINED                                              = 0xC00002A3
	STATUS_DS_ATTRIBUTE_OR_VALUE_EXISTS                                             = 0xC00002A4
	STATUS_DS_BUSY                                                                  = 0xC00002A5
	STATUS_DS_UNAVAILABLE                                                           = 0xC00002A6
	STATUS_DS_NO_RIDS_ALLOCATED                                                     = 0xC00002A7
	STATUS_DS_NO_MORE_RIDS                                                          = 0xC00002A8
	STATUS_DS_INCORRECT_ROLE_OWNER                                                  = 0xC00002A9
	STATUS_DS_RIDMGR_INIT_ERROR                                                     = 0xC00002AA
	STATUS_DS_OBJ_CLASS_VIOLATION                                                   = 0xC00002AB
	STATUS_DS_CANT_ON_NON_LEAF                                                      = 0xC00002AC
	STATUS_DS_CANT_ON_RDN                                                           = 0xC00002AD
	STATUS_DS_CANT_MOD_OBJ_CLASS                                                    = 0xC00002AE
	STATUS_DS_CROSS_DOM_MOVE_FAILED                                                 = 0xC00002AF
	STATUS_DS_GC_NOT_AVAILABLE                                                      = 0xC00002B0
	STATUS_DIRECTORY_SERVICE_REQUIRED                                               = 0xC00002B1
	STATUS_REPARSE_ATTRIBUTE_CONFLICT                                               = 0xC00002B2
	STATUS_CANT_ENABLE_DENY_ONLY                                                    = 0xC00002B3
	STATUS_FLOAT_MULTIPLE_FAULTS                                                    = 0xC00002B4
	STATUS_FLOAT_MULTIPLE_TRAPS                                                     = 0xC00002B5
	STATUS_DEVICE_REMOVED                                                           = 0xC00002B6
	STATUS_JOURNAL_DELETE_IN_PROGRESS                                               = 0xC00002B7
	STATUS_JOURNAL_NOT_ACTIVE                                                       = 0xC00002B8
	STATUS_NOINTERFACE                                                              = 0xC00002B9
	STATUS_DS_ADMIN_LIMIT_EXCEEDED                                                  = 0xC00002C1
	STATUS_DRIVER_FAILED_SLEEP                                                      = 0xC00002C2
	STATUS_MUTUAL_AUTHENTICATION_FAILED                                             = 0xC00002C3
	STATUS_CORRUPT_SYSTEM_FILE                                                      = 0xC00002C4
	STATUS_DATATYPE_MISALIGNMENT_ERROR                                              = 0xC00002C5
	STATUS_WMI_READ_ONLY                                                            = 0xC00002C6
	STATUS_WMI_SET_FAILURE                                                          = 0xC00002C7
	STATUS_COMMITMENT_MINIMUM                                                       = 0xC00002C8
	STATUS_REG_NAT_CONSUMPTION                                                      = 0xC00002C9
	STATUS_TRANSPORT_FULL                                                           = 0xC00002CA
	STATUS_DS_SAM_INIT_FAILURE                                                      = 0xC00002CB
	STATUS_ONLY_IF_CONNECTED                                                        = 0xC00002CC
	STATUS_DS_SENSITIVE_GROUP_VIOLATION                                             = 0xC00002CD
	STATUS_PNP_RESTART_ENUMERATION                                                  = 0xC00002CE
	STATUS_JOURNAL_ENTRY_DELETED                                                    = 0xC00002CF
	STATUS_DS_CANT_MOD_PRIMARYGROUPID                                               = 0xC00002D0
	STATUS_SYSTEM_IMAGE_BAD_SIGNATURE                                               = 0xC00002D1
	STATUS_PNP_REBOOT_REQUIRED                                                      = 0xC00002D2
	STATUS_POWER_STATE_INVALID                                                      = 0xC00002D3
	STATUS_DS_INVALID_GROUP_TYPE                                                    = 0xC00002D4
	STATUS_DS_NO_NEST_GLOBALGROUP_IN_MIXEDDOMAIN                                    = 0xC00002D5
	STATUS_DS_NO_NEST_LOCALGROUP_IN_MIXEDDOMAIN                                     = 0xC00002D6
	STATUS_DS_GLOBAL_CANT_HAVE_LOCAL_MEMBER                                         = 0xC00002D7
	STATUS_DS_GLOBAL_CANT_HAVE_UNIVERSAL_MEMBER                                     = 0xC00002D8
	STATUS_DS_UNIVERSAL_CANT_HAVE_LOCAL_MEMBER                                      = 0xC00002D9
	STATUS_DS_GLOBAL_CANT_HAVE_CROSSDOMAIN_MEMBER                                   = 0xC00002DA
	STATUS_DS_LOCAL_CANT_HAVE_CROSSDOMAIN_LOCAL_MEMBER                              = 0xC00002DB
	STATUS_DS_HAVE_PRIMARY_MEMBERS                                                  = 0xC00002DC
	STATUS_WMI_NOT_SUPPORTED                                                        = 0xC00002DD
	STATUS_INSUFFICIENT_POWER                                                       = 0xC00002DE
	STATUS_SAM_NEED_BOOTKEY_PASSWORD                                                = 0xC00002DF
	STATUS_SAM_NEED_BOOTKEY_FLOPPY                                                  = 0xC00002E0
	STATUS_DS_CANT_START                                                            = 0xC00002E1
	STATUS_DS_INIT_FAILURE                                                          = 0xC00002E2
	STATUS_SAM_INIT_FAILURE                                                         = 0xC00002E3
	STATUS_DS_GC_REQUIRED                                                           = 0xC00002E4
	STATUS_DS_LOCAL_MEMBER_OF_LOCAL_ONLY                                            = 0xC00002E5
	STATUS_DS_NO_FPO_IN_UNIVERSAL_GROUPS                                            = 0xC00002E6
	STATUS_DS_MACHINE_ACCOUNT_QUOTA_EXCEEDED                                        = 0xC00002E7
	STATUS_MULTIPLE_FAULT_VIOLATION                                                 = 0xC00002E8
	STATUS_CURRENT_DOMAIN_NOT_ALLOWED                                               = 0xC00002E9
	STATUS_CANNOT_MAKE                                                              = 0xC00002EA
	STATUS_SYSTEM_SHUTDOWN                                                          = 0xC00002EB
	STATUS_DS_INIT_FAILURE_CONSOLE                                                  = 0xC00002EC
	STATUS_DS_SAM_INIT_FAILURE_CONSOLE                                              = 0xC00002ED
	STATUS_UNFINISHED_CONTEXT_DELETED                                               = 0xC00002EE
	STATUS_NO_TGT_REPLY                                                             = 0xC00002EF
	STATUS_OBJECTID_NOT_FOUND                                                       = 0xC00002F0
	STATUS_NO_IP_ADDRESSES                                                          = 0xC00002F1
	STATUS_WRONG_CREDENTIAL_HANDLE                                                  = 0xC00002F2
	STATUS_CRYPTO_SYSTEM_INVALID                                                    = 0xC00002F3
	STATUS_MAX_REFERRALS_EXCEEDED                                                   = 0xC00002F4
	STATUS_MUST_BE_KDC                                                              = 0xC00002F5
	STATUS_STRONG_CRYPTO_NOT_SUPPORTED                                              = 0xC00002F6
	STATUS_TOO_MANY_PRINCIPALS                                                      = 0xC00002F7
	STATUS_NO_PA_DATA                                                               = 0xC00002F8
	STATUS_PKINIT_NAME_MISMATCH                                                     = 0xC00002F9
	STATUS_SMARTCARD_LOGON_REQUIRED                                                 = 0xC00002FA
	STATUS_KDC_INVALID_REQUEST                                                      = 0xC00002FB
	STATUS_KDC_UNABLE_TO_REFER                                                      = 0xC00002FC
	STATUS_KDC_UNKNOWN_ETYPE                                                        = 0xC00002FD
	STATUS_SHUTDOWN_IN_PROGRESS                                                     = 0xC00002FE
	STATUS_SERVER_SHUTDOWN_IN_PROGRESS                                              = 0xC00002FF
	STATUS_NOT_SUPPORTED_ON_SBS                                                     = 0xC0000300
	STATUS_WMI_GUID_DISCONNECTED                                                    = 0xC0000301
	STATUS_WMI_ALREADY_DISABLED                                                     = 0xC0000302
	STATUS_WMI_ALREADY_ENABLED                                                      = 0xC0000303
	STATUS_MFT_TOO_FRAGMENTED                                                       = 0xC0000304
	STATUS_COPY_PROTECTION_FAILURE                                                  = 0xC0000305
	STATUS_CSS_AUTHENTICATION_FAILURE                                               = 0xC0000306
	STATUS_CSS_KEY_NOT_PRESENT                                                      = 0xC0000307
	STATUS_CSS_KEY_NOT_ESTABLISHED                                                  = 0xC0000308
	STATUS_CSS_SCRAMBLED_SECTOR                                                     = 0xC0000309
	STATUS_CSS_REGION_MISMATCH                                                      = 0xC000030A
	STATUS_CSS_RESETS_EXHAUSTED                                                     = 0xC000030B
	STATUS_PKINIT_FAILURE                                                           = 0xC0000320
	STATUS_SMARTCARD_SUBSYSTEM_FAILURE                                              = 0xC0000321
	STATUS_NO_KERB_KEY                                                              = 0xC0000322
	STATUS_HOST_DOWN                                                                = 0xC0000350
	STATUS_UNSUPPORTED_PREAUTH                                                      = 0xC0000351
	STATUS_EFS_ALG_BLOB_TOO_BIG                                                     = 0xC0000352
	STATUS_PORT_NOT_SET                                                             = 0xC0000353
	STATUS_DEBUGGER_INACTIVE                                                        = 0xC0000354
	STATUS_DS_VERSION_CHECK_FAILURE                                                 = 0xC0000355
	STATUS_AUDITING_DISABLED                                                        = 0xC0000356
	STATUS_PRENT4_MACHINE_ACCOUNT                                                   = 0xC0000357
	STATUS_DS_AG_CANT_HAVE_UNIVERSAL_MEMBER                                         = 0xC0000358
	STATUS_INVALID_IMAGE_WIN_32                                                     = 0xC0000359
	STATUS_INVALID_IMAGE_WIN_64                                                     = 0xC000035A
	STATUS_BAD_BINDINGS                                                             = 0xC000035B
	STATUS_NETWORK_SESSION_EXPIRED                                                  = 0xC000035C
	STATUS_APPHELP_BLOCK                                                            = 0xC000035D
	STATUS_ALL_SIDS_FILTERED                                                        = 0xC000035E
	STATUS_NOT_SAFE_MODE_DRIVER                                                     = 0xC000035F
	STATUS_ACCESS_DISABLED_BY_POLICY_DEFAULT                                        = 0xC0000361
	STATUS_ACCESS_DISABLED_BY_POLICY_PATH                                           = 0xC0000362
	STATUS_ACCESS_DISABLED_BY_POLICY_PUBLISHER                                      = 0xC0000363
	STATUS_ACCESS_DISABLED_BY_POLICY_OTHER                                          = 0xC0000364
	STATUS_FAILED_DRIVER_ENTRY                                                      = 0xC0000365
	STATUS_DEVICE_ENUMERATION_ERROR                                                 = 0xC0000366
	STATUS_MOUNT_POINT_NOT_RESOLVED                                                 = 0xC0000368
	STATUS_INVALID_DEVICE_OBJECT_PARAMETER                                          = 0xC0000369
	STATUS_MCA_OCCURED                                                              = 0xC000036A
	STATUS_DRIVER_BLOCKED_CRITICAL                                                  = 0xC000036B
	STATUS_DRIVER_BLOCKED                                                           = 0xC000036C
	STATUS_DRIVER_DATABASE_ERROR                                                    = 0xC000036D
	STATUS_SYSTEM_HIVE_TOO_LARGE                                                    = 0xC000036E
	STATUS_INVALID_IMPORT_OF_NON_DLL                                                = 0xC000036F
	STATUS_DS_SHUTTING_DOWN                                                         = 0x40000370
	STATUS_NO_SECRETS                                                               = 0xC0000371
	STATUS_ACCESS_DISABLED_NO_SAFER_UI_BY_POLICY                                    = 0xC0000372
	STATUS_FAILED_STACK_SWITCH                                                      = 0xC0000373
	STATUS_HEAP_CORRUPTION                                                          = 0xC0000374
	STATUS_SMARTCARD_WRONG_PIN                                                      = 0xC0000380
	STATUS_SMARTCARD_CARD_BLOCKED                                                   = 0xC0000381
	STATUS_SMARTCARD_CARD_NOT_AUTHENTICATED                                         = 0xC0000382
	STATUS_SMARTCARD_NO_CARD                                                        = 0xC0000383
	STATUS_SMARTCARD_NO_KEY_CONTAINER                                               = 0xC0000384
	STATUS_SMARTCARD_NO_CERTIFICATE                                                 = 0xC0000385
	STATUS_SMARTCARD_NO_KEYSET                                                      = 0xC0000386
	STATUS_SMARTCARD_IO_ERROR                                                       = 0xC0000387
	STATUS_DOWNGRADE_DETECTED                                                       = 0xC0000388
	STATUS_SMARTCARD_CERT_REVOKED                                                   = 0xC0000389
	STATUS_ISSUING_CA_UNTRUSTED                                                     = 0xC000038A
	STATUS_REVOCATION_OFFLINE_C                                                     = 0xC000038B
	STATUS_PKINIT_CLIENT_FAILURE                                                    = 0xC000038C
	STATUS_SMARTCARD_CERT_EXPIRED                                                   = 0xC000038D
	STATUS_DRIVER_FAILED_PRIOR_UNLOAD                                               = 0xC000038E
	STATUS_SMARTCARD_SILENT_CONTEXT                                                 = 0xC000038F
	STATUS_PER_USER_TRUST_QUOTA_EXCEEDED                                            = 0xC0000401
	STATUS_ALL_USER_TRUST_QUOTA_EXCEEDED                                            = 0xC0000402
	STATUS_USER_DELETE_TRUST_QUOTA_EXCEEDED                                         = 0xC0000403
	STATUS_DS_NAME_NOT_UNIQUE                                                       = 0xC0000404
	STATUS_DS_DUPLICATE_ID_FOUND                                                    = 0xC0000405
	STATUS_DS_GROUP_CONVERSION_ERROR                                                = 0xC0000406
	STATUS_VOLSNAP_PREPARE_HIBERNATE                                                = 0xC0000407
	STATUS_USER2USER_REQUIRED                                                       = 0xC0000408
	STATUS_STACK_BUFFER_OVERRUN                                                     = 0xC0000409
	STATUS_NO_S4U_PROT_SUPPORT                                                      = 0xC000040A
	STATUS_CROSSREALM_DELEGATION_FAILURE                                            = 0xC000040B
	STATUS_REVOCATION_OFFLINE_KDC                                                   = 0xC000040C
	STATUS_ISSUING_CA_UNTRUSTED_KDC                                                 = 0xC000040D
	STATUS_KDC_CERT_EXPIRED                                                         = 0xC000040E
	STATUS_KDC_CERT_REVOKED                                                         = 0xC000040F
	STATUS_PARAMETER_QUOTA_EXCEEDED                                                 = 0xC0000410
	STATUS_HIBERNATION_FAILURE                                                      = 0xC0000411
	STATUS_DELAY_LOAD_FAILED                                                        = 0xC0000412
	STATUS_AUTHENTICATION_FIREWALL_FAILED                                           = 0xC0000413
	STATUS_VDM_DISALLOWED                                                           = 0xC0000414
	STATUS_HUNG_DISPLAY_DRIVER_THREAD                                               = 0xC0000415
	STATUS_INSUFFICIENT_RESOURCE_FOR_SPECIFIED_SHARED_SECTION_SIZE                  = 0xC0000416
	STATUS_INVALID_CRUNTIME_PARAMETER                                               = 0xC0000417
	STATUS_NTLM_BLOCKED                                                             = 0xC0000418
	STATUS_DS_SRC_SID_EXISTS_IN_FOREST                                              = 0xC0000419
	STATUS_DS_DOMAIN_NAME_EXISTS_IN_FOREST                                          = 0xC000041A
	STATUS_DS_FLAT_NAME_EXISTS_IN_FOREST                                            = 0xC000041B
	STATUS_INVALID_USER_PRINCIPAL_NAME                                              = 0xC000041C
	STATUS_FATAL_USER_CALLBACK_EXCEPTION                                            = 0xC000041D
	STATUS_ASSERTION_FAILURE                                                        = 0xC0000420
	STATUS_VERIFIER_STOP                                                            = 0xC0000421
	STATUS_CALLBACK_POP_STACK                                                       = 0xC0000423
	STATUS_INCOMPATIBLE_DRIVER_BLOCKED                                              = 0xC0000424
	STATUS_HIVE_UNLOADED                                                            = 0xC0000425
	STATUS_COMPRESSION_DISABLED                                                     = 0xC0000426
	STATUS_FILE_SYSTEM_LIMITATION                                                   = 0xC0000427
	STATUS_INVALID_IMAGE_HASH                                                       = 0xC0000428
	STATUS_NOT_CAPABLE                                                              = 0xC0000429
	STATUS_REQUEST_OUT_OF_SEQUENCE                                                  = 0xC000042A
	STATUS_IMPLEMENTATION_LIMIT                                                     = 0xC000042B
	STATUS_ELEVATION_REQUIRED                                                       = 0xC000042C
	STATUS_NO_SECURITY_CONTEXT                                                      = 0xC000042D
	STATUS_PKU2U_CERT_FAILURE                                                       = 0xC000042F
	STATUS_BEYOND_VDL                                                               = 0xC0000432
	STATUS_ENCOUNTERED_WRITE_IN_PROGRESS                                            = 0xC0000433
	STATUS_PTE_CHANGED                                                              = 0xC0000434
	STATUS_PURGE_FAILED                                                             = 0xC0000435
	STATUS_CRED_REQUIRES_CONFIRMATION                                               = 0xC0000440
	STATUS_CS_ENCRYPTION_INVALID_SERVER_RESPONSE                                    = 0xC0000441
	STATUS_CS_ENCRYPTION_UNSUPPORTED_SERVER                                         = 0xC0000442
	STATUS_CS_ENCRYPTION_EXISTING_ENCRYPTED_FILE                                    = 0xC0000443
	STATUS_CS_ENCRYPTION_NEW_ENCRYPTED_FILE                                         = 0xC0000444
	STATUS_CS_ENCRYPTION_FILE_NOT_CSE                                               = 0xC0000445
	STATUS_INVALID_LABEL                                                            = 0xC0000446
	STATUS_DRIVER_PROCESS_TERMINATED                                                = 0xC0000450
	STATUS_AMBIGUOUS_SYSTEM_DEVICE                                                  = 0xC0000451
	STATUS_SYSTEM_DEVICE_NOT_FOUND                                                  = 0xC0000452
	STATUS_RESTART_BOOT_APPLICATION                                                 = 0xC0000453
	STATUS_INSUFFICIENT_NVRAM_RESOURCES                                             = 0xC0000454
	STATUS_INVALID_TASK_NAME                                                        = 0xC0000500
	STATUS_INVALID_TASK_INDEX                                                       = 0xC0000501
	STATUS_THREAD_ALREADY_IN_TASK                                                   = 0xC0000502
	STATUS_CALLBACK_BYPASS                                                          = 0xC0000503
	STATUS_FAIL_FAST_EXCEPTION                                                      = 0xC0000602
	STATUS_IMAGE_CERT_REVOKED                                                       = 0xC0000603
	STATUS_PORT_CLOSED                                                              = 0xC0000700
	STATUS_MESSAGE_LOST                                                             = 0xC0000701
	STATUS_INVALID_MESSAGE                                                          = 0xC0000702
	STATUS_REQUEST_CANCELED                                                         = 0xC0000703
	STATUS_RECURSIVE_DISPATCH                                                       = 0xC0000704
	STATUS_LPC_RECEIVE_BUFFER_EXPECTED                                              = 0xC0000705
	STATUS_LPC_INVALID_CONNECTION_USAGE                                             = 0xC0000706
	STATUS_LPC_REQUESTS_NOT_ALLOWED                                                 = 0xC0000707
	STATUS_RESOURCE_IN_USE                                                          = 0xC0000708
	STATUS_HARDWARE_MEMORY_ERROR                                                    = 0xC0000709
	STATUS_THREADPOOL_HANDLE_EXCEPTION                                              = 0xC000070A
	STATUS_THREADPOOL_SET_EVENT_ON_COMPLETION_FAILED                                = 0xC000070B
	STATUS_THREADPOOL_RELEASE_SEMAPHORE_ON_COMPLETION_FAILED                        = 0xC000070C
	STATUS_THREADPOOL_RELEASE_MUTEX_ON_COMPLETION_FAILED                            = 0xC000070D
	STATUS_THREADPOOL_FREE_LIBRARY_ON_COMPLETION_FAILED                             = 0xC000070E
	STATUS_THREADPOOL_RELEASED_DURING_OPERATION                                     = 0xC000070F
	STATUS_CALLBACK_RETURNED_WHILE_IMPERSONATING                                    = 0xC0000710
	STATUS_APC_RETURNED_WHILE_IMPERSONATING                                         = 0xC0000711
	STATUS_PROCESS_IS_PROTECTED                                                     = 0xC0000712
	STATUS_MCA_EXCEPTION                                                            = 0xC0000713
	STATUS_CERTIFICATE_MAPPING_NOT_UNIQUE                                           = 0xC0000714
	STATUS_SYMLINK_CLASS_DISABLED                                                   = 0xC0000715
	STATUS_INVALID_IDN_NORMALIZATION                                                = 0xC0000716
	STATUS_NO_UNICODE_TRANSLATION                                                   = 0xC0000717
	STATUS_ALREADY_REGISTERED                                                       = 0xC0000718
	STATUS_CONTEXT_MISMATCH                                                         = 0xC0000719
	STATUS_PORT_ALREADY_HAS_COMPLETION_LIST                                         = 0xC000071A
	STATUS_CALLBACK_RETURNED_THREAD_PRIORITY                                        = 0xC000071B
	STATUS_INVALID_THREAD                                                           = 0xC000071C
	STATUS_CALLBACK_RETURNED_TRANSACTION                                            = 0xC000071D
	STATUS_CALLBACK_RETURNED_LDR_LOCK                                               = 0xC000071E
	STATUS_CALLBACK_RETURNED_LANG                                                   = 0xC000071F
	STATUS_CALLBACK_RETURNED_PRI_BACK                                               = 0xC0000720
	STATUS_CALLBACK_RETURNED_THREAD_AFFINITY                                        = 0xC0000721
	STATUS_DISK_REPAIR_DISABLED                                                     = 0xC0000800
	STATUS_DS_DOMAIN_RENAME_IN_PROGRESS                                             = 0xC0000801
	STATUS_DISK_QUOTA_EXCEEDED                                                      = 0xC0000802
	STATUS_DATA_LOST_REPAIR                                                         = 0x80000803
	STATUS_CONTENT_BLOCKED                                                          = 0xC0000804
	STATUS_BAD_CLUSTERS                                                             = 0xC0000805
	STATUS_VOLUME_DIRTY                                                             = 0xC0000806
	STATUS_FILE_CHECKED_OUT                                                         = 0xC0000901
	STATUS_CHECKOUT_REQUIRED                                                        = 0xC0000902
	STATUS_BAD_FILE_TYPE                                                            = 0xC0000903
	STATUS_FILE_TOO_LARGE                                                           = 0xC0000904
	STATUS_FORMS_AUTH_REQUIRED                                                      = 0xC0000905
	STATUS_VIRUS_INFECTED                                                           = 0xC0000906
	STATUS_VIRUS_DELETED                                                            = 0xC0000907
	STATUS_BAD_MCFG_TABLE                                                           = 0xC0000908
	STATUS_CANNOT_BREAK_OPLOCK                                                      = 0xC0000909
	STATUS_WOW_ASSERTION                                                            = 0xC0009898
	STATUS_INVALID_SIGNATURE                                                        = 0xC000A000
	STATUS_HMAC_NOT_SUPPORTED                                                       = 0xC000A001
	STATUS_AUTH_TAG_MISMATCH                                                        = 0xC000A002
	STATUS_IPSEC_QUEUE_OVERFLOW                                                     = 0xC000A010
	STATUS_ND_QUEUE_OVERFLOW                                                        = 0xC000A011
	STATUS_HOPLIMIT_EXCEEDED                                                        = 0xC000A012
	STATUS_PROTOCOL_NOT_SUPPORTED                                                   = 0xC000A013
	STATUS_FASTPATH_REJECTED                                                        = 0xC000A014
	STATUS_LOST_WRITEBEHIND_DATA_NETWORK_DISCONNECTED                               = 0xC000A080
	STATUS_LOST_WRITEBEHIND_DATA_NETWORK_SERVER_ERROR                               = 0xC000A081
	STATUS_LOST_WRITEBEHIND_DATA_LOCAL_DISK_ERROR                                   = 0xC000A082
	STATUS_XML_PARSE_ERROR                                                          = 0xC000A083
	STATUS_XMLDSIG_ERROR                                                            = 0xC000A084
	STATUS_WRONG_COMPARTMENT                                                        = 0xC000A085
	STATUS_AUTHIP_FAILURE                                                           = 0xC000A086
	STATUS_DS_OID_MAPPED_GROUP_CANT_HAVE_MEMBERS                                    = 0xC000A087
	STATUS_DS_OID_NOT_FOUND                                                         = 0xC000A088
	STATUS_HASH_NOT_SUPPORTED                                                       = 0xC000A100
	STATUS_HASH_NOT_PRESENT                                                         = 0xC000A101
	DBG_NO_STATE_CHANGE                                                             = 0xC0010001
	DBG_APP_NOT_IDLE                                                                = 0xC0010002
	RPC_NT_INVALID_STRING_BINDING                                                   = 0xC0020001
	RPC_NT_WRONG_KIND_OF_BINDING                                                    = 0xC0020002
	RPC_NT_INVALID_BINDING                                                          = 0xC0020003
	RPC_NT_PROTSEQ_NOT_SUPPORTED                                                    = 0xC0020004
	RPC_NT_INVALID_RPC_PROTSEQ                                                      = 0xC0020005
	RPC_NT_INVALID_STRING_UUID                                                      = 0xC0020006
	RPC_NT_INVALID_ENDPOINT_FORMAT                                                  = 0xC0020007
	RPC_NT_INVALID_NET_ADDR                                                         = 0xC0020008
	RPC_NT_NO_ENDPOINT_FOUND                                                        = 0xC0020009
	RPC_NT_INVALID_TIMEOUT                                                          = 0xC002000A
	RPC_NT_OBJECT_NOT_FOUND                                                         = 0xC002000B
	RPC_NT_ALREADY_REGISTERED                                                       = 0xC002000C
	RPC_NT_TYPE_ALREADY_REGISTERED                                                  = 0xC002000D
	RPC_NT_ALREADY_LISTENING                                                        = 0xC002000E
	RPC_NT_NO_PROTSEQS_REGISTERED                                                   = 0xC002000F
	RPC_NT_NOT_LISTENING                                                            = 0xC0020010
	RPC_NT_UNKNOWN_MGR_TYPE                                                         = 0xC0020011
	RPC_NT_UNKNOWN_IF                                                               = 0xC0020012
	RPC_NT_NO_BINDINGS                                                              = 0xC0020013
	RPC_NT_NO_PROTSEQS                                                              = 0xC0020014
	RPC_NT_CANT_CREATE_ENDPOINT                                                     = 0xC0020015
	RPC_NT_OUT_OF_RESOURCES                                                         = 0xC0020016
	RPC_NT_SERVER_UNAVAILABLE                                                       = 0xC0020017
	RPC_NT_SERVER_TOO_BUSY                                                          = 0xC0020018
	RPC_NT_INVALID_NETWORK_OPTIONS                                                  = 0xC0020019
	RPC_NT_NO_CALL_ACTIVE                                                           = 0xC002001A
	RPC_NT_CALL_FAILED                                                              = 0xC002001B
	RPC_NT_CALL_FAILED_DNE                                                          = 0xC002001C
	RPC_NT_PROTOCOL_ERROR                                                           = 0xC002001D
	RPC_NT_UNSUPPORTED_TRANS_SYN                                                    = 0xC002001F
	RPC_NT_UNSUPPORTED_TYPE                                                         = 0xC0020021
	RPC_NT_INVALID_TAG                                                              = 0xC0020022
	RPC_NT_INVALID_BOUND                                                            = 0xC0020023
	RPC_NT_NO_ENTRY_NAME                                                            = 0xC0020024
	RPC_NT_INVALID_NAME_SYNTAX                                                      = 0xC0020025
	RPC_NT_UNSUPPORTED_NAME_SYNTAX                                                  = 0xC0020026
	RPC_NT_UUID_NO_ADDRESS                                                          = 0xC0020028
	RPC_NT_DUPLICATE_ENDPOINT                                                       = 0xC0020029
	RPC_NT_UNKNOWN_AUTHN_TYPE                                                       = 0xC002002A
	RPC_NT_MAX_CALLS_TOO_SMALL                                                      = 0xC002002B
	RPC_NT_STRING_TOO_LONG                                                          = 0xC002002C
	RPC_NT_PROTSEQ_NOT_FOUND                                                        = 0xC002002D
	RPC_NT_PROCNUM_OUT_OF_RANGE                                                     = 0xC002002E
	RPC_NT_BINDING_HAS_NO_AUTH                                                      = 0xC002002F
	RPC_NT_UNKNOWN_AUTHN_SERVICE                                                    = 0xC0020030
	RPC_NT_UNKNOWN_AUTHN_LEVEL                                                      = 0xC0020031
	RPC_NT_INVALID_AUTH_IDENTITY                                                    = 0xC0020032
	RPC_NT_UNKNOWN_AUTHZ_SERVICE                                                    = 0xC0020033
	EPT_NT_INVALID_ENTRY                                                            = 0xC0020034
	EPT_NT_CANT_PERFORM_OP                                                          = 0xC0020035
	EPT_NT_NOT_REGISTERED                                                           = 0xC0020036
	RPC_NT_NOTHING_TO_EXPORT                                                        = 0xC0020037
	RPC_NT_INCOMPLETE_NAME                                                          = 0xC0020038
	RPC_NT_INVALID_VERS_OPTION                                                      = 0xC0020039
	RPC_NT_NO_MORE_MEMBERS                                                          = 0xC002003A
	RPC_NT_NOT_ALL_OBJS_UNEXPORTED                                                  = 0xC002003B
	RPC_NT_INTERFACE_NOT_FOUND                                                      = 0xC002003C
	RPC_NT_ENTRY_ALREADY_EXISTS                                                     = 0xC002003D
	RPC_NT_ENTRY_NOT_FOUND                                                          = 0xC002003E
	RPC_NT_NAME_SERVICE_UNAVAILABLE                                                 = 0xC002003F
	RPC_NT_INVALID_NAF_ID                                                           = 0xC0020040
	RPC_NT_CANNOT_SUPPORT                                                           = 0xC0020041
	RPC_NT_NO_CONTEXT_AVAILABLE                                                     = 0xC0020042
	RPC_NT_INTERNAL_ERROR                                                           = 0xC0020043
	RPC_NT_ZERO_DIVIDE                                                              = 0xC0020044
	RPC_NT_ADDRESS_ERROR                                                            = 0xC0020045
	RPC_NT_FP_DIV_ZERO                                                              = 0xC0020046
	RPC_NT_FP_UNDERFLOW                                                             = 0xC0020047
	RPC_NT_FP_OVERFLOW                                                              = 0xC0020048
	RPC_NT_NO_MORE_ENTRIES                                                          = 0xC0030001
	RPC_NT_SS_CHAR_TRANS_OPEN_FAIL                                                  = 0xC0030002
	RPC_NT_SS_CHAR_TRANS_SHORT_FILE                                                 = 0xC0030003
	RPC_NT_SS_IN_NULL_CONTEXT                                                       = 0xC0030004
	RPC_NT_SS_CONTEXT_MISMATCH                                                      = 0xC0030005
	RPC_NT_SS_CONTEXT_DAMAGED                                                       = 0xC0030006
	RPC_NT_SS_HANDLES_MISMATCH                                                      = 0xC0030007
	RPC_NT_SS_CANNOT_GET_CALL_HANDLE                                                = 0xC0030008
	RPC_NT_NULL_REF_POINTER                                                         = 0xC0030009
	RPC_NT_ENUM_VALUE_OUT_OF_RANGE                                                  = 0xC003000A
	RPC_NT_BYTE_COUNT_TOO_SMALL                                                     = 0xC003000B
	RPC_NT_BAD_STUB_DATA                                                            = 0xC003000C
	RPC_NT_CALL_IN_PROGRESS                                                         = 0xC0020049
	RPC_NT_NO_MORE_BINDINGS                                                         = 0xC002004A
	RPC_NT_GROUP_MEMBER_NOT_FOUND                                                   = 0xC002004B
	EPT_NT_CANT_CREATE                                                              = 0xC002004C
	RPC_NT_INVALID_OBJECT                                                           = 0xC002004D
	RPC_NT_NO_INTERFACES                                                            = 0xC002004F
	RPC_NT_CALL_CANCELLED                                                           = 0xC0020050
	RPC_NT_BINDING_INCOMPLETE                                                       = 0xC0020051
	RPC_NT_COMM_FAILURE                                                             = 0xC0020052
	RPC_NT_UNSUPPORTED_AUTHN_LEVEL                                                  = 0xC0020053
	RPC_NT_NO_PRINC_NAME                                                            = 0xC0020054
	RPC_NT_NOT_RPC_ERROR                                                            = 0xC0020055
	RPC_NT_UUID_LOCAL_ONLY                                                          = 0x40020056
	RPC_NT_SEC_PKG_ERROR                                                            = 0xC0020057
	RPC_NT_NOT_CANCELLED                                                            = 0xC0020058
	RPC_NT_INVALID_ES_ACTION                                                        = 0xC0030059
	RPC_NT_WRONG_ES_VERSION                                                         = 0xC003005A
	RPC_NT_WRONG_STUB_VERSION                                                       = 0xC003005B
	RPC_NT_INVALID_PIPE_OBJECT                                                      = 0xC003005C
	RPC_NT_INVALID_PIPE_OPERATION                                                   = 0xC003005D
	RPC_NT_WRONG_PIPE_VERSION                                                       = 0xC003005E
	RPC_NT_PIPE_CLOSED                                                              = 0xC003005F
	RPC_NT_PIPE_DISCIPLINE_ERROR                                                    = 0xC0030060
	RPC_NT_PIPE_EMPTY                                                               = 0xC0030061
	RPC_NT_INVALID_ASYNC_HANDLE                                                     = 0xC0020062
	RPC_NT_INVALID_ASYNC_CALL                                                       = 0xC0020063
	RPC_NT_PROXY_ACCESS_DENIED                                                      = 0xC0020064
	RPC_NT_COOKIE_AUTH_FAILED                                                       = 0xC0020065
	RPC_NT_SEND_INCOMPLETE                                                          = 0x400200AF
	STATUS_ACPI_INVALID_OPCODE                                                      = 0xC0140001
	STATUS_ACPI_STACK_OVERFLOW                                                      = 0xC0140002
	STATUS_ACPI_ASSERT_FAILED                                                       = 0xC0140003
	STATUS_ACPI_INVALID_INDEX                                                       = 0xC0140004
	STATUS_ACPI_INVALID_ARGUMENT                                                    = 0xC0140005
	STATUS_ACPI_FATAL                                                               = 0xC0140006
	STATUS_ACPI_INVALID_SUPERNAME                                                   = 0xC0140007
	STATUS_ACPI_INVALID_ARGTYPE                                                     = 0xC0140008
	STATUS_ACPI_INVALID_OBJTYPE                                                     = 0xC0140009
	STATUS_ACPI_INVALID_TARGETTYPE                                                  = 0xC014000A
	STATUS_ACPI_INCORRECT_ARGUMENT_COUNT                                            = 0xC014000B
	STATUS_ACPI_ADDRESS_NOT_MAPPED                                                  = 0xC014000C
	STATUS_ACPI_INVALID_EVENTTYPE                                                   = 0xC014000D
	STATUS_ACPI_HANDLER_COLLISION                                                   = 0xC014000E
	STATUS_ACPI_INVALID_DATA                                                        = 0xC014000F
	STATUS_ACPI_INVALID_REGION                                                      = 0xC0140010
	STATUS_ACPI_INVALID_ACCESS_SIZE                                                 = 0xC0140011
	STATUS_ACPI_ACQUIRE_GLOBAL_LOCK                                                 = 0xC0140012
	STATUS_ACPI_ALREADY_INITIALIZED                                                 = 0xC0140013
	STATUS_ACPI_NOT_INITIALIZED                                                     = 0xC0140014
	STATUS_ACPI_INVALID_MUTEX_LEVEL                                                 = 0xC0140015
	STATUS_ACPI_MUTEX_NOT_OWNED                                                     = 0xC0140016
	STATUS_ACPI_MUTEX_NOT_OWNER                                                     = 0xC0140017
	STATUS_ACPI_RS_ACCESS                                                           = 0xC0140018
	STATUS_ACPI_INVALID_TABLE                                                       = 0xC0140019
	STATUS_ACPI_REG_HANDLER_FAILED                                                  = 0xC0140020
	STATUS_ACPI_POWER_REQUEST_FAILED                                                = 0xC0140021
	STATUS_CTX_WINSTATION_NAME_INVALID                                              = 0xC00A0001
	STATUS_CTX_INVALID_PD                                                           = 0xC00A0002
	STATUS_CTX_PD_NOT_FOUND                                                         = 0xC00A0003
	STATUS_CTX_CDM_CONNECT                                                          = 0x400A0004
	STATUS_CTX_CDM_DISCONNECT                                                       = 0x400A0005
	STATUS_CTX_CLOSE_PENDING                                                        = 0xC00A0006
	STATUS_CTX_NO_OUTBUF                                                            = 0xC00A0007
	STATUS_CTX_MODEM_INF_NOT_FOUND                                                  = 0xC00A0008
	STATUS_CTX_INVALID_MODEMNAME                                                    = 0xC00A0009
	STATUS_CTX_RESPONSE_ERROR                                                       = 0xC00A000A
	STATUS_CTX_MODEM_RESPONSE_TIMEOUT                                               = 0xC00A000B
	STATUS_CTX_MODEM_RESPONSE_NO_CARRIER                                            = 0xC00A000C
	STATUS_CTX_MODEM_RESPONSE_NO_DIALTONE                                           = 0xC00A000D
	STATUS_CTX_MODEM_RESPONSE_BUSY                                                  = 0xC00A000E
	STATUS_CTX_MODEM_RESPONSE_VOICE                                                 = 0xC00A000F
	STATUS_CTX_TD_ERROR                                                             = 0xC00A0010
	STATUS_CTX_LICENSE_CLIENT_INVALID                                               = 0xC00A0012
	STATUS_CTX_LICENSE_NOT_AVAILABLE                                                = 0xC00A0013
	STATUS_CTX_LICENSE_EXPIRED                                                      = 0xC00A0014
	STATUS_CTX_WINSTATION_NOT_FOUND                                                 = 0xC00A0015
	STATUS_CTX_WINSTATION_NAME_COLLISION                                            = 0xC00A0016
	STATUS_CTX_WINSTATION_BUSY                                                      = 0xC00A0017
	STATUS_CTX_BAD_VIDEO_MODE                                                       = 0xC00A0018
	STATUS_CTX_GRAPHICS_INVALID                                                     = 0xC00A0022
	STATUS_CTX_NOT_CONSOLE                                                          = 0xC00A0024
	STATUS_CTX_CLIENT_QUERY_TIMEOUT                                                 = 0xC00A0026
	STATUS_CTX_CONSOLE_DISCONNECT                                                   = 0xC00A0027
	STATUS_CTX_CONSOLE_CONNECT                                                      = 0xC00A0028
	STATUS_CTX_SHADOW_DENIED                                                        = 0xC00A002A
	STATUS_CTX_WINSTATION_ACCESS_DENIED                                             = 0xC00A002B
	STATUS_CTX_INVALID_WD                                                           = 0xC00A002E
	STATUS_CTX_WD_NOT_FOUND                                                         = 0xC00A002F
	STATUS_CTX_SHADOW_INVALID                                                       = 0xC00A0030
	STATUS_CTX_SHADOW_DISABLED                                                      = 0xC00A0031
	STATUS_RDP_PROTOCOL_ERROR                                                       = 0xC00A0032
	STATUS_CTX_CLIENT_LICENSE_NOT_SET                                               = 0xC00A0033
	STATUS_CTX_CLIENT_LICENSE_IN_USE                                                = 0xC00A0034
	STATUS_CTX_SHADOW_ENDED_BY_MODE_CHANGE                                          = 0xC00A0035
	STATUS_CTX_SHADOW_NOT_RUNNING                                                   = 0xC00A0036
	STATUS_CTX_LOGON_DISABLED                                                       = 0xC00A0037
	STATUS_CTX_SECURITY_LAYER_ERROR                                                 = 0xC00A0038
	STATUS_TS_INCOMPATIBLE_SESSIONS                                                 = 0xC00A0039
	STATUS_TS_VIDEO_SUBSYSTEM_ERROR                                                 = 0xC00A003A
	STATUS_PNP_BAD_MPS_TABLE                                                        = 0xC0040035
	STATUS_PNP_TRANSLATION_FAILED                                                   = 0xC0040036
	STATUS_PNP_IRQ_TRANSLATION_FAILED                                               = 0xC0040037
	STATUS_PNP_INVALID_ID                                                           = 0xC0040038
	STATUS_IO_REISSUE_AS_CACHED                                                     = 0xC0040039
	STATUS_MUI_FILE_NOT_FOUND                                                       = 0xC00B0001
	STATUS_MUI_INVALID_FILE                                                         = 0xC00B0002
	STATUS_MUI_INVALID_RC_CONFIG                                                    = 0xC00B0003
	STATUS_MUI_INVALID_LOCALE_NAME                                                  = 0xC00B0004
	STATUS_MUI_INVALID_ULTIMATEFALLBACK_NAME                                        = 0xC00B0005
	STATUS_MUI_FILE_NOT_LOADED                                                      = 0xC00B0006
	STATUS_RESOURCE_ENUM_USER_STOP                                                  = 0xC00B0007
	STATUS_FLT_NO_HANDLER_DEFINED                                                   = 0xC01C0001
	STATUS_FLT_CONTEXT_ALREADY_DEFINED                                              = 0xC01C0002
	STATUS_FLT_INVALID_ASYNCHRONOUS_REQUEST                                         = 0xC01C0003
	STATUS_FLT_DISALLOW_FAST_IO                                                     = 0xC01C0004
	STATUS_FLT_INVALID_NAME_REQUEST                                                 = 0xC01C0005
	STATUS_FLT_NOT_SAFE_TO_POST_OPERATION                                           = 0xC01C0006
	STATUS_FLT_NOT_INITIALIZED                                                      = 0xC01C0007
	STATUS_FLT_FILTER_NOT_READY                                                     = 0xC01C0008
	STATUS_FLT_POST_OPERATION_CLEANUP                                               = 0xC01C0009
	STATUS_FLT_INTERNAL_ERROR                                                       = 0xC01C000A
	STATUS_FLT_DELETING_OBJECT                                                      = 0xC01C000B
	STATUS_FLT_MUST_BE_NONPAGED_POOL                                                = 0xC01C000C
	STATUS_FLT_DUPLICATE_ENTRY                                                      = 0xC01C000D
	STATUS_FLT_CBDQ_DISABLED                                                        = 0xC01C000E
	STATUS_FLT_DO_NOT_ATTACH                                                        = 0xC01C000F
	STATUS_FLT_DO_NOT_DETACH                                                        = 0xC01C0010
	STATUS_FLT_INSTANCE_ALTITUDE_COLLISION                                          = 0xC01C0011
	STATUS_FLT_INSTANCE_NAME_COLLISION                                              = 0xC01C0012
	STATUS_FLT_FILTER_NOT_FOUND                                                     = 0xC01C0013
	STATUS_FLT_VOLUME_NOT_FOUND                                                     = 0xC01C0014
	STATUS_FLT_INSTANCE_NOT_FOUND                                                   = 0xC01C0015
	STATUS_FLT_CONTEXT_ALLOCATION_NOT_FOUND                                         = 0xC01C0016
	STATUS_FLT_INVALID_CONTEXT_REGISTRATION                                         = 0xC01C0017
	STATUS_FLT_NAME_CACHE_MISS                                                      = 0xC01C0018
	STATUS_FLT_NO_DEVICE_OBJECT                                                     = 0xC01C0019
	STATUS_FLT_VOLUME_ALREADY_MOUNTED                                               = 0xC01C001A
	STATUS_FLT_ALREADY_ENLISTED                                                     = 0xC01C001B
	STATUS_FLT_CONTEXT_ALREADY_LINKED                                               = 0xC01C001C
	STATUS_FLT_NO_WAITER_FOR_REPLY                                                  = 0xC01C0020
	STATUS_SXS_SECTION_NOT_FOUND                                                    = 0xC0150001
	STATUS_SXS_CANT_GEN_ACTCTX                                                      = 0xC0150002
	STATUS_SXS_INVALID_ACTCTXDATA_FORMAT                                            = 0xC0150003
	STATUS_SXS_ASSEMBLY_NOT_FOUND                                                   = 0xC0150004
	STATUS_SXS_MANIFEST_FORMAT_ERROR                                                = 0xC0150005
	STATUS_SXS_MANIFEST_PARSE_ERROR                                                 = 0xC0150006
	STATUS_SXS_ACTIVATION_CONTEXT_DISABLED                                          = 0xC0150007
	STATUS_SXS_KEY_NOT_FOUND                                                        = 0xC0150008
	STATUS_SXS_VERSION_CONFLICT                                                     = 0xC0150009
	STATUS_SXS_WRONG_SECTION_TYPE                                                   = 0xC015000A
	STATUS_SXS_THREAD_QUERIES_DISABLED                                              = 0xC015000B
	STATUS_SXS_ASSEMBLY_MISSING                                                     = 0xC015000C
	STATUS_SXS_RELEASE_ACTIVATION_CONTEXT                                           = 0x4015000D
	STATUS_SXS_PROCESS_DEFAULT_ALREADY_SET                                          = 0xC015000E
	STATUS_SXS_EARLY_DEACTIVATION                                                   = 0xC015000F
	STATUS_SXS_INVALID_DEACTIVATION                                                 = 0xC0150010
	STATUS_SXS_MULTIPLE_DEACTIVATION                                                = 0xC0150011
	STATUS_SXS_SYSTEM_DEFAULT_ACTIVATION_CONTEXT_EMPTY                              = 0xC0150012
	STATUS_SXS_PROCESS_TERMINATION_REQUESTED                                        = 0xC0150013
	STATUS_SXS_CORRUPT_ACTIVATION_STACK                                             = 0xC0150014
	STATUS_SXS_CORRUPTION                                                           = 0xC0150015
	STATUS_SXS_INVALID_IDENTITY_ATTRIBUTE_VALUE                                     = 0xC0150016
	STATUS_SXS_INVALID_IDENTITY_ATTRIBUTE_NAME                                      = 0xC0150017
	STATUS_SXS_IDENTITY_DUPLICATE_ATTRIBUTE                                         = 0xC0150018
	STATUS_SXS_IDENTITY_PARSE_ERROR                                                 = 0xC0150019
	STATUS_SXS_COMPONENT_STORE_CORRUPT                                              = 0xC015001A
	STATUS_SXS_FILE_HASH_MISMATCH                                                   = 0xC015001B
	STATUS_SXS_MANIFEST_IDENTITY_SAME_BUT_CONTENTS_DIFFERENT                        = 0xC015001C
	STATUS_SXS_IDENTITIES_DIFFERENT                                                 = 0xC015001D
	STATUS_SXS_ASSEMBLY_IS_NOT_A_DEPLOYMENT                                         = 0xC015001E
	STATUS_SXS_FILE_NOT_PART_OF_ASSEMBLY                                            = 0xC015001F
	STATUS_ADVANCED_INSTALLER_FAILED                                                = 0xC0150020
	STATUS_XML_ENCODING_MISMATCH                                                    = 0xC0150021
	STATUS_SXS_MANIFEST_TOO_BIG                                                     = 0xC0150022
	STATUS_SXS_SETTING_NOT_REGISTERED                                               = 0xC0150023
	STATUS_SXS_TRANSACTION_CLOSURE_INCOMPLETE                                       = 0xC0150024
	STATUS_SMI_PRIMITIVE_INSTALLER_FAILED                                           = 0xC0150025
	STATUS_GENERIC_COMMAND_FAILED                                                   = 0xC0150026
	STATUS_SXS_FILE_HASH_MISSING                                                    = 0xC0150027
	STATUS_CLUSTER_INVALID_NODE                                                     = 0xC0130001
	STATUS_CLUSTER_NODE_EXISTS                                                      = 0xC0130002
	STATUS_CLUSTER_JOIN_IN_PROGRESS                                                 = 0xC0130003
	STATUS_CLUSTER_NODE_NOT_FOUND                                                   = 0xC0130004
	STATUS_CLUSTER_LOCAL_NODE_NOT_FOUND                                             = 0xC0130005
	STATUS_CLUSTER_NETWORK_EXISTS                                                   = 0xC0130006
	STATUS_CLUSTER_NETWORK_NOT_FOUND                                                = 0xC0130007
	STATUS_CLUSTER_NETINTERFACE_EXISTS                                              = 0xC0130008
	STATUS_CLUSTER_NETINTERFACE_NOT_FOUND                                           = 0xC0130009
	STATUS_CLUSTER_INVALID_REQUEST                                                  = 0xC013000A
	STATUS_CLUSTER_INVALID_NETWORK_PROVIDER                                         = 0xC013000B
	STATUS_CLUSTER_NODE_DOWN                                                        = 0xC013000C
	STATUS_CLUSTER_NODE_UNREACHABLE                                                 = 0xC013000D
	STATUS_CLUSTER_NODE_NOT_MEMBER                                                  = 0xC013000E
	STATUS_CLUSTER_JOIN_NOT_IN_PROGRESS                                             = 0xC013000F
	STATUS_CLUSTER_INVALID_NETWORK                                                  = 0xC0130010
	STATUS_CLUSTER_NO_NET_ADAPTERS                                                  = 0xC0130011
	STATUS_CLUSTER_NODE_UP                                                          = 0xC0130012
	STATUS_CLUSTER_NODE_PAUSED                                                      = 0xC0130013
	STATUS_CLUSTER_NODE_NOT_PAUSED                                                  = 0xC0130014
	STATUS_CLUSTER_NO_SECURITY_CONTEXT                                              = 0xC0130015
	STATUS_CLUSTER_NETWORK_NOT_INTERNAL                                             = 0xC0130016
	STATUS_CLUSTER_POISONED                                                         = 0xC0130017
	STATUS_CLUSTER_NON_CSV_PATH                                                     = 0xC0130018
	STATUS_CLUSTER_CSV_VOLUME_NOT_LOCAL                                             = 0xC0130019
	STATUS_TRANSACTIONAL_CONFLICT                                                   = 0xC0190001
	STATUS_INVALID_TRANSACTION                                                      = 0xC0190002
	STATUS_TRANSACTION_NOT_ACTIVE                                                   = 0xC0190003
	STATUS_TM_INITIALIZATION_FAILED                                                 = 0xC0190004
	STATUS_RM_NOT_ACTIVE                                                            = 0xC0190005
	STATUS_RM_METADATA_CORRUPT                                                      = 0xC0190006
	STATUS_TRANSACTION_NOT_JOINED                                                   = 0xC0190007
	STATUS_DIRECTORY_NOT_RM                                                         = 0xC0190008
	STATUS_COULD_NOT_RESIZE_LOG                                                     = 0x80190009
	STATUS_TRANSACTIONS_UNSUPPORTED_REMOTE                                          = 0xC019000A
	STATUS_LOG_RESIZE_INVALID_SIZE                                                  = 0xC019000B
	STATUS_REMOTE_FILE_VERSION_MISMATCH                                             = 0xC019000C
	STATUS_CRM_PROTOCOL_ALREADY_EXISTS                                              = 0xC019000F
	STATUS_TRANSACTION_PROPAGATION_FAILED                                           = 0xC0190010
	STATUS_CRM_PROTOCOL_NOT_FOUND                                                   = 0xC0190011
	STATUS_TRANSACTION_SUPERIOR_EXISTS                                              = 0xC0190012
	STATUS_TRANSACTION_REQUEST_NOT_VALID                                            = 0xC0190013
	STATUS_TRANSACTION_NOT_REQUESTED                                                = 0xC0190014
	STATUS_TRANSACTION_ALREADY_ABORTED                                              = 0xC0190015
	STATUS_TRANSACTION_ALREADY_COMMITTED                                            = 0xC0190016
	STATUS_TRANSACTION_INVALID_MARSHALL_BUFFER                                      = 0xC0190017
	STATUS_CURRENT_TRANSACTION_NOT_VALID                                            = 0xC0190018
	STATUS_LOG_GROWTH_FAILED                                                        = 0xC0190019
	STATUS_OBJECT_NO_LONGER_EXISTS                                                  = 0xC0190021
	STATUS_STREAM_MINIVERSION_NOT_FOUND                                             = 0xC0190022
	STATUS_STREAM_MINIVERSION_NOT_VALID                                             = 0xC0190023
	STATUS_MINIVERSION_INACCESSIBLE_FROM_SPECIFIED_TRANSACTION                      = 0xC0190024
	STATUS_CANT_OPEN_MINIVERSION_WITH_MODIFY_INTENT                                 = 0xC0190025
	STATUS_CANT_CREATE_MORE_STREAM_MINIVERSIONS                                     = 0xC0190026
	STATUS_HANDLE_NO_LONGER_VALID                                                   = 0xC0190028
	STATUS_NO_TXF_METADATA                                                          = 0x80190029
	STATUS_LOG_CORRUPTION_DETECTED                                                  = 0xC0190030
	STATUS_CANT_RECOVER_WITH_HANDLE_OPEN                                            = 0x80190031
	STATUS_RM_DISCONNECTED                                                          = 0xC0190032
	STATUS_ENLISTMENT_NOT_SUPERIOR                                                  = 0xC0190033
	STATUS_RECOVERY_NOT_NEEDED                                                      = 0x40190034
	STATUS_RM_ALREADY_STARTED                                                       = 0x40190035
	STATUS_FILE_IDENTITY_NOT_PERSISTENT                                             = 0xC0190036
	STATUS_CANT_BREAK_TRANSACTIONAL_DEPENDENCY                                      = 0xC0190037
	STATUS_CANT_CROSS_RM_BOUNDARY                                                   = 0xC0190038
	STATUS_TXF_DIR_NOT_EMPTY                                                        = 0xC0190039
	STATUS_INDOUBT_TRANSACTIONS_EXIST                                               = 0xC019003A
	STATUS_TM_VOLATILE                                                              = 0xC019003B
	STATUS_ROLLBACK_TIMER_EXPIRED                                                   = 0xC019003C
	STATUS_TXF_ATTRIBUTE_CORRUPT                                                    = 0xC019003D
	STATUS_EFS_NOT_ALLOWED_IN_TRANSACTION                                           = 0xC019003E
	STATUS_TRANSACTIONAL_OPEN_NOT_ALLOWED                                           = 0xC019003F
	STATUS_TRANSACTED_MAPPING_UNSUPPORTED_REMOTE                                    = 0xC0190040
	STATUS_TXF_METADATA_ALREADY_PRESENT                                             = 0x80190041
	STATUS_TRANSACTION_SCOPE_CALLBACKS_NOT_SET                                      = 0x80190042
	STATUS_TRANSACTION_REQUIRED_PROMOTION                                           = 0xC0190043
	STATUS_CANNOT_EXECUTE_FILE_IN_TRANSACTION                                       = 0xC0190044
	STATUS_TRANSACTIONS_NOT_FROZEN                                                  = 0xC0190045
	STATUS_TRANSACTION_FREEZE_IN_PROGRESS                                           = 0xC0190046
	STATUS_NOT_SNAPSHOT_VOLUME                                                      = 0xC0190047
	STATUS_NO_SAVEPOINT_WITH_OPEN_FILES                                             = 0xC0190048
	STATUS_SPARSE_NOT_ALLOWED_IN_TRANSACTION                                        = 0xC0190049
	STATUS_TM_IDENTITY_MISMATCH                                                     = 0xC019004A
	STATUS_FLOATED_SECTION                                                          = 0xC019004B
	STATUS_CANNOT_ACCEPT_TRANSACTED_WORK                                            = 0xC019004C
	STATUS_CANNOT_ABORT_TRANSACTIONS                                                = 0xC019004D
	STATUS_TRANSACTION_NOT_FOUND                                                    = 0xC019004E
	STATUS_RESOURCEMANAGER_NOT_FOUND                                                = 0xC019004F
	STATUS_ENLISTMENT_NOT_FOUND                                                     = 0xC0190050
	STATUS_TRANSACTIONMANAGER_NOT_FOUND                                             = 0xC0190051
	STATUS_TRANSACTIONMANAGER_NOT_ONLINE                                            = 0xC0190052
	STATUS_TRANSACTIONMANAGER_RECOVERY_NAME_COLLISION                               = 0xC0190053
	STATUS_TRANSACTION_NOT_ROOT                                                     = 0xC0190054
	STATUS_TRANSACTION_OBJECT_EXPIRED                                               = 0xC0190055
	STATUS_COMPRESSION_NOT_ALLOWED_IN_TRANSACTION                                   = 0xC0190056
	STATUS_TRANSACTION_RESPONSE_NOT_ENLISTED                                        = 0xC0190057
	STATUS_TRANSACTION_RECORD_TOO_LONG                                              = 0xC0190058
	STATUS_NO_LINK_TRACKING_IN_TRANSACTION                                          = 0xC0190059
	STATUS_OPERATION_NOT_SUPPORTED_IN_TRANSACTION                                   = 0xC019005A
	STATUS_TRANSACTION_INTEGRITY_VIOLATED                                           = 0xC019005B
	STATUS_TRANSACTIONMANAGER_IDENTITY_MISMATCH                                     = 0xC019005C
	STATUS_RM_CANNOT_BE_FROZEN_FOR_SNAPSHOT                                         = 0xC019005D
	STATUS_TRANSACTION_MUST_WRITETHROUGH                                            = 0xC019005E
	STATUS_TRANSACTION_NO_SUPERIOR                                                  = 0xC019005F
	STATUS_EXPIRED_HANDLE                                                           = 0xC0190060
	STATUS_TRANSACTION_NOT_ENLISTED                                                 = 0xC0190061
	STATUS_LOG_SECTOR_INVALID                                                       = 0xC01A0001
	STATUS_LOG_SECTOR_PARITY_INVALID                                                = 0xC01A0002
	STATUS_LOG_SECTOR_REMAPPED                                                      = 0xC01A0003
	STATUS_LOG_BLOCK_INCOMPLETE                                                     = 0xC01A0004
	STATUS_LOG_INVALID_RANGE                                                        = 0xC01A0005
	STATUS_LOG_BLOCKS_EXHAUSTED                                                     = 0xC01A0006
	STATUS_LOG_READ_CONTEXT_INVALID                                                 = 0xC01A0007
	STATUS_LOG_RESTART_INVALID                                                      = 0xC01A0008
	STATUS_LOG_BLOCK_VERSION                                                        = 0xC01A0009
	STATUS_LOG_BLOCK_INVALID                                                        = 0xC01A000A
	STATUS_LOG_READ_MODE_INVALID                                                    = 0xC01A000B
	STATUS_LOG_NO_RESTART                                                           = 0x401A000C
	STATUS_LOG_METADATA_CORRUPT                                                     = 0xC01A000D
	STATUS_LOG_METADATA_INVALID                                                     = 0xC01A000E
	STATUS_LOG_METADATA_INCONSISTENT                                                = 0xC01A000F
	STATUS_LOG_RESERVATION_INVALID                                                  = 0xC01A0010
	STATUS_LOG_CANT_DELETE                                                          = 0xC01A0011
	STATUS_LOG_CONTAINER_LIMIT_EXCEEDED                                             = 0xC01A0012
	STATUS_LOG_START_OF_LOG                                                         = 0xC01A0013
	STATUS_LOG_POLICY_ALREADY_INSTALLED                                             = 0xC01A0014
	STATUS_LOG_POLICY_NOT_INSTALLED                                                 = 0xC01A0015
	STATUS_LOG_POLICY_INVALID                                                       = 0xC01A0016
	STATUS_LOG_POLICY_CONFLICT                                                      = 0xC01A0017
	STATUS_LOG_PINNED_ARCHIVE_TAIL                                                  = 0xC01A0018
	STATUS_LOG_RECORD_NONEXISTENT                                                   = 0xC01A0019
	STATUS_LOG_RECORDS_RESERVED_INVALID                                             = 0xC01A001A
	STATUS_LOG_SPACE_RESERVED_INVALID                                               = 0xC01A001B
	STATUS_LOG_TAIL_INVALID                                                         = 0xC01A001C
	STATUS_LOG_FULL                                                                 = 0xC01A001D
	STATUS_LOG_MULTIPLEXED                                                          = 0xC01A001E
	STATUS_LOG_DEDICATED                                                            = 0xC01A001F
	STATUS_LOG_ARCHIVE_NOT_IN_PROGRESS                                              = 0xC01A0020
	STATUS_LOG_ARCHIVE_IN_PROGRESS                                                  = 0xC01A0021
	STATUS_LOG_EPHEMERAL                                                            = 0xC01A0022
	STATUS_LOG_NOT_ENOUGH_CONTAINERS                                                = 0xC01A0023
	STATUS_LOG_CLIENT_ALREADY_REGISTERED                                            = 0xC01A0024
	STATUS_LOG_CLIENT_NOT_REGISTERED                                                = 0xC01A0025
	STATUS_LOG_FULL_HANDLER_IN_PROGRESS                                             = 0xC01A0026
	STATUS_LOG_CONTAINER_READ_FAILED                                                = 0xC01A0027
	STATUS_LOG_CONTAINER_WRITE_FAILED                                               = 0xC01A0028
	STATUS_LOG_CONTAINER_OPEN_FAILED                                                = 0xC01A0029
	STATUS_LOG_CONTAINER_STATE_INVALID                                              = 0xC01A002A
	STATUS_LOG_STATE_INVALID                                                        = 0xC01A002B
	STATUS_LOG_PINNED                                                               = 0xC01A002C
	STATUS_LOG_METADATA_FLUSH_FAILED                                                = 0xC01A002D
	STATUS_LOG_INCONSISTENT_SECURITY                                                = 0xC01A002E
	STATUS_LOG_APPENDED_FLUSH_FAILED                                                = 0xC01A002F
	STATUS_LOG_PINNED_RESERVATION                                                   = 0xC01A0030
	STATUS_VIDEO_HUNG_DISPLAY_DRIVER_THREAD                                         = 0xC01B00EA
	STATUS_VIDEO_HUNG_DISPLAY_DRIVER_THREAD_RECOVERED                               = 0x801B00EB
	STATUS_VIDEO_DRIVER_DEBUG_REPORT_REQUEST                                        = 0x401B00EC
	STATUS_MONITOR_NO_DESCRIPTOR                                                    = 0xC01D0001
	STATUS_MONITOR_UNKNOWN_DESCRIPTOR_FORMAT                                        = 0xC01D0002
	STATUS_MONITOR_INVALID_DESCRIPTOR_CHECKSUM                                      = 0xC01D0003
	STATUS_MONITOR_INVALID_STANDARD_TIMING_BLOCK                                    = 0xC01D0004
	STATUS_MONITOR_WMI_DATABLOCK_REGISTRATION_FAILED                                = 0xC01D0005
	STATUS_MONITOR_INVALID_SERIAL_NUMBER_MONDSC_BLOCK                               = 0xC01D0006
	STATUS_MONITOR_INVALID_USER_FRIENDLY_MONDSC_BLOCK                               = 0xC01D0007
	STATUS_MONITOR_NO_MORE_DESCRIPTOR_DATA                                          = 0xC01D0008
	STATUS_MONITOR_INVALID_DETAILED_TIMING_BLOCK                                    = 0xC01D0009
	STATUS_MONITOR_INVALID_MANUFACTURE_DATE                                         = 0xC01D000A
	STATUS_GRAPHICS_NOT_EXCLUSIVE_MODE_OWNER                                        = 0xC01E0000
	STATUS_GRAPHICS_INSUFFICIENT_DMA_BUFFER                                         = 0xC01E0001
	STATUS_GRAPHICS_INVALID_DISPLAY_ADAPTER                                         = 0xC01E0002
	STATUS_GRAPHICS_ADAPTER_WAS_RESET                                               = 0xC01E0003
	STATUS_GRAPHICS_INVALID_DRIVER_MODEL                                            = 0xC01E0004
	STATUS_GRAPHICS_PRESENT_MODE_CHANGED                                            = 0xC01E0005
	STATUS_GRAPHICS_PRESENT_OCCLUDED                                                = 0xC01E0006
	STATUS_GRAPHICS_PRESENT_DENIED                                                  = 0xC01E0007
	STATUS_GRAPHICS_CANNOTCOLORCONVERT                                              = 0xC01E0008
	STATUS_GRAPHICS_DRIVER_MISMATCH                                                 = 0xC01E0009
	STATUS_GRAPHICS_PARTIAL_DATA_POPULATED                                          = 0x401E000A
	STATUS_GRAPHICS_PRESENT_REDIRECTION_DISABLED                                    = 0xC01E000B
	STATUS_GRAPHICS_PRESENT_UNOCCLUDED                                              = 0xC01E000C
	STATUS_GRAPHICS_NO_VIDEO_MEMORY                                                 = 0xC01E0100
	STATUS_GRAPHICS_CANT_LOCK_MEMORY                                                = 0xC01E0101
	STATUS_GRAPHICS_ALLOCATION_BUSY                                                 = 0xC01E0102
	STATUS_GRAPHICS_TOO_MANY_REFERENCES                                             = 0xC01E0103
	STATUS_GRAPHICS_TRY_AGAIN_LATER                                                 = 0xC01E0104
	STATUS_GRAPHICS_TRY_AGAIN_NOW                                                   = 0xC01E0105
	STATUS_GRAPHICS_ALLOCATION_INVALID                                              = 0xC01E0106
	STATUS_GRAPHICS_UNSWIZZLING_APERTURE_UNAVAILABLE                                = 0xC01E0107
	STATUS_GRAPHICS_UNSWIZZLING_APERTURE_UNSUPPORTED                                = 0xC01E0108
	STATUS_GRAPHICS_CANT_EVICT_PINNED_ALLOCATION                                    = 0xC01E0109
	STATUS_GRAPHICS_INVALID_ALLOCATION_USAGE                                        = 0xC01E0110
	STATUS_GRAPHICS_CANT_RENDER_LOCKED_ALLOCATION                                   = 0xC01E0111
	STATUS_GRAPHICS_ALLOCATION_CLOSED                                               = 0xC01E0112
	STATUS_GRAPHICS_INVALID_ALLOCATION_INSTANCE                                     = 0xC01E0113
	STATUS_GRAPHICS_INVALID_ALLOCATION_HANDLE                                       = 0xC01E0114
	STATUS_GRAPHICS_WRONG_ALLOCATION_DEVICE                                         = 0xC01E0115
	STATUS_GRAPHICS_ALLOCATION_CONTENT_LOST                                         = 0xC01E0116
	STATUS_GRAPHICS_GPU_EXCEPTION_ON_DEVICE                                         = 0xC01E0200
	STATUS_GRAPHICS_INVALID_VIDPN_TOPOLOGY                                          = 0xC01E0300
	STATUS_GRAPHICS_VIDPN_TOPOLOGY_NOT_SUPPORTED                                    = 0xC01E0301
	STATUS_GRAPHICS_VIDPN_TOPOLOGY_CURRENTLY_NOT_SUPPORTED                          = 0xC01E0302
	STATUS_GRAPHICS_INVALID_VIDPN                                                   = 0xC01E0303
	STATUS_GRAPHICS_INVALID_VIDEO_PRESENT_SOURCE                                    = 0xC01E0304
	STATUS_GRAPHICS_INVALID_VIDEO_PRESENT_TARGET                                    = 0xC01E0305
	STATUS_GRAPHICS_VIDPN_MODALITY_NOT_SUPPORTED                                    = 0xC01E0306
	STATUS_GRAPHICS_MODE_NOT_PINNED                                                 = 0x401E0307
	STATUS_GRAPHICS_INVALID_VIDPN_SOURCEMODESET                                     = 0xC01E0308
	STATUS_GRAPHICS_INVALID_VIDPN_TARGETMODESET                                     = 0xC01E0309
	STATUS_GRAPHICS_INVALID_FREQUENCY                                               = 0xC01E030A
	STATUS_GRAPHICS_INVALID_ACTIVE_REGION                                           = 0xC01E030B
	STATUS_GRAPHICS_INVALID_TOTAL_REGION                                            = 0xC01E030C
	STATUS_GRAPHICS_INVALID_VIDEO_PRESENT_SOURCE_MODE                               = 0xC01E0310
	STATUS_GRAPHICS_INVALID_VIDEO_PRESENT_TARGET_MODE                               = 0xC01E0311
	STATUS_GRAPHICS_PINNED_MODE_MUST_REMAIN_IN_SET                                  = 0xC01E0312
	STATUS_GRAPHICS_PATH_ALREADY_IN_TOPOLOGY                                        = 0xC01E0313
	STATUS_GRAPHICS_MODE_ALREADY_IN_MODESET                                         = 0xC01E0314
	STATUS_GRAPHICS_INVALID_VIDEOPRESENTSOURCESET                                   = 0xC01E0315
	STATUS_GRAPHICS_INVALID_VIDEOPRESENTTARGETSET                                   = 0xC01E0316
	STATUS_GRAPHICS_SOURCE_ALREADY_IN_SET                                           = 0xC01E0317
	STATUS_GRAPHICS_TARGET_ALREADY_IN_SET                                           = 0xC01E0318
	STATUS_GRAPHICS_INVALID_VIDPN_PRESENT_PATH                                      = 0xC01E0319
	STATUS_GRAPHICS_NO_RECOMMENDED_VIDPN_TOPOLOGY                                   = 0xC01E031A
	STATUS_GRAPHICS_INVALID_MONITOR_FREQUENCYRANGESET                               = 0xC01E031B
	STATUS_GRAPHICS_INVALID_MONITOR_FREQUENCYRANGE                                  = 0xC01E031C
	STATUS_GRAPHICS_FREQUENCYRANGE_NOT_IN_SET                                       = 0xC01E031D
	STATUS_GRAPHICS_NO_PREFERRED_MODE                                               = 0x401E031E
	STATUS_GRAPHICS_FREQUENCYRANGE_ALREADY_IN_SET                                   = 0xC01E031F
	STATUS_GRAPHICS_STALE_MODESET                                                   = 0xC01E0320
	STATUS_GRAPHICS_INVALID_MONITOR_SOURCEMODESET                                   = 0xC01E0321
	STATUS_GRAPHICS_INVALID_MONITOR_SOURCE_MODE                                     = 0xC01E0322
	STATUS_GRAPHICS_NO_RECOMMENDED_FUNCTIONAL_VIDPN                                 = 0xC01E0323
	STATUS_GRAPHICS_MODE_ID_MUST_BE_UNIQUE                                          = 0xC01E0324
	STATUS_GRAPHICS_EMPTY_ADAPTER_MONITOR_MODE_SUPPORT_INTERSECTION                 = 0xC01E0325
	STATUS_GRAPHICS_VIDEO_PRESENT_TARGETS_LESS_THAN_SOURCES                         = 0xC01E0326
	STATUS_GRAPHICS_PATH_NOT_IN_TOPOLOGY                                            = 0xC01E0327
	STATUS_GRAPHICS_ADAPTER_MUST_HAVE_AT_LEAST_ONE_SOURCE                           = 0xC01E0328
	STATUS_GRAPHICS_ADAPTER_MUST_HAVE_AT_LEAST_ONE_TARGET                           = 0xC01E0329
	STATUS_GRAPHICS_INVALID_MONITORDESCRIPTORSET                                    = 0xC01E032A
	STATUS_GRAPHICS_INVALID_MONITORDESCRIPTOR                                       = 0xC01E032B
	STATUS_GRAPHICS_MONITORDESCRIPTOR_NOT_IN_SET                                    = 0xC01E032C
	STATUS_GRAPHICS_MONITORDESCRIPTOR_ALREADY_IN_SET                                = 0xC01E032D
	STATUS_GRAPHICS_MONITORDESCRIPTOR_ID_MUST_BE_UNIQUE                             = 0xC01E032E
	STATUS_GRAPHICS_INVALID_VIDPN_TARGET_SUBSET_TYPE                                = 0xC01E032F
	STATUS_GRAPHICS_RESOURCES_NOT_RELATED                                           = 0xC01E0330
	STATUS_GRAPHICS_SOURCE_ID_MUST_BE_UNIQUE                                        = 0xC01E0331
	STATUS_GRAPHICS_TARGET_ID_MUST_BE_UNIQUE                                        = 0xC01E0332
	STATUS_GRAPHICS_NO_AVAILABLE_VIDPN_TARGET                                       = 0xC01E0333
	STATUS_GRAPHICS_MONITOR_COULD_NOT_BE_ASSOCIATED_WITH_ADAPTER                    = 0xC01E0334
	STATUS_GRAPHICS_NO_VIDPNMGR                                                     = 0xC01E0335
	STATUS_GRAPHICS_NO_ACTIVE_VIDPN                                                 = 0xC01E0336
	STATUS_GRAPHICS_STALE_VIDPN_TOPOLOGY                                            = 0xC01E0337
	STATUS_GRAPHICS_MONITOR_NOT_CONNECTED                                           = 0xC01E0338
	STATUS_GRAPHICS_SOURCE_NOT_IN_TOPOLOGY                                          = 0xC01E0339
	STATUS_GRAPHICS_INVALID_PRIMARYSURFACE_SIZE                                     = 0xC01E033A
	STATUS_GRAPHICS_INVALID_VISIBLEREGION_SIZE                                      = 0xC01E033B
	STATUS_GRAPHICS_INVALID_STRIDE                                                  = 0xC01E033C
	STATUS_GRAPHICS_INVALID_PIXELFORMAT                                             = 0xC01E033D
	STATUS_GRAPHICS_INVALID_COLORBASIS                                              = 0xC01E033E
	STATUS_GRAPHICS_INVALID_PIXELVALUEACCESSMODE                                    = 0xC01E033F
	STATUS_GRAPHICS_TARGET_NOT_IN_TOPOLOGY                                          = 0xC01E0340
	STATUS_GRAPHICS_NO_DISPLAY_MODE_MANAGEMENT_SUPPORT                              = 0xC01E0341
	STATUS_GRAPHICS_VIDPN_SOURCE_IN_USE                                             = 0xC01E0342
	STATUS_GRAPHICS_CANT_ACCESS_ACTIVE_VIDPN                                        = 0xC01E0343
	STATUS_GRAPHICS_INVALID_PATH_IMPORTANCE_ORDINAL                                 = 0xC01E0344
	STATUS_GRAPHICS_INVALID_PATH_CONTENT_GEOMETRY_TRANSFORMATION                    = 0xC01E0345
	STATUS_GRAPHICS_PATH_CONTENT_GEOMETRY_TRANSFORMATION_NOT_SUPPORTED              = 0xC01E0346
	STATUS_GRAPHICS_INVALID_GAMMA_RAMP                                              = 0xC01E0347
	STATUS_GRAPHICS_GAMMA_RAMP_NOT_SUPPORTED                                        = 0xC01E0348
	STATUS_GRAPHICS_MULTISAMPLING_NOT_SUPPORTED                                     = 0xC01E0349
	STATUS_GRAPHICS_MODE_NOT_IN_MODESET                                             = 0xC01E034A
	STATUS_GRAPHICS_DATASET_IS_EMPTY                                                = 0x401E034B
	STATUS_GRAPHICS_NO_MORE_ELEMENTS_IN_DATASET                                     = 0x401E034C
	STATUS_GRAPHICS_INVALID_VIDPN_TOPOLOGY_RECOMMENDATION_REASON                    = 0xC01E034D
	STATUS_GRAPHICS_INVALID_PATH_CONTENT_TYPE                                       = 0xC01E034E
	STATUS_GRAPHICS_INVALID_COPYPROTECTION_TYPE                                     = 0xC01E034F
	STATUS_GRAPHICS_UNASSIGNED_MODESET_ALREADY_EXISTS                               = 0xC01E0350
	STATUS_GRAPHICS_PATH_CONTENT_GEOMETRY_TRANSFORMATION_NOT_PINNED                 = 0x401E0351
	STATUS_GRAPHICS_INVALID_SCANLINE_ORDERING                                       = 0xC01E0352
	STATUS_GRAPHICS_TOPOLOGY_CHANGES_NOT_ALLOWED                                    = 0xC01E0353
	STATUS_GRAPHICS_NO_AVAILABLE_IMPORTANCE_ORDINALS                                = 0xC01E0354
	STATUS_GRAPHICS_INCOMPATIBLE_PRIVATE_FORMAT                                     = 0xC01E0355
	STATUS_GRAPHICS_INVALID_MODE_PRUNING_ALGORITHM                                  = 0xC01E0356
	STATUS_GRAPHICS_INVALID_MONITOR_CAPABILITY_ORIGIN                               = 0xC01E0357
	STATUS_GRAPHICS_INVALID_MONITOR_FREQUENCYRANGE_CONSTRAINT                       = 0xC01E0358
	STATUS_GRAPHICS_MAX_NUM_PATHS_REACHED                                           = 0xC01E0359
	STATUS_GRAPHICS_CANCEL_VIDPN_TOPOLOGY_AUGMENTATION                              = 0xC01E035A
	STATUS_GRAPHICS_INVALID_CLIENT_TYPE                                             = 0xC01E035B
	STATUS_GRAPHICS_CLIENTVIDPN_NOT_SET                                             = 0xC01E035C
	STATUS_GRAPHICS_SPECIFIED_CHILD_ALREADY_CONNECTED                               = 0xC01E0400
	STATUS_GRAPHICS_CHILD_DESCRIPTOR_NOT_SUPPORTED                                  = 0xC01E0401
	STATUS_GRAPHICS_UNKNOWN_CHILD_STATUS                                            = 0x401E042F
	STATUS_GRAPHICS_NOT_A_LINKED_ADAPTER                                            = 0xC01E0430
	STATUS_GRAPHICS_LEADLINK_NOT_ENUMERATED                                         = 0xC01E0431
	STATUS_GRAPHICS_CHAINLINKS_NOT_ENUMERATED                                       = 0xC01E0432
	STATUS_GRAPHICS_ADAPTER_CHAIN_NOT_READY                                         = 0xC01E0433
	STATUS_GRAPHICS_CHAINLINKS_NOT_STARTED                                          = 0xC01E0434
	STATUS_GRAPHICS_CHAINLINKS_NOT_POWERED_ON                                       = 0xC01E0435
	STATUS_GRAPHICS_INCONSISTENT_DEVICE_LINK_STATE                                  = 0xC01E0436
	STATUS_GRAPHICS_LEADLINK_START_DEFERRED                                         = 0x401E0437
	STATUS_GRAPHICS_NOT_POST_DEVICE_DRIVER                                          = 0xC01E0438
	STATUS_GRAPHICS_POLLING_TOO_FREQUENTLY                                          = 0x401E0439
	STATUS_GRAPHICS_START_DEFERRED                                                  = 0x401E043A
	STATUS_GRAPHICS_ADAPTER_ACCESS_NOT_EXCLUDED                                     = 0xC01E043B
	STATUS_GRAPHICS_OPM_NOT_SUPPORTED                                               = 0xC01E0500
	STATUS_GRAPHICS_COPP_NOT_SUPPORTED                                              = 0xC01E0501
	STATUS_GRAPHICS_UAB_NOT_SUPPORTED                                               = 0xC01E0502
	STATUS_GRAPHICS_OPM_INVALID_ENCRYPTED_PARAMETERS                                = 0xC01E0503
	STATUS_GRAPHICS_OPM_NO_PROTECTED_OUTPUTS_EXIST                                  = 0xC01E0505
	STATUS_GRAPHICS_OPM_INTERNAL_ERROR                                              = 0xC01E050B
	STATUS_GRAPHICS_OPM_INVALID_HANDLE                                              = 0xC01E050C
	STATUS_GRAPHICS_PVP_INVALID_CERTIFICATE_LENGTH                                  = 0xC01E050E
	STATUS_GRAPHICS_OPM_SPANNING_MODE_ENABLED                                       = 0xC01E050F
	STATUS_GRAPHICS_OPM_THEATER_MODE_ENABLED                                        = 0xC01E0510
	STATUS_GRAPHICS_PVP_HFS_FAILED                                                  = 0xC01E0511
	STATUS_GRAPHICS_OPM_INVALID_SRM                                                 = 0xC01E0512
	STATUS_GRAPHICS_OPM_OUTPUT_DOES_NOT_SUPPORT_HDCP                                = 0xC01E0513
	STATUS_GRAPHICS_OPM_OUTPUT_DOES_NOT_SUPPORT_ACP                                 = 0xC01E0514
	STATUS_GRAPHICS_OPM_OUTPUT_DOES_NOT_SUPPORT_CGMSA                               = 0xC01E0515
	STATUS_GRAPHICS_OPM_HDCP_SRM_NEVER_SET                                          = 0xC01E0516
	STATUS_GRAPHICS_OPM_RESOLUTION_TOO_HIGH                                         = 0xC01E0517
	STATUS_GRAPHICS_OPM_ALL_HDCP_HARDWARE_ALREADY_IN_USE                            = 0xC01E0518
	STATUS_GRAPHICS_OPM_PROTECTED_OUTPUT_NO_LONGER_EXISTS                           = 0xC01E051A
	STATUS_GRAPHICS_OPM_PROTECTED_OUTPUT_DOES_NOT_HAVE_COPP_SEMANTICS               = 0xC01E051C
	STATUS_GRAPHICS_OPM_INVALID_INFORMATION_REQUEST                                 = 0xC01E051D
	STATUS_GRAPHICS_OPM_DRIVER_INTERNAL_ERROR                                       = 0xC01E051E
	STATUS_GRAPHICS_OPM_PROTECTED_OUTPUT_DOES_NOT_HAVE_OPM_SEMANTICS                = 0xC01E051F
	STATUS_GRAPHICS_OPM_SIGNALING_NOT_SUPPORTED                                     = 0xC01E0520
	STATUS_GRAPHICS_OPM_INVALID_CONFIGURATION_REQUEST                               = 0xC01E0521
	STATUS_GRAPHICS_I2C_NOT_SUPPORTED                                               = 0xC01E0580
	STATUS_GRAPHICS_I2C_DEVICE_DOES_NOT_EXIST                                       = 0xC01E0581
	STATUS_GRAPHICS_I2C_ERROR_TRANSMITTING_DATA                                     = 0xC01E0582
	STATUS_GRAPHICS_I2C_ERROR_RECEIVING_DATA                                        = 0xC01E0583
	STATUS_GRAPHICS_DDCCI_VCP_NOT_SUPPORTED                                         = 0xC01E0584
	STATUS_GRAPHICS_DDCCI_INVALID_DATA                                              = 0xC01E0585
	STATUS_GRAPHICS_DDCCI_MONITOR_RETURNED_INVALID_TIMING_STATUS_BYTE               = 0xC01E0586
	STATUS_GRAPHICS_DDCCI_INVALID_CAPABILITIES_STRING                               = 0xC01E0587
	STATUS_GRAPHICS_MCA_INTERNAL_ERROR                                              = 0xC01E0588
	STATUS_GRAPHICS_DDCCI_INVALID_MESSAGE_COMMAND                                   = 0xC01E0589
	STATUS_GRAPHICS_DDCCI_INVALID_MESSAGE_LENGTH                                    = 0xC01E058A
	STATUS_GRAPHICS_DDCCI_INVALID_MESSAGE_CHECKSUM                                  = 0xC01E058B
	STATUS_GRAPHICS_INVALID_PHYSICAL_MONITOR_HANDLE                                 = 0xC01E058C
	STATUS_GRAPHICS_MONITOR_NO_LONGER_EXISTS                                        = 0xC01E058D
	STATUS_GRAPHICS_ONLY_CONSOLE_SESSION_SUPPORTED                                  = 0xC01E05E0
	STATUS_GRAPHICS_NO_DISPLAY_DEVICE_CORRESPONDS_TO_NAME                           = 0xC01E05E1
	STATUS_GRAPHICS_DISPLAY_DEVICE_NOT_ATTACHED_TO_DESKTOP                          = 0xC01E05E2
	STATUS_GRAPHICS_MIRRORING_DEVICES_NOT_SUPPORTED                                 = 0xC01E05E3
	STATUS_GRAPHICS_INVALID_POINTER                                                 = 0xC01E05E4
	STATUS_GRAPHICS_NO_MONITORS_CORRESPOND_TO_DISPLAY_DEVICE                        = 0xC01E05E5
	STATUS_GRAPHICS_PARAMETER_ARRAY_TOO_SMALL                                       = 0xC01E05E6
	STATUS_GRAPHICS_INTERNAL_ERROR                                                  = 0xC01E05E7
	STATUS_GRAPHICS_SESSION_TYPE_CHANGE_IN_PROGRESS                                 = 0xC01E05E8
	STATUS_FVE_LOCKED_VOLUME                                                        = 0xC0210000
	STATUS_FVE_NOT_ENCRYPTED                                                        = 0xC0210001
	STATUS_FVE_BAD_INFORMATION                                                      = 0xC0210002
	STATUS_FVE_TOO_SMALL                                                            = 0xC0210003
	STATUS_FVE_FAILED_WRONG_FS                                                      = 0xC0210004
	STATUS_FVE_BAD_PARTITION_SIZE                                                   = 0xC0210005
	STATUS_FVE_FS_NOT_EXTENDED                                                      = 0xC0210006
	STATUS_FVE_FS_MOUNTED                                                           = 0xC0210007
	STATUS_FVE_NO_LICENSE                                                           = 0xC0210008
	STATUS_FVE_ACTION_NOT_ALLOWED                                                   = 0xC0210009
	STATUS_FVE_BAD_DATA                                                             = 0xC021000A
	STATUS_FVE_VOLUME_NOT_BOUND                                                     = 0xC021000B
	STATUS_FVE_NOT_DATA_VOLUME                                                      = 0xC021000C
	STATUS_FVE_CONV_READ_ERROR                                                      = 0xC021000D
	STATUS_FVE_CONV_WRITE_ERROR                                                     = 0xC021000E
	STATUS_FVE_OVERLAPPED_UPDATE                                                    = 0xC021000F
	STATUS_FVE_FAILED_SECTOR_SIZE                                                   = 0xC0210010
	STATUS_FVE_FAILED_AUTHENTICATION                                                = 0xC0210011
	STATUS_FVE_NOT_OS_VOLUME                                                        = 0xC0210012
	STATUS_FVE_KEYFILE_NOT_FOUND                                                    = 0xC0210013
	STATUS_FVE_KEYFILE_INVALID                                                      = 0xC0210014
	STATUS_FVE_KEYFILE_NO_VMK                                                       = 0xC0210015
	STATUS_FVE_TPM_DISABLED                                                         = 0xC0210016
	STATUS_FVE_TPM_SRK_AUTH_NOT_ZERO                                                = 0xC0210017
	STATUS_FVE_TPM_INVALID_PCR                                                      = 0xC0210018
	STATUS_FVE_TPM_NO_VMK                                                           = 0xC0210019
	STATUS_FVE_PIN_INVALID                                                          = 0xC021001A
	STATUS_FVE_AUTH_INVALID_APPLICATION                                             = 0xC021001B
	STATUS_FVE_AUTH_INVALID_CONFIG                                                  = 0xC021001C
	STATUS_FVE_DEBUGGER_ENABLED                                                     = 0xC021001D
	STATUS_FVE_DRY_RUN_FAILED                                                       = 0xC021001E
	STATUS_FVE_BAD_METADATA_POINTER                                                 = 0xC021001F
	STATUS_FVE_OLD_METADATA_COPY                                                    = 0xC0210020
	STATUS_FVE_REBOOT_REQUIRED                                                      = 0xC0210021
	STATUS_FVE_RAW_ACCESS                                                           = 0xC0210022
	STATUS_FVE_RAW_BLOCKED                                                          = 0xC0210023
	STATUS_FVE_NO_AUTOUNLOCK_MASTER_KEY                                             = 0xC0210024
	STATUS_FVE_MOR_FAILED                                                           = 0xC0210025
	STATUS_FVE_NO_FEATURE_LICENSE                                                   = 0xC0210026
	STATUS_FVE_POLICY_USER_DISABLE_RDV_NOT_ALLOWED                                  = 0xC0210027
	STATUS_FVE_CONV_RECOVERY_FAILED                                                 = 0xC0210028
	STATUS_FVE_VIRTUALIZED_SPACE_TOO_BIG                                            = 0xC0210029
	STATUS_FVE_INVALID_DATUM_TYPE                                                   = 0xC021002A
	STATUS_FVE_VOLUME_TOO_SMALL                                                     = 0xC0210030
	STATUS_FVE_ENH_PIN_INVALID                                                      = 0xC0210031
	STATUS_FWP_CALLOUT_NOT_FOUND                                                    = 0xC0220001
	STATUS_FWP_CONDITION_NOT_FOUND                                                  = 0xC0220002
	STATUS_FWP_FILTER_NOT_FOUND                                                     = 0xC0220003
	STATUS_FWP_LAYER_NOT_FOUND                                                      = 0xC0220004
	STATUS_FWP_PROVIDER_NOT_FOUND                                                   = 0xC0220005
	STATUS_FWP_PROVIDER_CONTEXT_NOT_FOUND                                           = 0xC0220006
	STATUS_FWP_SUBLAYER_NOT_FOUND                                                   = 0xC0220007
	STATUS_FWP_NOT_FOUND                                                            = 0xC0220008
	STATUS_FWP_ALREADY_EXISTS                                                       = 0xC0220009
	STATUS_FWP_IN_USE                                                               = 0xC022000A
	STATUS_FWP_DYNAMIC_SESSION_IN_PROGRESS                                          = 0xC022000B
	STATUS_FWP_WRONG_SESSION                                                        = 0xC022000C
	STATUS_FWP_NO_TXN_IN_PROGRESS                                                   = 0xC022000D
	STATUS_FWP_TXN_IN_PROGRESS                                                      = 0xC022000E
	STATUS_FWP_TXN_ABORTED                                                          = 0xC022000F
	STATUS_FWP_SESSION_ABORTED                                                      = 0xC0220010
	STATUS_FWP_INCOMPATIBLE_TXN                                                     = 0xC0220011
	STATUS_FWP_TIMEOUT                                                              = 0xC0220012
	STATUS_FWP_NET_EVENTS_DISABLED                                                  = 0xC0220013
	STATUS_FWP_INCOMPATIBLE_LAYER                                                   = 0xC0220014
	STATUS_FWP_KM_CLIENTS_ONLY                                                      = 0xC0220015
	STATUS_FWP_LIFETIME_MISMATCH                                                    = 0xC0220016
	STATUS_FWP_BUILTIN_OBJECT                                                       = 0xC0220017
	STATUS_FWP_TOO_MANY_CALLOUTS                                                    = 0xC0220018
	STATUS_FWP_NOTIFICATION_DROPPED                                                 = 0xC0220019
	STATUS_FWP_TRAFFIC_MISMATCH                                                     = 0xC022001A
	STATUS_FWP_INCOMPATIBLE_SA_STATE                                                = 0xC022001B
	STATUS_FWP_NULL_POINTER                                                         = 0xC022001C
	STATUS_FWP_INVALID_ENUMERATOR                                                   = 0xC022001D
	STATUS_FWP_INVALID_FLAGS                                                        = 0xC022001E
	STATUS_FWP_INVALID_NET_MASK                                                     = 0xC022001F
	STATUS_FWP_INVALID_RANGE                                                        = 0xC0220020
	STATUS_FWP_INVALID_INTERVAL                                                     = 0xC0220021
	STATUS_FWP_ZERO_LENGTH_ARRAY                                                    = 0xC0220022
	STATUS_FWP_NULL_DISPLAY_NAME                                                    = 0xC0220023
	STATUS_FWP_INVALID_ACTION_TYPE                                                  = 0xC0220024
	STATUS_FWP_INVALID_WEIGHT                                                       = 0xC0220025
	STATUS_FWP_MATCH_TYPE_MISMATCH                                                  = 0xC0220026
	STATUS_FWP_TYPE_MISMATCH                                                        = 0xC0220027
	STATUS_FWP_OUT_OF_BOUNDS                                                        = 0xC0220028
	STATUS_FWP_RESERVED                                                             = 0xC0220029
	STATUS_FWP_DUPLICATE_CONDITION                                                  = 0xC022002A
	STATUS_FWP_DUPLICATE_KEYMOD                                                     = 0xC022002B
	STATUS_FWP_ACTION_INCOMPATIBLE_WITH_LAYER                                       = 0xC022002C
	STATUS_FWP_ACTION_INCOMPATIBLE_WITH_SUBLAYER                                    = 0xC022002D
	STATUS_FWP_CONTEXT_INCOMPATIBLE_WITH_LAYER                                      = 0xC022002E
	STATUS_FWP_CONTEXT_INCOMPATIBLE_WITH_CALLOUT                                    = 0xC022002F
	STATUS_FWP_INCOMPATIBLE_AUTH_METHOD                                             = 0xC0220030
	STATUS_FWP_INCOMPATIBLE_DH_GROUP                                                = 0xC0220031
	STATUS_FWP_EM_NOT_SUPPORTED                                                     = 0xC0220032
	STATUS_FWP_NEVER_MATCH                                                          = 0xC0220033
	STATUS_FWP_PROVIDER_CONTEXT_MISMATCH                                            = 0xC0220034
	STATUS_FWP_INVALID_PARAMETER                                                    = 0xC0220035
	STATUS_FWP_TOO_MANY_SUBLAYERS                                                   = 0xC0220036
	STATUS_FWP_CALLOUT_NOTIFICATION_FAILED                                          = 0xC0220037
	STATUS_FWP_INVALID_AUTH_TRANSFORM                                               = 0xC0220038
	STATUS_FWP_INVALID_CIPHER_TRANSFORM                                             = 0xC0220039
	STATUS_FWP_INCOMPATIBLE_CIPHER_TRANSFORM                                        = 0xC022003A
	STATUS_FWP_INVALID_TRANSFORM_COMBINATION                                        = 0xC022003B
	STATUS_FWP_DUPLICATE_AUTH_METHOD                                                = 0xC022003C
	STATUS_FWP_TCPIP_NOT_READY                                                      = 0xC0220100
	STATUS_FWP_INJECT_HANDLE_CLOSING                                                = 0xC0220101
	STATUS_FWP_INJECT_HANDLE_STALE                                                  = 0xC0220102
	STATUS_FWP_CANNOT_PEND                                                          = 0xC0220103
	STATUS_FWP_DROP_NOICMP                                                          = 0xC0220104
	STATUS_NDIS_CLOSING                                                             = 0xC0230002
	STATUS_NDIS_BAD_VERSION                                                         = 0xC0230004
	STATUS_NDIS_BAD_CHARACTERISTICS                                                 = 0xC0230005
	STATUS_NDIS_ADAPTER_NOT_FOUND                                                   = 0xC0230006
	STATUS_NDIS_OPEN_FAILED                                                         = 0xC0230007
	STATUS_NDIS_DEVICE_FAILED                                                       = 0xC0230008
	STATUS_NDIS_MULTICAST_FULL                                                      = 0xC0230009
	STATUS_NDIS_MULTICAST_EXISTS                                                    = 0xC023000A
	STATUS_NDIS_MULTICAST_NOT_FOUND                                                 = 0xC023000B
	STATUS_NDIS_REQUEST_ABORTED                                                     = 0xC023000C
	STATUS_NDIS_RESET_IN_PROGRESS                                                   = 0xC023000D
	STATUS_NDIS_NOT_SUPPORTED                                                       = 0xC02300BB
	STATUS_NDIS_INVALID_PACKET                                                      = 0xC023000F
	STATUS_NDIS_ADAPTER_NOT_READY                                                   = 0xC0230011
	STATUS_NDIS_INVALID_LENGTH                                                      = 0xC0230014
	STATUS_NDIS_INVALID_DATA                                                        = 0xC0230015
	STATUS_NDIS_BUFFER_TOO_SHORT                                                    = 0xC0230016
	STATUS_NDIS_INVALID_OID                                                         = 0xC0230017
	STATUS_NDIS_ADAPTER_REMOVED                                                     = 0xC0230018
	STATUS_NDIS_UNSUPPORTED_MEDIA                                                   = 0xC0230019
	STATUS_NDIS_GROUP_ADDRESS_IN_USE                                                = 0xC023001A
	STATUS_NDIS_FILE_NOT_FOUND                                                      = 0xC023001B
	STATUS_NDIS_ERROR_READING_FILE                                                  = 0xC023001C
	STATUS_NDIS_ALREADY_MAPPED                                                      = 0xC023001D
	STATUS_NDIS_RESOURCE_CONFLICT                                                   = 0xC023001E
	STATUS_NDIS_MEDIA_DISCONNECTED                                                  = 0xC023001F
	STATUS_NDIS_INVALID_ADDRESS                                                     = 0xC0230022
	STATUS_NDIS_INVALID_DEVICE_REQUEST                                              = 0xC0230010
	STATUS_NDIS_PAUSED                                                              = 0xC023002A
	STATUS_NDIS_INTERFACE_NOT_FOUND                                                 = 0xC023002B
	STATUS_NDIS_UNSUPPORTED_REVISION                                                = 0xC023002C
	STATUS_NDIS_INVALID_PORT                                                        = 0xC023002D
	STATUS_NDIS_INVALID_PORT_STATE                                                  = 0xC023002E
	STATUS_NDIS_LOW_POWER_STATE                                                     = 0xC023002F
	STATUS_NDIS_DOT11_AUTO_CONFIG_ENABLED                                           = 0xC0232000
	STATUS_NDIS_DOT11_MEDIA_IN_USE                                                  = 0xC0232001
	STATUS_NDIS_DOT11_POWER_STATE_INVALID                                           = 0xC0232002
	STATUS_NDIS_PM_WOL_PATTERN_LIST_FULL                                            = 0xC0232003
	STATUS_NDIS_PM_PROTOCOL_OFFLOAD_LIST_FULL                                       = 0xC0232004
	STATUS_NDIS_INDICATION_REQUIRED                                                 = 0x40230001
	STATUS_NDIS_OFFLOAD_POLICY                                                      = 0xC023100F
	STATUS_NDIS_OFFLOAD_CONNECTION_REJECTED                                         = 0xC0231012
	STATUS_NDIS_OFFLOAD_PATH_REJECTED                                               = 0xC0231013
	STATUS_HV_INVALID_HYPERCALL_CODE                                                = 0xC0350002
	STATUS_HV_INVALID_HYPERCALL_INPUT                                               = 0xC0350003
	STATUS_HV_INVALID_ALIGNMENT                                                     = 0xC0350004
	STATUS_HV_INVALID_PARAMETER                                                     = 0xC0350005
	STATUS_HV_ACCESS_DENIED                                                         = 0xC0350006
	STATUS_HV_INVALID_PARTITION_STATE                                               = 0xC0350007
	STATUS_HV_OPERATION_DENIED                                                      = 0xC0350008
	STATUS_HV_UNKNOWN_PROPERTY                                                      = 0xC0350009
	STATUS_HV_PROPERTY_VALUE_OUT_OF_RANGE                                           = 0xC035000A
	STATUS_HV_INSUFFICIENT_MEMORY                                                   = 0xC035000B
	STATUS_HV_PARTITION_TOO_DEEP                                                    = 0xC035000C
	STATUS_HV_INVALID_PARTITION_ID                                                  = 0xC035000D
	STATUS_HV_INVALID_VP_INDEX                                                      = 0xC035000E
	STATUS_HV_INVALID_PORT_ID                                                       = 0xC0350011
	STATUS_HV_INVALID_CONNECTION_ID                                                 = 0xC0350012
	STATUS_HV_INSUFFICIENT_BUFFERS                                                  = 0xC0350013
	STATUS_HV_NOT_ACKNOWLEDGED                                                      = 0xC0350014
	STATUS_HV_ACKNOWLEDGED                                                          = 0xC0350016
	STATUS_HV_INVALID_SAVE_RESTORE_STATE                                            = 0xC0350017
	STATUS_HV_INVALID_SYNIC_STATE                                                   = 0xC0350018
	STATUS_HV_OBJECT_IN_USE                                                         = 0xC0350019
	STATUS_HV_INVALID_PROXIMITY_DOMAIN_INFO                                         = 0xC035001A
	STATUS_HV_NO_DATA                                                               = 0xC035001B
	STATUS_HV_INACTIVE                                                              = 0xC035001C
	STATUS_HV_NO_RESOURCES                                                          = 0xC035001D
	STATUS_HV_FEATURE_UNAVAILABLE                                                   = 0xC035001E
	STATUS_HV_NOT_PRESENT                                                           = 0xC0351000
	STATUS_VID_DUPLICATE_HANDLER                                                    = 0xC0370001
	STATUS_VID_TOO_MANY_HANDLERS                                                    = 0xC0370002
	STATUS_VID_QUEUE_FULL                                                           = 0xC0370003
	STATUS_VID_HANDLER_NOT_PRESENT                                                  = 0xC0370004
	STATUS_VID_INVALID_OBJECT_NAME                                                  = 0xC0370005
	STATUS_VID_PARTITION_NAME_TOO_LONG                                              = 0xC0370006
	STATUS_VID_MESSAGE_QUEUE_NAME_TOO_LONG                                          = 0xC0370007
	STATUS_VID_PARTITION_ALREADY_EXISTS                                             = 0xC0370008
	STATUS_VID_PARTITION_DOES_NOT_EXIST                                             = 0xC0370009
	STATUS_VID_PARTITION_NAME_NOT_FOUND                                             = 0xC037000A
	STATUS_VID_MESSAGE_QUEUE_ALREADY_EXISTS                                         = 0xC037000B
	STATUS_VID_EXCEEDED_MBP_ENTRY_MAP_LIMIT                                         = 0xC037000C
	STATUS_VID_MB_STILL_REFERENCED                                                  = 0xC037000D
	STATUS_VID_CHILD_GPA_PAGE_SET_CORRUPTED                                         = 0xC037000E
	STATUS_VID_INVALID_NUMA_SETTINGS                                                = 0xC037000F
	STATUS_VID_INVALID_NUMA_NODE_INDEX                                              = 0xC0370010
	STATUS_VID_NOTIFICATION_QUEUE_ALREADY_ASSOCIATED                                = 0xC0370011
	STATUS_VID_INVALID_MEMORY_BLOCK_HANDLE                                          = 0xC0370012
	STATUS_VID_PAGE_RANGE_OVERFLOW                                                  = 0xC0370013
	STATUS_VID_INVALID_MESSAGE_QUEUE_HANDLE                                         = 0xC0370014
	STATUS_VID_INVALID_GPA_RANGE_HANDLE                                             = 0xC0370015
	STATUS_VID_NO_MEMORY_BLOCK_NOTIFICATION_QUEUE                                   = 0xC0370016
	STATUS_VID_MEMORY_BLOCK_LOCK_COUNT_EXCEEDED                                     = 0xC0370017
	STATUS_VID_INVALID_PPM_HANDLE                                                   = 0xC0370018
	STATUS_VID_MBPS_ARE_LOCKED                                                      = 0xC0370019
	STATUS_VID_MESSAGE_QUEUE_CLOSED                                                 = 0xC037001A
	STATUS_VID_VIRTUAL_PROCESSOR_LIMIT_EXCEEDED                                     = 0xC037001B
	STATUS_VID_STOP_PENDING                                                         = 0xC037001C
	STATUS_VID_INVALID_PROCESSOR_STATE                                              = 0xC037001D
	STATUS_VID_EXCEEDED_KM_CONTEXT_COUNT_LIMIT                                      = 0xC037001E
	STATUS_VID_KM_INTERFACE_ALREADY_INITIALIZED                                     = 0xC037001F
	STATUS_VID_MB_PROPERTY_ALREADY_SET_RESET                                        = 0xC0370020
	STATUS_VID_MMIO_RANGE_DESTROYED                                                 = 0xC0370021
	STATUS_VID_INVALID_CHILD_GPA_PAGE_SET                                           = 0xC0370022
	STATUS_VID_RESERVE_PAGE_SET_IS_BEING_USED                                       = 0xC0370023
	STATUS_VID_RESERVE_PAGE_SET_TOO_SMALL                                           = 0xC0370024
	STATUS_VID_MBP_ALREADY_LOCKED_USING_RESERVED_PAGE                               = 0xC0370025
	STATUS_VID_MBP_COUNT_EXCEEDED_LIMIT                                             = 0xC0370026
	STATUS_VID_SAVED_STATE_CORRUPT                                                  = 0xC0370027
	STATUS_VID_SAVED_STATE_UNRECOGNIZED_ITEM                                        = 0xC0370028
	STATUS_VID_SAVED_STATE_INCOMPATIBLE                                             = 0xC0370029
	STATUS_VID_REMOTE_NODE_PARENT_GPA_PAGES_USED                                    = 0x80370001
	STATUS_IPSEC_BAD_SPI                                                            = 0xC0360001
	STATUS_IPSEC_SA_LIFETIME_EXPIRED                                                = 0xC0360002
	STATUS_IPSEC_WRONG_SA                                                           = 0xC0360003
	STATUS_IPSEC_REPLAY_CHECK_FAILED                                                = 0xC0360004
	STATUS_IPSEC_INVALID_PACKET                                                     = 0xC0360005
	STATUS_IPSEC_INTEGRITY_CHECK_FAILED                                             = 0xC0360006
	STATUS_IPSEC_CLEAR_TEXT_DROP                                                    = 0xC0360007
	STATUS_IPSEC_AUTH_FIREWALL_DROP                                                 = 0xC0360008
	STATUS_IPSEC_THROTTLE_DROP                                                      = 0xC0360009
	STATUS_IPSEC_DOSP_BLOCK                                                         = 0xC0368000
	STATUS_IPSEC_DOSP_RECEIVED_MULTICAST                                            = 0xC0368001
	STATUS_IPSEC_DOSP_INVALID_PACKET                                                = 0xC0368002
	STATUS_IPSEC_DOSP_STATE_LOOKUP_FAILED                                           = 0xC0368003
	STATUS_IPSEC_DOSP_MAX_ENTRIES                                                   = 0xC0368004
	STATUS_IPSEC_DOSP_KEYMOD_NOT_ALLOWED                                            = 0xC0368005
	STATUS_IPSEC_DOSP_MAX_PER_IP_RATELIMIT_QUEUES                                   = 0xC0368006
	STATUS_VOLMGR_INCOMPLETE_REGENERATION                                           = 0x80380001
	STATUS_VOLMGR_INCOMPLETE_DISK_MIGRATION                                         = 0x80380002
	STATUS_VOLMGR_DATABASE_FULL                                                     = 0xC0380001
	STATUS_VOLMGR_DISK_CONFIGURATION_CORRUPTED                                      = 0xC0380002
	STATUS_VOLMGR_DISK_CONFIGURATION_NOT_IN_SYNC                                    = 0xC0380003
	STATUS_VOLMGR_PACK_CONFIG_UPDATE_FAILED                                         = 0xC0380004
	STATUS_VOLMGR_DISK_CONTAINS_NON_SIMPLE_VOLUME                                   = 0xC0380005
	STATUS_VOLMGR_DISK_DUPLICATE                                                    = 0xC0380006
	STATUS_VOLMGR_DISK_DYNAMIC                                                      = 0xC0380007
	STATUS_VOLMGR_DISK_ID_INVALID                                                   = 0xC0380008
	STATUS_VOLMGR_DISK_INVALID                                                      = 0xC0380009
	STATUS_VOLMGR_DISK_LAST_VOTER                                                   = 0xC038000A
	STATUS_VOLMGR_DISK_LAYOUT_INVALID                                               = 0xC038000B
	STATUS_VOLMGR_DISK_LAYOUT_NON_BASIC_BETWEEN_BASIC_PARTITIONS                    = 0xC038000C
	STATUS_VOLMGR_DISK_LAYOUT_NOT_CYLINDER_ALIGNED                                  = 0xC038000D
	STATUS_VOLMGR_DISK_LAYOUT_PARTITIONS_TOO_SMALL                                  = 0xC038000E
	STATUS_VOLMGR_DISK_LAYOUT_PRIMARY_BETWEEN_LOGICAL_PARTITIONS                    = 0xC038000F
	STATUS_VOLMGR_DISK_LAYOUT_TOO_MANY_PARTITIONS                                   = 0xC0380010
	STATUS_VOLMGR_DISK_MISSING                                                      = 0xC0380011
	STATUS_VOLMGR_DISK_NOT_EMPTY                                                    = 0xC0380012
	STATUS_VOLMGR_DISK_NOT_ENOUGH_SPACE                                             = 0xC0380013
	STATUS_VOLMGR_DISK_REVECTORING_FAILED                                           = 0xC0380014
	STATUS_VOLMGR_DISK_SECTOR_SIZE_INVALID                                          = 0xC0380015
	STATUS_VOLMGR_DISK_SET_NOT_CONTAINED                                            = 0xC0380016
	STATUS_VOLMGR_DISK_USED_BY_MULTIPLE_MEMBERS                                     = 0xC0380017
	STATUS_VOLMGR_DISK_USED_BY_MULTIPLE_PLEXES                                      = 0xC0380018
	STATUS_VOLMGR_DYNAMIC_DISK_NOT_SUPPORTED                                        = 0xC0380019
	STATUS_VOLMGR_EXTENT_ALREADY_USED                                               = 0xC038001A
	STATUS_VOLMGR_EXTENT_NOT_CONTIGUOUS                                             = 0xC038001B
	STATUS_VOLMGR_EXTENT_NOT_IN_PUBLIC_REGION                                       = 0xC038001C
	STATUS_VOLMGR_EXTENT_NOT_SECTOR_ALIGNED                                         = 0xC038001D
	STATUS_VOLMGR_EXTENT_OVERLAPS_EBR_PARTITION                                     = 0xC038001E
	STATUS_VOLMGR_EXTENT_VOLUME_LENGTHS_DO_NOT_MATCH                                = 0xC038001F
	STATUS_VOLMGR_FAULT_TOLERANT_NOT_SUPPORTED                                      = 0xC0380020
	STATUS_VOLMGR_INTERLEAVE_LENGTH_INVALID                                         = 0xC0380021
	STATUS_VOLMGR_MAXIMUM_REGISTERED_USERS                                          = 0xC0380022
	STATUS_VOLMGR_MEMBER_IN_SYNC                                                    = 0xC0380023
	STATUS_VOLMGR_MEMBER_INDEX_DUPLICATE                                            = 0xC0380024
	STATUS_VOLMGR_MEMBER_INDEX_INVALID                                              = 0xC0380025
	STATUS_VOLMGR_MEMBER_MISSING                                                    = 0xC0380026
	STATUS_VOLMGR_MEMBER_NOT_DETACHED                                               = 0xC0380027
	STATUS_VOLMGR_MEMBER_REGENERATING                                               = 0xC0380028
	STATUS_VOLMGR_ALL_DISKS_FAILED                                                  = 0xC0380029
	STATUS_VOLMGR_NO_REGISTERED_USERS                                               = 0xC038002A
	STATUS_VOLMGR_NO_SUCH_USER                                                      = 0xC038002B
	STATUS_VOLMGR_NOTIFICATION_RESET                                                = 0xC038002C
	STATUS_VOLMGR_NUMBER_OF_MEMBERS_INVALID                                         = 0xC038002D
	STATUS_VOLMGR_NUMBER_OF_PLEXES_INVALID                                          = 0xC038002E
	STATUS_VOLMGR_PACK_DUPLICATE                                                    = 0xC038002F
	STATUS_VOLMGR_PACK_ID_INVALID                                                   = 0xC0380030
	STATUS_VOLMGR_PACK_INVALID                                                      = 0xC0380031
	STATUS_VOLMGR_PACK_NAME_INVALID                                                 = 0xC0380032
	STATUS_VOLMGR_PACK_OFFLINE                                                      = 0xC0380033
	STATUS_VOLMGR_PACK_HAS_QUORUM                                                   = 0xC0380034
	STATUS_VOLMGR_PACK_WITHOUT_QUORUM                                               = 0xC0380035
	STATUS_VOLMGR_PARTITION_STYLE_INVALID                                           = 0xC0380036
	STATUS_VOLMGR_PARTITION_UPDATE_FAILED                                           = 0xC0380037
	STATUS_VOLMGR_PLEX_IN_SYNC                                                      = 0xC0380038
	STATUS_VOLMGR_PLEX_INDEX_DUPLICATE                                              = 0xC0380039
	STATUS_VOLMGR_PLEX_INDEX_INVALID                                                = 0xC038003A
	STATUS_VOLMGR_PLEX_LAST_ACTIVE                                                  = 0xC038003B
	STATUS_VOLMGR_PLEX_MISSING                                                      = 0xC038003C
	STATUS_VOLMGR_PLEX_REGENERATING                                                 = 0xC038003D
	STATUS_VOLMGR_PLEX_TYPE_INVALID                                                 = 0xC038003E
	STATUS_VOLMGR_PLEX_NOT_RAID5                                                    = 0xC038003F
	STATUS_VOLMGR_PLEX_NOT_SIMPLE                                                   = 0xC0380040
	STATUS_VOLMGR_STRUCTURE_SIZE_INVALID                                            = 0xC0380041
	STATUS_VOLMGR_TOO_MANY_NOTIFICATION_REQUESTS                                    = 0xC0380042
	STATUS_VOLMGR_TRANSACTION_IN_PROGRESS                                           = 0xC0380043
	STATUS_VOLMGR_UNEXPECTED_DISK_LAYOUT_CHANGE                                     = 0xC0380044
	STATUS_VOLMGR_VOLUME_CONTAINS_MISSING_DISK                                      = 0xC0380045
	STATUS_VOLMGR_VOLUME_ID_INVALID                                                 = 0xC0380046
	STATUS_VOLMGR_VOLUME_LENGTH_INVALID                                             = 0xC0380047
	STATUS_VOLMGR_VOLUME_LENGTH_NOT_SECTOR_SIZE_MULTIPLE                            = 0xC0380048
	STATUS_VOLMGR_VOLUME_NOT_MIRRORED                                               = 0xC0380049
	STATUS_VOLMGR_VOLUME_NOT_RETAINED                                               = 0xC038004A
	STATUS_VOLMGR_VOLUME_OFFLINE                                                    = 0xC038004B
	STATUS_VOLMGR_VOLUME_RETAINED                                                   = 0xC038004C
	STATUS_VOLMGR_NUMBER_OF_EXTENTS_INVALID                                         = 0xC038004D
	STATUS_VOLMGR_DIFFERENT_SECTOR_SIZE                                             = 0xC038004E
	STATUS_VOLMGR_BAD_BOOT_DISK                                                     = 0xC038004F
	STATUS_VOLMGR_PACK_CONFIG_OFFLINE                                               = 0xC0380050
	STATUS_VOLMGR_PACK_CONFIG_ONLINE                                                = 0xC0380051
	STATUS_VOLMGR_NOT_PRIMARY_PACK                                                  = 0xC0380052
	STATUS_VOLMGR_PACK_LOG_UPDATE_FAILED                                            = 0xC0380053
	STATUS_VOLMGR_NUMBER_OF_DISKS_IN_PLEX_INVALID                                   = 0xC0380054
	STATUS_VOLMGR_NUMBER_OF_DISKS_IN_MEMBER_INVALID                                 = 0xC0380055
	STATUS_VOLMGR_VOLUME_MIRRORED                                                   = 0xC0380056
	STATUS_VOLMGR_PLEX_NOT_SIMPLE_SPANNED                                           = 0xC0380057
	STATUS_VOLMGR_NO_VALID_LOG_COPIES                                               = 0xC0380058
	STATUS_VOLMGR_PRIMARY_PACK_PRESENT                                              = 0xC0380059
	STATUS_VOLMGR_NUMBER_OF_DISKS_INVALID                                           = 0xC038005A
	STATUS_VOLMGR_MIRROR_NOT_SUPPORTED                                              = 0xC038005B
	STATUS_VOLMGR_RAID5_NOT_SUPPORTED                                               = 0xC038005C
	STATUS_BCD_NOT_ALL_ENTRIES_IMPORTED                                             = 0x80390001
	STATUS_BCD_TOO_MANY_ELEMENTS                                                    = 0xC0390002
	STATUS_BCD_NOT_ALL_ENTRIES_SYNCHRONIZED                                         = 0x80390003
	STATUS_VHD_DRIVE_FOOTER_MISSING                                                 = 0xC03A0001
	STATUS_VHD_DRIVE_FOOTER_CHECKSUM_MISMATCH                                       = 0xC03A0002
	STATUS_VHD_DRIVE_FOOTER_CORRUPT                                                 = 0xC03A0003
	STATUS_VHD_FORMAT_UNKNOWN                                                       = 0xC03A0004
	STATUS_VHD_FORMAT_UNSUPPORTED_VERSION                                           = 0xC03A0005
	STATUS_VHD_SPARSE_HEADER_CHECKSUM_MISMATCH                                      = 0xC03A0006
	STATUS_VHD_SPARSE_HEADER_UNSUPPORTED_VERSION                                    = 0xC03A0007
	STATUS_VHD_SPARSE_HEADER_CORRUPT                                                = 0xC03A0008
	STATUS_VHD_BLOCK_ALLOCATION_FAILURE                                             = 0xC03A0009
	STATUS_VHD_BLOCK_ALLOCATION_TABLE_CORRUPT                                       = 0xC03A000A
	STATUS_VHD_INVALID_BLOCK_SIZE                                                   = 0xC03A000B
	STATUS_VHD_BITMAP_MISMATCH                                                      = 0xC03A000C
	STATUS_VHD_PARENT_VHD_NOT_FOUND                                                 = 0xC03A000D
	STATUS_VHD_CHILD_PARENT_ID_MISMATCH                                             = 0xC03A000E
	STATUS_VHD_CHILD_PARENT_TIMESTAMP_MISMATCH                                      = 0xC03A000F
	STATUS_VHD_METADATA_READ_FAILURE                                                = 0xC03A0010
	STATUS_VHD_METADATA_WRITE_FAILURE                                               = 0xC03A0011
	STATUS_VHD_INVALID_SIZE                                                         = 0xC03A0012
	STATUS_VHD_INVALID_FILE_SIZE                                                    = 0xC03A0013
	STATUS_VIRTDISK_PROVIDER_NOT_FOUND                                              = 0xC03A0014
	STATUS_VIRTDISK_NOT_VIRTUAL_DISK                                                = 0xC03A0015
	STATUS_VHD_PARENT_VHD_ACCESS_DENIED                                             = 0xC03A0016
	STATUS_VHD_CHILD_PARENT_SIZE_MISMATCH                                           = 0xC03A0017
	STATUS_VHD_DIFFERENCING_CHAIN_CYCLE_DETECTED                                    = 0xC03A0018
	STATUS_VHD_DIFFERENCING_CHAIN_ERROR_IN_PARENT                                   = 0xC03A0019
	STATUS_VIRTUAL_DISK_LIMITATION                                                  = 0xC03A001A
	STATUS_VHD_INVALID_TYPE                                                         = 0xC03A001B
	STATUS_VHD_INVALID_STATE                                                        = 0xC03A001C
	STATUS_VIRTDISK_UNSUPPORTED_DISK_SECTOR_SIZE                                    = 0xC03A001D
	STATUS_QUERY_STORAGE_ERROR                                                      = 0x803A0001
	STATUS_DIS_NOT_PRESENT                                                          = 0xC03C0001
	STATUS_DIS_ATTRIBUTE_NOT_FOUND                                                  = 0xC03C0002
	STATUS_DIS_UNRECOGNIZED_ATTRIBUTE                                               = 0xC03C0003
	STATUS_DIS_PARTIAL_DATA                                                         = 0xC03C0004
)

func (k NtstatusKind) String() string {
	switch k {
	case STATUS_SUCCESS:
		return "0x00000000"
	case STATUS_WAIT_1:
		return "0x00000001"
	case STATUS_WAIT_2:
		return "0x00000002"
	case STATUS_WAIT_3:
		return "0x00000003"
	case STATUS_WAIT_63:
		return "0x0000003F"
	case STATUS_ABANDONED:
		return "0x00000080"
	case STATUS_ABANDONED_WAIT_63:
		return "0x000000BF"
	case STATUS_USER_APC:
		return "0x000000C0"
	case STATUS_KERNEL_APC:
		return "0x00000100"
	case STATUS_ALERTED:
		return "0x00000101"
	case STATUS_TIMEOUT:
		return "0x00000102"
	case STATUS_PENDING:
		return "0x00000103"
	case STATUS_REPARSE:
		return "0x00000104"
	case STATUS_MORE_ENTRIES:
		return "0x00000105"
	case STATUS_NOT_ALL_ASSIGNED:
		return "0x00000106"
	case STATUS_SOME_NOT_MAPPED:
		return "0x00000107"
	case STATUS_OPLOCK_BREAK_IN_PROGRESS:
		return "0x00000108"
	case STATUS_VOLUME_MOUNTED:
		return "0x00000109"
	case STATUS_RXACT_COMMITTED:
		return "0x0000010A"
	case STATUS_NOTIFY_CLEANUP:
		return "0x0000010B"
	case STATUS_NOTIFY_ENUM_DIR:
		return "0x0000010C"
	case STATUS_NO_QUOTAS_FOR_ACCOUNT:
		return "0x0000010D"
	case STATUS_PRIMARY_TRANSPORT_CONNECT_FAILED:
		return "0x0000010E"
	case STATUS_PAGE_FAULT_TRANSITION:
		return "0x00000110"
	case STATUS_PAGE_FAULT_DEMAND_ZERO:
		return "0x00000111"
	case STATUS_PAGE_FAULT_COPY_ON_WRITE:
		return "0x00000112"
	case STATUS_PAGE_FAULT_GUARD_PAGE:
		return "0x00000113"
	case STATUS_PAGE_FAULT_PAGING_FILE:
		return "0x00000114"
	case STATUS_CACHE_PAGE_LOCKED:
		return "0x00000115"
	case STATUS_CRASH_DUMP:
		return "0x00000116"
	case STATUS_BUFFER_ALL_ZEROS:
		return "0x00000117"
	case STATUS_REPARSE_OBJECT:
		return "0x00000118"
	case STATUS_RESOURCE_REQUIREMENTS_CHANGED:
		return "0x00000119"
	case STATUS_TRANSLATION_COMPLETE:
		return "0x00000120"
	case STATUS_DS_MEMBERSHIP_EVALUATED_LOCALLY:
		return "0x00000121"
	case STATUS_NOTHING_TO_TERMINATE:
		return "0x00000122"
	case STATUS_PROCESS_NOT_IN_JOB:
		return "0x00000123"
	case STATUS_PROCESS_IN_JOB:
		return "0x00000124"
	case STATUS_VOLSNAP_HIBERNATE_READY:
		return "0x00000125"
	case STATUS_FSFILTER_OP_COMPLETED_SUCCESSFULLY:
		return "0x00000126"
	case STATUS_INTERRUPT_VECTOR_ALREADY_CONNECTED:
		return "0x00000127"
	case STATUS_INTERRUPT_STILL_CONNECTED:
		return "0x00000128"
	case STATUS_PROCESS_CLONED:
		return "0x00000129"
	case STATUS_FILE_LOCKED_WITH_ONLY_READERS:
		return "0x0000012A"
	case STATUS_FILE_LOCKED_WITH_WRITERS:
		return "0x0000012B"
	case STATUS_RESOURCEMANAGER_READ_ONLY:
		return "0x00000202"
	case STATUS_RING_PREVIOUSLY_EMPTY:
		return "0x00000210"
	case STATUS_RING_PREVIOUSLY_FULL:
		return "0x00000211"
	case STATUS_RING_PREVIOUSLY_ABOVE_QUOTA:
		return "0x00000212"
	case STATUS_RING_NEWLY_EMPTY:
		return "0x00000213"
	case STATUS_RING_SIGNAL_OPPOSITE_ENDPOINT:
		return "0x00000214"
	case STATUS_OPLOCK_SWITCHED_TO_NEW_HANDLE:
		return "0x00000215"
	case STATUS_OPLOCK_HANDLE_CLOSED:
		return "0x00000216"
	case STATUS_WAIT_FOR_OPLOCK:
		return "0x00000367"
	case DBG_EXCEPTION_HANDLED:
		return "0x00010001"
	case DBG_CONTINUE:
		return "0x00010002"
	case STATUS_FLT_IO_COMPLETE:
		return "0x001C0001"
	case STATUS_DIS_ATTRIBUTE_BUILT:
		return "0x003C0001"
	case STATUS_OBJECT_NAME_EXISTS:
		return "0x40000000"
	case STATUS_THREAD_WAS_SUSPENDED:
		return "0x40000001"
	case STATUS_WORKING_SET_LIMIT_RANGE:
		return "0x40000002"
	case STATUS_IMAGE_NOT_AT_BASE:
		return "0x40000003"
	case STATUS_RXACT_STATE_CREATED:
		return "0x40000004"
	case STATUS_SEGMENT_NOTIFICATION:
		return "0x40000005"
	case STATUS_LOCAL_USER_SESSION_KEY:
		return "0x40000006"
	case STATUS_BAD_CURRENT_DIRECTORY:
		return "0x40000007"
	case STATUS_SERIAL_MORE_WRITES:
		return "0x40000008"
	case STATUS_REGISTRY_RECOVERED:
		return "0x40000009"
	case STATUS_FT_READ_RECOVERY_FROM_BACKUP:
		return "0x4000000A"
	case STATUS_FT_WRITE_RECOVERY:
		return "0x4000000B"
	case STATUS_SERIAL_COUNTER_TIMEOUT:
		return "0x4000000C"
	case STATUS_NULL_LM_PASSWORD:
		return "0x4000000D"
	case STATUS_IMAGE_MACHINE_TYPE_MISMATCH:
		return "0x4000000E"
	case STATUS_RECEIVE_PARTIAL:
		return "0x4000000F"
	case STATUS_RECEIVE_EXPEDITED:
		return "0x40000010"
	case STATUS_RECEIVE_PARTIAL_EXPEDITED:
		return "0x40000011"
	case STATUS_EVENT_DONE:
		return "0x40000012"
	case STATUS_EVENT_PENDING:
		return "0x40000013"
	case STATUS_CHECKING_FILE_SYSTEM:
		return "0x40000014"
	case STATUS_FATAL_APP_EXIT:
		return "0x40000015"
	case STATUS_PREDEFINED_HANDLE:
		return "0x40000016"
	case STATUS_WAS_UNLOCKED:
		return "0x40000017"
	case STATUS_SERVICE_NOTIFICATION:
		return "0x40000018"
	case STATUS_WAS_LOCKED:
		return "0x40000019"
	case STATUS_LOG_HARD_ERROR:
		return "0x4000001A"
	case STATUS_ALREADY_WIN32:
		return "0x4000001B"
	case STATUS_WX86_UNSIMULATE:
		return "0x4000001C"
	case STATUS_WX86_CONTINUE:
		return "0x4000001D"
	case STATUS_WX86_SINGLE_STEP:
		return "0x4000001E"
	case STATUS_WX86_BREAKPOINT:
		return "0x4000001F"
	case STATUS_WX86_EXCEPTION_CONTINUE:
		return "0x40000020"
	case STATUS_WX86_EXCEPTION_LASTCHANCE:
		return "0x40000021"
	case STATUS_WX86_EXCEPTION_CHAIN:
		return "0x40000022"
	case STATUS_IMAGE_MACHINE_TYPE_MISMATCH_EXE:
		return "0x40000023"
	case STATUS_NO_YIELD_PERFORMED:
		return "0x40000024"
	case STATUS_TIMER_RESUME_IGNORED:
		return "0x40000025"
	case STATUS_ARBITRATION_UNHANDLED:
		return "0x40000026"
	case STATUS_CARDBUS_NOT_SUPPORTED:
		return "0x40000027"
	case STATUS_WX86_CREATEWX86TIB:
		return "0x40000028"
	case STATUS_MP_PROCESSOR_MISMATCH:
		return "0x40000029"
	case STATUS_HIBERNATED:
		return "0x4000002A"
	case STATUS_RESUME_HIBERNATION:
		return "0x4000002B"
	case STATUS_FIRMWARE_UPDATED:
		return "0x4000002C"
	case STATUS_DRIVERS_LEAKING_LOCKED_PAGES:
		return "0x4000002D"
	case STATUS_MESSAGE_RETRIEVED:
		return "0x4000002E"
	case STATUS_SYSTEM_POWERSTATE_TRANSITION:
		return "0x4000002F"
	case STATUS_ALPC_CHECK_COMPLETION_LIST:
		return "0x40000030"
	case STATUS_SYSTEM_POWERSTATE_COMPLEX_TRANSITION:
		return "0x40000031"
	case STATUS_ACCESS_AUDIT_BY_POLICY:
		return "0x40000032"
	case STATUS_ABANDON_HIBERFILE:
		return "0x40000033"
	case STATUS_BIZRULES_NOT_ENABLED:
		return "0x40000034"
	case DBG_REPLY_LATER:
		return "0x40010001"
	case DBG_UNABLE_TO_PROVIDE_HANDLE:
		return "0x40010002"
	case DBG_TERMINATE_THREAD:
		return "0x40010003"
	case DBG_TERMINATE_PROCESS:
		return "0x40010004"
	case DBG_CONTROL_C:
		return "0x40010005"
	case DBG_PRINTEXCEPTION_C:
		return "0x40010006"
	case DBG_RIPEXCEPTION:
		return "0x40010007"
	case DBG_CONTROL_BREAK:
		return "0x40010008"
	case DBG_COMMAND_EXCEPTION:
		return "0x40010009"
	case STATUS_HEURISTIC_DAMAGE_POSSIBLE:
		return "0x40190001"
	case STATUS_GUARD_PAGE_VIOLATION:
		return "0x80000001"
	case STATUS_DATATYPE_MISALIGNMENT:
		return "0x80000002"
	case STATUS_BREAKPOINT:
		return "0x80000003"
	case STATUS_SINGLE_STEP:
		return "0x80000004"
	case STATUS_BUFFER_OVERFLOW:
		return "0x80000005"
	case STATUS_NO_MORE_FILES:
		return "0x80000006"
	case STATUS_WAKE_SYSTEM_DEBUGGER:
		return "0x80000007"
	case STATUS_HANDLES_CLOSED:
		return "0x8000000A"
	case STATUS_NO_INHERITANCE:
		return "0x8000000B"
	case STATUS_GUID_SUBSTITUTION_MADE:
		return "0x8000000C"
	case STATUS_PARTIAL_COPY:
		return "0x8000000D"
	case STATUS_DEVICE_PAPER_EMPTY:
		return "0x8000000E"
	case STATUS_DEVICE_POWERED_OFF:
		return "0x8000000F"
	case STATUS_DEVICE_OFF_LINE:
		return "0x80000010"
	case STATUS_DEVICE_BUSY:
		return "0x80000011"
	case STATUS_NO_MORE_EAS:
		return "0x80000012"
	case STATUS_INVALID_EA_NAME:
		return "0x80000013"
	case STATUS_EA_LIST_INCONSISTENT:
		return "0x80000014"
	case STATUS_INVALID_EA_FLAG:
		return "0x80000015"
	case STATUS_VERIFY_REQUIRED:
		return "0x80000016"
	case STATUS_EXTRANEOUS_INFORMATION:
		return "0x80000017"
	case STATUS_RXACT_COMMIT_NECESSARY:
		return "0x80000018"
	case STATUS_NO_MORE_ENTRIES:
		return "0x8000001A"
	case STATUS_FILEMARK_DETECTED:
		return "0x8000001B"
	case STATUS_MEDIA_CHANGED:
		return "0x8000001C"
	case STATUS_BUS_RESET:
		return "0x8000001D"
	case STATUS_END_OF_MEDIA:
		return "0x8000001E"
	case STATUS_BEGINNING_OF_MEDIA:
		return "0x8000001F"
	case STATUS_MEDIA_CHECK:
		return "0x80000020"
	case STATUS_SETMARK_DETECTED:
		return "0x80000021"
	case STATUS_NO_DATA_DETECTED:
		return "0x80000022"
	case STATUS_REDIRECTOR_HAS_OPEN_HANDLES:
		return "0x80000023"
	case STATUS_SERVER_HAS_OPEN_HANDLES:
		return "0x80000024"
	case STATUS_ALREADY_DISCONNECTED:
		return "0x80000025"
	case STATUS_LONGJUMP:
		return "0x80000026"
	case STATUS_CLEANER_CARTRIDGE_INSTALLED:
		return "0x80000027"
	case STATUS_PLUGPLAY_QUERY_VETOED:
		return "0x80000028"
	case STATUS_UNWIND_CONSOLIDATE:
		return "0x80000029"
	case STATUS_REGISTRY_HIVE_RECOVERED:
		return "0x8000002A"
	case STATUS_DLL_MIGHT_BE_INSECURE:
		return "0x8000002B"
	case STATUS_DLL_MIGHT_BE_INCOMPATIBLE:
		return "0x8000002C"
	case STATUS_STOPPED_ON_SYMLINK:
		return "0x8000002D"
	case STATUS_CANNOT_GRANT_REQUESTED_OPLOCK:
		return "0x8000002E"
	case STATUS_NO_ACE_CONDITION:
		return "0x8000002F"
	case DBG_EXCEPTION_NOT_HANDLED:
		return "0x80010001"
	case STATUS_CLUSTER_NODE_ALREADY_UP:
		return "0x80130001"
	case STATUS_CLUSTER_NODE_ALREADY_DOWN:
		return "0x80130002"
	case STATUS_CLUSTER_NETWORK_ALREADY_ONLINE:
		return "0x80130003"
	case STATUS_CLUSTER_NETWORK_ALREADY_OFFLINE:
		return "0x80130004"
	case STATUS_CLUSTER_NODE_ALREADY_MEMBER:
		return "0x80130005"
	case STATUS_FLT_BUFFER_TOO_SMALL:
		return "0x801C0001"
	case STATUS_FVE_PARTIAL_METADATA:
		return "0x80210001"
	case STATUS_FVE_TRANSIENT_STATE:
		return "0x80210002"
	case STATUS_UNSUCCESSFUL:
		return "0xC0000001"
	case STATUS_NOT_IMPLEMENTED:
		return "0xC0000002"
	case STATUS_INVALID_INFO_CLASS:
		return "0xC0000003"
	case STATUS_INFO_LENGTH_MISMATCH:
		return "0xC0000004"
	case STATUS_ACCESS_VIOLATION:
		return "0xC0000005"
	case STATUS_IN_PAGE_ERROR:
		return "0xC0000006"
	case STATUS_PAGEFILE_QUOTA:
		return "0xC0000007"
	case STATUS_INVALID_HANDLE:
		return "0xC0000008"
	case STATUS_BAD_INITIAL_STACK:
		return "0xC0000009"
	case STATUS_BAD_INITIAL_PC:
		return "0xC000000A"
	case STATUS_INVALID_CID:
		return "0xC000000B"
	case STATUS_TIMER_NOT_CANCELED:
		return "0xC000000C"
	case STATUS_INVALID_PARAMETER:
		return "0xC000000D"
	case STATUS_NO_SUCH_DEVICE:
		return "0xC000000E"
	case STATUS_NO_SUCH_FILE:
		return "0xC000000F"
	case STATUS_INVALID_DEVICE_REQUEST:
		return "0xC0000010"
	case STATUS_END_OF_FILE:
		return "0xC0000011"
	case STATUS_WRONG_VOLUME:
		return "0xC0000012"
	case STATUS_NO_MEDIA_IN_DEVICE:
		return "0xC0000013"
	case STATUS_UNRECOGNIZED_MEDIA:
		return "0xC0000014"
	case STATUS_NONEXISTENT_SECTOR:
		return "0xC0000015"
	case STATUS_MORE_PROCESSING_REQUIRED:
		return "0xC0000016"
	case STATUS_NO_MEMORY:
		return "0xC0000017"
	case STATUS_CONFLICTING_ADDRESSES:
		return "0xC0000018"
	case STATUS_NOT_MAPPED_VIEW:
		return "0xC0000019"
	case STATUS_UNABLE_TO_FREE_VM:
		return "0xC000001A"
	case STATUS_UNABLE_TO_DELETE_SECTION:
		return "0xC000001B"
	case STATUS_INVALID_SYSTEM_SERVICE:
		return "0xC000001C"
	case STATUS_ILLEGAL_INSTRUCTION:
		return "0xC000001D"
	case STATUS_INVALID_LOCK_SEQUENCE:
		return "0xC000001E"
	case STATUS_INVALID_VIEW_SIZE:
		return "0xC000001F"
	case STATUS_INVALID_FILE_FOR_SECTION:
		return "0xC0000020"
	case STATUS_ALREADY_COMMITTED:
		return "0xC0000021"
	case STATUS_ACCESS_DENIED:
		return "0xC0000022"
	case STATUS_BUFFER_TOO_SMALL:
		return "0xC0000023"
	case STATUS_OBJECT_TYPE_MISMATCH:
		return "0xC0000024"
	case STATUS_NONCONTINUABLE_EXCEPTION:
		return "0xC0000025"
	case STATUS_INVALID_DISPOSITION:
		return "0xC0000026"
	case STATUS_UNWIND:
		return "0xC0000027"
	case STATUS_BAD_STACK:
		return "0xC0000028"
	case STATUS_INVALID_UNWIND_TARGET:
		return "0xC0000029"
	case STATUS_NOT_LOCKED:
		return "0xC000002A"
	case STATUS_PARITY_ERROR:
		return "0xC000002B"
	case STATUS_UNABLE_TO_DECOMMIT_VM:
		return "0xC000002C"
	case STATUS_NOT_COMMITTED:
		return "0xC000002D"
	case STATUS_INVALID_PORT_ATTRIBUTES:
		return "0xC000002E"
	case STATUS_PORT_MESSAGE_TOO_LONG:
		return "0xC000002F"
	case STATUS_INVALID_PARAMETER_MIX:
		return "0xC0000030"
	case STATUS_INVALID_QUOTA_LOWER:
		return "0xC0000031"
	case STATUS_DISK_CORRUPT_ERROR:
		return "0xC0000032"
	case STATUS_OBJECT_NAME_INVALID:
		return "0xC0000033"
	case STATUS_OBJECT_NAME_NOT_FOUND:
		return "0xC0000034"
	case STATUS_OBJECT_NAME_COLLISION:
		return "0xC0000035"
	case STATUS_PORT_DISCONNECTED:
		return "0xC0000037"
	case STATUS_DEVICE_ALREADY_ATTACHED:
		return "0xC0000038"
	case STATUS_OBJECT_PATH_INVALID:
		return "0xC0000039"
	case STATUS_OBJECT_PATH_NOT_FOUND:
		return "0xC000003A"
	case STATUS_OBJECT_PATH_SYNTAX_BAD:
		return "0xC000003B"
	case STATUS_DATA_OVERRUN:
		return "0xC000003C"
	case STATUS_DATA_LATE_ERROR:
		return "0xC000003D"
	case STATUS_DATA_ERROR:
		return "0xC000003E"
	case STATUS_CRC_ERROR:
		return "0xC000003F"
	case STATUS_SECTION_TOO_BIG:
		return "0xC0000040"
	case STATUS_PORT_CONNECTION_REFUSED:
		return "0xC0000041"
	case STATUS_INVALID_PORT_HANDLE:
		return "0xC0000042"
	case STATUS_SHARING_VIOLATION:
		return "0xC0000043"
	case STATUS_QUOTA_EXCEEDED:
		return "0xC0000044"
	case STATUS_INVALID_PAGE_PROTECTION:
		return "0xC0000045"
	case STATUS_MUTANT_NOT_OWNED:
		return "0xC0000046"
	case STATUS_SEMAPHORE_LIMIT_EXCEEDED:
		return "0xC0000047"
	case STATUS_PORT_ALREADY_SET:
		return "0xC0000048"
	case STATUS_SECTION_NOT_IMAGE:
		return "0xC0000049"
	case STATUS_SUSPEND_COUNT_EXCEEDED:
		return "0xC000004A"
	case STATUS_THREAD_IS_TERMINATING:
		return "0xC000004B"
	case STATUS_BAD_WORKING_SET_LIMIT:
		return "0xC000004C"
	case STATUS_INCOMPATIBLE_FILE_MAP:
		return "0xC000004D"
	case STATUS_SECTION_PROTECTION:
		return "0xC000004E"
	case STATUS_EAS_NOT_SUPPORTED:
		return "0xC000004F"
	case STATUS_EA_TOO_LARGE:
		return "0xC0000050"
	case STATUS_NONEXISTENT_EA_ENTRY:
		return "0xC0000051"
	case STATUS_NO_EAS_ON_FILE:
		return "0xC0000052"
	case STATUS_EA_CORRUPT_ERROR:
		return "0xC0000053"
	case STATUS_FILE_LOCK_CONFLICT:
		return "0xC0000054"
	case STATUS_LOCK_NOT_GRANTED:
		return "0xC0000055"
	case STATUS_DELETE_PENDING:
		return "0xC0000056"
	case STATUS_CTL_FILE_NOT_SUPPORTED:
		return "0xC0000057"
	case STATUS_UNKNOWN_REVISION:
		return "0xC0000058"
	case STATUS_REVISION_MISMATCH:
		return "0xC0000059"
	case STATUS_INVALID_OWNER:
		return "0xC000005A"
	case STATUS_INVALID_PRIMARY_GROUP:
		return "0xC000005B"
	case STATUS_NO_IMPERSONATION_TOKEN:
		return "0xC000005C"
	case STATUS_CANT_DISABLE_MANDATORY:
		return "0xC000005D"
	case STATUS_NO_LOGON_SERVERS:
		return "0xC000005E"
	case STATUS_NO_SUCH_LOGON_SESSION:
		return "0xC000005F"
	case STATUS_NO_SUCH_PRIVILEGE:
		return "0xC0000060"
	case STATUS_PRIVILEGE_NOT_HELD:
		return "0xC0000061"
	case STATUS_INVALID_ACCOUNT_NAME:
		return "0xC0000062"
	case STATUS_USER_EXISTS:
		return "0xC0000063"
	case STATUS_NO_SUCH_USER:
		return "0xC0000064"
	case STATUS_GROUP_EXISTS:
		return "0xC0000065"
	case STATUS_NO_SUCH_GROUP:
		return "0xC0000066"
	case STATUS_MEMBER_IN_GROUP:
		return "0xC0000067"
	case STATUS_MEMBER_NOT_IN_GROUP:
		return "0xC0000068"
	case STATUS_LAST_ADMIN:
		return "0xC0000069"
	case STATUS_WRONG_PASSWORD:
		return "0xC000006A"
	case STATUS_ILL_FORMED_PASSWORD:
		return "0xC000006B"
	case STATUS_PASSWORD_RESTRICTION:
		return "0xC000006C"
	case STATUS_LOGON_FAILURE:
		return "0xC000006D"
	case STATUS_ACCOUNT_RESTRICTION:
		return "0xC000006E"
	case STATUS_INVALID_LOGON_HOURS:
		return "0xC000006F"
	case STATUS_INVALID_WORKSTATION:
		return "0xC0000070"
	case STATUS_PASSWORD_EXPIRED:
		return "0xC0000071"
	case STATUS_ACCOUNT_DISABLED:
		return "0xC0000072"
	case STATUS_NONE_MAPPED:
		return "0xC0000073"
	case STATUS_TOO_MANY_LUIDS_REQUESTED:
		return "0xC0000074"
	case STATUS_LUIDS_EXHAUSTED:
		return "0xC0000075"
	case STATUS_INVALID_SUB_AUTHORITY:
		return "0xC0000076"
	case STATUS_INVALID_ACL:
		return "0xC0000077"
	case STATUS_INVALID_SID:
		return "0xC0000078"
	case STATUS_INVALID_SECURITY_DESCR:
		return "0xC0000079"
	case STATUS_PROCEDURE_NOT_FOUND:
		return "0xC000007A"
	case STATUS_INVALID_IMAGE_FORMAT:
		return "0xC000007B"
	case STATUS_NO_TOKEN:
		return "0xC000007C"
	case STATUS_BAD_INHERITANCE_ACL:
		return "0xC000007D"
	case STATUS_RANGE_NOT_LOCKED:
		return "0xC000007E"
	case STATUS_DISK_FULL:
		return "0xC000007F"
	case STATUS_SERVER_DISABLED:
		return "0xC0000080"
	case STATUS_SERVER_NOT_DISABLED:
		return "0xC0000081"
	case STATUS_TOO_MANY_GUIDS_REQUESTED:
		return "0xC0000082"
	case STATUS_GUIDS_EXHAUSTED:
		return "0xC0000083"
	case STATUS_INVALID_ID_AUTHORITY:
		return "0xC0000084"
	case STATUS_AGENTS_EXHAUSTED:
		return "0xC0000085"
	case STATUS_INVALID_VOLUME_LABEL:
		return "0xC0000086"
	case STATUS_SECTION_NOT_EXTENDED:
		return "0xC0000087"
	case STATUS_NOT_MAPPED_DATA:
		return "0xC0000088"
	case STATUS_RESOURCE_DATA_NOT_FOUND:
		return "0xC0000089"
	case STATUS_RESOURCE_TYPE_NOT_FOUND:
		return "0xC000008A"
	case STATUS_RESOURCE_NAME_NOT_FOUND:
		return "0xC000008B"
	case STATUS_ARRAY_BOUNDS_EXCEEDED:
		return "0xC000008C"
	case STATUS_FLOAT_DENORMAL_OPERAND:
		return "0xC000008D"
	case STATUS_FLOAT_DIVIDE_BY_ZERO:
		return "0xC000008E"
	case STATUS_FLOAT_INEXACT_RESULT:
		return "0xC000008F"
	case STATUS_FLOAT_INVALID_OPERATION:
		return "0xC0000090"
	case STATUS_FLOAT_OVERFLOW:
		return "0xC0000091"
	case STATUS_FLOAT_STACK_CHECK:
		return "0xC0000092"
	case STATUS_FLOAT_UNDERFLOW:
		return "0xC0000093"
	case STATUS_INTEGER_DIVIDE_BY_ZERO:
		return "0xC0000094"
	case STATUS_INTEGER_OVERFLOW:
		return "0xC0000095"
	case STATUS_PRIVILEGED_INSTRUCTION:
		return "0xC0000096"
	case STATUS_TOO_MANY_PAGING_FILES:
		return "0xC0000097"
	case STATUS_FILE_INVALID:
		return "0xC0000098"
	case STATUS_ALLOTTED_SPACE_EXCEEDED:
		return "0xC0000099"
	case STATUS_INSUFFICIENT_RESOURCES:
		return "0xC000009A"
	case STATUS_DFS_EXIT_PATH_FOUND:
		return "0xC000009B"
	case STATUS_DEVICE_DATA_ERROR:
		return "0xC000009C"
	case STATUS_DEVICE_NOT_CONNECTED:
		return "0xC000009D"
	case STATUS_DEVICE_POWER_FAILURE:
		return "0xC000009E"
	case STATUS_FREE_VM_NOT_AT_BASE:
		return "0xC000009F"
	case STATUS_MEMORY_NOT_ALLOCATED:
		return "0xC00000A0"
	case STATUS_WORKING_SET_QUOTA:
		return "0xC00000A1"
	case STATUS_MEDIA_WRITE_PROTECTED:
		return "0xC00000A2"
	case STATUS_DEVICE_NOT_READY:
		return "0xC00000A3"
	case STATUS_INVALID_GROUP_ATTRIBUTES:
		return "0xC00000A4"
	case STATUS_BAD_IMPERSONATION_LEVEL:
		return "0xC00000A5"
	case STATUS_CANT_OPEN_ANONYMOUS:
		return "0xC00000A6"
	case STATUS_BAD_VALIDATION_CLASS:
		return "0xC00000A7"
	case STATUS_BAD_TOKEN_TYPE:
		return "0xC00000A8"
	case STATUS_BAD_MASTER_BOOT_RECORD:
		return "0xC00000A9"
	case STATUS_INSTRUCTION_MISALIGNMENT:
		return "0xC00000AA"
	case STATUS_INSTANCE_NOT_AVAILABLE:
		return "0xC00000AB"
	case STATUS_PIPE_NOT_AVAILABLE:
		return "0xC00000AC"
	case STATUS_INVALID_PIPE_STATE:
		return "0xC00000AD"
	case STATUS_PIPE_BUSY:
		return "0xC00000AE"
	case STATUS_ILLEGAL_FUNCTION:
		return "0xC00000AF"
	case STATUS_PIPE_DISCONNECTED:
		return "0xC00000B0"
	case STATUS_PIPE_CLOSING:
		return "0xC00000B1"
	case STATUS_PIPE_CONNECTED:
		return "0xC00000B2"
	case STATUS_PIPE_LISTENING:
		return "0xC00000B3"
	case STATUS_INVALID_READ_MODE:
		return "0xC00000B4"
	case STATUS_IO_TIMEOUT:
		return "0xC00000B5"
	case STATUS_FILE_FORCED_CLOSED:
		return "0xC00000B6"
	case STATUS_PROFILING_NOT_STARTED:
		return "0xC00000B7"
	case STATUS_PROFILING_NOT_STOPPED:
		return "0xC00000B8"
	case STATUS_COULD_NOT_INTERPRET:
		return "0xC00000B9"
	case STATUS_FILE_IS_A_DIRECTORY:
		return "0xC00000BA"
	case STATUS_NOT_SUPPORTED:
		return "0xC00000BB"
	case STATUS_REMOTE_NOT_LISTENING:
		return "0xC00000BC"
	case STATUS_DUPLICATE_NAME:
		return "0xC00000BD"
	case STATUS_BAD_NETWORK_PATH:
		return "0xC00000BE"
	case STATUS_NETWORK_BUSY:
		return "0xC00000BF"
	case STATUS_DEVICE_DOES_NOT_EXIST:
		return "0xC00000C0"
	case STATUS_TOO_MANY_COMMANDS:
		return "0xC00000C1"
	case STATUS_ADAPTER_HARDWARE_ERROR:
		return "0xC00000C2"
	case STATUS_INVALID_NETWORK_RESPONSE:
		return "0xC00000C3"
	case STATUS_UNEXPECTED_NETWORK_ERROR:
		return "0xC00000C4"
	case STATUS_BAD_REMOTE_ADAPTER:
		return "0xC00000C5"
	case STATUS_PRINT_QUEUE_FULL:
		return "0xC00000C6"
	case STATUS_NO_SPOOL_SPACE:
		return "0xC00000C7"
	case STATUS_PRINT_CANCELLED:
		return "0xC00000C8"
	case STATUS_NETWORK_NAME_DELETED:
		return "0xC00000C9"
	case STATUS_NETWORK_ACCESS_DENIED:
		return "0xC00000CA"
	case STATUS_BAD_DEVICE_TYPE:
		return "0xC00000CB"
	case STATUS_BAD_NETWORK_NAME:
		return "0xC00000CC"
	case STATUS_TOO_MANY_NAMES:
		return "0xC00000CD"
	case STATUS_TOO_MANY_SESSIONS:
		return "0xC00000CE"
	case STATUS_SHARING_PAUSED:
		return "0xC00000CF"
	case STATUS_REQUEST_NOT_ACCEPTED:
		return "0xC00000D0"
	case STATUS_REDIRECTOR_PAUSED:
		return "0xC00000D1"
	case STATUS_NET_WRITE_FAULT:
		return "0xC00000D2"
	case STATUS_PROFILING_AT_LIMIT:
		return "0xC00000D3"
	case STATUS_NOT_SAME_DEVICE:
		return "0xC00000D4"
	case STATUS_FILE_RENAMED:
		return "0xC00000D5"
	case STATUS_VIRTUAL_CIRCUIT_CLOSED:
		return "0xC00000D6"
	case STATUS_NO_SECURITY_ON_OBJECT:
		return "0xC00000D7"
	case STATUS_CANT_WAIT:
		return "0xC00000D8"
	case STATUS_PIPE_EMPTY:
		return "0xC00000D9"
	case STATUS_CANT_ACCESS_DOMAIN_INFO:
		return "0xC00000DA"
	case STATUS_CANT_TERMINATE_SELF:
		return "0xC00000DB"
	case STATUS_INVALID_SERVER_STATE:
		return "0xC00000DC"
	case STATUS_INVALID_DOMAIN_STATE:
		return "0xC00000DD"
	case STATUS_INVALID_DOMAIN_ROLE:
		return "0xC00000DE"
	case STATUS_NO_SUCH_DOMAIN:
		return "0xC00000DF"
	case STATUS_DOMAIN_EXISTS:
		return "0xC00000E0"
	case STATUS_DOMAIN_LIMIT_EXCEEDED:
		return "0xC00000E1"
	case STATUS_OPLOCK_NOT_GRANTED:
		return "0xC00000E2"
	case STATUS_INVALID_OPLOCK_PROTOCOL:
		return "0xC00000E3"
	case STATUS_INTERNAL_DB_CORRUPTION:
		return "0xC00000E4"
	case STATUS_INTERNAL_ERROR:
		return "0xC00000E5"
	case STATUS_GENERIC_NOT_MAPPED:
		return "0xC00000E6"
	case STATUS_BAD_DESCRIPTOR_FORMAT:
		return "0xC00000E7"
	case STATUS_INVALID_USER_BUFFER:
		return "0xC00000E8"
	case STATUS_UNEXPECTED_IO_ERROR:
		return "0xC00000E9"
	case STATUS_UNEXPECTED_MM_CREATE_ERR:
		return "0xC00000EA"
	case STATUS_UNEXPECTED_MM_MAP_ERROR:
		return "0xC00000EB"
	case STATUS_UNEXPECTED_MM_EXTEND_ERR:
		return "0xC00000EC"
	case STATUS_NOT_LOGON_PROCESS:
		return "0xC00000ED"
	case STATUS_LOGON_SESSION_EXISTS:
		return "0xC00000EE"
	case STATUS_INVALID_PARAMETER_1:
		return "0xC00000EF"
	case STATUS_INVALID_PARAMETER_2:
		return "0xC00000F0"
	case STATUS_INVALID_PARAMETER_3:
		return "0xC00000F1"
	case STATUS_INVALID_PARAMETER_4:
		return "0xC00000F2"
	case STATUS_INVALID_PARAMETER_5:
		return "0xC00000F3"
	case STATUS_INVALID_PARAMETER_6:
		return "0xC00000F4"
	case STATUS_INVALID_PARAMETER_7:
		return "0xC00000F5"
	case STATUS_INVALID_PARAMETER_8:
		return "0xC00000F6"
	case STATUS_INVALID_PARAMETER_9:
		return "0xC00000F7"
	case STATUS_INVALID_PARAMETER_10:
		return "0xC00000F8"
	case STATUS_INVALID_PARAMETER_11:
		return "0xC00000F9"
	case STATUS_INVALID_PARAMETER_12:
		return "0xC00000FA"
	case STATUS_REDIRECTOR_NOT_STARTED:
		return "0xC00000FB"
	case STATUS_REDIRECTOR_STARTED:
		return "0xC00000FC"
	case STATUS_STACK_OVERFLOW:
		return "0xC00000FD"
	case STATUS_NO_SUCH_PACKAGE:
		return "0xC00000FE"
	case STATUS_BAD_FUNCTION_TABLE:
		return "0xC00000FF"
	case STATUS_VARIABLE_NOT_FOUND:
		return "0xC0000100"
	case STATUS_DIRECTORY_NOT_EMPTY:
		return "0xC0000101"
	case STATUS_FILE_CORRUPT_ERROR:
		return "0xC0000102"
	case STATUS_NOT_A_DIRECTORY:
		return "0xC0000103"
	case STATUS_BAD_LOGON_SESSION_STATE:
		return "0xC0000104"
	case STATUS_LOGON_SESSION_COLLISION:
		return "0xC0000105"
	case STATUS_NAME_TOO_LONG:
		return "0xC0000106"
	case STATUS_FILES_OPEN:
		return "0xC0000107"
	case STATUS_CONNECTION_IN_USE:
		return "0xC0000108"
	case STATUS_MESSAGE_NOT_FOUND:
		return "0xC0000109"
	case STATUS_PROCESS_IS_TERMINATING:
		return "0xC000010A"
	case STATUS_INVALID_LOGON_TYPE:
		return "0xC000010B"
	case STATUS_NO_GUID_TRANSLATION:
		return "0xC000010C"
	case STATUS_CANNOT_IMPERSONATE:
		return "0xC000010D"
	case STATUS_IMAGE_ALREADY_LOADED:
		return "0xC000010E"
	case STATUS_ABIOS_NOT_PRESENT:
		return "0xC000010F"
	case STATUS_ABIOS_LID_NOT_EXIST:
		return "0xC0000110"
	case STATUS_ABIOS_LID_ALREADY_OWNED:
		return "0xC0000111"
	case STATUS_ABIOS_NOT_LID_OWNER:
		return "0xC0000112"
	case STATUS_ABIOS_INVALID_COMMAND:
		return "0xC0000113"
	case STATUS_ABIOS_INVALID_LID:
		return "0xC0000114"
	case STATUS_ABIOS_SELECTOR_NOT_AVAILABLE:
		return "0xC0000115"
	case STATUS_ABIOS_INVALID_SELECTOR:
		return "0xC0000116"
	case STATUS_NO_LDT:
		return "0xC0000117"
	case STATUS_INVALID_LDT_SIZE:
		return "0xC0000118"
	case STATUS_INVALID_LDT_OFFSET:
		return "0xC0000119"
	case STATUS_INVALID_LDT_DESCRIPTOR:
		return "0xC000011A"
	case STATUS_INVALID_IMAGE_NE_FORMAT:
		return "0xC000011B"
	case STATUS_RXACT_INVALID_STATE:
		return "0xC000011C"
	case STATUS_RXACT_COMMIT_FAILURE:
		return "0xC000011D"
	case STATUS_MAPPED_FILE_SIZE_ZERO:
		return "0xC000011E"
	case STATUS_TOO_MANY_OPENED_FILES:
		return "0xC000011F"
	case STATUS_CANCELLED:
		return "0xC0000120"
	case STATUS_CANNOT_DELETE:
		return "0xC0000121"
	case STATUS_INVALID_COMPUTER_NAME:
		return "0xC0000122"
	case STATUS_FILE_DELETED:
		return "0xC0000123"
	case STATUS_SPECIAL_ACCOUNT:
		return "0xC0000124"
	case STATUS_SPECIAL_GROUP:
		return "0xC0000125"
	case STATUS_SPECIAL_USER:
		return "0xC0000126"
	case STATUS_MEMBERS_PRIMARY_GROUP:
		return "0xC0000127"
	case STATUS_FILE_CLOSED:
		return "0xC0000128"
	case STATUS_TOO_MANY_THREADS:
		return "0xC0000129"
	case STATUS_THREAD_NOT_IN_PROCESS:
		return "0xC000012A"
	case STATUS_TOKEN_ALREADY_IN_USE:
		return "0xC000012B"
	case STATUS_PAGEFILE_QUOTA_EXCEEDED:
		return "0xC000012C"
	case STATUS_COMMITMENT_LIMIT:
		return "0xC000012D"
	case STATUS_INVALID_IMAGE_LE_FORMAT:
		return "0xC000012E"
	case STATUS_INVALID_IMAGE_NOT_MZ:
		return "0xC000012F"
	case STATUS_INVALID_IMAGE_PROTECT:
		return "0xC0000130"
	case STATUS_INVALID_IMAGE_WIN_16:
		return "0xC0000131"
	case STATUS_LOGON_SERVER_CONFLICT:
		return "0xC0000132"
	case STATUS_TIME_DIFFERENCE_AT_DC:
		return "0xC0000133"
	case STATUS_SYNCHRONIZATION_REQUIRED:
		return "0xC0000134"
	case STATUS_DLL_NOT_FOUND:
		return "0xC0000135"
	case STATUS_OPEN_FAILED:
		return "0xC0000136"
	case STATUS_IO_PRIVILEGE_FAILED:
		return "0xC0000137"
	case STATUS_ORDINAL_NOT_FOUND:
		return "0xC0000138"
	case STATUS_ENTRYPOINT_NOT_FOUND:
		return "0xC0000139"
	case STATUS_CONTROL_C_EXIT:
		return "0xC000013A"
	case STATUS_LOCAL_DISCONNECT:
		return "0xC000013B"
	case STATUS_REMOTE_DISCONNECT:
		return "0xC000013C"
	case STATUS_REMOTE_RESOURCES:
		return "0xC000013D"
	case STATUS_LINK_FAILED:
		return "0xC000013E"
	case STATUS_LINK_TIMEOUT:
		return "0xC000013F"
	case STATUS_INVALID_CONNECTION:
		return "0xC0000140"
	case STATUS_INVALID_ADDRESS:
		return "0xC0000141"
	case STATUS_DLL_INIT_FAILED:
		return "0xC0000142"
	case STATUS_MISSING_SYSTEMFILE:
		return "0xC0000143"
	case STATUS_UNHANDLED_EXCEPTION:
		return "0xC0000144"
	case STATUS_APP_INIT_FAILURE:
		return "0xC0000145"
	case STATUS_PAGEFILE_CREATE_FAILED:
		return "0xC0000146"
	case STATUS_NO_PAGEFILE:
		return "0xC0000147"
	case STATUS_INVALID_LEVEL:
		return "0xC0000148"
	case STATUS_WRONG_PASSWORD_CORE:
		return "0xC0000149"
	case STATUS_ILLEGAL_FLOAT_CONTEXT:
		return "0xC000014A"
	case STATUS_PIPE_BROKEN:
		return "0xC000014B"
	case STATUS_REGISTRY_CORRUPT:
		return "0xC000014C"
	case STATUS_REGISTRY_IO_FAILED:
		return "0xC000014D"
	case STATUS_NO_EVENT_PAIR:
		return "0xC000014E"
	case STATUS_UNRECOGNIZED_VOLUME:
		return "0xC000014F"
	case STATUS_SERIAL_NO_DEVICE_INITED:
		return "0xC0000150"
	case STATUS_NO_SUCH_ALIAS:
		return "0xC0000151"
	case STATUS_MEMBER_NOT_IN_ALIAS:
		return "0xC0000152"
	case STATUS_MEMBER_IN_ALIAS:
		return "0xC0000153"
	case STATUS_ALIAS_EXISTS:
		return "0xC0000154"
	case STATUS_LOGON_NOT_GRANTED:
		return "0xC0000155"
	case STATUS_TOO_MANY_SECRETS:
		return "0xC0000156"
	case STATUS_SECRET_TOO_LONG:
		return "0xC0000157"
	case STATUS_INTERNAL_DB_ERROR:
		return "0xC0000158"
	case STATUS_FULLSCREEN_MODE:
		return "0xC0000159"
	case STATUS_TOO_MANY_CONTEXT_IDS:
		return "0xC000015A"
	case STATUS_LOGON_TYPE_NOT_GRANTED:
		return "0xC000015B"
	case STATUS_NOT_REGISTRY_FILE:
		return "0xC000015C"
	case STATUS_NT_CROSS_ENCRYPTION_REQUIRED:
		return "0xC000015D"
	case STATUS_DOMAIN_CTRLR_CONFIG_ERROR:
		return "0xC000015E"
	case STATUS_FT_MISSING_MEMBER:
		return "0xC000015F"
	case STATUS_ILL_FORMED_SERVICE_ENTRY:
		return "0xC0000160"
	case STATUS_ILLEGAL_CHARACTER:
		return "0xC0000161"
	case STATUS_UNMAPPABLE_CHARACTER:
		return "0xC0000162"
	case STATUS_UNDEFINED_CHARACTER:
		return "0xC0000163"
	case STATUS_FLOPPY_VOLUME:
		return "0xC0000164"
	case STATUS_FLOPPY_ID_MARK_NOT_FOUND:
		return "0xC0000165"
	case STATUS_FLOPPY_WRONG_CYLINDER:
		return "0xC0000166"
	case STATUS_FLOPPY_UNKNOWN_ERROR:
		return "0xC0000167"
	case STATUS_FLOPPY_BAD_REGISTERS:
		return "0xC0000168"
	case STATUS_DISK_RECALIBRATE_FAILED:
		return "0xC0000169"
	case STATUS_DISK_OPERATION_FAILED:
		return "0xC000016A"
	case STATUS_DISK_RESET_FAILED:
		return "0xC000016B"
	case STATUS_SHARED_IRQ_BUSY:
		return "0xC000016C"
	case STATUS_FT_ORPHANING:
		return "0xC000016D"
	case STATUS_BIOS_FAILED_TO_CONNECT_INTERRUPT:
		return "0xC000016E"
	case STATUS_PARTITION_FAILURE:
		return "0xC0000172"
	case STATUS_INVALID_BLOCK_LENGTH:
		return "0xC0000173"
	case STATUS_DEVICE_NOT_PARTITIONED:
		return "0xC0000174"
	case STATUS_UNABLE_TO_LOCK_MEDIA:
		return "0xC0000175"
	case STATUS_UNABLE_TO_UNLOAD_MEDIA:
		return "0xC0000176"
	case STATUS_EOM_OVERFLOW:
		return "0xC0000177"
	case STATUS_NO_MEDIA:
		return "0xC0000178"
	case STATUS_NO_SUCH_MEMBER:
		return "0xC000017A"
	case STATUS_INVALID_MEMBER:
		return "0xC000017B"
	case STATUS_KEY_DELETED:
		return "0xC000017C"
	case STATUS_NO_LOG_SPACE:
		return "0xC000017D"
	case STATUS_TOO_MANY_SIDS:
		return "0xC000017E"
	case STATUS_LM_CROSS_ENCRYPTION_REQUIRED:
		return "0xC000017F"
	case STATUS_KEY_HAS_CHILDREN:
		return "0xC0000180"
	case STATUS_CHILD_MUST_BE_VOLATILE:
		return "0xC0000181"
	case STATUS_DEVICE_CONFIGURATION_ERROR:
		return "0xC0000182"
	case STATUS_DRIVER_INTERNAL_ERROR:
		return "0xC0000183"
	case STATUS_INVALID_DEVICE_STATE:
		return "0xC0000184"
	case STATUS_IO_DEVICE_ERROR:
		return "0xC0000185"
	case STATUS_DEVICE_PROTOCOL_ERROR:
		return "0xC0000186"
	case STATUS_BACKUP_CONTROLLER:
		return "0xC0000187"
	case STATUS_LOG_FILE_FULL:
		return "0xC0000188"
	case STATUS_TOO_LATE:
		return "0xC0000189"
	case STATUS_NO_TRUST_LSA_SECRET:
		return "0xC000018A"
	case STATUS_NO_TRUST_SAM_ACCOUNT:
		return "0xC000018B"
	case STATUS_TRUSTED_DOMAIN_FAILURE:
		return "0xC000018C"
	case STATUS_TRUSTED_RELATIONSHIP_FAILURE:
		return "0xC000018D"
	case STATUS_EVENTLOG_FILE_CORRUPT:
		return "0xC000018E"
	case STATUS_EVENTLOG_CANT_START:
		return "0xC000018F"
	case STATUS_TRUST_FAILURE:
		return "0xC0000190"
	case STATUS_MUTANT_LIMIT_EXCEEDED:
		return "0xC0000191"
	case STATUS_NETLOGON_NOT_STARTED:
		return "0xC0000192"
	case STATUS_ACCOUNT_EXPIRED:
		return "0xC0000193"
	case STATUS_POSSIBLE_DEADLOCK:
		return "0xC0000194"
	case STATUS_NETWORK_CREDENTIAL_CONFLICT:
		return "0xC0000195"
	case STATUS_REMOTE_SESSION_LIMIT:
		return "0xC0000196"
	case STATUS_EVENTLOG_FILE_CHANGED:
		return "0xC0000197"
	case STATUS_NOLOGON_INTERDOMAIN_TRUST_ACCOUNT:
		return "0xC0000198"
	case STATUS_NOLOGON_WORKSTATION_TRUST_ACCOUNT:
		return "0xC0000199"
	case STATUS_NOLOGON_SERVER_TRUST_ACCOUNT:
		return "0xC000019A"
	case STATUS_DOMAIN_TRUST_INCONSISTENT:
		return "0xC000019B"
	case STATUS_FS_DRIVER_REQUIRED:
		return "0xC000019C"
	case STATUS_IMAGE_ALREADY_LOADED_AS_DLL:
		return "0xC000019D"
	case STATUS_INCOMPATIBLE_WITH_GLOBAL_SHORT_NAME_REGISTRY_SETTING:
		return "0xC000019E"
	case STATUS_SHORT_NAMES_NOT_ENABLED_ON_VOLUME:
		return "0xC000019F"
	case STATUS_SECURITY_STREAM_IS_INCONSISTENT:
		return "0xC00001A0"
	case STATUS_INVALID_LOCK_RANGE:
		return "0xC00001A1"
	case STATUS_INVALID_ACE_CONDITION:
		return "0xC00001A2"
	case STATUS_IMAGE_SUBSYSTEM_NOT_PRESENT:
		return "0xC00001A3"
	case STATUS_NOTIFICATION_GUID_ALREADY_DEFINED:
		return "0xC00001A4"
	case STATUS_NETWORK_OPEN_RESTRICTION:
		return "0xC0000201"
	case STATUS_NO_USER_SESSION_KEY:
		return "0xC0000202"
	case STATUS_USER_SESSION_DELETED:
		return "0xC0000203"
	case STATUS_RESOURCE_LANG_NOT_FOUND:
		return "0xC0000204"
	case STATUS_INSUFF_SERVER_RESOURCES:
		return "0xC0000205"
	case STATUS_INVALID_BUFFER_SIZE:
		return "0xC0000206"
	case STATUS_INVALID_ADDRESS_COMPONENT:
		return "0xC0000207"
	case STATUS_INVALID_ADDRESS_WILDCARD:
		return "0xC0000208"
	case STATUS_TOO_MANY_ADDRESSES:
		return "0xC0000209"
	case STATUS_ADDRESS_ALREADY_EXISTS:
		return "0xC000020A"
	case STATUS_ADDRESS_CLOSED:
		return "0xC000020B"
	case STATUS_CONNECTION_DISCONNECTED:
		return "0xC000020C"
	case STATUS_CONNECTION_RESET:
		return "0xC000020D"
	case STATUS_TOO_MANY_NODES:
		return "0xC000020E"
	case STATUS_TRANSACTION_ABORTED:
		return "0xC000020F"
	case STATUS_TRANSACTION_TIMED_OUT:
		return "0xC0000210"
	case STATUS_TRANSACTION_NO_RELEASE:
		return "0xC0000211"
	case STATUS_TRANSACTION_NO_MATCH:
		return "0xC0000212"
	case STATUS_TRANSACTION_RESPONDED:
		return "0xC0000213"
	case STATUS_TRANSACTION_INVALID_ID:
		return "0xC0000214"
	case STATUS_TRANSACTION_INVALID_TYPE:
		return "0xC0000215"
	case STATUS_NOT_SERVER_SESSION:
		return "0xC0000216"
	case STATUS_NOT_CLIENT_SESSION:
		return "0xC0000217"
	case STATUS_CANNOT_LOAD_REGISTRY_FILE:
		return "0xC0000218"
	case STATUS_DEBUG_ATTACH_FAILED:
		return "0xC0000219"
	case STATUS_SYSTEM_PROCESS_TERMINATED:
		return "0xC000021A"
	case STATUS_DATA_NOT_ACCEPTED:
		return "0xC000021B"
	case STATUS_NO_BROWSER_SERVERS_FOUND:
		return "0xC000021C"
	case STATUS_VDM_HARD_ERROR:
		return "0xC000021D"
	case STATUS_DRIVER_CANCEL_TIMEOUT:
		return "0xC000021E"
	case STATUS_REPLY_MESSAGE_MISMATCH:
		return "0xC000021F"
	case STATUS_MAPPED_ALIGNMENT:
		return "0xC0000220"
	case STATUS_IMAGE_CHECKSUM_MISMATCH:
		return "0xC0000221"
	case STATUS_LOST_WRITEBEHIND_DATA:
		return "0xC0000222"
	case STATUS_CLIENT_SERVER_PARAMETERS_INVALID:
		return "0xC0000223"
	case STATUS_PASSWORD_MUST_CHANGE:
		return "0xC0000224"
	case STATUS_NOT_FOUND:
		return "0xC0000225"
	case STATUS_NOT_TINY_STREAM:
		return "0xC0000226"
	case STATUS_RECOVERY_FAILURE:
		return "0xC0000227"
	case STATUS_STACK_OVERFLOW_READ:
		return "0xC0000228"
	case STATUS_FAIL_CHECK:
		return "0xC0000229"
	case STATUS_DUPLICATE_OBJECTID:
		return "0xC000022A"
	case STATUS_OBJECTID_EXISTS:
		return "0xC000022B"
	case STATUS_CONVERT_TO_LARGE:
		return "0xC000022C"
	case STATUS_RETRY:
		return "0xC000022D"
	case STATUS_FOUND_OUT_OF_SCOPE:
		return "0xC000022E"
	case STATUS_ALLOCATE_BUCKET:
		return "0xC000022F"
	case STATUS_PROPSET_NOT_FOUND:
		return "0xC0000230"
	case STATUS_MARSHALL_OVERFLOW:
		return "0xC0000231"
	case STATUS_INVALID_VARIANT:
		return "0xC0000232"
	case STATUS_DOMAIN_CONTROLLER_NOT_FOUND:
		return "0xC0000233"
	case STATUS_ACCOUNT_LOCKED_OUT:
		return "0xC0000234"
	case STATUS_HANDLE_NOT_CLOSABLE:
		return "0xC0000235"
	case STATUS_CONNECTION_REFUSED:
		return "0xC0000236"
	case STATUS_GRACEFUL_DISCONNECT:
		return "0xC0000237"
	case STATUS_ADDRESS_ALREADY_ASSOCIATED:
		return "0xC0000238"
	case STATUS_ADDRESS_NOT_ASSOCIATED:
		return "0xC0000239"
	case STATUS_CONNECTION_INVALID:
		return "0xC000023A"
	case STATUS_CONNECTION_ACTIVE:
		return "0xC000023B"
	case STATUS_NETWORK_UNREACHABLE:
		return "0xC000023C"
	case STATUS_HOST_UNREACHABLE:
		return "0xC000023D"
	case STATUS_PROTOCOL_UNREACHABLE:
		return "0xC000023E"
	case STATUS_PORT_UNREACHABLE:
		return "0xC000023F"
	case STATUS_REQUEST_ABORTED:
		return "0xC0000240"
	case STATUS_CONNECTION_ABORTED:
		return "0xC0000241"
	case STATUS_BAD_COMPRESSION_BUFFER:
		return "0xC0000242"
	case STATUS_USER_MAPPED_FILE:
		return "0xC0000243"
	case STATUS_AUDIT_FAILED:
		return "0xC0000244"
	case STATUS_TIMER_RESOLUTION_NOT_SET:
		return "0xC0000245"
	case STATUS_CONNECTION_COUNT_LIMIT:
		return "0xC0000246"
	case STATUS_LOGIN_TIME_RESTRICTION:
		return "0xC0000247"
	case STATUS_LOGIN_WKSTA_RESTRICTION:
		return "0xC0000248"
	case STATUS_IMAGE_MP_UP_MISMATCH:
		return "0xC0000249"
	case STATUS_INSUFFICIENT_LOGON_INFO:
		return "0xC0000250"
	case STATUS_BAD_DLL_ENTRYPOINT:
		return "0xC0000251"
	case STATUS_BAD_SERVICE_ENTRYPOINT:
		return "0xC0000252"
	case STATUS_LPC_REPLY_LOST:
		return "0xC0000253"
	case STATUS_IP_ADDRESS_CONFLICT1:
		return "0xC0000254"
	case STATUS_IP_ADDRESS_CONFLICT2:
		return "0xC0000255"
	case STATUS_REGISTRY_QUOTA_LIMIT:
		return "0xC0000256"
	case STATUS_PATH_NOT_COVERED:
		return "0xC0000257"
	case STATUS_NO_CALLBACK_ACTIVE:
		return "0xC0000258"
	case STATUS_LICENSE_QUOTA_EXCEEDED:
		return "0xC0000259"
	case STATUS_PWD_TOO_SHORT:
		return "0xC000025A"
	case STATUS_PWD_TOO_RECENT:
		return "0xC000025B"
	case STATUS_PWD_HISTORY_CONFLICT:
		return "0xC000025C"
	case STATUS_PLUGPLAY_NO_DEVICE:
		return "0xC000025E"
	case STATUS_UNSUPPORTED_COMPRESSION:
		return "0xC000025F"
	case STATUS_INVALID_HW_PROFILE:
		return "0xC0000260"
	case STATUS_INVALID_PLUGPLAY_DEVICE_PATH:
		return "0xC0000261"
	case STATUS_DRIVER_ORDINAL_NOT_FOUND:
		return "0xC0000262"
	case STATUS_DRIVER_ENTRYPOINT_NOT_FOUND:
		return "0xC0000263"
	case STATUS_RESOURCE_NOT_OWNED:
		return "0xC0000264"
	case STATUS_TOO_MANY_LINKS:
		return "0xC0000265"
	case STATUS_QUOTA_LIST_INCONSISTENT:
		return "0xC0000266"
	case STATUS_FILE_IS_OFFLINE:
		return "0xC0000267"
	case STATUS_EVALUATION_EXPIRATION:
		return "0xC0000268"
	case STATUS_ILLEGAL_DLL_RELOCATION:
		return "0xC0000269"
	case STATUS_LICENSE_VIOLATION:
		return "0xC000026A"
	case STATUS_DLL_INIT_FAILED_LOGOFF:
		return "0xC000026B"
	case STATUS_DRIVER_UNABLE_TO_LOAD:
		return "0xC000026C"
	case STATUS_DFS_UNAVAILABLE:
		return "0xC000026D"
	case STATUS_VOLUME_DISMOUNTED:
		return "0xC000026E"
	case STATUS_WX86_INTERNAL_ERROR:
		return "0xC000026F"
	case STATUS_WX86_FLOAT_STACK_CHECK:
		return "0xC0000270"
	case STATUS_VALIDATE_CONTINUE:
		return "0xC0000271"
	case STATUS_NO_MATCH:
		return "0xC0000272"
	case STATUS_NO_MORE_MATCHES:
		return "0xC0000273"
	case STATUS_NOT_A_REPARSE_POINT:
		return "0xC0000275"
	case STATUS_IO_REPARSE_TAG_INVALID:
		return "0xC0000276"
	case STATUS_IO_REPARSE_TAG_MISMATCH:
		return "0xC0000277"
	case STATUS_IO_REPARSE_DATA_INVALID:
		return "0xC0000278"
	case STATUS_IO_REPARSE_TAG_NOT_HANDLED:
		return "0xC0000279"
	case STATUS_REPARSE_POINT_NOT_RESOLVED:
		return "0xC0000280"
	case STATUS_DIRECTORY_IS_A_REPARSE_POINT:
		return "0xC0000281"
	case STATUS_RANGE_LIST_CONFLICT:
		return "0xC0000282"
	case STATUS_SOURCE_ELEMENT_EMPTY:
		return "0xC0000283"
	case STATUS_DESTINATION_ELEMENT_FULL:
		return "0xC0000284"
	case STATUS_ILLEGAL_ELEMENT_ADDRESS:
		return "0xC0000285"
	case STATUS_MAGAZINE_NOT_PRESENT:
		return "0xC0000286"
	case STATUS_REINITIALIZATION_NEEDED:
		return "0xC0000287"
	case STATUS_DEVICE_REQUIRES_CLEANING:
		return "0x80000288"
	case STATUS_DEVICE_DOOR_OPEN:
		return "0x80000289"
	case STATUS_ENCRYPTION_FAILED:
		return "0xC000028A"
	case STATUS_DECRYPTION_FAILED:
		return "0xC000028B"
	case STATUS_RANGE_NOT_FOUND:
		return "0xC000028C"
	case STATUS_NO_RECOVERY_POLICY:
		return "0xC000028D"
	case STATUS_NO_EFS:
		return "0xC000028E"
	case STATUS_WRONG_EFS:
		return "0xC000028F"
	case STATUS_NO_USER_KEYS:
		return "0xC0000290"
	case STATUS_FILE_NOT_ENCRYPTED:
		return "0xC0000291"
	case STATUS_NOT_EXPORT_FORMAT:
		return "0xC0000292"
	case STATUS_FILE_ENCRYPTED:
		return "0xC0000293"
	case STATUS_WAKE_SYSTEM:
		return "0x40000294"
	case STATUS_WMI_GUID_NOT_FOUND:
		return "0xC0000295"
	case STATUS_WMI_INSTANCE_NOT_FOUND:
		return "0xC0000296"
	case STATUS_WMI_ITEMID_NOT_FOUND:
		return "0xC0000297"
	case STATUS_WMI_TRY_AGAIN:
		return "0xC0000298"
	case STATUS_SHARED_POLICY:
		return "0xC0000299"
	case STATUS_POLICY_OBJECT_NOT_FOUND:
		return "0xC000029A"
	case STATUS_POLICY_ONLY_IN_DS:
		return "0xC000029B"
	case STATUS_VOLUME_NOT_UPGRADED:
		return "0xC000029C"
	case STATUS_REMOTE_STORAGE_NOT_ACTIVE:
		return "0xC000029D"
	case STATUS_REMOTE_STORAGE_MEDIA_ERROR:
		return "0xC000029E"
	case STATUS_NO_TRACKING_SERVICE:
		return "0xC000029F"
	case STATUS_SERVER_SID_MISMATCH:
		return "0xC00002A0"
	case STATUS_DS_NO_ATTRIBUTE_OR_VALUE:
		return "0xC00002A1"
	case STATUS_DS_INVALID_ATTRIBUTE_SYNTAX:
		return "0xC00002A2"
	case STATUS_DS_ATTRIBUTE_TYPE_UNDEFINED:
		return "0xC00002A3"
	case STATUS_DS_ATTRIBUTE_OR_VALUE_EXISTS:
		return "0xC00002A4"
	case STATUS_DS_BUSY:
		return "0xC00002A5"
	case STATUS_DS_UNAVAILABLE:
		return "0xC00002A6"
	case STATUS_DS_NO_RIDS_ALLOCATED:
		return "0xC00002A7"
	case STATUS_DS_NO_MORE_RIDS:
		return "0xC00002A8"
	case STATUS_DS_INCORRECT_ROLE_OWNER:
		return "0xC00002A9"
	case STATUS_DS_RIDMGR_INIT_ERROR:
		return "0xC00002AA"
	case STATUS_DS_OBJ_CLASS_VIOLATION:
		return "0xC00002AB"
	case STATUS_DS_CANT_ON_NON_LEAF:
		return "0xC00002AC"
	case STATUS_DS_CANT_ON_RDN:
		return "0xC00002AD"
	case STATUS_DS_CANT_MOD_OBJ_CLASS:
		return "0xC00002AE"
	case STATUS_DS_CROSS_DOM_MOVE_FAILED:
		return "0xC00002AF"
	case STATUS_DS_GC_NOT_AVAILABLE:
		return "0xC00002B0"
	case STATUS_DIRECTORY_SERVICE_REQUIRED:
		return "0xC00002B1"
	case STATUS_REPARSE_ATTRIBUTE_CONFLICT:
		return "0xC00002B2"
	case STATUS_CANT_ENABLE_DENY_ONLY:
		return "0xC00002B3"
	case STATUS_FLOAT_MULTIPLE_FAULTS:
		return "0xC00002B4"
	case STATUS_FLOAT_MULTIPLE_TRAPS:
		return "0xC00002B5"
	case STATUS_DEVICE_REMOVED:
		return "0xC00002B6"
	case STATUS_JOURNAL_DELETE_IN_PROGRESS:
		return "0xC00002B7"
	case STATUS_JOURNAL_NOT_ACTIVE:
		return "0xC00002B8"
	case STATUS_NOINTERFACE:
		return "0xC00002B9"
	case STATUS_DS_ADMIN_LIMIT_EXCEEDED:
		return "0xC00002C1"
	case STATUS_DRIVER_FAILED_SLEEP:
		return "0xC00002C2"
	case STATUS_MUTUAL_AUTHENTICATION_FAILED:
		return "0xC00002C3"
	case STATUS_CORRUPT_SYSTEM_FILE:
		return "0xC00002C4"
	case STATUS_DATATYPE_MISALIGNMENT_ERROR:
		return "0xC00002C5"
	case STATUS_WMI_READ_ONLY:
		return "0xC00002C6"
	case STATUS_WMI_SET_FAILURE:
		return "0xC00002C7"
	case STATUS_COMMITMENT_MINIMUM:
		return "0xC00002C8"
	case STATUS_REG_NAT_CONSUMPTION:
		return "0xC00002C9"
	case STATUS_TRANSPORT_FULL:
		return "0xC00002CA"
	case STATUS_DS_SAM_INIT_FAILURE:
		return "0xC00002CB"
	case STATUS_ONLY_IF_CONNECTED:
		return "0xC00002CC"
	case STATUS_DS_SENSITIVE_GROUP_VIOLATION:
		return "0xC00002CD"
	case STATUS_PNP_RESTART_ENUMERATION:
		return "0xC00002CE"
	case STATUS_JOURNAL_ENTRY_DELETED:
		return "0xC00002CF"
	case STATUS_DS_CANT_MOD_PRIMARYGROUPID:
		return "0xC00002D0"
	case STATUS_SYSTEM_IMAGE_BAD_SIGNATURE:
		return "0xC00002D1"
	case STATUS_PNP_REBOOT_REQUIRED:
		return "0xC00002D2"
	case STATUS_POWER_STATE_INVALID:
		return "0xC00002D3"
	case STATUS_DS_INVALID_GROUP_TYPE:
		return "0xC00002D4"
	case STATUS_DS_NO_NEST_GLOBALGROUP_IN_MIXEDDOMAIN:
		return "0xC00002D5"
	case STATUS_DS_NO_NEST_LOCALGROUP_IN_MIXEDDOMAIN:
		return "0xC00002D6"
	case STATUS_DS_GLOBAL_CANT_HAVE_LOCAL_MEMBER:
		return "0xC00002D7"
	case STATUS_DS_GLOBAL_CANT_HAVE_UNIVERSAL_MEMBER:
		return "0xC00002D8"
	case STATUS_DS_UNIVERSAL_CANT_HAVE_LOCAL_MEMBER:
		return "0xC00002D9"
	case STATUS_DS_GLOBAL_CANT_HAVE_CROSSDOMAIN_MEMBER:
		return "0xC00002DA"
	case STATUS_DS_LOCAL_CANT_HAVE_CROSSDOMAIN_LOCAL_MEMBER:
		return "0xC00002DB"
	case STATUS_DS_HAVE_PRIMARY_MEMBERS:
		return "0xC00002DC"
	case STATUS_WMI_NOT_SUPPORTED:
		return "0xC00002DD"
	case STATUS_INSUFFICIENT_POWER:
		return "0xC00002DE"
	case STATUS_SAM_NEED_BOOTKEY_PASSWORD:
		return "0xC00002DF"
	case STATUS_SAM_NEED_BOOTKEY_FLOPPY:
		return "0xC00002E0"
	case STATUS_DS_CANT_START:
		return "0xC00002E1"
	case STATUS_DS_INIT_FAILURE:
		return "0xC00002E2"
	case STATUS_SAM_INIT_FAILURE:
		return "0xC00002E3"
	case STATUS_DS_GC_REQUIRED:
		return "0xC00002E4"
	case STATUS_DS_LOCAL_MEMBER_OF_LOCAL_ONLY:
		return "0xC00002E5"
	case STATUS_DS_NO_FPO_IN_UNIVERSAL_GROUPS:
		return "0xC00002E6"
	case STATUS_DS_MACHINE_ACCOUNT_QUOTA_EXCEEDED:
		return "0xC00002E7"
	case STATUS_MULTIPLE_FAULT_VIOLATION:
		return "0xC00002E8"
	case STATUS_CURRENT_DOMAIN_NOT_ALLOWED:
		return "0xC00002E9"
	case STATUS_CANNOT_MAKE:
		return "0xC00002EA"
	case STATUS_SYSTEM_SHUTDOWN:
		return "0xC00002EB"
	case STATUS_DS_INIT_FAILURE_CONSOLE:
		return "0xC00002EC"
	case STATUS_DS_SAM_INIT_FAILURE_CONSOLE:
		return "0xC00002ED"
	case STATUS_UNFINISHED_CONTEXT_DELETED:
		return "0xC00002EE"
	case STATUS_NO_TGT_REPLY:
		return "0xC00002EF"
	case STATUS_OBJECTID_NOT_FOUND:
		return "0xC00002F0"
	case STATUS_NO_IP_ADDRESSES:
		return "0xC00002F1"
	case STATUS_WRONG_CREDENTIAL_HANDLE:
		return "0xC00002F2"
	case STATUS_CRYPTO_SYSTEM_INVALID:
		return "0xC00002F3"
	case STATUS_MAX_REFERRALS_EXCEEDED:
		return "0xC00002F4"
	case STATUS_MUST_BE_KDC:
		return "0xC00002F5"
	case STATUS_STRONG_CRYPTO_NOT_SUPPORTED:
		return "0xC00002F6"
	case STATUS_TOO_MANY_PRINCIPALS:
		return "0xC00002F7"
	case STATUS_NO_PA_DATA:
		return "0xC00002F8"
	case STATUS_PKINIT_NAME_MISMATCH:
		return "0xC00002F9"
	case STATUS_SMARTCARD_LOGON_REQUIRED:
		return "0xC00002FA"
	case STATUS_KDC_INVALID_REQUEST:
		return "0xC00002FB"
	case STATUS_KDC_UNABLE_TO_REFER:
		return "0xC00002FC"
	case STATUS_KDC_UNKNOWN_ETYPE:
		return "0xC00002FD"
	case STATUS_SHUTDOWN_IN_PROGRESS:
		return "0xC00002FE"
	case STATUS_SERVER_SHUTDOWN_IN_PROGRESS:
		return "0xC00002FF"
	case STATUS_NOT_SUPPORTED_ON_SBS:
		return "0xC0000300"
	case STATUS_WMI_GUID_DISCONNECTED:
		return "0xC0000301"
	case STATUS_WMI_ALREADY_DISABLED:
		return "0xC0000302"
	case STATUS_WMI_ALREADY_ENABLED:
		return "0xC0000303"
	case STATUS_MFT_TOO_FRAGMENTED:
		return "0xC0000304"
	case STATUS_COPY_PROTECTION_FAILURE:
		return "0xC0000305"
	case STATUS_CSS_AUTHENTICATION_FAILURE:
		return "0xC0000306"
	case STATUS_CSS_KEY_NOT_PRESENT:
		return "0xC0000307"
	case STATUS_CSS_KEY_NOT_ESTABLISHED:
		return "0xC0000308"
	case STATUS_CSS_SCRAMBLED_SECTOR:
		return "0xC0000309"
	case STATUS_CSS_REGION_MISMATCH:
		return "0xC000030A"
	case STATUS_CSS_RESETS_EXHAUSTED:
		return "0xC000030B"
	case STATUS_PKINIT_FAILURE:
		return "0xC0000320"
	case STATUS_SMARTCARD_SUBSYSTEM_FAILURE:
		return "0xC0000321"
	case STATUS_NO_KERB_KEY:
		return "0xC0000322"
	case STATUS_HOST_DOWN:
		return "0xC0000350"
	case STATUS_UNSUPPORTED_PREAUTH:
		return "0xC0000351"
	case STATUS_EFS_ALG_BLOB_TOO_BIG:
		return "0xC0000352"
	case STATUS_PORT_NOT_SET:
		return "0xC0000353"
	case STATUS_DEBUGGER_INACTIVE:
		return "0xC0000354"
	case STATUS_DS_VERSION_CHECK_FAILURE:
		return "0xC0000355"
	case STATUS_AUDITING_DISABLED:
		return "0xC0000356"
	case STATUS_PRENT4_MACHINE_ACCOUNT:
		return "0xC0000357"
	case STATUS_DS_AG_CANT_HAVE_UNIVERSAL_MEMBER:
		return "0xC0000358"
	case STATUS_INVALID_IMAGE_WIN_32:
		return "0xC0000359"
	case STATUS_INVALID_IMAGE_WIN_64:
		return "0xC000035A"
	case STATUS_BAD_BINDINGS:
		return "0xC000035B"
	case STATUS_NETWORK_SESSION_EXPIRED:
		return "0xC000035C"
	case STATUS_APPHELP_BLOCK:
		return "0xC000035D"
	case STATUS_ALL_SIDS_FILTERED:
		return "0xC000035E"
	case STATUS_NOT_SAFE_MODE_DRIVER:
		return "0xC000035F"
	case STATUS_ACCESS_DISABLED_BY_POLICY_DEFAULT:
		return "0xC0000361"
	case STATUS_ACCESS_DISABLED_BY_POLICY_PATH:
		return "0xC0000362"
	case STATUS_ACCESS_DISABLED_BY_POLICY_PUBLISHER:
		return "0xC0000363"
	case STATUS_ACCESS_DISABLED_BY_POLICY_OTHER:
		return "0xC0000364"
	case STATUS_FAILED_DRIVER_ENTRY:
		return "0xC0000365"
	case STATUS_DEVICE_ENUMERATION_ERROR:
		return "0xC0000366"
	case STATUS_MOUNT_POINT_NOT_RESOLVED:
		return "0xC0000368"
	case STATUS_INVALID_DEVICE_OBJECT_PARAMETER:
		return "0xC0000369"
	case STATUS_MCA_OCCURED:
		return "0xC000036A"
	case STATUS_DRIVER_BLOCKED_CRITICAL:
		return "0xC000036B"
	case STATUS_DRIVER_BLOCKED:
		return "0xC000036C"
	case STATUS_DRIVER_DATABASE_ERROR:
		return "0xC000036D"
	case STATUS_SYSTEM_HIVE_TOO_LARGE:
		return "0xC000036E"
	case STATUS_INVALID_IMPORT_OF_NON_DLL:
		return "0xC000036F"
	case STATUS_DS_SHUTTING_DOWN:
		return "0x40000370"
	case STATUS_NO_SECRETS:
		return "0xC0000371"
	case STATUS_ACCESS_DISABLED_NO_SAFER_UI_BY_POLICY:
		return "0xC0000372"
	case STATUS_FAILED_STACK_SWITCH:
		return "0xC0000373"
	case STATUS_HEAP_CORRUPTION:
		return "0xC0000374"
	case STATUS_SMARTCARD_WRONG_PIN:
		return "0xC0000380"
	case STATUS_SMARTCARD_CARD_BLOCKED:
		return "0xC0000381"
	case STATUS_SMARTCARD_CARD_NOT_AUTHENTICATED:
		return "0xC0000382"
	case STATUS_SMARTCARD_NO_CARD:
		return "0xC0000383"
	case STATUS_SMARTCARD_NO_KEY_CONTAINER:
		return "0xC0000384"
	case STATUS_SMARTCARD_NO_CERTIFICATE:
		return "0xC0000385"
	case STATUS_SMARTCARD_NO_KEYSET:
		return "0xC0000386"
	case STATUS_SMARTCARD_IO_ERROR:
		return "0xC0000387"
	case STATUS_DOWNGRADE_DETECTED:
		return "0xC0000388"
	case STATUS_SMARTCARD_CERT_REVOKED:
		return "0xC0000389"
	case STATUS_ISSUING_CA_UNTRUSTED:
		return "0xC000038A"
	case STATUS_REVOCATION_OFFLINE_C:
		return "0xC000038B"
	case STATUS_PKINIT_CLIENT_FAILURE:
		return "0xC000038C"
	case STATUS_SMARTCARD_CERT_EXPIRED:
		return "0xC000038D"
	case STATUS_DRIVER_FAILED_PRIOR_UNLOAD:
		return "0xC000038E"
	case STATUS_SMARTCARD_SILENT_CONTEXT:
		return "0xC000038F"
	case STATUS_PER_USER_TRUST_QUOTA_EXCEEDED:
		return "0xC0000401"
	case STATUS_ALL_USER_TRUST_QUOTA_EXCEEDED:
		return "0xC0000402"
	case STATUS_USER_DELETE_TRUST_QUOTA_EXCEEDED:
		return "0xC0000403"
	case STATUS_DS_NAME_NOT_UNIQUE:
		return "0xC0000404"
	case STATUS_DS_DUPLICATE_ID_FOUND:
		return "0xC0000405"
	case STATUS_DS_GROUP_CONVERSION_ERROR:
		return "0xC0000406"
	case STATUS_VOLSNAP_PREPARE_HIBERNATE:
		return "0xC0000407"
	case STATUS_USER2USER_REQUIRED:
		return "0xC0000408"
	case STATUS_STACK_BUFFER_OVERRUN:
		return "0xC0000409"
	case STATUS_NO_S4U_PROT_SUPPORT:
		return "0xC000040A"
	case STATUS_CROSSREALM_DELEGATION_FAILURE:
		return "0xC000040B"
	case STATUS_REVOCATION_OFFLINE_KDC:
		return "0xC000040C"
	case STATUS_ISSUING_CA_UNTRUSTED_KDC:
		return "0xC000040D"
	case STATUS_KDC_CERT_EXPIRED:
		return "0xC000040E"
	case STATUS_KDC_CERT_REVOKED:
		return "0xC000040F"
	case STATUS_PARAMETER_QUOTA_EXCEEDED:
		return "0xC0000410"
	case STATUS_HIBERNATION_FAILURE:
		return "0xC0000411"
	case STATUS_DELAY_LOAD_FAILED:
		return "0xC0000412"
	case STATUS_AUTHENTICATION_FIREWALL_FAILED:
		return "0xC0000413"
	case STATUS_VDM_DISALLOWED:
		return "0xC0000414"
	case STATUS_HUNG_DISPLAY_DRIVER_THREAD:
		return "0xC0000415"
	case STATUS_INSUFFICIENT_RESOURCE_FOR_SPECIFIED_SHARED_SECTION_SIZE:
		return "0xC0000416"
	case STATUS_INVALID_CRUNTIME_PARAMETER:
		return "0xC0000417"
	case STATUS_NTLM_BLOCKED:
		return "0xC0000418"
	case STATUS_DS_SRC_SID_EXISTS_IN_FOREST:
		return "0xC0000419"
	case STATUS_DS_DOMAIN_NAME_EXISTS_IN_FOREST:
		return "0xC000041A"
	case STATUS_DS_FLAT_NAME_EXISTS_IN_FOREST:
		return "0xC000041B"
	case STATUS_INVALID_USER_PRINCIPAL_NAME:
		return "0xC000041C"
	case STATUS_FATAL_USER_CALLBACK_EXCEPTION:
		return "0xC000041D"
	case STATUS_ASSERTION_FAILURE:
		return "0xC0000420"
	case STATUS_VERIFIER_STOP:
		return "0xC0000421"
	case STATUS_CALLBACK_POP_STACK:
		return "0xC0000423"
	case STATUS_INCOMPATIBLE_DRIVER_BLOCKED:
		return "0xC0000424"
	case STATUS_HIVE_UNLOADED:
		return "0xC0000425"
	case STATUS_COMPRESSION_DISABLED:
		return "0xC0000426"
	case STATUS_FILE_SYSTEM_LIMITATION:
		return "0xC0000427"
	case STATUS_INVALID_IMAGE_HASH:
		return "0xC0000428"
	case STATUS_NOT_CAPABLE:
		return "0xC0000429"
	case STATUS_REQUEST_OUT_OF_SEQUENCE:
		return "0xC000042A"
	case STATUS_IMPLEMENTATION_LIMIT:
		return "0xC000042B"
	case STATUS_ELEVATION_REQUIRED:
		return "0xC000042C"
	case STATUS_NO_SECURITY_CONTEXT:
		return "0xC000042D"
	case STATUS_PKU2U_CERT_FAILURE:
		return "0xC000042F"
	case STATUS_BEYOND_VDL:
		return "0xC0000432"
	case STATUS_ENCOUNTERED_WRITE_IN_PROGRESS:
		return "0xC0000433"
	case STATUS_PTE_CHANGED:
		return "0xC0000434"
	case STATUS_PURGE_FAILED:
		return "0xC0000435"
	case STATUS_CRED_REQUIRES_CONFIRMATION:
		return "0xC0000440"
	case STATUS_CS_ENCRYPTION_INVALID_SERVER_RESPONSE:
		return "0xC0000441"
	case STATUS_CS_ENCRYPTION_UNSUPPORTED_SERVER:
		return "0xC0000442"
	case STATUS_CS_ENCRYPTION_EXISTING_ENCRYPTED_FILE:
		return "0xC0000443"
	case STATUS_CS_ENCRYPTION_NEW_ENCRYPTED_FILE:
		return "0xC0000444"
	case STATUS_CS_ENCRYPTION_FILE_NOT_CSE:
		return "0xC0000445"
	case STATUS_INVALID_LABEL:
		return "0xC0000446"
	case STATUS_DRIVER_PROCESS_TERMINATED:
		return "0xC0000450"
	case STATUS_AMBIGUOUS_SYSTEM_DEVICE:
		return "0xC0000451"
	case STATUS_SYSTEM_DEVICE_NOT_FOUND:
		return "0xC0000452"
	case STATUS_RESTART_BOOT_APPLICATION:
		return "0xC0000453"
	case STATUS_INSUFFICIENT_NVRAM_RESOURCES:
		return "0xC0000454"
	case STATUS_INVALID_TASK_NAME:
		return "0xC0000500"
	case STATUS_INVALID_TASK_INDEX:
		return "0xC0000501"
	case STATUS_THREAD_ALREADY_IN_TASK:
		return "0xC0000502"
	case STATUS_CALLBACK_BYPASS:
		return "0xC0000503"
	case STATUS_FAIL_FAST_EXCEPTION:
		return "0xC0000602"
	case STATUS_IMAGE_CERT_REVOKED:
		return "0xC0000603"
	case STATUS_PORT_CLOSED:
		return "0xC0000700"
	case STATUS_MESSAGE_LOST:
		return "0xC0000701"
	case STATUS_INVALID_MESSAGE:
		return "0xC0000702"
	case STATUS_REQUEST_CANCELED:
		return "0xC0000703"
	case STATUS_RECURSIVE_DISPATCH:
		return "0xC0000704"
	case STATUS_LPC_RECEIVE_BUFFER_EXPECTED:
		return "0xC0000705"
	case STATUS_LPC_INVALID_CONNECTION_USAGE:
		return "0xC0000706"
	case STATUS_LPC_REQUESTS_NOT_ALLOWED:
		return "0xC0000707"
	case STATUS_RESOURCE_IN_USE:
		return "0xC0000708"
	case STATUS_HARDWARE_MEMORY_ERROR:
		return "0xC0000709"
	case STATUS_THREADPOOL_HANDLE_EXCEPTION:
		return "0xC000070A"
	case STATUS_THREADPOOL_SET_EVENT_ON_COMPLETION_FAILED:
		return "0xC000070B"
	case STATUS_THREADPOOL_RELEASE_SEMAPHORE_ON_COMPLETION_FAILED:
		return "0xC000070C"
	case STATUS_THREADPOOL_RELEASE_MUTEX_ON_COMPLETION_FAILED:
		return "0xC000070D"
	case STATUS_THREADPOOL_FREE_LIBRARY_ON_COMPLETION_FAILED:
		return "0xC000070E"
	case STATUS_THREADPOOL_RELEASED_DURING_OPERATION:
		return "0xC000070F"
	case STATUS_CALLBACK_RETURNED_WHILE_IMPERSONATING:
		return "0xC0000710"
	case STATUS_APC_RETURNED_WHILE_IMPERSONATING:
		return "0xC0000711"
	case STATUS_PROCESS_IS_PROTECTED:
		return "0xC0000712"
	case STATUS_MCA_EXCEPTION:
		return "0xC0000713"
	case STATUS_CERTIFICATE_MAPPING_NOT_UNIQUE:
		return "0xC0000714"
	case STATUS_SYMLINK_CLASS_DISABLED:
		return "0xC0000715"
	case STATUS_INVALID_IDN_NORMALIZATION:
		return "0xC0000716"
	case STATUS_NO_UNICODE_TRANSLATION:
		return "0xC0000717"
	case STATUS_ALREADY_REGISTERED:
		return "0xC0000718"
	case STATUS_CONTEXT_MISMATCH:
		return "0xC0000719"
	case STATUS_PORT_ALREADY_HAS_COMPLETION_LIST:
		return "0xC000071A"
	case STATUS_CALLBACK_RETURNED_THREAD_PRIORITY:
		return "0xC000071B"
	case STATUS_INVALID_THREAD:
		return "0xC000071C"
	case STATUS_CALLBACK_RETURNED_TRANSACTION:
		return "0xC000071D"
	case STATUS_CALLBACK_RETURNED_LDR_LOCK:
		return "0xC000071E"
	case STATUS_CALLBACK_RETURNED_LANG:
		return "0xC000071F"
	case STATUS_CALLBACK_RETURNED_PRI_BACK:
		return "0xC0000720"
	case STATUS_CALLBACK_RETURNED_THREAD_AFFINITY:
		return "0xC0000721"
	case STATUS_DISK_REPAIR_DISABLED:
		return "0xC0000800"
	case STATUS_DS_DOMAIN_RENAME_IN_PROGRESS:
		return "0xC0000801"
	case STATUS_DISK_QUOTA_EXCEEDED:
		return "0xC0000802"
	case STATUS_DATA_LOST_REPAIR:
		return "0x80000803"
	case STATUS_CONTENT_BLOCKED:
		return "0xC0000804"
	case STATUS_BAD_CLUSTERS:
		return "0xC0000805"
	case STATUS_VOLUME_DIRTY:
		return "0xC0000806"
	case STATUS_FILE_CHECKED_OUT:
		return "0xC0000901"
	case STATUS_CHECKOUT_REQUIRED:
		return "0xC0000902"
	case STATUS_BAD_FILE_TYPE:
		return "0xC0000903"
	case STATUS_FILE_TOO_LARGE:
		return "0xC0000904"
	case STATUS_FORMS_AUTH_REQUIRED:
		return "0xC0000905"
	case STATUS_VIRUS_INFECTED:
		return "0xC0000906"
	case STATUS_VIRUS_DELETED:
		return "0xC0000907"
	case STATUS_BAD_MCFG_TABLE:
		return "0xC0000908"
	case STATUS_CANNOT_BREAK_OPLOCK:
		return "0xC0000909"
	case STATUS_WOW_ASSERTION:
		return "0xC0009898"
	case STATUS_INVALID_SIGNATURE:
		return "0xC000A000"
	case STATUS_HMAC_NOT_SUPPORTED:
		return "0xC000A001"
	case STATUS_AUTH_TAG_MISMATCH:
		return "0xC000A002"
	case STATUS_IPSEC_QUEUE_OVERFLOW:
		return "0xC000A010"
	case STATUS_ND_QUEUE_OVERFLOW:
		return "0xC000A011"
	case STATUS_HOPLIMIT_EXCEEDED:
		return "0xC000A012"
	case STATUS_PROTOCOL_NOT_SUPPORTED:
		return "0xC000A013"
	case STATUS_FASTPATH_REJECTED:
		return "0xC000A014"
	case STATUS_LOST_WRITEBEHIND_DATA_NETWORK_DISCONNECTED:
		return "0xC000A080"
	case STATUS_LOST_WRITEBEHIND_DATA_NETWORK_SERVER_ERROR:
		return "0xC000A081"
	case STATUS_LOST_WRITEBEHIND_DATA_LOCAL_DISK_ERROR:
		return "0xC000A082"
	case STATUS_XML_PARSE_ERROR:
		return "0xC000A083"
	case STATUS_XMLDSIG_ERROR:
		return "0xC000A084"
	case STATUS_WRONG_COMPARTMENT:
		return "0xC000A085"
	case STATUS_AUTHIP_FAILURE:
		return "0xC000A086"
	case STATUS_DS_OID_MAPPED_GROUP_CANT_HAVE_MEMBERS:
		return "0xC000A087"
	case STATUS_DS_OID_NOT_FOUND:
		return "0xC000A088"
	case STATUS_HASH_NOT_SUPPORTED:
		return "0xC000A100"
	case STATUS_HASH_NOT_PRESENT:
		return "0xC000A101"
	case DBG_NO_STATE_CHANGE:
		return "0xC0010001"
	case DBG_APP_NOT_IDLE:
		return "0xC0010002"
	case RPC_NT_INVALID_STRING_BINDING:
		return "0xC0020001"
	case RPC_NT_WRONG_KIND_OF_BINDING:
		return "0xC0020002"
	case RPC_NT_INVALID_BINDING:
		return "0xC0020003"
	case RPC_NT_PROTSEQ_NOT_SUPPORTED:
		return "0xC0020004"
	case RPC_NT_INVALID_RPC_PROTSEQ:
		return "0xC0020005"
	case RPC_NT_INVALID_STRING_UUID:
		return "0xC0020006"
	case RPC_NT_INVALID_ENDPOINT_FORMAT:
		return "0xC0020007"
	case RPC_NT_INVALID_NET_ADDR:
		return "0xC0020008"
	case RPC_NT_NO_ENDPOINT_FOUND:
		return "0xC0020009"
	case RPC_NT_INVALID_TIMEOUT:
		return "0xC002000A"
	case RPC_NT_OBJECT_NOT_FOUND:
		return "0xC002000B"
	case RPC_NT_ALREADY_REGISTERED:
		return "0xC002000C"
	case RPC_NT_TYPE_ALREADY_REGISTERED:
		return "0xC002000D"
	case RPC_NT_ALREADY_LISTENING:
		return "0xC002000E"
	case RPC_NT_NO_PROTSEQS_REGISTERED:
		return "0xC002000F"
	case RPC_NT_NOT_LISTENING:
		return "0xC0020010"
	case RPC_NT_UNKNOWN_MGR_TYPE:
		return "0xC0020011"
	case RPC_NT_UNKNOWN_IF:
		return "0xC0020012"
	case RPC_NT_NO_BINDINGS:
		return "0xC0020013"
	case RPC_NT_NO_PROTSEQS:
		return "0xC0020014"
	case RPC_NT_CANT_CREATE_ENDPOINT:
		return "0xC0020015"
	case RPC_NT_OUT_OF_RESOURCES:
		return "0xC0020016"
	case RPC_NT_SERVER_UNAVAILABLE:
		return "0xC0020017"
	case RPC_NT_SERVER_TOO_BUSY:
		return "0xC0020018"
	case RPC_NT_INVALID_NETWORK_OPTIONS:
		return "0xC0020019"
	case RPC_NT_NO_CALL_ACTIVE:
		return "0xC002001A"
	case RPC_NT_CALL_FAILED:
		return "0xC002001B"
	case RPC_NT_CALL_FAILED_DNE:
		return "0xC002001C"
	case RPC_NT_PROTOCOL_ERROR:
		return "0xC002001D"
	case RPC_NT_UNSUPPORTED_TRANS_SYN:
		return "0xC002001F"
	case RPC_NT_UNSUPPORTED_TYPE:
		return "0xC0020021"
	case RPC_NT_INVALID_TAG:
		return "0xC0020022"
	case RPC_NT_INVALID_BOUND:
		return "0xC0020023"
	case RPC_NT_NO_ENTRY_NAME:
		return "0xC0020024"
	case RPC_NT_INVALID_NAME_SYNTAX:
		return "0xC0020025"
	case RPC_NT_UNSUPPORTED_NAME_SYNTAX:
		return "0xC0020026"
	case RPC_NT_UUID_NO_ADDRESS:
		return "0xC0020028"
	case RPC_NT_DUPLICATE_ENDPOINT:
		return "0xC0020029"
	case RPC_NT_UNKNOWN_AUTHN_TYPE:
		return "0xC002002A"
	case RPC_NT_MAX_CALLS_TOO_SMALL:
		return "0xC002002B"
	case RPC_NT_STRING_TOO_LONG:
		return "0xC002002C"
	case RPC_NT_PROTSEQ_NOT_FOUND:
		return "0xC002002D"
	case RPC_NT_PROCNUM_OUT_OF_RANGE:
		return "0xC002002E"
	case RPC_NT_BINDING_HAS_NO_AUTH:
		return "0xC002002F"
	case RPC_NT_UNKNOWN_AUTHN_SERVICE:
		return "0xC0020030"
	case RPC_NT_UNKNOWN_AUTHN_LEVEL:
		return "0xC0020031"
	case RPC_NT_INVALID_AUTH_IDENTITY:
		return "0xC0020032"
	case RPC_NT_UNKNOWN_AUTHZ_SERVICE:
		return "0xC0020033"
	case EPT_NT_INVALID_ENTRY:
		return "0xC0020034"
	case EPT_NT_CANT_PERFORM_OP:
		return "0xC0020035"
	case EPT_NT_NOT_REGISTERED:
		return "0xC0020036"
	case RPC_NT_NOTHING_TO_EXPORT:
		return "0xC0020037"
	case RPC_NT_INCOMPLETE_NAME:
		return "0xC0020038"
	case RPC_NT_INVALID_VERS_OPTION:
		return "0xC0020039"
	case RPC_NT_NO_MORE_MEMBERS:
		return "0xC002003A"
	case RPC_NT_NOT_ALL_OBJS_UNEXPORTED:
		return "0xC002003B"
	case RPC_NT_INTERFACE_NOT_FOUND:
		return "0xC002003C"
	case RPC_NT_ENTRY_ALREADY_EXISTS:
		return "0xC002003D"
	case RPC_NT_ENTRY_NOT_FOUND:
		return "0xC002003E"
	case RPC_NT_NAME_SERVICE_UNAVAILABLE:
		return "0xC002003F"
	case RPC_NT_INVALID_NAF_ID:
		return "0xC0020040"
	case RPC_NT_CANNOT_SUPPORT:
		return "0xC0020041"
	case RPC_NT_NO_CONTEXT_AVAILABLE:
		return "0xC0020042"
	case RPC_NT_INTERNAL_ERROR:
		return "0xC0020043"
	case RPC_NT_ZERO_DIVIDE:
		return "0xC0020044"
	case RPC_NT_ADDRESS_ERROR:
		return "0xC0020045"
	case RPC_NT_FP_DIV_ZERO:
		return "0xC0020046"
	case RPC_NT_FP_UNDERFLOW:
		return "0xC0020047"
	case RPC_NT_FP_OVERFLOW:
		return "0xC0020048"
	case RPC_NT_NO_MORE_ENTRIES:
		return "0xC0030001"
	case RPC_NT_SS_CHAR_TRANS_OPEN_FAIL:
		return "0xC0030002"
	case RPC_NT_SS_CHAR_TRANS_SHORT_FILE:
		return "0xC0030003"
	case RPC_NT_SS_IN_NULL_CONTEXT:
		return "0xC0030004"
	case RPC_NT_SS_CONTEXT_MISMATCH:
		return "0xC0030005"
	case RPC_NT_SS_CONTEXT_DAMAGED:
		return "0xC0030006"
	case RPC_NT_SS_HANDLES_MISMATCH:
		return "0xC0030007"
	case RPC_NT_SS_CANNOT_GET_CALL_HANDLE:
		return "0xC0030008"
	case RPC_NT_NULL_REF_POINTER:
		return "0xC0030009"
	case RPC_NT_ENUM_VALUE_OUT_OF_RANGE:
		return "0xC003000A"
	case RPC_NT_BYTE_COUNT_TOO_SMALL:
		return "0xC003000B"
	case RPC_NT_BAD_STUB_DATA:
		return "0xC003000C"
	case RPC_NT_CALL_IN_PROGRESS:
		return "0xC0020049"
	case RPC_NT_NO_MORE_BINDINGS:
		return "0xC002004A"
	case RPC_NT_GROUP_MEMBER_NOT_FOUND:
		return "0xC002004B"
	case EPT_NT_CANT_CREATE:
		return "0xC002004C"
	case RPC_NT_INVALID_OBJECT:
		return "0xC002004D"
	case RPC_NT_NO_INTERFACES:
		return "0xC002004F"
	case RPC_NT_CALL_CANCELLED:
		return "0xC0020050"
	case RPC_NT_BINDING_INCOMPLETE:
		return "0xC0020051"
	case RPC_NT_COMM_FAILURE:
		return "0xC0020052"
	case RPC_NT_UNSUPPORTED_AUTHN_LEVEL:
		return "0xC0020053"
	case RPC_NT_NO_PRINC_NAME:
		return "0xC0020054"
	case RPC_NT_NOT_RPC_ERROR:
		return "0xC0020055"
	case RPC_NT_UUID_LOCAL_ONLY:
		return "0x40020056"
	case RPC_NT_SEC_PKG_ERROR:
		return "0xC0020057"
	case RPC_NT_NOT_CANCELLED:
		return "0xC0020058"
	case RPC_NT_INVALID_ES_ACTION:
		return "0xC0030059"
	case RPC_NT_WRONG_ES_VERSION:
		return "0xC003005A"
	case RPC_NT_WRONG_STUB_VERSION:
		return "0xC003005B"
	case RPC_NT_INVALID_PIPE_OBJECT:
		return "0xC003005C"
	case RPC_NT_INVALID_PIPE_OPERATION:
		return "0xC003005D"
	case RPC_NT_WRONG_PIPE_VERSION:
		return "0xC003005E"
	case RPC_NT_PIPE_CLOSED:
		return "0xC003005F"
	case RPC_NT_PIPE_DISCIPLINE_ERROR:
		return "0xC0030060"
	case RPC_NT_PIPE_EMPTY:
		return "0xC0030061"
	case RPC_NT_INVALID_ASYNC_HANDLE:
		return "0xC0020062"
	case RPC_NT_INVALID_ASYNC_CALL:
		return "0xC0020063"
	case RPC_NT_PROXY_ACCESS_DENIED:
		return "0xC0020064"
	case RPC_NT_COOKIE_AUTH_FAILED:
		return "0xC0020065"
	case RPC_NT_SEND_INCOMPLETE:
		return "0x400200AF"
	case STATUS_ACPI_INVALID_OPCODE:
		return "0xC0140001"
	case STATUS_ACPI_STACK_OVERFLOW:
		return "0xC0140002"
	case STATUS_ACPI_ASSERT_FAILED:
		return "0xC0140003"
	case STATUS_ACPI_INVALID_INDEX:
		return "0xC0140004"
	case STATUS_ACPI_INVALID_ARGUMENT:
		return "0xC0140005"
	case STATUS_ACPI_FATAL:
		return "0xC0140006"
	case STATUS_ACPI_INVALID_SUPERNAME:
		return "0xC0140007"
	case STATUS_ACPI_INVALID_ARGTYPE:
		return "0xC0140008"
	case STATUS_ACPI_INVALID_OBJTYPE:
		return "0xC0140009"
	case STATUS_ACPI_INVALID_TARGETTYPE:
		return "0xC014000A"
	case STATUS_ACPI_INCORRECT_ARGUMENT_COUNT:
		return "0xC014000B"
	case STATUS_ACPI_ADDRESS_NOT_MAPPED:
		return "0xC014000C"
	case STATUS_ACPI_INVALID_EVENTTYPE:
		return "0xC014000D"
	case STATUS_ACPI_HANDLER_COLLISION:
		return "0xC014000E"
	case STATUS_ACPI_INVALID_DATA:
		return "0xC014000F"
	case STATUS_ACPI_INVALID_REGION:
		return "0xC0140010"
	case STATUS_ACPI_INVALID_ACCESS_SIZE:
		return "0xC0140011"
	case STATUS_ACPI_ACQUIRE_GLOBAL_LOCK:
		return "0xC0140012"
	case STATUS_ACPI_ALREADY_INITIALIZED:
		return "0xC0140013"
	case STATUS_ACPI_NOT_INITIALIZED:
		return "0xC0140014"
	case STATUS_ACPI_INVALID_MUTEX_LEVEL:
		return "0xC0140015"
	case STATUS_ACPI_MUTEX_NOT_OWNED:
		return "0xC0140016"
	case STATUS_ACPI_MUTEX_NOT_OWNER:
		return "0xC0140017"
	case STATUS_ACPI_RS_ACCESS:
		return "0xC0140018"
	case STATUS_ACPI_INVALID_TABLE:
		return "0xC0140019"
	case STATUS_ACPI_REG_HANDLER_FAILED:
		return "0xC0140020"
	case STATUS_ACPI_POWER_REQUEST_FAILED:
		return "0xC0140021"
	case STATUS_CTX_WINSTATION_NAME_INVALID:
		return "0xC00A0001"
	case STATUS_CTX_INVALID_PD:
		return "0xC00A0002"
	case STATUS_CTX_PD_NOT_FOUND:
		return "0xC00A0003"
	case STATUS_CTX_CDM_CONNECT:
		return "0x400A0004"
	case STATUS_CTX_CDM_DISCONNECT:
		return "0x400A0005"
	case STATUS_CTX_CLOSE_PENDING:
		return "0xC00A0006"
	case STATUS_CTX_NO_OUTBUF:
		return "0xC00A0007"
	case STATUS_CTX_MODEM_INF_NOT_FOUND:
		return "0xC00A0008"
	case STATUS_CTX_INVALID_MODEMNAME:
		return "0xC00A0009"
	case STATUS_CTX_RESPONSE_ERROR:
		return "0xC00A000A"
	case STATUS_CTX_MODEM_RESPONSE_TIMEOUT:
		return "0xC00A000B"
	case STATUS_CTX_MODEM_RESPONSE_NO_CARRIER:
		return "0xC00A000C"
	case STATUS_CTX_MODEM_RESPONSE_NO_DIALTONE:
		return "0xC00A000D"
	case STATUS_CTX_MODEM_RESPONSE_BUSY:
		return "0xC00A000E"
	case STATUS_CTX_MODEM_RESPONSE_VOICE:
		return "0xC00A000F"
	case STATUS_CTX_TD_ERROR:
		return "0xC00A0010"
	case STATUS_CTX_LICENSE_CLIENT_INVALID:
		return "0xC00A0012"
	case STATUS_CTX_LICENSE_NOT_AVAILABLE:
		return "0xC00A0013"
	case STATUS_CTX_LICENSE_EXPIRED:
		return "0xC00A0014"
	case STATUS_CTX_WINSTATION_NOT_FOUND:
		return "0xC00A0015"
	case STATUS_CTX_WINSTATION_NAME_COLLISION:
		return "0xC00A0016"
	case STATUS_CTX_WINSTATION_BUSY:
		return "0xC00A0017"
	case STATUS_CTX_BAD_VIDEO_MODE:
		return "0xC00A0018"
	case STATUS_CTX_GRAPHICS_INVALID:
		return "0xC00A0022"
	case STATUS_CTX_NOT_CONSOLE:
		return "0xC00A0024"
	case STATUS_CTX_CLIENT_QUERY_TIMEOUT:
		return "0xC00A0026"
	case STATUS_CTX_CONSOLE_DISCONNECT:
		return "0xC00A0027"
	case STATUS_CTX_CONSOLE_CONNECT:
		return "0xC00A0028"
	case STATUS_CTX_SHADOW_DENIED:
		return "0xC00A002A"
	case STATUS_CTX_WINSTATION_ACCESS_DENIED:
		return "0xC00A002B"
	case STATUS_CTX_INVALID_WD:
		return "0xC00A002E"
	case STATUS_CTX_WD_NOT_FOUND:
		return "0xC00A002F"
	case STATUS_CTX_SHADOW_INVALID:
		return "0xC00A0030"
	case STATUS_CTX_SHADOW_DISABLED:
		return "0xC00A0031"
	case STATUS_RDP_PROTOCOL_ERROR:
		return "0xC00A0032"
	case STATUS_CTX_CLIENT_LICENSE_NOT_SET:
		return "0xC00A0033"
	case STATUS_CTX_CLIENT_LICENSE_IN_USE:
		return "0xC00A0034"
	case STATUS_CTX_SHADOW_ENDED_BY_MODE_CHANGE:
		return "0xC00A0035"
	case STATUS_CTX_SHADOW_NOT_RUNNING:
		return "0xC00A0036"
	case STATUS_CTX_LOGON_DISABLED:
		return "0xC00A0037"
	case STATUS_CTX_SECURITY_LAYER_ERROR:
		return "0xC00A0038"
	case STATUS_TS_INCOMPATIBLE_SESSIONS:
		return "0xC00A0039"
	case STATUS_TS_VIDEO_SUBSYSTEM_ERROR:
		return "0xC00A003A"
	case STATUS_PNP_BAD_MPS_TABLE:
		return "0xC0040035"
	case STATUS_PNP_TRANSLATION_FAILED:
		return "0xC0040036"
	case STATUS_PNP_IRQ_TRANSLATION_FAILED:
		return "0xC0040037"
	case STATUS_PNP_INVALID_ID:
		return "0xC0040038"
	case STATUS_IO_REISSUE_AS_CACHED:
		return "0xC0040039"
	case STATUS_MUI_FILE_NOT_FOUND:
		return "0xC00B0001"
	case STATUS_MUI_INVALID_FILE:
		return "0xC00B0002"
	case STATUS_MUI_INVALID_RC_CONFIG:
		return "0xC00B0003"
	case STATUS_MUI_INVALID_LOCALE_NAME:
		return "0xC00B0004"
	case STATUS_MUI_INVALID_ULTIMATEFALLBACK_NAME:
		return "0xC00B0005"
	case STATUS_MUI_FILE_NOT_LOADED:
		return "0xC00B0006"
	case STATUS_RESOURCE_ENUM_USER_STOP:
		return "0xC00B0007"
	case STATUS_FLT_NO_HANDLER_DEFINED:
		return "0xC01C0001"
	case STATUS_FLT_CONTEXT_ALREADY_DEFINED:
		return "0xC01C0002"
	case STATUS_FLT_INVALID_ASYNCHRONOUS_REQUEST:
		return "0xC01C0003"
	case STATUS_FLT_DISALLOW_FAST_IO:
		return "0xC01C0004"
	case STATUS_FLT_INVALID_NAME_REQUEST:
		return "0xC01C0005"
	case STATUS_FLT_NOT_SAFE_TO_POST_OPERATION:
		return "0xC01C0006"
	case STATUS_FLT_NOT_INITIALIZED:
		return "0xC01C0007"
	case STATUS_FLT_FILTER_NOT_READY:
		return "0xC01C0008"
	case STATUS_FLT_POST_OPERATION_CLEANUP:
		return "0xC01C0009"
	case STATUS_FLT_INTERNAL_ERROR:
		return "0xC01C000A"
	case STATUS_FLT_DELETING_OBJECT:
		return "0xC01C000B"
	case STATUS_FLT_MUST_BE_NONPAGED_POOL:
		return "0xC01C000C"
	case STATUS_FLT_DUPLICATE_ENTRY:
		return "0xC01C000D"
	case STATUS_FLT_CBDQ_DISABLED:
		return "0xC01C000E"
	case STATUS_FLT_DO_NOT_ATTACH:
		return "0xC01C000F"
	case STATUS_FLT_DO_NOT_DETACH:
		return "0xC01C0010"
	case STATUS_FLT_INSTANCE_ALTITUDE_COLLISION:
		return "0xC01C0011"
	case STATUS_FLT_INSTANCE_NAME_COLLISION:
		return "0xC01C0012"
	case STATUS_FLT_FILTER_NOT_FOUND:
		return "0xC01C0013"
	case STATUS_FLT_VOLUME_NOT_FOUND:
		return "0xC01C0014"
	case STATUS_FLT_INSTANCE_NOT_FOUND:
		return "0xC01C0015"
	case STATUS_FLT_CONTEXT_ALLOCATION_NOT_FOUND:
		return "0xC01C0016"
	case STATUS_FLT_INVALID_CONTEXT_REGISTRATION:
		return "0xC01C0017"
	case STATUS_FLT_NAME_CACHE_MISS:
		return "0xC01C0018"
	case STATUS_FLT_NO_DEVICE_OBJECT:
		return "0xC01C0019"
	case STATUS_FLT_VOLUME_ALREADY_MOUNTED:
		return "0xC01C001A"
	case STATUS_FLT_ALREADY_ENLISTED:
		return "0xC01C001B"
	case STATUS_FLT_CONTEXT_ALREADY_LINKED:
		return "0xC01C001C"
	case STATUS_FLT_NO_WAITER_FOR_REPLY:
		return "0xC01C0020"
	case STATUS_SXS_SECTION_NOT_FOUND:
		return "0xC0150001"
	case STATUS_SXS_CANT_GEN_ACTCTX:
		return "0xC0150002"
	case STATUS_SXS_INVALID_ACTCTXDATA_FORMAT:
		return "0xC0150003"
	case STATUS_SXS_ASSEMBLY_NOT_FOUND:
		return "0xC0150004"
	case STATUS_SXS_MANIFEST_FORMAT_ERROR:
		return "0xC0150005"
	case STATUS_SXS_MANIFEST_PARSE_ERROR:
		return "0xC0150006"
	case STATUS_SXS_ACTIVATION_CONTEXT_DISABLED:
		return "0xC0150007"
	case STATUS_SXS_KEY_NOT_FOUND:
		return "0xC0150008"
	case STATUS_SXS_VERSION_CONFLICT:
		return "0xC0150009"
	case STATUS_SXS_WRONG_SECTION_TYPE:
		return "0xC015000A"
	case STATUS_SXS_THREAD_QUERIES_DISABLED:
		return "0xC015000B"
	case STATUS_SXS_ASSEMBLY_MISSING:
		return "0xC015000C"
	case STATUS_SXS_RELEASE_ACTIVATION_CONTEXT:
		return "0x4015000D"
	case STATUS_SXS_PROCESS_DEFAULT_ALREADY_SET:
		return "0xC015000E"
	case STATUS_SXS_EARLY_DEACTIVATION:
		return "0xC015000F"
	case STATUS_SXS_INVALID_DEACTIVATION:
		return "0xC0150010"
	case STATUS_SXS_MULTIPLE_DEACTIVATION:
		return "0xC0150011"
	case STATUS_SXS_SYSTEM_DEFAULT_ACTIVATION_CONTEXT_EMPTY:
		return "0xC0150012"
	case STATUS_SXS_PROCESS_TERMINATION_REQUESTED:
		return "0xC0150013"
	case STATUS_SXS_CORRUPT_ACTIVATION_STACK:
		return "0xC0150014"
	case STATUS_SXS_CORRUPTION:
		return "0xC0150015"
	case STATUS_SXS_INVALID_IDENTITY_ATTRIBUTE_VALUE:
		return "0xC0150016"
	case STATUS_SXS_INVALID_IDENTITY_ATTRIBUTE_NAME:
		return "0xC0150017"
	case STATUS_SXS_IDENTITY_DUPLICATE_ATTRIBUTE:
		return "0xC0150018"
	case STATUS_SXS_IDENTITY_PARSE_ERROR:
		return "0xC0150019"
	case STATUS_SXS_COMPONENT_STORE_CORRUPT:
		return "0xC015001A"
	case STATUS_SXS_FILE_HASH_MISMATCH:
		return "0xC015001B"
	case STATUS_SXS_MANIFEST_IDENTITY_SAME_BUT_CONTENTS_DIFFERENT:
		return "0xC015001C"
	case STATUS_SXS_IDENTITIES_DIFFERENT:
		return "0xC015001D"
	case STATUS_SXS_ASSEMBLY_IS_NOT_A_DEPLOYMENT:
		return "0xC015001E"
	case STATUS_SXS_FILE_NOT_PART_OF_ASSEMBLY:
		return "0xC015001F"
	case STATUS_ADVANCED_INSTALLER_FAILED:
		return "0xC0150020"
	case STATUS_XML_ENCODING_MISMATCH:
		return "0xC0150021"
	case STATUS_SXS_MANIFEST_TOO_BIG:
		return "0xC0150022"
	case STATUS_SXS_SETTING_NOT_REGISTERED:
		return "0xC0150023"
	case STATUS_SXS_TRANSACTION_CLOSURE_INCOMPLETE:
		return "0xC0150024"
	case STATUS_SMI_PRIMITIVE_INSTALLER_FAILED:
		return "0xC0150025"
	case STATUS_GENERIC_COMMAND_FAILED:
		return "0xC0150026"
	case STATUS_SXS_FILE_HASH_MISSING:
		return "0xC0150027"
	case STATUS_CLUSTER_INVALID_NODE:
		return "0xC0130001"
	case STATUS_CLUSTER_NODE_EXISTS:
		return "0xC0130002"
	case STATUS_CLUSTER_JOIN_IN_PROGRESS:
		return "0xC0130003"
	case STATUS_CLUSTER_NODE_NOT_FOUND:
		return "0xC0130004"
	case STATUS_CLUSTER_LOCAL_NODE_NOT_FOUND:
		return "0xC0130005"
	case STATUS_CLUSTER_NETWORK_EXISTS:
		return "0xC0130006"
	case STATUS_CLUSTER_NETWORK_NOT_FOUND:
		return "0xC0130007"
	case STATUS_CLUSTER_NETINTERFACE_EXISTS:
		return "0xC0130008"
	case STATUS_CLUSTER_NETINTERFACE_NOT_FOUND:
		return "0xC0130009"
	case STATUS_CLUSTER_INVALID_REQUEST:
		return "0xC013000A"
	case STATUS_CLUSTER_INVALID_NETWORK_PROVIDER:
		return "0xC013000B"
	case STATUS_CLUSTER_NODE_DOWN:
		return "0xC013000C"
	case STATUS_CLUSTER_NODE_UNREACHABLE:
		return "0xC013000D"
	case STATUS_CLUSTER_NODE_NOT_MEMBER:
		return "0xC013000E"
	case STATUS_CLUSTER_JOIN_NOT_IN_PROGRESS:
		return "0xC013000F"
	case STATUS_CLUSTER_INVALID_NETWORK:
		return "0xC0130010"
	case STATUS_CLUSTER_NO_NET_ADAPTERS:
		return "0xC0130011"
	case STATUS_CLUSTER_NODE_UP:
		return "0xC0130012"
	case STATUS_CLUSTER_NODE_PAUSED:
		return "0xC0130013"
	case STATUS_CLUSTER_NODE_NOT_PAUSED:
		return "0xC0130014"
	case STATUS_CLUSTER_NO_SECURITY_CONTEXT:
		return "0xC0130015"
	case STATUS_CLUSTER_NETWORK_NOT_INTERNAL:
		return "0xC0130016"
	case STATUS_CLUSTER_POISONED:
		return "0xC0130017"
	case STATUS_CLUSTER_NON_CSV_PATH:
		return "0xC0130018"
	case STATUS_CLUSTER_CSV_VOLUME_NOT_LOCAL:
		return "0xC0130019"
	case STATUS_TRANSACTIONAL_CONFLICT:
		return "0xC0190001"
	case STATUS_INVALID_TRANSACTION:
		return "0xC0190002"
	case STATUS_TRANSACTION_NOT_ACTIVE:
		return "0xC0190003"
	case STATUS_TM_INITIALIZATION_FAILED:
		return "0xC0190004"
	case STATUS_RM_NOT_ACTIVE:
		return "0xC0190005"
	case STATUS_RM_METADATA_CORRUPT:
		return "0xC0190006"
	case STATUS_TRANSACTION_NOT_JOINED:
		return "0xC0190007"
	case STATUS_DIRECTORY_NOT_RM:
		return "0xC0190008"
	case STATUS_COULD_NOT_RESIZE_LOG:
		return "0x80190009"
	case STATUS_TRANSACTIONS_UNSUPPORTED_REMOTE:
		return "0xC019000A"
	case STATUS_LOG_RESIZE_INVALID_SIZE:
		return "0xC019000B"
	case STATUS_REMOTE_FILE_VERSION_MISMATCH:
		return "0xC019000C"
	case STATUS_CRM_PROTOCOL_ALREADY_EXISTS:
		return "0xC019000F"
	case STATUS_TRANSACTION_PROPAGATION_FAILED:
		return "0xC0190010"
	case STATUS_CRM_PROTOCOL_NOT_FOUND:
		return "0xC0190011"
	case STATUS_TRANSACTION_SUPERIOR_EXISTS:
		return "0xC0190012"
	case STATUS_TRANSACTION_REQUEST_NOT_VALID:
		return "0xC0190013"
	case STATUS_TRANSACTION_NOT_REQUESTED:
		return "0xC0190014"
	case STATUS_TRANSACTION_ALREADY_ABORTED:
		return "0xC0190015"
	case STATUS_TRANSACTION_ALREADY_COMMITTED:
		return "0xC0190016"
	case STATUS_TRANSACTION_INVALID_MARSHALL_BUFFER:
		return "0xC0190017"
	case STATUS_CURRENT_TRANSACTION_NOT_VALID:
		return "0xC0190018"
	case STATUS_LOG_GROWTH_FAILED:
		return "0xC0190019"
	case STATUS_OBJECT_NO_LONGER_EXISTS:
		return "0xC0190021"
	case STATUS_STREAM_MINIVERSION_NOT_FOUND:
		return "0xC0190022"
	case STATUS_STREAM_MINIVERSION_NOT_VALID:
		return "0xC0190023"
	case STATUS_MINIVERSION_INACCESSIBLE_FROM_SPECIFIED_TRANSACTION:
		return "0xC0190024"
	case STATUS_CANT_OPEN_MINIVERSION_WITH_MODIFY_INTENT:
		return "0xC0190025"
	case STATUS_CANT_CREATE_MORE_STREAM_MINIVERSIONS:
		return "0xC0190026"
	case STATUS_HANDLE_NO_LONGER_VALID:
		return "0xC0190028"
	case STATUS_NO_TXF_METADATA:
		return "0x80190029"
	case STATUS_LOG_CORRUPTION_DETECTED:
		return "0xC0190030"
	case STATUS_CANT_RECOVER_WITH_HANDLE_OPEN:
		return "0x80190031"
	case STATUS_RM_DISCONNECTED:
		return "0xC0190032"
	case STATUS_ENLISTMENT_NOT_SUPERIOR:
		return "0xC0190033"
	case STATUS_RECOVERY_NOT_NEEDED:
		return "0x40190034"
	case STATUS_RM_ALREADY_STARTED:
		return "0x40190035"
	case STATUS_FILE_IDENTITY_NOT_PERSISTENT:
		return "0xC0190036"
	case STATUS_CANT_BREAK_TRANSACTIONAL_DEPENDENCY:
		return "0xC0190037"
	case STATUS_CANT_CROSS_RM_BOUNDARY:
		return "0xC0190038"
	case STATUS_TXF_DIR_NOT_EMPTY:
		return "0xC0190039"
	case STATUS_INDOUBT_TRANSACTIONS_EXIST:
		return "0xC019003A"
	case STATUS_TM_VOLATILE:
		return "0xC019003B"
	case STATUS_ROLLBACK_TIMER_EXPIRED:
		return "0xC019003C"
	case STATUS_TXF_ATTRIBUTE_CORRUPT:
		return "0xC019003D"
	case STATUS_EFS_NOT_ALLOWED_IN_TRANSACTION:
		return "0xC019003E"
	case STATUS_TRANSACTIONAL_OPEN_NOT_ALLOWED:
		return "0xC019003F"
	case STATUS_TRANSACTED_MAPPING_UNSUPPORTED_REMOTE:
		return "0xC0190040"
	case STATUS_TXF_METADATA_ALREADY_PRESENT:
		return "0x80190041"
	case STATUS_TRANSACTION_SCOPE_CALLBACKS_NOT_SET:
		return "0x80190042"
	case STATUS_TRANSACTION_REQUIRED_PROMOTION:
		return "0xC0190043"
	case STATUS_CANNOT_EXECUTE_FILE_IN_TRANSACTION:
		return "0xC0190044"
	case STATUS_TRANSACTIONS_NOT_FROZEN:
		return "0xC0190045"
	case STATUS_TRANSACTION_FREEZE_IN_PROGRESS:
		return "0xC0190046"
	case STATUS_NOT_SNAPSHOT_VOLUME:
		return "0xC0190047"
	case STATUS_NO_SAVEPOINT_WITH_OPEN_FILES:
		return "0xC0190048"
	case STATUS_SPARSE_NOT_ALLOWED_IN_TRANSACTION:
		return "0xC0190049"
	case STATUS_TM_IDENTITY_MISMATCH:
		return "0xC019004A"
	case STATUS_FLOATED_SECTION:
		return "0xC019004B"
	case STATUS_CANNOT_ACCEPT_TRANSACTED_WORK:
		return "0xC019004C"
	case STATUS_CANNOT_ABORT_TRANSACTIONS:
		return "0xC019004D"
	case STATUS_TRANSACTION_NOT_FOUND:
		return "0xC019004E"
	case STATUS_RESOURCEMANAGER_NOT_FOUND:
		return "0xC019004F"
	case STATUS_ENLISTMENT_NOT_FOUND:
		return "0xC0190050"
	case STATUS_TRANSACTIONMANAGER_NOT_FOUND:
		return "0xC0190051"
	case STATUS_TRANSACTIONMANAGER_NOT_ONLINE:
		return "0xC0190052"
	case STATUS_TRANSACTIONMANAGER_RECOVERY_NAME_COLLISION:
		return "0xC0190053"
	case STATUS_TRANSACTION_NOT_ROOT:
		return "0xC0190054"
	case STATUS_TRANSACTION_OBJECT_EXPIRED:
		return "0xC0190055"
	case STATUS_COMPRESSION_NOT_ALLOWED_IN_TRANSACTION:
		return "0xC0190056"
	case STATUS_TRANSACTION_RESPONSE_NOT_ENLISTED:
		return "0xC0190057"
	case STATUS_TRANSACTION_RECORD_TOO_LONG:
		return "0xC0190058"
	case STATUS_NO_LINK_TRACKING_IN_TRANSACTION:
		return "0xC0190059"
	case STATUS_OPERATION_NOT_SUPPORTED_IN_TRANSACTION:
		return "0xC019005A"
	case STATUS_TRANSACTION_INTEGRITY_VIOLATED:
		return "0xC019005B"
	case STATUS_TRANSACTIONMANAGER_IDENTITY_MISMATCH:
		return "0xC019005C"
	case STATUS_RM_CANNOT_BE_FROZEN_FOR_SNAPSHOT:
		return "0xC019005D"
	case STATUS_TRANSACTION_MUST_WRITETHROUGH:
		return "0xC019005E"
	case STATUS_TRANSACTION_NO_SUPERIOR:
		return "0xC019005F"
	case STATUS_EXPIRED_HANDLE:
		return "0xC0190060"
	case STATUS_TRANSACTION_NOT_ENLISTED:
		return "0xC0190061"
	case STATUS_LOG_SECTOR_INVALID:
		return "0xC01A0001"
	case STATUS_LOG_SECTOR_PARITY_INVALID:
		return "0xC01A0002"
	case STATUS_LOG_SECTOR_REMAPPED:
		return "0xC01A0003"
	case STATUS_LOG_BLOCK_INCOMPLETE:
		return "0xC01A0004"
	case STATUS_LOG_INVALID_RANGE:
		return "0xC01A0005"
	case STATUS_LOG_BLOCKS_EXHAUSTED:
		return "0xC01A0006"
	case STATUS_LOG_READ_CONTEXT_INVALID:
		return "0xC01A0007"
	case STATUS_LOG_RESTART_INVALID:
		return "0xC01A0008"
	case STATUS_LOG_BLOCK_VERSION:
		return "0xC01A0009"
	case STATUS_LOG_BLOCK_INVALID:
		return "0xC01A000A"
	case STATUS_LOG_READ_MODE_INVALID:
		return "0xC01A000B"
	case STATUS_LOG_NO_RESTART:
		return "0x401A000C"
	case STATUS_LOG_METADATA_CORRUPT:
		return "0xC01A000D"
	case STATUS_LOG_METADATA_INVALID:
		return "0xC01A000E"
	case STATUS_LOG_METADATA_INCONSISTENT:
		return "0xC01A000F"
	case STATUS_LOG_RESERVATION_INVALID:
		return "0xC01A0010"
	case STATUS_LOG_CANT_DELETE:
		return "0xC01A0011"
	case STATUS_LOG_CONTAINER_LIMIT_EXCEEDED:
		return "0xC01A0012"
	case STATUS_LOG_START_OF_LOG:
		return "0xC01A0013"
	case STATUS_LOG_POLICY_ALREADY_INSTALLED:
		return "0xC01A0014"
	case STATUS_LOG_POLICY_NOT_INSTALLED:
		return "0xC01A0015"
	case STATUS_LOG_POLICY_INVALID:
		return "0xC01A0016"
	case STATUS_LOG_POLICY_CONFLICT:
		return "0xC01A0017"
	case STATUS_LOG_PINNED_ARCHIVE_TAIL:
		return "0xC01A0018"
	case STATUS_LOG_RECORD_NONEXISTENT:
		return "0xC01A0019"
	case STATUS_LOG_RECORDS_RESERVED_INVALID:
		return "0xC01A001A"
	case STATUS_LOG_SPACE_RESERVED_INVALID:
		return "0xC01A001B"
	case STATUS_LOG_TAIL_INVALID:
		return "0xC01A001C"
	case STATUS_LOG_FULL:
		return "0xC01A001D"
	case STATUS_LOG_MULTIPLEXED:
		return "0xC01A001E"
	case STATUS_LOG_DEDICATED:
		return "0xC01A001F"
	case STATUS_LOG_ARCHIVE_NOT_IN_PROGRESS:
		return "0xC01A0020"
	case STATUS_LOG_ARCHIVE_IN_PROGRESS:
		return "0xC01A0021"
	case STATUS_LOG_EPHEMERAL:
		return "0xC01A0022"
	case STATUS_LOG_NOT_ENOUGH_CONTAINERS:
		return "0xC01A0023"
	case STATUS_LOG_CLIENT_ALREADY_REGISTERED:
		return "0xC01A0024"
	case STATUS_LOG_CLIENT_NOT_REGISTERED:
		return "0xC01A0025"
	case STATUS_LOG_FULL_HANDLER_IN_PROGRESS:
		return "0xC01A0026"
	case STATUS_LOG_CONTAINER_READ_FAILED:
		return "0xC01A0027"
	case STATUS_LOG_CONTAINER_WRITE_FAILED:
		return "0xC01A0028"
	case STATUS_LOG_CONTAINER_OPEN_FAILED:
		return "0xC01A0029"
	case STATUS_LOG_CONTAINER_STATE_INVALID:
		return "0xC01A002A"
	case STATUS_LOG_STATE_INVALID:
		return "0xC01A002B"
	case STATUS_LOG_PINNED:
		return "0xC01A002C"
	case STATUS_LOG_METADATA_FLUSH_FAILED:
		return "0xC01A002D"
	case STATUS_LOG_INCONSISTENT_SECURITY:
		return "0xC01A002E"
	case STATUS_LOG_APPENDED_FLUSH_FAILED:
		return "0xC01A002F"
	case STATUS_LOG_PINNED_RESERVATION:
		return "0xC01A0030"
	case STATUS_VIDEO_HUNG_DISPLAY_DRIVER_THREAD:
		return "0xC01B00EA"
	case STATUS_VIDEO_HUNG_DISPLAY_DRIVER_THREAD_RECOVERED:
		return "0x801B00EB"
	case STATUS_VIDEO_DRIVER_DEBUG_REPORT_REQUEST:
		return "0x401B00EC"
	case STATUS_MONITOR_NO_DESCRIPTOR:
		return "0xC01D0001"
	case STATUS_MONITOR_UNKNOWN_DESCRIPTOR_FORMAT:
		return "0xC01D0002"
	case STATUS_MONITOR_INVALID_DESCRIPTOR_CHECKSUM:
		return "0xC01D0003"
	case STATUS_MONITOR_INVALID_STANDARD_TIMING_BLOCK:
		return "0xC01D0004"
	case STATUS_MONITOR_WMI_DATABLOCK_REGISTRATION_FAILED:
		return "0xC01D0005"
	case STATUS_MONITOR_INVALID_SERIAL_NUMBER_MONDSC_BLOCK:
		return "0xC01D0006"
	case STATUS_MONITOR_INVALID_USER_FRIENDLY_MONDSC_BLOCK:
		return "0xC01D0007"
	case STATUS_MONITOR_NO_MORE_DESCRIPTOR_DATA:
		return "0xC01D0008"
	case STATUS_MONITOR_INVALID_DETAILED_TIMING_BLOCK:
		return "0xC01D0009"
	case STATUS_MONITOR_INVALID_MANUFACTURE_DATE:
		return "0xC01D000A"
	case STATUS_GRAPHICS_NOT_EXCLUSIVE_MODE_OWNER:
		return "0xC01E0000"
	case STATUS_GRAPHICS_INSUFFICIENT_DMA_BUFFER:
		return "0xC01E0001"
	case STATUS_GRAPHICS_INVALID_DISPLAY_ADAPTER:
		return "0xC01E0002"
	case STATUS_GRAPHICS_ADAPTER_WAS_RESET:
		return "0xC01E0003"
	case STATUS_GRAPHICS_INVALID_DRIVER_MODEL:
		return "0xC01E0004"
	case STATUS_GRAPHICS_PRESENT_MODE_CHANGED:
		return "0xC01E0005"
	case STATUS_GRAPHICS_PRESENT_OCCLUDED:
		return "0xC01E0006"
	case STATUS_GRAPHICS_PRESENT_DENIED:
		return "0xC01E0007"
	case STATUS_GRAPHICS_CANNOTCOLORCONVERT:
		return "0xC01E0008"
	case STATUS_GRAPHICS_DRIVER_MISMATCH:
		return "0xC01E0009"
	case STATUS_GRAPHICS_PARTIAL_DATA_POPULATED:
		return "0x401E000A"
	case STATUS_GRAPHICS_PRESENT_REDIRECTION_DISABLED:
		return "0xC01E000B"
	case STATUS_GRAPHICS_PRESENT_UNOCCLUDED:
		return "0xC01E000C"
	case STATUS_GRAPHICS_NO_VIDEO_MEMORY:
		return "0xC01E0100"
	case STATUS_GRAPHICS_CANT_LOCK_MEMORY:
		return "0xC01E0101"
	case STATUS_GRAPHICS_ALLOCATION_BUSY:
		return "0xC01E0102"
	case STATUS_GRAPHICS_TOO_MANY_REFERENCES:
		return "0xC01E0103"
	case STATUS_GRAPHICS_TRY_AGAIN_LATER:
		return "0xC01E0104"
	case STATUS_GRAPHICS_TRY_AGAIN_NOW:
		return "0xC01E0105"
	case STATUS_GRAPHICS_ALLOCATION_INVALID:
		return "0xC01E0106"
	case STATUS_GRAPHICS_UNSWIZZLING_APERTURE_UNAVAILABLE:
		return "0xC01E0107"
	case STATUS_GRAPHICS_UNSWIZZLING_APERTURE_UNSUPPORTED:
		return "0xC01E0108"
	case STATUS_GRAPHICS_CANT_EVICT_PINNED_ALLOCATION:
		return "0xC01E0109"
	case STATUS_GRAPHICS_INVALID_ALLOCATION_USAGE:
		return "0xC01E0110"
	case STATUS_GRAPHICS_CANT_RENDER_LOCKED_ALLOCATION:
		return "0xC01E0111"
	case STATUS_GRAPHICS_ALLOCATION_CLOSED:
		return "0xC01E0112"
	case STATUS_GRAPHICS_INVALID_ALLOCATION_INSTANCE:
		return "0xC01E0113"
	case STATUS_GRAPHICS_INVALID_ALLOCATION_HANDLE:
		return "0xC01E0114"
	case STATUS_GRAPHICS_WRONG_ALLOCATION_DEVICE:
		return "0xC01E0115"
	case STATUS_GRAPHICS_ALLOCATION_CONTENT_LOST:
		return "0xC01E0116"
	case STATUS_GRAPHICS_GPU_EXCEPTION_ON_DEVICE:
		return "0xC01E0200"
	case STATUS_GRAPHICS_INVALID_VIDPN_TOPOLOGY:
		return "0xC01E0300"
	case STATUS_GRAPHICS_VIDPN_TOPOLOGY_NOT_SUPPORTED:
		return "0xC01E0301"
	case STATUS_GRAPHICS_VIDPN_TOPOLOGY_CURRENTLY_NOT_SUPPORTED:
		return "0xC01E0302"
	case STATUS_GRAPHICS_INVALID_VIDPN:
		return "0xC01E0303"
	case STATUS_GRAPHICS_INVALID_VIDEO_PRESENT_SOURCE:
		return "0xC01E0304"
	case STATUS_GRAPHICS_INVALID_VIDEO_PRESENT_TARGET:
		return "0xC01E0305"
	case STATUS_GRAPHICS_VIDPN_MODALITY_NOT_SUPPORTED:
		return "0xC01E0306"
	case STATUS_GRAPHICS_MODE_NOT_PINNED:
		return "0x401E0307"
	case STATUS_GRAPHICS_INVALID_VIDPN_SOURCEMODESET:
		return "0xC01E0308"
	case STATUS_GRAPHICS_INVALID_VIDPN_TARGETMODESET:
		return "0xC01E0309"
	case STATUS_GRAPHICS_INVALID_FREQUENCY:
		return "0xC01E030A"
	case STATUS_GRAPHICS_INVALID_ACTIVE_REGION:
		return "0xC01E030B"
	case STATUS_GRAPHICS_INVALID_TOTAL_REGION:
		return "0xC01E030C"
	case STATUS_GRAPHICS_INVALID_VIDEO_PRESENT_SOURCE_MODE:
		return "0xC01E0310"
	case STATUS_GRAPHICS_INVALID_VIDEO_PRESENT_TARGET_MODE:
		return "0xC01E0311"
	case STATUS_GRAPHICS_PINNED_MODE_MUST_REMAIN_IN_SET:
		return "0xC01E0312"
	case STATUS_GRAPHICS_PATH_ALREADY_IN_TOPOLOGY:
		return "0xC01E0313"
	case STATUS_GRAPHICS_MODE_ALREADY_IN_MODESET:
		return "0xC01E0314"
	case STATUS_GRAPHICS_INVALID_VIDEOPRESENTSOURCESET:
		return "0xC01E0315"
	case STATUS_GRAPHICS_INVALID_VIDEOPRESENTTARGETSET:
		return "0xC01E0316"
	case STATUS_GRAPHICS_SOURCE_ALREADY_IN_SET:
		return "0xC01E0317"
	case STATUS_GRAPHICS_TARGET_ALREADY_IN_SET:
		return "0xC01E0318"
	case STATUS_GRAPHICS_INVALID_VIDPN_PRESENT_PATH:
		return "0xC01E0319"
	case STATUS_GRAPHICS_NO_RECOMMENDED_VIDPN_TOPOLOGY:
		return "0xC01E031A"
	case STATUS_GRAPHICS_INVALID_MONITOR_FREQUENCYRANGESET:
		return "0xC01E031B"
	case STATUS_GRAPHICS_INVALID_MONITOR_FREQUENCYRANGE:
		return "0xC01E031C"
	case STATUS_GRAPHICS_FREQUENCYRANGE_NOT_IN_SET:
		return "0xC01E031D"
	case STATUS_GRAPHICS_NO_PREFERRED_MODE:
		return "0x401E031E"
	case STATUS_GRAPHICS_FREQUENCYRANGE_ALREADY_IN_SET:
		return "0xC01E031F"
	case STATUS_GRAPHICS_STALE_MODESET:
		return "0xC01E0320"
	case STATUS_GRAPHICS_INVALID_MONITOR_SOURCEMODESET:
		return "0xC01E0321"
	case STATUS_GRAPHICS_INVALID_MONITOR_SOURCE_MODE:
		return "0xC01E0322"
	case STATUS_GRAPHICS_NO_RECOMMENDED_FUNCTIONAL_VIDPN:
		return "0xC01E0323"
	case STATUS_GRAPHICS_MODE_ID_MUST_BE_UNIQUE:
		return "0xC01E0324"
	case STATUS_GRAPHICS_EMPTY_ADAPTER_MONITOR_MODE_SUPPORT_INTERSECTION:
		return "0xC01E0325"
	case STATUS_GRAPHICS_VIDEO_PRESENT_TARGETS_LESS_THAN_SOURCES:
		return "0xC01E0326"
	case STATUS_GRAPHICS_PATH_NOT_IN_TOPOLOGY:
		return "0xC01E0327"
	case STATUS_GRAPHICS_ADAPTER_MUST_HAVE_AT_LEAST_ONE_SOURCE:
		return "0xC01E0328"
	case STATUS_GRAPHICS_ADAPTER_MUST_HAVE_AT_LEAST_ONE_TARGET:
		return "0xC01E0329"
	case STATUS_GRAPHICS_INVALID_MONITORDESCRIPTORSET:
		return "0xC01E032A"
	case STATUS_GRAPHICS_INVALID_MONITORDESCRIPTOR:
		return "0xC01E032B"
	case STATUS_GRAPHICS_MONITORDESCRIPTOR_NOT_IN_SET:
		return "0xC01E032C"
	case STATUS_GRAPHICS_MONITORDESCRIPTOR_ALREADY_IN_SET:
		return "0xC01E032D"
	case STATUS_GRAPHICS_MONITORDESCRIPTOR_ID_MUST_BE_UNIQUE:
		return "0xC01E032E"
	case STATUS_GRAPHICS_INVALID_VIDPN_TARGET_SUBSET_TYPE:
		return "0xC01E032F"
	case STATUS_GRAPHICS_RESOURCES_NOT_RELATED:
		return "0xC01E0330"
	case STATUS_GRAPHICS_SOURCE_ID_MUST_BE_UNIQUE:
		return "0xC01E0331"
	case STATUS_GRAPHICS_TARGET_ID_MUST_BE_UNIQUE:
		return "0xC01E0332"
	case STATUS_GRAPHICS_NO_AVAILABLE_VIDPN_TARGET:
		return "0xC01E0333"
	case STATUS_GRAPHICS_MONITOR_COULD_NOT_BE_ASSOCIATED_WITH_ADAPTER:
		return "0xC01E0334"
	case STATUS_GRAPHICS_NO_VIDPNMGR:
		return "0xC01E0335"
	case STATUS_GRAPHICS_NO_ACTIVE_VIDPN:
		return "0xC01E0336"
	case STATUS_GRAPHICS_STALE_VIDPN_TOPOLOGY:
		return "0xC01E0337"
	case STATUS_GRAPHICS_MONITOR_NOT_CONNECTED:
		return "0xC01E0338"
	case STATUS_GRAPHICS_SOURCE_NOT_IN_TOPOLOGY:
		return "0xC01E0339"
	case STATUS_GRAPHICS_INVALID_PRIMARYSURFACE_SIZE:
		return "0xC01E033A"
	case STATUS_GRAPHICS_INVALID_VISIBLEREGION_SIZE:
		return "0xC01E033B"
	case STATUS_GRAPHICS_INVALID_STRIDE:
		return "0xC01E033C"
	case STATUS_GRAPHICS_INVALID_PIXELFORMAT:
		return "0xC01E033D"
	case STATUS_GRAPHICS_INVALID_COLORBASIS:
		return "0xC01E033E"
	case STATUS_GRAPHICS_INVALID_PIXELVALUEACCESSMODE:
		return "0xC01E033F"
	case STATUS_GRAPHICS_TARGET_NOT_IN_TOPOLOGY:
		return "0xC01E0340"
	case STATUS_GRAPHICS_NO_DISPLAY_MODE_MANAGEMENT_SUPPORT:
		return "0xC01E0341"
	case STATUS_GRAPHICS_VIDPN_SOURCE_IN_USE:
		return "0xC01E0342"
	case STATUS_GRAPHICS_CANT_ACCESS_ACTIVE_VIDPN:
		return "0xC01E0343"
	case STATUS_GRAPHICS_INVALID_PATH_IMPORTANCE_ORDINAL:
		return "0xC01E0344"
	case STATUS_GRAPHICS_INVALID_PATH_CONTENT_GEOMETRY_TRANSFORMATION:
		return "0xC01E0345"
	case STATUS_GRAPHICS_PATH_CONTENT_GEOMETRY_TRANSFORMATION_NOT_SUPPORTED:
		return "0xC01E0346"
	case STATUS_GRAPHICS_INVALID_GAMMA_RAMP:
		return "0xC01E0347"
	case STATUS_GRAPHICS_GAMMA_RAMP_NOT_SUPPORTED:
		return "0xC01E0348"
	case STATUS_GRAPHICS_MULTISAMPLING_NOT_SUPPORTED:
		return "0xC01E0349"
	case STATUS_GRAPHICS_MODE_NOT_IN_MODESET:
		return "0xC01E034A"
	case STATUS_GRAPHICS_DATASET_IS_EMPTY:
		return "0x401E034B"
	case STATUS_GRAPHICS_NO_MORE_ELEMENTS_IN_DATASET:
		return "0x401E034C"
	case STATUS_GRAPHICS_INVALID_VIDPN_TOPOLOGY_RECOMMENDATION_REASON:
		return "0xC01E034D"
	case STATUS_GRAPHICS_INVALID_PATH_CONTENT_TYPE:
		return "0xC01E034E"
	case STATUS_GRAPHICS_INVALID_COPYPROTECTION_TYPE:
		return "0xC01E034F"
	case STATUS_GRAPHICS_UNASSIGNED_MODESET_ALREADY_EXISTS:
		return "0xC01E0350"
	case STATUS_GRAPHICS_PATH_CONTENT_GEOMETRY_TRANSFORMATION_NOT_PINNED:
		return "0x401E0351"
	case STATUS_GRAPHICS_INVALID_SCANLINE_ORDERING:
		return "0xC01E0352"
	case STATUS_GRAPHICS_TOPOLOGY_CHANGES_NOT_ALLOWED:
		return "0xC01E0353"
	case STATUS_GRAPHICS_NO_AVAILABLE_IMPORTANCE_ORDINALS:
		return "0xC01E0354"
	case STATUS_GRAPHICS_INCOMPATIBLE_PRIVATE_FORMAT:
		return "0xC01E0355"
	case STATUS_GRAPHICS_INVALID_MODE_PRUNING_ALGORITHM:
		return "0xC01E0356"
	case STATUS_GRAPHICS_INVALID_MONITOR_CAPABILITY_ORIGIN:
		return "0xC01E0357"
	case STATUS_GRAPHICS_INVALID_MONITOR_FREQUENCYRANGE_CONSTRAINT:
		return "0xC01E0358"
	case STATUS_GRAPHICS_MAX_NUM_PATHS_REACHED:
		return "0xC01E0359"
	case STATUS_GRAPHICS_CANCEL_VIDPN_TOPOLOGY_AUGMENTATION:
		return "0xC01E035A"
	case STATUS_GRAPHICS_INVALID_CLIENT_TYPE:
		return "0xC01E035B"
	case STATUS_GRAPHICS_CLIENTVIDPN_NOT_SET:
		return "0xC01E035C"
	case STATUS_GRAPHICS_SPECIFIED_CHILD_ALREADY_CONNECTED:
		return "0xC01E0400"
	case STATUS_GRAPHICS_CHILD_DESCRIPTOR_NOT_SUPPORTED:
		return "0xC01E0401"
	case STATUS_GRAPHICS_UNKNOWN_CHILD_STATUS:
		return "0x401E042F"
	case STATUS_GRAPHICS_NOT_A_LINKED_ADAPTER:
		return "0xC01E0430"
	case STATUS_GRAPHICS_LEADLINK_NOT_ENUMERATED:
		return "0xC01E0431"
	case STATUS_GRAPHICS_CHAINLINKS_NOT_ENUMERATED:
		return "0xC01E0432"
	case STATUS_GRAPHICS_ADAPTER_CHAIN_NOT_READY:
		return "0xC01E0433"
	case STATUS_GRAPHICS_CHAINLINKS_NOT_STARTED:
		return "0xC01E0434"
	case STATUS_GRAPHICS_CHAINLINKS_NOT_POWERED_ON:
		return "0xC01E0435"
	case STATUS_GRAPHICS_INCONSISTENT_DEVICE_LINK_STATE:
		return "0xC01E0436"
	case STATUS_GRAPHICS_LEADLINK_START_DEFERRED:
		return "0x401E0437"
	case STATUS_GRAPHICS_NOT_POST_DEVICE_DRIVER:
		return "0xC01E0438"
	case STATUS_GRAPHICS_POLLING_TOO_FREQUENTLY:
		return "0x401E0439"
	case STATUS_GRAPHICS_START_DEFERRED:
		return "0x401E043A"
	case STATUS_GRAPHICS_ADAPTER_ACCESS_NOT_EXCLUDED:
		return "0xC01E043B"
	case STATUS_GRAPHICS_OPM_NOT_SUPPORTED:
		return "0xC01E0500"
	case STATUS_GRAPHICS_COPP_NOT_SUPPORTED:
		return "0xC01E0501"
	case STATUS_GRAPHICS_UAB_NOT_SUPPORTED:
		return "0xC01E0502"
	case STATUS_GRAPHICS_OPM_INVALID_ENCRYPTED_PARAMETERS:
		return "0xC01E0503"
	case STATUS_GRAPHICS_OPM_NO_PROTECTED_OUTPUTS_EXIST:
		return "0xC01E0505"
	case STATUS_GRAPHICS_OPM_INTERNAL_ERROR:
		return "0xC01E050B"
	case STATUS_GRAPHICS_OPM_INVALID_HANDLE:
		return "0xC01E050C"
	case STATUS_GRAPHICS_PVP_INVALID_CERTIFICATE_LENGTH:
		return "0xC01E050E"
	case STATUS_GRAPHICS_OPM_SPANNING_MODE_ENABLED:
		return "0xC01E050F"
	case STATUS_GRAPHICS_OPM_THEATER_MODE_ENABLED:
		return "0xC01E0510"
	case STATUS_GRAPHICS_PVP_HFS_FAILED:
		return "0xC01E0511"
	case STATUS_GRAPHICS_OPM_INVALID_SRM:
		return "0xC01E0512"
	case STATUS_GRAPHICS_OPM_OUTPUT_DOES_NOT_SUPPORT_HDCP:
		return "0xC01E0513"
	case STATUS_GRAPHICS_OPM_OUTPUT_DOES_NOT_SUPPORT_ACP:
		return "0xC01E0514"
	case STATUS_GRAPHICS_OPM_OUTPUT_DOES_NOT_SUPPORT_CGMSA:
		return "0xC01E0515"
	case STATUS_GRAPHICS_OPM_HDCP_SRM_NEVER_SET:
		return "0xC01E0516"
	case STATUS_GRAPHICS_OPM_RESOLUTION_TOO_HIGH:
		return "0xC01E0517"
	case STATUS_GRAPHICS_OPM_ALL_HDCP_HARDWARE_ALREADY_IN_USE:
		return "0xC01E0518"
	case STATUS_GRAPHICS_OPM_PROTECTED_OUTPUT_NO_LONGER_EXISTS:
		return "0xC01E051A"
	case STATUS_GRAPHICS_OPM_PROTECTED_OUTPUT_DOES_NOT_HAVE_COPP_SEMANTICS:
		return "0xC01E051C"
	case STATUS_GRAPHICS_OPM_INVALID_INFORMATION_REQUEST:
		return "0xC01E051D"
	case STATUS_GRAPHICS_OPM_DRIVER_INTERNAL_ERROR:
		return "0xC01E051E"
	case STATUS_GRAPHICS_OPM_PROTECTED_OUTPUT_DOES_NOT_HAVE_OPM_SEMANTICS:
		return "0xC01E051F"
	case STATUS_GRAPHICS_OPM_SIGNALING_NOT_SUPPORTED:
		return "0xC01E0520"
	case STATUS_GRAPHICS_OPM_INVALID_CONFIGURATION_REQUEST:
		return "0xC01E0521"
	case STATUS_GRAPHICS_I2C_NOT_SUPPORTED:
		return "0xC01E0580"
	case STATUS_GRAPHICS_I2C_DEVICE_DOES_NOT_EXIST:
		return "0xC01E0581"
	case STATUS_GRAPHICS_I2C_ERROR_TRANSMITTING_DATA:
		return "0xC01E0582"
	case STATUS_GRAPHICS_I2C_ERROR_RECEIVING_DATA:
		return "0xC01E0583"
	case STATUS_GRAPHICS_DDCCI_VCP_NOT_SUPPORTED:
		return "0xC01E0584"
	case STATUS_GRAPHICS_DDCCI_INVALID_DATA:
		return "0xC01E0585"
	case STATUS_GRAPHICS_DDCCI_MONITOR_RETURNED_INVALID_TIMING_STATUS_BYTE:
		return "0xC01E0586"
	case STATUS_GRAPHICS_DDCCI_INVALID_CAPABILITIES_STRING:
		return "0xC01E0587"
	case STATUS_GRAPHICS_MCA_INTERNAL_ERROR:
		return "0xC01E0588"
	case STATUS_GRAPHICS_DDCCI_INVALID_MESSAGE_COMMAND:
		return "0xC01E0589"
	case STATUS_GRAPHICS_DDCCI_INVALID_MESSAGE_LENGTH:
		return "0xC01E058A"
	case STATUS_GRAPHICS_DDCCI_INVALID_MESSAGE_CHECKSUM:
		return "0xC01E058B"
	case STATUS_GRAPHICS_INVALID_PHYSICAL_MONITOR_HANDLE:
		return "0xC01E058C"
	case STATUS_GRAPHICS_MONITOR_NO_LONGER_EXISTS:
		return "0xC01E058D"
	case STATUS_GRAPHICS_ONLY_CONSOLE_SESSION_SUPPORTED:
		return "0xC01E05E0"
	case STATUS_GRAPHICS_NO_DISPLAY_DEVICE_CORRESPONDS_TO_NAME:
		return "0xC01E05E1"
	case STATUS_GRAPHICS_DISPLAY_DEVICE_NOT_ATTACHED_TO_DESKTOP:
		return "0xC01E05E2"
	case STATUS_GRAPHICS_MIRRORING_DEVICES_NOT_SUPPORTED:
		return "0xC01E05E3"
	case STATUS_GRAPHICS_INVALID_POINTER:
		return "0xC01E05E4"
	case STATUS_GRAPHICS_NO_MONITORS_CORRESPOND_TO_DISPLAY_DEVICE:
		return "0xC01E05E5"
	case STATUS_GRAPHICS_PARAMETER_ARRAY_TOO_SMALL:
		return "0xC01E05E6"
	case STATUS_GRAPHICS_INTERNAL_ERROR:
		return "0xC01E05E7"
	case STATUS_GRAPHICS_SESSION_TYPE_CHANGE_IN_PROGRESS:
		return "0xC01E05E8"
	case STATUS_FVE_LOCKED_VOLUME:
		return "0xC0210000"
	case STATUS_FVE_NOT_ENCRYPTED:
		return "0xC0210001"
	case STATUS_FVE_BAD_INFORMATION:
		return "0xC0210002"
	case STATUS_FVE_TOO_SMALL:
		return "0xC0210003"
	case STATUS_FVE_FAILED_WRONG_FS:
		return "0xC0210004"
	case STATUS_FVE_BAD_PARTITION_SIZE:
		return "0xC0210005"
	case STATUS_FVE_FS_NOT_EXTENDED:
		return "0xC0210006"
	case STATUS_FVE_FS_MOUNTED:
		return "0xC0210007"
	case STATUS_FVE_NO_LICENSE:
		return "0xC0210008"
	case STATUS_FVE_ACTION_NOT_ALLOWED:
		return "0xC0210009"
	case STATUS_FVE_BAD_DATA:
		return "0xC021000A"
	case STATUS_FVE_VOLUME_NOT_BOUND:
		return "0xC021000B"
	case STATUS_FVE_NOT_DATA_VOLUME:
		return "0xC021000C"
	case STATUS_FVE_CONV_READ_ERROR:
		return "0xC021000D"
	case STATUS_FVE_CONV_WRITE_ERROR:
		return "0xC021000E"
	case STATUS_FVE_OVERLAPPED_UPDATE:
		return "0xC021000F"
	case STATUS_FVE_FAILED_SECTOR_SIZE:
		return "0xC0210010"
	case STATUS_FVE_FAILED_AUTHENTICATION:
		return "0xC0210011"
	case STATUS_FVE_NOT_OS_VOLUME:
		return "0xC0210012"
	case STATUS_FVE_KEYFILE_NOT_FOUND:
		return "0xC0210013"
	case STATUS_FVE_KEYFILE_INVALID:
		return "0xC0210014"
	case STATUS_FVE_KEYFILE_NO_VMK:
		return "0xC0210015"
	case STATUS_FVE_TPM_DISABLED:
		return "0xC0210016"
	case STATUS_FVE_TPM_SRK_AUTH_NOT_ZERO:
		return "0xC0210017"
	case STATUS_FVE_TPM_INVALID_PCR:
		return "0xC0210018"
	case STATUS_FVE_TPM_NO_VMK:
		return "0xC0210019"
	case STATUS_FVE_PIN_INVALID:
		return "0xC021001A"
	case STATUS_FVE_AUTH_INVALID_APPLICATION:
		return "0xC021001B"
	case STATUS_FVE_AUTH_INVALID_CONFIG:
		return "0xC021001C"
	case STATUS_FVE_DEBUGGER_ENABLED:
		return "0xC021001D"
	case STATUS_FVE_DRY_RUN_FAILED:
		return "0xC021001E"
	case STATUS_FVE_BAD_METADATA_POINTER:
		return "0xC021001F"
	case STATUS_FVE_OLD_METADATA_COPY:
		return "0xC0210020"
	case STATUS_FVE_REBOOT_REQUIRED:
		return "0xC0210021"
	case STATUS_FVE_RAW_ACCESS:
		return "0xC0210022"
	case STATUS_FVE_RAW_BLOCKED:
		return "0xC0210023"
	case STATUS_FVE_NO_AUTOUNLOCK_MASTER_KEY:
		return "0xC0210024"
	case STATUS_FVE_MOR_FAILED:
		return "0xC0210025"
	case STATUS_FVE_NO_FEATURE_LICENSE:
		return "0xC0210026"
	case STATUS_FVE_POLICY_USER_DISABLE_RDV_NOT_ALLOWED:
		return "0xC0210027"
	case STATUS_FVE_CONV_RECOVERY_FAILED:
		return "0xC0210028"
	case STATUS_FVE_VIRTUALIZED_SPACE_TOO_BIG:
		return "0xC0210029"
	case STATUS_FVE_INVALID_DATUM_TYPE:
		return "0xC021002A"
	case STATUS_FVE_VOLUME_TOO_SMALL:
		return "0xC0210030"
	case STATUS_FVE_ENH_PIN_INVALID:
		return "0xC0210031"
	case STATUS_FWP_CALLOUT_NOT_FOUND:
		return "0xC0220001"
	case STATUS_FWP_CONDITION_NOT_FOUND:
		return "0xC0220002"
	case STATUS_FWP_FILTER_NOT_FOUND:
		return "0xC0220003"
	case STATUS_FWP_LAYER_NOT_FOUND:
		return "0xC0220004"
	case STATUS_FWP_PROVIDER_NOT_FOUND:
		return "0xC0220005"
	case STATUS_FWP_PROVIDER_CONTEXT_NOT_FOUND:
		return "0xC0220006"
	case STATUS_FWP_SUBLAYER_NOT_FOUND:
		return "0xC0220007"
	case STATUS_FWP_NOT_FOUND:
		return "0xC0220008"
	case STATUS_FWP_ALREADY_EXISTS:
		return "0xC0220009"
	case STATUS_FWP_IN_USE:
		return "0xC022000A"
	case STATUS_FWP_DYNAMIC_SESSION_IN_PROGRESS:
		return "0xC022000B"
	case STATUS_FWP_WRONG_SESSION:
		return "0xC022000C"
	case STATUS_FWP_NO_TXN_IN_PROGRESS:
		return "0xC022000D"
	case STATUS_FWP_TXN_IN_PROGRESS:
		return "0xC022000E"
	case STATUS_FWP_TXN_ABORTED:
		return "0xC022000F"
	case STATUS_FWP_SESSION_ABORTED:
		return "0xC0220010"
	case STATUS_FWP_INCOMPATIBLE_TXN:
		return "0xC0220011"
	case STATUS_FWP_TIMEOUT:
		return "0xC0220012"
	case STATUS_FWP_NET_EVENTS_DISABLED:
		return "0xC0220013"
	case STATUS_FWP_INCOMPATIBLE_LAYER:
		return "0xC0220014"
	case STATUS_FWP_KM_CLIENTS_ONLY:
		return "0xC0220015"
	case STATUS_FWP_LIFETIME_MISMATCH:
		return "0xC0220016"
	case STATUS_FWP_BUILTIN_OBJECT:
		return "0xC0220017"
	case STATUS_FWP_TOO_MANY_CALLOUTS:
		return "0xC0220018"
	case STATUS_FWP_NOTIFICATION_DROPPED:
		return "0xC0220019"
	case STATUS_FWP_TRAFFIC_MISMATCH:
		return "0xC022001A"
	case STATUS_FWP_INCOMPATIBLE_SA_STATE:
		return "0xC022001B"
	case STATUS_FWP_NULL_POINTER:
		return "0xC022001C"
	case STATUS_FWP_INVALID_ENUMERATOR:
		return "0xC022001D"
	case STATUS_FWP_INVALID_FLAGS:
		return "0xC022001E"
	case STATUS_FWP_INVALID_NET_MASK:
		return "0xC022001F"
	case STATUS_FWP_INVALID_RANGE:
		return "0xC0220020"
	case STATUS_FWP_INVALID_INTERVAL:
		return "0xC0220021"
	case STATUS_FWP_ZERO_LENGTH_ARRAY:
		return "0xC0220022"
	case STATUS_FWP_NULL_DISPLAY_NAME:
		return "0xC0220023"
	case STATUS_FWP_INVALID_ACTION_TYPE:
		return "0xC0220024"
	case STATUS_FWP_INVALID_WEIGHT:
		return "0xC0220025"
	case STATUS_FWP_MATCH_TYPE_MISMATCH:
		return "0xC0220026"
	case STATUS_FWP_TYPE_MISMATCH:
		return "0xC0220027"
	case STATUS_FWP_OUT_OF_BOUNDS:
		return "0xC0220028"
	case STATUS_FWP_RESERVED:
		return "0xC0220029"
	case STATUS_FWP_DUPLICATE_CONDITION:
		return "0xC022002A"
	case STATUS_FWP_DUPLICATE_KEYMOD:
		return "0xC022002B"
	case STATUS_FWP_ACTION_INCOMPATIBLE_WITH_LAYER:
		return "0xC022002C"
	case STATUS_FWP_ACTION_INCOMPATIBLE_WITH_SUBLAYER:
		return "0xC022002D"
	case STATUS_FWP_CONTEXT_INCOMPATIBLE_WITH_LAYER:
		return "0xC022002E"
	case STATUS_FWP_CONTEXT_INCOMPATIBLE_WITH_CALLOUT:
		return "0xC022002F"
	case STATUS_FWP_INCOMPATIBLE_AUTH_METHOD:
		return "0xC0220030"
	case STATUS_FWP_INCOMPATIBLE_DH_GROUP:
		return "0xC0220031"
	case STATUS_FWP_EM_NOT_SUPPORTED:
		return "0xC0220032"
	case STATUS_FWP_NEVER_MATCH:
		return "0xC0220033"
	case STATUS_FWP_PROVIDER_CONTEXT_MISMATCH:
		return "0xC0220034"
	case STATUS_FWP_INVALID_PARAMETER:
		return "0xC0220035"
	case STATUS_FWP_TOO_MANY_SUBLAYERS:
		return "0xC0220036"
	case STATUS_FWP_CALLOUT_NOTIFICATION_FAILED:
		return "0xC0220037"
	case STATUS_FWP_INVALID_AUTH_TRANSFORM:
		return "0xC0220038"
	case STATUS_FWP_INVALID_CIPHER_TRANSFORM:
		return "0xC0220039"
	case STATUS_FWP_INCOMPATIBLE_CIPHER_TRANSFORM:
		return "0xC022003A"
	case STATUS_FWP_INVALID_TRANSFORM_COMBINATION:
		return "0xC022003B"
	case STATUS_FWP_DUPLICATE_AUTH_METHOD:
		return "0xC022003C"
	case STATUS_FWP_TCPIP_NOT_READY:
		return "0xC0220100"
	case STATUS_FWP_INJECT_HANDLE_CLOSING:
		return "0xC0220101"
	case STATUS_FWP_INJECT_HANDLE_STALE:
		return "0xC0220102"
	case STATUS_FWP_CANNOT_PEND:
		return "0xC0220103"
	case STATUS_FWP_DROP_NOICMP:
		return "0xC0220104"
	case STATUS_NDIS_CLOSING:
		return "0xC0230002"
	case STATUS_NDIS_BAD_VERSION:
		return "0xC0230004"
	case STATUS_NDIS_BAD_CHARACTERISTICS:
		return "0xC0230005"
	case STATUS_NDIS_ADAPTER_NOT_FOUND:
		return "0xC0230006"
	case STATUS_NDIS_OPEN_FAILED:
		return "0xC0230007"
	case STATUS_NDIS_DEVICE_FAILED:
		return "0xC0230008"
	case STATUS_NDIS_MULTICAST_FULL:
		return "0xC0230009"
	case STATUS_NDIS_MULTICAST_EXISTS:
		return "0xC023000A"
	case STATUS_NDIS_MULTICAST_NOT_FOUND:
		return "0xC023000B"
	case STATUS_NDIS_REQUEST_ABORTED:
		return "0xC023000C"
	case STATUS_NDIS_RESET_IN_PROGRESS:
		return "0xC023000D"
	case STATUS_NDIS_NOT_SUPPORTED:
		return "0xC02300BB"
	case STATUS_NDIS_INVALID_PACKET:
		return "0xC023000F"
	case STATUS_NDIS_ADAPTER_NOT_READY:
		return "0xC0230011"
	case STATUS_NDIS_INVALID_LENGTH:
		return "0xC0230014"
	case STATUS_NDIS_INVALID_DATA:
		return "0xC0230015"
	case STATUS_NDIS_BUFFER_TOO_SHORT:
		return "0xC0230016"
	case STATUS_NDIS_INVALID_OID:
		return "0xC0230017"
	case STATUS_NDIS_ADAPTER_REMOVED:
		return "0xC0230018"
	case STATUS_NDIS_UNSUPPORTED_MEDIA:
		return "0xC0230019"
	case STATUS_NDIS_GROUP_ADDRESS_IN_USE:
		return "0xC023001A"
	case STATUS_NDIS_FILE_NOT_FOUND:
		return "0xC023001B"
	case STATUS_NDIS_ERROR_READING_FILE:
		return "0xC023001C"
	case STATUS_NDIS_ALREADY_MAPPED:
		return "0xC023001D"
	case STATUS_NDIS_RESOURCE_CONFLICT:
		return "0xC023001E"
	case STATUS_NDIS_MEDIA_DISCONNECTED:
		return "0xC023001F"
	case STATUS_NDIS_INVALID_ADDRESS:
		return "0xC0230022"
	case STATUS_NDIS_INVALID_DEVICE_REQUEST:
		return "0xC0230010"
	case STATUS_NDIS_PAUSED:
		return "0xC023002A"
	case STATUS_NDIS_INTERFACE_NOT_FOUND:
		return "0xC023002B"
	case STATUS_NDIS_UNSUPPORTED_REVISION:
		return "0xC023002C"
	case STATUS_NDIS_INVALID_PORT:
		return "0xC023002D"
	case STATUS_NDIS_INVALID_PORT_STATE:
		return "0xC023002E"
	case STATUS_NDIS_LOW_POWER_STATE:
		return "0xC023002F"
	case STATUS_NDIS_DOT11_AUTO_CONFIG_ENABLED:
		return "0xC0232000"
	case STATUS_NDIS_DOT11_MEDIA_IN_USE:
		return "0xC0232001"
	case STATUS_NDIS_DOT11_POWER_STATE_INVALID:
		return "0xC0232002"
	case STATUS_NDIS_PM_WOL_PATTERN_LIST_FULL:
		return "0xC0232003"
	case STATUS_NDIS_PM_PROTOCOL_OFFLOAD_LIST_FULL:
		return "0xC0232004"
	case STATUS_NDIS_INDICATION_REQUIRED:
		return "0x40230001"
	case STATUS_NDIS_OFFLOAD_POLICY:
		return "0xC023100F"
	case STATUS_NDIS_OFFLOAD_CONNECTION_REJECTED:
		return "0xC0231012"
	case STATUS_NDIS_OFFLOAD_PATH_REJECTED:
		return "0xC0231013"
	case STATUS_HV_INVALID_HYPERCALL_CODE:
		return "0xC0350002"
	case STATUS_HV_INVALID_HYPERCALL_INPUT:
		return "0xC0350003"
	case STATUS_HV_INVALID_ALIGNMENT:
		return "0xC0350004"
	case STATUS_HV_INVALID_PARAMETER:
		return "0xC0350005"
	case STATUS_HV_ACCESS_DENIED:
		return "0xC0350006"
	case STATUS_HV_INVALID_PARTITION_STATE:
		return "0xC0350007"
	case STATUS_HV_OPERATION_DENIED:
		return "0xC0350008"
	case STATUS_HV_UNKNOWN_PROPERTY:
		return "0xC0350009"
	case STATUS_HV_PROPERTY_VALUE_OUT_OF_RANGE:
		return "0xC035000A"
	case STATUS_HV_INSUFFICIENT_MEMORY:
		return "0xC035000B"
	case STATUS_HV_PARTITION_TOO_DEEP:
		return "0xC035000C"
	case STATUS_HV_INVALID_PARTITION_ID:
		return "0xC035000D"
	case STATUS_HV_INVALID_VP_INDEX:
		return "0xC035000E"
	case STATUS_HV_INVALID_PORT_ID:
		return "0xC0350011"
	case STATUS_HV_INVALID_CONNECTION_ID:
		return "0xC0350012"
	case STATUS_HV_INSUFFICIENT_BUFFERS:
		return "0xC0350013"
	case STATUS_HV_NOT_ACKNOWLEDGED:
		return "0xC0350014"
	case STATUS_HV_ACKNOWLEDGED:
		return "0xC0350016"
	case STATUS_HV_INVALID_SAVE_RESTORE_STATE:
		return "0xC0350017"
	case STATUS_HV_INVALID_SYNIC_STATE:
		return "0xC0350018"
	case STATUS_HV_OBJECT_IN_USE:
		return "0xC0350019"
	case STATUS_HV_INVALID_PROXIMITY_DOMAIN_INFO:
		return "0xC035001A"
	case STATUS_HV_NO_DATA:
		return "0xC035001B"
	case STATUS_HV_INACTIVE:
		return "0xC035001C"
	case STATUS_HV_NO_RESOURCES:
		return "0xC035001D"
	case STATUS_HV_FEATURE_UNAVAILABLE:
		return "0xC035001E"
	case STATUS_HV_NOT_PRESENT:
		return "0xC0351000"
	case STATUS_VID_DUPLICATE_HANDLER:
		return "0xC0370001"
	case STATUS_VID_TOO_MANY_HANDLERS:
		return "0xC0370002"
	case STATUS_VID_QUEUE_FULL:
		return "0xC0370003"
	case STATUS_VID_HANDLER_NOT_PRESENT:
		return "0xC0370004"
	case STATUS_VID_INVALID_OBJECT_NAME:
		return "0xC0370005"
	case STATUS_VID_PARTITION_NAME_TOO_LONG:
		return "0xC0370006"
	case STATUS_VID_MESSAGE_QUEUE_NAME_TOO_LONG:
		return "0xC0370007"
	case STATUS_VID_PARTITION_ALREADY_EXISTS:
		return "0xC0370008"
	case STATUS_VID_PARTITION_DOES_NOT_EXIST:
		return "0xC0370009"
	case STATUS_VID_PARTITION_NAME_NOT_FOUND:
		return "0xC037000A"
	case STATUS_VID_MESSAGE_QUEUE_ALREADY_EXISTS:
		return "0xC037000B"
	case STATUS_VID_EXCEEDED_MBP_ENTRY_MAP_LIMIT:
		return "0xC037000C"
	case STATUS_VID_MB_STILL_REFERENCED:
		return "0xC037000D"
	case STATUS_VID_CHILD_GPA_PAGE_SET_CORRUPTED:
		return "0xC037000E"
	case STATUS_VID_INVALID_NUMA_SETTINGS:
		return "0xC037000F"
	case STATUS_VID_INVALID_NUMA_NODE_INDEX:
		return "0xC0370010"
	case STATUS_VID_NOTIFICATION_QUEUE_ALREADY_ASSOCIATED:
		return "0xC0370011"
	case STATUS_VID_INVALID_MEMORY_BLOCK_HANDLE:
		return "0xC0370012"
	case STATUS_VID_PAGE_RANGE_OVERFLOW:
		return "0xC0370013"
	case STATUS_VID_INVALID_MESSAGE_QUEUE_HANDLE:
		return "0xC0370014"
	case STATUS_VID_INVALID_GPA_RANGE_HANDLE:
		return "0xC0370015"
	case STATUS_VID_NO_MEMORY_BLOCK_NOTIFICATION_QUEUE:
		return "0xC0370016"
	case STATUS_VID_MEMORY_BLOCK_LOCK_COUNT_EXCEEDED:
		return "0xC0370017"
	case STATUS_VID_INVALID_PPM_HANDLE:
		return "0xC0370018"
	case STATUS_VID_MBPS_ARE_LOCKED:
		return "0xC0370019"
	case STATUS_VID_MESSAGE_QUEUE_CLOSED:
		return "0xC037001A"
	case STATUS_VID_VIRTUAL_PROCESSOR_LIMIT_EXCEEDED:
		return "0xC037001B"
	case STATUS_VID_STOP_PENDING:
		return "0xC037001C"
	case STATUS_VID_INVALID_PROCESSOR_STATE:
		return "0xC037001D"
	case STATUS_VID_EXCEEDED_KM_CONTEXT_COUNT_LIMIT:
		return "0xC037001E"
	case STATUS_VID_KM_INTERFACE_ALREADY_INITIALIZED:
		return "0xC037001F"
	case STATUS_VID_MB_PROPERTY_ALREADY_SET_RESET:
		return "0xC0370020"
	case STATUS_VID_MMIO_RANGE_DESTROYED:
		return "0xC0370021"
	case STATUS_VID_INVALID_CHILD_GPA_PAGE_SET:
		return "0xC0370022"
	case STATUS_VID_RESERVE_PAGE_SET_IS_BEING_USED:
		return "0xC0370023"
	case STATUS_VID_RESERVE_PAGE_SET_TOO_SMALL:
		return "0xC0370024"
	case STATUS_VID_MBP_ALREADY_LOCKED_USING_RESERVED_PAGE:
		return "0xC0370025"
	case STATUS_VID_MBP_COUNT_EXCEEDED_LIMIT:
		return "0xC0370026"
	case STATUS_VID_SAVED_STATE_CORRUPT:
		return "0xC0370027"
	case STATUS_VID_SAVED_STATE_UNRECOGNIZED_ITEM:
		return "0xC0370028"
	case STATUS_VID_SAVED_STATE_INCOMPATIBLE:
		return "0xC0370029"
	case STATUS_VID_REMOTE_NODE_PARENT_GPA_PAGES_USED:
		return "0x80370001"
	case STATUS_IPSEC_BAD_SPI:
		return "0xC0360001"
	case STATUS_IPSEC_SA_LIFETIME_EXPIRED:
		return "0xC0360002"
	case STATUS_IPSEC_WRONG_SA:
		return "0xC0360003"
	case STATUS_IPSEC_REPLAY_CHECK_FAILED:
		return "0xC0360004"
	case STATUS_IPSEC_INVALID_PACKET:
		return "0xC0360005"
	case STATUS_IPSEC_INTEGRITY_CHECK_FAILED:
		return "0xC0360006"
	case STATUS_IPSEC_CLEAR_TEXT_DROP:
		return "0xC0360007"
	case STATUS_IPSEC_AUTH_FIREWALL_DROP:
		return "0xC0360008"
	case STATUS_IPSEC_THROTTLE_DROP:
		return "0xC0360009"
	case STATUS_IPSEC_DOSP_BLOCK:
		return "0xC0368000"
	case STATUS_IPSEC_DOSP_RECEIVED_MULTICAST:
		return "0xC0368001"
	case STATUS_IPSEC_DOSP_INVALID_PACKET:
		return "0xC0368002"
	case STATUS_IPSEC_DOSP_STATE_LOOKUP_FAILED:
		return "0xC0368003"
	case STATUS_IPSEC_DOSP_MAX_ENTRIES:
		return "0xC0368004"
	case STATUS_IPSEC_DOSP_KEYMOD_NOT_ALLOWED:
		return "0xC0368005"
	case STATUS_IPSEC_DOSP_MAX_PER_IP_RATELIMIT_QUEUES:
		return "0xC0368006"
	case STATUS_VOLMGR_INCOMPLETE_REGENERATION:
		return "0x80380001"
	case STATUS_VOLMGR_INCOMPLETE_DISK_MIGRATION:
		return "0x80380002"
	case STATUS_VOLMGR_DATABASE_FULL:
		return "0xC0380001"
	case STATUS_VOLMGR_DISK_CONFIGURATION_CORRUPTED:
		return "0xC0380002"
	case STATUS_VOLMGR_DISK_CONFIGURATION_NOT_IN_SYNC:
		return "0xC0380003"
	case STATUS_VOLMGR_PACK_CONFIG_UPDATE_FAILED:
		return "0xC0380004"
	case STATUS_VOLMGR_DISK_CONTAINS_NON_SIMPLE_VOLUME:
		return "0xC0380005"
	case STATUS_VOLMGR_DISK_DUPLICATE:
		return "0xC0380006"
	case STATUS_VOLMGR_DISK_DYNAMIC:
		return "0xC0380007"
	case STATUS_VOLMGR_DISK_ID_INVALID:
		return "0xC0380008"
	case STATUS_VOLMGR_DISK_INVALID:
		return "0xC0380009"
	case STATUS_VOLMGR_DISK_LAST_VOTER:
		return "0xC038000A"
	case STATUS_VOLMGR_DISK_LAYOUT_INVALID:
		return "0xC038000B"
	case STATUS_VOLMGR_DISK_LAYOUT_NON_BASIC_BETWEEN_BASIC_PARTITIONS:
		return "0xC038000C"
	case STATUS_VOLMGR_DISK_LAYOUT_NOT_CYLINDER_ALIGNED:
		return "0xC038000D"
	case STATUS_VOLMGR_DISK_LAYOUT_PARTITIONS_TOO_SMALL:
		return "0xC038000E"
	case STATUS_VOLMGR_DISK_LAYOUT_PRIMARY_BETWEEN_LOGICAL_PARTITIONS:
		return "0xC038000F"
	case STATUS_VOLMGR_DISK_LAYOUT_TOO_MANY_PARTITIONS:
		return "0xC0380010"
	case STATUS_VOLMGR_DISK_MISSING:
		return "0xC0380011"
	case STATUS_VOLMGR_DISK_NOT_EMPTY:
		return "0xC0380012"
	case STATUS_VOLMGR_DISK_NOT_ENOUGH_SPACE:
		return "0xC0380013"
	case STATUS_VOLMGR_DISK_REVECTORING_FAILED:
		return "0xC0380014"
	case STATUS_VOLMGR_DISK_SECTOR_SIZE_INVALID:
		return "0xC0380015"
	case STATUS_VOLMGR_DISK_SET_NOT_CONTAINED:
		return "0xC0380016"
	case STATUS_VOLMGR_DISK_USED_BY_MULTIPLE_MEMBERS:
		return "0xC0380017"
	case STATUS_VOLMGR_DISK_USED_BY_MULTIPLE_PLEXES:
		return "0xC0380018"
	case STATUS_VOLMGR_DYNAMIC_DISK_NOT_SUPPORTED:
		return "0xC0380019"
	case STATUS_VOLMGR_EXTENT_ALREADY_USED:
		return "0xC038001A"
	case STATUS_VOLMGR_EXTENT_NOT_CONTIGUOUS:
		return "0xC038001B"
	case STATUS_VOLMGR_EXTENT_NOT_IN_PUBLIC_REGION:
		return "0xC038001C"
	case STATUS_VOLMGR_EXTENT_NOT_SECTOR_ALIGNED:
		return "0xC038001D"
	case STATUS_VOLMGR_EXTENT_OVERLAPS_EBR_PARTITION:
		return "0xC038001E"
	case STATUS_VOLMGR_EXTENT_VOLUME_LENGTHS_DO_NOT_MATCH:
		return "0xC038001F"
	case STATUS_VOLMGR_FAULT_TOLERANT_NOT_SUPPORTED:
		return "0xC0380020"
	case STATUS_VOLMGR_INTERLEAVE_LENGTH_INVALID:
		return "0xC0380021"
	case STATUS_VOLMGR_MAXIMUM_REGISTERED_USERS:
		return "0xC0380022"
	case STATUS_VOLMGR_MEMBER_IN_SYNC:
		return "0xC0380023"
	case STATUS_VOLMGR_MEMBER_INDEX_DUPLICATE:
		return "0xC0380024"
	case STATUS_VOLMGR_MEMBER_INDEX_INVALID:
		return "0xC0380025"
	case STATUS_VOLMGR_MEMBER_MISSING:
		return "0xC0380026"
	case STATUS_VOLMGR_MEMBER_NOT_DETACHED:
		return "0xC0380027"
	case STATUS_VOLMGR_MEMBER_REGENERATING:
		return "0xC0380028"
	case STATUS_VOLMGR_ALL_DISKS_FAILED:
		return "0xC0380029"
	case STATUS_VOLMGR_NO_REGISTERED_USERS:
		return "0xC038002A"
	case STATUS_VOLMGR_NO_SUCH_USER:
		return "0xC038002B"
	case STATUS_VOLMGR_NOTIFICATION_RESET:
		return "0xC038002C"
	case STATUS_VOLMGR_NUMBER_OF_MEMBERS_INVALID:
		return "0xC038002D"
	case STATUS_VOLMGR_NUMBER_OF_PLEXES_INVALID:
		return "0xC038002E"
	case STATUS_VOLMGR_PACK_DUPLICATE:
		return "0xC038002F"
	case STATUS_VOLMGR_PACK_ID_INVALID:
		return "0xC0380030"
	case STATUS_VOLMGR_PACK_INVALID:
		return "0xC0380031"
	case STATUS_VOLMGR_PACK_NAME_INVALID:
		return "0xC0380032"
	case STATUS_VOLMGR_PACK_OFFLINE:
		return "0xC0380033"
	case STATUS_VOLMGR_PACK_HAS_QUORUM:
		return "0xC0380034"
	case STATUS_VOLMGR_PACK_WITHOUT_QUORUM:
		return "0xC0380035"
	case STATUS_VOLMGR_PARTITION_STYLE_INVALID:
		return "0xC0380036"
	case STATUS_VOLMGR_PARTITION_UPDATE_FAILED:
		return "0xC0380037"
	case STATUS_VOLMGR_PLEX_IN_SYNC:
		return "0xC0380038"
	case STATUS_VOLMGR_PLEX_INDEX_DUPLICATE:
		return "0xC0380039"
	case STATUS_VOLMGR_PLEX_INDEX_INVALID:
		return "0xC038003A"
	case STATUS_VOLMGR_PLEX_LAST_ACTIVE:
		return "0xC038003B"
	case STATUS_VOLMGR_PLEX_MISSING:
		return "0xC038003C"
	case STATUS_VOLMGR_PLEX_REGENERATING:
		return "0xC038003D"
	case STATUS_VOLMGR_PLEX_TYPE_INVALID:
		return "0xC038003E"
	case STATUS_VOLMGR_PLEX_NOT_RAID5:
		return "0xC038003F"
	case STATUS_VOLMGR_PLEX_NOT_SIMPLE:
		return "0xC0380040"
	case STATUS_VOLMGR_STRUCTURE_SIZE_INVALID:
		return "0xC0380041"
	case STATUS_VOLMGR_TOO_MANY_NOTIFICATION_REQUESTS:
		return "0xC0380042"
	case STATUS_VOLMGR_TRANSACTION_IN_PROGRESS:
		return "0xC0380043"
	case STATUS_VOLMGR_UNEXPECTED_DISK_LAYOUT_CHANGE:
		return "0xC0380044"
	case STATUS_VOLMGR_VOLUME_CONTAINS_MISSING_DISK:
		return "0xC0380045"
	case STATUS_VOLMGR_VOLUME_ID_INVALID:
		return "0xC0380046"
	case STATUS_VOLMGR_VOLUME_LENGTH_INVALID:
		return "0xC0380047"
	case STATUS_VOLMGR_VOLUME_LENGTH_NOT_SECTOR_SIZE_MULTIPLE:
		return "0xC0380048"
	case STATUS_VOLMGR_VOLUME_NOT_MIRRORED:
		return "0xC0380049"
	case STATUS_VOLMGR_VOLUME_NOT_RETAINED:
		return "0xC038004A"
	case STATUS_VOLMGR_VOLUME_OFFLINE:
		return "0xC038004B"
	case STATUS_VOLMGR_VOLUME_RETAINED:
		return "0xC038004C"
	case STATUS_VOLMGR_NUMBER_OF_EXTENTS_INVALID:
		return "0xC038004D"
	case STATUS_VOLMGR_DIFFERENT_SECTOR_SIZE:
		return "0xC038004E"
	case STATUS_VOLMGR_BAD_BOOT_DISK:
		return "0xC038004F"
	case STATUS_VOLMGR_PACK_CONFIG_OFFLINE:
		return "0xC0380050"
	case STATUS_VOLMGR_PACK_CONFIG_ONLINE:
		return "0xC0380051"
	case STATUS_VOLMGR_NOT_PRIMARY_PACK:
		return "0xC0380052"
	case STATUS_VOLMGR_PACK_LOG_UPDATE_FAILED:
		return "0xC0380053"
	case STATUS_VOLMGR_NUMBER_OF_DISKS_IN_PLEX_INVALID:
		return "0xC0380054"
	case STATUS_VOLMGR_NUMBER_OF_DISKS_IN_MEMBER_INVALID:
		return "0xC0380055"
	case STATUS_VOLMGR_VOLUME_MIRRORED:
		return "0xC0380056"
	case STATUS_VOLMGR_PLEX_NOT_SIMPLE_SPANNED:
		return "0xC0380057"
	case STATUS_VOLMGR_NO_VALID_LOG_COPIES:
		return "0xC0380058"
	case STATUS_VOLMGR_PRIMARY_PACK_PRESENT:
		return "0xC0380059"
	case STATUS_VOLMGR_NUMBER_OF_DISKS_INVALID:
		return "0xC038005A"
	case STATUS_VOLMGR_MIRROR_NOT_SUPPORTED:
		return "0xC038005B"
	case STATUS_VOLMGR_RAID5_NOT_SUPPORTED:
		return "0xC038005C"
	case STATUS_BCD_NOT_ALL_ENTRIES_IMPORTED:
		return "0x80390001"
	case STATUS_BCD_TOO_MANY_ELEMENTS:
		return "0xC0390002"
	case STATUS_BCD_NOT_ALL_ENTRIES_SYNCHRONIZED:
		return "0x80390003"
	case STATUS_VHD_DRIVE_FOOTER_MISSING:
		return "0xC03A0001"
	case STATUS_VHD_DRIVE_FOOTER_CHECKSUM_MISMATCH:
		return "0xC03A0002"
	case STATUS_VHD_DRIVE_FOOTER_CORRUPT:
		return "0xC03A0003"
	case STATUS_VHD_FORMAT_UNKNOWN:
		return "0xC03A0004"
	case STATUS_VHD_FORMAT_UNSUPPORTED_VERSION:
		return "0xC03A0005"
	case STATUS_VHD_SPARSE_HEADER_CHECKSUM_MISMATCH:
		return "0xC03A0006"
	case STATUS_VHD_SPARSE_HEADER_UNSUPPORTED_VERSION:
		return "0xC03A0007"
	case STATUS_VHD_SPARSE_HEADER_CORRUPT:
		return "0xC03A0008"
	case STATUS_VHD_BLOCK_ALLOCATION_FAILURE:
		return "0xC03A0009"
	case STATUS_VHD_BLOCK_ALLOCATION_TABLE_CORRUPT:
		return "0xC03A000A"
	case STATUS_VHD_INVALID_BLOCK_SIZE:
		return "0xC03A000B"
	case STATUS_VHD_BITMAP_MISMATCH:
		return "0xC03A000C"
	case STATUS_VHD_PARENT_VHD_NOT_FOUND:
		return "0xC03A000D"
	case STATUS_VHD_CHILD_PARENT_ID_MISMATCH:
		return "0xC03A000E"
	case STATUS_VHD_CHILD_PARENT_TIMESTAMP_MISMATCH:
		return "0xC03A000F"
	case STATUS_VHD_METADATA_READ_FAILURE:
		return "0xC03A0010"
	case STATUS_VHD_METADATA_WRITE_FAILURE:
		return "0xC03A0011"
	case STATUS_VHD_INVALID_SIZE:
		return "0xC03A0012"
	case STATUS_VHD_INVALID_FILE_SIZE:
		return "0xC03A0013"
	case STATUS_VIRTDISK_PROVIDER_NOT_FOUND:
		return "0xC03A0014"
	case STATUS_VIRTDISK_NOT_VIRTUAL_DISK:
		return "0xC03A0015"
	case STATUS_VHD_PARENT_VHD_ACCESS_DENIED:
		return "0xC03A0016"
	case STATUS_VHD_CHILD_PARENT_SIZE_MISMATCH:
		return "0xC03A0017"
	case STATUS_VHD_DIFFERENCING_CHAIN_CYCLE_DETECTED:
		return "0xC03A0018"
	case STATUS_VHD_DIFFERENCING_CHAIN_ERROR_IN_PARENT:
		return "0xC03A0019"
	case STATUS_VIRTUAL_DISK_LIMITATION:
		return "0xC03A001A"
	case STATUS_VHD_INVALID_TYPE:
		return "0xC03A001B"
	case STATUS_VHD_INVALID_STATE:
		return "0xC03A001C"
	case STATUS_VIRTDISK_UNSUPPORTED_DISK_SECTOR_SIZE:
		return "0xC03A001D"
	case STATUS_QUERY_STORAGE_ERROR:
		return "0x803A0001"
	case STATUS_DIS_NOT_PRESENT:
		return "0xC03C0001"
	case STATUS_DIS_ATTRIBUTE_NOT_FOUND:
		return "0xC03C0002"
	case STATUS_DIS_UNRECOGNIZED_ATTRIBUTE:
		return "0xC03C0003"
	case STATUS_DIS_PARTIAL_DATA:
		return "0xC03C0004"
	default:
		return "unknown"
	}
}

func (k NtstatusKind) Elements() []NtstatusKind {
	return []NtstatusKind{
		STATUS_SUCCESS,
		STATUS_WAIT_1,
		STATUS_WAIT_2,
		STATUS_WAIT_3,
		STATUS_WAIT_63,
		STATUS_ABANDONED,
		STATUS_ABANDONED_WAIT_63,
		STATUS_USER_APC,
		STATUS_KERNEL_APC,
		STATUS_ALERTED,
		STATUS_TIMEOUT,
		STATUS_PENDING,
		STATUS_REPARSE,
		STATUS_MORE_ENTRIES,
		STATUS_NOT_ALL_ASSIGNED,
		STATUS_SOME_NOT_MAPPED,
		STATUS_OPLOCK_BREAK_IN_PROGRESS,
		STATUS_VOLUME_MOUNTED,
		STATUS_RXACT_COMMITTED,
		STATUS_NOTIFY_CLEANUP,
		STATUS_NOTIFY_ENUM_DIR,
		STATUS_NO_QUOTAS_FOR_ACCOUNT,
		STATUS_PRIMARY_TRANSPORT_CONNECT_FAILED,
		STATUS_PAGE_FAULT_TRANSITION,
		STATUS_PAGE_FAULT_DEMAND_ZERO,
		STATUS_PAGE_FAULT_COPY_ON_WRITE,
		STATUS_PAGE_FAULT_GUARD_PAGE,
		STATUS_PAGE_FAULT_PAGING_FILE,
		STATUS_CACHE_PAGE_LOCKED,
		STATUS_CRASH_DUMP,
		STATUS_BUFFER_ALL_ZEROS,
		STATUS_REPARSE_OBJECT,
		STATUS_RESOURCE_REQUIREMENTS_CHANGED,
		STATUS_TRANSLATION_COMPLETE,
		STATUS_DS_MEMBERSHIP_EVALUATED_LOCALLY,
		STATUS_NOTHING_TO_TERMINATE,
		STATUS_PROCESS_NOT_IN_JOB,
		STATUS_PROCESS_IN_JOB,
		STATUS_VOLSNAP_HIBERNATE_READY,
		STATUS_FSFILTER_OP_COMPLETED_SUCCESSFULLY,
		STATUS_INTERRUPT_VECTOR_ALREADY_CONNECTED,
		STATUS_INTERRUPT_STILL_CONNECTED,
		STATUS_PROCESS_CLONED,
		STATUS_FILE_LOCKED_WITH_ONLY_READERS,
		STATUS_FILE_LOCKED_WITH_WRITERS,
		STATUS_RESOURCEMANAGER_READ_ONLY,
		STATUS_RING_PREVIOUSLY_EMPTY,
		STATUS_RING_PREVIOUSLY_FULL,
		STATUS_RING_PREVIOUSLY_ABOVE_QUOTA,
		STATUS_RING_NEWLY_EMPTY,
		STATUS_RING_SIGNAL_OPPOSITE_ENDPOINT,
		STATUS_OPLOCK_SWITCHED_TO_NEW_HANDLE,
		STATUS_OPLOCK_HANDLE_CLOSED,
		STATUS_WAIT_FOR_OPLOCK,
		DBG_EXCEPTION_HANDLED,
		DBG_CONTINUE,
		STATUS_FLT_IO_COMPLETE,
		STATUS_DIS_ATTRIBUTE_BUILT,
		STATUS_OBJECT_NAME_EXISTS,
		STATUS_THREAD_WAS_SUSPENDED,
		STATUS_WORKING_SET_LIMIT_RANGE,
		STATUS_IMAGE_NOT_AT_BASE,
		STATUS_RXACT_STATE_CREATED,
		STATUS_SEGMENT_NOTIFICATION,
		STATUS_LOCAL_USER_SESSION_KEY,
		STATUS_BAD_CURRENT_DIRECTORY,
		STATUS_SERIAL_MORE_WRITES,
		STATUS_REGISTRY_RECOVERED,
		STATUS_FT_READ_RECOVERY_FROM_BACKUP,
		STATUS_FT_WRITE_RECOVERY,
		STATUS_SERIAL_COUNTER_TIMEOUT,
		STATUS_NULL_LM_PASSWORD,
		STATUS_IMAGE_MACHINE_TYPE_MISMATCH,
		STATUS_RECEIVE_PARTIAL,
		STATUS_RECEIVE_EXPEDITED,
		STATUS_RECEIVE_PARTIAL_EXPEDITED,
		STATUS_EVENT_DONE,
		STATUS_EVENT_PENDING,
		STATUS_CHECKING_FILE_SYSTEM,
		STATUS_FATAL_APP_EXIT,
		STATUS_PREDEFINED_HANDLE,
		STATUS_WAS_UNLOCKED,
		STATUS_SERVICE_NOTIFICATION,
		STATUS_WAS_LOCKED,
		STATUS_LOG_HARD_ERROR,
		STATUS_ALREADY_WIN32,
		STATUS_WX86_UNSIMULATE,
		STATUS_WX86_CONTINUE,
		STATUS_WX86_SINGLE_STEP,
		STATUS_WX86_BREAKPOINT,
		STATUS_WX86_EXCEPTION_CONTINUE,
		STATUS_WX86_EXCEPTION_LASTCHANCE,
		STATUS_WX86_EXCEPTION_CHAIN,
		STATUS_IMAGE_MACHINE_TYPE_MISMATCH_EXE,
		STATUS_NO_YIELD_PERFORMED,
		STATUS_TIMER_RESUME_IGNORED,
		STATUS_ARBITRATION_UNHANDLED,
		STATUS_CARDBUS_NOT_SUPPORTED,
		STATUS_WX86_CREATEWX86TIB,
		STATUS_MP_PROCESSOR_MISMATCH,
		STATUS_HIBERNATED,
		STATUS_RESUME_HIBERNATION,
		STATUS_FIRMWARE_UPDATED,
		STATUS_DRIVERS_LEAKING_LOCKED_PAGES,
		STATUS_MESSAGE_RETRIEVED,
		STATUS_SYSTEM_POWERSTATE_TRANSITION,
		STATUS_ALPC_CHECK_COMPLETION_LIST,
		STATUS_SYSTEM_POWERSTATE_COMPLEX_TRANSITION,
		STATUS_ACCESS_AUDIT_BY_POLICY,
		STATUS_ABANDON_HIBERFILE,
		STATUS_BIZRULES_NOT_ENABLED,
		DBG_REPLY_LATER,
		DBG_UNABLE_TO_PROVIDE_HANDLE,
		DBG_TERMINATE_THREAD,
		DBG_TERMINATE_PROCESS,
		DBG_CONTROL_C,
		DBG_PRINTEXCEPTION_C,
		DBG_RIPEXCEPTION,
		DBG_CONTROL_BREAK,
		DBG_COMMAND_EXCEPTION,
		STATUS_HEURISTIC_DAMAGE_POSSIBLE,
		STATUS_GUARD_PAGE_VIOLATION,
		STATUS_DATATYPE_MISALIGNMENT,
		STATUS_BREAKPOINT,
		STATUS_SINGLE_STEP,
		STATUS_BUFFER_OVERFLOW,
		STATUS_NO_MORE_FILES,
		STATUS_WAKE_SYSTEM_DEBUGGER,
		STATUS_HANDLES_CLOSED,
		STATUS_NO_INHERITANCE,
		STATUS_GUID_SUBSTITUTION_MADE,
		STATUS_PARTIAL_COPY,
		STATUS_DEVICE_PAPER_EMPTY,
		STATUS_DEVICE_POWERED_OFF,
		STATUS_DEVICE_OFF_LINE,
		STATUS_DEVICE_BUSY,
		STATUS_NO_MORE_EAS,
		STATUS_INVALID_EA_NAME,
		STATUS_EA_LIST_INCONSISTENT,
		STATUS_INVALID_EA_FLAG,
		STATUS_VERIFY_REQUIRED,
		STATUS_EXTRANEOUS_INFORMATION,
		STATUS_RXACT_COMMIT_NECESSARY,
		STATUS_NO_MORE_ENTRIES,
		STATUS_FILEMARK_DETECTED,
		STATUS_MEDIA_CHANGED,
		STATUS_BUS_RESET,
		STATUS_END_OF_MEDIA,
		STATUS_BEGINNING_OF_MEDIA,
		STATUS_MEDIA_CHECK,
		STATUS_SETMARK_DETECTED,
		STATUS_NO_DATA_DETECTED,
		STATUS_REDIRECTOR_HAS_OPEN_HANDLES,
		STATUS_SERVER_HAS_OPEN_HANDLES,
		STATUS_ALREADY_DISCONNECTED,
		STATUS_LONGJUMP,
		STATUS_CLEANER_CARTRIDGE_INSTALLED,
		STATUS_PLUGPLAY_QUERY_VETOED,
		STATUS_UNWIND_CONSOLIDATE,
		STATUS_REGISTRY_HIVE_RECOVERED,
		STATUS_DLL_MIGHT_BE_INSECURE,
		STATUS_DLL_MIGHT_BE_INCOMPATIBLE,
		STATUS_STOPPED_ON_SYMLINK,
		STATUS_CANNOT_GRANT_REQUESTED_OPLOCK,
		STATUS_NO_ACE_CONDITION,
		DBG_EXCEPTION_NOT_HANDLED,
		STATUS_CLUSTER_NODE_ALREADY_UP,
		STATUS_CLUSTER_NODE_ALREADY_DOWN,
		STATUS_CLUSTER_NETWORK_ALREADY_ONLINE,
		STATUS_CLUSTER_NETWORK_ALREADY_OFFLINE,
		STATUS_CLUSTER_NODE_ALREADY_MEMBER,
		STATUS_FLT_BUFFER_TOO_SMALL,
		STATUS_FVE_PARTIAL_METADATA,
		STATUS_FVE_TRANSIENT_STATE,
		STATUS_UNSUCCESSFUL,
		STATUS_NOT_IMPLEMENTED,
		STATUS_INVALID_INFO_CLASS,
		STATUS_INFO_LENGTH_MISMATCH,
		STATUS_ACCESS_VIOLATION,
		STATUS_IN_PAGE_ERROR,
		STATUS_PAGEFILE_QUOTA,
		STATUS_INVALID_HANDLE,
		STATUS_BAD_INITIAL_STACK,
		STATUS_BAD_INITIAL_PC,
		STATUS_INVALID_CID,
		STATUS_TIMER_NOT_CANCELED,
		STATUS_INVALID_PARAMETER,
		STATUS_NO_SUCH_DEVICE,
		STATUS_NO_SUCH_FILE,
		STATUS_INVALID_DEVICE_REQUEST,
		STATUS_END_OF_FILE,
		STATUS_WRONG_VOLUME,
		STATUS_NO_MEDIA_IN_DEVICE,
		STATUS_UNRECOGNIZED_MEDIA,
		STATUS_NONEXISTENT_SECTOR,
		STATUS_MORE_PROCESSING_REQUIRED,
		STATUS_NO_MEMORY,
		STATUS_CONFLICTING_ADDRESSES,
		STATUS_NOT_MAPPED_VIEW,
		STATUS_UNABLE_TO_FREE_VM,
		STATUS_UNABLE_TO_DELETE_SECTION,
		STATUS_INVALID_SYSTEM_SERVICE,
		STATUS_ILLEGAL_INSTRUCTION,
		STATUS_INVALID_LOCK_SEQUENCE,
		STATUS_INVALID_VIEW_SIZE,
		STATUS_INVALID_FILE_FOR_SECTION,
		STATUS_ALREADY_COMMITTED,
		STATUS_ACCESS_DENIED,
		STATUS_BUFFER_TOO_SMALL,
		STATUS_OBJECT_TYPE_MISMATCH,
		STATUS_NONCONTINUABLE_EXCEPTION,
		STATUS_INVALID_DISPOSITION,
		STATUS_UNWIND,
		STATUS_BAD_STACK,
		STATUS_INVALID_UNWIND_TARGET,
		STATUS_NOT_LOCKED,
		STATUS_PARITY_ERROR,
		STATUS_UNABLE_TO_DECOMMIT_VM,
		STATUS_NOT_COMMITTED,
		STATUS_INVALID_PORT_ATTRIBUTES,
		STATUS_PORT_MESSAGE_TOO_LONG,
		STATUS_INVALID_PARAMETER_MIX,
		STATUS_INVALID_QUOTA_LOWER,
		STATUS_DISK_CORRUPT_ERROR,
		STATUS_OBJECT_NAME_INVALID,
		STATUS_OBJECT_NAME_NOT_FOUND,
		STATUS_OBJECT_NAME_COLLISION,
		STATUS_PORT_DISCONNECTED,
		STATUS_DEVICE_ALREADY_ATTACHED,
		STATUS_OBJECT_PATH_INVALID,
		STATUS_OBJECT_PATH_NOT_FOUND,
		STATUS_OBJECT_PATH_SYNTAX_BAD,
		STATUS_DATA_OVERRUN,
		STATUS_DATA_LATE_ERROR,
		STATUS_DATA_ERROR,
		STATUS_CRC_ERROR,
		STATUS_SECTION_TOO_BIG,
		STATUS_PORT_CONNECTION_REFUSED,
		STATUS_INVALID_PORT_HANDLE,
		STATUS_SHARING_VIOLATION,
		STATUS_QUOTA_EXCEEDED,
		STATUS_INVALID_PAGE_PROTECTION,
		STATUS_MUTANT_NOT_OWNED,
		STATUS_SEMAPHORE_LIMIT_EXCEEDED,
		STATUS_PORT_ALREADY_SET,
		STATUS_SECTION_NOT_IMAGE,
		STATUS_SUSPEND_COUNT_EXCEEDED,
		STATUS_THREAD_IS_TERMINATING,
		STATUS_BAD_WORKING_SET_LIMIT,
		STATUS_INCOMPATIBLE_FILE_MAP,
		STATUS_SECTION_PROTECTION,
		STATUS_EAS_NOT_SUPPORTED,
		STATUS_EA_TOO_LARGE,
		STATUS_NONEXISTENT_EA_ENTRY,
		STATUS_NO_EAS_ON_FILE,
		STATUS_EA_CORRUPT_ERROR,
		STATUS_FILE_LOCK_CONFLICT,
		STATUS_LOCK_NOT_GRANTED,
		STATUS_DELETE_PENDING,
		STATUS_CTL_FILE_NOT_SUPPORTED,
		STATUS_UNKNOWN_REVISION,
		STATUS_REVISION_MISMATCH,
		STATUS_INVALID_OWNER,
		STATUS_INVALID_PRIMARY_GROUP,
		STATUS_NO_IMPERSONATION_TOKEN,
		STATUS_CANT_DISABLE_MANDATORY,
		STATUS_NO_LOGON_SERVERS,
		STATUS_NO_SUCH_LOGON_SESSION,
		STATUS_NO_SUCH_PRIVILEGE,
		STATUS_PRIVILEGE_NOT_HELD,
		STATUS_INVALID_ACCOUNT_NAME,
		STATUS_USER_EXISTS,
		STATUS_NO_SUCH_USER,
		STATUS_GROUP_EXISTS,
		STATUS_NO_SUCH_GROUP,
		STATUS_MEMBER_IN_GROUP,
		STATUS_MEMBER_NOT_IN_GROUP,
		STATUS_LAST_ADMIN,
		STATUS_WRONG_PASSWORD,
		STATUS_ILL_FORMED_PASSWORD,
		STATUS_PASSWORD_RESTRICTION,
		STATUS_LOGON_FAILURE,
		STATUS_ACCOUNT_RESTRICTION,
		STATUS_INVALID_LOGON_HOURS,
		STATUS_INVALID_WORKSTATION,
		STATUS_PASSWORD_EXPIRED,
		STATUS_ACCOUNT_DISABLED,
		STATUS_NONE_MAPPED,
		STATUS_TOO_MANY_LUIDS_REQUESTED,
		STATUS_LUIDS_EXHAUSTED,
		STATUS_INVALID_SUB_AUTHORITY,
		STATUS_INVALID_ACL,
		STATUS_INVALID_SID,
		STATUS_INVALID_SECURITY_DESCR,
		STATUS_PROCEDURE_NOT_FOUND,
		STATUS_INVALID_IMAGE_FORMAT,
		STATUS_NO_TOKEN,
		STATUS_BAD_INHERITANCE_ACL,
		STATUS_RANGE_NOT_LOCKED,
		STATUS_DISK_FULL,
		STATUS_SERVER_DISABLED,
		STATUS_SERVER_NOT_DISABLED,
		STATUS_TOO_MANY_GUIDS_REQUESTED,
		STATUS_GUIDS_EXHAUSTED,
		STATUS_INVALID_ID_AUTHORITY,
		STATUS_AGENTS_EXHAUSTED,
		STATUS_INVALID_VOLUME_LABEL,
		STATUS_SECTION_NOT_EXTENDED,
		STATUS_NOT_MAPPED_DATA,
		STATUS_RESOURCE_DATA_NOT_FOUND,
		STATUS_RESOURCE_TYPE_NOT_FOUND,
		STATUS_RESOURCE_NAME_NOT_FOUND,
		STATUS_ARRAY_BOUNDS_EXCEEDED,
		STATUS_FLOAT_DENORMAL_OPERAND,
		STATUS_FLOAT_DIVIDE_BY_ZERO,
		STATUS_FLOAT_INEXACT_RESULT,
		STATUS_FLOAT_INVALID_OPERATION,
		STATUS_FLOAT_OVERFLOW,
		STATUS_FLOAT_STACK_CHECK,
		STATUS_FLOAT_UNDERFLOW,
		STATUS_INTEGER_DIVIDE_BY_ZERO,
		STATUS_INTEGER_OVERFLOW,
		STATUS_PRIVILEGED_INSTRUCTION,
		STATUS_TOO_MANY_PAGING_FILES,
		STATUS_FILE_INVALID,
		STATUS_ALLOTTED_SPACE_EXCEEDED,
		STATUS_INSUFFICIENT_RESOURCES,
		STATUS_DFS_EXIT_PATH_FOUND,
		STATUS_DEVICE_DATA_ERROR,
		STATUS_DEVICE_NOT_CONNECTED,
		STATUS_DEVICE_POWER_FAILURE,
		STATUS_FREE_VM_NOT_AT_BASE,
		STATUS_MEMORY_NOT_ALLOCATED,
		STATUS_WORKING_SET_QUOTA,
		STATUS_MEDIA_WRITE_PROTECTED,
		STATUS_DEVICE_NOT_READY,
		STATUS_INVALID_GROUP_ATTRIBUTES,
		STATUS_BAD_IMPERSONATION_LEVEL,
		STATUS_CANT_OPEN_ANONYMOUS,
		STATUS_BAD_VALIDATION_CLASS,
		STATUS_BAD_TOKEN_TYPE,
		STATUS_BAD_MASTER_BOOT_RECORD,
		STATUS_INSTRUCTION_MISALIGNMENT,
		STATUS_INSTANCE_NOT_AVAILABLE,
		STATUS_PIPE_NOT_AVAILABLE,
		STATUS_INVALID_PIPE_STATE,
		STATUS_PIPE_BUSY,
		STATUS_ILLEGAL_FUNCTION,
		STATUS_PIPE_DISCONNECTED,
		STATUS_PIPE_CLOSING,
		STATUS_PIPE_CONNECTED,
		STATUS_PIPE_LISTENING,
		STATUS_INVALID_READ_MODE,
		STATUS_IO_TIMEOUT,
		STATUS_FILE_FORCED_CLOSED,
		STATUS_PROFILING_NOT_STARTED,
		STATUS_PROFILING_NOT_STOPPED,
		STATUS_COULD_NOT_INTERPRET,
		STATUS_FILE_IS_A_DIRECTORY,
		STATUS_NOT_SUPPORTED,
		STATUS_REMOTE_NOT_LISTENING,
		STATUS_DUPLICATE_NAME,
		STATUS_BAD_NETWORK_PATH,
		STATUS_NETWORK_BUSY,
		STATUS_DEVICE_DOES_NOT_EXIST,
		STATUS_TOO_MANY_COMMANDS,
		STATUS_ADAPTER_HARDWARE_ERROR,
		STATUS_INVALID_NETWORK_RESPONSE,
		STATUS_UNEXPECTED_NETWORK_ERROR,
		STATUS_BAD_REMOTE_ADAPTER,
		STATUS_PRINT_QUEUE_FULL,
		STATUS_NO_SPOOL_SPACE,
		STATUS_PRINT_CANCELLED,
		STATUS_NETWORK_NAME_DELETED,
		STATUS_NETWORK_ACCESS_DENIED,
		STATUS_BAD_DEVICE_TYPE,
		STATUS_BAD_NETWORK_NAME,
		STATUS_TOO_MANY_NAMES,
		STATUS_TOO_MANY_SESSIONS,
		STATUS_SHARING_PAUSED,
		STATUS_REQUEST_NOT_ACCEPTED,
		STATUS_REDIRECTOR_PAUSED,
		STATUS_NET_WRITE_FAULT,
		STATUS_PROFILING_AT_LIMIT,
		STATUS_NOT_SAME_DEVICE,
		STATUS_FILE_RENAMED,
		STATUS_VIRTUAL_CIRCUIT_CLOSED,
		STATUS_NO_SECURITY_ON_OBJECT,
		STATUS_CANT_WAIT,
		STATUS_PIPE_EMPTY,
		STATUS_CANT_ACCESS_DOMAIN_INFO,
		STATUS_CANT_TERMINATE_SELF,
		STATUS_INVALID_SERVER_STATE,
		STATUS_INVALID_DOMAIN_STATE,
		STATUS_INVALID_DOMAIN_ROLE,
		STATUS_NO_SUCH_DOMAIN,
		STATUS_DOMAIN_EXISTS,
		STATUS_DOMAIN_LIMIT_EXCEEDED,
		STATUS_OPLOCK_NOT_GRANTED,
		STATUS_INVALID_OPLOCK_PROTOCOL,
		STATUS_INTERNAL_DB_CORRUPTION,
		STATUS_INTERNAL_ERROR,
		STATUS_GENERIC_NOT_MAPPED,
		STATUS_BAD_DESCRIPTOR_FORMAT,
		STATUS_INVALID_USER_BUFFER,
		STATUS_UNEXPECTED_IO_ERROR,
		STATUS_UNEXPECTED_MM_CREATE_ERR,
		STATUS_UNEXPECTED_MM_MAP_ERROR,
		STATUS_UNEXPECTED_MM_EXTEND_ERR,
		STATUS_NOT_LOGON_PROCESS,
		STATUS_LOGON_SESSION_EXISTS,
		STATUS_INVALID_PARAMETER_1,
		STATUS_INVALID_PARAMETER_2,
		STATUS_INVALID_PARAMETER_3,
		STATUS_INVALID_PARAMETER_4,
		STATUS_INVALID_PARAMETER_5,
		STATUS_INVALID_PARAMETER_6,
		STATUS_INVALID_PARAMETER_7,
		STATUS_INVALID_PARAMETER_8,
		STATUS_INVALID_PARAMETER_9,
		STATUS_INVALID_PARAMETER_10,
		STATUS_INVALID_PARAMETER_11,
		STATUS_INVALID_PARAMETER_12,
		STATUS_REDIRECTOR_NOT_STARTED,
		STATUS_REDIRECTOR_STARTED,
		STATUS_STACK_OVERFLOW,
		STATUS_NO_SUCH_PACKAGE,
		STATUS_BAD_FUNCTION_TABLE,
		STATUS_VARIABLE_NOT_FOUND,
		STATUS_DIRECTORY_NOT_EMPTY,
		STATUS_FILE_CORRUPT_ERROR,
		STATUS_NOT_A_DIRECTORY,
		STATUS_BAD_LOGON_SESSION_STATE,
		STATUS_LOGON_SESSION_COLLISION,
		STATUS_NAME_TOO_LONG,
		STATUS_FILES_OPEN,
		STATUS_CONNECTION_IN_USE,
		STATUS_MESSAGE_NOT_FOUND,
		STATUS_PROCESS_IS_TERMINATING,
		STATUS_INVALID_LOGON_TYPE,
		STATUS_NO_GUID_TRANSLATION,
		STATUS_CANNOT_IMPERSONATE,
		STATUS_IMAGE_ALREADY_LOADED,
		STATUS_ABIOS_NOT_PRESENT,
		STATUS_ABIOS_LID_NOT_EXIST,
		STATUS_ABIOS_LID_ALREADY_OWNED,
		STATUS_ABIOS_NOT_LID_OWNER,
		STATUS_ABIOS_INVALID_COMMAND,
		STATUS_ABIOS_INVALID_LID,
		STATUS_ABIOS_SELECTOR_NOT_AVAILABLE,
		STATUS_ABIOS_INVALID_SELECTOR,
		STATUS_NO_LDT,
		STATUS_INVALID_LDT_SIZE,
		STATUS_INVALID_LDT_OFFSET,
		STATUS_INVALID_LDT_DESCRIPTOR,
		STATUS_INVALID_IMAGE_NE_FORMAT,
		STATUS_RXACT_INVALID_STATE,
		STATUS_RXACT_COMMIT_FAILURE,
		STATUS_MAPPED_FILE_SIZE_ZERO,
		STATUS_TOO_MANY_OPENED_FILES,
		STATUS_CANCELLED,
		STATUS_CANNOT_DELETE,
		STATUS_INVALID_COMPUTER_NAME,
		STATUS_FILE_DELETED,
		STATUS_SPECIAL_ACCOUNT,
		STATUS_SPECIAL_GROUP,
		STATUS_SPECIAL_USER,
		STATUS_MEMBERS_PRIMARY_GROUP,
		STATUS_FILE_CLOSED,
		STATUS_TOO_MANY_THREADS,
		STATUS_THREAD_NOT_IN_PROCESS,
		STATUS_TOKEN_ALREADY_IN_USE,
		STATUS_PAGEFILE_QUOTA_EXCEEDED,
		STATUS_COMMITMENT_LIMIT,
		STATUS_INVALID_IMAGE_LE_FORMAT,
		STATUS_INVALID_IMAGE_NOT_MZ,
		STATUS_INVALID_IMAGE_PROTECT,
		STATUS_INVALID_IMAGE_WIN_16,
		STATUS_LOGON_SERVER_CONFLICT,
		STATUS_TIME_DIFFERENCE_AT_DC,
		STATUS_SYNCHRONIZATION_REQUIRED,
		STATUS_DLL_NOT_FOUND,
		STATUS_OPEN_FAILED,
		STATUS_IO_PRIVILEGE_FAILED,
		STATUS_ORDINAL_NOT_FOUND,
		STATUS_ENTRYPOINT_NOT_FOUND,
		STATUS_CONTROL_C_EXIT,
		STATUS_LOCAL_DISCONNECT,
		STATUS_REMOTE_DISCONNECT,
		STATUS_REMOTE_RESOURCES,
		STATUS_LINK_FAILED,
		STATUS_LINK_TIMEOUT,
		STATUS_INVALID_CONNECTION,
		STATUS_INVALID_ADDRESS,
		STATUS_DLL_INIT_FAILED,
		STATUS_MISSING_SYSTEMFILE,
		STATUS_UNHANDLED_EXCEPTION,
		STATUS_APP_INIT_FAILURE,
		STATUS_PAGEFILE_CREATE_FAILED,
		STATUS_NO_PAGEFILE,
		STATUS_INVALID_LEVEL,
		STATUS_WRONG_PASSWORD_CORE,
		STATUS_ILLEGAL_FLOAT_CONTEXT,
		STATUS_PIPE_BROKEN,
		STATUS_REGISTRY_CORRUPT,
		STATUS_REGISTRY_IO_FAILED,
		STATUS_NO_EVENT_PAIR,
		STATUS_UNRECOGNIZED_VOLUME,
		STATUS_SERIAL_NO_DEVICE_INITED,
		STATUS_NO_SUCH_ALIAS,
		STATUS_MEMBER_NOT_IN_ALIAS,
		STATUS_MEMBER_IN_ALIAS,
		STATUS_ALIAS_EXISTS,
		STATUS_LOGON_NOT_GRANTED,
		STATUS_TOO_MANY_SECRETS,
		STATUS_SECRET_TOO_LONG,
		STATUS_INTERNAL_DB_ERROR,
		STATUS_FULLSCREEN_MODE,
		STATUS_TOO_MANY_CONTEXT_IDS,
		STATUS_LOGON_TYPE_NOT_GRANTED,
		STATUS_NOT_REGISTRY_FILE,
		STATUS_NT_CROSS_ENCRYPTION_REQUIRED,
		STATUS_DOMAIN_CTRLR_CONFIG_ERROR,
		STATUS_FT_MISSING_MEMBER,
		STATUS_ILL_FORMED_SERVICE_ENTRY,
		STATUS_ILLEGAL_CHARACTER,
		STATUS_UNMAPPABLE_CHARACTER,
		STATUS_UNDEFINED_CHARACTER,
		STATUS_FLOPPY_VOLUME,
		STATUS_FLOPPY_ID_MARK_NOT_FOUND,
		STATUS_FLOPPY_WRONG_CYLINDER,
		STATUS_FLOPPY_UNKNOWN_ERROR,
		STATUS_FLOPPY_BAD_REGISTERS,
		STATUS_DISK_RECALIBRATE_FAILED,
		STATUS_DISK_OPERATION_FAILED,
		STATUS_DISK_RESET_FAILED,
		STATUS_SHARED_IRQ_BUSY,
		STATUS_FT_ORPHANING,
		STATUS_BIOS_FAILED_TO_CONNECT_INTERRUPT,
		STATUS_PARTITION_FAILURE,
		STATUS_INVALID_BLOCK_LENGTH,
		STATUS_DEVICE_NOT_PARTITIONED,
		STATUS_UNABLE_TO_LOCK_MEDIA,
		STATUS_UNABLE_TO_UNLOAD_MEDIA,
		STATUS_EOM_OVERFLOW,
		STATUS_NO_MEDIA,
		STATUS_NO_SUCH_MEMBER,
		STATUS_INVALID_MEMBER,
		STATUS_KEY_DELETED,
		STATUS_NO_LOG_SPACE,
		STATUS_TOO_MANY_SIDS,
		STATUS_LM_CROSS_ENCRYPTION_REQUIRED,
		STATUS_KEY_HAS_CHILDREN,
		STATUS_CHILD_MUST_BE_VOLATILE,
		STATUS_DEVICE_CONFIGURATION_ERROR,
		STATUS_DRIVER_INTERNAL_ERROR,
		STATUS_INVALID_DEVICE_STATE,
		STATUS_IO_DEVICE_ERROR,
		STATUS_DEVICE_PROTOCOL_ERROR,
		STATUS_BACKUP_CONTROLLER,
		STATUS_LOG_FILE_FULL,
		STATUS_TOO_LATE,
		STATUS_NO_TRUST_LSA_SECRET,
		STATUS_NO_TRUST_SAM_ACCOUNT,
		STATUS_TRUSTED_DOMAIN_FAILURE,
		STATUS_TRUSTED_RELATIONSHIP_FAILURE,
		STATUS_EVENTLOG_FILE_CORRUPT,
		STATUS_EVENTLOG_CANT_START,
		STATUS_TRUST_FAILURE,
		STATUS_MUTANT_LIMIT_EXCEEDED,
		STATUS_NETLOGON_NOT_STARTED,
		STATUS_ACCOUNT_EXPIRED,
		STATUS_POSSIBLE_DEADLOCK,
		STATUS_NETWORK_CREDENTIAL_CONFLICT,
		STATUS_REMOTE_SESSION_LIMIT,
		STATUS_EVENTLOG_FILE_CHANGED,
		STATUS_NOLOGON_INTERDOMAIN_TRUST_ACCOUNT,
		STATUS_NOLOGON_WORKSTATION_TRUST_ACCOUNT,
		STATUS_NOLOGON_SERVER_TRUST_ACCOUNT,
		STATUS_DOMAIN_TRUST_INCONSISTENT,
		STATUS_FS_DRIVER_REQUIRED,
		STATUS_IMAGE_ALREADY_LOADED_AS_DLL,
		STATUS_INCOMPATIBLE_WITH_GLOBAL_SHORT_NAME_REGISTRY_SETTING,
		STATUS_SHORT_NAMES_NOT_ENABLED_ON_VOLUME,
		STATUS_SECURITY_STREAM_IS_INCONSISTENT,
		STATUS_INVALID_LOCK_RANGE,
		STATUS_INVALID_ACE_CONDITION,
		STATUS_IMAGE_SUBSYSTEM_NOT_PRESENT,
		STATUS_NOTIFICATION_GUID_ALREADY_DEFINED,
		STATUS_NETWORK_OPEN_RESTRICTION,
		STATUS_NO_USER_SESSION_KEY,
		STATUS_USER_SESSION_DELETED,
		STATUS_RESOURCE_LANG_NOT_FOUND,
		STATUS_INSUFF_SERVER_RESOURCES,
		STATUS_INVALID_BUFFER_SIZE,
		STATUS_INVALID_ADDRESS_COMPONENT,
		STATUS_INVALID_ADDRESS_WILDCARD,
		STATUS_TOO_MANY_ADDRESSES,
		STATUS_ADDRESS_ALREADY_EXISTS,
		STATUS_ADDRESS_CLOSED,
		STATUS_CONNECTION_DISCONNECTED,
		STATUS_CONNECTION_RESET,
		STATUS_TOO_MANY_NODES,
		STATUS_TRANSACTION_ABORTED,
		STATUS_TRANSACTION_TIMED_OUT,
		STATUS_TRANSACTION_NO_RELEASE,
		STATUS_TRANSACTION_NO_MATCH,
		STATUS_TRANSACTION_RESPONDED,
		STATUS_TRANSACTION_INVALID_ID,
		STATUS_TRANSACTION_INVALID_TYPE,
		STATUS_NOT_SERVER_SESSION,
		STATUS_NOT_CLIENT_SESSION,
		STATUS_CANNOT_LOAD_REGISTRY_FILE,
		STATUS_DEBUG_ATTACH_FAILED,
		STATUS_SYSTEM_PROCESS_TERMINATED,
		STATUS_DATA_NOT_ACCEPTED,
		STATUS_NO_BROWSER_SERVERS_FOUND,
		STATUS_VDM_HARD_ERROR,
		STATUS_DRIVER_CANCEL_TIMEOUT,
		STATUS_REPLY_MESSAGE_MISMATCH,
		STATUS_MAPPED_ALIGNMENT,
		STATUS_IMAGE_CHECKSUM_MISMATCH,
		STATUS_LOST_WRITEBEHIND_DATA,
		STATUS_CLIENT_SERVER_PARAMETERS_INVALID,
		STATUS_PASSWORD_MUST_CHANGE,
		STATUS_NOT_FOUND,
		STATUS_NOT_TINY_STREAM,
		STATUS_RECOVERY_FAILURE,
		STATUS_STACK_OVERFLOW_READ,
		STATUS_FAIL_CHECK,
		STATUS_DUPLICATE_OBJECTID,
		STATUS_OBJECTID_EXISTS,
		STATUS_CONVERT_TO_LARGE,
		STATUS_RETRY,
		STATUS_FOUND_OUT_OF_SCOPE,
		STATUS_ALLOCATE_BUCKET,
		STATUS_PROPSET_NOT_FOUND,
		STATUS_MARSHALL_OVERFLOW,
		STATUS_INVALID_VARIANT,
		STATUS_DOMAIN_CONTROLLER_NOT_FOUND,
		STATUS_ACCOUNT_LOCKED_OUT,
		STATUS_HANDLE_NOT_CLOSABLE,
		STATUS_CONNECTION_REFUSED,
		STATUS_GRACEFUL_DISCONNECT,
		STATUS_ADDRESS_ALREADY_ASSOCIATED,
		STATUS_ADDRESS_NOT_ASSOCIATED,
		STATUS_CONNECTION_INVALID,
		STATUS_CONNECTION_ACTIVE,
		STATUS_NETWORK_UNREACHABLE,
		STATUS_HOST_UNREACHABLE,
		STATUS_PROTOCOL_UNREACHABLE,
		STATUS_PORT_UNREACHABLE,
		STATUS_REQUEST_ABORTED,
		STATUS_CONNECTION_ABORTED,
		STATUS_BAD_COMPRESSION_BUFFER,
		STATUS_USER_MAPPED_FILE,
		STATUS_AUDIT_FAILED,
		STATUS_TIMER_RESOLUTION_NOT_SET,
		STATUS_CONNECTION_COUNT_LIMIT,
		STATUS_LOGIN_TIME_RESTRICTION,
		STATUS_LOGIN_WKSTA_RESTRICTION,
		STATUS_IMAGE_MP_UP_MISMATCH,
		STATUS_INSUFFICIENT_LOGON_INFO,
		STATUS_BAD_DLL_ENTRYPOINT,
		STATUS_BAD_SERVICE_ENTRYPOINT,
		STATUS_LPC_REPLY_LOST,
		STATUS_IP_ADDRESS_CONFLICT1,
		STATUS_IP_ADDRESS_CONFLICT2,
		STATUS_REGISTRY_QUOTA_LIMIT,
		STATUS_PATH_NOT_COVERED,
		STATUS_NO_CALLBACK_ACTIVE,
		STATUS_LICENSE_QUOTA_EXCEEDED,
		STATUS_PWD_TOO_SHORT,
		STATUS_PWD_TOO_RECENT,
		STATUS_PWD_HISTORY_CONFLICT,
		STATUS_PLUGPLAY_NO_DEVICE,
		STATUS_UNSUPPORTED_COMPRESSION,
		STATUS_INVALID_HW_PROFILE,
		STATUS_INVALID_PLUGPLAY_DEVICE_PATH,
		STATUS_DRIVER_ORDINAL_NOT_FOUND,
		STATUS_DRIVER_ENTRYPOINT_NOT_FOUND,
		STATUS_RESOURCE_NOT_OWNED,
		STATUS_TOO_MANY_LINKS,
		STATUS_QUOTA_LIST_INCONSISTENT,
		STATUS_FILE_IS_OFFLINE,
		STATUS_EVALUATION_EXPIRATION,
		STATUS_ILLEGAL_DLL_RELOCATION,
		STATUS_LICENSE_VIOLATION,
		STATUS_DLL_INIT_FAILED_LOGOFF,
		STATUS_DRIVER_UNABLE_TO_LOAD,
		STATUS_DFS_UNAVAILABLE,
		STATUS_VOLUME_DISMOUNTED,
		STATUS_WX86_INTERNAL_ERROR,
		STATUS_WX86_FLOAT_STACK_CHECK,
		STATUS_VALIDATE_CONTINUE,
		STATUS_NO_MATCH,
		STATUS_NO_MORE_MATCHES,
		STATUS_NOT_A_REPARSE_POINT,
		STATUS_IO_REPARSE_TAG_INVALID,
		STATUS_IO_REPARSE_TAG_MISMATCH,
		STATUS_IO_REPARSE_DATA_INVALID,
		STATUS_IO_REPARSE_TAG_NOT_HANDLED,
		STATUS_REPARSE_POINT_NOT_RESOLVED,
		STATUS_DIRECTORY_IS_A_REPARSE_POINT,
		STATUS_RANGE_LIST_CONFLICT,
		STATUS_SOURCE_ELEMENT_EMPTY,
		STATUS_DESTINATION_ELEMENT_FULL,
		STATUS_ILLEGAL_ELEMENT_ADDRESS,
		STATUS_MAGAZINE_NOT_PRESENT,
		STATUS_REINITIALIZATION_NEEDED,
		STATUS_DEVICE_REQUIRES_CLEANING,
		STATUS_DEVICE_DOOR_OPEN,
		STATUS_ENCRYPTION_FAILED,
		STATUS_DECRYPTION_FAILED,
		STATUS_RANGE_NOT_FOUND,
		STATUS_NO_RECOVERY_POLICY,
		STATUS_NO_EFS,
		STATUS_WRONG_EFS,
		STATUS_NO_USER_KEYS,
		STATUS_FILE_NOT_ENCRYPTED,
		STATUS_NOT_EXPORT_FORMAT,
		STATUS_FILE_ENCRYPTED,
		STATUS_WAKE_SYSTEM,
		STATUS_WMI_GUID_NOT_FOUND,
		STATUS_WMI_INSTANCE_NOT_FOUND,
		STATUS_WMI_ITEMID_NOT_FOUND,
		STATUS_WMI_TRY_AGAIN,
		STATUS_SHARED_POLICY,
		STATUS_POLICY_OBJECT_NOT_FOUND,
		STATUS_POLICY_ONLY_IN_DS,
		STATUS_VOLUME_NOT_UPGRADED,
		STATUS_REMOTE_STORAGE_NOT_ACTIVE,
		STATUS_REMOTE_STORAGE_MEDIA_ERROR,
		STATUS_NO_TRACKING_SERVICE,
		STATUS_SERVER_SID_MISMATCH,
		STATUS_DS_NO_ATTRIBUTE_OR_VALUE,
		STATUS_DS_INVALID_ATTRIBUTE_SYNTAX,
		STATUS_DS_ATTRIBUTE_TYPE_UNDEFINED,
		STATUS_DS_ATTRIBUTE_OR_VALUE_EXISTS,
		STATUS_DS_BUSY,
		STATUS_DS_UNAVAILABLE,
		STATUS_DS_NO_RIDS_ALLOCATED,
		STATUS_DS_NO_MORE_RIDS,
		STATUS_DS_INCORRECT_ROLE_OWNER,
		STATUS_DS_RIDMGR_INIT_ERROR,
		STATUS_DS_OBJ_CLASS_VIOLATION,
		STATUS_DS_CANT_ON_NON_LEAF,
		STATUS_DS_CANT_ON_RDN,
		STATUS_DS_CANT_MOD_OBJ_CLASS,
		STATUS_DS_CROSS_DOM_MOVE_FAILED,
		STATUS_DS_GC_NOT_AVAILABLE,
		STATUS_DIRECTORY_SERVICE_REQUIRED,
		STATUS_REPARSE_ATTRIBUTE_CONFLICT,
		STATUS_CANT_ENABLE_DENY_ONLY,
		STATUS_FLOAT_MULTIPLE_FAULTS,
		STATUS_FLOAT_MULTIPLE_TRAPS,
		STATUS_DEVICE_REMOVED,
		STATUS_JOURNAL_DELETE_IN_PROGRESS,
		STATUS_JOURNAL_NOT_ACTIVE,
		STATUS_NOINTERFACE,
		STATUS_DS_ADMIN_LIMIT_EXCEEDED,
		STATUS_DRIVER_FAILED_SLEEP,
		STATUS_MUTUAL_AUTHENTICATION_FAILED,
		STATUS_CORRUPT_SYSTEM_FILE,
		STATUS_DATATYPE_MISALIGNMENT_ERROR,
		STATUS_WMI_READ_ONLY,
		STATUS_WMI_SET_FAILURE,
		STATUS_COMMITMENT_MINIMUM,
		STATUS_REG_NAT_CONSUMPTION,
		STATUS_TRANSPORT_FULL,
		STATUS_DS_SAM_INIT_FAILURE,
		STATUS_ONLY_IF_CONNECTED,
		STATUS_DS_SENSITIVE_GROUP_VIOLATION,
		STATUS_PNP_RESTART_ENUMERATION,
		STATUS_JOURNAL_ENTRY_DELETED,
		STATUS_DS_CANT_MOD_PRIMARYGROUPID,
		STATUS_SYSTEM_IMAGE_BAD_SIGNATURE,
		STATUS_PNP_REBOOT_REQUIRED,
		STATUS_POWER_STATE_INVALID,
		STATUS_DS_INVALID_GROUP_TYPE,
		STATUS_DS_NO_NEST_GLOBALGROUP_IN_MIXEDDOMAIN,
		STATUS_DS_NO_NEST_LOCALGROUP_IN_MIXEDDOMAIN,
		STATUS_DS_GLOBAL_CANT_HAVE_LOCAL_MEMBER,
		STATUS_DS_GLOBAL_CANT_HAVE_UNIVERSAL_MEMBER,
		STATUS_DS_UNIVERSAL_CANT_HAVE_LOCAL_MEMBER,
		STATUS_DS_GLOBAL_CANT_HAVE_CROSSDOMAIN_MEMBER,
		STATUS_DS_LOCAL_CANT_HAVE_CROSSDOMAIN_LOCAL_MEMBER,
		STATUS_DS_HAVE_PRIMARY_MEMBERS,
		STATUS_WMI_NOT_SUPPORTED,
		STATUS_INSUFFICIENT_POWER,
		STATUS_SAM_NEED_BOOTKEY_PASSWORD,
		STATUS_SAM_NEED_BOOTKEY_FLOPPY,
		STATUS_DS_CANT_START,
		STATUS_DS_INIT_FAILURE,
		STATUS_SAM_INIT_FAILURE,
		STATUS_DS_GC_REQUIRED,
		STATUS_DS_LOCAL_MEMBER_OF_LOCAL_ONLY,
		STATUS_DS_NO_FPO_IN_UNIVERSAL_GROUPS,
		STATUS_DS_MACHINE_ACCOUNT_QUOTA_EXCEEDED,
		STATUS_MULTIPLE_FAULT_VIOLATION,
		STATUS_CURRENT_DOMAIN_NOT_ALLOWED,
		STATUS_CANNOT_MAKE,
		STATUS_SYSTEM_SHUTDOWN,
		STATUS_DS_INIT_FAILURE_CONSOLE,
		STATUS_DS_SAM_INIT_FAILURE_CONSOLE,
		STATUS_UNFINISHED_CONTEXT_DELETED,
		STATUS_NO_TGT_REPLY,
		STATUS_OBJECTID_NOT_FOUND,
		STATUS_NO_IP_ADDRESSES,
		STATUS_WRONG_CREDENTIAL_HANDLE,
		STATUS_CRYPTO_SYSTEM_INVALID,
		STATUS_MAX_REFERRALS_EXCEEDED,
		STATUS_MUST_BE_KDC,
		STATUS_STRONG_CRYPTO_NOT_SUPPORTED,
		STATUS_TOO_MANY_PRINCIPALS,
		STATUS_NO_PA_DATA,
		STATUS_PKINIT_NAME_MISMATCH,
		STATUS_SMARTCARD_LOGON_REQUIRED,
		STATUS_KDC_INVALID_REQUEST,
		STATUS_KDC_UNABLE_TO_REFER,
		STATUS_KDC_UNKNOWN_ETYPE,
		STATUS_SHUTDOWN_IN_PROGRESS,
		STATUS_SERVER_SHUTDOWN_IN_PROGRESS,
		STATUS_NOT_SUPPORTED_ON_SBS,
		STATUS_WMI_GUID_DISCONNECTED,
		STATUS_WMI_ALREADY_DISABLED,
		STATUS_WMI_ALREADY_ENABLED,
		STATUS_MFT_TOO_FRAGMENTED,
		STATUS_COPY_PROTECTION_FAILURE,
		STATUS_CSS_AUTHENTICATION_FAILURE,
		STATUS_CSS_KEY_NOT_PRESENT,
		STATUS_CSS_KEY_NOT_ESTABLISHED,
		STATUS_CSS_SCRAMBLED_SECTOR,
		STATUS_CSS_REGION_MISMATCH,
		STATUS_CSS_RESETS_EXHAUSTED,
		STATUS_PKINIT_FAILURE,
		STATUS_SMARTCARD_SUBSYSTEM_FAILURE,
		STATUS_NO_KERB_KEY,
		STATUS_HOST_DOWN,
		STATUS_UNSUPPORTED_PREAUTH,
		STATUS_EFS_ALG_BLOB_TOO_BIG,
		STATUS_PORT_NOT_SET,
		STATUS_DEBUGGER_INACTIVE,
		STATUS_DS_VERSION_CHECK_FAILURE,
		STATUS_AUDITING_DISABLED,
		STATUS_PRENT4_MACHINE_ACCOUNT,
		STATUS_DS_AG_CANT_HAVE_UNIVERSAL_MEMBER,
		STATUS_INVALID_IMAGE_WIN_32,
		STATUS_INVALID_IMAGE_WIN_64,
		STATUS_BAD_BINDINGS,
		STATUS_NETWORK_SESSION_EXPIRED,
		STATUS_APPHELP_BLOCK,
		STATUS_ALL_SIDS_FILTERED,
		STATUS_NOT_SAFE_MODE_DRIVER,
		STATUS_ACCESS_DISABLED_BY_POLICY_DEFAULT,
		STATUS_ACCESS_DISABLED_BY_POLICY_PATH,
		STATUS_ACCESS_DISABLED_BY_POLICY_PUBLISHER,
		STATUS_ACCESS_DISABLED_BY_POLICY_OTHER,
		STATUS_FAILED_DRIVER_ENTRY,
		STATUS_DEVICE_ENUMERATION_ERROR,
		STATUS_MOUNT_POINT_NOT_RESOLVED,
		STATUS_INVALID_DEVICE_OBJECT_PARAMETER,
		STATUS_MCA_OCCURED,
		STATUS_DRIVER_BLOCKED_CRITICAL,
		STATUS_DRIVER_BLOCKED,
		STATUS_DRIVER_DATABASE_ERROR,
		STATUS_SYSTEM_HIVE_TOO_LARGE,
		STATUS_INVALID_IMPORT_OF_NON_DLL,
		STATUS_DS_SHUTTING_DOWN,
		STATUS_NO_SECRETS,
		STATUS_ACCESS_DISABLED_NO_SAFER_UI_BY_POLICY,
		STATUS_FAILED_STACK_SWITCH,
		STATUS_HEAP_CORRUPTION,
		STATUS_SMARTCARD_WRONG_PIN,
		STATUS_SMARTCARD_CARD_BLOCKED,
		STATUS_SMARTCARD_CARD_NOT_AUTHENTICATED,
		STATUS_SMARTCARD_NO_CARD,
		STATUS_SMARTCARD_NO_KEY_CONTAINER,
		STATUS_SMARTCARD_NO_CERTIFICATE,
		STATUS_SMARTCARD_NO_KEYSET,
		STATUS_SMARTCARD_IO_ERROR,
		STATUS_DOWNGRADE_DETECTED,
		STATUS_SMARTCARD_CERT_REVOKED,
		STATUS_ISSUING_CA_UNTRUSTED,
		STATUS_REVOCATION_OFFLINE_C,
		STATUS_PKINIT_CLIENT_FAILURE,
		STATUS_SMARTCARD_CERT_EXPIRED,
		STATUS_DRIVER_FAILED_PRIOR_UNLOAD,
		STATUS_SMARTCARD_SILENT_CONTEXT,
		STATUS_PER_USER_TRUST_QUOTA_EXCEEDED,
		STATUS_ALL_USER_TRUST_QUOTA_EXCEEDED,
		STATUS_USER_DELETE_TRUST_QUOTA_EXCEEDED,
		STATUS_DS_NAME_NOT_UNIQUE,
		STATUS_DS_DUPLICATE_ID_FOUND,
		STATUS_DS_GROUP_CONVERSION_ERROR,
		STATUS_VOLSNAP_PREPARE_HIBERNATE,
		STATUS_USER2USER_REQUIRED,
		STATUS_STACK_BUFFER_OVERRUN,
		STATUS_NO_S4U_PROT_SUPPORT,
		STATUS_CROSSREALM_DELEGATION_FAILURE,
		STATUS_REVOCATION_OFFLINE_KDC,
		STATUS_ISSUING_CA_UNTRUSTED_KDC,
		STATUS_KDC_CERT_EXPIRED,
		STATUS_KDC_CERT_REVOKED,
		STATUS_PARAMETER_QUOTA_EXCEEDED,
		STATUS_HIBERNATION_FAILURE,
		STATUS_DELAY_LOAD_FAILED,
		STATUS_AUTHENTICATION_FIREWALL_FAILED,
		STATUS_VDM_DISALLOWED,
		STATUS_HUNG_DISPLAY_DRIVER_THREAD,
		STATUS_INSUFFICIENT_RESOURCE_FOR_SPECIFIED_SHARED_SECTION_SIZE,
		STATUS_INVALID_CRUNTIME_PARAMETER,
		STATUS_NTLM_BLOCKED,
		STATUS_DS_SRC_SID_EXISTS_IN_FOREST,
		STATUS_DS_DOMAIN_NAME_EXISTS_IN_FOREST,
		STATUS_DS_FLAT_NAME_EXISTS_IN_FOREST,
		STATUS_INVALID_USER_PRINCIPAL_NAME,
		STATUS_FATAL_USER_CALLBACK_EXCEPTION,
		STATUS_ASSERTION_FAILURE,
		STATUS_VERIFIER_STOP,
		STATUS_CALLBACK_POP_STACK,
		STATUS_INCOMPATIBLE_DRIVER_BLOCKED,
		STATUS_HIVE_UNLOADED,
		STATUS_COMPRESSION_DISABLED,
		STATUS_FILE_SYSTEM_LIMITATION,
		STATUS_INVALID_IMAGE_HASH,
		STATUS_NOT_CAPABLE,
		STATUS_REQUEST_OUT_OF_SEQUENCE,
		STATUS_IMPLEMENTATION_LIMIT,
		STATUS_ELEVATION_REQUIRED,
		STATUS_NO_SECURITY_CONTEXT,
		STATUS_PKU2U_CERT_FAILURE,
		STATUS_BEYOND_VDL,
		STATUS_ENCOUNTERED_WRITE_IN_PROGRESS,
		STATUS_PTE_CHANGED,
		STATUS_PURGE_FAILED,
		STATUS_CRED_REQUIRES_CONFIRMATION,
		STATUS_CS_ENCRYPTION_INVALID_SERVER_RESPONSE,
		STATUS_CS_ENCRYPTION_UNSUPPORTED_SERVER,
		STATUS_CS_ENCRYPTION_EXISTING_ENCRYPTED_FILE,
		STATUS_CS_ENCRYPTION_NEW_ENCRYPTED_FILE,
		STATUS_CS_ENCRYPTION_FILE_NOT_CSE,
		STATUS_INVALID_LABEL,
		STATUS_DRIVER_PROCESS_TERMINATED,
		STATUS_AMBIGUOUS_SYSTEM_DEVICE,
		STATUS_SYSTEM_DEVICE_NOT_FOUND,
		STATUS_RESTART_BOOT_APPLICATION,
		STATUS_INSUFFICIENT_NVRAM_RESOURCES,
		STATUS_INVALID_TASK_NAME,
		STATUS_INVALID_TASK_INDEX,
		STATUS_THREAD_ALREADY_IN_TASK,
		STATUS_CALLBACK_BYPASS,
		STATUS_FAIL_FAST_EXCEPTION,
		STATUS_IMAGE_CERT_REVOKED,
		STATUS_PORT_CLOSED,
		STATUS_MESSAGE_LOST,
		STATUS_INVALID_MESSAGE,
		STATUS_REQUEST_CANCELED,
		STATUS_RECURSIVE_DISPATCH,
		STATUS_LPC_RECEIVE_BUFFER_EXPECTED,
		STATUS_LPC_INVALID_CONNECTION_USAGE,
		STATUS_LPC_REQUESTS_NOT_ALLOWED,
		STATUS_RESOURCE_IN_USE,
		STATUS_HARDWARE_MEMORY_ERROR,
		STATUS_THREADPOOL_HANDLE_EXCEPTION,
		STATUS_THREADPOOL_SET_EVENT_ON_COMPLETION_FAILED,
		STATUS_THREADPOOL_RELEASE_SEMAPHORE_ON_COMPLETION_FAILED,
		STATUS_THREADPOOL_RELEASE_MUTEX_ON_COMPLETION_FAILED,
		STATUS_THREADPOOL_FREE_LIBRARY_ON_COMPLETION_FAILED,
		STATUS_THREADPOOL_RELEASED_DURING_OPERATION,
		STATUS_CALLBACK_RETURNED_WHILE_IMPERSONATING,
		STATUS_APC_RETURNED_WHILE_IMPERSONATING,
		STATUS_PROCESS_IS_PROTECTED,
		STATUS_MCA_EXCEPTION,
		STATUS_CERTIFICATE_MAPPING_NOT_UNIQUE,
		STATUS_SYMLINK_CLASS_DISABLED,
		STATUS_INVALID_IDN_NORMALIZATION,
		STATUS_NO_UNICODE_TRANSLATION,
		STATUS_ALREADY_REGISTERED,
		STATUS_CONTEXT_MISMATCH,
		STATUS_PORT_ALREADY_HAS_COMPLETION_LIST,
		STATUS_CALLBACK_RETURNED_THREAD_PRIORITY,
		STATUS_INVALID_THREAD,
		STATUS_CALLBACK_RETURNED_TRANSACTION,
		STATUS_CALLBACK_RETURNED_LDR_LOCK,
		STATUS_CALLBACK_RETURNED_LANG,
		STATUS_CALLBACK_RETURNED_PRI_BACK,
		STATUS_CALLBACK_RETURNED_THREAD_AFFINITY,
		STATUS_DISK_REPAIR_DISABLED,
		STATUS_DS_DOMAIN_RENAME_IN_PROGRESS,
		STATUS_DISK_QUOTA_EXCEEDED,
		STATUS_DATA_LOST_REPAIR,
		STATUS_CONTENT_BLOCKED,
		STATUS_BAD_CLUSTERS,
		STATUS_VOLUME_DIRTY,
		STATUS_FILE_CHECKED_OUT,
		STATUS_CHECKOUT_REQUIRED,
		STATUS_BAD_FILE_TYPE,
		STATUS_FILE_TOO_LARGE,
		STATUS_FORMS_AUTH_REQUIRED,
		STATUS_VIRUS_INFECTED,
		STATUS_VIRUS_DELETED,
		STATUS_BAD_MCFG_TABLE,
		STATUS_CANNOT_BREAK_OPLOCK,
		STATUS_WOW_ASSERTION,
		STATUS_INVALID_SIGNATURE,
		STATUS_HMAC_NOT_SUPPORTED,
		STATUS_AUTH_TAG_MISMATCH,
		STATUS_IPSEC_QUEUE_OVERFLOW,
		STATUS_ND_QUEUE_OVERFLOW,
		STATUS_HOPLIMIT_EXCEEDED,
		STATUS_PROTOCOL_NOT_SUPPORTED,
		STATUS_FASTPATH_REJECTED,
		STATUS_LOST_WRITEBEHIND_DATA_NETWORK_DISCONNECTED,
		STATUS_LOST_WRITEBEHIND_DATA_NETWORK_SERVER_ERROR,
		STATUS_LOST_WRITEBEHIND_DATA_LOCAL_DISK_ERROR,
		STATUS_XML_PARSE_ERROR,
		STATUS_XMLDSIG_ERROR,
		STATUS_WRONG_COMPARTMENT,
		STATUS_AUTHIP_FAILURE,
		STATUS_DS_OID_MAPPED_GROUP_CANT_HAVE_MEMBERS,
		STATUS_DS_OID_NOT_FOUND,
		STATUS_HASH_NOT_SUPPORTED,
		STATUS_HASH_NOT_PRESENT,
		DBG_NO_STATE_CHANGE,
		DBG_APP_NOT_IDLE,
		RPC_NT_INVALID_STRING_BINDING,
		RPC_NT_WRONG_KIND_OF_BINDING,
		RPC_NT_INVALID_BINDING,
		RPC_NT_PROTSEQ_NOT_SUPPORTED,
		RPC_NT_INVALID_RPC_PROTSEQ,
		RPC_NT_INVALID_STRING_UUID,
		RPC_NT_INVALID_ENDPOINT_FORMAT,
		RPC_NT_INVALID_NET_ADDR,
		RPC_NT_NO_ENDPOINT_FOUND,
		RPC_NT_INVALID_TIMEOUT,
		RPC_NT_OBJECT_NOT_FOUND,
		RPC_NT_ALREADY_REGISTERED,
		RPC_NT_TYPE_ALREADY_REGISTERED,
		RPC_NT_ALREADY_LISTENING,
		RPC_NT_NO_PROTSEQS_REGISTERED,
		RPC_NT_NOT_LISTENING,
		RPC_NT_UNKNOWN_MGR_TYPE,
		RPC_NT_UNKNOWN_IF,
		RPC_NT_NO_BINDINGS,
		RPC_NT_NO_PROTSEQS,
		RPC_NT_CANT_CREATE_ENDPOINT,
		RPC_NT_OUT_OF_RESOURCES,
		RPC_NT_SERVER_UNAVAILABLE,
		RPC_NT_SERVER_TOO_BUSY,
		RPC_NT_INVALID_NETWORK_OPTIONS,
		RPC_NT_NO_CALL_ACTIVE,
		RPC_NT_CALL_FAILED,
		RPC_NT_CALL_FAILED_DNE,
		RPC_NT_PROTOCOL_ERROR,
		RPC_NT_UNSUPPORTED_TRANS_SYN,
		RPC_NT_UNSUPPORTED_TYPE,
		RPC_NT_INVALID_TAG,
		RPC_NT_INVALID_BOUND,
		RPC_NT_NO_ENTRY_NAME,
		RPC_NT_INVALID_NAME_SYNTAX,
		RPC_NT_UNSUPPORTED_NAME_SYNTAX,
		RPC_NT_UUID_NO_ADDRESS,
		RPC_NT_DUPLICATE_ENDPOINT,
		RPC_NT_UNKNOWN_AUTHN_TYPE,
		RPC_NT_MAX_CALLS_TOO_SMALL,
		RPC_NT_STRING_TOO_LONG,
		RPC_NT_PROTSEQ_NOT_FOUND,
		RPC_NT_PROCNUM_OUT_OF_RANGE,
		RPC_NT_BINDING_HAS_NO_AUTH,
		RPC_NT_UNKNOWN_AUTHN_SERVICE,
		RPC_NT_UNKNOWN_AUTHN_LEVEL,
		RPC_NT_INVALID_AUTH_IDENTITY,
		RPC_NT_UNKNOWN_AUTHZ_SERVICE,
		EPT_NT_INVALID_ENTRY,
		EPT_NT_CANT_PERFORM_OP,
		EPT_NT_NOT_REGISTERED,
		RPC_NT_NOTHING_TO_EXPORT,
		RPC_NT_INCOMPLETE_NAME,
		RPC_NT_INVALID_VERS_OPTION,
		RPC_NT_NO_MORE_MEMBERS,
		RPC_NT_NOT_ALL_OBJS_UNEXPORTED,
		RPC_NT_INTERFACE_NOT_FOUND,
		RPC_NT_ENTRY_ALREADY_EXISTS,
		RPC_NT_ENTRY_NOT_FOUND,
		RPC_NT_NAME_SERVICE_UNAVAILABLE,
		RPC_NT_INVALID_NAF_ID,
		RPC_NT_CANNOT_SUPPORT,
		RPC_NT_NO_CONTEXT_AVAILABLE,
		RPC_NT_INTERNAL_ERROR,
		RPC_NT_ZERO_DIVIDE,
		RPC_NT_ADDRESS_ERROR,
		RPC_NT_FP_DIV_ZERO,
		RPC_NT_FP_UNDERFLOW,
		RPC_NT_FP_OVERFLOW,
		RPC_NT_NO_MORE_ENTRIES,
		RPC_NT_SS_CHAR_TRANS_OPEN_FAIL,
		RPC_NT_SS_CHAR_TRANS_SHORT_FILE,
		RPC_NT_SS_IN_NULL_CONTEXT,
		RPC_NT_SS_CONTEXT_MISMATCH,
		RPC_NT_SS_CONTEXT_DAMAGED,
		RPC_NT_SS_HANDLES_MISMATCH,
		RPC_NT_SS_CANNOT_GET_CALL_HANDLE,
		RPC_NT_NULL_REF_POINTER,
		RPC_NT_ENUM_VALUE_OUT_OF_RANGE,
		RPC_NT_BYTE_COUNT_TOO_SMALL,
		RPC_NT_BAD_STUB_DATA,
		RPC_NT_CALL_IN_PROGRESS,
		RPC_NT_NO_MORE_BINDINGS,
		RPC_NT_GROUP_MEMBER_NOT_FOUND,
		EPT_NT_CANT_CREATE,
		RPC_NT_INVALID_OBJECT,
		RPC_NT_NO_INTERFACES,
		RPC_NT_CALL_CANCELLED,
		RPC_NT_BINDING_INCOMPLETE,
		RPC_NT_COMM_FAILURE,
		RPC_NT_UNSUPPORTED_AUTHN_LEVEL,
		RPC_NT_NO_PRINC_NAME,
		RPC_NT_NOT_RPC_ERROR,
		RPC_NT_UUID_LOCAL_ONLY,
		RPC_NT_SEC_PKG_ERROR,
		RPC_NT_NOT_CANCELLED,
		RPC_NT_INVALID_ES_ACTION,
		RPC_NT_WRONG_ES_VERSION,
		RPC_NT_WRONG_STUB_VERSION,
		RPC_NT_INVALID_PIPE_OBJECT,
		RPC_NT_INVALID_PIPE_OPERATION,
		RPC_NT_WRONG_PIPE_VERSION,
		RPC_NT_PIPE_CLOSED,
		RPC_NT_PIPE_DISCIPLINE_ERROR,
		RPC_NT_PIPE_EMPTY,
		RPC_NT_INVALID_ASYNC_HANDLE,
		RPC_NT_INVALID_ASYNC_CALL,
		RPC_NT_PROXY_ACCESS_DENIED,
		RPC_NT_COOKIE_AUTH_FAILED,
		RPC_NT_SEND_INCOMPLETE,
		STATUS_ACPI_INVALID_OPCODE,
		STATUS_ACPI_STACK_OVERFLOW,
		STATUS_ACPI_ASSERT_FAILED,
		STATUS_ACPI_INVALID_INDEX,
		STATUS_ACPI_INVALID_ARGUMENT,
		STATUS_ACPI_FATAL,
		STATUS_ACPI_INVALID_SUPERNAME,
		STATUS_ACPI_INVALID_ARGTYPE,
		STATUS_ACPI_INVALID_OBJTYPE,
		STATUS_ACPI_INVALID_TARGETTYPE,
		STATUS_ACPI_INCORRECT_ARGUMENT_COUNT,
		STATUS_ACPI_ADDRESS_NOT_MAPPED,
		STATUS_ACPI_INVALID_EVENTTYPE,
		STATUS_ACPI_HANDLER_COLLISION,
		STATUS_ACPI_INVALID_DATA,
		STATUS_ACPI_INVALID_REGION,
		STATUS_ACPI_INVALID_ACCESS_SIZE,
		STATUS_ACPI_ACQUIRE_GLOBAL_LOCK,
		STATUS_ACPI_ALREADY_INITIALIZED,
		STATUS_ACPI_NOT_INITIALIZED,
		STATUS_ACPI_INVALID_MUTEX_LEVEL,
		STATUS_ACPI_MUTEX_NOT_OWNED,
		STATUS_ACPI_MUTEX_NOT_OWNER,
		STATUS_ACPI_RS_ACCESS,
		STATUS_ACPI_INVALID_TABLE,
		STATUS_ACPI_REG_HANDLER_FAILED,
		STATUS_ACPI_POWER_REQUEST_FAILED,
		STATUS_CTX_WINSTATION_NAME_INVALID,
		STATUS_CTX_INVALID_PD,
		STATUS_CTX_PD_NOT_FOUND,
		STATUS_CTX_CDM_CONNECT,
		STATUS_CTX_CDM_DISCONNECT,
		STATUS_CTX_CLOSE_PENDING,
		STATUS_CTX_NO_OUTBUF,
		STATUS_CTX_MODEM_INF_NOT_FOUND,
		STATUS_CTX_INVALID_MODEMNAME,
		STATUS_CTX_RESPONSE_ERROR,
		STATUS_CTX_MODEM_RESPONSE_TIMEOUT,
		STATUS_CTX_MODEM_RESPONSE_NO_CARRIER,
		STATUS_CTX_MODEM_RESPONSE_NO_DIALTONE,
		STATUS_CTX_MODEM_RESPONSE_BUSY,
		STATUS_CTX_MODEM_RESPONSE_VOICE,
		STATUS_CTX_TD_ERROR,
		STATUS_CTX_LICENSE_CLIENT_INVALID,
		STATUS_CTX_LICENSE_NOT_AVAILABLE,
		STATUS_CTX_LICENSE_EXPIRED,
		STATUS_CTX_WINSTATION_NOT_FOUND,
		STATUS_CTX_WINSTATION_NAME_COLLISION,
		STATUS_CTX_WINSTATION_BUSY,
		STATUS_CTX_BAD_VIDEO_MODE,
		STATUS_CTX_GRAPHICS_INVALID,
		STATUS_CTX_NOT_CONSOLE,
		STATUS_CTX_CLIENT_QUERY_TIMEOUT,
		STATUS_CTX_CONSOLE_DISCONNECT,
		STATUS_CTX_CONSOLE_CONNECT,
		STATUS_CTX_SHADOW_DENIED,
		STATUS_CTX_WINSTATION_ACCESS_DENIED,
		STATUS_CTX_INVALID_WD,
		STATUS_CTX_WD_NOT_FOUND,
		STATUS_CTX_SHADOW_INVALID,
		STATUS_CTX_SHADOW_DISABLED,
		STATUS_RDP_PROTOCOL_ERROR,
		STATUS_CTX_CLIENT_LICENSE_NOT_SET,
		STATUS_CTX_CLIENT_LICENSE_IN_USE,
		STATUS_CTX_SHADOW_ENDED_BY_MODE_CHANGE,
		STATUS_CTX_SHADOW_NOT_RUNNING,
		STATUS_CTX_LOGON_DISABLED,
		STATUS_CTX_SECURITY_LAYER_ERROR,
		STATUS_TS_INCOMPATIBLE_SESSIONS,
		STATUS_TS_VIDEO_SUBSYSTEM_ERROR,
		STATUS_PNP_BAD_MPS_TABLE,
		STATUS_PNP_TRANSLATION_FAILED,
		STATUS_PNP_IRQ_TRANSLATION_FAILED,
		STATUS_PNP_INVALID_ID,
		STATUS_IO_REISSUE_AS_CACHED,
		STATUS_MUI_FILE_NOT_FOUND,
		STATUS_MUI_INVALID_FILE,
		STATUS_MUI_INVALID_RC_CONFIG,
		STATUS_MUI_INVALID_LOCALE_NAME,
		STATUS_MUI_INVALID_ULTIMATEFALLBACK_NAME,
		STATUS_MUI_FILE_NOT_LOADED,
		STATUS_RESOURCE_ENUM_USER_STOP,
		STATUS_FLT_NO_HANDLER_DEFINED,
		STATUS_FLT_CONTEXT_ALREADY_DEFINED,
		STATUS_FLT_INVALID_ASYNCHRONOUS_REQUEST,
		STATUS_FLT_DISALLOW_FAST_IO,
		STATUS_FLT_INVALID_NAME_REQUEST,
		STATUS_FLT_NOT_SAFE_TO_POST_OPERATION,
		STATUS_FLT_NOT_INITIALIZED,
		STATUS_FLT_FILTER_NOT_READY,
		STATUS_FLT_POST_OPERATION_CLEANUP,
		STATUS_FLT_INTERNAL_ERROR,
		STATUS_FLT_DELETING_OBJECT,
		STATUS_FLT_MUST_BE_NONPAGED_POOL,
		STATUS_FLT_DUPLICATE_ENTRY,
		STATUS_FLT_CBDQ_DISABLED,
		STATUS_FLT_DO_NOT_ATTACH,
		STATUS_FLT_DO_NOT_DETACH,
		STATUS_FLT_INSTANCE_ALTITUDE_COLLISION,
		STATUS_FLT_INSTANCE_NAME_COLLISION,
		STATUS_FLT_FILTER_NOT_FOUND,
		STATUS_FLT_VOLUME_NOT_FOUND,
		STATUS_FLT_INSTANCE_NOT_FOUND,
		STATUS_FLT_CONTEXT_ALLOCATION_NOT_FOUND,
		STATUS_FLT_INVALID_CONTEXT_REGISTRATION,
		STATUS_FLT_NAME_CACHE_MISS,
		STATUS_FLT_NO_DEVICE_OBJECT,
		STATUS_FLT_VOLUME_ALREADY_MOUNTED,
		STATUS_FLT_ALREADY_ENLISTED,
		STATUS_FLT_CONTEXT_ALREADY_LINKED,
		STATUS_FLT_NO_WAITER_FOR_REPLY,
		STATUS_SXS_SECTION_NOT_FOUND,
		STATUS_SXS_CANT_GEN_ACTCTX,
		STATUS_SXS_INVALID_ACTCTXDATA_FORMAT,
		STATUS_SXS_ASSEMBLY_NOT_FOUND,
		STATUS_SXS_MANIFEST_FORMAT_ERROR,
		STATUS_SXS_MANIFEST_PARSE_ERROR,
		STATUS_SXS_ACTIVATION_CONTEXT_DISABLED,
		STATUS_SXS_KEY_NOT_FOUND,
		STATUS_SXS_VERSION_CONFLICT,
		STATUS_SXS_WRONG_SECTION_TYPE,
		STATUS_SXS_THREAD_QUERIES_DISABLED,
		STATUS_SXS_ASSEMBLY_MISSING,
		STATUS_SXS_RELEASE_ACTIVATION_CONTEXT,
		STATUS_SXS_PROCESS_DEFAULT_ALREADY_SET,
		STATUS_SXS_EARLY_DEACTIVATION,
		STATUS_SXS_INVALID_DEACTIVATION,
		STATUS_SXS_MULTIPLE_DEACTIVATION,
		STATUS_SXS_SYSTEM_DEFAULT_ACTIVATION_CONTEXT_EMPTY,
		STATUS_SXS_PROCESS_TERMINATION_REQUESTED,
		STATUS_SXS_CORRUPT_ACTIVATION_STACK,
		STATUS_SXS_CORRUPTION,
		STATUS_SXS_INVALID_IDENTITY_ATTRIBUTE_VALUE,
		STATUS_SXS_INVALID_IDENTITY_ATTRIBUTE_NAME,
		STATUS_SXS_IDENTITY_DUPLICATE_ATTRIBUTE,
		STATUS_SXS_IDENTITY_PARSE_ERROR,
		STATUS_SXS_COMPONENT_STORE_CORRUPT,
		STATUS_SXS_FILE_HASH_MISMATCH,
		STATUS_SXS_MANIFEST_IDENTITY_SAME_BUT_CONTENTS_DIFFERENT,
		STATUS_SXS_IDENTITIES_DIFFERENT,
		STATUS_SXS_ASSEMBLY_IS_NOT_A_DEPLOYMENT,
		STATUS_SXS_FILE_NOT_PART_OF_ASSEMBLY,
		STATUS_ADVANCED_INSTALLER_FAILED,
		STATUS_XML_ENCODING_MISMATCH,
		STATUS_SXS_MANIFEST_TOO_BIG,
		STATUS_SXS_SETTING_NOT_REGISTERED,
		STATUS_SXS_TRANSACTION_CLOSURE_INCOMPLETE,
		STATUS_SMI_PRIMITIVE_INSTALLER_FAILED,
		STATUS_GENERIC_COMMAND_FAILED,
		STATUS_SXS_FILE_HASH_MISSING,
		STATUS_CLUSTER_INVALID_NODE,
		STATUS_CLUSTER_NODE_EXISTS,
		STATUS_CLUSTER_JOIN_IN_PROGRESS,
		STATUS_CLUSTER_NODE_NOT_FOUND,
		STATUS_CLUSTER_LOCAL_NODE_NOT_FOUND,
		STATUS_CLUSTER_NETWORK_EXISTS,
		STATUS_CLUSTER_NETWORK_NOT_FOUND,
		STATUS_CLUSTER_NETINTERFACE_EXISTS,
		STATUS_CLUSTER_NETINTERFACE_NOT_FOUND,
		STATUS_CLUSTER_INVALID_REQUEST,
		STATUS_CLUSTER_INVALID_NETWORK_PROVIDER,
		STATUS_CLUSTER_NODE_DOWN,
		STATUS_CLUSTER_NODE_UNREACHABLE,
		STATUS_CLUSTER_NODE_NOT_MEMBER,
		STATUS_CLUSTER_JOIN_NOT_IN_PROGRESS,
		STATUS_CLUSTER_INVALID_NETWORK,
		STATUS_CLUSTER_NO_NET_ADAPTERS,
		STATUS_CLUSTER_NODE_UP,
		STATUS_CLUSTER_NODE_PAUSED,
		STATUS_CLUSTER_NODE_NOT_PAUSED,
		STATUS_CLUSTER_NO_SECURITY_CONTEXT,
		STATUS_CLUSTER_NETWORK_NOT_INTERNAL,
		STATUS_CLUSTER_POISONED,
		STATUS_CLUSTER_NON_CSV_PATH,
		STATUS_CLUSTER_CSV_VOLUME_NOT_LOCAL,
		STATUS_TRANSACTIONAL_CONFLICT,
		STATUS_INVALID_TRANSACTION,
		STATUS_TRANSACTION_NOT_ACTIVE,
		STATUS_TM_INITIALIZATION_FAILED,
		STATUS_RM_NOT_ACTIVE,
		STATUS_RM_METADATA_CORRUPT,
		STATUS_TRANSACTION_NOT_JOINED,
		STATUS_DIRECTORY_NOT_RM,
		STATUS_COULD_NOT_RESIZE_LOG,
		STATUS_TRANSACTIONS_UNSUPPORTED_REMOTE,
		STATUS_LOG_RESIZE_INVALID_SIZE,
		STATUS_REMOTE_FILE_VERSION_MISMATCH,
		STATUS_CRM_PROTOCOL_ALREADY_EXISTS,
		STATUS_TRANSACTION_PROPAGATION_FAILED,
		STATUS_CRM_PROTOCOL_NOT_FOUND,
		STATUS_TRANSACTION_SUPERIOR_EXISTS,
		STATUS_TRANSACTION_REQUEST_NOT_VALID,
		STATUS_TRANSACTION_NOT_REQUESTED,
		STATUS_TRANSACTION_ALREADY_ABORTED,
		STATUS_TRANSACTION_ALREADY_COMMITTED,
		STATUS_TRANSACTION_INVALID_MARSHALL_BUFFER,
		STATUS_CURRENT_TRANSACTION_NOT_VALID,
		STATUS_LOG_GROWTH_FAILED,
		STATUS_OBJECT_NO_LONGER_EXISTS,
		STATUS_STREAM_MINIVERSION_NOT_FOUND,
		STATUS_STREAM_MINIVERSION_NOT_VALID,
		STATUS_MINIVERSION_INACCESSIBLE_FROM_SPECIFIED_TRANSACTION,
		STATUS_CANT_OPEN_MINIVERSION_WITH_MODIFY_INTENT,
		STATUS_CANT_CREATE_MORE_STREAM_MINIVERSIONS,
		STATUS_HANDLE_NO_LONGER_VALID,
		STATUS_NO_TXF_METADATA,
		STATUS_LOG_CORRUPTION_DETECTED,
		STATUS_CANT_RECOVER_WITH_HANDLE_OPEN,
		STATUS_RM_DISCONNECTED,
		STATUS_ENLISTMENT_NOT_SUPERIOR,
		STATUS_RECOVERY_NOT_NEEDED,
		STATUS_RM_ALREADY_STARTED,
		STATUS_FILE_IDENTITY_NOT_PERSISTENT,
		STATUS_CANT_BREAK_TRANSACTIONAL_DEPENDENCY,
		STATUS_CANT_CROSS_RM_BOUNDARY,
		STATUS_TXF_DIR_NOT_EMPTY,
		STATUS_INDOUBT_TRANSACTIONS_EXIST,
		STATUS_TM_VOLATILE,
		STATUS_ROLLBACK_TIMER_EXPIRED,
		STATUS_TXF_ATTRIBUTE_CORRUPT,
		STATUS_EFS_NOT_ALLOWED_IN_TRANSACTION,
		STATUS_TRANSACTIONAL_OPEN_NOT_ALLOWED,
		STATUS_TRANSACTED_MAPPING_UNSUPPORTED_REMOTE,
		STATUS_TXF_METADATA_ALREADY_PRESENT,
		STATUS_TRANSACTION_SCOPE_CALLBACKS_NOT_SET,
		STATUS_TRANSACTION_REQUIRED_PROMOTION,
		STATUS_CANNOT_EXECUTE_FILE_IN_TRANSACTION,
		STATUS_TRANSACTIONS_NOT_FROZEN,
		STATUS_TRANSACTION_FREEZE_IN_PROGRESS,
		STATUS_NOT_SNAPSHOT_VOLUME,
		STATUS_NO_SAVEPOINT_WITH_OPEN_FILES,
		STATUS_SPARSE_NOT_ALLOWED_IN_TRANSACTION,
		STATUS_TM_IDENTITY_MISMATCH,
		STATUS_FLOATED_SECTION,
		STATUS_CANNOT_ACCEPT_TRANSACTED_WORK,
		STATUS_CANNOT_ABORT_TRANSACTIONS,
		STATUS_TRANSACTION_NOT_FOUND,
		STATUS_RESOURCEMANAGER_NOT_FOUND,
		STATUS_ENLISTMENT_NOT_FOUND,
		STATUS_TRANSACTIONMANAGER_NOT_FOUND,
		STATUS_TRANSACTIONMANAGER_NOT_ONLINE,
		STATUS_TRANSACTIONMANAGER_RECOVERY_NAME_COLLISION,
		STATUS_TRANSACTION_NOT_ROOT,
		STATUS_TRANSACTION_OBJECT_EXPIRED,
		STATUS_COMPRESSION_NOT_ALLOWED_IN_TRANSACTION,
		STATUS_TRANSACTION_RESPONSE_NOT_ENLISTED,
		STATUS_TRANSACTION_RECORD_TOO_LONG,
		STATUS_NO_LINK_TRACKING_IN_TRANSACTION,
		STATUS_OPERATION_NOT_SUPPORTED_IN_TRANSACTION,
		STATUS_TRANSACTION_INTEGRITY_VIOLATED,
		STATUS_TRANSACTIONMANAGER_IDENTITY_MISMATCH,
		STATUS_RM_CANNOT_BE_FROZEN_FOR_SNAPSHOT,
		STATUS_TRANSACTION_MUST_WRITETHROUGH,
		STATUS_TRANSACTION_NO_SUPERIOR,
		STATUS_EXPIRED_HANDLE,
		STATUS_TRANSACTION_NOT_ENLISTED,
		STATUS_LOG_SECTOR_INVALID,
		STATUS_LOG_SECTOR_PARITY_INVALID,
		STATUS_LOG_SECTOR_REMAPPED,
		STATUS_LOG_BLOCK_INCOMPLETE,
		STATUS_LOG_INVALID_RANGE,
		STATUS_LOG_BLOCKS_EXHAUSTED,
		STATUS_LOG_READ_CONTEXT_INVALID,
		STATUS_LOG_RESTART_INVALID,
		STATUS_LOG_BLOCK_VERSION,
		STATUS_LOG_BLOCK_INVALID,
		STATUS_LOG_READ_MODE_INVALID,
		STATUS_LOG_NO_RESTART,
		STATUS_LOG_METADATA_CORRUPT,
		STATUS_LOG_METADATA_INVALID,
		STATUS_LOG_METADATA_INCONSISTENT,
		STATUS_LOG_RESERVATION_INVALID,
		STATUS_LOG_CANT_DELETE,
		STATUS_LOG_CONTAINER_LIMIT_EXCEEDED,
		STATUS_LOG_START_OF_LOG,
		STATUS_LOG_POLICY_ALREADY_INSTALLED,
		STATUS_LOG_POLICY_NOT_INSTALLED,
		STATUS_LOG_POLICY_INVALID,
		STATUS_LOG_POLICY_CONFLICT,
		STATUS_LOG_PINNED_ARCHIVE_TAIL,
		STATUS_LOG_RECORD_NONEXISTENT,
		STATUS_LOG_RECORDS_RESERVED_INVALID,
		STATUS_LOG_SPACE_RESERVED_INVALID,
		STATUS_LOG_TAIL_INVALID,
		STATUS_LOG_FULL,
		STATUS_LOG_MULTIPLEXED,
		STATUS_LOG_DEDICATED,
		STATUS_LOG_ARCHIVE_NOT_IN_PROGRESS,
		STATUS_LOG_ARCHIVE_IN_PROGRESS,
		STATUS_LOG_EPHEMERAL,
		STATUS_LOG_NOT_ENOUGH_CONTAINERS,
		STATUS_LOG_CLIENT_ALREADY_REGISTERED,
		STATUS_LOG_CLIENT_NOT_REGISTERED,
		STATUS_LOG_FULL_HANDLER_IN_PROGRESS,
		STATUS_LOG_CONTAINER_READ_FAILED,
		STATUS_LOG_CONTAINER_WRITE_FAILED,
		STATUS_LOG_CONTAINER_OPEN_FAILED,
		STATUS_LOG_CONTAINER_STATE_INVALID,
		STATUS_LOG_STATE_INVALID,
		STATUS_LOG_PINNED,
		STATUS_LOG_METADATA_FLUSH_FAILED,
		STATUS_LOG_INCONSISTENT_SECURITY,
		STATUS_LOG_APPENDED_FLUSH_FAILED,
		STATUS_LOG_PINNED_RESERVATION,
		STATUS_VIDEO_HUNG_DISPLAY_DRIVER_THREAD,
		STATUS_VIDEO_HUNG_DISPLAY_DRIVER_THREAD_RECOVERED,
		STATUS_VIDEO_DRIVER_DEBUG_REPORT_REQUEST,
		STATUS_MONITOR_NO_DESCRIPTOR,
		STATUS_MONITOR_UNKNOWN_DESCRIPTOR_FORMAT,
		STATUS_MONITOR_INVALID_DESCRIPTOR_CHECKSUM,
		STATUS_MONITOR_INVALID_STANDARD_TIMING_BLOCK,
		STATUS_MONITOR_WMI_DATABLOCK_REGISTRATION_FAILED,
		STATUS_MONITOR_INVALID_SERIAL_NUMBER_MONDSC_BLOCK,
		STATUS_MONITOR_INVALID_USER_FRIENDLY_MONDSC_BLOCK,
		STATUS_MONITOR_NO_MORE_DESCRIPTOR_DATA,
		STATUS_MONITOR_INVALID_DETAILED_TIMING_BLOCK,
		STATUS_MONITOR_INVALID_MANUFACTURE_DATE,
		STATUS_GRAPHICS_NOT_EXCLUSIVE_MODE_OWNER,
		STATUS_GRAPHICS_INSUFFICIENT_DMA_BUFFER,
		STATUS_GRAPHICS_INVALID_DISPLAY_ADAPTER,
		STATUS_GRAPHICS_ADAPTER_WAS_RESET,
		STATUS_GRAPHICS_INVALID_DRIVER_MODEL,
		STATUS_GRAPHICS_PRESENT_MODE_CHANGED,
		STATUS_GRAPHICS_PRESENT_OCCLUDED,
		STATUS_GRAPHICS_PRESENT_DENIED,
		STATUS_GRAPHICS_CANNOTCOLORCONVERT,
		STATUS_GRAPHICS_DRIVER_MISMATCH,
		STATUS_GRAPHICS_PARTIAL_DATA_POPULATED,
		STATUS_GRAPHICS_PRESENT_REDIRECTION_DISABLED,
		STATUS_GRAPHICS_PRESENT_UNOCCLUDED,
		STATUS_GRAPHICS_NO_VIDEO_MEMORY,
		STATUS_GRAPHICS_CANT_LOCK_MEMORY,
		STATUS_GRAPHICS_ALLOCATION_BUSY,
		STATUS_GRAPHICS_TOO_MANY_REFERENCES,
		STATUS_GRAPHICS_TRY_AGAIN_LATER,
		STATUS_GRAPHICS_TRY_AGAIN_NOW,
		STATUS_GRAPHICS_ALLOCATION_INVALID,
		STATUS_GRAPHICS_UNSWIZZLING_APERTURE_UNAVAILABLE,
		STATUS_GRAPHICS_UNSWIZZLING_APERTURE_UNSUPPORTED,
		STATUS_GRAPHICS_CANT_EVICT_PINNED_ALLOCATION,
		STATUS_GRAPHICS_INVALID_ALLOCATION_USAGE,
		STATUS_GRAPHICS_CANT_RENDER_LOCKED_ALLOCATION,
		STATUS_GRAPHICS_ALLOCATION_CLOSED,
		STATUS_GRAPHICS_INVALID_ALLOCATION_INSTANCE,
		STATUS_GRAPHICS_INVALID_ALLOCATION_HANDLE,
		STATUS_GRAPHICS_WRONG_ALLOCATION_DEVICE,
		STATUS_GRAPHICS_ALLOCATION_CONTENT_LOST,
		STATUS_GRAPHICS_GPU_EXCEPTION_ON_DEVICE,
		STATUS_GRAPHICS_INVALID_VIDPN_TOPOLOGY,
		STATUS_GRAPHICS_VIDPN_TOPOLOGY_NOT_SUPPORTED,
		STATUS_GRAPHICS_VIDPN_TOPOLOGY_CURRENTLY_NOT_SUPPORTED,
		STATUS_GRAPHICS_INVALID_VIDPN,
		STATUS_GRAPHICS_INVALID_VIDEO_PRESENT_SOURCE,
		STATUS_GRAPHICS_INVALID_VIDEO_PRESENT_TARGET,
		STATUS_GRAPHICS_VIDPN_MODALITY_NOT_SUPPORTED,
		STATUS_GRAPHICS_MODE_NOT_PINNED,
		STATUS_GRAPHICS_INVALID_VIDPN_SOURCEMODESET,
		STATUS_GRAPHICS_INVALID_VIDPN_TARGETMODESET,
		STATUS_GRAPHICS_INVALID_FREQUENCY,
		STATUS_GRAPHICS_INVALID_ACTIVE_REGION,
		STATUS_GRAPHICS_INVALID_TOTAL_REGION,
		STATUS_GRAPHICS_INVALID_VIDEO_PRESENT_SOURCE_MODE,
		STATUS_GRAPHICS_INVALID_VIDEO_PRESENT_TARGET_MODE,
		STATUS_GRAPHICS_PINNED_MODE_MUST_REMAIN_IN_SET,
		STATUS_GRAPHICS_PATH_ALREADY_IN_TOPOLOGY,
		STATUS_GRAPHICS_MODE_ALREADY_IN_MODESET,
		STATUS_GRAPHICS_INVALID_VIDEOPRESENTSOURCESET,
		STATUS_GRAPHICS_INVALID_VIDEOPRESENTTARGETSET,
		STATUS_GRAPHICS_SOURCE_ALREADY_IN_SET,
		STATUS_GRAPHICS_TARGET_ALREADY_IN_SET,
		STATUS_GRAPHICS_INVALID_VIDPN_PRESENT_PATH,
		STATUS_GRAPHICS_NO_RECOMMENDED_VIDPN_TOPOLOGY,
		STATUS_GRAPHICS_INVALID_MONITOR_FREQUENCYRANGESET,
		STATUS_GRAPHICS_INVALID_MONITOR_FREQUENCYRANGE,
		STATUS_GRAPHICS_FREQUENCYRANGE_NOT_IN_SET,
		STATUS_GRAPHICS_NO_PREFERRED_MODE,
		STATUS_GRAPHICS_FREQUENCYRANGE_ALREADY_IN_SET,
		STATUS_GRAPHICS_STALE_MODESET,
		STATUS_GRAPHICS_INVALID_MONITOR_SOURCEMODESET,
		STATUS_GRAPHICS_INVALID_MONITOR_SOURCE_MODE,
		STATUS_GRAPHICS_NO_RECOMMENDED_FUNCTIONAL_VIDPN,
		STATUS_GRAPHICS_MODE_ID_MUST_BE_UNIQUE,
		STATUS_GRAPHICS_EMPTY_ADAPTER_MONITOR_MODE_SUPPORT_INTERSECTION,
		STATUS_GRAPHICS_VIDEO_PRESENT_TARGETS_LESS_THAN_SOURCES,
		STATUS_GRAPHICS_PATH_NOT_IN_TOPOLOGY,
		STATUS_GRAPHICS_ADAPTER_MUST_HAVE_AT_LEAST_ONE_SOURCE,
		STATUS_GRAPHICS_ADAPTER_MUST_HAVE_AT_LEAST_ONE_TARGET,
		STATUS_GRAPHICS_INVALID_MONITORDESCRIPTORSET,
		STATUS_GRAPHICS_INVALID_MONITORDESCRIPTOR,
		STATUS_GRAPHICS_MONITORDESCRIPTOR_NOT_IN_SET,
		STATUS_GRAPHICS_MONITORDESCRIPTOR_ALREADY_IN_SET,
		STATUS_GRAPHICS_MONITORDESCRIPTOR_ID_MUST_BE_UNIQUE,
		STATUS_GRAPHICS_INVALID_VIDPN_TARGET_SUBSET_TYPE,
		STATUS_GRAPHICS_RESOURCES_NOT_RELATED,
		STATUS_GRAPHICS_SOURCE_ID_MUST_BE_UNIQUE,
		STATUS_GRAPHICS_TARGET_ID_MUST_BE_UNIQUE,
		STATUS_GRAPHICS_NO_AVAILABLE_VIDPN_TARGET,
		STATUS_GRAPHICS_MONITOR_COULD_NOT_BE_ASSOCIATED_WITH_ADAPTER,
		STATUS_GRAPHICS_NO_VIDPNMGR,
		STATUS_GRAPHICS_NO_ACTIVE_VIDPN,
		STATUS_GRAPHICS_STALE_VIDPN_TOPOLOGY,
		STATUS_GRAPHICS_MONITOR_NOT_CONNECTED,
		STATUS_GRAPHICS_SOURCE_NOT_IN_TOPOLOGY,
		STATUS_GRAPHICS_INVALID_PRIMARYSURFACE_SIZE,
		STATUS_GRAPHICS_INVALID_VISIBLEREGION_SIZE,
		STATUS_GRAPHICS_INVALID_STRIDE,
		STATUS_GRAPHICS_INVALID_PIXELFORMAT,
		STATUS_GRAPHICS_INVALID_COLORBASIS,
		STATUS_GRAPHICS_INVALID_PIXELVALUEACCESSMODE,
		STATUS_GRAPHICS_TARGET_NOT_IN_TOPOLOGY,
		STATUS_GRAPHICS_NO_DISPLAY_MODE_MANAGEMENT_SUPPORT,
		STATUS_GRAPHICS_VIDPN_SOURCE_IN_USE,
		STATUS_GRAPHICS_CANT_ACCESS_ACTIVE_VIDPN,
		STATUS_GRAPHICS_INVALID_PATH_IMPORTANCE_ORDINAL,
		STATUS_GRAPHICS_INVALID_PATH_CONTENT_GEOMETRY_TRANSFORMATION,
		STATUS_GRAPHICS_PATH_CONTENT_GEOMETRY_TRANSFORMATION_NOT_SUPPORTED,
		STATUS_GRAPHICS_INVALID_GAMMA_RAMP,
		STATUS_GRAPHICS_GAMMA_RAMP_NOT_SUPPORTED,
		STATUS_GRAPHICS_MULTISAMPLING_NOT_SUPPORTED,
		STATUS_GRAPHICS_MODE_NOT_IN_MODESET,
		STATUS_GRAPHICS_DATASET_IS_EMPTY,
		STATUS_GRAPHICS_NO_MORE_ELEMENTS_IN_DATASET,
		STATUS_GRAPHICS_INVALID_VIDPN_TOPOLOGY_RECOMMENDATION_REASON,
		STATUS_GRAPHICS_INVALID_PATH_CONTENT_TYPE,
		STATUS_GRAPHICS_INVALID_COPYPROTECTION_TYPE,
		STATUS_GRAPHICS_UNASSIGNED_MODESET_ALREADY_EXISTS,
		STATUS_GRAPHICS_PATH_CONTENT_GEOMETRY_TRANSFORMATION_NOT_PINNED,
		STATUS_GRAPHICS_INVALID_SCANLINE_ORDERING,
		STATUS_GRAPHICS_TOPOLOGY_CHANGES_NOT_ALLOWED,
		STATUS_GRAPHICS_NO_AVAILABLE_IMPORTANCE_ORDINALS,
		STATUS_GRAPHICS_INCOMPATIBLE_PRIVATE_FORMAT,
		STATUS_GRAPHICS_INVALID_MODE_PRUNING_ALGORITHM,
		STATUS_GRAPHICS_INVALID_MONITOR_CAPABILITY_ORIGIN,
		STATUS_GRAPHICS_INVALID_MONITOR_FREQUENCYRANGE_CONSTRAINT,
		STATUS_GRAPHICS_MAX_NUM_PATHS_REACHED,
		STATUS_GRAPHICS_CANCEL_VIDPN_TOPOLOGY_AUGMENTATION,
		STATUS_GRAPHICS_INVALID_CLIENT_TYPE,
		STATUS_GRAPHICS_CLIENTVIDPN_NOT_SET,
		STATUS_GRAPHICS_SPECIFIED_CHILD_ALREADY_CONNECTED,
		STATUS_GRAPHICS_CHILD_DESCRIPTOR_NOT_SUPPORTED,
		STATUS_GRAPHICS_UNKNOWN_CHILD_STATUS,
		STATUS_GRAPHICS_NOT_A_LINKED_ADAPTER,
		STATUS_GRAPHICS_LEADLINK_NOT_ENUMERATED,
		STATUS_GRAPHICS_CHAINLINKS_NOT_ENUMERATED,
		STATUS_GRAPHICS_ADAPTER_CHAIN_NOT_READY,
		STATUS_GRAPHICS_CHAINLINKS_NOT_STARTED,
		STATUS_GRAPHICS_CHAINLINKS_NOT_POWERED_ON,
		STATUS_GRAPHICS_INCONSISTENT_DEVICE_LINK_STATE,
		STATUS_GRAPHICS_LEADLINK_START_DEFERRED,
		STATUS_GRAPHICS_NOT_POST_DEVICE_DRIVER,
		STATUS_GRAPHICS_POLLING_TOO_FREQUENTLY,
		STATUS_GRAPHICS_START_DEFERRED,
		STATUS_GRAPHICS_ADAPTER_ACCESS_NOT_EXCLUDED,
		STATUS_GRAPHICS_OPM_NOT_SUPPORTED,
		STATUS_GRAPHICS_COPP_NOT_SUPPORTED,
		STATUS_GRAPHICS_UAB_NOT_SUPPORTED,
		STATUS_GRAPHICS_OPM_INVALID_ENCRYPTED_PARAMETERS,
		STATUS_GRAPHICS_OPM_NO_PROTECTED_OUTPUTS_EXIST,
		STATUS_GRAPHICS_OPM_INTERNAL_ERROR,
		STATUS_GRAPHICS_OPM_INVALID_HANDLE,
		STATUS_GRAPHICS_PVP_INVALID_CERTIFICATE_LENGTH,
		STATUS_GRAPHICS_OPM_SPANNING_MODE_ENABLED,
		STATUS_GRAPHICS_OPM_THEATER_MODE_ENABLED,
		STATUS_GRAPHICS_PVP_HFS_FAILED,
		STATUS_GRAPHICS_OPM_INVALID_SRM,
		STATUS_GRAPHICS_OPM_OUTPUT_DOES_NOT_SUPPORT_HDCP,
		STATUS_GRAPHICS_OPM_OUTPUT_DOES_NOT_SUPPORT_ACP,
		STATUS_GRAPHICS_OPM_OUTPUT_DOES_NOT_SUPPORT_CGMSA,
		STATUS_GRAPHICS_OPM_HDCP_SRM_NEVER_SET,
		STATUS_GRAPHICS_OPM_RESOLUTION_TOO_HIGH,
		STATUS_GRAPHICS_OPM_ALL_HDCP_HARDWARE_ALREADY_IN_USE,
		STATUS_GRAPHICS_OPM_PROTECTED_OUTPUT_NO_LONGER_EXISTS,
		STATUS_GRAPHICS_OPM_PROTECTED_OUTPUT_DOES_NOT_HAVE_COPP_SEMANTICS,
		STATUS_GRAPHICS_OPM_INVALID_INFORMATION_REQUEST,
		STATUS_GRAPHICS_OPM_DRIVER_INTERNAL_ERROR,
		STATUS_GRAPHICS_OPM_PROTECTED_OUTPUT_DOES_NOT_HAVE_OPM_SEMANTICS,
		STATUS_GRAPHICS_OPM_SIGNALING_NOT_SUPPORTED,
		STATUS_GRAPHICS_OPM_INVALID_CONFIGURATION_REQUEST,
		STATUS_GRAPHICS_I2C_NOT_SUPPORTED,
		STATUS_GRAPHICS_I2C_DEVICE_DOES_NOT_EXIST,
		STATUS_GRAPHICS_I2C_ERROR_TRANSMITTING_DATA,
		STATUS_GRAPHICS_I2C_ERROR_RECEIVING_DATA,
		STATUS_GRAPHICS_DDCCI_VCP_NOT_SUPPORTED,
		STATUS_GRAPHICS_DDCCI_INVALID_DATA,
		STATUS_GRAPHICS_DDCCI_MONITOR_RETURNED_INVALID_TIMING_STATUS_BYTE,
		STATUS_GRAPHICS_DDCCI_INVALID_CAPABILITIES_STRING,
		STATUS_GRAPHICS_MCA_INTERNAL_ERROR,
		STATUS_GRAPHICS_DDCCI_INVALID_MESSAGE_COMMAND,
		STATUS_GRAPHICS_DDCCI_INVALID_MESSAGE_LENGTH,
		STATUS_GRAPHICS_DDCCI_INVALID_MESSAGE_CHECKSUM,
		STATUS_GRAPHICS_INVALID_PHYSICAL_MONITOR_HANDLE,
		STATUS_GRAPHICS_MONITOR_NO_LONGER_EXISTS,
		STATUS_GRAPHICS_ONLY_CONSOLE_SESSION_SUPPORTED,
		STATUS_GRAPHICS_NO_DISPLAY_DEVICE_CORRESPONDS_TO_NAME,
		STATUS_GRAPHICS_DISPLAY_DEVICE_NOT_ATTACHED_TO_DESKTOP,
		STATUS_GRAPHICS_MIRRORING_DEVICES_NOT_SUPPORTED,
		STATUS_GRAPHICS_INVALID_POINTER,
		STATUS_GRAPHICS_NO_MONITORS_CORRESPOND_TO_DISPLAY_DEVICE,
		STATUS_GRAPHICS_PARAMETER_ARRAY_TOO_SMALL,
		STATUS_GRAPHICS_INTERNAL_ERROR,
		STATUS_GRAPHICS_SESSION_TYPE_CHANGE_IN_PROGRESS,
		STATUS_FVE_LOCKED_VOLUME,
		STATUS_FVE_NOT_ENCRYPTED,
		STATUS_FVE_BAD_INFORMATION,
		STATUS_FVE_TOO_SMALL,
		STATUS_FVE_FAILED_WRONG_FS,
		STATUS_FVE_BAD_PARTITION_SIZE,
		STATUS_FVE_FS_NOT_EXTENDED,
		STATUS_FVE_FS_MOUNTED,
		STATUS_FVE_NO_LICENSE,
		STATUS_FVE_ACTION_NOT_ALLOWED,
		STATUS_FVE_BAD_DATA,
		STATUS_FVE_VOLUME_NOT_BOUND,
		STATUS_FVE_NOT_DATA_VOLUME,
		STATUS_FVE_CONV_READ_ERROR,
		STATUS_FVE_CONV_WRITE_ERROR,
		STATUS_FVE_OVERLAPPED_UPDATE,
		STATUS_FVE_FAILED_SECTOR_SIZE,
		STATUS_FVE_FAILED_AUTHENTICATION,
		STATUS_FVE_NOT_OS_VOLUME,
		STATUS_FVE_KEYFILE_NOT_FOUND,
		STATUS_FVE_KEYFILE_INVALID,
		STATUS_FVE_KEYFILE_NO_VMK,
		STATUS_FVE_TPM_DISABLED,
		STATUS_FVE_TPM_SRK_AUTH_NOT_ZERO,
		STATUS_FVE_TPM_INVALID_PCR,
		STATUS_FVE_TPM_NO_VMK,
		STATUS_FVE_PIN_INVALID,
		STATUS_FVE_AUTH_INVALID_APPLICATION,
		STATUS_FVE_AUTH_INVALID_CONFIG,
		STATUS_FVE_DEBUGGER_ENABLED,
		STATUS_FVE_DRY_RUN_FAILED,
		STATUS_FVE_BAD_METADATA_POINTER,
		STATUS_FVE_OLD_METADATA_COPY,
		STATUS_FVE_REBOOT_REQUIRED,
		STATUS_FVE_RAW_ACCESS,
		STATUS_FVE_RAW_BLOCKED,
		STATUS_FVE_NO_AUTOUNLOCK_MASTER_KEY,
		STATUS_FVE_MOR_FAILED,
		STATUS_FVE_NO_FEATURE_LICENSE,
		STATUS_FVE_POLICY_USER_DISABLE_RDV_NOT_ALLOWED,
		STATUS_FVE_CONV_RECOVERY_FAILED,
		STATUS_FVE_VIRTUALIZED_SPACE_TOO_BIG,
		STATUS_FVE_INVALID_DATUM_TYPE,
		STATUS_FVE_VOLUME_TOO_SMALL,
		STATUS_FVE_ENH_PIN_INVALID,
		STATUS_FWP_CALLOUT_NOT_FOUND,
		STATUS_FWP_CONDITION_NOT_FOUND,
		STATUS_FWP_FILTER_NOT_FOUND,
		STATUS_FWP_LAYER_NOT_FOUND,
		STATUS_FWP_PROVIDER_NOT_FOUND,
		STATUS_FWP_PROVIDER_CONTEXT_NOT_FOUND,
		STATUS_FWP_SUBLAYER_NOT_FOUND,
		STATUS_FWP_NOT_FOUND,
		STATUS_FWP_ALREADY_EXISTS,
		STATUS_FWP_IN_USE,
		STATUS_FWP_DYNAMIC_SESSION_IN_PROGRESS,
		STATUS_FWP_WRONG_SESSION,
		STATUS_FWP_NO_TXN_IN_PROGRESS,
		STATUS_FWP_TXN_IN_PROGRESS,
		STATUS_FWP_TXN_ABORTED,
		STATUS_FWP_SESSION_ABORTED,
		STATUS_FWP_INCOMPATIBLE_TXN,
		STATUS_FWP_TIMEOUT,
		STATUS_FWP_NET_EVENTS_DISABLED,
		STATUS_FWP_INCOMPATIBLE_LAYER,
		STATUS_FWP_KM_CLIENTS_ONLY,
		STATUS_FWP_LIFETIME_MISMATCH,
		STATUS_FWP_BUILTIN_OBJECT,
		STATUS_FWP_TOO_MANY_CALLOUTS,
		STATUS_FWP_NOTIFICATION_DROPPED,
		STATUS_FWP_TRAFFIC_MISMATCH,
		STATUS_FWP_INCOMPATIBLE_SA_STATE,
		STATUS_FWP_NULL_POINTER,
		STATUS_FWP_INVALID_ENUMERATOR,
		STATUS_FWP_INVALID_FLAGS,
		STATUS_FWP_INVALID_NET_MASK,
		STATUS_FWP_INVALID_RANGE,
		STATUS_FWP_INVALID_INTERVAL,
		STATUS_FWP_ZERO_LENGTH_ARRAY,
		STATUS_FWP_NULL_DISPLAY_NAME,
		STATUS_FWP_INVALID_ACTION_TYPE,
		STATUS_FWP_INVALID_WEIGHT,
		STATUS_FWP_MATCH_TYPE_MISMATCH,
		STATUS_FWP_TYPE_MISMATCH,
		STATUS_FWP_OUT_OF_BOUNDS,
		STATUS_FWP_RESERVED,
		STATUS_FWP_DUPLICATE_CONDITION,
		STATUS_FWP_DUPLICATE_KEYMOD,
		STATUS_FWP_ACTION_INCOMPATIBLE_WITH_LAYER,
		STATUS_FWP_ACTION_INCOMPATIBLE_WITH_SUBLAYER,
		STATUS_FWP_CONTEXT_INCOMPATIBLE_WITH_LAYER,
		STATUS_FWP_CONTEXT_INCOMPATIBLE_WITH_CALLOUT,
		STATUS_FWP_INCOMPATIBLE_AUTH_METHOD,
		STATUS_FWP_INCOMPATIBLE_DH_GROUP,
		STATUS_FWP_EM_NOT_SUPPORTED,
		STATUS_FWP_NEVER_MATCH,
		STATUS_FWP_PROVIDER_CONTEXT_MISMATCH,
		STATUS_FWP_INVALID_PARAMETER,
		STATUS_FWP_TOO_MANY_SUBLAYERS,
		STATUS_FWP_CALLOUT_NOTIFICATION_FAILED,
		STATUS_FWP_INVALID_AUTH_TRANSFORM,
		STATUS_FWP_INVALID_CIPHER_TRANSFORM,
		STATUS_FWP_INCOMPATIBLE_CIPHER_TRANSFORM,
		STATUS_FWP_INVALID_TRANSFORM_COMBINATION,
		STATUS_FWP_DUPLICATE_AUTH_METHOD,
		STATUS_FWP_TCPIP_NOT_READY,
		STATUS_FWP_INJECT_HANDLE_CLOSING,
		STATUS_FWP_INJECT_HANDLE_STALE,
		STATUS_FWP_CANNOT_PEND,
		STATUS_FWP_DROP_NOICMP,
		STATUS_NDIS_CLOSING,
		STATUS_NDIS_BAD_VERSION,
		STATUS_NDIS_BAD_CHARACTERISTICS,
		STATUS_NDIS_ADAPTER_NOT_FOUND,
		STATUS_NDIS_OPEN_FAILED,
		STATUS_NDIS_DEVICE_FAILED,
		STATUS_NDIS_MULTICAST_FULL,
		STATUS_NDIS_MULTICAST_EXISTS,
		STATUS_NDIS_MULTICAST_NOT_FOUND,
		STATUS_NDIS_REQUEST_ABORTED,
		STATUS_NDIS_RESET_IN_PROGRESS,
		STATUS_NDIS_NOT_SUPPORTED,
		STATUS_NDIS_INVALID_PACKET,
		STATUS_NDIS_ADAPTER_NOT_READY,
		STATUS_NDIS_INVALID_LENGTH,
		STATUS_NDIS_INVALID_DATA,
		STATUS_NDIS_BUFFER_TOO_SHORT,
		STATUS_NDIS_INVALID_OID,
		STATUS_NDIS_ADAPTER_REMOVED,
		STATUS_NDIS_UNSUPPORTED_MEDIA,
		STATUS_NDIS_GROUP_ADDRESS_IN_USE,
		STATUS_NDIS_FILE_NOT_FOUND,
		STATUS_NDIS_ERROR_READING_FILE,
		STATUS_NDIS_ALREADY_MAPPED,
		STATUS_NDIS_RESOURCE_CONFLICT,
		STATUS_NDIS_MEDIA_DISCONNECTED,
		STATUS_NDIS_INVALID_ADDRESS,
		STATUS_NDIS_INVALID_DEVICE_REQUEST,
		STATUS_NDIS_PAUSED,
		STATUS_NDIS_INTERFACE_NOT_FOUND,
		STATUS_NDIS_UNSUPPORTED_REVISION,
		STATUS_NDIS_INVALID_PORT,
		STATUS_NDIS_INVALID_PORT_STATE,
		STATUS_NDIS_LOW_POWER_STATE,
		STATUS_NDIS_DOT11_AUTO_CONFIG_ENABLED,
		STATUS_NDIS_DOT11_MEDIA_IN_USE,
		STATUS_NDIS_DOT11_POWER_STATE_INVALID,
		STATUS_NDIS_PM_WOL_PATTERN_LIST_FULL,
		STATUS_NDIS_PM_PROTOCOL_OFFLOAD_LIST_FULL,
		STATUS_NDIS_INDICATION_REQUIRED,
		STATUS_NDIS_OFFLOAD_POLICY,
		STATUS_NDIS_OFFLOAD_CONNECTION_REJECTED,
		STATUS_NDIS_OFFLOAD_PATH_REJECTED,
		STATUS_HV_INVALID_HYPERCALL_CODE,
		STATUS_HV_INVALID_HYPERCALL_INPUT,
		STATUS_HV_INVALID_ALIGNMENT,
		STATUS_HV_INVALID_PARAMETER,
		STATUS_HV_ACCESS_DENIED,
		STATUS_HV_INVALID_PARTITION_STATE,
		STATUS_HV_OPERATION_DENIED,
		STATUS_HV_UNKNOWN_PROPERTY,
		STATUS_HV_PROPERTY_VALUE_OUT_OF_RANGE,
		STATUS_HV_INSUFFICIENT_MEMORY,
		STATUS_HV_PARTITION_TOO_DEEP,
		STATUS_HV_INVALID_PARTITION_ID,
		STATUS_HV_INVALID_VP_INDEX,
		STATUS_HV_INVALID_PORT_ID,
		STATUS_HV_INVALID_CONNECTION_ID,
		STATUS_HV_INSUFFICIENT_BUFFERS,
		STATUS_HV_NOT_ACKNOWLEDGED,
		STATUS_HV_ACKNOWLEDGED,
		STATUS_HV_INVALID_SAVE_RESTORE_STATE,
		STATUS_HV_INVALID_SYNIC_STATE,
		STATUS_HV_OBJECT_IN_USE,
		STATUS_HV_INVALID_PROXIMITY_DOMAIN_INFO,
		STATUS_HV_NO_DATA,
		STATUS_HV_INACTIVE,
		STATUS_HV_NO_RESOURCES,
		STATUS_HV_FEATURE_UNAVAILABLE,
		STATUS_HV_NOT_PRESENT,
		STATUS_VID_DUPLICATE_HANDLER,
		STATUS_VID_TOO_MANY_HANDLERS,
		STATUS_VID_QUEUE_FULL,
		STATUS_VID_HANDLER_NOT_PRESENT,
		STATUS_VID_INVALID_OBJECT_NAME,
		STATUS_VID_PARTITION_NAME_TOO_LONG,
		STATUS_VID_MESSAGE_QUEUE_NAME_TOO_LONG,
		STATUS_VID_PARTITION_ALREADY_EXISTS,
		STATUS_VID_PARTITION_DOES_NOT_EXIST,
		STATUS_VID_PARTITION_NAME_NOT_FOUND,
		STATUS_VID_MESSAGE_QUEUE_ALREADY_EXISTS,
		STATUS_VID_EXCEEDED_MBP_ENTRY_MAP_LIMIT,
		STATUS_VID_MB_STILL_REFERENCED,
		STATUS_VID_CHILD_GPA_PAGE_SET_CORRUPTED,
		STATUS_VID_INVALID_NUMA_SETTINGS,
		STATUS_VID_INVALID_NUMA_NODE_INDEX,
		STATUS_VID_NOTIFICATION_QUEUE_ALREADY_ASSOCIATED,
		STATUS_VID_INVALID_MEMORY_BLOCK_HANDLE,
		STATUS_VID_PAGE_RANGE_OVERFLOW,
		STATUS_VID_INVALID_MESSAGE_QUEUE_HANDLE,
		STATUS_VID_INVALID_GPA_RANGE_HANDLE,
		STATUS_VID_NO_MEMORY_BLOCK_NOTIFICATION_QUEUE,
		STATUS_VID_MEMORY_BLOCK_LOCK_COUNT_EXCEEDED,
		STATUS_VID_INVALID_PPM_HANDLE,
		STATUS_VID_MBPS_ARE_LOCKED,
		STATUS_VID_MESSAGE_QUEUE_CLOSED,
		STATUS_VID_VIRTUAL_PROCESSOR_LIMIT_EXCEEDED,
		STATUS_VID_STOP_PENDING,
		STATUS_VID_INVALID_PROCESSOR_STATE,
		STATUS_VID_EXCEEDED_KM_CONTEXT_COUNT_LIMIT,
		STATUS_VID_KM_INTERFACE_ALREADY_INITIALIZED,
		STATUS_VID_MB_PROPERTY_ALREADY_SET_RESET,
		STATUS_VID_MMIO_RANGE_DESTROYED,
		STATUS_VID_INVALID_CHILD_GPA_PAGE_SET,
		STATUS_VID_RESERVE_PAGE_SET_IS_BEING_USED,
		STATUS_VID_RESERVE_PAGE_SET_TOO_SMALL,
		STATUS_VID_MBP_ALREADY_LOCKED_USING_RESERVED_PAGE,
		STATUS_VID_MBP_COUNT_EXCEEDED_LIMIT,
		STATUS_VID_SAVED_STATE_CORRUPT,
		STATUS_VID_SAVED_STATE_UNRECOGNIZED_ITEM,
		STATUS_VID_SAVED_STATE_INCOMPATIBLE,
		STATUS_VID_REMOTE_NODE_PARENT_GPA_PAGES_USED,
		STATUS_IPSEC_BAD_SPI,
		STATUS_IPSEC_SA_LIFETIME_EXPIRED,
		STATUS_IPSEC_WRONG_SA,
		STATUS_IPSEC_REPLAY_CHECK_FAILED,
		STATUS_IPSEC_INVALID_PACKET,
		STATUS_IPSEC_INTEGRITY_CHECK_FAILED,
		STATUS_IPSEC_CLEAR_TEXT_DROP,
		STATUS_IPSEC_AUTH_FIREWALL_DROP,
		STATUS_IPSEC_THROTTLE_DROP,
		STATUS_IPSEC_DOSP_BLOCK,
		STATUS_IPSEC_DOSP_RECEIVED_MULTICAST,
		STATUS_IPSEC_DOSP_INVALID_PACKET,
		STATUS_IPSEC_DOSP_STATE_LOOKUP_FAILED,
		STATUS_IPSEC_DOSP_MAX_ENTRIES,
		STATUS_IPSEC_DOSP_KEYMOD_NOT_ALLOWED,
		STATUS_IPSEC_DOSP_MAX_PER_IP_RATELIMIT_QUEUES,
		STATUS_VOLMGR_INCOMPLETE_REGENERATION,
		STATUS_VOLMGR_INCOMPLETE_DISK_MIGRATION,
		STATUS_VOLMGR_DATABASE_FULL,
		STATUS_VOLMGR_DISK_CONFIGURATION_CORRUPTED,
		STATUS_VOLMGR_DISK_CONFIGURATION_NOT_IN_SYNC,
		STATUS_VOLMGR_PACK_CONFIG_UPDATE_FAILED,
		STATUS_VOLMGR_DISK_CONTAINS_NON_SIMPLE_VOLUME,
		STATUS_VOLMGR_DISK_DUPLICATE,
		STATUS_VOLMGR_DISK_DYNAMIC,
		STATUS_VOLMGR_DISK_ID_INVALID,
		STATUS_VOLMGR_DISK_INVALID,
		STATUS_VOLMGR_DISK_LAST_VOTER,
		STATUS_VOLMGR_DISK_LAYOUT_INVALID,
		STATUS_VOLMGR_DISK_LAYOUT_NON_BASIC_BETWEEN_BASIC_PARTITIONS,
		STATUS_VOLMGR_DISK_LAYOUT_NOT_CYLINDER_ALIGNED,
		STATUS_VOLMGR_DISK_LAYOUT_PARTITIONS_TOO_SMALL,
		STATUS_VOLMGR_DISK_LAYOUT_PRIMARY_BETWEEN_LOGICAL_PARTITIONS,
		STATUS_VOLMGR_DISK_LAYOUT_TOO_MANY_PARTITIONS,
		STATUS_VOLMGR_DISK_MISSING,
		STATUS_VOLMGR_DISK_NOT_EMPTY,
		STATUS_VOLMGR_DISK_NOT_ENOUGH_SPACE,
		STATUS_VOLMGR_DISK_REVECTORING_FAILED,
		STATUS_VOLMGR_DISK_SECTOR_SIZE_INVALID,
		STATUS_VOLMGR_DISK_SET_NOT_CONTAINED,
		STATUS_VOLMGR_DISK_USED_BY_MULTIPLE_MEMBERS,
		STATUS_VOLMGR_DISK_USED_BY_MULTIPLE_PLEXES,
		STATUS_VOLMGR_DYNAMIC_DISK_NOT_SUPPORTED,
		STATUS_VOLMGR_EXTENT_ALREADY_USED,
		STATUS_VOLMGR_EXTENT_NOT_CONTIGUOUS,
		STATUS_VOLMGR_EXTENT_NOT_IN_PUBLIC_REGION,
		STATUS_VOLMGR_EXTENT_NOT_SECTOR_ALIGNED,
		STATUS_VOLMGR_EXTENT_OVERLAPS_EBR_PARTITION,
		STATUS_VOLMGR_EXTENT_VOLUME_LENGTHS_DO_NOT_MATCH,
		STATUS_VOLMGR_FAULT_TOLERANT_NOT_SUPPORTED,
		STATUS_VOLMGR_INTERLEAVE_LENGTH_INVALID,
		STATUS_VOLMGR_MAXIMUM_REGISTERED_USERS,
		STATUS_VOLMGR_MEMBER_IN_SYNC,
		STATUS_VOLMGR_MEMBER_INDEX_DUPLICATE,
		STATUS_VOLMGR_MEMBER_INDEX_INVALID,
		STATUS_VOLMGR_MEMBER_MISSING,
		STATUS_VOLMGR_MEMBER_NOT_DETACHED,
		STATUS_VOLMGR_MEMBER_REGENERATING,
		STATUS_VOLMGR_ALL_DISKS_FAILED,
		STATUS_VOLMGR_NO_REGISTERED_USERS,
		STATUS_VOLMGR_NO_SUCH_USER,
		STATUS_VOLMGR_NOTIFICATION_RESET,
		STATUS_VOLMGR_NUMBER_OF_MEMBERS_INVALID,
		STATUS_VOLMGR_NUMBER_OF_PLEXES_INVALID,
		STATUS_VOLMGR_PACK_DUPLICATE,
		STATUS_VOLMGR_PACK_ID_INVALID,
		STATUS_VOLMGR_PACK_INVALID,
		STATUS_VOLMGR_PACK_NAME_INVALID,
		STATUS_VOLMGR_PACK_OFFLINE,
		STATUS_VOLMGR_PACK_HAS_QUORUM,
		STATUS_VOLMGR_PACK_WITHOUT_QUORUM,
		STATUS_VOLMGR_PARTITION_STYLE_INVALID,
		STATUS_VOLMGR_PARTITION_UPDATE_FAILED,
		STATUS_VOLMGR_PLEX_IN_SYNC,
		STATUS_VOLMGR_PLEX_INDEX_DUPLICATE,
		STATUS_VOLMGR_PLEX_INDEX_INVALID,
		STATUS_VOLMGR_PLEX_LAST_ACTIVE,
		STATUS_VOLMGR_PLEX_MISSING,
		STATUS_VOLMGR_PLEX_REGENERATING,
		STATUS_VOLMGR_PLEX_TYPE_INVALID,
		STATUS_VOLMGR_PLEX_NOT_RAID5,
		STATUS_VOLMGR_PLEX_NOT_SIMPLE,
		STATUS_VOLMGR_STRUCTURE_SIZE_INVALID,
		STATUS_VOLMGR_TOO_MANY_NOTIFICATION_REQUESTS,
		STATUS_VOLMGR_TRANSACTION_IN_PROGRESS,
		STATUS_VOLMGR_UNEXPECTED_DISK_LAYOUT_CHANGE,
		STATUS_VOLMGR_VOLUME_CONTAINS_MISSING_DISK,
		STATUS_VOLMGR_VOLUME_ID_INVALID,
		STATUS_VOLMGR_VOLUME_LENGTH_INVALID,
		STATUS_VOLMGR_VOLUME_LENGTH_NOT_SECTOR_SIZE_MULTIPLE,
		STATUS_VOLMGR_VOLUME_NOT_MIRRORED,
		STATUS_VOLMGR_VOLUME_NOT_RETAINED,
		STATUS_VOLMGR_VOLUME_OFFLINE,
		STATUS_VOLMGR_VOLUME_RETAINED,
		STATUS_VOLMGR_NUMBER_OF_EXTENTS_INVALID,
		STATUS_VOLMGR_DIFFERENT_SECTOR_SIZE,
		STATUS_VOLMGR_BAD_BOOT_DISK,
		STATUS_VOLMGR_PACK_CONFIG_OFFLINE,
		STATUS_VOLMGR_PACK_CONFIG_ONLINE,
		STATUS_VOLMGR_NOT_PRIMARY_PACK,
		STATUS_VOLMGR_PACK_LOG_UPDATE_FAILED,
		STATUS_VOLMGR_NUMBER_OF_DISKS_IN_PLEX_INVALID,
		STATUS_VOLMGR_NUMBER_OF_DISKS_IN_MEMBER_INVALID,
		STATUS_VOLMGR_VOLUME_MIRRORED,
		STATUS_VOLMGR_PLEX_NOT_SIMPLE_SPANNED,
		STATUS_VOLMGR_NO_VALID_LOG_COPIES,
		STATUS_VOLMGR_PRIMARY_PACK_PRESENT,
		STATUS_VOLMGR_NUMBER_OF_DISKS_INVALID,
		STATUS_VOLMGR_MIRROR_NOT_SUPPORTED,
		STATUS_VOLMGR_RAID5_NOT_SUPPORTED,
		STATUS_BCD_NOT_ALL_ENTRIES_IMPORTED,
		STATUS_BCD_TOO_MANY_ELEMENTS,
		STATUS_BCD_NOT_ALL_ENTRIES_SYNCHRONIZED,
		STATUS_VHD_DRIVE_FOOTER_MISSING,
		STATUS_VHD_DRIVE_FOOTER_CHECKSUM_MISMATCH,
		STATUS_VHD_DRIVE_FOOTER_CORRUPT,
		STATUS_VHD_FORMAT_UNKNOWN,
		STATUS_VHD_FORMAT_UNSUPPORTED_VERSION,
		STATUS_VHD_SPARSE_HEADER_CHECKSUM_MISMATCH,
		STATUS_VHD_SPARSE_HEADER_UNSUPPORTED_VERSION,
		STATUS_VHD_SPARSE_HEADER_CORRUPT,
		STATUS_VHD_BLOCK_ALLOCATION_FAILURE,
		STATUS_VHD_BLOCK_ALLOCATION_TABLE_CORRUPT,
		STATUS_VHD_INVALID_BLOCK_SIZE,
		STATUS_VHD_BITMAP_MISMATCH,
		STATUS_VHD_PARENT_VHD_NOT_FOUND,
		STATUS_VHD_CHILD_PARENT_ID_MISMATCH,
		STATUS_VHD_CHILD_PARENT_TIMESTAMP_MISMATCH,
		STATUS_VHD_METADATA_READ_FAILURE,
		STATUS_VHD_METADATA_WRITE_FAILURE,
		STATUS_VHD_INVALID_SIZE,
		STATUS_VHD_INVALID_FILE_SIZE,
		STATUS_VIRTDISK_PROVIDER_NOT_FOUND,
		STATUS_VIRTDISK_NOT_VIRTUAL_DISK,
		STATUS_VHD_PARENT_VHD_ACCESS_DENIED,
		STATUS_VHD_CHILD_PARENT_SIZE_MISMATCH,
		STATUS_VHD_DIFFERENCING_CHAIN_CYCLE_DETECTED,
		STATUS_VHD_DIFFERENCING_CHAIN_ERROR_IN_PARENT,
		STATUS_VIRTUAL_DISK_LIMITATION,
		STATUS_VHD_INVALID_TYPE,
		STATUS_VHD_INVALID_STATE,
		STATUS_VIRTDISK_UNSUPPORTED_DISK_SECTOR_SIZE,
		STATUS_QUERY_STORAGE_ERROR,
		STATUS_DIS_NOT_PRESENT,
		STATUS_DIS_ATTRIBUTE_NOT_FOUND,
		STATUS_DIS_UNRECOGNIZED_ATTRIBUTE,
		STATUS_DIS_PARTIAL_DATA,
	}
}
