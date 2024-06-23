package constants

type ErrorsKind int

const (
	ERROR_SUCCESS                                                         ErrorsKind = 0
	ERROR_INVALID_FUNCTION                                                           = 1
	ERROR_FILE_NOT_FOUND                                                             = 2
	ERROR_PATH_NOT_FOUND                                                             = 3
	ERROR_TOO_MANY_OPEN_FILES                                                        = 4
	ERROR_ACCESS_DENIED                                                              = 5
	ERROR_INVALID_HANDLE                                                             = 6
	ERROR_ARENA_TRASHED                                                              = 7
	ERROR_NOT_ENOUGH_MEMORY                                                          = 8
	ERROR_INVALID_BLOCK                                                              = 9
	ERROR_BAD_ENVIRONMENT                                                            = 10
	ERROR_BAD_FORMAT                                                                 = 11
	ERROR_INVALID_ACCESS                                                             = 12
	ERROR_INVALID_DATA                                                               = 13
	ERROR_OUTOFMEMORY                                                                = 14
	ERROR_INVALID_DRIVE                                                              = 15
	ERROR_CURRENT_DIRECTORY                                                          = 16
	ERROR_NOT_SAME_DEVICE                                                            = 17
	ERROR_NO_MORE_FILES                                                              = 18
	ERROR_WRITE_PROTECT                                                              = 19
	ERROR_BAD_UNIT                                                                   = 20
	ERROR_NOT_READY                                                                  = 21
	ERROR_BAD_COMMAND                                                                = 22
	ERROR_CRC                                                                        = 23
	ERROR_BAD_LENGTH                                                                 = 24
	ERROR_SEEK                                                                       = 25
	ERROR_NOT_DOS_DISK                                                               = 26
	ERROR_SECTOR_NOT_FOUND                                                           = 27
	ERROR_OUT_OF_PAPER                                                               = 28
	ERROR_WRITE_FAULT                                                                = 29
	ERROR_READ_FAULT                                                                 = 30
	ERROR_GEN_FAILURE                                                                = 31
	ERROR_SHARING_VIOLATION                                                          = 32
	ERROR_LOCK_VIOLATION                                                             = 33
	ERROR_WRONG_DISK                                                                 = 34
	ERROR_SHARING_BUFFER_EXCEEDED                                                    = 36
	ERROR_HANDLE_EOF                                                                 = 38
	ERROR_HANDLE_DISK_FULL                                                           = 39
	ERROR_NOT_SUPPORTED                                                              = 50
	ERROR_REM_NOT_LIST                                                               = 51
	ERROR_DUP_NAME                                                                   = 52
	ERROR_BAD_NETPATH                                                                = 53
	ERROR_NETWORK_BUSY                                                               = 54
	ERROR_DEV_NOT_EXIST                                                              = 55
	ERROR_TOO_MANY_CMDS                                                              = 56
	ERROR_ADAP_HDW_ERR                                                               = 57
	ERROR_BAD_NET_RESP                                                               = 58
	ERROR_UNEXP_NET_ERR                                                              = 59
	ERROR_BAD_REM_ADAP                                                               = 60
	ERROR_PRINTQ_FULL                                                                = 61
	ERROR_NO_SPOOL_SPACE                                                             = 62
	ERROR_PRINT_CANCELLED                                                            = 63
	ERROR_NETNAME_DELETED                                                            = 64
	ERROR_NETWORK_ACCESS_DENIED                                                      = 65
	ERROR_BAD_DEV_TYPE                                                               = 66
	ERROR_BAD_NET_NAME                                                               = 67
	ERROR_TOO_MANY_NAMES                                                             = 68
	ERROR_TOO_MANY_SESS                                                              = 69
	ERROR_SHARING_PAUSED                                                             = 70
	ERROR_REQ_NOT_ACCEP                                                              = 71
	ERROR_REDIR_PAUSED                                                               = 72
	ERROR_FILE_EXISTS                                                                = 80
	ERROR_CANNOT_MAKE                                                                = 82
	ERROR_FAIL_I24                                                                   = 83
	ERROR_OUT_OF_STRUCTURES                                                          = 84
	ERROR_ALREADY_ASSIGNED                                                           = 85
	ERROR_INVALID_PASSWORD                                                           = 86
	ERROR_INVALID_PARAMETER                                                          = 87
	ERROR_NET_WRITE_FAULT                                                            = 88
	ERROR_NO_PROC_SLOTS                                                              = 89
	ERROR_TOO_MANY_SEMAPHORES                                                        = 100
	ERROR_EXCL_SEM_ALREADY_OWNED                                                     = 101
	ERROR_SEM_IS_SET                                                                 = 102
	ERROR_TOO_MANY_SEM_REQUESTS                                                      = 103
	ERROR_INVALID_AT_INTERRUPT_TIME                                                  = 104
	ERROR_SEM_OWNER_DIED                                                             = 105
	ERROR_SEM_USER_LIMIT                                                             = 106
	ERROR_DISK_CHANGE                                                                = 107
	ERROR_DRIVE_LOCKED                                                               = 108
	ERROR_BROKEN_PIPE                                                                = 109
	ERROR_OPEN_FAILED                                                                = 110
	ERROR_BUFFER_OVERFLOW                                                            = 111
	ERROR_DISK_FULL                                                                  = 112
	ERROR_NO_MORE_SEARCH_HANDLES                                                     = 113
	ERROR_INVALID_TARGET_HANDLE                                                      = 114
	ERROR_INVALID_CATEGORY                                                           = 117
	ERROR_INVALID_VERIFY_SWITCH                                                      = 118
	ERROR_BAD_DRIVER_LEVEL                                                           = 119
	ERROR_CALL_NOT_IMPLEMENTED                                                       = 120
	ERROR_SEM_TIMEOUT                                                                = 121
	ERROR_INSUFFICIENT_BUFFER                                                        = 122
	ERROR_INVALID_NAME                                                               = 123
	ERROR_INVALID_LEVEL                                                              = 124
	ERROR_NO_VOLUME_LABEL                                                            = 125
	ERROR_MOD_NOT_FOUND                                                              = 126
	ERROR_PROC_NOT_FOUND                                                             = 127
	ERROR_WAIT_NO_CHILDREN                                                           = 128
	ERROR_CHILD_NOT_COMPLETE                                                         = 129
	ERROR_DIRECT_ACCESS_HANDLE                                                       = 130
	ERROR_NEGATIVE_SEEK                                                              = 131
	ERROR_SEEK_ON_DEVICE                                                             = 132
	ERROR_IS_JOIN_TARGET                                                             = 133
	ERROR_IS_JOINED                                                                  = 134
	ERROR_IS_SUBSTED                                                                 = 135
	ERROR_NOT_JOINED                                                                 = 136
	ERROR_NOT_SUBSTED                                                                = 137
	ERROR_JOIN_TO_JOIN                                                               = 138
	ERROR_SUBST_TO_SUBST                                                             = 139
	ERROR_JOIN_TO_SUBST                                                              = 140
	ERROR_SUBST_TO_JOIN                                                              = 141
	ERROR_BUSY_DRIVE                                                                 = 142
	ERROR_SAME_DRIVE                                                                 = 143
	ERROR_DIR_NOT_ROOT                                                               = 144
	ERROR_DIR_NOT_EMPTY                                                              = 145
	ERROR_IS_SUBST_PATH                                                              = 146
	ERROR_IS_JOIN_PATH                                                               = 147
	ERROR_PATH_BUSY                                                                  = 148
	ERROR_IS_SUBST_TARGET                                                            = 149
	ERROR_SYSTEM_TRACE                                                               = 150
	ERROR_INVALID_EVENT_COUNT                                                        = 151
	ERROR_TOO_MANY_MUXWAITERS                                                        = 152
	ERROR_INVALID_LIST_FORMAT                                                        = 153
	ERROR_LABEL_TOO_LONG                                                             = 154
	ERROR_TOO_MANY_TCBS                                                              = 155
	ERROR_SIGNAL_REFUSED                                                             = 156
	ERROR_DISCARDED                                                                  = 157
	ERROR_NOT_LOCKED                                                                 = 158
	ERROR_BAD_THREADID_ADDR                                                          = 159
	ERROR_BAD_ARGUMENTS                                                              = 160
	ERROR_BAD_PATHNAME                                                               = 161
	ERROR_SIGNAL_PENDING                                                             = 162
	ERROR_MAX_THRDS_REACHED                                                          = 164
	ERROR_LOCK_FAILED                                                                = 167
	ERROR_BUSY                                                                       = 170
	ERROR_CANCEL_VIOLATION                                                           = 173
	ERROR_ATOMIC_LOCKS_NOT_SUPPORTED                                                 = 174
	ERROR_INVALID_SEGMENT_NUMBER                                                     = 180
	ERROR_INVALID_ORDINAL                                                            = 182
	ERROR_ALREADY_EXISTS                                                             = 183
	ERROR_INVALID_FLAG_NUMBER                                                        = 186
	ERROR_SEM_NOT_FOUND                                                              = 187
	ERROR_INVALID_STARTING_CODESEG                                                   = 188
	ERROR_INVALID_STACKSEG                                                           = 189
	ERROR_INVALID_MODULETYPE                                                         = 190
	ERROR_INVALID_EXE_SIGNATURE                                                      = 191
	ERROR_EXE_MARKED_INVALID                                                         = 192
	ERROR_BAD_EXE_FORMAT                                                             = 193
	ERROR_ITERATED_DATA_EXCEEDS_64k                                                  = 194
	ERROR_INVALID_MINALLOCSIZE                                                       = 195
	ERROR_DYNLINK_FROM_INVALID_RING                                                  = 196
	ERROR_IOPL_NOT_ENABLED                                                           = 197
	ERROR_INVALID_SEGDPL                                                             = 198
	ERROR_AUTODATASEG_EXCEEDS_64k                                                    = 199
	ERROR_RING2SEG_MUST_BE_MOVABLE                                                   = 200
	ERROR_RELOC_CHAIN_XEEDS_SEGLIM                                                   = 201
	ERROR_INFLOOP_IN_RELOC_CHAIN                                                     = 202
	ERROR_ENVVAR_NOT_FOUND                                                           = 203
	ERROR_NO_SIGNAL_SENT                                                             = 205
	ERROR_FILENAME_EXCED_RANGE                                                       = 206
	ERROR_RING2_STACK_IN_USE                                                         = 207
	ERROR_META_EXPANSION_TOO_LONG                                                    = 208
	ERROR_INVALID_SIGNAL_NUMBER                                                      = 209
	ERROR_THREAD_1_INACTIVE                                                          = 210
	ERROR_LOCKED                                                                     = 212
	ERROR_TOO_MANY_MODULES                                                           = 214
	ERROR_NESTING_NOT_ALLOWED                                                        = 215
	ERROR_EXE_MACHINE_TYPE_MISMATCH                                                  = 216
	ERROR_EXE_CANNOT_MODIFY_SIGNED_BINARY                                            = 217
	ERROR_EXE_CANNOT_MODIFY_STRONG_SIGNED_BINARY                                     = 218
	ERROR_FILE_CHECKED_OUT                                                           = 220
	ERROR_CHECKOUT_REQUIRED                                                          = 221
	ERROR_BAD_FILE_TYPE                                                              = 222
	ERROR_FILE_TOO_LARGE                                                             = 223
	ERROR_FORMS_AUTH_REQUIRED                                                        = 224
	ERROR_VIRUS_INFECTED                                                             = 225
	ERROR_VIRUS_DELETED                                                              = 226
	ERROR_PIPE_LOCAL                                                                 = 229
	ERROR_BAD_PIPE                                                                   = 230
	ERROR_PIPE_BUSY                                                                  = 231
	ERROR_NO_DATA                                                                    = 232
	ERROR_PIPE_NOT_CONNECTED                                                         = 233
	ERROR_MORE_DATA                                                                  = 234
	ERROR_VC_DISCONNECTED                                                            = 240
	ERROR_INVALID_EA_NAME                                                            = 254
	ERROR_EA_LIST_INCONSISTENT                                                       = 255
	WAIT_TIMEOUT                                                                     = 258
	ERROR_NO_MORE_ITEMS                                                              = 259
	ERROR_CANNOT_COPY                                                                = 266
	ERROR_DIRECTORY                                                                  = 267
	ERROR_EAS_DIDNT_FIT                                                              = 275
	ERROR_EA_FILE_CORRUPT                                                            = 276
	ERROR_EA_TABLE_FULL                                                              = 277
	ERROR_INVALID_EA_HANDLE                                                          = 278
	ERROR_EAS_NOT_SUPPORTED                                                          = 282
	ERROR_NOT_OWNER                                                                  = 288
	ERROR_TOO_MANY_POSTS                                                             = 298
	ERROR_PARTIAL_COPY                                                               = 299
	ERROR_OPLOCK_NOT_GRANTED                                                         = 300
	ERROR_INVALID_OPLOCK_PROTOCOL                                                    = 301
	ERROR_DISK_TOO_FRAGMENTED                                                        = 302
	ERROR_DELETE_PENDING                                                             = 303
	ERROR_INCOMPATIBLE_WITH_GLOBAL_SHORT_NAME_REGISTRY_SETTING                       = 304
	ERROR_SHORT_NAMES_NOT_ENABLED_ON_VOLUME                                          = 305
	ERROR_SECURITY_STREAM_IS_INCONSISTENT                                            = 306
	ERROR_INVALID_LOCK_RANGE                                                         = 307
	ERROR_IMAGE_SUBSYSTEM_NOT_PRESENT                                                = 308
	ERROR_NOTIFICATION_GUID_ALREADY_DEFINED                                          = 309
	ERROR_MR_MID_NOT_FOUND                                                           = 317
	ERROR_SCOPE_NOT_FOUND                                                            = 318
	ERROR_FAIL_NOACTION_REBOOT                                                       = 350
	ERROR_FAIL_SHUTDOWN                                                              = 351
	ERROR_FAIL_RESTART                                                               = 352
	ERROR_MAX_SESSIONS_REACHED                                                       = 353
	ERROR_THREAD_MODE_ALREADY_BACKGROUND                                             = 400
	ERROR_THREAD_MODE_NOT_BACKGROUND                                                 = 401
	ERROR_PROCESS_MODE_ALREADY_BACKGROUND                                            = 402
	ERROR_PROCESS_MODE_NOT_BACKGROUND                                                = 403
	ERROR_INVALID_ADDRESS                                                            = 487
	ERROR_USER_PROFILE_LOAD                                                          = 500
	ERROR_ARITHMETIC_OVERFLOW                                                        = 534
	ERROR_PIPE_CONNECTED                                                             = 535
	ERROR_PIPE_LISTENING                                                             = 536
	ERROR_VERIFIER_STOP                                                              = 537
	ERROR_ABIOS_ERROR                                                                = 538
	ERROR_WX86_WARNING                                                               = 539
	ERROR_WX86_ERROR                                                                 = 540
	ERROR_TIMER_NOT_CANCELED                                                         = 541
	ERROR_UNWIND                                                                     = 542
	ERROR_BAD_STACK                                                                  = 543
	ERROR_INVALID_UNWIND_TARGET                                                      = 544
	ERROR_INVALID_PORT_ATTRIBUTES                                                    = 545
	ERROR_PORT_MESSAGE_TOO_LONG                                                      = 546
	ERROR_INVALID_QUOTA_LOWER                                                        = 547
	ERROR_DEVICE_ALREADY_ATTACHED                                                    = 548
	ERROR_INSTRUCTION_MISALIGNMENT                                                   = 549
	ERROR_PROFILING_NOT_STARTED                                                      = 550
	ERROR_PROFILING_NOT_STOPPED                                                      = 551
	ERROR_COULD_NOT_INTERPRET                                                        = 552
	ERROR_PROFILING_AT_LIMIT                                                         = 553
	ERROR_CANT_WAIT                                                                  = 554
	ERROR_CANT_TERMINATE_SELF                                                        = 555
	ERROR_UNEXPECTED_MM_CREATE_ERR                                                   = 556
	ERROR_UNEXPECTED_MM_MAP_ERROR                                                    = 557
	ERROR_UNEXPECTED_MM_EXTEND_ERR                                                   = 558
	ERROR_BAD_FUNCTION_TABLE                                                         = 559
	ERROR_NO_GUID_TRANSLATION                                                        = 560
	ERROR_INVALID_LDT_SIZE                                                           = 561
	ERROR_INVALID_LDT_OFFSET                                                         = 563
	ERROR_INVALID_LDT_DESCRIPTOR                                                     = 564
	ERROR_TOO_MANY_THREADS                                                           = 565
	ERROR_THREAD_NOT_IN_PROCESS                                                      = 566
	ERROR_PAGEFILE_QUOTA_EXCEEDED                                                    = 567
	ERROR_LOGON_SERVER_CONFLICT                                                      = 568
	ERROR_SYNCHRONIZATION_REQUIRED                                                   = 569
	ERROR_NET_OPEN_FAILED                                                            = 570
	ERROR_IO_PRIVILEGE_FAILED                                                        = 571
	ERROR_CONTROL_C_EXIT                                                             = 572
	ERROR_MISSING_SYSTEMFILE                                                         = 573
	ERROR_UNHANDLED_EXCEPTION                                                        = 574
	ERROR_APP_INIT_FAILURE                                                           = 575
	ERROR_PAGEFILE_CREATE_FAILED                                                     = 576
	ERROR_INVALID_IMAGE_HASH                                                         = 577
	ERROR_NO_PAGEFILE                                                                = 578
	ERROR_ILLEGAL_FLOAT_CONTEXT                                                      = 579
	ERROR_NO_EVENT_PAIR                                                              = 580
	ERROR_DOMAIN_CTRLR_CONFIG_ERROR                                                  = 581
	ERROR_ILLEGAL_CHARACTER                                                          = 582
	ERROR_UNDEFINED_CHARACTER                                                        = 583
	ERROR_FLOPPY_VOLUME                                                              = 584
	ERROR_BIOS_FAILED_TO_CONNECT_INTERRUPT                                           = 585
	ERROR_BACKUP_CONTROLLER                                                          = 586
	ERROR_MUTANT_LIMIT_EXCEEDED                                                      = 587
	ERROR_FS_DRIVER_REQUIRED                                                         = 588
	ERROR_CANNOT_LOAD_REGISTRY_FILE                                                  = 589
	ERROR_DEBUG_ATTACH_FAILED                                                        = 590
	ERROR_SYSTEM_PROCESS_TERMINATED                                                  = 591
	ERROR_DATA_NOT_ACCEPTED                                                          = 592
	ERROR_VDM_HARD_ERROR                                                             = 593
	ERROR_DRIVER_CANCEL_TIMEOUT                                                      = 594
	ERROR_REPLY_MESSAGE_MISMATCH                                                     = 595
	ERROR_LOST_WRITEBEHIND_DATA                                                      = 596
	ERROR_CLIENT_SERVER_PARAMETERS_INVALID                                           = 597
	ERROR_NOT_TINY_STREAM                                                            = 598
	ERROR_STACK_OVERFLOW_READ                                                        = 599
	ERROR_CONVERT_TO_LARGE                                                           = 600
	ERROR_FOUND_OUT_OF_SCOPE                                                         = 601
	ERROR_ALLOCATE_BUCKET                                                            = 602
	ERROR_MARSHALL_OVERFLOW                                                          = 603
	ERROR_INVALID_VARIANT                                                            = 604
	ERROR_BAD_COMPRESSION_BUFFER                                                     = 605
	ERROR_AUDIT_FAILED                                                               = 606
	ERROR_TIMER_RESOLUTION_NOT_SET                                                   = 607
	ERROR_INSUFFICIENT_LOGON_INFO                                                    = 608
	ERROR_BAD_DLL_ENTRYPOINT                                                         = 609
	ERROR_BAD_SERVICE_ENTRYPOINT                                                     = 610
	ERROR_IP_ADDRESS_CONFLICT1                                                       = 611
	ERROR_IP_ADDRESS_CONFLICT2                                                       = 612
	ERROR_REGISTRY_QUOTA_LIMIT                                                       = 613
	ERROR_NO_CALLBACK_ACTIVE                                                         = 614
	ERROR_PWD_TOO_SHORT                                                              = 615
	ERROR_PWD_TOO_RECENT                                                             = 616
	ERROR_PWD_HISTORY_CONFLICT                                                       = 617
	ERROR_UNSUPPORTED_COMPRESSION                                                    = 618
	ERROR_INVALID_HW_PROFILE                                                         = 619
	ERROR_INVALID_PLUGPLAY_DEVICE_PATH                                               = 620
	ERROR_QUOTA_LIST_INCONSISTENT                                                    = 621
	ERROR_EVALUATION_EXPIRATION                                                      = 622
	ERROR_ILLEGAL_DLL_RELOCATION                                                     = 623
	ERROR_DLL_INIT_FAILED_LOGOFF                                                     = 624
	ERROR_VALIDATE_CONTINUE                                                          = 625
	ERROR_NO_MORE_MATCHES                                                            = 626
	ERROR_RANGE_LIST_CONFLICT                                                        = 627
	ERROR_SERVER_SID_MISMATCH                                                        = 628
	ERROR_CANT_ENABLE_DENY_ONLY                                                      = 629
	ERROR_FLOAT_MULTIPLE_FAULTS                                                      = 630
	ERROR_FLOAT_MULTIPLE_TRAPS                                                       = 631
	ERROR_NOINTERFACE                                                                = 632
	ERROR_DRIVER_FAILED_SLEEP                                                        = 633
	ERROR_CORRUPT_SYSTEM_FILE                                                        = 634
	ERROR_COMMITMENT_MINIMUM                                                         = 635
	ERROR_PNP_RESTART_ENUMERATION                                                    = 636
	ERROR_SYSTEM_IMAGE_BAD_SIGNATURE                                                 = 637
	ERROR_PNP_REBOOT_REQUIRED                                                        = 638
	ERROR_INSUFFICIENT_POWER                                                         = 639
	ERROR_MULTIPLE_FAULT_VIOLATION                                                   = 640
	ERROR_SYSTEM_SHUTDOWN                                                            = 641
	ERROR_PORT_NOT_SET                                                               = 642
	ERROR_DS_VERSION_CHECK_FAILURE                                                   = 643
	ERROR_RANGE_NOT_FOUND                                                            = 644
	ERROR_NOT_SAFE_MODE_DRIVER                                                       = 646
	ERROR_FAILED_DRIVER_ENTRY                                                        = 647
	ERROR_DEVICE_ENUMERATION_ERROR                                                   = 648
	ERROR_MOUNT_POINT_NOT_RESOLVED                                                   = 649
	ERROR_INVALID_DEVICE_OBJECT_PARAMETER                                            = 650
	ERROR_MCA_OCCURED                                                                = 651
	ERROR_DRIVER_DATABASE_ERROR                                                      = 652
	ERROR_SYSTEM_HIVE_TOO_LARGE                                                      = 653
	ERROR_DRIVER_FAILED_PRIOR_UNLOAD                                                 = 654
	ERROR_VOLSNAP_PREPARE_HIBERNATE                                                  = 655
	ERROR_HIBERNATION_FAILURE                                                        = 656
	ERROR_FILE_SYSTEM_LIMITATION                                                     = 665
	ERROR_ASSERTION_FAILURE                                                          = 668
	ERROR_ACPI_ERROR                                                                 = 669
	ERROR_WOW_ASSERTION                                                              = 670
	ERROR_PNP_BAD_MPS_TABLE                                                          = 671
	ERROR_PNP_TRANSLATION_FAILED                                                     = 672
	ERROR_PNP_IRQ_TRANSLATION_FAILED                                                 = 673
	ERROR_PNP_INVALID_ID                                                             = 674
	ERROR_WAKE_SYSTEM_DEBUGGER                                                       = 675
	ERROR_HANDLES_CLOSED                                                             = 676
	ERROR_EXTRANEOUS_INFORMATION                                                     = 677
	ERROR_RXACT_COMMIT_NECESSARY                                                     = 678
	ERROR_MEDIA_CHECK                                                                = 679
	ERROR_GUID_SUBSTITUTION_MADE                                                     = 680
	ERROR_STOPPED_ON_SYMLINK                                                         = 681
	ERROR_LONGJUMP                                                                   = 682
	ERROR_PLUGPLAY_QUERY_VETOED                                                      = 683
	ERROR_UNWIND_CONSOLIDATE                                                         = 684
	ERROR_REGISTRY_HIVE_RECOVERED                                                    = 685
	ERROR_DLL_MIGHT_BE_INSECURE                                                      = 686
	ERROR_DLL_MIGHT_BE_INCOMPATIBLE                                                  = 687
	ERROR_DBG_EXCEPTION_NOT_HANDLED                                                  = 688
	ERROR_DBG_REPLY_LATER                                                            = 689
	ERROR_DBG_UNABLE_TO_PROVIDE_HANDLE                                               = 690
	ERROR_DBG_TERMINATE_THREAD                                                       = 691
	ERROR_DBG_TERMINATE_PROCESS                                                      = 692
	ERROR_DBG_CONTROL_C                                                              = 693
	ERROR_DBG_PRINTEXCEPTION_C                                                       = 694
	ERROR_DBG_RIPEXCEPTION                                                           = 695
	ERROR_DBG_CONTROL_BREAK                                                          = 696
	ERROR_DBG_COMMAND_EXCEPTION                                                      = 697
	ERROR_OBJECT_NAME_EXISTS                                                         = 698
	ERROR_THREAD_WAS_SUSPENDED                                                       = 699
	ERROR_IMAGE_NOT_AT_BASE                                                          = 700
	ERROR_RXACT_STATE_CREATED                                                        = 701
	ERROR_SEGMENT_NOTIFICATION                                                       = 702
	ERROR_BAD_CURRENT_DIRECTORY                                                      = 703
	ERROR_FT_READ_RECOVERY_FROM_BACKUP                                               = 704
	ERROR_FT_WRITE_RECOVERY                                                          = 705
	ERROR_IMAGE_MACHINE_TYPE_MISMATCH                                                = 706
	ERROR_RECEIVE_PARTIAL                                                            = 707
	ERROR_RECEIVE_EXPEDITED                                                          = 708
	ERROR_RECEIVE_PARTIAL_EXPEDITED                                                  = 709
	ERROR_EVENT_DONE                                                                 = 710
	ERROR_EVENT_PENDING                                                              = 711
	ERROR_CHECKING_FILE_SYSTEM                                                       = 712
	ERROR_FATAL_APP_EXIT                                                             = 713
	ERROR_PREDEFINED_HANDLE                                                          = 714
	ERROR_WAS_UNLOCKED                                                               = 715
	ERROR_SERVICE_NOTIFICATION                                                       = 716
	ERROR_WAS_LOCKED                                                                 = 717
	ERROR_LOG_HARD_ERROR                                                             = 718
	ERROR_ALREADY_WIN32                                                              = 719
	ERROR_IMAGE_MACHINE_TYPE_MISMATCH_EXE                                            = 720
	ERROR_NO_YIELD_PERFORMED                                                         = 721
	ERROR_TIMER_RESUME_IGNORED                                                       = 722
	ERROR_ARBITRATION_UNHANDLED                                                      = 723
	ERROR_CARDBUS_NOT_SUPPORTED                                                      = 724
	ERROR_MP_PROCESSOR_MISMATCH                                                      = 725
	ERROR_HIBERNATED                                                                 = 726
	ERROR_RESUME_HIBERNATION                                                         = 727
	ERROR_FIRMWARE_UPDATED                                                           = 728
	ERROR_DRIVERS_LEAKING_LOCKED_PAGES                                               = 729
	ERROR_WAKE_SYSTEM                                                                = 730
	ERROR_WAIT_1                                                                     = 731
	ERROR_WAIT_2                                                                     = 732
	ERROR_WAIT_3                                                                     = 733
	ERROR_WAIT_63                                                                    = 734
	ERROR_ABANDONED_WAIT_0                                                           = 735
	ERROR_ABANDONED_WAIT_63                                                          = 736
	ERROR_USER_APC                                                                   = 737
	ERROR_KERNEL_APC                                                                 = 738
	ERROR_ALERTED                                                                    = 739
	ERROR_ELEVATION_REQUIRED                                                         = 740
	ERROR_REPARSE                                                                    = 741
	ERROR_OPLOCK_BREAK_IN_PROGRESS                                                   = 742
	ERROR_VOLUME_MOUNTED                                                             = 743
	ERROR_RXACT_COMMITTED                                                            = 744
	ERROR_NOTIFY_CLEANUP                                                             = 745
	ERROR_PRIMARY_TRANSPORT_CONNECT_FAILED                                           = 746
	ERROR_PAGE_FAULT_TRANSITION                                                      = 747
	ERROR_PAGE_FAULT_DEMAND_ZERO                                                     = 748
	ERROR_PAGE_FAULT_COPY_ON_WRITE                                                   = 749
	ERROR_PAGE_FAULT_GUARD_PAGE                                                      = 750
	ERROR_PAGE_FAULT_PAGING_FILE                                                     = 751
	ERROR_CACHE_PAGE_LOCKED                                                          = 752
	ERROR_CRASH_DUMP                                                                 = 753
	ERROR_BUFFER_ALL_ZEROS                                                           = 754
	ERROR_REPARSE_OBJECT                                                             = 755
	ERROR_RESOURCE_REQUIREMENTS_CHANGED                                              = 756
	ERROR_TRANSLATION_COMPLETE                                                       = 757
	ERROR_NOTHING_TO_TERMINATE                                                       = 758
	ERROR_PROCESS_NOT_IN_JOB                                                         = 759
	ERROR_PROCESS_IN_JOB                                                             = 760
	ERROR_VOLSNAP_HIBERNATE_READY                                                    = 761
	ERROR_FSFILTER_OP_COMPLETED_SUCCESSFULLY                                         = 762
	ERROR_INTERRUPT_VECTOR_ALREADY_CONNECTED                                         = 763
	ERROR_INTERRUPT_STILL_CONNECTED                                                  = 764
	ERROR_WAIT_FOR_OPLOCK                                                            = 765
	ERROR_DBG_EXCEPTION_HANDLED                                                      = 766
	ERROR_DBG_CONTINUE                                                               = 767
	ERROR_CALLBACK_POP_STACK                                                         = 768
	ERROR_COMPRESSION_DISABLED                                                       = 769
	ERROR_CANTFETCHBACKWARDS                                                         = 770
	ERROR_CANTSCROLLBACKWARDS                                                        = 771
	ERROR_ROWSNOTRELEASED                                                            = 772
	ERROR_BAD_ACCESSOR_FLAGS                                                         = 773
	ERROR_ERRORS_ENCOUNTERED                                                         = 774
	ERROR_NOT_CAPABLE                                                                = 775
	ERROR_REQUEST_OUT_OF_SEQUENCE                                                    = 776
	ERROR_VERSION_PARSE_ERROR                                                        = 777
	ERROR_BADSTARTPOSITION                                                           = 778
	ERROR_MEMORY_HARDWARE                                                            = 779
	ERROR_DISK_REPAIR_DISABLED                                                       = 780
	ERROR_INSUFFICIENT_RESOURCE_FOR_SPECIFIED_SHARED_SECTION_SIZE                    = 781
	ERROR_SYSTEM_POWERSTATE_TRANSITION                                               = 782
	ERROR_SYSTEM_POWERSTATE_COMPLEX_TRANSITION                                       = 783
	ERROR_MCA_EXCEPTION                                                              = 784
	ERROR_ACCESS_AUDIT_BY_POLICY                                                     = 785
	ERROR_ACCESS_DISABLED_NO_SAFER_UI_BY_POLICY                                      = 786
	ERROR_ABANDON_HIBERFILE                                                          = 787
	ERROR_LOST_WRITEBEHIND_DATA_NETWORK_DISCONNECTED                                 = 788
	ERROR_LOST_WRITEBEHIND_DATA_NETWORK_SERVER_ERROR                                 = 789
	ERROR_LOST_WRITEBEHIND_DATA_LOCAL_DISK_ERROR                                     = 790
	ERROR_BAD_MCFG_TABLE                                                             = 791
	ERROR_OPLOCK_SWITCHED_TO_NEW_HANDLE                                              = 800
	ERROR_CANNOT_GRANT_REQUESTED_OPLOCK                                              = 801
	ERROR_CANNOT_BREAK_OPLOCK                                                        = 802
	ERROR_OPLOCK_HANDLE_CLOSED                                                       = 803
	ERROR_NO_ACE_CONDITION                                                           = 804
	ERROR_INVALID_ACE_CONDITION                                                      = 805
	ERROR_EA_ACCESS_DENIED                                                           = 994
	ERROR_OPERATION_ABORTED                                                          = 995
	ERROR_IO_INCOMPLETE                                                              = 996
	ERROR_IO_PENDING                                                                 = 997
	ERROR_NOACCESS                                                                   = 998
	ERROR_SWAPERROR                                                                  = 999
	ERROR_STACK_OVERFLOW                                                             = 1001
	ERROR_INVALID_MESSAGE                                                            = 1002
	ERROR_CAN_NOT_COMPLETE                                                           = 1003
	ERROR_INVALID_FLAGS                                                              = 1004
	ERROR_UNRECOGNIZED_VOLUME                                                        = 1005
	ERROR_FILE_INVALID                                                               = 1006
	ERROR_FULLSCREEN_MODE                                                            = 1007
	ERROR_NO_TOKEN                                                                   = 1008
	ERROR_BADDB                                                                      = 1009
	ERROR_BADKEY                                                                     = 1010
	ERROR_CANTOPEN                                                                   = 1011
	ERROR_CANTREAD                                                                   = 1012
	ERROR_CANTWRITE                                                                  = 1013
	ERROR_REGISTRY_RECOVERED                                                         = 1014
	ERROR_REGISTRY_CORRUPT                                                           = 1015
	ERROR_REGISTRY_IO_FAILED                                                         = 1016
	ERROR_NOT_REGISTRY_FILE                                                          = 1017
	ERROR_KEY_DELETED                                                                = 1018
	ERROR_NO_LOG_SPACE                                                               = 1019
	ERROR_KEY_HAS_CHILDREN                                                           = 1020
	ERROR_CHILD_MUST_BE_VOLATILE                                                     = 1021
	ERROR_NOTIFY_ENUM_DIR                                                            = 1022
	ERROR_DEPENDENT_SERVICES_RUNNING                                                 = 1051
	ERROR_INVALID_SERVICE_CONTROL                                                    = 1052
	ERROR_SERVICE_REQUEST_TIMEOUT                                                    = 1053
	ERROR_SERVICE_NO_THREAD                                                          = 1054
	ERROR_SERVICE_DATABASE_LOCKED                                                    = 1055
	ERROR_SERVICE_ALREADY_RUNNING                                                    = 1056
	ERROR_INVALID_SERVICE_ACCOUNT                                                    = 1057
	ERROR_SERVICE_DISABLED                                                           = 1058
	ERROR_CIRCULAR_DEPENDENCY                                                        = 1059
	ERROR_SERVICE_DOES_NOT_EXIST                                                     = 1060
	ERROR_SERVICE_CANNOT_ACCEPT_CTRL                                                 = 1061
	ERROR_SERVICE_NOT_ACTIVE                                                         = 1062
	ERROR_FAILED_SERVICE_CONTROLLER_CONNECT                                          = 1063
	ERROR_EXCEPTION_IN_SERVICE                                                       = 1064
	ERROR_DATABASE_DOES_NOT_EXIST                                                    = 1065
	ERROR_SERVICE_SPECIFIC_ERROR                                                     = 1066
	ERROR_PROCESS_ABORTED                                                            = 1067
	ERROR_SERVICE_DEPENDENCY_FAIL                                                    = 1068
	ERROR_SERVICE_LOGON_FAILED                                                       = 1069
	ERROR_SERVICE_START_HANG                                                         = 1070
	ERROR_INVALID_SERVICE_LOCK                                                       = 1071
	ERROR_SERVICE_MARKED_FOR_DELETE                                                  = 1072
	ERROR_SERVICE_EXISTS                                                             = 1073
	ERROR_ALREADY_RUNNING_LKG                                                        = 1074
	ERROR_SERVICE_DEPENDENCY_DELETED                                                 = 1075
	ERROR_BOOT_ALREADY_ACCEPTED                                                      = 1076
	ERROR_SERVICE_NEVER_STARTED                                                      = 1077
	ERROR_DUPLICATE_SERVICE_NAME                                                     = 1078
	ERROR_DIFFERENT_SERVICE_ACCOUNT                                                  = 1079
	ERROR_CANNOT_DETECT_DRIVER_FAILURE                                               = 1080
	ERROR_CANNOT_DETECT_PROCESS_ABORT                                                = 1081
	ERROR_NO_RECOVERY_PROGRAM                                                        = 1082
	ERROR_SERVICE_NOT_IN_EXE                                                         = 1083
	ERROR_NOT_SAFEBOOT_SERVICE                                                       = 1084
	ERROR_END_OF_MEDIA                                                               = 1100
	ERROR_FILEMARK_DETECTED                                                          = 1101
	ERROR_BEGINNING_OF_MEDIA                                                         = 1102
	ERROR_SETMARK_DETECTED                                                           = 1103
	ERROR_NO_DATA_DETECTED                                                           = 1104
	ERROR_PARTITION_FAILURE                                                          = 1105
	ERROR_INVALID_BLOCK_LENGTH                                                       = 1106
	ERROR_DEVICE_NOT_PARTITIONED                                                     = 1107
	ERROR_UNABLE_TO_LOCK_MEDIA                                                       = 1108
	ERROR_UNABLE_TO_UNLOAD_MEDIA                                                     = 1109
	ERROR_MEDIA_CHANGED                                                              = 1110
	ERROR_BUS_RESET                                                                  = 1111
	ERROR_NO_MEDIA_IN_DRIVE                                                          = 1112
	ERROR_NO_UNICODE_TRANSLATION                                                     = 1113
	ERROR_DLL_INIT_FAILED                                                            = 1114
	ERROR_SHUTDOWN_IN_PROGRESS                                                       = 1115
	ERROR_NO_SHUTDOWN_IN_PROGRESS                                                    = 1116
	ERROR_IO_DEVICE                                                                  = 1117
	ERROR_SERIAL_NO_DEVICE                                                           = 1118
	ERROR_IRQ_BUSY                                                                   = 1119
	ERROR_MORE_WRITES                                                                = 1120
	ERROR_COUNTER_TIMEOUT                                                            = 1121
	ERROR_FLOPPY_ID_MARK_NOT_FOUND                                                   = 1122
	ERROR_FLOPPY_WRONG_CYLINDER                                                      = 1123
	ERROR_FLOPPY_UNKNOWN_ERROR                                                       = 1124
	ERROR_FLOPPY_BAD_REGISTERS                                                       = 1125
	ERROR_DISK_RECALIBRATE_FAILED                                                    = 1126
	ERROR_DISK_OPERATION_FAILED                                                      = 1127
	ERROR_DISK_RESET_FAILED                                                          = 1128
	ERROR_EOM_OVERFLOW                                                               = 1129
	ERROR_NOT_ENOUGH_SERVER_MEMORY                                                   = 1130
	ERROR_POSSIBLE_DEADLOCK                                                          = 1131
	ERROR_MAPPED_ALIGNMENT                                                           = 1132
	ERROR_SET_POWER_STATE_VETOED                                                     = 1140
	ERROR_SET_POWER_STATE_FAILED                                                     = 1141
	ERROR_TOO_MANY_LINKS                                                             = 1142
	ERROR_OLD_WIN_VERSION                                                            = 1150
	ERROR_APP_WRONG_OS                                                               = 1151
	ERROR_SINGLE_INSTANCE_APP                                                        = 1152
	ERROR_RMODE_APP                                                                  = 1153
	ERROR_INVALID_DLL                                                                = 1154
	ERROR_NO_ASSOCIATION                                                             = 1155
	ERROR_DDE_FAIL                                                                   = 1156
	ERROR_DLL_NOT_FOUND                                                              = 1157
	ERROR_NO_MORE_USER_HANDLES                                                       = 1158
	ERROR_MESSAGE_SYNC_ONLY                                                          = 1159
	ERROR_SOURCE_ELEMENT_EMPTY                                                       = 1160
	ERROR_DESTINATION_ELEMENT_FULL                                                   = 1161
	ERROR_ILLEGAL_ELEMENT_ADDRESS                                                    = 1162
	ERROR_MAGAZINE_NOT_PRESENT                                                       = 1163
	ERROR_DEVICE_REINITIALIZATION_NEEDED                                             = 1164
	ERROR_DEVICE_REQUIRES_CLEANING                                                   = 1165
	ERROR_DEVICE_DOOR_OPEN                                                           = 1166
	ERROR_DEVICE_NOT_CONNECTED                                                       = 1167
	ERROR_NOT_FOUND                                                                  = 1168
	ERROR_NO_MATCH                                                                   = 1169
	ERROR_SET_NOT_FOUND                                                              = 1170
	ERROR_POINT_NOT_FOUND                                                            = 1171
	ERROR_NO_TRACKING_SERVICE                                                        = 1172
	ERROR_NO_VOLUME_ID                                                               = 1173
	ERROR_UNABLE_TO_REMOVE_REPLACED                                                  = 1175
	ERROR_UNABLE_TO_MOVE_REPLACEMENT                                                 = 1176
	ERROR_UNABLE_TO_MOVE_REPLACEMENT_2                                               = 1177
	ERROR_JOURNAL_DELETE_IN_PROGRESS                                                 = 1178
	ERROR_JOURNAL_NOT_ACTIVE                                                         = 1179
	ERROR_POTENTIAL_FILE_FOUND                                                       = 1180
	ERROR_JOURNAL_ENTRY_DELETED                                                      = 1181
	ERROR_SHUTDOWN_IS_SCHEDULED                                                      = 1190
	ERROR_SHUTDOWN_USERS_LOGGED_ON                                                   = 1191
	ERROR_BAD_DEVICE                                                                 = 1200
	ERROR_CONNECTION_UNAVAIL                                                         = 1201
	ERROR_DEVICE_ALREADY_REMEMBERED                                                  = 1202
	ERROR_NO_NET_OR_BAD_PATH                                                         = 1203
	ERROR_BAD_PROVIDER                                                               = 1204
	ERROR_CANNOT_OPEN_PROFILE                                                        = 1205
	ERROR_BAD_PROFILE                                                                = 1206
	ERROR_NOT_CONTAINER                                                              = 1207
	ERROR_EXTENDED_ERROR                                                             = 1208
	ERROR_INVALID_GROUPNAME                                                          = 1209
	ERROR_INVALID_COMPUTERNAME                                                       = 1210
	ERROR_INVALID_EVENTNAME                                                          = 1211
	ERROR_INVALID_DOMAINNAME                                                         = 1212
	ERROR_INVALID_SERVICENAME                                                        = 1213
	ERROR_INVALID_NETNAME                                                            = 1214
	ERROR_INVALID_SHARENAME                                                          = 1215
	ERROR_INVALID_PASSWORDNAME                                                       = 1216
	ERROR_INVALID_MESSAGENAME                                                        = 1217
	ERROR_INVALID_MESSAGEDEST                                                        = 1218
	ERROR_SESSION_CREDENTIAL_CONFLICT                                                = 1219
	ERROR_REMOTE_SESSION_LIMIT_EXCEEDED                                              = 1220
	ERROR_DUP_DOMAINNAME                                                             = 1221
	ERROR_NO_NETWORK                                                                 = 1222
	ERROR_CANCELLED                                                                  = 1223
	ERROR_USER_MAPPED_FILE                                                           = 1224
	ERROR_CONNECTION_REFUSED                                                         = 1225
	ERROR_GRACEFUL_DISCONNECT                                                        = 1226
	ERROR_ADDRESS_ALREADY_ASSOCIATED                                                 = 1227
	ERROR_ADDRESS_NOT_ASSOCIATED                                                     = 1228
	ERROR_CONNECTION_INVALID                                                         = 1229
	ERROR_CONNECTION_ACTIVE                                                          = 1230
	ERROR_NETWORK_UNREACHABLE                                                        = 1231
	ERROR_HOST_UNREACHABLE                                                           = 1232
	ERROR_PROTOCOL_UNREACHABLE                                                       = 1233
	ERROR_PORT_UNREACHABLE                                                           = 1234
	ERROR_REQUEST_ABORTED                                                            = 1235
	ERROR_CONNECTION_ABORTED                                                         = 1236
	ERROR_RETRY                                                                      = 1237
	ERROR_CONNECTION_COUNT_LIMIT                                                     = 1238
	ERROR_LOGIN_TIME_RESTRICTION                                                     = 1239
	ERROR_LOGIN_WKSTA_RESTRICTION                                                    = 1240
	ERROR_INCORRECT_ADDRESS                                                          = 1241
	ERROR_ALREADY_REGISTERED                                                         = 1242
	ERROR_SERVICE_NOT_FOUND                                                          = 1243
	ERROR_NOT_AUTHENTICATED                                                          = 1244
	ERROR_NOT_LOGGED_ON                                                              = 1245
	ERROR_CONTINUE                                                                   = 1246
	ERROR_ALREADY_INITIALIZED                                                        = 1247
	ERROR_NO_MORE_DEVICES                                                            = 1248
	ERROR_NO_SUCH_SITE                                                               = 1249
	ERROR_DOMAIN_CONTROLLER_EXISTS                                                   = 1250
	ERROR_ONLY_IF_CONNECTED                                                          = 1251
	ERROR_OVERRIDE_NOCHANGES                                                         = 1252
	ERROR_BAD_USER_PROFILE                                                           = 1253
	ERROR_NOT_SUPPORTED_ON_SBS                                                       = 1254
	ERROR_SERVER_SHUTDOWN_IN_PROGRESS                                                = 1255
	ERROR_HOST_DOWN                                                                  = 1256
	ERROR_NON_ACCOUNT_SID                                                            = 1257
	ERROR_NON_DOMAIN_SID                                                             = 1258
	ERROR_APPHELP_BLOCK                                                              = 1259
	ERROR_ACCESS_DISABLED_BY_POLICY                                                  = 1260
	ERROR_REG_NAT_CONSUMPTION                                                        = 1261
	ERROR_CSCSHARE_OFFLINE                                                           = 1262
	ERROR_PKINIT_FAILURE                                                             = 1263
	ERROR_SMARTCARD_SUBSYSTEM_FAILURE                                                = 1264
	ERROR_DOWNGRADE_DETECTED                                                         = 1265
	ERROR_MACHINE_LOCKED                                                             = 1271
	ERROR_CALLBACK_SUPPLIED_INVALID_DATA                                             = 1273
	ERROR_SYNC_FOREGROUND_REFRESH_REQUIRED                                           = 1274
	ERROR_DRIVER_BLOCKED                                                             = 1275
	ERROR_INVALID_IMPORT_OF_NON_DLL                                                  = 1276
	ERROR_ACCESS_DISABLED_WEBBLADE                                                   = 1277
	ERROR_ACCESS_DISABLED_WEBBLADE_TAMPER                                            = 1278
	ERROR_RECOVERY_FAILURE                                                           = 1279
	ERROR_ALREADY_FIBER                                                              = 1280
	ERROR_ALREADY_THREAD                                                             = 1281
	ERROR_STACK_BUFFER_OVERRUN                                                       = 1282
	ERROR_PARAMETER_QUOTA_EXCEEDED                                                   = 1283
	ERROR_DEBUGGER_INACTIVE                                                          = 1284
	ERROR_DELAY_LOAD_FAILED                                                          = 1285
	ERROR_VDM_DISALLOWED                                                             = 1286
	ERROR_UNIDENTIFIED_ERROR                                                         = 1287
	ERROR_INVALID_CRUNTIME_PARAMETER                                                 = 1288
	ERROR_BEYOND_VDL                                                                 = 1289
	ERROR_INCOMPATIBLE_SERVICE_SID_TYPE                                              = 1290
	ERROR_DRIVER_PROCESS_TERMINATED                                                  = 1291
	ERROR_IMPLEMENTATION_LIMIT                                                       = 1292
	ERROR_PROCESS_IS_PROTECTED                                                       = 1293
	ERROR_SERVICE_NOTIFY_CLIENT_LAGGING                                              = 1294
	ERROR_DISK_QUOTA_EXCEEDED                                                        = 1295
	ERROR_CONTENT_BLOCKED                                                            = 1296
	ERROR_INCOMPATIBLE_SERVICE_PRIVILEGE                                             = 1297
	ERROR_APP_HANG                                                                   = 1298
	ERROR_INVALID_LABEL                                                              = 1299
	ERROR_NOT_ALL_ASSIGNED                                                           = 1300
	ERROR_SOME_NOT_MAPPED                                                            = 1301
	ERROR_NO_QUOTAS_FOR_ACCOUNT                                                      = 1302
	ERROR_LOCAL_USER_SESSION_KEY                                                     = 1303
	ERROR_NULL_LM_PASSWORD                                                           = 1304
	ERROR_UNKNOWN_REVISION                                                           = 1305
	ERROR_REVISION_MISMATCH                                                          = 1306
	ERROR_INVALID_OWNER                                                              = 1307
	ERROR_INVALID_PRIMARY_GROUP                                                      = 1308
	ERROR_NO_IMPERSONATION_TOKEN                                                     = 1309
	ERROR_CANT_DISABLE_MANDATORY                                                     = 1310
	ERROR_NO_LOGON_SERVERS                                                           = 1311
	ERROR_NO_SUCH_LOGON_SESSION                                                      = 1312
	ERROR_NO_SUCH_PRIVILEGE                                                          = 1313
	ERROR_PRIVILEGE_NOT_HELD                                                         = 1314
	ERROR_INVALID_ACCOUNT_NAME                                                       = 1315
	ERROR_USER_EXISTS                                                                = 1316
	ERROR_NO_SUCH_USER                                                               = 1317
	ERROR_GROUP_EXISTS                                                               = 1318
	ERROR_NO_SUCH_GROUP                                                              = 1319
	ERROR_MEMBER_IN_GROUP                                                            = 1320
	ERROR_MEMBER_NOT_IN_GROUP                                                        = 1321
	ERROR_LAST_ADMIN                                                                 = 1322
	ERROR_WRONG_PASSWORD                                                             = 1323
	ERROR_ILL_FORMED_PASSWORD                                                        = 1324
	ERROR_PASSWORD_RESTRICTION                                                       = 1325
	ERROR_LOGON_FAILURE                                                              = 1326
	ERROR_ACCOUNT_RESTRICTION                                                        = 1327
	ERROR_INVALID_LOGON_HOURS                                                        = 1328
	ERROR_INVALID_WORKSTATION                                                        = 1329
	ERROR_PASSWORD_EXPIRED                                                           = 1330
	ERROR_ACCOUNT_DISABLED                                                           = 1331
	ERROR_NONE_MAPPED                                                                = 1332
	ERROR_TOO_MANY_LUIDS_REQUESTED                                                   = 1333
	ERROR_LUIDS_EXHAUSTED                                                            = 1334
	ERROR_INVALID_SUB_AUTHORITY                                                      = 1335
	ERROR_INVALID_ACL                                                                = 1336
	ERROR_INVALID_SID                                                                = 1337
	ERROR_INVALID_SECURITY_DESCR                                                     = 1338
	ERROR_BAD_INHERITANCE_ACL                                                        = 1340
	ERROR_SERVER_DISABLED                                                            = 1341
	ERROR_SERVER_NOT_DISABLED                                                        = 1342
	ERROR_INVALID_ID_AUTHORITY                                                       = 1343
	ERROR_ALLOTTED_SPACE_EXCEEDED                                                    = 1344
	ERROR_INVALID_GROUP_ATTRIBUTES                                                   = 1345
	ERROR_BAD_IMPERSONATION_LEVEL                                                    = 1346
	ERROR_CANT_OPEN_ANONYMOUS                                                        = 1347
	ERROR_BAD_VALIDATION_CLASS                                                       = 1348
	ERROR_BAD_TOKEN_TYPE                                                             = 1349
	ERROR_NO_SECURITY_ON_OBJECT                                                      = 1350
	ERROR_CANT_ACCESS_DOMAIN_INFO                                                    = 1351
	ERROR_INVALID_SERVER_STATE                                                       = 1352
	ERROR_INVALID_DOMAIN_STATE                                                       = 1353
	ERROR_INVALID_DOMAIN_ROLE                                                        = 1354
	ERROR_NO_SUCH_DOMAIN                                                             = 1355
	ERROR_DOMAIN_EXISTS                                                              = 1356
	ERROR_DOMAIN_LIMIT_EXCEEDED                                                      = 1357
	ERROR_INTERNAL_DB_CORRUPTION                                                     = 1358
	ERROR_INTERNAL_ERROR                                                             = 1359
	ERROR_GENERIC_NOT_MAPPED                                                         = 1360
	ERROR_BAD_DESCRIPTOR_FORMAT                                                      = 1361
	ERROR_NOT_LOGON_PROCESS                                                          = 1362
	ERROR_LOGON_SESSION_EXISTS                                                       = 1363
	ERROR_NO_SUCH_PACKAGE                                                            = 1364
	ERROR_BAD_LOGON_SESSION_STATE                                                    = 1365
	ERROR_LOGON_SESSION_COLLISION                                                    = 1366
	ERROR_INVALID_LOGON_TYPE                                                         = 1367
	ERROR_CANNOT_IMPERSONATE                                                         = 1368
	ERROR_RXACT_INVALID_STATE                                                        = 1369
	ERROR_RXACT_COMMIT_FAILURE                                                       = 1370
	ERROR_SPECIAL_ACCOUNT                                                            = 1371
	ERROR_SPECIAL_GROUP                                                              = 1372
	ERROR_SPECIAL_USER                                                               = 1373
	ERROR_MEMBERS_PRIMARY_GROUP                                                      = 1374
	ERROR_TOKEN_ALREADY_IN_USE                                                       = 1375
	ERROR_NO_SUCH_ALIAS                                                              = 1376
	ERROR_MEMBER_NOT_IN_ALIAS                                                        = 1377
	ERROR_MEMBER_IN_ALIAS                                                            = 1378
	ERROR_ALIAS_EXISTS                                                               = 1379
	ERROR_LOGON_NOT_GRANTED                                                          = 1380
	ERROR_TOO_MANY_SECRETS                                                           = 1381
	ERROR_SECRET_TOO_LONG                                                            = 1382
	ERROR_INTERNAL_DB_ERROR                                                          = 1383
	ERROR_TOO_MANY_CONTEXT_IDS                                                       = 1384
	ERROR_LOGON_TYPE_NOT_GRANTED                                                     = 1385
	ERROR_NT_CROSS_ENCRYPTION_REQUIRED                                               = 1386
	ERROR_NO_SUCH_MEMBER                                                             = 1387
	ERROR_INVALID_MEMBER                                                             = 1388
	ERROR_TOO_MANY_SIDS                                                              = 1389
	ERROR_LM_CROSS_ENCRYPTION_REQUIRED                                               = 1390
	ERROR_NO_INHERITANCE                                                             = 1391
	ERROR_FILE_CORRUPT                                                               = 1392
	ERROR_DISK_CORRUPT                                                               = 1393
	ERROR_NO_USER_SESSION_KEY                                                        = 1394
	ERROR_LICENSE_QUOTA_EXCEEDED                                                     = 1395
	ERROR_WRONG_TARGET_NAME                                                          = 1396
	ERROR_MUTUAL_AUTH_FAILED                                                         = 1397
	ERROR_TIME_SKEW                                                                  = 1398
	ERROR_CURRENT_DOMAIN_NOT_ALLOWED                                                 = 1399
	ERROR_INVALID_WINDOW_HANDLE                                                      = 1400
	ERROR_INVALID_MENU_HANDLE                                                        = 1401
	ERROR_INVALID_CURSOR_HANDLE                                                      = 1402
	ERROR_INVALID_ACCEL_HANDLE                                                       = 1403
	ERROR_INVALID_HOOK_HANDLE                                                        = 1404
	ERROR_INVALID_DWP_HANDLE                                                         = 1405
	ERROR_TLW_WITH_WSCHILD                                                           = 1406
	ERROR_CANNOT_FIND_WND_CLASS                                                      = 1407
	ERROR_WINDOW_OF_OTHER_THREAD                                                     = 1408
	ERROR_HOTKEY_ALREADY_REGISTERED                                                  = 1409
	ERROR_CLASS_ALREADY_EXISTS                                                       = 1410
	ERROR_CLASS_DOES_NOT_EXIST                                                       = 1411
	ERROR_CLASS_HAS_WINDOWS                                                          = 1412
	ERROR_INVALID_INDEX                                                              = 1413
	ERROR_INVALID_ICON_HANDLE                                                        = 1414
	ERROR_PRIVATE_DIALOG_INDEX                                                       = 1415
	ERROR_LISTBOX_ID_NOT_FOUND                                                       = 1416
	ERROR_NO_WILDCARD_CHARACTERS                                                     = 1417
	ERROR_CLIPBOARD_NOT_OPEN                                                         = 1418
	ERROR_HOTKEY_NOT_REGISTERED                                                      = 1419
	ERROR_WINDOW_NOT_DIALOG                                                          = 1420
	ERROR_CONTROL_ID_NOT_FOUND                                                       = 1421
	ERROR_INVALID_COMBOBOX_MESSAGE                                                   = 1422
	ERROR_WINDOW_NOT_COMBOBOX                                                        = 1423
	ERROR_INVALID_EDIT_HEIGHT                                                        = 1424
	ERROR_DC_NOT_FOUND                                                               = 1425
	ERROR_INVALID_HOOK_FILTER                                                        = 1426
	ERROR_INVALID_FILTER_PROC                                                        = 1427
	ERROR_HOOK_NEEDS_HMOD                                                            = 1428
	ERROR_GLOBAL_ONLY_HOOK                                                           = 1429
	ERROR_JOURNAL_HOOK_SET                                                           = 1430
	ERROR_HOOK_NOT_INSTALLED                                                         = 1431
	ERROR_INVALID_LB_MESSAGE                                                         = 1432
	ERROR_SETCOUNT_ON_BAD_LB                                                         = 1433
	ERROR_LB_WITHOUT_TABSTOPS                                                        = 1434
	ERROR_DESTROY_OBJECT_OF_OTHER_THREAD                                             = 1435
	ERROR_CHILD_WINDOW_MENU                                                          = 1436
	ERROR_NO_SYSTEM_MENU                                                             = 1437
	ERROR_INVALID_MSGBOX_STYLE                                                       = 1438
	ERROR_INVALID_SPI_VALUE                                                          = 1439
	ERROR_SCREEN_ALREADY_LOCKED                                                      = 1440
	ERROR_HWNDS_HAVE_DIFF_PARENT                                                     = 1441
	ERROR_NOT_CHILD_WINDOW                                                           = 1442
	ERROR_INVALID_GW_COMMAND                                                         = 1443
	ERROR_INVALID_THREAD_ID                                                          = 1444
	ERROR_NON_MDICHILD_WINDOW                                                        = 1445
	ERROR_POPUP_ALREADY_ACTIVE                                                       = 1446
	ERROR_NO_SCROLLBARS                                                              = 1447
	ERROR_INVALID_SCROLLBAR_RANGE                                                    = 1448
	ERROR_INVALID_SHOWWIN_COMMAND                                                    = 1449
	ERROR_NO_SYSTEM_RESOURCES                                                        = 1450
	ERROR_NONPAGED_SYSTEM_RESOURCES                                                  = 1451
	ERROR_PAGED_SYSTEM_RESOURCES                                                     = 1452
	ERROR_WORKING_SET_QUOTA                                                          = 1453
	ERROR_PAGEFILE_QUOTA                                                             = 1454
	ERROR_COMMITMENT_LIMIT                                                           = 1455
	ERROR_MENU_ITEM_NOT_FOUND                                                        = 1456
	ERROR_INVALID_KEYBOARD_HANDLE                                                    = 1457
	ERROR_HOOK_TYPE_NOT_ALLOWED                                                      = 1458
	ERROR_REQUIRES_INTERACTIVE_WINDOWSTATION                                         = 1459
	ERROR_TIMEOUT                                                                    = 1460
	ERROR_INVALID_MONITOR_HANDLE                                                     = 1461
	ERROR_INCORRECT_SIZE                                                             = 1462
	ERROR_SYMLINK_CLASS_DISABLED                                                     = 1463
	ERROR_SYMLINK_NOT_SUPPORTED                                                      = 1464
	ERROR_XML_PARSE_ERROR                                                            = 1465
	ERROR_XMLDSIG_ERROR                                                              = 1466
	ERROR_RESTART_APPLICATION                                                        = 1467
	ERROR_WRONG_COMPARTMENT                                                          = 1468
	ERROR_AUTHIP_FAILURE                                                             = 1469
	ERROR_NO_NVRAM_RESOURCES                                                         = 1470
	ERROR_EVENTLOG_FILE_CORRUPT                                                      = 1500
	ERROR_EVENTLOG_CANT_START                                                        = 1501
	ERROR_LOG_FILE_FULL                                                              = 1502
	ERROR_EVENTLOG_FILE_CHANGED                                                      = 1503
	ERROR_INVALID_TASK_NAME                                                          = 1550
	ERROR_INVALID_TASK_INDEX                                                         = 1551
	ERROR_THREAD_ALREADY_IN_TASK                                                     = 1552
	ERROR_INSTALL_SERVICE_FAILURE                                                    = 1601
	ERROR_INSTALL_USEREXIT                                                           = 1602
	ERROR_INSTALL_FAILURE                                                            = 1603
	ERROR_INSTALL_SUSPEND                                                            = 1604
	ERROR_UNKNOWN_PRODUCT                                                            = 1605
	ERROR_UNKNOWN_FEATURE                                                            = 1606
	ERROR_UNKNOWN_COMPONENT                                                          = 1607
	ERROR_UNKNOWN_PROPERTY                                                           = 1608
	ERROR_INVALID_HANDLE_STATE                                                       = 1609
	ERROR_BAD_CONFIGURATION                                                          = 1610
	ERROR_INDEX_ABSENT                                                               = 1611
	ERROR_INSTALL_SOURCE_ABSENT                                                      = 1612
	ERROR_INSTALL_PACKAGE_VERSION                                                    = 1613
	ERROR_PRODUCT_UNINSTALLED                                                        = 1614
	ERROR_BAD_QUERY_SYNTAX                                                           = 1615
	ERROR_INVALID_FIELD                                                              = 1616
	ERROR_DEVICE_REMOVED                                                             = 1617
	ERROR_INSTALL_ALREADY_RUNNING                                                    = 1618
	ERROR_INSTALL_PACKAGE_OPEN_FAILED                                                = 1619
	ERROR_INSTALL_PACKAGE_INVALID                                                    = 1620
	ERROR_INSTALL_UI_FAILURE                                                         = 1621
	ERROR_INSTALL_LOG_FAILURE                                                        = 1622
	ERROR_INSTALL_LANGUAGE_UNSUPPORTED                                               = 1623
	ERROR_INSTALL_TRANSFORM_FAILURE                                                  = 1624
	ERROR_INSTALL_PACKAGE_REJECTED                                                   = 1625
	ERROR_FUNCTION_NOT_CALLED                                                        = 1626
	ERROR_FUNCTION_FAILED                                                            = 1627
	ERROR_INVALID_TABLE                                                              = 1628
	ERROR_DATATYPE_MISMATCH                                                          = 1629
	ERROR_UNSUPPORTED_TYPE                                                           = 1630
	ERROR_CREATE_FAILED                                                              = 1631
	ERROR_INSTALL_TEMP_UNWRITABLE                                                    = 1632
	ERROR_INSTALL_PLATFORM_UNSUPPORTED                                               = 1633
	ERROR_INSTALL_NOTUSED                                                            = 1634
	ERROR_PATCH_PACKAGE_OPEN_FAILED                                                  = 1635
	ERROR_PATCH_PACKAGE_INVALID                                                      = 1636
	ERROR_PATCH_PACKAGE_UNSUPPORTED                                                  = 1637
	ERROR_PRODUCT_VERSION                                                            = 1638
	ERROR_INVALID_COMMAND_LINE                                                       = 1639
	ERROR_INSTALL_REMOTE_DISALLOWED                                                  = 1640
	ERROR_SUCCESS_REBOOT_INITIATED                                                   = 1641
	ERROR_PATCH_TARGET_NOT_FOUND                                                     = 1642
	ERROR_PATCH_PACKAGE_REJECTED                                                     = 1643
	ERROR_INSTALL_TRANSFORM_REJECTED                                                 = 1644
	ERROR_INSTALL_REMOTE_PROHIBITED                                                  = 1645
	ERROR_PATCH_REMOVAL_UNSUPPORTED                                                  = 1646
	ERROR_UNKNOWN_PATCH                                                              = 1647
	ERROR_PATCH_NO_SEQUENCE                                                          = 1648
	ERROR_PATCH_REMOVAL_DISALLOWED                                                   = 1649
	ERROR_INVALID_PATCH_XML                                                          = 1650
	ERROR_PATCH_MANAGED_ADVERTISED_PRODUCT                                           = 1651
	ERROR_INSTALL_SERVICE_SAFEBOOT                                                   = 1652
	ERROR_FAIL_FAST_EXCEPTION                                                        = 1653
	RPC_S_INVALID_STRING_BINDING                                                     = 1700
	RPC_S_WRONG_KIND_OF_BINDING                                                      = 1701
	RPC_S_INVALID_BINDING                                                            = 1702
	RPC_S_PROTSEQ_NOT_SUPPORTED                                                      = 1703
	RPC_S_INVALID_RPC_PROTSEQ                                                        = 1704
	RPC_S_INVALID_STRING_UUID                                                        = 1705
	RPC_S_INVALID_ENDPOINT_FORMAT                                                    = 1706
	RPC_S_INVALID_NET_ADDR                                                           = 1707
	RPC_S_NO_ENDPOINT_FOUND                                                          = 1708
	RPC_S_INVALID_TIMEOUT                                                            = 1709
	RPC_S_OBJECT_NOT_FOUND                                                           = 1710
	RPC_S_ALREADY_REGISTERED                                                         = 1711
	RPC_S_TYPE_ALREADY_REGISTERED                                                    = 1712
	RPC_S_ALREADY_LISTENING                                                          = 1713
	RPC_S_NO_PROTSEQS_REGISTERED                                                     = 1714
	RPC_S_NOT_LISTENING                                                              = 1715
	RPC_S_UNKNOWN_MGR_TYPE                                                           = 1716
	RPC_S_UNKNOWN_IF                                                                 = 1717
	RPC_S_NO_BINDINGS                                                                = 1718
	RPC_S_NO_PROTSEQS                                                                = 1719
	RPC_S_CANT_CREATE_ENDPOINT                                                       = 1720
	RPC_S_OUT_OF_RESOURCES                                                           = 1721
	RPC_S_SERVER_UNAVAILABLE                                                         = 1722
	RPC_S_SERVER_TOO_BUSY                                                            = 1723
	RPC_S_INVALID_NETWORK_OPTIONS                                                    = 1724
	RPC_S_NO_CALL_ACTIVE                                                             = 1725
	RPC_S_CALL_FAILED                                                                = 1726
	RPC_S_CALL_FAILED_DNE                                                            = 1727
	RPC_S_PROTOCOL_ERROR                                                             = 1728
	RPC_S_PROXY_ACCESS_DENIED                                                        = 1729
	RPC_S_UNSUPPORTED_TRANS_SYN                                                      = 1730
	RPC_S_UNSUPPORTED_TYPE                                                           = 1732
	RPC_S_INVALID_TAG                                                                = 1733
	RPC_S_INVALID_BOUND                                                              = 1734
	RPC_S_NO_ENTRY_NAME                                                              = 1735
	RPC_S_INVALID_NAME_SYNTAX                                                        = 1736
	RPC_S_UNSUPPORTED_NAME_SYNTAX                                                    = 1737
	RPC_S_UUID_NO_ADDRESS                                                            = 1739
	RPC_S_DUPLICATE_ENDPOINT                                                         = 1740
	RPC_S_UNKNOWN_AUTHN_TYPE                                                         = 1741
	RPC_S_MAX_CALLS_TOO_SMALL                                                        = 1742
	RPC_S_STRING_TOO_LONG                                                            = 1743
	RPC_S_PROTSEQ_NOT_FOUND                                                          = 1744
	RPC_S_PROCNUM_OUT_OF_RANGE                                                       = 1745
	RPC_S_BINDING_HAS_NO_AUTH                                                        = 1746
	RPC_S_UNKNOWN_AUTHN_SERVICE                                                      = 1747
	RPC_S_UNKNOWN_AUTHN_LEVEL                                                        = 1748
	RPC_S_INVALID_AUTH_IDENTITY                                                      = 1749
	RPC_S_UNKNOWN_AUTHZ_SERVICE                                                      = 1750
	EPT_S_INVALID_ENTRY                                                              = 1751
	EPT_S_CANT_PERFORM_OP                                                            = 1752
	EPT_S_NOT_REGISTERED                                                             = 1753
	RPC_S_NOTHING_TO_EXPORT                                                          = 1754
	RPC_S_INCOMPLETE_NAME                                                            = 1755
	RPC_S_INVALID_VERS_OPTION                                                        = 1756
	RPC_S_NO_MORE_MEMBERS                                                            = 1757
	RPC_S_NOT_ALL_OBJS_UNEXPORTED                                                    = 1758
	RPC_S_INTERFACE_NOT_FOUND                                                        = 1759
	RPC_S_ENTRY_ALREADY_EXISTS                                                       = 1760
	RPC_S_ENTRY_NOT_FOUND                                                            = 1761
	RPC_S_NAME_SERVICE_UNAVAILABLE                                                   = 1762
	RPC_S_INVALID_NAF_ID                                                             = 1763
	RPC_S_CANNOT_SUPPORT                                                             = 1764
	RPC_S_NO_CONTEXT_AVAILABLE                                                       = 1765
	RPC_S_INTERNAL_ERROR                                                             = 1766
	RPC_S_ZERO_DIVIDE                                                                = 1767
	RPC_S_ADDRESS_ERROR                                                              = 1768
	RPC_S_FP_DIV_ZERO                                                                = 1769
	RPC_S_FP_UNDERFLOW                                                               = 1770
	RPC_S_FP_OVERFLOW                                                                = 1771
	RPC_X_NO_MORE_ENTRIES                                                            = 1772
	RPC_X_SS_CHAR_TRANS_OPEN_FAIL                                                    = 1773
	RPC_X_SS_CHAR_TRANS_SHORT_FILE                                                   = 1774
	RPC_X_SS_IN_NULL_CONTEXT                                                         = 1775
	RPC_X_SS_CONTEXT_DAMAGED                                                         = 1777
	RPC_X_SS_HANDLES_MISMATCH                                                        = 1778
	RPC_X_SS_CANNOT_GET_CALL_HANDLE                                                  = 1779
	RPC_X_NULL_REF_POINTER                                                           = 1780
	RPC_X_ENUM_VALUE_OUT_OF_RANGE                                                    = 1781
	RPC_X_BYTE_COUNT_TOO_SMALL                                                       = 1782
	RPC_X_BAD_STUB_DATA                                                              = 1783
	ERROR_INVALID_USER_BUFFER                                                        = 1784
	ERROR_UNRECOGNIZED_MEDIA                                                         = 1785
	ERROR_NO_TRUST_LSA_SECRET                                                        = 1786
	ERROR_NO_TRUST_SAM_ACCOUNT                                                       = 1787
	ERROR_TRUSTED_DOMAIN_FAILURE                                                     = 1788
	ERROR_TRUSTED_RELATIONSHIP_FAILURE                                               = 1789
	ERROR_TRUST_FAILURE                                                              = 1790
	RPC_S_CALL_IN_PROGRESS                                                           = 1791
	ERROR_NETLOGON_NOT_STARTED                                                       = 1792
	ERROR_ACCOUNT_EXPIRED                                                            = 1793
	ERROR_REDIRECTOR_HAS_OPEN_HANDLES                                                = 1794
	ERROR_PRINTER_DRIVER_ALREADY_INSTALLED                                           = 1795
	ERROR_UNKNOWN_PORT                                                               = 1796
	ERROR_UNKNOWN_PRINTER_DRIVER                                                     = 1797
	ERROR_UNKNOWN_PRINTPROCESSOR                                                     = 1798
	ERROR_INVALID_SEPARATOR_FILE                                                     = 1799
	ERROR_INVALID_PRIORITY                                                           = 1800
	ERROR_INVALID_PRINTER_NAME                                                       = 1801
	ERROR_PRINTER_ALREADY_EXISTS                                                     = 1802
	ERROR_INVALID_PRINTER_COMMAND                                                    = 1803
	ERROR_INVALID_DATATYPE                                                           = 1804
	ERROR_INVALID_ENVIRONMENT                                                        = 1805
	RPC_S_NO_MORE_BINDINGS                                                           = 1806
	ERROR_NOLOGON_INTERDOMAIN_TRUST_ACCOUNT                                          = 1807
	ERROR_NOLOGON_WORKSTATION_TRUST_ACCOUNT                                          = 1808
	ERROR_NOLOGON_SERVER_TRUST_ACCOUNT                                               = 1809
	ERROR_DOMAIN_TRUST_INCONSISTENT                                                  = 1810
	ERROR_SERVER_HAS_OPEN_HANDLES                                                    = 1811
	ERROR_RESOURCE_DATA_NOT_FOUND                                                    = 1812
	ERROR_RESOURCE_TYPE_NOT_FOUND                                                    = 1813
	ERROR_RESOURCE_NAME_NOT_FOUND                                                    = 1814
	ERROR_RESOURCE_LANG_NOT_FOUND                                                    = 1815
	ERROR_NOT_ENOUGH_QUOTA                                                           = 1816
	RPC_S_NO_INTERFACES                                                              = 1817
	RPC_S_CALL_CANCELLED                                                             = 1818
	RPC_S_BINDING_INCOMPLETE                                                         = 1819
	RPC_S_COMM_FAILURE                                                               = 1820
	RPC_S_UNSUPPORTED_AUTHN_LEVEL                                                    = 1821
	RPC_S_NO_PRINC_NAME                                                              = 1822
	RPC_S_NOT_RPC_ERROR                                                              = 1823
	RPC_S_UUID_LOCAL_ONLY                                                            = 1824
	RPC_S_SEC_PKG_ERROR                                                              = 1825
	RPC_S_NOT_CANCELLED                                                              = 1826
	RPC_X_INVALID_ES_ACTION                                                          = 1827
	RPC_X_WRONG_ES_VERSION                                                           = 1828
	RPC_X_WRONG_STUB_VERSION                                                         = 1829
	RPC_X_INVALID_PIPE_OBJECT                                                        = 1830
	RPC_X_WRONG_PIPE_ORDER                                                           = 1831
	RPC_X_WRONG_PIPE_VERSION                                                         = 1832
	RPC_S_COOKIE_AUTH_FAILED                                                         = 1833
	RPC_S_GROUP_MEMBER_NOT_FOUND                                                     = 1898
	EPT_S_CANT_CREATE                                                                = 1899
	RPC_S_INVALID_OBJECT                                                             = 1900
	ERROR_INVALID_TIME                                                               = 1901
	ERROR_INVALID_FORM_NAME                                                          = 1902
	ERROR_INVALID_FORM_SIZE                                                          = 1903
	ERROR_ALREADY_WAITING                                                            = 1904
	ERROR_PRINTER_DELETED                                                            = 1905
	ERROR_INVALID_PRINTER_STATE                                                      = 1906
	ERROR_PASSWORD_MUST_CHANGE                                                       = 1907
	ERROR_DOMAIN_CONTROLLER_NOT_FOUND                                                = 1908
	ERROR_ACCOUNT_LOCKED_OUT                                                         = 1909
	OR_INVALID_OXID                                                                  = 1910
	OR_INVALID_OID                                                                   = 1911
	OR_INVALID_SET                                                                   = 1912
	RPC_S_SEND_INCOMPLETE                                                            = 1913
	RPC_S_INVALID_ASYNC_HANDLE                                                       = 1914
	RPC_S_INVALID_ASYNC_CALL                                                         = 1915
	RPC_X_PIPE_CLOSED                                                                = 1916
	RPC_X_PIPE_DISCIPLINE_ERROR                                                      = 1917
	RPC_X_PIPE_EMPTY                                                                 = 1918
	ERROR_NO_SITENAME                                                                = 1919
	ERROR_CANT_ACCESS_FILE                                                           = 1920
	ERROR_CANT_RESOLVE_FILENAME                                                      = 1921
	RPC_S_ENTRY_TYPE_MISMATCH                                                        = 1922
	RPC_S_NOT_ALL_OBJS_EXPORTED                                                      = 1923
	RPC_S_INTERFACE_NOT_EXPORTED                                                     = 1924
	RPC_S_PROFILE_NOT_ADDED                                                          = 1925
	RPC_S_PRF_ELT_NOT_ADDED                                                          = 1926
	RPC_S_PRF_ELT_NOT_REMOVED                                                        = 1927
	RPC_S_GRP_ELT_NOT_ADDED                                                          = 1928
	RPC_S_GRP_ELT_NOT_REMOVED                                                        = 1929
	ERROR_KM_DRIVER_BLOCKED                                                          = 1930
	ERROR_CONTEXT_EXPIRED                                                            = 1931
	ERROR_PER_USER_TRUST_QUOTA_EXCEEDED                                              = 1932
	ERROR_ALL_USER_TRUST_QUOTA_EXCEEDED                                              = 1933
	ERROR_USER_DELETE_TRUST_QUOTA_EXCEEDED                                           = 1934
	ERROR_AUTHENTICATION_FIREWALL_FAILED                                             = 1935
	ERROR_REMOTE_PRINT_CONNECTIONS_BLOCKED                                           = 1936
	ERROR_NTLM_BLOCKED                                                               = 1937
	ERROR_INVALID_PIXEL_FORMAT                                                       = 2000
	ERROR_BAD_DRIVER                                                                 = 2001
	ERROR_INVALID_WINDOW_STYLE                                                       = 2002
	ERROR_METAFILE_NOT_SUPPORTED                                                     = 2003
	ERROR_TRANSFORM_NOT_SUPPORTED                                                    = 2004
	ERROR_CLIPPING_NOT_SUPPORTED                                                     = 2005
	ERROR_INVALID_CMM                                                                = 2010
	ERROR_INVALID_PROFILE                                                            = 2011
	ERROR_TAG_NOT_FOUND                                                              = 2012
	ERROR_TAG_NOT_PRESENT                                                            = 2013
	ERROR_DUPLICATE_TAG                                                              = 2014
	ERROR_PROFILE_NOT_ASSOCIATED_WITH_DEVICE                                         = 2015
	ERROR_PROFILE_NOT_FOUND                                                          = 2016
	ERROR_INVALID_COLORSPACE                                                         = 2017
	ERROR_ICM_NOT_ENABLED                                                            = 2018
	ERROR_DELETING_ICM_XFORM                                                         = 2019
	ERROR_INVALID_TRANSFORM                                                          = 2020
	ERROR_COLORSPACE_MISMATCH                                                        = 2021
	ERROR_INVALID_COLORINDEX                                                         = 2022
	ERROR_PROFILE_DOES_NOT_MATCH_DEVICE                                              = 2023
	ERROR_CONNECTED_OTHER_PASSWORD                                                   = 2108
	ERROR_CONNECTED_OTHER_PASSWORD_DEFAULT                                           = 2109
	ERROR_BAD_USERNAME                                                               = 2202
	ERROR_NOT_CONNECTED                                                              = 2250
	ERROR_OPEN_FILES                                                                 = 2401
	ERROR_ACTIVE_CONNECTIONS                                                         = 2402
	ERROR_DEVICE_IN_USE                                                              = 2404
	ERROR_UNKNOWN_PRINT_MONITOR                                                      = 3000
	ERROR_PRINTER_DRIVER_IN_USE                                                      = 3001
	ERROR_SPOOL_FILE_NOT_FOUND                                                       = 3002
	ERROR_SPL_NO_STARTDOC                                                            = 3003
	ERROR_SPL_NO_ADDJOB                                                              = 3004
	ERROR_PRINT_PROCESSOR_ALREADY_INSTALLED                                          = 3005
	ERROR_PRINT_MONITOR_ALREADY_INSTALLED                                            = 3006
	ERROR_INVALID_PRINT_MONITOR                                                      = 3007
	ERROR_PRINT_MONITOR_IN_USE                                                       = 3008
	ERROR_PRINTER_HAS_JOBS_QUEUED                                                    = 3009
	ERROR_SUCCESS_REBOOT_REQUIRED                                                    = 3010
	ERROR_SUCCESS_RESTART_REQUIRED                                                   = 3011
	ERROR_PRINTER_NOT_FOUND                                                          = 3012
	ERROR_PRINTER_DRIVER_WARNED                                                      = 3013
	ERROR_PRINTER_DRIVER_BLOCKED                                                     = 3014
	ERROR_PRINTER_DRIVER_PACKAGE_IN_USE                                              = 3015
	ERROR_CORE_DRIVER_PACKAGE_NOT_FOUND                                              = 3016
	ERROR_FAIL_REBOOT_REQUIRED                                                       = 3017
	ERROR_FAIL_REBOOT_INITIATED                                                      = 3018
	ERROR_PRINTER_DRIVER_DOWNLOAD_NEEDED                                             = 3019
	ERROR_PRINT_JOB_RESTART_REQUIRED                                                 = 3020
	ERROR_IO_REISSUE_AS_CACHED                                                       = 3950
	ERROR_WINS_INTERNAL                                                              = 4000
	ERROR_CAN_NOT_DEL_LOCAL_WINS                                                     = 4001
	ERROR_STATIC_INIT                                                                = 4002
	ERROR_INC_BACKUP                                                                 = 4003
	ERROR_FULL_BACKUP                                                                = 4004
	ERROR_REC_NON_EXISTENT                                                           = 4005
	ERROR_RPL_NOT_ALLOWED                                                            = 4006
	PEERDIST_ERROR_CONTENTINFO_VERSION_UNSUPPORTED                                   = 4050
	PEERDIST_ERROR_CANNOT_PARSE_CONTENTINFO                                          = 4051
	PEERDIST_ERROR_MISSING_DATA                                                      = 4052
	PEERDIST_ERROR_NO_MORE                                                           = 4053
	PEERDIST_ERROR_NOT_INITIALIZED                                                   = 4054
	PEERDIST_ERROR_ALREADY_INITIALIZED                                               = 4055
	PEERDIST_ERROR_SHUTDOWN_IN_PROGRESS                                              = 4056
	PEERDIST_ERROR_INVALIDATED                                                       = 4057
	PEERDIST_ERROR_ALREADY_EXISTS                                                    = 4058
	PEERDIST_ERROR_OPERATION_NOTFOUND                                                = 4059
	PEERDIST_ERROR_ALREADY_COMPLETED                                                 = 4060
	PEERDIST_ERROR_OUT_OF_BOUNDS                                                     = 4061
	PEERDIST_ERROR_VERSION_UNSUPPORTED                                               = 4062
	PEERDIST_ERROR_INVALID_CONFIGURATION                                             = 4063
	PEERDIST_ERROR_NOT_LICENSED                                                      = 4064
	PEERDIST_ERROR_SERVICE_UNAVAILABLE                                               = 4065
	ERROR_DHCP_ADDRESS_CONFLICT                                                      = 4100
	ERROR_WMI_GUID_NOT_FOUND                                                         = 4200
	ERROR_WMI_INSTANCE_NOT_FOUND                                                     = 4201
	ERROR_WMI_ITEMID_NOT_FOUND                                                       = 4202
	ERROR_WMI_TRY_AGAIN                                                              = 4203
	ERROR_WMI_DP_NOT_FOUND                                                           = 4204
	ERROR_WMI_UNRESOLVED_INSTANCE_REF                                                = 4205
	ERROR_WMI_ALREADY_ENABLED                                                        = 4206
	ERROR_WMI_GUID_DISCONNECTED                                                      = 4207
	ERROR_WMI_SERVER_UNAVAILABLE                                                     = 4208
	ERROR_WMI_DP_FAILED                                                              = 4209
	ERROR_WMI_INVALID_MOF                                                            = 4210
	ERROR_WMI_INVALID_REGINFO                                                        = 4211
	ERROR_WMI_ALREADY_DISABLED                                                       = 4212
	ERROR_WMI_READ_ONLY                                                              = 4213
	ERROR_WMI_SET_FAILURE                                                            = 4214
	ERROR_INVALID_MEDIA                                                              = 4300
	ERROR_INVALID_LIBRARY                                                            = 4301
	ERROR_INVALID_MEDIA_POOL                                                         = 4302
	ERROR_DRIVE_MEDIA_MISMATCH                                                       = 4303
	ERROR_MEDIA_OFFLINE                                                              = 4304
	ERROR_LIBRARY_OFFLINE                                                            = 4305
	ERROR_EMPTY                                                                      = 4306
	ERROR_NOT_EMPTY                                                                  = 4307
	ERROR_MEDIA_UNAVAILABLE                                                          = 4308
	ERROR_RESOURCE_DISABLED                                                          = 4309
	ERROR_INVALID_CLEANER                                                            = 4310
	ERROR_UNABLE_TO_CLEAN                                                            = 4311
	ERROR_OBJECT_NOT_FOUND                                                           = 4312
	ERROR_DATABASE_FAILURE                                                           = 4313
	ERROR_DATABASE_FULL                                                              = 4314
	ERROR_MEDIA_INCOMPATIBLE                                                         = 4315
	ERROR_RESOURCE_NOT_PRESENT                                                       = 4316
	ERROR_INVALID_OPERATION                                                          = 4317
	ERROR_MEDIA_NOT_AVAILABLE                                                        = 4318
	ERROR_DEVICE_NOT_AVAILABLE                                                       = 4319
	ERROR_REQUEST_REFUSED                                                            = 4320
	ERROR_INVALID_DRIVE_OBJECT                                                       = 4321
	ERROR_LIBRARY_FULL                                                               = 4322
	ERROR_MEDIUM_NOT_ACCESSIBLE                                                      = 4323
	ERROR_UNABLE_TO_LOAD_MEDIUM                                                      = 4324
	ERROR_UNABLE_TO_INVENTORY_DRIVE                                                  = 4325
	ERROR_UNABLE_TO_INVENTORY_SLOT                                                   = 4326
	ERROR_UNABLE_TO_INVENTORY_TRANSPORT                                              = 4327
	ERROR_TRANSPORT_FULL                                                             = 4328
	ERROR_CONTROLLING_IEPORT                                                         = 4329
	ERROR_UNABLE_TO_EJECT_MOUNTED_MEDIA                                              = 4330
	ERROR_CLEANER_SLOT_SET                                                           = 4331
	ERROR_CLEANER_SLOT_NOT_SET                                                       = 4332
	ERROR_CLEANER_CARTRIDGE_SPENT                                                    = 4333
	ERROR_UNEXPECTED_OMID                                                            = 4334
	ERROR_CANT_DELETE_LAST_ITEM                                                      = 4335
	ERROR_MESSAGE_EXCEEDS_MAX_SIZE                                                   = 4336
	ERROR_VOLUME_CONTAINS_SYS_FILES                                                  = 4337
	ERROR_INDIGENOUS_TYPE                                                            = 4338
	ERROR_NO_SUPPORTING_DRIVES                                                       = 4339
	ERROR_CLEANER_CARTRIDGE_INSTALLED                                                = 4340
	ERROR_IEPORT_FULL                                                                = 4341
	ERROR_FILE_OFFLINE                                                               = 4350
	ERROR_REMOTE_STORAGE_NOT_ACTIVE                                                  = 4351
	ERROR_REMOTE_STORAGE_MEDIA_ERROR                                                 = 4352
	ERROR_NOT_A_REPARSE_POINT                                                        = 4390
	ERROR_REPARSE_ATTRIBUTE_CONFLICT                                                 = 4391
	ERROR_INVALID_REPARSE_DATA                                                       = 4392
	ERROR_REPARSE_TAG_INVALID                                                        = 4393
	ERROR_REPARSE_TAG_MISMATCH                                                       = 4394
	ERROR_VOLUME_NOT_SIS_ENABLED                                                     = 4500
	ERROR_DEPENDENT_RESOURCE_EXISTS                                                  = 5001
	ERROR_DEPENDENCY_NOT_FOUND                                                       = 5002
	ERROR_DEPENDENCY_ALREADY_EXISTS                                                  = 5003
	ERROR_RESOURCE_NOT_ONLINE                                                        = 5004
	ERROR_HOST_NODE_NOT_AVAILABLE                                                    = 5005
	ERROR_RESOURCE_NOT_AVAILABLE                                                     = 5006
	ERROR_RESOURCE_NOT_FOUND                                                         = 5007
	ERROR_SHUTDOWN_CLUSTER                                                           = 5008
	ERROR_CANT_EVICT_ACTIVE_NODE                                                     = 5009
	ERROR_OBJECT_ALREADY_EXISTS                                                      = 5010
	ERROR_OBJECT_IN_LIST                                                             = 5011
	ERROR_GROUP_NOT_AVAILABLE                                                        = 5012
	ERROR_GROUP_NOT_FOUND                                                            = 5013
	ERROR_GROUP_NOT_ONLINE                                                           = 5014
	ERROR_HOST_NODE_NOT_RESOURCE_OWNER                                               = 5015
	ERROR_HOST_NODE_NOT_GROUP_OWNER                                                  = 5016
	ERROR_RESMON_CREATE_FAILED                                                       = 5017
	ERROR_RESMON_ONLINE_FAILED                                                       = 5018
	ERROR_RESOURCE_ONLINE                                                            = 5019
	ERROR_QUORUM_RESOURCE                                                            = 5020
	ERROR_NOT_QUORUM_CAPABLE                                                         = 5021
	ERROR_CLUSTER_SHUTTING_DOWN                                                      = 5022
	ERROR_INVALID_STATE                                                              = 5023
	ERROR_RESOURCE_PROPERTIES_STORED                                                 = 5024
	ERROR_NOT_QUORUM_CLASS                                                           = 5025
	ERROR_CORE_RESOURCE                                                              = 5026
	ERROR_QUORUM_RESOURCE_ONLINE_FAILED                                              = 5027
	ERROR_QUORUMLOG_OPEN_FAILED                                                      = 5028
	ERROR_CLUSTERLOG_CORRUPT                                                         = 5029
	ERROR_CLUSTERLOG_RECORD_EXCEEDS_MAXSIZE                                          = 5030
	ERROR_CLUSTERLOG_EXCEEDS_MAXSIZE                                                 = 5031
	ERROR_CLUSTERLOG_CHKPOINT_NOT_FOUND                                              = 5032
	ERROR_CLUSTERLOG_NOT_ENOUGH_SPACE                                                = 5033
	ERROR_QUORUM_OWNER_ALIVE                                                         = 5034
	ERROR_NETWORK_NOT_AVAILABLE                                                      = 5035
	ERROR_NODE_NOT_AVAILABLE                                                         = 5036
	ERROR_ALL_NODES_NOT_AVAILABLE                                                    = 5037
	ERROR_RESOURCE_FAILED                                                            = 5038
	ERROR_CLUSTER_INVALID_NODE                                                       = 5039
	ERROR_CLUSTER_NODE_EXISTS                                                        = 5040
	ERROR_CLUSTER_JOIN_IN_PROGRESS                                                   = 5041
	ERROR_CLUSTER_NODE_NOT_FOUND                                                     = 5042
	ERROR_CLUSTER_LOCAL_NODE_NOT_FOUND                                               = 5043
	ERROR_CLUSTER_NETWORK_EXISTS                                                     = 5044
	ERROR_CLUSTER_NETWORK_NOT_FOUND                                                  = 5045
	ERROR_CLUSTER_NETINTERFACE_EXISTS                                                = 5046
	ERROR_CLUSTER_NETINTERFACE_NOT_FOUND                                             = 5047
	ERROR_CLUSTER_INVALID_REQUEST                                                    = 5048
	ERROR_CLUSTER_INVALID_NETWORK_PROVIDER                                           = 5049
	ERROR_CLUSTER_NODE_DOWN                                                          = 5050
	ERROR_CLUSTER_NODE_UNREACHABLE                                                   = 5051
	ERROR_CLUSTER_NODE_NOT_MEMBER                                                    = 5052
	ERROR_CLUSTER_JOIN_NOT_IN_PROGRESS                                               = 5053
	ERROR_CLUSTER_INVALID_NETWORK                                                    = 5054
	ERROR_CLUSTER_NODE_UP                                                            = 5056
	ERROR_CLUSTER_IPADDR_IN_USE                                                      = 5057
	ERROR_CLUSTER_NODE_NOT_PAUSED                                                    = 5058
	ERROR_CLUSTER_NO_SECURITY_CONTEXT                                                = 5059
	ERROR_CLUSTER_NETWORK_NOT_INTERNAL                                               = 5060
	ERROR_CLUSTER_NODE_ALREADY_UP                                                    = 5061
	ERROR_CLUSTER_NODE_ALREADY_DOWN                                                  = 5062
	ERROR_CLUSTER_NETWORK_ALREADY_ONLINE                                             = 5063
	ERROR_CLUSTER_NETWORK_ALREADY_OFFLINE                                            = 5064
	ERROR_CLUSTER_NODE_ALREADY_MEMBER                                                = 5065
	ERROR_CLUSTER_LAST_INTERNAL_NETWORK                                              = 5066
	ERROR_CLUSTER_NETWORK_HAS_DEPENDENTS                                             = 5067
	ERROR_INVALID_OPERATION_ON_QUORUM                                                = 5068
	ERROR_DEPENDENCY_NOT_ALLOWED                                                     = 5069
	ERROR_CLUSTER_NODE_PAUSED                                                        = 5070
	ERROR_NODE_CANT_HOST_RESOURCE                                                    = 5071
	ERROR_CLUSTER_NODE_NOT_READY                                                     = 5072
	ERROR_CLUSTER_NODE_SHUTTING_DOWN                                                 = 5073
	ERROR_CLUSTER_JOIN_ABORTED                                                       = 5074
	ERROR_CLUSTER_INCOMPATIBLE_VERSIONS                                              = 5075
	ERROR_CLUSTER_MAXNUM_OF_RESOURCES_EXCEEDED                                       = 5076
	ERROR_CLUSTER_SYSTEM_CONFIG_CHANGED                                              = 5077
	ERROR_CLUSTER_RESOURCE_TYPE_NOT_FOUND                                            = 5078
	ERROR_CLUSTER_RESTYPE_NOT_SUPPORTED                                              = 5079
	ERROR_CLUSTER_RESNAME_NOT_FOUND                                                  = 5080
	ERROR_CLUSTER_NO_RPC_PACKAGES_REGISTERED                                         = 5081
	ERROR_CLUSTER_OWNER_NOT_IN_PREFLIST                                              = 5082
	ERROR_CLUSTER_DATABASE_SEQMISMATCH                                               = 5083
	ERROR_RESMON_INVALID_STATE                                                       = 5084
	ERROR_CLUSTER_GUM_NOT_LOCKER                                                     = 5085
	ERROR_QUORUM_DISK_NOT_FOUND                                                      = 5086
	ERROR_DATABASE_BACKUP_CORRUPT                                                    = 5087
	ERROR_CLUSTER_NODE_ALREADY_HAS_DFS_ROOT                                          = 5088
	ERROR_RESOURCE_PROPERTY_UNCHANGEABLE                                             = 5089
	ERROR_CLUSTER_MEMBERSHIP_INVALID_STATE                                           = 5890
	ERROR_CLUSTER_QUORUMLOG_NOT_FOUND                                                = 5891
	ERROR_CLUSTER_MEMBERSHIP_HALT                                                    = 5892
	ERROR_CLUSTER_INSTANCE_ID_MISMATCH                                               = 5893
	ERROR_CLUSTER_NETWORK_NOT_FOUND_FOR_IP                                           = 5894
	ERROR_CLUSTER_PROPERTY_DATA_TYPE_MISMATCH                                        = 5895
	ERROR_CLUSTER_EVICT_WITHOUT_CLEANUP                                              = 5896
	ERROR_CLUSTER_PARAMETER_MISMATCH                                                 = 5897
	ERROR_NODE_CANNOT_BE_CLUSTERED                                                   = 5898
	ERROR_CLUSTER_WRONG_OS_VERSION                                                   = 5899
	ERROR_CLUSTER_CANT_CREATE_DUP_CLUSTER_NAME                                       = 5900
	ERROR_CLUSCFG_ALREADY_COMMITTED                                                  = 5901
	ERROR_CLUSCFG_ROLLBACK_FAILED                                                    = 5902
	ERROR_CLUSCFG_SYSTEM_DISK_DRIVE_LETTER_CONFLICT                                  = 5903
	ERROR_CLUSTER_OLD_VERSION                                                        = 5904
	ERROR_CLUSTER_MISMATCHED_COMPUTER_ACCT_NAME                                      = 5905
	ERROR_CLUSTER_NO_NET_ADAPTERS                                                    = 5906
	ERROR_CLUSTER_POISONED                                                           = 5907
	ERROR_CLUSTER_GROUP_MOVING                                                       = 5908
	ERROR_CLUSTER_RESOURCE_TYPE_BUSY                                                 = 5909
	ERROR_RESOURCE_CALL_TIMED_OUT                                                    = 5910
	ERROR_INVALID_CLUSTER_IPV6_ADDRESS                                               = 5911
	ERROR_CLUSTER_INTERNAL_INVALID_FUNCTION                                          = 5912
	ERROR_CLUSTER_PARAMETER_OUT_OF_BOUNDS                                            = 5913
	ERROR_CLUSTER_PARTIAL_SEND                                                       = 5914
	ERROR_CLUSTER_REGISTRY_INVALID_FUNCTION                                          = 5915
	ERROR_CLUSTER_INVALID_STRING_TERMINATION                                         = 5916
	ERROR_CLUSTER_INVALID_STRING_FORMAT                                              = 5917
	ERROR_CLUSTER_DATABASE_TRANSACTION_IN_PROGRESS                                   = 5918
	ERROR_CLUSTER_DATABASE_TRANSACTION_NOT_IN_PROGRESS                               = 5919
	ERROR_CLUSTER_NULL_DATA                                                          = 5920
	ERROR_CLUSTER_PARTIAL_READ                                                       = 5921
	ERROR_CLUSTER_PARTIAL_WRITE                                                      = 5922
	ERROR_CLUSTER_CANT_DESERIALIZE_DATA                                              = 5923
	ERROR_DEPENDENT_RESOURCE_PROPERTY_CONFLICT                                       = 5924
	ERROR_CLUSTER_NO_QUORUM                                                          = 5925
	ERROR_CLUSTER_INVALID_IPV6_NETWORK                                               = 5926
	ERROR_CLUSTER_INVALID_IPV6_TUNNEL_NETWORK                                        = 5927
	ERROR_QUORUM_NOT_ALLOWED_IN_THIS_GROUP                                           = 5928
	ERROR_DEPENDENCY_TREE_TOO_COMPLEX                                                = 5929
	ERROR_EXCEPTION_IN_RESOURCE_CALL                                                 = 5930
	ERROR_CLUSTER_RHS_FAILED_INITIALIZATION                                          = 5931
	ERROR_CLUSTER_NOT_INSTALLED                                                      = 5932
	ERROR_CLUSTER_RESOURCES_MUST_BE_ONLINE_ON_THE_SAME_NODE                          = 5933
	ERROR_CLUSTER_MAX_NODES_IN_CLUSTER                                               = 5934
	ERROR_CLUSTER_TOO_MANY_NODES                                                     = 5935
	ERROR_CLUSTER_OBJECT_ALREADY_USED                                                = 5936
	ERROR_NONCORE_GROUPS_FOUND                                                       = 5937
	ERROR_FILE_SHARE_RESOURCE_CONFLICT                                               = 5938
	ERROR_CLUSTER_EVICT_INVALID_REQUEST                                              = 5939
	ERROR_CLUSTER_SINGLETON_RESOURCE                                                 = 5940
	ERROR_CLUSTER_GROUP_SINGLETON_RESOURCE                                           = 5941
	ERROR_CLUSTER_RESOURCE_PROVIDER_FAILED                                           = 5942
	ERROR_CLUSTER_RESOURCE_CONFIGURATION_ERROR                                       = 5943
	ERROR_CLUSTER_GROUP_BUSY                                                         = 5944
	ERROR_CLUSTER_NOT_SHARED_VOLUME                                                  = 5945
	ERROR_CLUSTER_INVALID_SECURITY_DESCRIPTOR                                        = 5946
	ERROR_CLUSTER_SHARED_VOLUMES_IN_USE                                              = 5947
	ERROR_CLUSTER_USE_SHARED_VOLUMES_API                                             = 5948
	ERROR_CLUSTER_BACKUP_IN_PROGRESS                                                 = 5949
	ERROR_NON_CSV_PATH                                                               = 5950
	ERROR_CSV_VOLUME_NOT_LOCAL                                                       = 5951
	ERROR_CLUSTER_WATCHDOG_TERMINATING                                               = 5952
	ERROR_ENCRYPTION_FAILED                                                          = 6000
	ERROR_DECRYPTION_FAILED                                                          = 6001
	ERROR_FILE_ENCRYPTED                                                             = 6002
	ERROR_NO_RECOVERY_POLICY                                                         = 6003
	ERROR_NO_EFS                                                                     = 6004
	ERROR_WRONG_EFS                                                                  = 6005
	ERROR_NO_USER_KEYS                                                               = 6006
	ERROR_FILE_NOT_ENCRYPTED                                                         = 6007
	ERROR_NOT_EXPORT_FORMAT                                                          = 6008
	ERROR_FILE_READ_ONLY                                                             = 6009
	ERROR_DIR_EFS_DISALLOWED                                                         = 6010
	ERROR_EFS_SERVER_NOT_TRUSTED                                                     = 6011
	ERROR_BAD_RECOVERY_POLICY                                                        = 6012
	ERROR_EFS_ALG_BLOB_TOO_BIG                                                       = 6013
	ERROR_VOLUME_NOT_SUPPORT_EFS                                                     = 6014
	ERROR_EFS_DISABLED                                                               = 6015
	ERROR_EFS_VERSION_NOT_SUPPORT                                                    = 6016
	ERROR_CS_ENCRYPTION_INVALID_SERVER_RESPONSE                                      = 6017
	ERROR_CS_ENCRYPTION_UNSUPPORTED_SERVER                                           = 6018
	ERROR_CS_ENCRYPTION_EXISTING_ENCRYPTED_FILE                                      = 6019
	ERROR_CS_ENCRYPTION_NEW_ENCRYPTED_FILE                                           = 6020
	ERROR_CS_ENCRYPTION_FILE_NOT_CSE                                                 = 6021
	ERROR_ENCRYPTION_POLICY_DENIES_OPERATION                                         = 6022
	ERROR_NO_BROWSER_SERVERS_FOUND                                                   = 6118
	SCHED_E_SERVICE_NOT_LOCALSYSTEM                                                  = 6200
	ERROR_LOG_SECTOR_INVALID                                                         = 6600
	ERROR_LOG_SECTOR_PARITY_INVALID                                                  = 6601
	ERROR_LOG_SECTOR_REMAPPED                                                        = 6602
	ERROR_LOG_BLOCK_INCOMPLETE                                                       = 6603
	ERROR_LOG_INVALID_RANGE                                                          = 6604
	ERROR_LOG_BLOCKS_EXHAUSTED                                                       = 6605
	ERROR_LOG_READ_CONTEXT_INVALID                                                   = 6606
	ERROR_LOG_RESTART_INVALID                                                        = 6607
	ERROR_LOG_BLOCK_VERSION                                                          = 6608
	ERROR_LOG_BLOCK_INVALID                                                          = 6609
	ERROR_LOG_READ_MODE_INVALID                                                      = 6610
	ERROR_LOG_NO_RESTART                                                             = 6611
	ERROR_LOG_METADATA_CORRUPT                                                       = 6612
	ERROR_LOG_METADATA_INVALID                                                       = 6613
	ERROR_LOG_METADATA_INCONSISTENT                                                  = 6614
	ERROR_LOG_RESERVATION_INVALID                                                    = 6615
	ERROR_LOG_CANT_DELETE                                                            = 6616
	ERROR_LOG_CONTAINER_LIMIT_EXCEEDED                                               = 6617
	ERROR_LOG_START_OF_LOG                                                           = 6618
	ERROR_LOG_POLICY_ALREADY_INSTALLED                                               = 6619
	ERROR_LOG_POLICY_NOT_INSTALLED                                                   = 6620
	ERROR_LOG_POLICY_INVALID                                                         = 6621
	ERROR_LOG_POLICY_CONFLICT                                                        = 6622
	ERROR_LOG_PINNED_ARCHIVE_TAIL                                                    = 6623
	ERROR_LOG_RECORD_NONEXISTENT                                                     = 6624
	ERROR_LOG_RECORDS_RESERVED_INVALID                                               = 6625
	ERROR_LOG_SPACE_RESERVED_INVALID                                                 = 6626
	ERROR_LOG_TAIL_INVALID                                                           = 6627
	ERROR_LOG_FULL                                                                   = 6628
	ERROR_COULD_NOT_RESIZE_LOG                                                       = 6629
	ERROR_LOG_MULTIPLEXED                                                            = 6630
	ERROR_LOG_DEDICATED                                                              = 6631
	ERROR_LOG_ARCHIVE_NOT_IN_PROGRESS                                                = 6632
	ERROR_LOG_ARCHIVE_IN_PROGRESS                                                    = 6633
	ERROR_LOG_EPHEMERAL                                                              = 6634
	ERROR_LOG_NOT_ENOUGH_CONTAINERS                                                  = 6635
	ERROR_LOG_CLIENT_ALREADY_REGISTERED                                              = 6636
	ERROR_LOG_CLIENT_NOT_REGISTERED                                                  = 6637
	ERROR_LOG_FULL_HANDLER_IN_PROGRESS                                               = 6638
	ERROR_LOG_CONTAINER_READ_FAILED                                                  = 6639
	ERROR_LOG_CONTAINER_WRITE_FAILED                                                 = 6640
	ERROR_LOG_CONTAINER_OPEN_FAILED                                                  = 6641
	ERROR_LOG_CONTAINER_STATE_INVALID                                                = 6642
	ERROR_LOG_STATE_INVALID                                                          = 6643
	ERROR_LOG_PINNED                                                                 = 6644
	ERROR_LOG_METADATA_FLUSH_FAILED                                                  = 6645
	ERROR_LOG_INCONSISTENT_SECURITY                                                  = 6646
	ERROR_LOG_APPENDED_FLUSH_FAILED                                                  = 6647
	ERROR_LOG_PINNED_RESERVATION                                                     = 6648
	ERROR_INVALID_TRANSACTION                                                        = 6700
	ERROR_TRANSACTION_NOT_ACTIVE                                                     = 6701
	ERROR_TRANSACTION_REQUEST_NOT_VALID                                              = 6702
	ERROR_TRANSACTION_NOT_REQUESTED                                                  = 6703
	ERROR_TRANSACTION_ALREADY_ABORTED                                                = 6704
	ERROR_TRANSACTION_ALREADY_COMMITTED                                              = 6705
	ERROR_TM_INITIALIZATION_FAILED                                                   = 6706
	ERROR_RESOURCEMANAGER_READ_ONLY                                                  = 6707
	ERROR_TRANSACTION_NOT_JOINED                                                     = 6708
	ERROR_TRANSACTION_SUPERIOR_EXISTS                                                = 6709
	ERROR_CRM_PROTOCOL_ALREADY_EXISTS                                                = 6710
	ERROR_TRANSACTION_PROPAGATION_FAILED                                             = 6711
	ERROR_CRM_PROTOCOL_NOT_FOUND                                                     = 6712
	ERROR_TRANSACTION_INVALID_MARSHALL_BUFFER                                        = 6713
	ERROR_CURRENT_TRANSACTION_NOT_VALID                                              = 6714
	ERROR_TRANSACTION_NOT_FOUND                                                      = 6715
	ERROR_RESOURCEMANAGER_NOT_FOUND                                                  = 6716
	ERROR_ENLISTMENT_NOT_FOUND                                                       = 6717
	ERROR_TRANSACTIONMANAGER_NOT_FOUND                                               = 6718
	ERROR_TRANSACTIONMANAGER_NOT_ONLINE                                              = 6719
	ERROR_TRANSACTIONMANAGER_RECOVERY_NAME_COLLISION                                 = 6720
	ERROR_TRANSACTION_NOT_ROOT                                                       = 6721
	ERROR_TRANSACTION_OBJECT_EXPIRED                                                 = 6722
	ERROR_TRANSACTION_RESPONSE_NOT_ENLISTED                                          = 6723
	ERROR_TRANSACTION_RECORD_TOO_LONG                                                = 6724
	ERROR_IMPLICIT_TRANSACTION_NOT_SUPPORTED                                         = 6725
	ERROR_TRANSACTION_INTEGRITY_VIOLATED                                             = 6726
	ERROR_TRANSACTIONMANAGER_IDENTITY_MISMATCH                                       = 6727
	ERROR_RM_CANNOT_BE_FROZEN_FOR_SNAPSHOT                                           = 6728
	ERROR_TRANSACTION_MUST_WRITETHROUGH                                              = 6729
	ERROR_TRANSACTION_NO_SUPERIOR                                                    = 6730
	ERROR_HEURISTIC_DAMAGE_POSSIBLE                                                  = 6731
	ERROR_TRANSACTIONAL_CONFLICT                                                     = 6800
	ERROR_RM_NOT_ACTIVE                                                              = 6801
	ERROR_RM_METADATA_CORRUPT                                                        = 6802
	ERROR_DIRECTORY_NOT_RM                                                           = 6803
	ERROR_TRANSACTIONS_UNSUPPORTED_REMOTE                                            = 6805
	ERROR_LOG_RESIZE_INVALID_SIZE                                                    = 6806
	ERROR_OBJECT_NO_LONGER_EXISTS                                                    = 6807
	ERROR_STREAM_MINIVERSION_NOT_FOUND                                               = 6808
	ERROR_STREAM_MINIVERSION_NOT_VALID                                               = 6809
	ERROR_MINIVERSION_INACCESSIBLE_FROM_SPECIFIED_TRANSACTION                        = 6810
	ERROR_CANT_OPEN_MINIVERSION_WITH_MODIFY_INTENT                                   = 6811
	ERROR_CANT_CREATE_MORE_STREAM_MINIVERSIONS                                       = 6812
	ERROR_REMOTE_FILE_VERSION_MISMATCH                                               = 6814
	ERROR_HANDLE_NO_LONGER_VALID                                                     = 6815
	ERROR_NO_TXF_METADATA                                                            = 6816
	ERROR_LOG_CORRUPTION_DETECTED                                                    = 6817
	ERROR_CANT_RECOVER_WITH_HANDLE_OPEN                                              = 6818
	ERROR_RM_DISCONNECTED                                                            = 6819
	ERROR_ENLISTMENT_NOT_SUPERIOR                                                    = 6820
	ERROR_RECOVERY_NOT_NEEDED                                                        = 6821
	ERROR_RM_ALREADY_STARTED                                                         = 6822
	ERROR_FILE_IDENTITY_NOT_PERSISTENT                                               = 6823
	ERROR_CANT_BREAK_TRANSACTIONAL_DEPENDENCY                                        = 6824
	ERROR_CANT_CROSS_RM_BOUNDARY                                                     = 6825
	ERROR_TXF_DIR_NOT_EMPTY                                                          = 6826
	ERROR_INDOUBT_TRANSACTIONS_EXIST                                                 = 6827
	ERROR_TM_VOLATILE                                                                = 6828
	ERROR_ROLLBACK_TIMER_EXPIRED                                                     = 6829
	ERROR_TXF_ATTRIBUTE_CORRUPT                                                      = 6830
	ERROR_EFS_NOT_ALLOWED_IN_TRANSACTION                                             = 6831
	ERROR_TRANSACTIONAL_OPEN_NOT_ALLOWED                                             = 6832
	ERROR_LOG_GROWTH_FAILED                                                          = 6833
	ERROR_TRANSACTED_MAPPING_UNSUPPORTED_REMOTE                                      = 6834
	ERROR_TXF_METADATA_ALREADY_PRESENT                                               = 6835
	ERROR_TRANSACTION_SCOPE_CALLBACKS_NOT_SET                                        = 6836
	ERROR_TRANSACTION_REQUIRED_PROMOTION                                             = 6837
	ERROR_CANNOT_EXECUTE_FILE_IN_TRANSACTION                                         = 6838
	ERROR_TRANSACTIONS_NOT_FROZEN                                                    = 6839
	ERROR_TRANSACTION_FREEZE_IN_PROGRESS                                             = 6840
	ERROR_NOT_SNAPSHOT_VOLUME                                                        = 6841
	ERROR_NO_SAVEPOINT_WITH_OPEN_FILES                                               = 6842
	ERROR_DATA_LOST_REPAIR                                                           = 6843
	ERROR_SPARSE_NOT_ALLOWED_IN_TRANSACTION                                          = 6844
	ERROR_TM_IDENTITY_MISMATCH                                                       = 6845
	ERROR_FLOATED_SECTION                                                            = 6846
	ERROR_CANNOT_ACCEPT_TRANSACTED_WORK                                              = 6847
	ERROR_CANNOT_ABORT_TRANSACTIONS                                                  = 6848
	ERROR_BAD_CLUSTERS                                                               = 6849
	ERROR_COMPRESSION_NOT_ALLOWED_IN_TRANSACTION                                     = 6850
	ERROR_VOLUME_DIRTY                                                               = 6851
	ERROR_NO_LINK_TRACKING_IN_TRANSACTION                                            = 6852
	ERROR_OPERATION_NOT_SUPPORTED_IN_TRANSACTION                                     = 6853
	ERROR_EXPIRED_HANDLE                                                             = 6854
	ERROR_TRANSACTION_NOT_ENLISTED                                                   = 6855
	ERROR_CTX_WINSTATION_NAME_INVALID                                                = 7001
	ERROR_CTX_INVALID_PD                                                             = 7002
	ERROR_CTX_PD_NOT_FOUND                                                           = 7003
	ERROR_CTX_WD_NOT_FOUND                                                           = 7004
	ERROR_CTX_CANNOT_MAKE_EVENTLOG_ENTRY                                             = 7005
	ERROR_CTX_SERVICE_NAME_COLLISION                                                 = 7006
	ERROR_CTX_CLOSE_PENDING                                                          = 7007
	ERROR_CTX_NO_OUTBUF                                                              = 7008
	ERROR_CTX_MODEM_INF_NOT_FOUND                                                    = 7009
	ERROR_CTX_INVALID_MODEMNAME                                                      = 7010
	ERROR_CTX_MODEM_RESPONSE_ERROR                                                   = 7011
	ERROR_CTX_MODEM_RESPONSE_TIMEOUT                                                 = 7012
	ERROR_CTX_MODEM_RESPONSE_NO_CARRIER                                              = 7013
	ERROR_CTX_MODEM_RESPONSE_NO_DIALTONE                                             = 7014
	ERROR_CTX_MODEM_RESPONSE_BUSY                                                    = 7015
	ERROR_CTX_MODEM_RESPONSE_VOICE                                                   = 7016
	ERROR_CTX_TD_ERROR                                                               = 7017
	ERROR_CTX_WINSTATION_NOT_FOUND                                                   = 7022
	ERROR_CTX_WINSTATION_ALREADY_EXISTS                                              = 7023
	ERROR_CTX_WINSTATION_BUSY                                                        = 7024
	ERROR_CTX_BAD_VIDEO_MODE                                                         = 7025
	ERROR_CTX_GRAPHICS_INVALID                                                       = 7035
	ERROR_CTX_LOGON_DISABLED                                                         = 7037
	ERROR_CTX_NOT_CONSOLE                                                            = 7038
	ERROR_CTX_CLIENT_QUERY_TIMEOUT                                                   = 7040
	ERROR_CTX_CONSOLE_DISCONNECT                                                     = 7041
	ERROR_CTX_CONSOLE_CONNECT                                                        = 7042
	ERROR_CTX_SHADOW_DENIED                                                          = 7044
	ERROR_CTX_WINSTATION_ACCESS_DENIED                                               = 7045
	ERROR_CTX_INVALID_WD                                                             = 7049
	ERROR_CTX_SHADOW_INVALID                                                         = 7050
	ERROR_CTX_SHADOW_DISABLED                                                        = 7051
	ERROR_CTX_CLIENT_LICENSE_IN_USE                                                  = 7052
	ERROR_CTX_CLIENT_LICENSE_NOT_SET                                                 = 7053
	ERROR_CTX_LICENSE_NOT_AVAILABLE                                                  = 7054
	ERROR_CTX_LICENSE_CLIENT_INVALID                                                 = 7055
	ERROR_CTX_LICENSE_EXPIRED                                                        = 7056
	ERROR_CTX_SHADOW_NOT_RUNNING                                                     = 7057
	ERROR_CTX_SHADOW_ENDED_BY_MODE_CHANGE                                            = 7058
	ERROR_ACTIVATION_COUNT_EXCEEDED                                                  = 7059
	ERROR_CTX_WINSTATIONS_DISABLED                                                   = 7060
	ERROR_CTX_ENCRYPTION_LEVEL_REQUIRED                                              = 7061
	ERROR_CTX_SESSION_IN_USE                                                         = 7062
	ERROR_CTX_NO_FORCE_LOGOFF                                                        = 7063
	ERROR_CTX_ACCOUNT_RESTRICTION                                                    = 7064
	ERROR_RDP_PROTOCOL_ERROR                                                         = 7065
	ERROR_CTX_CDM_CONNECT                                                            = 7066
	ERROR_CTX_CDM_DISCONNECT                                                         = 7067
	ERROR_CTX_SECURITY_LAYER_ERROR                                                   = 7068
	ERROR_TS_INCOMPATIBLE_SESSIONS                                                   = 7069
	ERROR_TS_VIDEO_SUBSYSTEM_ERROR                                                   = 7070
	FRS_ERR_INVALID_API_SEQUENCE                                                     = 8001
	FRS_ERR_STARTING_SERVICE                                                         = 8002
	FRS_ERR_STOPPING_SERVICE                                                         = 8003
	FRS_ERR_INTERNAL_API                                                             = 8004
	FRS_ERR_INTERNAL                                                                 = 8005
	FRS_ERR_SERVICE_COMM                                                             = 8006
	FRS_ERR_INSUFFICIENT_PRIV                                                        = 8007
	FRS_ERR_AUTHENTICATION                                                           = 8008
	FRS_ERR_PARENT_INSUFFICIENT_PRIV                                                 = 8009
	FRS_ERR_PARENT_AUTHENTICATION                                                    = 8010
	FRS_ERR_CHILD_TO_PARENT_COMM                                                     = 8011
	FRS_ERR_PARENT_TO_CHILD_COMM                                                     = 8012
	FRS_ERR_SYSVOL_POPULATE                                                          = 8013
	FRS_ERR_SYSVOL_POPULATE_TIMEOUT                                                  = 8014
	FRS_ERR_SYSVOL_IS_BUSY                                                           = 8015
	FRS_ERR_SYSVOL_DEMOTE                                                            = 8016
	FRS_ERR_INVALID_SERVICE_PARAMETER                                                = 8017
	ERROR_DS_NOT_INSTALLED                                                           = 8200
	ERROR_DS_MEMBERSHIP_EVALUATED_LOCALLY                                            = 8201
	ERROR_DS_NO_ATTRIBUTE_OR_VALUE                                                   = 8202
	ERROR_DS_INVALID_ATTRIBUTE_SYNTAX                                                = 8203
	ERROR_DS_ATTRIBUTE_TYPE_UNDEFINED                                                = 8204
	ERROR_DS_ATTRIBUTE_OR_VALUE_EXISTS                                               = 8205
	ERROR_DS_BUSY                                                                    = 8206
	ERROR_DS_UNAVAILABLE                                                             = 8207
	ERROR_DS_NO_RIDS_ALLOCATED                                                       = 8208
	ERROR_DS_NO_MORE_RIDS                                                            = 8209
	ERROR_DS_INCORRECT_ROLE_OWNER                                                    = 8210
	ERROR_DS_RIDMGR_INIT_ERROR                                                       = 8211
	ERROR_DS_OBJ_CLASS_VIOLATION                                                     = 8212
	ERROR_DS_CANT_ON_NON_LEAF                                                        = 8213
	ERROR_DS_CANT_ON_RDN                                                             = 8214
	ERROR_DS_CANT_MOD_OBJ_CLASS                                                      = 8215
	ERROR_DS_CROSS_DOM_MOVE_ERROR                                                    = 8216
	ERROR_DS_GC_NOT_AVAILABLE                                                        = 8217
	ERROR_SHARED_POLICY                                                              = 8218
	ERROR_POLICY_OBJECT_NOT_FOUND                                                    = 8219
	ERROR_POLICY_ONLY_IN_DS                                                          = 8220
	ERROR_PROMOTION_ACTIVE                                                           = 8221
	ERROR_NO_PROMOTION_ACTIVE                                                        = 8222
	ERROR_DS_OPERATIONS_ERROR                                                        = 8224
	ERROR_DS_PROTOCOL_ERROR                                                          = 8225
	ERROR_DS_TIMELIMIT_EXCEEDED                                                      = 8226
	ERROR_DS_SIZELIMIT_EXCEEDED                                                      = 8227
	ERROR_DS_ADMIN_LIMIT_EXCEEDED                                                    = 8228
	ERROR_DS_COMPARE_FALSE                                                           = 8229
	ERROR_DS_COMPARE_TRUE                                                            = 8230
	ERROR_DS_AUTH_METHOD_NOT_SUPPORTED                                               = 8231
	ERROR_DS_STRONG_AUTH_REQUIRED                                                    = 8232
	ERROR_DS_INAPPROPRIATE_AUTH                                                      = 8233
	ERROR_DS_AUTH_UNKNOWN                                                            = 8234
	ERROR_DS_REFERRAL                                                                = 8235
	ERROR_DS_UNAVAILABLE_CRIT_EXTENSION                                              = 8236
	ERROR_DS_CONFIDENTIALITY_REQUIRED                                                = 8237
	ERROR_DS_INAPPROPRIATE_MATCHING                                                  = 8238
	ERROR_DS_CONSTRAINT_VIOLATION                                                    = 8239
	ERROR_DS_NO_SUCH_OBJECT                                                          = 8240
	ERROR_DS_ALIAS_PROBLEM                                                           = 8241
	ERROR_DS_INVALID_DN_SYNTAX                                                       = 8242
	ERROR_DS_IS_LEAF                                                                 = 8243
	ERROR_DS_ALIAS_DEREF_PROBLEM                                                     = 8244
	ERROR_DS_UNWILLING_TO_PERFORM                                                    = 8245
	ERROR_DS_LOOP_DETECT                                                             = 8246
	ERROR_DS_NAMING_VIOLATION                                                        = 8247
	ERROR_DS_OBJECT_RESULTS_TOO_LARGE                                                = 8248
	ERROR_DS_AFFECTS_MULTIPLE_DSAS                                                   = 8249
	ERROR_DS_SERVER_DOWN                                                             = 8250
	ERROR_DS_LOCAL_ERROR                                                             = 8251
	ERROR_DS_ENCODING_ERROR                                                          = 8252
	ERROR_DS_DECODING_ERROR                                                          = 8253
	ERROR_DS_FILTER_UNKNOWN                                                          = 8254
	ERROR_DS_PARAM_ERROR                                                             = 8255
	ERROR_DS_NOT_SUPPORTED                                                           = 8256
	ERROR_DS_NO_RESULTS_RETURNED                                                     = 8257
	ERROR_DS_CONTROL_NOT_FOUND                                                       = 8258
	ERROR_DS_CLIENT_LOOP                                                             = 8259
	ERROR_DS_REFERRAL_LIMIT_EXCEEDED                                                 = 8260
	ERROR_DS_SORT_CONTROL_MISSING                                                    = 8261
	ERROR_DS_OFFSET_RANGE_ERROR                                                      = 8262
	ERROR_DS_ROOT_MUST_BE_NC                                                         = 8301
	ERROR_DS_ADD_REPLICA_INHIBITED                                                   = 8302
	ERROR_DS_ATT_NOT_DEF_IN_SCHEMA                                                   = 8303
	ERROR_DS_MAX_OBJ_SIZE_EXCEEDED                                                   = 8304
	ERROR_DS_OBJ_STRING_NAME_EXISTS                                                  = 8305
	ERROR_DS_NO_RDN_DEFINED_IN_SCHEMA                                                = 8306
	ERROR_DS_RDN_DOESNT_MATCH_SCHEMA                                                 = 8307
	ERROR_DS_NO_REQUESTED_ATTS_FOUND                                                 = 8308
	ERROR_DS_USER_BUFFER_TO_SMALL                                                    = 8309
	ERROR_DS_ATT_IS_NOT_ON_OBJ                                                       = 8310
	ERROR_DS_ILLEGAL_MOD_OPERATION                                                   = 8311
	ERROR_DS_OBJ_TOO_LARGE                                                           = 8312
	ERROR_DS_BAD_INSTANCE_TYPE                                                       = 8313
	ERROR_DS_MASTERDSA_REQUIRED                                                      = 8314
	ERROR_DS_OBJECT_CLASS_REQUIRED                                                   = 8315
	ERROR_DS_MISSING_REQUIRED_ATT                                                    = 8316
	ERROR_DS_ATT_NOT_DEF_FOR_CLASS                                                   = 8317
	ERROR_DS_ATT_ALREADY_EXISTS                                                      = 8318
	ERROR_DS_CANT_ADD_ATT_VALUES                                                     = 8320
	ERROR_DS_SINGLE_VALUE_CONSTRAINT                                                 = 8321
	ERROR_DS_RANGE_CONSTRAINT                                                        = 8322
	ERROR_DS_ATT_VAL_ALREADY_EXISTS                                                  = 8323
	ERROR_DS_CANT_REM_MISSING_ATT                                                    = 8324
	ERROR_DS_CANT_REM_MISSING_ATT_VAL                                                = 8325
	ERROR_DS_ROOT_CANT_BE_SUBREF                                                     = 8326
	ERROR_DS_NO_CHAINING                                                             = 8327
	ERROR_DS_NO_CHAINED_EVAL                                                         = 8328
	ERROR_DS_NO_PARENT_OBJECT                                                        = 8329
	ERROR_DS_PARENT_IS_AN_ALIAS                                                      = 8330
	ERROR_DS_CANT_MIX_MASTER_AND_REPS                                                = 8331
	ERROR_DS_CHILDREN_EXIST                                                          = 8332
	ERROR_DS_OBJ_NOT_FOUND                                                           = 8333
	ERROR_DS_ALIASED_OBJ_MISSING                                                     = 8334
	ERROR_DS_BAD_NAME_SYNTAX                                                         = 8335
	ERROR_DS_ALIAS_POINTS_TO_ALIAS                                                   = 8336
	ERROR_DS_CANT_DEREF_ALIAS                                                        = 8337
	ERROR_DS_OUT_OF_SCOPE                                                            = 8338
	ERROR_DS_OBJECT_BEING_REMOVED                                                    = 8339
	ERROR_DS_CANT_DELETE_DSA_OBJ                                                     = 8340
	ERROR_DS_GENERIC_ERROR                                                           = 8341
	ERROR_DS_DSA_MUST_BE_INT_MASTER                                                  = 8342
	ERROR_DS_CLASS_NOT_DSA                                                           = 8343
	ERROR_DS_INSUFF_ACCESS_RIGHTS                                                    = 8344
	ERROR_DS_ILLEGAL_SUPERIOR                                                        = 8345
	ERROR_DS_ATTRIBUTE_OWNED_BY_SAM                                                  = 8346
	ERROR_DS_NAME_TOO_MANY_PARTS                                                     = 8347
	ERROR_DS_NAME_TOO_LONG                                                           = 8348
	ERROR_DS_NAME_VALUE_TOO_LONG                                                     = 8349
	ERROR_DS_NAME_UNPARSEABLE                                                        = 8350
	ERROR_DS_NAME_TYPE_UNKNOWN                                                       = 8351
	ERROR_DS_NOT_AN_OBJECT                                                           = 8352
	ERROR_DS_SEC_DESC_TOO_SHORT                                                      = 8353
	ERROR_DS_SEC_DESC_INVALID                                                        = 8354
	ERROR_DS_NO_DELETED_NAME                                                         = 8355
	ERROR_DS_SUBREF_MUST_HAVE_PARENT                                                 = 8356
	ERROR_DS_NCNAME_MUST_BE_NC                                                       = 8357
	ERROR_DS_CANT_ADD_SYSTEM_ONLY                                                    = 8358
	ERROR_DS_CLASS_MUST_BE_CONCRETE                                                  = 8359
	ERROR_DS_INVALID_DMD                                                             = 8360
	ERROR_DS_OBJ_GUID_EXISTS                                                         = 8361
	ERROR_DS_NOT_ON_BACKLINK                                                         = 8362
	ERROR_DS_NO_CROSSREF_FOR_NC                                                      = 8363
	ERROR_DS_SHUTTING_DOWN                                                           = 8364
	ERROR_DS_UNKNOWN_OPERATION                                                       = 8365
	ERROR_DS_INVALID_ROLE_OWNER                                                      = 8366
	ERROR_DS_COULDNT_CONTACT_FSMO                                                    = 8367
	ERROR_DS_CROSS_NC_DN_RENAME                                                      = 8368
	ERROR_DS_CANT_MOD_SYSTEM_ONLY                                                    = 8369
	ERROR_DS_REPLICATOR_ONLY                                                         = 8370
	ERROR_DS_OBJ_CLASS_NOT_DEFINED                                                   = 8371
	ERROR_DS_OBJ_CLASS_NOT_SUBCLASS                                                  = 8372
	ERROR_DS_NAME_REFERENCE_INVALID                                                  = 8373
	ERROR_DS_CROSS_REF_EXISTS                                                        = 8374
	ERROR_DS_CANT_DEL_MASTER_CROSSREF                                                = 8375
	ERROR_DS_SUBTREE_NOTIFY_NOT_NC_HEAD                                              = 8376
	ERROR_DS_NOTIFY_FILTER_TOO_COMPLEX                                               = 8377
	ERROR_DS_DUP_RDN                                                                 = 8378
	ERROR_DS_DUP_OID                                                                 = 8379
	ERROR_DS_DUP_MAPI_ID                                                             = 8380
	ERROR_DS_DUP_SCHEMA_ID_GUID                                                      = 8381
	ERROR_DS_DUP_LDAP_DISPLAY_NAME                                                   = 8382
	ERROR_DS_SEMANTIC_ATT_TEST                                                       = 8383
	ERROR_DS_SYNTAX_MISMATCH                                                         = 8384
	ERROR_DS_EXISTS_IN_MUST_HAVE                                                     = 8385
	ERROR_DS_EXISTS_IN_MAY_HAVE                                                      = 8386
	ERROR_DS_NONEXISTENT_MAY_HAVE                                                    = 8387
	ERROR_DS_NONEXISTENT_MUST_HAVE                                                   = 8388
	ERROR_DS_AUX_CLS_TEST_FAIL                                                       = 8389
	ERROR_DS_NONEXISTENT_POSS_SUP                                                    = 8390
	ERROR_DS_SUB_CLS_TEST_FAIL                                                       = 8391
	ERROR_DS_BAD_RDN_ATT_ID_SYNTAX                                                   = 8392
	ERROR_DS_EXISTS_IN_AUX_CLS                                                       = 8393
	ERROR_DS_EXISTS_IN_SUB_CLS                                                       = 8394
	ERROR_DS_EXISTS_IN_POSS_SUP                                                      = 8395
	ERROR_DS_RECALCSCHEMA_FAILED                                                     = 8396
	ERROR_DS_TREE_DELETE_NOT_FINISHED                                                = 8397
	ERROR_DS_CANT_DELETE                                                             = 8398
	ERROR_DS_ATT_SCHEMA_REQ_ID                                                       = 8399
	ERROR_DS_BAD_ATT_SCHEMA_SYNTAX                                                   = 8400
	ERROR_DS_CANT_CACHE_ATT                                                          = 8401
	ERROR_DS_CANT_CACHE_CLASS                                                        = 8402
	ERROR_DS_CANT_REMOVE_ATT_CACHE                                                   = 8403
	ERROR_DS_CANT_REMOVE_CLASS_CACHE                                                 = 8404
	ERROR_DS_CANT_RETRIEVE_DN                                                        = 8405
	ERROR_DS_MISSING_SUPREF                                                          = 8406
	ERROR_DS_CANT_RETRIEVE_INSTANCE                                                  = 8407
	ERROR_DS_CODE_INCONSISTENCY                                                      = 8408
	ERROR_DS_DATABASE_ERROR                                                          = 8409
	ERROR_DS_GOVERNSID_MISSING                                                       = 8410
	ERROR_DS_MISSING_EXPECTED_ATT                                                    = 8411
	ERROR_DS_NCNAME_MISSING_CR_REF                                                   = 8412
	ERROR_DS_SECURITY_CHECKING_ERROR                                                 = 8413
	ERROR_DS_SCHEMA_NOT_LOADED                                                       = 8414
	ERROR_DS_SCHEMA_ALLOC_FAILED                                                     = 8415
	ERROR_DS_ATT_SCHEMA_REQ_SYNTAX                                                   = 8416
	ERROR_DS_GCVERIFY_ERROR                                                          = 8417
	ERROR_DS_DRA_SCHEMA_MISMATCH                                                     = 8418
	ERROR_DS_CANT_FIND_DSA_OBJ                                                       = 8419
	ERROR_DS_CANT_FIND_EXPECTED_NC                                                   = 8420
	ERROR_DS_CANT_FIND_NC_IN_CACHE                                                   = 8421
	ERROR_DS_CANT_RETRIEVE_CHILD                                                     = 8422
	ERROR_DS_SECURITY_ILLEGAL_MODIFY                                                 = 8423
	ERROR_DS_CANT_REPLACE_HIDDEN_REC                                                 = 8424
	ERROR_DS_BAD_HIERARCHY_FILE                                                      = 8425
	ERROR_DS_BUILD_HIERARCHY_TABLE_FAILED                                            = 8426
	ERROR_DS_CONFIG_PARAM_MISSING                                                    = 8427
	ERROR_DS_COUNTING_AB_INDICES_FAILED                                              = 8428
	ERROR_DS_HIERARCHY_TABLE_MALLOC_FAILED                                           = 8429
	ERROR_DS_INTERNAL_FAILURE                                                        = 8430
	ERROR_DS_UNKNOWN_ERROR                                                           = 8431
	ERROR_DS_ROOT_REQUIRES_CLASS_TOP                                                 = 8432
	ERROR_DS_REFUSING_FSMO_ROLES                                                     = 8433
	ERROR_DS_MISSING_FSMO_SETTINGS                                                   = 8434
	ERROR_DS_UNABLE_TO_SURRENDER_ROLES                                               = 8435
	ERROR_DS_DRA_GENERIC                                                             = 8436
	ERROR_DS_DRA_INVALID_PARAMETER                                                   = 8437
	ERROR_DS_DRA_BUSY                                                                = 8438
	ERROR_DS_DRA_BAD_DN                                                              = 8439
	ERROR_DS_DRA_BAD_NC                                                              = 8440
	ERROR_DS_DRA_DN_EXISTS                                                           = 8441
	ERROR_DS_DRA_INTERNAL_ERROR                                                      = 8442
	ERROR_DS_DRA_INCONSISTENT_DIT                                                    = 8443
	ERROR_DS_DRA_CONNECTION_FAILED                                                   = 8444
	ERROR_DS_DRA_BAD_INSTANCE_TYPE                                                   = 8445
	ERROR_DS_DRA_OUT_OF_MEM                                                          = 8446
	ERROR_DS_DRA_MAIL_PROBLEM                                                        = 8447
	ERROR_DS_DRA_REF_ALREADY_EXISTS                                                  = 8448
	ERROR_DS_DRA_REF_NOT_FOUND                                                       = 8449
	ERROR_DS_DRA_OBJ_IS_REP_SOURCE                                                   = 8450
	ERROR_DS_DRA_DB_ERROR                                                            = 8451
	ERROR_DS_DRA_NO_REPLICA                                                          = 8452
	ERROR_DS_DRA_ACCESS_DENIED                                                       = 8453
	ERROR_DS_DRA_NOT_SUPPORTED                                                       = 8454
	ERROR_DS_DRA_RPC_CANCELLED                                                       = 8455
	ERROR_DS_DRA_SOURCE_DISABLED                                                     = 8456
	ERROR_DS_DRA_SINK_DISABLED                                                       = 8457
	ERROR_DS_DRA_NAME_COLLISION                                                      = 8458
	ERROR_DS_DRA_SOURCE_REINSTALLED                                                  = 8459
	ERROR_DS_DRA_MISSING_PARENT                                                      = 8460
	ERROR_DS_DRA_PREEMPTED                                                           = 8461
	ERROR_DS_DRA_ABANDON_SYNC                                                        = 8462
	ERROR_DS_DRA_SHUTDOWN                                                            = 8463
	ERROR_DS_DRA_INCOMPATIBLE_PARTIAL_SET                                            = 8464
	ERROR_DS_DRA_SOURCE_IS_PARTIAL_REPLICA                                           = 8465
	ERROR_DS_DRA_EXTN_CONNECTION_FAILED                                              = 8466
	ERROR_DS_INSTALL_SCHEMA_MISMATCH                                                 = 8467
	ERROR_DS_DUP_LINK_ID                                                             = 8468
	ERROR_DS_NAME_ERROR_RESOLVING                                                    = 8469
	ERROR_DS_NAME_ERROR_NOT_FOUND                                                    = 8470
	ERROR_DS_NAME_ERROR_NOT_UNIQUE                                                   = 8471
	ERROR_DS_NAME_ERROR_NO_MAPPING                                                   = 8472
	ERROR_DS_NAME_ERROR_DOMAIN_ONLY                                                  = 8473
	ERROR_DS_NAME_ERROR_NO_SYNTACTICAL_MAPPING                                       = 8474
	ERROR_DS_CONSTRUCTED_ATT_MOD                                                     = 8475
	ERROR_DS_WRONG_OM_OBJ_CLASS                                                      = 8476
	ERROR_DS_DRA_REPL_PENDING                                                        = 8477
	ERROR_DS_DS_REQUIRED                                                             = 8478
	ERROR_DS_INVALID_LDAP_DISPLAY_NAME                                               = 8479
	ERROR_DS_NON_BASE_SEARCH                                                         = 8480
	ERROR_DS_CANT_RETRIEVE_ATTS                                                      = 8481
	ERROR_DS_BACKLINK_WITHOUT_LINK                                                   = 8482
	ERROR_DS_EPOCH_MISMATCH                                                          = 8483
	ERROR_DS_SRC_NAME_MISMATCH                                                       = 8484
	ERROR_DS_SRC_AND_DST_NC_IDENTICAL                                                = 8485
	ERROR_DS_DST_NC_MISMATCH                                                         = 8486
	ERROR_DS_NOT_AUTHORITIVE_FOR_DST_NC                                              = 8487
	ERROR_DS_SRC_GUID_MISMATCH                                                       = 8488
	ERROR_DS_CANT_MOVE_DELETED_OBJECT                                                = 8489
	ERROR_DS_PDC_OPERATION_IN_PROGRESS                                               = 8490
	ERROR_DS_CROSS_DOMAIN_CLEANUP_REQD                                               = 8491
	ERROR_DS_ILLEGAL_XDOM_MOVE_OPERATION                                             = 8492
	ERROR_DS_CANT_WITH_ACCT_GROUP_MEMBERSHPS                                         = 8493
	ERROR_DS_NC_MUST_HAVE_NC_PARENT                                                  = 8494
	ERROR_DS_CR_IMPOSSIBLE_TO_VALIDATE                                               = 8495
	ERROR_DS_DST_DOMAIN_NOT_NATIVE                                                   = 8496
	ERROR_DS_MISSING_INFRASTRUCTURE_CONTAINER                                        = 8497
	ERROR_DS_CANT_MOVE_ACCOUNT_GROUP                                                 = 8498
	ERROR_DS_CANT_MOVE_RESOURCE_GROUP                                                = 8499
	ERROR_DS_INVALID_SEARCH_FLAG                                                     = 8500
	ERROR_DS_NO_TREE_DELETE_ABOVE_NC                                                 = 8501
	ERROR_DS_COULDNT_LOCK_TREE_FOR_DELETE                                            = 8502
	ERROR_DS_COULDNT_IDENTIFY_OBJECTS_FOR_TREE_DELETE                                = 8503
	ERROR_DS_SAM_INIT_FAILURE                                                        = 8504
	ERROR_DS_SENSITIVE_GROUP_VIOLATION                                               = 8505
	ERROR_DS_CANT_MOD_PRIMARYGROUPID                                                 = 8506
	ERROR_DS_ILLEGAL_BASE_SCHEMA_MOD                                                 = 8507
	ERROR_DS_NONSAFE_SCHEMA_CHANGE                                                   = 8508
	ERROR_DS_SCHEMA_UPDATE_DISALLOWED                                                = 8509
	ERROR_DS_CANT_CREATE_UNDER_SCHEMA                                                = 8510
	ERROR_DS_INSTALL_NO_SRC_SCH_VERSION                                              = 8511
	ERROR_DS_INSTALL_NO_SCH_VERSION_IN_INIFILE                                       = 8512
	ERROR_DS_INVALID_GROUP_TYPE                                                      = 8513
	ERROR_DS_NO_NEST_GLOBALGROUP_IN_MIXEDDOMAIN                                      = 8514
	ERROR_DS_NO_NEST_LOCALGROUP_IN_MIXEDDOMAIN                                       = 8515
	ERROR_DS_GLOBAL_CANT_HAVE_LOCAL_MEMBER                                           = 8516
	ERROR_DS_GLOBAL_CANT_HAVE_UNIVERSAL_MEMBER                                       = 8517
	ERROR_DS_UNIVERSAL_CANT_HAVE_LOCAL_MEMBER                                        = 8518
	ERROR_DS_GLOBAL_CANT_HAVE_CROSSDOMAIN_MEMBER                                     = 8519
	ERROR_DS_LOCAL_CANT_HAVE_CROSSDOMAIN_LOCAL_MEMBER                                = 8520
	ERROR_DS_HAVE_PRIMARY_MEMBERS                                                    = 8521
	ERROR_DS_STRING_SD_CONVERSION_FAILED                                             = 8522
	ERROR_DS_NAMING_MASTER_GC                                                        = 8523
	ERROR_DS_DNS_LOOKUP_FAILURE                                                      = 8524
	ERROR_DS_COULDNT_UPDATE_SPNS                                                     = 8525
	ERROR_DS_CANT_RETRIEVE_SD                                                        = 8526
	ERROR_DS_KEY_NOT_UNIQUE                                                          = 8527
	ERROR_DS_WRONG_LINKED_ATT_SYNTAX                                                 = 8528
	ERROR_DS_SAM_NEED_BOOTKEY_PASSWORD                                               = 8529
	ERROR_DS_SAM_NEED_BOOTKEY_FLOPPY                                                 = 8530
	ERROR_DS_CANT_START                                                              = 8531
	ERROR_DS_INIT_FAILURE                                                            = 8532
	ERROR_DS_NO_PKT_PRIVACY_ON_CONNECTION                                            = 8533
	ERROR_DS_SOURCE_DOMAIN_IN_FOREST                                                 = 8534
	ERROR_DS_DESTINATION_DOMAIN_NOT_IN_FOREST                                        = 8535
	ERROR_DS_DESTINATION_AUDITING_NOT_ENABLED                                        = 8536
	ERROR_DS_CANT_FIND_DC_FOR_SRC_DOMAIN                                             = 8537
	ERROR_DS_SRC_OBJ_NOT_GROUP_OR_USER                                               = 8538
	ERROR_DS_SRC_SID_EXISTS_IN_FOREST                                                = 8539
	ERROR_DS_SRC_AND_DST_OBJECT_CLASS_MISMATCH                                       = 8540
	ERROR_SAM_INIT_FAILURE                                                           = 8541
	ERROR_DS_DRA_SCHEMA_INFO_SHIP                                                    = 8542
	ERROR_DS_DRA_SCHEMA_CONFLICT                                                     = 8543
	ERROR_DS_DRA_EARLIER_SCHEMA_CONFLICT                                             = 8544
	ERROR_DS_DRA_OBJ_NC_MISMATCH                                                     = 8545
	ERROR_DS_NC_STILL_HAS_DSAS                                                       = 8546
	ERROR_DS_GC_REQUIRED                                                             = 8547
	ERROR_DS_LOCAL_MEMBER_OF_LOCAL_ONLY                                              = 8548
	ERROR_DS_NO_FPO_IN_UNIVERSAL_GROUPS                                              = 8549
	ERROR_DS_CANT_ADD_TO_GC                                                          = 8550
	ERROR_DS_NO_CHECKPOINT_WITH_PDC                                                  = 8551
	ERROR_DS_SOURCE_AUDITING_NOT_ENABLED                                             = 8552
	ERROR_DS_CANT_CREATE_IN_NONDOMAIN_NC                                             = 8553
	ERROR_DS_INVALID_NAME_FOR_SPN                                                    = 8554
	ERROR_DS_FILTER_USES_CONTRUCTED_ATTRS                                            = 8555
	ERROR_DS_UNICODEPWD_NOT_IN_QUOTES                                                = 8556
	ERROR_DS_MACHINE_ACCOUNT_QUOTA_EXCEEDED                                          = 8557
	ERROR_DS_MUST_BE_RUN_ON_DST_DC                                                   = 8558
	ERROR_DS_SRC_DC_MUST_BE_SP4_OR_GREATER                                           = 8559
	ERROR_DS_CANT_TREE_DELETE_CRITICAL_OBJ                                           = 8560
	ERROR_DS_INIT_FAILURE_CONSOLE                                                    = 8561
	ERROR_DS_SAM_INIT_FAILURE_CONSOLE                                                = 8562
	ERROR_DS_FOREST_VERSION_TOO_HIGH                                                 = 8563
	ERROR_DS_DOMAIN_VERSION_TOO_HIGH                                                 = 8564
	ERROR_DS_FOREST_VERSION_TOO_LOW                                                  = 8565
	ERROR_DS_DOMAIN_VERSION_TOO_LOW                                                  = 8566
	ERROR_DS_INCOMPATIBLE_VERSION                                                    = 8567
	ERROR_DS_LOW_DSA_VERSION                                                         = 8568
	ERROR_DS_NO_BEHAVIOR_VERSION_IN_MIXEDDOMAIN                                      = 8569
	ERROR_DS_NOT_SUPPORTED_SORT_ORDER                                                = 8570
	ERROR_DS_NAME_NOT_UNIQUE                                                         = 8571
	ERROR_DS_MACHINE_ACCOUNT_CREATED_PRENT4                                          = 8572
	ERROR_DS_OUT_OF_VERSION_STORE                                                    = 8573
	ERROR_DS_INCOMPATIBLE_CONTROLS_USED                                              = 8574
	ERROR_DS_NO_REF_DOMAIN                                                           = 8575
	ERROR_DS_RESERVED_LINK_ID                                                        = 8576
	ERROR_DS_LINK_ID_NOT_AVAILABLE                                                   = 8577
	ERROR_DS_AG_CANT_HAVE_UNIVERSAL_MEMBER                                           = 8578
	ERROR_DS_MODIFYDN_DISALLOWED_BY_INSTANCE_TYPE                                    = 8579
	ERROR_DS_NO_OBJECT_MOVE_IN_SCHEMA_NC                                             = 8580
	ERROR_DS_MODIFYDN_DISALLOWED_BY_FLAG                                             = 8581
	ERROR_DS_MODIFYDN_WRONG_GRANDPARENT                                              = 8582
	ERROR_DS_NAME_ERROR_TRUST_REFERRAL                                               = 8583
	ERROR_NOT_SUPPORTED_ON_STANDARD_SERVER                                           = 8584
	ERROR_DS_CANT_ACCESS_REMOTE_PART_OF_AD                                           = 8585
	ERROR_DS_CR_IMPOSSIBLE_TO_VALIDATE_V2                                            = 8586
	ERROR_DS_THREAD_LIMIT_EXCEEDED                                                   = 8587
	ERROR_DS_NOT_CLOSEST                                                             = 8588
	ERROR_DS_CANT_DERIVE_SPN_WITHOUT_SERVER_REF                                      = 8589
	ERROR_DS_SINGLE_USER_MODE_FAILED                                                 = 8590
	ERROR_DS_NTDSCRIPT_SYNTAX_ERROR                                                  = 8591
	ERROR_DS_NTDSCRIPT_PROCESS_ERROR                                                 = 8592
	ERROR_DS_DIFFERENT_REPL_EPOCHS                                                   = 8593
	ERROR_DS_DRS_EXTENSIONS_CHANGED                                                  = 8594
	ERROR_DS_REPLICA_SET_CHANGE_NOT_ALLOWED_ON_DISABLED_CR                           = 8595
	ERROR_DS_NO_MSDS_INTID                                                           = 8596
	ERROR_DS_DUP_MSDS_INTID                                                          = 8597
	ERROR_DS_EXISTS_IN_RDNATTID                                                      = 8598
	ERROR_DS_AUTHORIZATION_FAILED                                                    = 8599
	ERROR_DS_INVALID_SCRIPT                                                          = 8600
	ERROR_DS_REMOTE_CROSSREF_OP_FAILED                                               = 8601
	ERROR_DS_CROSS_REF_BUSY                                                          = 8602
	ERROR_DS_CANT_DERIVE_SPN_FOR_DELETED_DOMAIN                                      = 8603
	ERROR_DS_CANT_DEMOTE_WITH_WRITEABLE_NC                                           = 8604
	ERROR_DS_DUPLICATE_ID_FOUND                                                      = 8605
	ERROR_DS_INSUFFICIENT_ATTR_TO_CREATE_OBJECT                                      = 8606
	ERROR_DS_GROUP_CONVERSION_ERROR                                                  = 8607
	ERROR_DS_CANT_MOVE_APP_BASIC_GROUP                                               = 8608
	ERROR_DS_CANT_MOVE_APP_QUERY_GROUP                                               = 8609
	ERROR_DS_ROLE_NOT_VERIFIED                                                       = 8610
	ERROR_DS_WKO_CONTAINER_CANNOT_BE_SPECIAL                                         = 8611
	ERROR_DS_DOMAIN_RENAME_IN_PROGRESS                                               = 8612
	ERROR_DS_EXISTING_AD_CHILD_NC                                                    = 8613
	ERROR_DS_REPL_LIFETIME_EXCEEDED                                                  = 8614
	ERROR_DS_DISALLOWED_IN_SYSTEM_CONTAINER                                          = 8615
	ERROR_DS_LDAP_SEND_QUEUE_FULL                                                    = 8616
	ERROR_DS_DRA_OUT_SCHEDULE_WINDOW                                                 = 8617
	ERROR_DS_POLICY_NOT_KNOWN                                                        = 8618
	ERROR_NO_SITE_SETTINGS_OBJECT                                                    = 8619
	ERROR_NO_SECRETS                                                                 = 8620
	ERROR_NO_WRITABLE_DC_FOUND                                                       = 8621
	ERROR_DS_NO_SERVER_OBJECT                                                        = 8622
	ERROR_DS_NO_NTDSA_OBJECT                                                         = 8623
	ERROR_DS_NON_ASQ_SEARCH                                                          = 8624
	ERROR_DS_AUDIT_FAILURE                                                           = 8625
	ERROR_DS_INVALID_SEARCH_FLAG_SUBTREE                                             = 8626
	ERROR_DS_INVALID_SEARCH_FLAG_TUPLE                                               = 8627
	ERROR_DS_HIERARCHY_TABLE_TOO_DEEP                                                = 8628
	ERROR_DS_DRA_CORRUPT_UTD_VECTOR                                                  = 8629
	ERROR_DS_DRA_SECRETS_DENIED                                                      = 8630
	ERROR_DS_RESERVED_MAPI_ID                                                        = 8631
	ERROR_DS_MAPI_ID_NOT_AVAILABLE                                                   = 8632
	ERROR_DS_DRA_MISSING_KRBTGT_SECRET                                               = 8633
	ERROR_DS_DOMAIN_NAME_EXISTS_IN_FOREST                                            = 8634
	ERROR_DS_FLAT_NAME_EXISTS_IN_FOREST                                              = 8635
	ERROR_INVALID_USER_PRINCIPAL_NAME                                                = 8636
	ERROR_DS_OID_MAPPED_GROUP_CANT_HAVE_MEMBERS                                      = 8637
	ERROR_DS_OID_NOT_FOUND                                                           = 8638
	ERROR_DS_DRA_RECYCLED_TARGET                                                     = 8639
	DNS_ERROR_RESPONSE_CODES_BASE                                                    = 9000
	DNS_ERROR_MASK                                                                   = 0x00002328
	DNS_ERROR_RCODE_FORMAT_ERROR                                                     = 9001
	DNS_ERROR_RCODE_SERVER_FAILURE                                                   = 9002
	DNS_ERROR_RCODE_NAME_ERROR                                                       = 9003
	DNS_ERROR_RCODE_NOT_IMPLEMENTED                                                  = 9004
	DNS_ERROR_RCODE_REFUSED                                                          = 9005
	DNS_ERROR_RCODE_YXDOMAIN                                                         = 9006
	DNS_ERROR_RCODE_YXRRSET                                                          = 9007
	DNS_ERROR_RCODE_NXRRSET                                                          = 9008
	DNS_ERROR_RCODE_NOTAUTH                                                          = 9009
	DNS_ERROR_RCODE_NOTZONE                                                          = 9010
	DNS_ERROR_RCODE_BADSIG                                                           = 9016
	DNS_ERROR_RCODE_BADKEY                                                           = 9017
	DNS_ERROR_RCODE_BADTIME                                                          = 9018
	DNS_ERROR_PACKET_FMT_BASE                                                        = 9500
	DNS_INFO_NO_RECORDS                                                              = 9501
	DNS_ERROR_BAD_PACKET                                                             = 9502
	DNS_ERROR_NO_PACKET                                                              = 9503
	DNS_ERROR_RCODE                                                                  = 9504
	DNS_ERROR_UNSECURE_PACKET                                                        = 9505
	DNS_ERROR_GENERAL_API_BASE                                                       = 9550
	DNS_ERROR_INVALID_TYPE                                                           = 9551
	DNS_ERROR_INVALID_IP_ADDRESS                                                     = 9552
	DNS_ERROR_INVALID_PROPERTY                                                       = 9553
	DNS_ERROR_TRY_AGAIN_LATER                                                        = 9554
	DNS_ERROR_NOT_UNIQUE                                                             = 9555
	DNS_ERROR_NON_RFC_NAME                                                           = 9556
	DNS_STATUS_FQDN                                                                  = 9557
	DNS_STATUS_DOTTED_NAME                                                           = 9558
	DNS_STATUS_SINGLE_PART_NAME                                                      = 9559
	DNS_ERROR_INVALID_NAME_CHAR                                                      = 9560
	DNS_ERROR_NUMERIC_NAME                                                           = 9561
	DNS_ERROR_NOT_ALLOWED_ON_ROOT_SERVER                                             = 9562
	DNS_ERROR_NOT_ALLOWED_UNDER_DELEGATION                                           = 9563
	DNS_ERROR_CANNOT_FIND_ROOT_HINTS                                                 = 9564
	DNS_ERROR_INCONSISTENT_ROOT_HINTS                                                = 9565
	DNS_ERROR_DWORD_VALUE_TOO_SMALL                                                  = 9566
	DNS_ERROR_DWORD_VALUE_TOO_LARGE                                                  = 9567
	DNS_ERROR_BACKGROUND_LOADING                                                     = 9568
	DNS_ERROR_NOT_ALLOWED_ON_RODC                                                    = 9569
	DNS_ERROR_NOT_ALLOWED_UNDER_DNAME                                                = 9570
	DNS_ERROR_DELEGATION_REQUIRED                                                    = 9571
	DNS_ERROR_INVALID_POLICY_TABLE                                                   = 9572
	DNS_ERROR_ZONE_BASE                                                              = 9600
	DNS_ERROR_ZONE_DOES_NOT_EXIST                                                    = 9601
	DNS_ERROR_NO_ZONE_INFO                                                           = 9602
	DNS_ERROR_INVALID_ZONE_OPERATION                                                 = 9603
	DNS_ERROR_ZONE_CONFIGURATION_ERROR                                               = 9604
	DNS_ERROR_ZONE_HAS_NO_SOA_RECORD                                                 = 9605
	DNS_ERROR_ZONE_HAS_NO_NS_RECORDS                                                 = 9606
	DNS_ERROR_ZONE_LOCKED                                                            = 9607
	DNS_ERROR_ZONE_CREATION_FAILED                                                   = 9608
	DNS_ERROR_ZONE_ALREADY_EXISTS                                                    = 9609
	DNS_ERROR_AUTOZONE_ALREADY_EXISTS                                                = 9610
	DNS_ERROR_INVALID_ZONE_TYPE                                                      = 9611
	DNS_ERROR_SECONDARY_REQUIRES_MASTER_IP                                           = 9612
	DNS_ERROR_ZONE_NOT_SECONDARY                                                     = 9613
	DNS_ERROR_NEED_SECONDARY_ADDRESSES                                               = 9614
	DNS_ERROR_WINS_INIT_FAILED                                                       = 9615
	DNS_ERROR_NEED_WINS_SERVERS                                                      = 9616
	DNS_ERROR_NBSTAT_INIT_FAILED                                                     = 9617
	DNS_ERROR_SOA_DELETE_INVALID                                                     = 9618
	DNS_ERROR_FORWARDER_ALREADY_EXISTS                                               = 9619
	DNS_ERROR_ZONE_REQUIRES_MASTER_IP                                                = 9620
	DNS_ERROR_ZONE_IS_SHUTDOWN                                                       = 9621
	DNS_ERROR_DATAFILE_BASE                                                          = 9650
	DNS_ERROR_PRIMARY_REQUIRES_DATAFILE                                              = 9651
	DNS_ERROR_INVALID_DATAFILE_NAME                                                  = 9652
	DNS_ERROR_DATAFILE_OPEN_FAILURE                                                  = 9653
	DNS_ERROR_FILE_WRITEBACK_FAILED                                                  = 9654
	DNS_ERROR_DATAFILE_PARSING                                                       = 9655
	DNS_ERROR_DATABASE_BASE                                                          = 9700
	DNS_ERROR_RECORD_DOES_NOT_EXIST                                                  = 9701
	DNS_ERROR_RECORD_FORMAT                                                          = 9702
	DNS_ERROR_NODE_CREATION_FAILED                                                   = 9703
	DNS_ERROR_UNKNOWN_RECORD_TYPE                                                    = 9704
	DNS_ERROR_RECORD_TIMED_OUT                                                       = 9705
	DNS_ERROR_NAME_NOT_IN_ZONE                                                       = 9706
	DNS_ERROR_CNAME_LOOP                                                             = 9707
	DNS_ERROR_NODE_IS_CNAME                                                          = 9708
	DNS_ERROR_CNAME_COLLISION                                                        = 9709
	DNS_ERROR_RECORD_ONLY_AT_ZONE_ROOT                                               = 9710
	DNS_ERROR_RECORD_ALREADY_EXISTS                                                  = 9711
	DNS_ERROR_SECONDARY_DATA                                                         = 9712
	DNS_ERROR_NO_CREATE_CACHE_DATA                                                   = 9713
	DNS_ERROR_NAME_DOES_NOT_EXIST                                                    = 9714
	DNS_WARNING_PTR_CREATE_FAILED                                                    = 9715
	DNS_WARNING_DOMAIN_UNDELETED                                                     = 9716
	DNS_ERROR_DS_UNAVAILABLE                                                         = 9717
	DNS_ERROR_DS_ZONE_ALREADY_EXISTS                                                 = 9718
	DNS_ERROR_NO_BOOTFILE_IF_DS_ZONE                                                 = 9719
	DNS_ERROR_NODE_IS_DNAME                                                          = 9720
	DNS_ERROR_DNAME_COLLISION                                                        = 9721
	DNS_ERROR_ALIAS_LOOP                                                             = 9722
	DNS_ERROR_OPERATION_BASE                                                         = 9750
	DNS_INFO_AXFR_COMPLETE                                                           = 9751
	DNS_ERROR_AXFR                                                                   = 9752
	DNS_INFO_ADDED_LOCAL_WINS                                                        = 9753
	DNS_ERROR_SECURE_BASE                                                            = 9800
	DNS_STATUS_CONTINUE_NEEDED                                                       = 9801
	DNS_ERROR_SETUP_BASE                                                             = 9850
	DNS_ERROR_NO_TCPIP                                                               = 9851
	DNS_ERROR_NO_DNS_SERVERS                                                         = 9852
	DNS_ERROR_DP_BASE                                                                = 9900
	DNS_ERROR_DP_DOES_NOT_EXIST                                                      = 9901
	DNS_ERROR_DP_ALREADY_EXISTS                                                      = 9902
	DNS_ERROR_DP_NOT_ENLISTED                                                        = 9903
	DNS_ERROR_DP_ALREADY_ENLISTED                                                    = 9904
	DNS_ERROR_DP_NOT_AVAILABLE                                                       = 9905
	DNS_ERROR_DP_FSMO_ERROR                                                          = 9906
	WSABASEERR                                                                       = 10000
	WSAEINTR                                                                         = 10004
	WSAEBADF                                                                         = 10009
	WSAEACCES                                                                        = 10013
	WSAEFAULT                                                                        = 10014
	WSAEINVAL                                                                        = 10022
	WSAEMFILE                                                                        = 10024
	WSAEWOULDBLOCK                                                                   = 10035
	WSAEINPROGRESS                                                                   = 10036
	WSAEALREADY                                                                      = 10037
	WSAENOTSOCK                                                                      = 10038
	WSAEDESTADDRREQ                                                                  = 10039
	WSAEMSGSIZE                                                                      = 10040
	WSAEPROTOTYPE                                                                    = 10041
	WSAENOPROTOOPT                                                                   = 10042
	WSAEPROTONOSUPPORT                                                               = 10043
	WSAESOCKTNOSUPPORT                                                               = 10044
	WSAEOPNOTSUPP                                                                    = 10045
	WSAEPFNOSUPPORT                                                                  = 10046
	WSAEAFNOSUPPORT                                                                  = 10047
	WSAEADDRINUSE                                                                    = 10048
	WSAEADDRNOTAVAIL                                                                 = 10049
	WSAENETDOWN                                                                      = 10050
	WSAENETUNREACH                                                                   = 10051
	WSAENETRESET                                                                     = 10052
	WSAECONNABORTED                                                                  = 10053
	WSAECONNRESET                                                                    = 10054
	WSAENOBUFS                                                                       = 10055
	WSAEISCONN                                                                       = 10056
	WSAENOTCONN                                                                      = 10057
	WSAESHUTDOWN                                                                     = 10058
	WSAETOOMANYREFS                                                                  = 10059
	WSAETIMEDOUT                                                                     = 10060
	WSAECONNREFUSED                                                                  = 10061
	WSAELOOP                                                                         = 10062
	WSAENAMETOOLONG                                                                  = 10063
	WSAEHOSTDOWN                                                                     = 10064
	WSAEHOSTUNREACH                                                                  = 10065
	WSAENOTEMPTY                                                                     = 10066
	WSAEPROCLIM                                                                      = 10067
	WSAEUSERS                                                                        = 10068
	WSAEDQUOT                                                                        = 10069
	WSAESTALE                                                                        = 10070
	WSAEREMOTE                                                                       = 10071
	WSASYSNOTREADY                                                                   = 10091
	WSAVERNOTSUPPORTED                                                               = 10092
	WSANOTINITIALISED                                                                = 10093
	WSAEDISCON                                                                       = 10101
	WSAENOMORE                                                                       = 10102
	WSAECANCELLED                                                                    = 10103
	WSAEINVALIDPROCTABLE                                                             = 10104
	WSAEINVALIDPROVIDER                                                              = 10105
	WSAEPROVIDERFAILEDINIT                                                           = 10106
	WSASYSCALLFAILURE                                                                = 10107
	WSASERVICE_NOT_FOUND                                                             = 10108
	WSATYPE_NOT_FOUND                                                                = 10109
	WSA_E_NO_MORE                                                                    = 10110
	WSA_E_CANCELLED                                                                  = 10111
	WSAEREFUSED                                                                      = 10112
	WSAHOST_NOT_FOUND                                                                = 11001
	WSATRY_AGAIN                                                                     = 11002
	WSANO_RECOVERY                                                                   = 11003
	WSANO_DATA                                                                       = 11004
	WSA_QOS_RECEIVERS                                                                = 11005
	WSA_QOS_SENDERS                                                                  = 11006
	WSA_QOS_NO_SENDERS                                                               = 11007
	WSA_QOS_NO_RECEIVERS                                                             = 11008
	WSA_QOS_REQUEST_CONFIRMED                                                        = 11009
	WSA_QOS_ADMISSION_FAILURE                                                        = 11010
	WSA_QOS_POLICY_FAILURE                                                           = 11011
	WSA_QOS_BAD_STYLE                                                                = 11012
	WSA_QOS_BAD_OBJECT                                                               = 11013
	WSA_QOS_TRAFFIC_CTRL_ERROR                                                       = 11014
	WSA_QOS_GENERIC_ERROR                                                            = 11015
	WSA_QOS_ESERVICETYPE                                                             = 11016
	WSA_QOS_EFLOWSPEC                                                                = 11017
	WSA_QOS_EPROVSPECBUF                                                             = 11018
	WSA_QOS_EFILTERSTYLE                                                             = 11019
	WSA_QOS_EFILTERTYPE                                                              = 11020
	WSA_QOS_EFILTERCOUNT                                                             = 11021
	WSA_QOS_EOBJLENGTH                                                               = 11022
	WSA_QOS_EFLOWCOUNT                                                               = 11023
	WSA_QOS_EUNKOWNPSOBJ                                                             = 11024
	WSA_QOS_EPOLICYOBJ                                                               = 11025
	WSA_QOS_EFLOWDESC                                                                = 11026
	WSA_QOS_EPSFLOWSPEC                                                              = 11027
	WSA_QOS_EPSFILTERSPEC                                                            = 11028
	WSA_QOS_ESDMODEOBJ                                                               = 11029
	WSA_QOS_ESHAPERATEOBJ                                                            = 11030
	WSA_QOS_RESERVED_PETYPE                                                          = 11031
	WSA_SECURE_HOST_NOT_FOUND                                                        = 11032
	WSA_IPSEC_NAME_POLICY_ERROR                                                      = 11033
	ERROR_IPSEC_QM_POLICY_EXISTS                                                     = 13000
	ERROR_IPSEC_QM_POLICY_NOT_FOUND                                                  = 13001
	ERROR_IPSEC_QM_POLICY_IN_USE                                                     = 13002
	ERROR_IPSEC_MM_POLICY_EXISTS                                                     = 13003
	ERROR_IPSEC_MM_POLICY_NOT_FOUND                                                  = 13004
	ERROR_IPSEC_MM_POLICY_IN_USE                                                     = 13005
	ERROR_IPSEC_MM_FILTER_EXISTS                                                     = 13006
	ERROR_IPSEC_MM_FILTER_NOT_FOUND                                                  = 13007
	ERROR_IPSEC_TRANSPORT_FILTER_EXISTS                                              = 13008
	ERROR_IPSEC_TRANSPORT_FILTER_NOT_FOUND                                           = 13009
	ERROR_IPSEC_MM_AUTH_EXISTS                                                       = 13010
	ERROR_IPSEC_MM_AUTH_NOT_FOUND                                                    = 13011
	ERROR_IPSEC_MM_AUTH_IN_USE                                                       = 13012
	ERROR_IPSEC_DEFAULT_MM_POLICY_NOT_FOUND                                          = 13013
	ERROR_IPSEC_DEFAULT_MM_AUTH_NOT_FOUND                                            = 13014
	ERROR_IPSEC_DEFAULT_QM_POLICY_NOT_FOUND                                          = 13015
	ERROR_IPSEC_TUNNEL_FILTER_EXISTS                                                 = 13016
	ERROR_IPSEC_TUNNEL_FILTER_NOT_FOUND                                              = 13017
	ERROR_IPSEC_MM_FILTER_PENDING_DELETION                                           = 13018
	ERROR_IPSEC_TRANSPORT_FILTER_PENDING_DELETION                                    = 13019
	ERROR_IPSEC_TUNNEL_FILTER_PENDING_DELETION                                       = 13020
	ERROR_IPSEC_MM_POLICY_PENDING_DELETION                                           = 13021
	ERROR_IPSEC_MM_AUTH_PENDING_DELETION                                             = 13022
	ERROR_IPSEC_QM_POLICY_PENDING_DELETION                                           = 13023
	WARNING_IPSEC_MM_POLICY_PRUNED                                                   = 13024
	WARNING_IPSEC_QM_POLICY_PRUNED                                                   = 13025
	ERROR_IPSEC_IKE_NEG_STATUS_BEGIN                                                 = 13800
	ERROR_IPSEC_IKE_AUTH_FAIL                                                        = 13801
	ERROR_IPSEC_IKE_ATTRIB_FAIL                                                      = 13802
	ERROR_IPSEC_IKE_NEGOTIATION_PENDING                                              = 13803
	ERROR_IPSEC_IKE_GENERAL_PROCESSING_ERROR                                         = 13804
	ERROR_IPSEC_IKE_TIMED_OUT                                                        = 13805
	ERROR_IPSEC_IKE_NO_CERT                                                          = 13806
	ERROR_IPSEC_IKE_SA_DELETED                                                       = 13807
	ERROR_IPSEC_IKE_SA_REAPED                                                        = 13808
	ERROR_IPSEC_IKE_MM_ACQUIRE_DROP                                                  = 13809
	ERROR_IPSEC_IKE_QM_ACQUIRE_DROP                                                  = 13810
	ERROR_IPSEC_IKE_QUEUE_DROP_MM                                                    = 13811
	ERROR_IPSEC_IKE_QUEUE_DROP_NO_MM                                                 = 13812
	ERROR_IPSEC_IKE_DROP_NO_RESPONSE                                                 = 13813
	ERROR_IPSEC_IKE_MM_DELAY_DROP                                                    = 13814
	ERROR_IPSEC_IKE_QM_DELAY_DROP                                                    = 13815
	ERROR_IPSEC_IKE_ERROR                                                            = 13816
	ERROR_IPSEC_IKE_CRL_FAILED                                                       = 13817
	ERROR_IPSEC_IKE_INVALID_KEY_USAGE                                                = 13818
	ERROR_IPSEC_IKE_INVALID_CERT_TYPE                                                = 13819
	ERROR_IPSEC_IKE_NO_PRIVATE_KEY                                                   = 13820
	ERROR_IPSEC_IKE_SIMULTANEOUS_REKEY                                               = 13821
	ERROR_IPSEC_IKE_DH_FAIL                                                          = 13822
	ERROR_IPSEC_IKE_CRITICAL_PAYLOAD_NOT_RECOGNIZED                                  = 13823
	ERROR_IPSEC_IKE_INVALID_HEADER                                                   = 13824
	ERROR_IPSEC_IKE_NO_POLICY                                                        = 13825
	ERROR_IPSEC_IKE_INVALID_SIGNATURE                                                = 13826
	ERROR_IPSEC_IKE_KERBEROS_ERROR                                                   = 13827
	ERROR_IPSEC_IKE_NO_PUBLIC_KEY                                                    = 13828
	ERROR_IPSEC_IKE_PROCESS_ERR                                                      = 13829
	ERROR_IPSEC_IKE_PROCESS_ERR_SA                                                   = 13830
	ERROR_IPSEC_IKE_PROCESS_ERR_PROP                                                 = 13831
	ERROR_IPSEC_IKE_PROCESS_ERR_TRANS                                                = 13832
	ERROR_IPSEC_IKE_PROCESS_ERR_KE                                                   = 13833
	ERROR_IPSEC_IKE_PROCESS_ERR_ID                                                   = 13834
	ERROR_IPSEC_IKE_PROCESS_ERR_CERT                                                 = 13835
	ERROR_IPSEC_IKE_PROCESS_ERR_CERT_REQ                                             = 13836
	ERROR_IPSEC_IKE_PROCESS_ERR_HASH                                                 = 13837
	ERROR_IPSEC_IKE_PROCESS_ERR_SIG                                                  = 13838
	ERROR_IPSEC_IKE_PROCESS_ERR_NONCE                                                = 13839
	ERROR_IPSEC_IKE_PROCESS_ERR_NOTIFY                                               = 13840
	ERROR_IPSEC_IKE_PROCESS_ERR_DELETE                                               = 13841
	ERROR_IPSEC_IKE_PROCESS_ERR_VENDOR                                               = 13842
	ERROR_IPSEC_IKE_INVALID_PAYLOAD                                                  = 13843
	ERROR_IPSEC_IKE_LOAD_SOFT_SA                                                     = 13844
	ERROR_IPSEC_IKE_SOFT_SA_TORN_DOWN                                                = 13845
	ERROR_IPSEC_IKE_INVALID_COOKIE                                                   = 13846
	ERROR_IPSEC_IKE_NO_PEER_CERT                                                     = 13847
	ERROR_IPSEC_IKE_PEER_CRL_FAILED                                                  = 13848
	ERROR_IPSEC_IKE_POLICY_CHANGE                                                    = 13849
	ERROR_IPSEC_IKE_NO_MM_POLICY                                                     = 13850
	ERROR_IPSEC_IKE_NOTCBPRIV                                                        = 13851
	ERROR_IPSEC_IKE_SECLOADFAIL                                                      = 13852
	ERROR_IPSEC_IKE_FAILSSPINIT                                                      = 13853
	ERROR_IPSEC_IKE_FAILQUERYSSP                                                     = 13854
	ERROR_IPSEC_IKE_SRVACQFAIL                                                       = 13855
	ERROR_IPSEC_IKE_SRVQUERYCRED                                                     = 13856
	ERROR_IPSEC_IKE_GETSPIFAIL                                                       = 13857
	ERROR_IPSEC_IKE_INVALID_FILTER                                                   = 13858
	ERROR_IPSEC_IKE_OUT_OF_MEMORY                                                    = 13859
	ERROR_IPSEC_IKE_ADD_UPDATE_KEY_FAILED                                            = 13860
	ERROR_IPSEC_IKE_INVALID_POLICY                                                   = 13861
	ERROR_IPSEC_IKE_UNKNOWN_DOI                                                      = 13862
	ERROR_IPSEC_IKE_INVALID_SITUATION                                                = 13863
	ERROR_IPSEC_IKE_DH_FAILURE                                                       = 13864
	ERROR_IPSEC_IKE_INVALID_GROUP                                                    = 13865
	ERROR_IPSEC_IKE_ENCRYPT                                                          = 13866
	ERROR_IPSEC_IKE_DECRYPT                                                          = 13867
	ERROR_IPSEC_IKE_POLICY_MATCH                                                     = 13868
	ERROR_IPSEC_IKE_UNSUPPORTED_ID                                                   = 13869
	ERROR_IPSEC_IKE_INVALID_HASH                                                     = 13870
	ERROR_IPSEC_IKE_INVALID_HASH_ALG                                                 = 13871
	ERROR_IPSEC_IKE_INVALID_HASH_SIZE                                                = 13872
	ERROR_IPSEC_IKE_INVALID_ENCRYPT_ALG                                              = 13873
	ERROR_IPSEC_IKE_INVALID_AUTH_ALG                                                 = 13874
	ERROR_IPSEC_IKE_INVALID_SIG                                                      = 13875
	ERROR_IPSEC_IKE_LOAD_FAILED                                                      = 13876
	ERROR_IPSEC_IKE_RPC_DELETE                                                       = 13877
	ERROR_IPSEC_IKE_BENIGN_REINIT                                                    = 13878
	ERROR_IPSEC_IKE_INVALID_RESPONDER_LIFETIME_NOTIFY                                = 13879
	ERROR_IPSEC_IKE_INVALID_MAJOR_VERSION                                            = 13880
	ERROR_IPSEC_IKE_INVALID_CERT_KEYLEN                                              = 13881
	ERROR_IPSEC_IKE_MM_LIMIT                                                         = 13882
	ERROR_IPSEC_IKE_NEGOTIATION_DISABLED                                             = 13883
	ERROR_IPSEC_IKE_QM_LIMIT                                                         = 13884
	ERROR_IPSEC_IKE_MM_EXPIRED                                                       = 13885
	ERROR_IPSEC_IKE_PEER_MM_ASSUMED_INVALID                                          = 13886
	ERROR_IPSEC_IKE_CERT_CHAIN_POLICY_MISMATCH                                       = 13887
	ERROR_IPSEC_IKE_UNEXPECTED_MESSAGE_ID                                            = 13888
	ERROR_IPSEC_IKE_INVALID_AUTH_PAYLOAD                                             = 13889
	ERROR_IPSEC_IKE_DOS_COOKIE_SENT                                                  = 13890
	ERROR_IPSEC_IKE_SHUTTING_DOWN                                                    = 13891
	ERROR_IPSEC_IKE_CGA_AUTH_FAILED                                                  = 13892
	ERROR_IPSEC_IKE_PROCESS_ERR_NATOA                                                = 13893
	ERROR_IPSEC_IKE_INVALID_MM_FOR_QM                                                = 13894
	ERROR_IPSEC_IKE_QM_EXPIRED                                                       = 13895
	ERROR_IPSEC_IKE_TOO_MANY_FILTERS                                                 = 13896
	ERROR_IPSEC_IKE_NEG_STATUS_END                                                   = 13897
	ERROR_IPSEC_IKE_KILL_DUMMY_NAP_TUNNEL                                            = 13898
	ERROR_IPSEC_IKE_INNER_IP_ASSIGNMENT_FAILURE                                      = 13899
	ERROR_IPSEC_IKE_REQUIRE_CP_PAYLOAD_MISSING                                       = 13900
	ERROR_IPSEC_KEY_MODULE_IMPERSONATION_NEGOTIATION_PENDING                         = 13901
	ERROR_IPSEC_IKE_COEXISTENCE_SUPPRESS                                             = 13902
	ERROR_IPSEC_IKE_RATELIMIT_DROP                                                   = 13903
	ERROR_IPSEC_IKE_PEER_DOESNT_SUPPORT_MOBIKE                                       = 13904
	ERROR_IPSEC_IKE_AUTHORIZATION_FAILURE                                            = 13905
	ERROR_IPSEC_IKE_STRONG_CRED_AUTHORIZATION_FAILURE                                = 13906
	ERROR_IPSEC_IKE_AUTHORIZATION_FAILURE_WITH_OPTIONAL_RETRY                        = 13907
	ERROR_IPSEC_IKE_STRONG_CRED_AUTHORIZATION_AND_CERTMAP_FAILURE                    = 13908
	ERROR_IPSEC_IKE_NEG_STATUS_EXTENDED_END                                          = 13909
	ERROR_IPSEC_BAD_SPI                                                              = 13910
	ERROR_IPSEC_SA_LIFETIME_EXPIRED                                                  = 13911
	ERROR_IPSEC_WRONG_SA                                                             = 13912
	ERROR_IPSEC_REPLAY_CHECK_FAILED                                                  = 13913
	ERROR_IPSEC_INVALID_PACKET                                                       = 13914
	ERROR_IPSEC_INTEGRITY_CHECK_FAILED                                               = 13915
	ERROR_IPSEC_CLEAR_TEXT_DROP                                                      = 13916
	ERROR_IPSEC_AUTH_FIREWALL_DROP                                                   = 13917
	ERROR_IPSEC_THROTTLE_DROP                                                        = 13918
	ERROR_IPSEC_DOSP_BLOCK                                                           = 13925
	ERROR_IPSEC_DOSP_RECEIVED_MULTICAST                                              = 13926
	ERROR_IPSEC_DOSP_INVALID_PACKET                                                  = 13927
	ERROR_IPSEC_DOSP_STATE_LOOKUP_FAILED                                             = 13928
	ERROR_IPSEC_DOSP_MAX_ENTRIES                                                     = 13929
	ERROR_IPSEC_DOSP_KEYMOD_NOT_ALLOWED                                              = 13930
	ERROR_IPSEC_DOSP_NOT_INSTALLED                                                   = 13931
	ERROR_IPSEC_DOSP_MAX_PER_IP_RATELIMIT_QUEUES                                     = 13932
	ERROR_SXS_SECTION_NOT_FOUND                                                      = 14000
	ERROR_SXS_CANT_GEN_ACTCTX                                                        = 14001
	ERROR_SXS_INVALID_ACTCTXDATA_FORMAT                                              = 14002
	ERROR_SXS_ASSEMBLY_NOT_FOUND                                                     = 14003
	ERROR_SXS_MANIFEST_FORMAT_ERROR                                                  = 14004
	ERROR_SXS_MANIFEST_PARSE_ERROR                                                   = 14005
	ERROR_SXS_ACTIVATION_CONTEXT_DISABLED                                            = 14006
	ERROR_SXS_KEY_NOT_FOUND                                                          = 14007
	ERROR_SXS_VERSION_CONFLICT                                                       = 14008
	ERROR_SXS_WRONG_SECTION_TYPE                                                     = 14009
	ERROR_SXS_THREAD_QUERIES_DISABLED                                                = 14010
	ERROR_SXS_PROCESS_DEFAULT_ALREADY_SET                                            = 14011
	ERROR_SXS_UNKNOWN_ENCODING_GROUP                                                 = 14012
	ERROR_SXS_UNKNOWN_ENCODING                                                       = 14013
	ERROR_SXS_INVALID_XML_NAMESPACE_URI                                              = 14014
	ERROR_SXS_ROOT_MANIFEST_DEPENDENCY_NOT_INSTALLED                                 = 14015
	ERROR_SXS_LEAF_MANIFEST_DEPENDENCY_NOT_INSTALLED                                 = 14016
	ERROR_SXS_INVALID_ASSEMBLY_IDENTITY_ATTRIBUTE                                    = 14017
	ERROR_SXS_MANIFEST_MISSING_REQUIRED_DEFAULT_NAMESPACE                            = 14018
	ERROR_SXS_MANIFEST_INVALID_REQUIRED_DEFAULT_NAMESPACE                            = 14019
	ERROR_SXS_PRIVATE_MANIFEST_CROSS_PATH_WITH_REPARSE_POINT                         = 14020
	ERROR_SXS_DUPLICATE_DLL_NAME                                                     = 14021
	ERROR_SXS_DUPLICATE_WINDOWCLASS_NAME                                             = 14022
	ERROR_SXS_DUPLICATE_CLSID                                                        = 14023
	ERROR_SXS_DUPLICATE_IID                                                          = 14024
	ERROR_SXS_DUPLICATE_TLBID                                                        = 14025
	ERROR_SXS_DUPLICATE_PROGID                                                       = 14026
	ERROR_SXS_DUPLICATE_ASSEMBLY_NAME                                                = 14027
	ERROR_SXS_FILE_HASH_MISMATCH                                                     = 14028
	ERROR_SXS_POLICY_PARSE_ERROR                                                     = 14029
	ERROR_SXS_XML_E_MISSINGQUOTE                                                     = 14030
	ERROR_SXS_XML_E_COMMENTSYNTAX                                                    = 14031
	ERROR_SXS_XML_E_BADSTARTNAMECHAR                                                 = 14032
	ERROR_SXS_XML_E_BADNAMECHAR                                                      = 14033
	ERROR_SXS_XML_E_BADCHARINSTRING                                                  = 14034
	ERROR_SXS_XML_E_XMLDECLSYNTAX                                                    = 14035
	ERROR_SXS_XML_E_BADCHARDATA                                                      = 14036
	ERROR_SXS_XML_E_MISSINGWHITESPACE                                                = 14037
	ERROR_SXS_XML_E_EXPECTINGTAGEND                                                  = 14038
	ERROR_SXS_XML_E_MISSINGSEMICOLON                                                 = 14039
	ERROR_SXS_XML_E_UNBALANCEDPAREN                                                  = 14040
	ERROR_SXS_XML_E_INTERNALERROR                                                    = 14041
	ERROR_SXS_XML_E_UNEXPECTED_WHITESPACE                                            = 14042
	ERROR_SXS_XML_E_INCOMPLETE_ENCODING                                              = 14043
	ERROR_SXS_XML_E_MISSING_PAREN                                                    = 14044
	ERROR_SXS_XML_E_EXPECTINGCLOSEQUOTE                                              = 14045
	ERROR_SXS_XML_E_MULTIPLE_COLONS                                                  = 14046
	ERROR_SXS_XML_E_INVALID_DECIMAL                                                  = 14047
	ERROR_SXS_XML_E_INVALID_HEXIDECIMAL                                              = 14048
	ERROR_SXS_XML_E_INVALID_UNICODE                                                  = 14049
	ERROR_SXS_XML_E_WHITESPACEORQUESTIONMARK                                         = 14050
	ERROR_SXS_XML_E_UNEXPECTEDENDTAG                                                 = 14051
	ERROR_SXS_XML_E_UNCLOSEDTAG                                                      = 14052
	ERROR_SXS_XML_E_DUPLICATEATTRIBUTE                                               = 14053
	ERROR_SXS_XML_E_MULTIPLEROOTS                                                    = 14054
	ERROR_SXS_XML_E_INVALIDATROOTLEVEL                                               = 14055
	ERROR_SXS_XML_E_BADXMLDECL                                                       = 14056
	ERROR_SXS_XML_E_MISSINGROOT                                                      = 14057
	ERROR_SXS_XML_E_UNEXPECTEDEOF                                                    = 14058
	ERROR_SXS_XML_E_BADPEREFINSUBSET                                                 = 14059
	ERROR_SXS_XML_E_UNCLOSEDSTARTTAG                                                 = 14060
	ERROR_SXS_XML_E_UNCLOSEDENDTAG                                                   = 14061
	ERROR_SXS_XML_E_UNCLOSEDSTRING                                                   = 14062
	ERROR_SXS_XML_E_UNCLOSEDCOMMENT                                                  = 14063
	ERROR_SXS_XML_E_UNCLOSEDDECL                                                     = 14064
	ERROR_SXS_XML_E_UNCLOSEDCDATA                                                    = 14065
	ERROR_SXS_XML_E_RESERVEDNAMESPACE                                                = 14066
	ERROR_SXS_XML_E_INVALIDENCODING                                                  = 14067
	ERROR_SXS_XML_E_INVALIDSWITCH                                                    = 14068
	ERROR_SXS_XML_E_BADXMLCASE                                                       = 14069
	ERROR_SXS_XML_E_INVALID_STANDALONE                                               = 14070
	ERROR_SXS_XML_E_UNEXPECTED_STANDALONE                                            = 14071
	ERROR_SXS_XML_E_INVALID_VERSION                                                  = 14072
	ERROR_SXS_XML_E_MISSINGEQUALS                                                    = 14073
	ERROR_SXS_PROTECTION_RECOVERY_FAILED                                             = 14074
	ERROR_SXS_PROTECTION_PUBLIC_KEY_TOO_SHORT                                        = 14075
	ERROR_SXS_PROTECTION_CATALOG_NOT_VALID                                           = 14076
	ERROR_SXS_UNTRANSLATABLE_HRESULT                                                 = 14077
	ERROR_SXS_PROTECTION_CATALOG_FILE_MISSING                                        = 14078
	ERROR_SXS_MISSING_ASSEMBLY_IDENTITY_ATTRIBUTE                                    = 14079
	ERROR_SXS_INVALID_ASSEMBLY_IDENTITY_ATTRIBUTE_NAME                               = 14080
	ERROR_SXS_ASSEMBLY_MISSING                                                       = 14081
	ERROR_SXS_CORRUPT_ACTIVATION_STACK                                               = 14082
	ERROR_SXS_CORRUPTION                                                             = 14083
	ERROR_SXS_EARLY_DEACTIVATION                                                     = 14084
	ERROR_SXS_INVALID_DEACTIVATION                                                   = 14085
	ERROR_SXS_MULTIPLE_DEACTIVATION                                                  = 14086
	ERROR_SXS_PROCESS_TERMINATION_REQUESTED                                          = 14087
	ERROR_SXS_RELEASE_ACTIVATION_CONTEXT                                             = 14088
	ERROR_SXS_SYSTEM_DEFAULT_ACTIVATION_CONTEXT_EMPTY                                = 14089
	ERROR_SXS_INVALID_IDENTITY_ATTRIBUTE_VALUE                                       = 14090
	ERROR_SXS_INVALID_IDENTITY_ATTRIBUTE_NAME                                        = 14091
	ERROR_SXS_IDENTITY_DUPLICATE_ATTRIBUTE                                           = 14092
	ERROR_SXS_IDENTITY_PARSE_ERROR                                                   = 14093
	ERROR_MALFORMED_SUBSTITUTION_STRING                                              = 14094
	ERROR_SXS_INCORRECT_PUBLIC_KEY_TOKEN                                             = 14095
	ERROR_UNMAPPED_SUBSTITUTION_STRING                                               = 14096
	ERROR_SXS_ASSEMBLY_NOT_LOCKED                                                    = 14097
	ERROR_SXS_COMPONENT_STORE_CORRUPT                                                = 14098
	ERROR_ADVANCED_INSTALLER_FAILED                                                  = 14099
	ERROR_XML_ENCODING_MISMATCH                                                      = 14100
	ERROR_SXS_MANIFEST_IDENTITY_SAME_BUT_CONTENTS_DIFFERENT                          = 14101
	ERROR_SXS_IDENTITIES_DIFFERENT                                                   = 14102
	ERROR_SXS_ASSEMBLY_IS_NOT_A_DEPLOYMENT                                           = 14103
	ERROR_SXS_FILE_NOT_PART_OF_ASSEMBLY                                              = 14104
	ERROR_SXS_MANIFEST_TOO_BIG                                                       = 14105
	ERROR_SXS_SETTING_NOT_REGISTERED                                                 = 14106
	ERROR_SXS_TRANSACTION_CLOSURE_INCOMPLETE                                         = 14107
	ERROR_SMI_PRIMITIVE_INSTALLER_FAILED                                             = 14108
	ERROR_GENERIC_COMMAND_FAILED                                                     = 14109
	ERROR_SXS_FILE_HASH_MISSING                                                      = 14110
	ERROR_EVT_INVALID_CHANNEL_PATH                                                   = 15000
	ERROR_EVT_INVALID_QUERY                                                          = 15001
	ERROR_EVT_PUBLISHER_METADATA_NOT_FOUND                                           = 15002
	ERROR_EVT_EVENT_TEMPLATE_NOT_FOUND                                               = 15003
	ERROR_EVT_INVALID_PUBLISHER_NAME                                                 = 15004
	ERROR_EVT_INVALID_EVENT_DATA                                                     = 15005
	ERROR_EVT_CHANNEL_NOT_FOUND                                                      = 15007
	ERROR_EVT_MALFORMED_XML_TEXT                                                     = 15008
	ERROR_EVT_SUBSCRIPTION_TO_DIRECT_CHANNEL                                         = 15009
	ERROR_EVT_CONFIGURATION_ERROR                                                    = 15010
	ERROR_EVT_QUERY_RESULT_STALE                                                     = 15011
	ERROR_EVT_QUERY_RESULT_INVALID_POSITION                                          = 15012
	ERROR_EVT_NON_VALIDATING_MSXML                                                   = 15013
	ERROR_EVT_FILTER_ALREADYSCOPED                                                   = 15014
	ERROR_EVT_FILTER_NOTELTSET                                                       = 15015
	ERROR_EVT_FILTER_INVARG                                                          = 15016
	ERROR_EVT_FILTER_INVTEST                                                         = 15017
	ERROR_EVT_FILTER_INVTYPE                                                         = 15018
	ERROR_EVT_FILTER_PARSEERR                                                        = 15019
	ERROR_EVT_FILTER_UNSUPPORTEDOP                                                   = 15020
	ERROR_EVT_FILTER_UNEXPECTEDTOKEN                                                 = 15021
	ERROR_EVT_INVALID_OPERATION_OVER_ENABLED_DIRECT_CHANNEL                          = 15022
	ERROR_EVT_INVALID_CHANNEL_PROPERTY_VALUE                                         = 15023
	ERROR_EVT_INVALID_PUBLISHER_PROPERTY_VALUE                                       = 15024
	ERROR_EVT_CHANNEL_CANNOT_ACTIVATE                                                = 15025
	ERROR_EVT_FILTER_TOO_COMPLEX                                                     = 15026
	ERROR_EVT_MESSAGE_NOT_FOUND                                                      = 15027
	ERROR_EVT_MESSAGE_ID_NOT_FOUND                                                   = 15028
	ERROR_EVT_UNRESOLVED_VALUE_INSERT                                                = 15029
	ERROR_EVT_UNRESOLVED_PARAMETER_INSERT                                            = 15030
	ERROR_EVT_MAX_INSERTS_REACHED                                                    = 15031
	ERROR_EVT_EVENT_DEFINITION_NOT_FOUND                                             = 15032
	ERROR_EVT_MESSAGE_LOCALE_NOT_FOUND                                               = 15033
	ERROR_EVT_VERSION_TOO_OLD                                                        = 15034
	ERROR_EVT_VERSION_TOO_NEW                                                        = 15035
	ERROR_EVT_CANNOT_OPEN_CHANNEL_OF_QUERY                                           = 15036
	ERROR_EVT_PUBLISHER_DISABLED                                                     = 15037
	ERROR_EVT_FILTER_OUT_OF_RANGE                                                    = 15038
	ERROR_EC_SUBSCRIPTION_CANNOT_ACTIVATE                                            = 15080
	ERROR_EC_LOG_DISABLED                                                            = 15081
	ERROR_EC_CIRCULAR_FORWARDING                                                     = 15082
	ERROR_EC_CREDSTORE_FULL                                                          = 15083
	ERROR_EC_CRED_NOT_FOUND                                                          = 15084
	ERROR_EC_NO_ACTIVE_CHANNEL                                                       = 15085
	ERROR_MUI_FILE_NOT_FOUND                                                         = 15100
	ERROR_MUI_INVALID_FILE                                                           = 15101
	ERROR_MUI_INVALID_RC_CONFIG                                                      = 15102
	ERROR_MUI_INVALID_LOCALE_NAME                                                    = 15103
	ERROR_MUI_INVALID_ULTIMATEFALLBACK_NAME                                          = 15104
	ERROR_MUI_FILE_NOT_LOADED                                                        = 15105
	ERROR_RESOURCE_ENUM_USER_STOP                                                    = 15106
	ERROR_MUI_INTLSETTINGS_UILANG_NOT_INSTALLED                                      = 15107
	ERROR_MUI_INTLSETTINGS_INVALID_LOCALE_NAME                                       = 15108
	ERROR_MCA_INVALID_CAPABILITIES_STRING                                            = 15200
	ERROR_MCA_INVALID_VCP_VERSION                                                    = 15201
	ERROR_MCA_MONITOR_VIOLATES_MCCS_SPECIFICATION                                    = 15202
	ERROR_MCA_MCCS_VERSION_MISMATCH                                                  = 15203
	ERROR_MCA_UNSUPPORTED_MCCS_VERSION                                               = 15204
	ERROR_MCA_INTERNAL_ERROR                                                         = 15205
	ERROR_MCA_INVALID_TECHNOLOGY_TYPE_RETURNED                                       = 15206
	ERROR_MCA_UNSUPPORTED_COLOR_TEMPERATURE                                          = 15207
	ERROR_AMBIGUOUS_SYSTEM_DEVICE                                                    = 15250
	ERROR_SYSTEM_DEVICE_NOT_FOUND                                                    = 15299
	ERROR_HASH_NOT_SUPPORTED                                                         = 15300
	ERROR_HASH_NOT_PRESENT                                                           = 15301
	SEVERITY_SUCCESS                                                                 = 0
	SEVERITY_ERROR                                                                   = 1
	FACILITY_NT_BIT                                                                  = 0x10000000
	E_UNEXPECTED                                                                     = 0x8000FFFF
	E_NOTIMPL                                                                        = 0x80004001
	E_OUTOFMEMORY                                                                    = 0x8007000E
	E_INVALIDARG                                                                     = 0x80070057
	E_NOINTERFACE                                                                    = 0x80004002
	E_POINTER                                                                        = 0x80004003
	E_HANDLE                                                                         = 0x80070006
	E_ABORT                                                                          = 0x80004004
	E_FAIL                                                                           = 0x80004005
	E_ACCESSDENIED                                                                   = 0x80070005
	E_PENDING                                                                        = 0x8000000A
	CO_E_INIT_TLS                                                                    = 0x80004006
	CO_E_INIT_SHARED_ALLOCATOR                                                       = 0x80004007
	CO_E_INIT_MEMORY_ALLOCATOR                                                       = 0x80004008
	CO_E_INIT_CLASS_CACHE                                                            = 0x80004009
	CO_E_INIT_RPC_CHANNEL                                                            = 0x8000400A
	CO_E_INIT_TLS_SET_CHANNEL_CONTROL                                                = 0x8000400B
	CO_E_INIT_TLS_CHANNEL_CONTROL                                                    = 0x8000400C
	CO_E_INIT_UNACCEPTED_USER_ALLOCATOR                                              = 0x8000400D
	CO_E_INIT_SCM_MUTEX_EXISTS                                                       = 0x8000400E
	CO_E_INIT_SCM_FILE_MAPPING_EXISTS                                                = 0x8000400F
	CO_E_INIT_SCM_MAP_VIEW_OF_FILE                                                   = 0x80004010
	CO_E_INIT_SCM_EXEC_FAILURE                                                       = 0x80004011
	CO_E_INIT_ONLY_SINGLE_THREADED                                                   = 0x80004012
	CO_E_CANT_REMOTE                                                                 = 0x80004013
	CO_E_BAD_SERVER_NAME                                                             = 0x80004014
	CO_E_WRONG_SERVER_IDENTITY                                                       = 0x80004015
	CO_E_OLE1DDE_DISABLED                                                            = 0x80004016
	CO_E_RUNAS_SYNTAX                                                                = 0x80004017
	CO_E_CREATEPROCESS_FAILURE                                                       = 0x80004018
	CO_E_RUNAS_CREATEPROCESS_FAILURE                                                 = 0x80004019
	CO_E_RUNAS_LOGON_FAILURE                                                         = 0x8000401A
	CO_E_LAUNCH_PERMSSION_DENIED                                                     = 0x8000401B
	CO_E_START_SERVICE_FAILURE                                                       = 0x8000401C
	CO_E_REMOTE_COMMUNICATION_FAILURE                                                = 0x8000401D
	CO_E_SERVER_START_TIMEOUT                                                        = 0x8000401E
	CO_E_CLSREG_INCONSISTENT                                                         = 0x8000401F
	CO_E_IIDREG_INCONSISTENT                                                         = 0x80004020
	CO_E_NOT_SUPPORTED                                                               = 0x80004021
	CO_E_RELOAD_DLL                                                                  = 0x80004022
	CO_E_MSI_ERROR                                                                   = 0x80004023
	CO_E_ATTEMPT_TO_CREATE_OUTSIDE_CLIENT_CONTEXT                                    = 0x80004024
	CO_E_SERVER_PAUSED                                                               = 0x80004025
	CO_E_SERVER_NOT_PAUSED                                                           = 0x80004026
	CO_E_CLASS_DISABLED                                                              = 0x80004027
	CO_E_CLRNOTAVAILABLE                                                             = 0x80004028
	CO_E_ASYNC_WORK_REJECTED                                                         = 0x80004029
	CO_E_SERVER_INIT_TIMEOUT                                                         = 0x8000402A
	CO_E_NO_SECCTX_IN_ACTIVATE                                                       = 0x8000402B
	CO_E_TRACKER_CONFIG                                                              = 0x80004030
	CO_E_THREADPOOL_CONFIG                                                           = 0x80004031
	CO_E_SXS_CONFIG                                                                  = 0x80004032
	CO_E_MALFORMED_SPN                                                               = 0x80004033
	S_OK                                                                             = 0
	S_FALSE                                                                          = 1
	OLE_E_FIRST                                                                      = 0x80040000
	OLE_E_LAST                                                                       = 0x800400FF
	OLE_S_FIRST                                                                      = 0x00040000
	OLE_S_LAST                                                                       = 0x000400FF
	OLE_E_OLEVERB                                                                    = 0x80040000
	OLE_E_ADVF                                                                       = 0x80040001
	OLE_E_ENUM_NOMORE                                                                = 0x80040002
	OLE_E_ADVISENOTSUPPORTED                                                         = 0x80040003
	OLE_E_NOCONNECTION                                                               = 0x80040004
	OLE_E_NOTRUNNING                                                                 = 0x80040005
	OLE_E_NOCACHE                                                                    = 0x80040006
	OLE_E_BLANK                                                                      = 0x80040007
	OLE_E_CLASSDIFF                                                                  = 0x80040008
	OLE_E_CANT_GETMONIKER                                                            = 0x80040009
	OLE_E_CANT_BINDTOSOURCE                                                          = 0x8004000A
	OLE_E_STATIC                                                                     = 0x8004000B
	OLE_E_PROMPTSAVECANCELLED                                                        = 0x8004000C
	OLE_E_INVALIDRECT                                                                = 0x8004000D
	OLE_E_WRONGCOMPOBJ                                                               = 0x8004000E
	OLE_E_INVALIDHWND                                                                = 0x8004000F
	OLE_E_NOT_INPLACEACTIVE                                                          = 0x80040010
	OLE_E_CANTCONVERT                                                                = 0x80040011
	OLE_E_NOSTORAGE                                                                  = 0x80040012
	DV_E_FORMATETC                                                                   = 0x80040064
	DV_E_DVTARGETDEVICE                                                              = 0x80040065
	DV_E_STGMEDIUM                                                                   = 0x80040066
	DV_E_STATDATA                                                                    = 0x80040067
	DV_E_LINDEX                                                                      = 0x80040068
	DV_E_TYMED                                                                       = 0x80040069
	DV_E_CLIPFORMAT                                                                  = 0x8004006A
	DV_E_DVASPECT                                                                    = 0x8004006B
	DV_E_DVTARGETDEVICE_SIZE                                                         = 0x8004006C
	DV_E_NOIVIEWOBJECT                                                               = 0x8004006D
	DRAGDROP_E_FIRST                                                                 = 0x80040100
	DRAGDROP_E_LAST                                                                  = 0x8004010F
	DRAGDROP_S_FIRST                                                                 = 0x00040100
	DRAGDROP_S_LAST                                                                  = 0x0004010F
	DRAGDROP_E_NOTREGISTERED                                                         = 0x80040100
	DRAGDROP_E_ALREADYREGISTERED                                                     = 0x80040101
	DRAGDROP_E_INVALIDHWND                                                           = 0x80040102
	CLASSFACTORY_E_FIRST                                                             = 0x80040110
	CLASSFACTORY_E_LAST                                                              = 0x8004011F
	CLASSFACTORY_S_FIRST                                                             = 0x00040110
	CLASSFACTORY_S_LAST                                                              = 0x0004011F
	CLASS_E_NOAGGREGATION                                                            = 0x80040110
	CLASS_E_CLASSNOTAVAILABLE                                                        = 0x80040111
	CLASS_E_NOTLICENSED                                                              = 0x80040112
	MARSHAL_E_FIRST                                                                  = 0x80040120
	MARSHAL_E_LAST                                                                   = 0x8004012F
	MARSHAL_S_FIRST                                                                  = 0x00040120
	MARSHAL_S_LAST                                                                   = 0x0004012F
	DATA_E_FIRST                                                                     = 0x80040130
	DATA_E_LAST                                                                      = 0x8004013F
	DATA_S_FIRST                                                                     = 0x00040130
	DATA_S_LAST                                                                      = 0x0004013F
	VIEW_E_FIRST                                                                     = 0x80040140
	VIEW_E_LAST                                                                      = 0x8004014F
	VIEW_S_FIRST                                                                     = 0x00040140
	VIEW_S_LAST                                                                      = 0x0004014F
	VIEW_E_DRAW                                                                      = 0x80040140
	REGDB_E_FIRST                                                                    = 0x80040150
	REGDB_E_LAST                                                                     = 0x8004015F
	REGDB_S_FIRST                                                                    = 0x00040150
	REGDB_S_LAST                                                                     = 0x0004015F
	REGDB_E_READREGDB                                                                = 0x80040150
	REGDB_E_WRITEREGDB                                                               = 0x80040151
	REGDB_E_KEYMISSING                                                               = 0x80040152
	REGDB_E_INVALIDVALUE                                                             = 0x80040153
	REGDB_E_CLASSNOTREG                                                              = 0x80040154
	REGDB_E_IIDNOTREG                                                                = 0x80040155
	REGDB_E_BADTHREADINGMODEL                                                        = 0x80040156
	CAT_E_FIRST                                                                      = 0x80040160
	CAT_E_LAST                                                                       = 0x80040161
	CAT_E_CATIDNOEXIST                                                               = 0x80040160
	CAT_E_NODESCRIPTION                                                              = 0x80040161
	CS_E_FIRST                                                                       = 0x80040164
	CS_E_LAST                                                                        = 0x8004016F
	CS_E_PACKAGE_NOTFOUND                                                            = 0x80040164
	CS_E_NOT_DELETABLE                                                               = 0x80040165
	CS_E_CLASS_NOTFOUND                                                              = 0x80040166
	CS_E_INVALID_VERSION                                                             = 0x80040167
	CS_E_NO_CLASSSTORE                                                               = 0x80040168
	CS_E_OBJECT_NOTFOUND                                                             = 0x80040169
	CS_E_OBJECT_ALREADY_EXISTS                                                       = 0x8004016A
	CS_E_INVALID_PATH                                                                = 0x8004016B
	CS_E_NETWORK_ERROR                                                               = 0x8004016C
	CS_E_ADMIN_LIMIT_EXCEEDED                                                        = 0x8004016D
	CS_E_SCHEMA_MISMATCH                                                             = 0x8004016E
	CS_E_INTERNAL_ERROR                                                              = 0x8004016F
	CACHE_E_FIRST                                                                    = 0x80040170
	CACHE_E_LAST                                                                     = 0x8004017F
	CACHE_S_FIRST                                                                    = 0x00040170
	CACHE_S_LAST                                                                     = 0x0004017F
	CACHE_E_NOCACHE_UPDATED                                                          = 0x80040170
	OLEOBJ_E_FIRST                                                                   = 0x80040180
	OLEOBJ_E_LAST                                                                    = 0x8004018F
	OLEOBJ_S_FIRST                                                                   = 0x00040180
	OLEOBJ_S_LAST                                                                    = 0x0004018F
	OLEOBJ_E_NOVERBS                                                                 = 0x80040180
	OLEOBJ_E_INVALIDVERB                                                             = 0x80040181
	CLIENTSITE_E_FIRST                                                               = 0x80040190
	CLIENTSITE_E_LAST                                                                = 0x8004019F
	CLIENTSITE_S_FIRST                                                               = 0x00040190
	CLIENTSITE_S_LAST                                                                = 0x0004019F
	INPLACE_E_NOTUNDOABLE                                                            = 0x800401A0
	INPLACE_E_NOTOOLSPACE                                                            = 0x800401A1
	INPLACE_E_FIRST                                                                  = 0x800401A0
	INPLACE_E_LAST                                                                   = 0x800401AF
	INPLACE_S_FIRST                                                                  = 0x000401A0
	INPLACE_S_LAST                                                                   = 0x000401AF
	ENUM_E_FIRST                                                                     = 0x800401B0
	ENUM_E_LAST                                                                      = 0x800401BF
	ENUM_S_FIRST                                                                     = 0x000401B0
	ENUM_S_LAST                                                                      = 0x000401BF
	CONVERT10_E_FIRST                                                                = 0x800401C0
	CONVERT10_E_LAST                                                                 = 0x800401CF
	CONVERT10_S_FIRST                                                                = 0x000401C0
	CONVERT10_S_LAST                                                                 = 0x000401CF
	CONVERT10_E_OLESTREAM_GET                                                        = 0x800401C0
	CONVERT10_E_OLESTREAM_PUT                                                        = 0x800401C1
	CONVERT10_E_OLESTREAM_FMT                                                        = 0x800401C2
	CONVERT10_E_OLESTREAM_BITMAP_TO_DIB                                              = 0x800401C3
	CONVERT10_E_STG_FMT                                                              = 0x800401C4
	CONVERT10_E_STG_NO_STD_STREAM                                                    = 0x800401C5
	CONVERT10_E_STG_DIB_TO_BITMAP                                                    = 0x800401C6
	CLIPBRD_E_FIRST                                                                  = 0x800401D0
	CLIPBRD_E_LAST                                                                   = 0x800401DF
	CLIPBRD_S_FIRST                                                                  = 0x000401D0
	CLIPBRD_S_LAST                                                                   = 0x000401DF
	CLIPBRD_E_CANT_OPEN                                                              = 0x800401D0
	CLIPBRD_E_CANT_EMPTY                                                             = 0x800401D1
	CLIPBRD_E_CANT_SET                                                               = 0x800401D2
	CLIPBRD_E_BAD_DATA                                                               = 0x800401D3
	CLIPBRD_E_CANT_CLOSE                                                             = 0x800401D4
	MK_E_FIRST                                                                       = 0x800401E0
	MK_E_LAST                                                                        = 0x800401EF
	MK_S_FIRST                                                                       = 0x000401E0
	MK_S_LAST                                                                        = 0x000401EF
	MK_E_CONNECTMANUALLY                                                             = 0x800401E0
	MK_E_EXCEEDEDDEADLINE                                                            = 0x800401E1
	MK_E_NEEDGENERIC                                                                 = 0x800401E2
	MK_E_UNAVAILABLE                                                                 = 0x800401E3
	MK_E_SYNTAX                                                                      = 0x800401E4
	MK_E_NOOBJECT                                                                    = 0x800401E5
	MK_E_INVALIDEXTENSION                                                            = 0x800401E6
	MK_E_INTERMEDIATEINTERFACENOTSUPPORTED                                           = 0x800401E7
	MK_E_NOTBINDABLE                                                                 = 0x800401E8
	MK_E_NOTBOUND                                                                    = 0x800401E9
	MK_E_CANTOPENFILE                                                                = 0x800401EA
	MK_E_MUSTBOTHERUSER                                                              = 0x800401EB
	MK_E_NOINVERSE                                                                   = 0x800401EC
	MK_E_NOSTORAGE                                                                   = 0x800401ED
	MK_E_NOPREFIX                                                                    = 0x800401EE
	MK_E_ENUMERATION_FAILED                                                          = 0x800401EF
	CO_E_FIRST                                                                       = 0x800401F0
	CO_E_LAST                                                                        = 0x800401FF
	CO_S_FIRST                                                                       = 0x000401F0
	CO_S_LAST                                                                        = 0x000401FF
	CO_E_NOTINITIALIZED                                                              = 0x800401F0
	CO_E_ALREADYINITIALIZED                                                          = 0x800401F1
	CO_E_CANTDETERMINECLASS                                                          = 0x800401F2
	CO_E_CLASSSTRING                                                                 = 0x800401F3
	CO_E_IIDSTRING                                                                   = 0x800401F4
	CO_E_APPNOTFOUND                                                                 = 0x800401F5
	CO_E_APPSINGLEUSE                                                                = 0x800401F6
	CO_E_ERRORINAPP                                                                  = 0x800401F7
	CO_E_DLLNOTFOUND                                                                 = 0x800401F8
	CO_E_ERRORINDLL                                                                  = 0x800401F9
	CO_E_WRONGOSFORAPP                                                               = 0x800401FA
	CO_E_OBJNOTREG                                                                   = 0x800401FB
	CO_E_OBJISREG                                                                    = 0x800401FC
	CO_E_OBJNOTCONNECTED                                                             = 0x800401FD
	CO_E_APPDIDNTREG                                                                 = 0x800401FE
	CO_E_RELEASED                                                                    = 0x800401FF
	EVENT_E_FIRST                                                                    = 0x80040200
	EVENT_E_LAST                                                                     = 0x8004021F
	EVENT_S_FIRST                                                                    = 0x00040200
	EVENT_S_LAST                                                                     = 0x0004021F
	EVENT_S_SOME_SUBSCRIBERS_FAILED                                                  = 0x00040200
	EVENT_E_ALL_SUBSCRIBERS_FAILED                                                   = 0x80040201
	EVENT_S_NOSUBSCRIBERS                                                            = 0x00040202
	EVENT_E_QUERYSYNTAX                                                              = 0x80040203
	EVENT_E_QUERYFIELD                                                               = 0x80040204
	EVENT_E_INTERNALEXCEPTION                                                        = 0x80040205
	EVENT_E_INTERNALERROR                                                            = 0x80040206
	EVENT_E_INVALID_PER_USER_SID                                                     = 0x80040207
	EVENT_E_USER_EXCEPTION                                                           = 0x80040208
	EVENT_E_TOO_MANY_METHODS                                                         = 0x80040209
	EVENT_E_MISSING_EVENTCLASS                                                       = 0x8004020A
	EVENT_E_NOT_ALL_REMOVED                                                          = 0x8004020B
	EVENT_E_COMPLUS_NOT_INSTALLED                                                    = 0x8004020C
	EVENT_E_CANT_MODIFY_OR_DELETE_UNCONFIGURED_OBJECT                                = 0x8004020D
	EVENT_E_CANT_MODIFY_OR_DELETE_CONFIGURED_OBJECT                                  = 0x8004020E
	EVENT_E_INVALID_EVENT_CLASS_PARTITION                                            = 0x8004020F
	EVENT_E_PER_USER_SID_NOT_LOGGED_ON                                               = 0x80040210
	XACT_E_FIRST                                                                     = 0x8004D000
	XACT_E_LAST                                                                      = 0x8004D02B
	XACT_S_FIRST                                                                     = 0x0004D000
	XACT_S_LAST                                                                      = 0x0004D010
	XACT_E_ALREADYOTHERSINGLEPHASE                                                   = 0x8004D000
	XACT_E_CANTRETAIN                                                                = 0x8004D001
	XACT_E_COMMITFAILED                                                              = 0x8004D002
	XACT_E_COMMITPREVENTED                                                           = 0x8004D003
	XACT_E_HEURISTICABORT                                                            = 0x8004D004
	XACT_E_HEURISTICCOMMIT                                                           = 0x8004D005
	XACT_E_HEURISTICDAMAGE                                                           = 0x8004D006
	XACT_E_HEURISTICDANGER                                                           = 0x8004D007
	XACT_E_ISOLATIONLEVEL                                                            = 0x8004D008
	XACT_E_NOASYNC                                                                   = 0x8004D009
	XACT_E_NOENLIST                                                                  = 0x8004D00A
	XACT_E_NOISORETAIN                                                               = 0x8004D00B
	XACT_E_NORESOURCE                                                                = 0x8004D00C
	XACT_E_NOTCURRENT                                                                = 0x8004D00D
	XACT_E_NOTRANSACTION                                                             = 0x8004D00E
	XACT_E_NOTSUPPORTED                                                              = 0x8004D00F
	XACT_E_UNKNOWNRMGRID                                                             = 0x8004D010
	XACT_E_WRONGSTATE                                                                = 0x8004D011
	XACT_E_WRONGUOW                                                                  = 0x8004D012
	XACT_E_XTIONEXISTS                                                               = 0x8004D013
	XACT_E_NOIMPORTOBJECT                                                            = 0x8004D014
	XACT_E_INVALIDCOOKIE                                                             = 0x8004D015
	XACT_E_INDOUBT                                                                   = 0x8004D016
	XACT_E_NOTIMEOUT                                                                 = 0x8004D017
	XACT_E_ALREADYINPROGRESS                                                         = 0x8004D018
	XACT_E_ABORTED                                                                   = 0x8004D019
	XACT_E_LOGFULL                                                                   = 0x8004D01A
	XACT_E_TMNOTAVAILABLE                                                            = 0x8004D01B
	XACT_E_CONNECTION_DOWN                                                           = 0x8004D01C
	XACT_E_CONNECTION_DENIED                                                         = 0x8004D01D
	XACT_E_REENLISTTIMEOUT                                                           = 0x8004D01E
	XACT_E_TIP_CONNECT_FAILED                                                        = 0x8004D01F
	XACT_E_TIP_PROTOCOL_ERROR                                                        = 0x8004D020
	XACT_E_TIP_PULL_FAILED                                                           = 0x8004D021
	XACT_E_DEST_TMNOTAVAILABLE                                                       = 0x8004D022
	XACT_E_TIP_DISABLED                                                              = 0x8004D023
	XACT_E_NETWORK_TX_DISABLED                                                       = 0x8004D024
	XACT_E_PARTNER_NETWORK_TX_DISABLED                                               = 0x8004D025
	XACT_E_XA_TX_DISABLED                                                            = 0x8004D026
	XACT_E_UNABLE_TO_READ_DTC_CONFIG                                                 = 0x8004D027
	XACT_E_UNABLE_TO_LOAD_DTC_PROXY                                                  = 0x8004D028
	XACT_E_ABORTING                                                                  = 0x8004D029
	XACT_E_PUSH_COMM_FAILURE                                                         = 0x8004D02A
	XACT_E_PULL_COMM_FAILURE                                                         = 0x8004D02B
	XACT_E_LU_TX_DISABLED                                                            = 0x8004D02C
	XACT_E_CLERKNOTFOUND                                                             = 0x8004D080
	XACT_E_CLERKEXISTS                                                               = 0x8004D081
	XACT_E_RECOVERYINPROGRESS                                                        = 0x8004D082
	XACT_E_TRANSACTIONCLOSED                                                         = 0x8004D083
	XACT_E_INVALIDLSN                                                                = 0x8004D084
	XACT_E_REPLAYREQUEST                                                             = 0x8004D085
	XACT_S_ASYNC                                                                     = 0x0004D000
	XACT_S_DEFECT                                                                    = 0x0004D001
	XACT_S_READONLY                                                                  = 0x0004D002
	XACT_S_SOMENORETAIN                                                              = 0x0004D003
	XACT_S_OKINFORM                                                                  = 0x0004D004
	XACT_S_MADECHANGESCONTENT                                                        = 0x0004D005
	XACT_S_MADECHANGESINFORM                                                         = 0x0004D006
	XACT_S_ALLNORETAIN                                                               = 0x0004D007
	XACT_S_ABORTING                                                                  = 0x0004D008
	XACT_S_SINGLEPHASE                                                               = 0x0004D009
	XACT_S_LOCALLY_OK                                                                = 0x0004D00A
	XACT_S_LASTRESOURCEMANAGER                                                       = 0x0004D010
	CONTEXT_E_FIRST                                                                  = 0x8004E000
	CONTEXT_E_LAST                                                                   = 0x8004E02F
	CONTEXT_S_FIRST                                                                  = 0x0004E000
	CONTEXT_S_LAST                                                                   = 0x0004E02F
	CONTEXT_E_ABORTED                                                                = 0x8004E002
	CONTEXT_E_ABORTING                                                               = 0x8004E003
	CONTEXT_E_NOCONTEXT                                                              = 0x8004E004
	CONTEXT_E_WOULD_DEADLOCK                                                         = 0x8004E005
	CONTEXT_E_SYNCH_TIMEOUT                                                          = 0x8004E006
	CONTEXT_E_OLDREF                                                                 = 0x8004E007
	CONTEXT_E_ROLENOTFOUND                                                           = 0x8004E00C
	CONTEXT_E_TMNOTAVAILABLE                                                         = 0x8004E00F
	CO_E_ACTIVATIONFAILED                                                            = 0x8004E021
	CO_E_ACTIVATIONFAILED_EVENTLOGGED                                                = 0x8004E022
	CO_E_ACTIVATIONFAILED_CATALOGERROR                                               = 0x8004E023
	CO_E_ACTIVATIONFAILED_TIMEOUT                                                    = 0x8004E024
	CO_E_INITIALIZATIONFAILED                                                        = 0x8004E025
	CONTEXT_E_NOJIT                                                                  = 0x8004E026
	CONTEXT_E_NOTRANSACTION                                                          = 0x8004E027
	CO_E_THREADINGMODEL_CHANGED                                                      = 0x8004E028
	CO_E_NOIISINTRINSICS                                                             = 0x8004E029
	CO_E_NOCOOKIES                                                                   = 0x8004E02A
	CO_E_DBERROR                                                                     = 0x8004E02B
	CO_E_NOTPOOLED                                                                   = 0x8004E02C
	CO_E_NOTCONSTRUCTED                                                              = 0x8004E02D
	CO_E_NOSYNCHRONIZATION                                                           = 0x8004E02E
	CO_E_ISOLEVELMISMATCH                                                            = 0x8004E02F
	CO_E_CALL_OUT_OF_TX_SCOPE_NOT_ALLOWED                                            = 0x8004E030
	CO_E_EXIT_TRANSACTION_SCOPE_NOT_CALLED                                           = 0x8004E031
	OLE_S_USEREG                                                                     = 0x00040000
	OLE_S_STATIC                                                                     = 0x00040001
	OLE_S_MAC_CLIPFORMAT                                                             = 0x00040002
	DRAGDROP_S_DROP                                                                  = 0x00040100
	DRAGDROP_S_CANCEL                                                                = 0x00040101
	DRAGDROP_S_USEDEFAULTCURSORS                                                     = 0x00040102
	DATA_S_SAMEFORMATETC                                                             = 0x00040130
	VIEW_S_ALREADY_FROZEN                                                            = 0x00040140
	CACHE_S_FORMATETC_NOTSUPPORTED                                                   = 0x00040170
	CACHE_S_SAMECACHE                                                                = 0x00040171
	CACHE_S_SOMECACHES_NOTUPDATED                                                    = 0x00040172
	OLEOBJ_S_INVALIDVERB                                                             = 0x00040180
	OLEOBJ_S_CANNOT_DOVERB_NOW                                                       = 0x00040181
	OLEOBJ_S_INVALIDHWND                                                             = 0x00040182
	INPLACE_S_TRUNCATED                                                              = 0x000401A0
	CONVERT10_S_NO_PRESENTATION                                                      = 0x000401C0
	MK_S_REDUCED_TO_SELF                                                             = 0x000401E2
	MK_S_ME                                                                          = 0x000401E4
	MK_S_HIM                                                                         = 0x000401E5
	MK_S_US                                                                          = 0x000401E6
	MK_S_MONIKERALREADYREGISTERED                                                    = 0x000401E7
	SCHED_S_TASK_READY                                                               = 0x00041300
	SCHED_S_TASK_RUNNING                                                             = 0x00041301
	SCHED_S_TASK_DISABLED                                                            = 0x00041302
	SCHED_S_TASK_HAS_NOT_RUN                                                         = 0x00041303
	SCHED_S_TASK_NO_MORE_RUNS                                                        = 0x00041304
	SCHED_S_TASK_NOT_SCHEDULED                                                       = 0x00041305
	SCHED_S_TASK_TERMINATED                                                          = 0x00041306
	SCHED_S_TASK_NO_VALID_TRIGGERS                                                   = 0x00041307
	SCHED_S_EVENT_TRIGGER                                                            = 0x00041308
	SCHED_E_TRIGGER_NOT_FOUND                                                        = 0x80041309
	SCHED_E_TASK_NOT_READY                                                           = 0x8004130A
	SCHED_E_TASK_NOT_RUNNING                                                         = 0x8004130B
	SCHED_E_SERVICE_NOT_INSTALLED                                                    = 0x8004130C
	SCHED_E_CANNOT_OPEN_TASK                                                         = 0x8004130D
	SCHED_E_INVALID_TASK                                                             = 0x8004130E
	SCHED_E_ACCOUNT_INFORMATION_NOT_SET                                              = 0x8004130F
	SCHED_E_ACCOUNT_NAME_NOT_FOUND                                                   = 0x80041310
	SCHED_E_ACCOUNT_DBASE_CORRUPT                                                    = 0x80041311
	SCHED_E_NO_SECURITY_SERVICES                                                     = 0x80041312
	SCHED_E_UNKNOWN_OBJECT_VERSION                                                   = 0x80041313
	SCHED_E_UNSUPPORTED_ACCOUNT_OPTION                                               = 0x80041314
	SCHED_E_SERVICE_NOT_RUNNING                                                      = 0x80041315
	SCHED_E_UNEXPECTEDNODE                                                           = 0x80041316
	SCHED_E_NAMESPACE                                                                = 0x80041317
	SCHED_E_INVALIDVALUE                                                             = 0x80041318
	SCHED_E_MISSINGNODE                                                              = 0x80041319
	SCHED_E_MALFORMEDXML                                                             = 0x8004131A
	SCHED_S_SOME_TRIGGERS_FAILED                                                     = 0x0004131B
	SCHED_S_BATCH_LOGON_PROBLEM                                                      = 0x0004131C
	SCHED_E_TOO_MANY_NODES                                                           = 0x8004131D
	SCHED_E_PAST_END_BOUNDARY                                                        = 0x8004131E
	SCHED_E_ALREADY_RUNNING                                                          = 0x8004131F
	SCHED_E_USER_NOT_LOGGED_ON                                                       = 0x80041320
	SCHED_E_INVALID_TASK_HASH                                                        = 0x80041321
	SCHED_E_SERVICE_NOT_AVAILABLE                                                    = 0x80041322
	SCHED_E_SERVICE_TOO_BUSY                                                         = 0x80041323
	SCHED_E_TASK_ATTEMPTED                                                           = 0x80041324
	SCHED_S_TASK_QUEUED                                                              = 0x00041325
	SCHED_E_TASK_DISABLED                                                            = 0x80041326
	SCHED_E_TASK_NOT_V1_COMPAT                                                       = 0x80041327
	SCHED_E_START_ON_DEMAND                                                          = 0x80041328
	CO_E_CLASS_CREATE_FAILED                                                         = 0x80080001
	CO_E_SCM_ERROR                                                                   = 0x80080002
	CO_E_SCM_RPC_FAILURE                                                             = 0x80080003
	CO_E_BAD_PATH                                                                    = 0x80080004
	CO_E_SERVER_EXEC_FAILURE                                                         = 0x80080005
	CO_E_OBJSRV_RPC_FAILURE                                                          = 0x80080006
	MK_E_NO_NORMALIZED                                                               = 0x80080007
	CO_E_SERVER_STOPPING                                                             = 0x80080008
	MEM_E_INVALID_ROOT                                                               = 0x80080009
	MEM_E_INVALID_LINK                                                               = 0x80080010
	MEM_E_INVALID_SIZE                                                               = 0x80080011
	CO_S_NOTALLINTERFACES                                                            = 0x00080012
	CO_S_MACHINENAMENOTFOUND                                                         = 0x00080013
	CO_E_MISSING_DISPLAYNAME                                                         = 0x80080015
	CO_E_RUNAS_VALUE_MUST_BE_AAA                                                     = 0x80080016
	CO_E_ELEVATION_DISABLED                                                          = 0x80080017
	DISP_E_UNKNOWNINTERFACE                                                          = 0x80020001
	DISP_E_MEMBERNOTFOUND                                                            = 0x80020003
	DISP_E_PARAMNOTFOUND                                                             = 0x80020004
	DISP_E_TYPEMISMATCH                                                              = 0x80020005
	DISP_E_UNKNOWNNAME                                                               = 0x80020006
	DISP_E_NONAMEDARGS                                                               = 0x80020007
	DISP_E_BADVARTYPE                                                                = 0x80020008
	DISP_E_EXCEPTION                                                                 = 0x80020009
	DISP_E_OVERFLOW                                                                  = 0x8002000A
	DISP_E_BADINDEX                                                                  = 0x8002000B
	DISP_E_UNKNOWNLCID                                                               = 0x8002000C
	DISP_E_ARRAYISLOCKED                                                             = 0x8002000D
	DISP_E_BADPARAMCOUNT                                                             = 0x8002000E
	DISP_E_PARAMNOTOPTIONAL                                                          = 0x8002000F
	DISP_E_BADCALLEE                                                                 = 0x80020010
	DISP_E_NOTACOLLECTION                                                            = 0x80020011
	DISP_E_DIVBYZERO                                                                 = 0x80020012
	DISP_E_BUFFERTOOSMALL                                                            = 0x80020013
	TYPE_E_BUFFERTOOSMALL                                                            = 0x80028016
	TYPE_E_FIELDNOTFOUND                                                             = 0x80028017
	TYPE_E_INVDATAREAD                                                               = 0x80028018
	TYPE_E_UNSUPFORMAT                                                               = 0x80028019
	TYPE_E_REGISTRYACCESS                                                            = 0x8002801C
	TYPE_E_LIBNOTREGISTERED                                                          = 0x8002801D
	TYPE_E_UNDEFINEDTYPE                                                             = 0x80028027
	TYPE_E_QUALIFIEDNAMEDISALLOWED                                                   = 0x80028028
	TYPE_E_INVALIDSTATE                                                              = 0x80028029
	TYPE_E_WRONGTYPEKIND                                                             = 0x8002802A
	TYPE_E_ELEMENTNOTFOUND                                                           = 0x8002802B
	TYPE_E_AMBIGUOUSNAME                                                             = 0x8002802C
	TYPE_E_NAMECONFLICT                                                              = 0x8002802D
	TYPE_E_UNKNOWNLCID                                                               = 0x8002802E
	TYPE_E_DLLFUNCTIONNOTFOUND                                                       = 0x8002802F
	TYPE_E_BADMODULEKIND                                                             = 0x800288BD
	TYPE_E_SIZETOOBIG                                                                = 0x800288C5
	TYPE_E_DUPLICATEID                                                               = 0x800288C6
	TYPE_E_INVALIDID                                                                 = 0x800288CF
	TYPE_E_TYPEMISMATCH                                                              = 0x80028CA0
	TYPE_E_OUTOFBOUNDS                                                               = 0x80028CA1
	TYPE_E_IOERROR                                                                   = 0x80028CA2
	TYPE_E_CANTCREATETMPFILE                                                         = 0x80028CA3
	TYPE_E_CANTLOADLIBRARY                                                           = 0x80029C4A
	TYPE_E_INCONSISTENTPROPFUNCS                                                     = 0x80029C83
	TYPE_E_CIRCULARTYPE                                                              = 0x80029C84
	STG_E_INVALIDFUNCTION                                                            = 0x80030001
	STG_E_FILENOTFOUND                                                               = 0x80030002
	STG_E_PATHNOTFOUND                                                               = 0x80030003
	STG_E_TOOMANYOPENFILES                                                           = 0x80030004
	STG_E_ACCESSDENIED                                                               = 0x80030005
	STG_E_INVALIDHANDLE                                                              = 0x80030006
	STG_E_INSUFFICIENTMEMORY                                                         = 0x80030008
	STG_E_INVALIDPOINTER                                                             = 0x80030009
	STG_E_NOMOREFILES                                                                = 0x80030012
	STG_E_DISKISWRITEPROTECTED                                                       = 0x80030013
	STG_E_SEEKERROR                                                                  = 0x80030019
	STG_E_WRITEFAULT                                                                 = 0x8003001D
	STG_E_READFAULT                                                                  = 0x8003001E
	STG_E_SHAREVIOLATION                                                             = 0x80030020
	STG_E_LOCKVIOLATION                                                              = 0x80030021
	STG_E_FILEALREADYEXISTS                                                          = 0x80030050
	STG_E_INVALIDPARAMETER                                                           = 0x80030057
	STG_E_MEDIUMFULL                                                                 = 0x80030070
	STG_E_PROPSETMISMATCHED                                                          = 0x800300F0
	STG_E_ABNORMALAPIEXIT                                                            = 0x800300FA
	STG_E_INVALIDHEADER                                                              = 0x800300FB
	STG_E_INVALIDNAME                                                                = 0x800300FC
	STG_E_UNKNOWN                                                                    = 0x800300FD
	STG_E_UNIMPLEMENTEDFUNCTION                                                      = 0x800300FE
	STG_E_INVALIDFLAG                                                                = 0x800300FF
	STG_E_INUSE                                                                      = 0x80030100
	STG_E_NOTCURRENT                                                                 = 0x80030101
	STG_E_REVERTED                                                                   = 0x80030102
	STG_E_CANTSAVE                                                                   = 0x80030103
	STG_E_OLDFORMAT                                                                  = 0x80030104
	STG_E_OLDDLL                                                                     = 0x80030105
	STG_E_SHAREREQUIRED                                                              = 0x80030106
	STG_E_NOTFILEBASEDSTORAGE                                                        = 0x80030107
	STG_E_EXTANTMARSHALLINGS                                                         = 0x80030108
	STG_E_DOCFILECORRUPT                                                             = 0x80030109
	STG_E_BADBASEADDRESS                                                             = 0x80030110
	STG_E_DOCFILETOOLARGE                                                            = 0x80030111
	STG_E_NOTSIMPLEFORMAT                                                            = 0x80030112
	STG_E_INCOMPLETE                                                                 = 0x80030201
	STG_E_TERMINATED                                                                 = 0x80030202
	STG_S_CONVERTED                                                                  = 0x00030200
	STG_S_BLOCK                                                                      = 0x00030201
	STG_S_RETRYNOW                                                                   = 0x00030202
	STG_S_MONITORING                                                                 = 0x00030203
	STG_S_MULTIPLEOPENS                                                              = 0x00030204
	STG_S_CONSOLIDATIONFAILED                                                        = 0x00030205
	STG_S_CANNOTCONSOLIDATE                                                          = 0x00030206
	STG_E_STATUS_COPY_PROTECTION_FAILURE                                             = 0x80030305
	STG_E_CSS_AUTHENTICATION_FAILURE                                                 = 0x80030306
	STG_E_CSS_KEY_NOT_PRESENT                                                        = 0x80030307
	STG_E_CSS_KEY_NOT_ESTABLISHED                                                    = 0x80030308
	STG_E_CSS_SCRAMBLED_SECTOR                                                       = 0x80030309
	STG_E_CSS_REGION_MISMATCH                                                        = 0x8003030A
	STG_E_RESETS_EXHAUSTED                                                           = 0x8003030B
	RPC_E_CALL_REJECTED                                                              = 0x80010001
	RPC_E_CALL_CANCELED                                                              = 0x80010002
	RPC_E_CANTPOST_INSENDCALL                                                        = 0x80010003
	RPC_E_CANTCALLOUT_INASYNCCALL                                                    = 0x80010004
	RPC_E_CANTCALLOUT_INEXTERNALCALL                                                 = 0x80010005
	RPC_E_CONNECTION_TERMINATED                                                      = 0x80010006
	RPC_E_SERVER_DIED                                                                = 0x80010007
	RPC_E_CLIENT_DIED                                                                = 0x80010008
	RPC_E_INVALID_DATAPACKET                                                         = 0x80010009
	RPC_E_CANTTRANSMIT_CALL                                                          = 0x8001000A
	RPC_E_CLIENT_CANTMARSHAL_DATA                                                    = 0x8001000B
	RPC_E_CLIENT_CANTUNMARSHAL_DATA                                                  = 0x8001000C
	RPC_E_SERVER_CANTMARSHAL_DATA                                                    = 0x8001000D
	RPC_E_SERVER_CANTUNMARSHAL_DATA                                                  = 0x8001000E
	RPC_E_INVALID_DATA                                                               = 0x8001000F
	RPC_E_INVALID_PARAMETER                                                          = 0x80010010
	RPC_E_CANTCALLOUT_AGAIN                                                          = 0x80010011
	RPC_E_SERVER_DIED_DNE                                                            = 0x80010012
	RPC_E_SYS_CALL_FAILED                                                            = 0x80010100
	RPC_E_OUT_OF_RESOURCES                                                           = 0x80010101
	RPC_E_ATTEMPTED_MULTITHREAD                                                      = 0x80010102
	RPC_E_NOT_REGISTERED                                                             = 0x80010103
	RPC_E_FAULT                                                                      = 0x80010104
	RPC_E_SERVERFAULT                                                                = 0x80010105
	RPC_E_CHANGED_MODE                                                               = 0x80010106
	RPC_E_INVALIDMETHOD                                                              = 0x80010107
	RPC_E_DISCONNECTED                                                               = 0x80010108
	RPC_E_RETRY                                                                      = 0x80010109
	RPC_E_SERVERCALL_RETRYLATER                                                      = 0x8001010A
	RPC_E_SERVERCALL_REJECTED                                                        = 0x8001010B
	RPC_E_INVALID_CALLDATA                                                           = 0x8001010C
	RPC_E_CANTCALLOUT_ININPUTSYNCCALL                                                = 0x8001010D
	RPC_E_WRONG_THREAD                                                               = 0x8001010E
	RPC_E_THREAD_NOT_INIT                                                            = 0x8001010F
	RPC_E_VERSION_MISMATCH                                                           = 0x80010110
	RPC_E_INVALID_HEADER                                                             = 0x80010111
	RPC_E_INVALID_EXTENSION                                                          = 0x80010112
	RPC_E_INVALID_IPID                                                               = 0x80010113
	RPC_E_INVALID_OBJECT                                                             = 0x80010114
	RPC_S_CALLPENDING                                                                = 0x80010115
	RPC_S_WAITONTIMER                                                                = 0x80010116
	RPC_E_CALL_COMPLETE                                                              = 0x80010117
	RPC_E_UNSECURE_CALL                                                              = 0x80010118
	RPC_E_TOO_LATE                                                                   = 0x80010119
	RPC_E_NO_GOOD_SECURITY_PACKAGES                                                  = 0x8001011A
	RPC_E_ACCESS_DENIED                                                              = 0x8001011B
	RPC_E_REMOTE_DISABLED                                                            = 0x8001011C
	RPC_E_INVALID_OBJREF                                                             = 0x8001011D
	RPC_E_NO_CONTEXT                                                                 = 0x8001011E
	RPC_E_TIMEOUT                                                                    = 0x8001011F
	RPC_E_NO_SYNC                                                                    = 0x80010120
	RPC_E_FULLSIC_REQUIRED                                                           = 0x80010121
	RPC_E_INVALID_STD_NAME                                                           = 0x80010122
	CO_E_FAILEDTOIMPERSONATE                                                         = 0x80010123
	CO_E_FAILEDTOGETSECCTX                                                           = 0x80010124
	CO_E_FAILEDTOOPENTHREADTOKEN                                                     = 0x80010125
	CO_E_FAILEDTOGETTOKENINFO                                                        = 0x80010126
	CO_E_TRUSTEEDOESNTMATCHCLIENT                                                    = 0x80010127
	CO_E_FAILEDTOQUERYCLIENTBLANKET                                                  = 0x80010128
	CO_E_FAILEDTOSETDACL                                                             = 0x80010129
	CO_E_ACCESSCHECKFAILED                                                           = 0x8001012A
	CO_E_NETACCESSAPIFAILED                                                          = 0x8001012B
	CO_E_WRONGTRUSTEENAMESYNTAX                                                      = 0x8001012C
	CO_E_INVALIDSID                                                                  = 0x8001012D
	CO_E_CONVERSIONFAILED                                                            = 0x8001012E
	CO_E_NOMATCHINGSIDFOUND                                                          = 0x8001012F
	CO_E_LOOKUPACCSIDFAILED                                                          = 0x80010130
	CO_E_NOMATCHINGNAMEFOUND                                                         = 0x80010131
	CO_E_LOOKUPACCNAMEFAILED                                                         = 0x80010132
	CO_E_SETSERLHNDLFAILED                                                           = 0x80010133
	CO_E_FAILEDTOGETWINDIR                                                           = 0x80010134
	CO_E_PATHTOOLONG                                                                 = 0x80010135
	CO_E_FAILEDTOGENUUID                                                             = 0x80010136
	CO_E_FAILEDTOCREATEFILE                                                          = 0x80010137
	CO_E_FAILEDTOCLOSEHANDLE                                                         = 0x80010138
	CO_E_EXCEEDSYSACLLIMIT                                                           = 0x80010139
	CO_E_ACESINWRONGORDER                                                            = 0x8001013A
	CO_E_INCOMPATIBLESTREAMVERSION                                                   = 0x8001013B
	CO_E_FAILEDTOOPENPROCESSTOKEN                                                    = 0x8001013C
	CO_E_DECODEFAILED                                                                = 0x8001013D
	CO_E_ACNOTINITIALIZED                                                            = 0x8001013F
	CO_E_CANCEL_DISABLED                                                             = 0x80010140
	RPC_E_UNEXPECTED                                                                 = 0x8001FFFF
	ERROR_AUDITING_DISABLED                                                          = 0xC0090001
	ERROR_ALL_SIDS_FILTERED                                                          = 0xC0090002
	ERROR_BIZRULES_NOT_ENABLED                                                       = 0xC0090003
	NTE_BAD_UID                                                                      = 0x80090001
	NTE_BAD_HASH                                                                     = 0x80090002
	NTE_BAD_KEY                                                                      = 0x80090003
	NTE_BAD_LEN                                                                      = 0x80090004
	NTE_BAD_DATA                                                                     = 0x80090005
	NTE_BAD_SIGNATURE                                                                = 0x80090006
	NTE_BAD_VER                                                                      = 0x80090007
	NTE_BAD_ALGID                                                                    = 0x80090008
	NTE_BAD_FLAGS                                                                    = 0x80090009
	NTE_BAD_TYPE                                                                     = 0x8009000A
	NTE_BAD_KEY_STATE                                                                = 0x8009000B
	NTE_BAD_HASH_STATE                                                               = 0x8009000C
	NTE_NO_KEY                                                                       = 0x8009000D
	NTE_NO_MEMORY                                                                    = 0x8009000E
	NTE_EXISTS                                                                       = 0x8009000F
	NTE_PERM                                                                         = 0x80090010
	NTE_NOT_FOUND                                                                    = 0x80090011
	NTE_DOUBLE_ENCRYPT                                                               = 0x80090012
	NTE_BAD_PROVIDER                                                                 = 0x80090013
	NTE_BAD_PROV_TYPE                                                                = 0x80090014
	NTE_BAD_PUBLIC_KEY                                                               = 0x80090015
	NTE_BAD_KEYSET                                                                   = 0x80090016
	NTE_PROV_TYPE_NOT_DEF                                                            = 0x80090017
	NTE_PROV_TYPE_ENTRY_BAD                                                          = 0x80090018
	NTE_KEYSET_NOT_DEF                                                               = 0x80090019
	NTE_KEYSET_ENTRY_BAD                                                             = 0x8009001A
	NTE_PROV_TYPE_NO_MATCH                                                           = 0x8009001B
	NTE_SIGNATURE_FILE_BAD                                                           = 0x8009001C
	NTE_PROVIDER_DLL_FAIL                                                            = 0x8009001D
	NTE_PROV_DLL_NOT_FOUND                                                           = 0x8009001E
	NTE_BAD_KEYSET_PARAM                                                             = 0x8009001F
	NTE_FAIL                                                                         = 0x80090020
	NTE_SYS_ERR                                                                      = 0x80090021
	NTE_SILENT_CONTEXT                                                               = 0x80090022
	NTE_TOKEN_KEYSET_STORAGE_FULL                                                    = 0x80090023
	NTE_TEMPORARY_PROFILE                                                            = 0x80090024
	NTE_FIXEDPARAMETER                                                               = 0x80090025
	NTE_INVALID_HANDLE                                                               = 0x80090026
	NTE_INVALID_PARAMETER                                                            = 0x80090027
	NTE_BUFFER_TOO_SMALL                                                             = 0x80090028
	NTE_NOT_SUPPORTED                                                                = 0x80090029
	NTE_NO_MORE_ITEMS                                                                = 0x8009002A
	NTE_BUFFERS_OVERLAP                                                              = 0x8009002B
	NTE_DECRYPTION_FAILURE                                                           = 0x8009002C
	NTE_INTERNAL_ERROR                                                               = 0x8009002D
	NTE_UI_REQUIRED                                                                  = 0x8009002E
	NTE_HMAC_NOT_SUPPORTED                                                           = 0x8009002F
	SEC_E_INSUFFICIENT_MEMORY                                                        = 0x80090300
	SEC_E_INVALID_HANDLE                                                             = 0x80090301
	SEC_E_UNSUPPORTED_FUNCTION                                                       = 0x80090302
	SEC_E_TARGET_UNKNOWN                                                             = 0x80090303
	SEC_E_INTERNAL_ERROR                                                             = 0x80090304
	SEC_E_SECPKG_NOT_FOUND                                                           = 0x80090305
	SEC_E_NOT_OWNER                                                                  = 0x80090306
	SEC_E_CANNOT_INSTALL                                                             = 0x80090307
	SEC_E_INVALID_TOKEN                                                              = 0x80090308
	SEC_E_CANNOT_PACK                                                                = 0x80090309
	SEC_E_QOP_NOT_SUPPORTED                                                          = 0x8009030A
	SEC_E_NO_IMPERSONATION                                                           = 0x8009030B
	SEC_E_LOGON_DENIED                                                               = 0x8009030C
	SEC_E_UNKNOWN_CREDENTIALS                                                        = 0x8009030D
	SEC_E_NO_CREDENTIALS                                                             = 0x8009030E
	SEC_E_MESSAGE_ALTERED                                                            = 0x8009030F
	SEC_E_OUT_OF_SEQUENCE                                                            = 0x80090310
	SEC_E_NO_AUTHENTICATING_AUTHORITY                                                = 0x80090311
	SEC_I_CONTINUE_NEEDED                                                            = 0x00090312
	SEC_I_COMPLETE_NEEDED                                                            = 0x00090313
	SEC_I_COMPLETE_AND_CONTINUE                                                      = 0x00090314
	SEC_I_LOCAL_LOGON                                                                = 0x00090315
	SEC_E_BAD_PKGID                                                                  = 0x80090316
	SEC_E_CONTEXT_EXPIRED                                                            = 0x80090317
	SEC_I_CONTEXT_EXPIRED                                                            = 0x00090317
	SEC_E_INCOMPLETE_MESSAGE                                                         = 0x80090318
	SEC_E_INCOMPLETE_CREDENTIALS                                                     = 0x80090320
	SEC_E_BUFFER_TOO_SMALL                                                           = 0x80090321
	SEC_I_INCOMPLETE_CREDENTIALS                                                     = 0x00090320
	SEC_I_RENEGOTIATE                                                                = 0x00090321
	SEC_E_WRONG_PRINCIPAL                                                            = 0x80090322
	SEC_I_NO_LSA_CONTEXT                                                             = 0x00090323
	SEC_E_TIME_SKEW                                                                  = 0x80090324
	SEC_E_UNTRUSTED_ROOT                                                             = 0x80090325
	SEC_E_ILLEGAL_MESSAGE                                                            = 0x80090326
	SEC_E_CERT_UNKNOWN                                                               = 0x80090327
	SEC_E_CERT_EXPIRED                                                               = 0x80090328
	SEC_E_ENCRYPT_FAILURE                                                            = 0x80090329
	SEC_E_DECRYPT_FAILURE                                                            = 0x80090330
	SEC_E_ALGORITHM_MISMATCH                                                         = 0x80090331
	SEC_E_SECURITY_QOS_FAILED                                                        = 0x80090332
	SEC_E_UNFINISHED_CONTEXT_DELETED                                                 = 0x80090333
	SEC_E_NO_TGT_REPLY                                                               = 0x80090334
	SEC_E_NO_IP_ADDRESSES                                                            = 0x80090335
	SEC_E_WRONG_CREDENTIAL_HANDLE                                                    = 0x80090336
	SEC_E_CRYPTO_SYSTEM_INVALID                                                      = 0x80090337
	SEC_E_MAX_REFERRALS_EXCEEDED                                                     = 0x80090338
	SEC_E_MUST_BE_KDC                                                                = 0x80090339
	SEC_E_STRONG_CRYPTO_NOT_SUPPORTED                                                = 0x8009033A
	SEC_E_TOO_MANY_PRINCIPALS                                                        = 0x8009033B
	SEC_E_NO_PA_DATA                                                                 = 0x8009033C
	SEC_E_PKINIT_NAME_MISMATCH                                                       = 0x8009033D
	SEC_E_SMARTCARD_LOGON_REQUIRED                                                   = 0x8009033E
	SEC_E_SHUTDOWN_IN_PROGRESS                                                       = 0x8009033F
	SEC_E_KDC_INVALID_REQUEST                                                        = 0x80090340
	SEC_E_KDC_UNABLE_TO_REFER                                                        = 0x80090341
	SEC_E_KDC_UNKNOWN_ETYPE                                                          = 0x80090342
	SEC_E_UNSUPPORTED_PREAUTH                                                        = 0x80090343
	SEC_E_DELEGATION_REQUIRED                                                        = 0x80090345
	SEC_E_BAD_BINDINGS                                                               = 0x80090346
	SEC_E_MULTIPLE_ACCOUNTS                                                          = 0x80090347
	SEC_E_NO_KERB_KEY                                                                = 0x80090348
	SEC_E_CERT_WRONG_USAGE                                                           = 0x80090349
	SEC_E_DOWNGRADE_DETECTED                                                         = 0x80090350
	SEC_E_SMARTCARD_CERT_REVOKED                                                     = 0x80090351
	SEC_E_ISSUING_CA_UNTRUSTED                                                       = 0x80090352
	SEC_E_REVOCATION_OFFLINE_C                                                       = 0x80090353
	SEC_E_PKINIT_CLIENT_FAILURE                                                      = 0x80090354
	SEC_E_SMARTCARD_CERT_EXPIRED                                                     = 0x80090355
	SEC_E_NO_S4U_PROT_SUPPORT                                                        = 0x80090356
	SEC_E_CROSSREALM_DELEGATION_FAILURE                                              = 0x80090357
	SEC_E_REVOCATION_OFFLINE_KDC                                                     = 0x80090358
	SEC_E_ISSUING_CA_UNTRUSTED_KDC                                                   = 0x80090359
	SEC_E_KDC_CERT_EXPIRED                                                           = 0x8009035A
	SEC_E_KDC_CERT_REVOKED                                                           = 0x8009035B
	SEC_I_SIGNATURE_NEEDED                                                           = 0x0009035C
	SEC_E_INVALID_PARAMETER                                                          = 0x8009035D
	SEC_E_DELEGATION_POLICY                                                          = 0x8009035E
	SEC_E_POLICY_NLTM_ONLY                                                           = 0x8009035F
	SEC_I_NO_RENEGOTIATION                                                           = 0x00090360
	SEC_E_NO_CONTEXT                                                                 = 0x80090361
	SEC_E_PKU2U_CERT_FAILURE                                                         = 0x80090362
	SEC_E_MUTUAL_AUTH_FAILED                                                         = 0x80090363
	CRYPT_E_MSG_ERROR                                                                = 0x80091001
	CRYPT_E_UNKNOWN_ALGO                                                             = 0x80091002
	CRYPT_E_OID_FORMAT                                                               = 0x80091003
	CRYPT_E_INVALID_MSG_TYPE                                                         = 0x80091004
	CRYPT_E_UNEXPECTED_ENCODING                                                      = 0x80091005
	CRYPT_E_AUTH_ATTR_MISSING                                                        = 0x80091006
	CRYPT_E_HASH_VALUE                                                               = 0x80091007
	CRYPT_E_INVALID_INDEX                                                            = 0x80091008
	CRYPT_E_ALREADY_DECRYPTED                                                        = 0x80091009
	CRYPT_E_NOT_DECRYPTED                                                            = 0x8009100A
	CRYPT_E_RECIPIENT_NOT_FOUND                                                      = 0x8009100B
	CRYPT_E_CONTROL_TYPE                                                             = 0x8009100C
	CRYPT_E_ISSUER_SERIALNUMBER                                                      = 0x8009100D
	CRYPT_E_SIGNER_NOT_FOUND                                                         = 0x8009100E
	CRYPT_E_ATTRIBUTES_MISSING                                                       = 0x8009100F
	CRYPT_E_STREAM_MSG_NOT_READY                                                     = 0x80091010
	CRYPT_E_STREAM_INSUFFICIENT_DATA                                                 = 0x80091011
	CRYPT_I_NEW_PROTECTION_REQUIRED                                                  = 0x00091012
	CRYPT_E_BAD_LEN                                                                  = 0x80092001
	CRYPT_E_BAD_ENCODE                                                               = 0x80092002
	CRYPT_E_FILE_ERROR                                                               = 0x80092003
	CRYPT_E_NOT_FOUND                                                                = 0x80092004
	CRYPT_E_EXISTS                                                                   = 0x80092005
	CRYPT_E_NO_PROVIDER                                                              = 0x80092006
	CRYPT_E_SELF_SIGNED                                                              = 0x80092007
	CRYPT_E_DELETED_PREV                                                             = 0x80092008
	CRYPT_E_NO_MATCH                                                                 = 0x80092009
	CRYPT_E_UNEXPECTED_MSG_TYPE                                                      = 0x8009200A
	CRYPT_E_NO_KEY_PROPERTY                                                          = 0x8009200B
	CRYPT_E_NO_DECRYPT_CERT                                                          = 0x8009200C
	CRYPT_E_BAD_MSG                                                                  = 0x8009200D
	CRYPT_E_NO_SIGNER                                                                = 0x8009200E
	CRYPT_E_PENDING_CLOSE                                                            = 0x8009200F
	CRYPT_E_REVOKED                                                                  = 0x80092010
	CRYPT_E_NO_REVOCATION_DLL                                                        = 0x80092011
	CRYPT_E_NO_REVOCATION_CHECK                                                      = 0x80092012
	CRYPT_E_REVOCATION_OFFLINE                                                       = 0x80092013
	CRYPT_E_NOT_IN_REVOCATION_DATABASE                                               = 0x80092014
	CRYPT_E_INVALID_NUMERIC_STRING                                                   = 0x80092020
	CRYPT_E_INVALID_PRINTABLE_STRING                                                 = 0x80092021
	CRYPT_E_INVALID_IA5_STRING                                                       = 0x80092022
	CRYPT_E_INVALID_X500_STRING                                                      = 0x80092023
	CRYPT_E_NOT_CHAR_STRING                                                          = 0x80092024
	CRYPT_E_FILERESIZED                                                              = 0x80092025
	CRYPT_E_SECURITY_SETTINGS                                                        = 0x80092026
	CRYPT_E_NO_VERIFY_USAGE_DLL                                                      = 0x80092027
	CRYPT_E_NO_VERIFY_USAGE_CHECK                                                    = 0x80092028
	CRYPT_E_VERIFY_USAGE_OFFLINE                                                     = 0x80092029
	CRYPT_E_NOT_IN_CTL                                                               = 0x8009202A
	CRYPT_E_NO_TRUSTED_SIGNER                                                        = 0x8009202B
	CRYPT_E_MISSING_PUBKEY_PARA                                                      = 0x8009202C
	CRYPT_E_OSS_ERROR                                                                = 0x80093000
	OSS_MORE_BUF                                                                     = 0x80093001
	OSS_NEGATIVE_UINTEGER                                                            = 0x80093002
	OSS_PDU_RANGE                                                                    = 0x80093003
	OSS_MORE_INPUT                                                                   = 0x80093004
	OSS_DATA_ERROR                                                                   = 0x80093005
	OSS_BAD_ARG                                                                      = 0x80093006
	OSS_BAD_VERSION                                                                  = 0x80093007
	OSS_OUT_MEMORY                                                                   = 0x80093008
	OSS_PDU_MISMATCH                                                                 = 0x80093009
	OSS_LIMITED                                                                      = 0x8009300A
	OSS_BAD_PTR                                                                      = 0x8009300B
	OSS_BAD_TIME                                                                     = 0x8009300C
	OSS_INDEFINITE_NOT_SUPPORTED                                                     = 0x8009300D
	OSS_MEM_ERROR                                                                    = 0x8009300E
	OSS_BAD_TABLE                                                                    = 0x8009300F
	OSS_TOO_LONG                                                                     = 0x80093010
	OSS_CONSTRAINT_VIOLATED                                                          = 0x80093011
	OSS_FATAL_ERROR                                                                  = 0x80093012
	OSS_ACCESS_SERIALIZATION_ERROR                                                   = 0x80093013
	OSS_NULL_TBL                                                                     = 0x80093014
	OSS_NULL_FCN                                                                     = 0x80093015
	OSS_BAD_ENCRULES                                                                 = 0x80093016
	OSS_UNAVAIL_ENCRULES                                                             = 0x80093017
	OSS_CANT_OPEN_TRACE_WINDOW                                                       = 0x80093018
	OSS_UNIMPLEMENTED                                                                = 0x80093019
	OSS_OID_DLL_NOT_LINKED                                                           = 0x8009301A
	OSS_CANT_OPEN_TRACE_FILE                                                         = 0x8009301B
	OSS_TRACE_FILE_ALREADY_OPEN                                                      = 0x8009301C
	OSS_TABLE_MISMATCH                                                               = 0x8009301D
	OSS_TYPE_NOT_SUPPORTED                                                           = 0x8009301E
	OSS_REAL_DLL_NOT_LINKED                                                          = 0x8009301F
	OSS_REAL_CODE_NOT_LINKED                                                         = 0x80093020
	OSS_OUT_OF_RANGE                                                                 = 0x80093021
	OSS_COPIER_DLL_NOT_LINKED                                                        = 0x80093022
	OSS_CONSTRAINT_DLL_NOT_LINKED                                                    = 0x80093023
	OSS_COMPARATOR_DLL_NOT_LINKED                                                    = 0x80093024
	OSS_COMPARATOR_CODE_NOT_LINKED                                                   = 0x80093025
	OSS_MEM_MGR_DLL_NOT_LINKED                                                       = 0x80093026
	OSS_PDV_DLL_NOT_LINKED                                                           = 0x80093027
	OSS_PDV_CODE_NOT_LINKED                                                          = 0x80093028
	OSS_API_DLL_NOT_LINKED                                                           = 0x80093029
	OSS_BERDER_DLL_NOT_LINKED                                                        = 0x8009302A
	OSS_PER_DLL_NOT_LINKED                                                           = 0x8009302B
	OSS_OPEN_TYPE_ERROR                                                              = 0x8009302C
	OSS_MUTEX_NOT_CREATED                                                            = 0x8009302D
	OSS_CANT_CLOSE_TRACE_FILE                                                        = 0x8009302E
	CRYPT_E_ASN1_ERROR                                                               = 0x80093100
	CRYPT_E_ASN1_INTERNAL                                                            = 0x80093101
	CRYPT_E_ASN1_EOD                                                                 = 0x80093102
	CRYPT_E_ASN1_CORRUPT                                                             = 0x80093103
	CRYPT_E_ASN1_LARGE                                                               = 0x80093104
	CRYPT_E_ASN1_CONSTRAINT                                                          = 0x80093105
	CRYPT_E_ASN1_MEMORY                                                              = 0x80093106
	CRYPT_E_ASN1_OVERFLOW                                                            = 0x80093107
	CRYPT_E_ASN1_BADPDU                                                              = 0x80093108
	CRYPT_E_ASN1_BADARGS                                                             = 0x80093109
	CRYPT_E_ASN1_BADREAL                                                             = 0x8009310A
	CRYPT_E_ASN1_BADTAG                                                              = 0x8009310B
	CRYPT_E_ASN1_CHOICE                                                              = 0x8009310C
	CRYPT_E_ASN1_RULE                                                                = 0x8009310D
	CRYPT_E_ASN1_UTF8                                                                = 0x8009310E
	CRYPT_E_ASN1_PDU_TYPE                                                            = 0x80093133
	CRYPT_E_ASN1_NYI                                                                 = 0x80093134
	CRYPT_E_ASN1_EXTENDED                                                            = 0x80093201
	CRYPT_E_ASN1_NOEOD                                                               = 0x80093202
	CERTSRV_E_BAD_REQUESTSUBJECT                                                     = 0x80094001
	CERTSRV_E_NO_REQUEST                                                             = 0x80094002
	CERTSRV_E_BAD_REQUESTSTATUS                                                      = 0x80094003
	CERTSRV_E_PROPERTY_EMPTY                                                         = 0x80094004
	CERTSRV_E_INVALID_CA_CERTIFICATE                                                 = 0x80094005
	CERTSRV_E_SERVER_SUSPENDED                                                       = 0x80094006
	CERTSRV_E_ENCODING_LENGTH                                                        = 0x80094007
	CERTSRV_E_ROLECONFLICT                                                           = 0x80094008
	CERTSRV_E_RESTRICTEDOFFICER                                                      = 0x80094009
	CERTSRV_E_KEY_ARCHIVAL_NOT_CONFIGURED                                            = 0x8009400A
	CERTSRV_E_NO_VALID_KRA                                                           = 0x8009400B
	CERTSRV_E_BAD_REQUEST_KEY_ARCHIVAL                                               = 0x8009400C
	CERTSRV_E_NO_CAADMIN_DEFINED                                                     = 0x8009400D
	CERTSRV_E_BAD_RENEWAL_CERT_ATTRIBUTE                                             = 0x8009400E
	CERTSRV_E_NO_DB_SESSIONS                                                         = 0x8009400F
	CERTSRV_E_ALIGNMENT_FAULT                                                        = 0x80094010
	CERTSRV_E_ENROLL_DENIED                                                          = 0x80094011
	CERTSRV_E_TEMPLATE_DENIED                                                        = 0x80094012
	CERTSRV_E_DOWNLEVEL_DC_SSL_OR_UPGRADE                                            = 0x80094013
	CERTSRV_E_ADMIN_DENIED_REQUEST                                                   = 0x80094014
	CERTSRV_E_NO_POLICY_SERVER                                                       = 0x80094015
	CERTSRV_E_UNSUPPORTED_CERT_TYPE                                                  = 0x80094800
	CERTSRV_E_NO_CERT_TYPE                                                           = 0x80094801
	CERTSRV_E_TEMPLATE_CONFLICT                                                      = 0x80094802
	CERTSRV_E_SUBJECT_ALT_NAME_REQUIRED                                              = 0x80094803
	CERTSRV_E_ARCHIVED_KEY_REQUIRED                                                  = 0x80094804
	CERTSRV_E_SMIME_REQUIRED                                                         = 0x80094805
	CERTSRV_E_BAD_RENEWAL_SUBJECT                                                    = 0x80094806
	CERTSRV_E_BAD_TEMPLATE_VERSION                                                   = 0x80094807
	CERTSRV_E_TEMPLATE_POLICY_REQUIRED                                               = 0x80094808
	CERTSRV_E_SIGNATURE_POLICY_REQUIRED                                              = 0x80094809
	CERTSRV_E_SIGNATURE_COUNT                                                        = 0x8009480A
	CERTSRV_E_SIGNATURE_REJECTED                                                     = 0x8009480B
	CERTSRV_E_ISSUANCE_POLICY_REQUIRED                                               = 0x8009480C
	CERTSRV_E_SUBJECT_UPN_REQUIRED                                                   = 0x8009480D
	CERTSRV_E_SUBJECT_DIRECTORY_GUID_REQUIRED                                        = 0x8009480E
	CERTSRV_E_SUBJECT_DNS_REQUIRED                                                   = 0x8009480F
	CERTSRV_E_ARCHIVED_KEY_UNEXPECTED                                                = 0x80094810
	CERTSRV_E_KEY_LENGTH                                                             = 0x80094811
	CERTSRV_E_SUBJECT_EMAIL_REQUIRED                                                 = 0x80094812
	CERTSRV_E_UNKNOWN_CERT_TYPE                                                      = 0x80094813
	CERTSRV_E_CERT_TYPE_OVERLAP                                                      = 0x80094814
	CERTSRV_E_TOO_MANY_SIGNATURES                                                    = 0x80094815
	XENROLL_E_KEY_NOT_EXPORTABLE                                                     = 0x80095000
	XENROLL_E_CANNOT_ADD_ROOT_CERT                                                   = 0x80095001
	XENROLL_E_RESPONSE_KA_HASH_NOT_FOUND                                             = 0x80095002
	XENROLL_E_RESPONSE_UNEXPECTED_KA_HASH                                            = 0x80095003
	XENROLL_E_RESPONSE_KA_HASH_MISMATCH                                              = 0x80095004
	XENROLL_E_KEYSPEC_SMIME_MISMATCH                                                 = 0x80095005
	TRUST_E_SYSTEM_ERROR                                                             = 0x80096001
	TRUST_E_NO_SIGNER_CERT                                                           = 0x80096002
	TRUST_E_COUNTER_SIGNER                                                           = 0x80096003
	TRUST_E_CERT_SIGNATURE                                                           = 0x80096004
	TRUST_E_TIME_STAMP                                                               = 0x80096005
	TRUST_E_BAD_DIGEST                                                               = 0x80096010
	TRUST_E_BASIC_CONSTRAINTS                                                        = 0x80096019
	TRUST_E_FINANCIAL_CRITERIA                                                       = 0x8009601E
	MSSIPOTF_E_OUTOFMEMRANGE                                                         = 0x80097001
	MSSIPOTF_E_CANTGETOBJECT                                                         = 0x80097002
	MSSIPOTF_E_NOHEADTABLE                                                           = 0x80097003
	MSSIPOTF_E_BAD_MAGICNUMBER                                                       = 0x80097004
	MSSIPOTF_E_BAD_OFFSET_TABLE                                                      = 0x80097005
	MSSIPOTF_E_TABLE_TAGORDER                                                        = 0x80097006
	MSSIPOTF_E_TABLE_LONGWORD                                                        = 0x80097007
	MSSIPOTF_E_BAD_FIRST_TABLE_PLACEMENT                                             = 0x80097008
	MSSIPOTF_E_TABLES_OVERLAP                                                        = 0x80097009
	MSSIPOTF_E_TABLE_PADBYTES                                                        = 0x8009700A
	MSSIPOTF_E_FILETOOSMALL                                                          = 0x8009700B
	MSSIPOTF_E_TABLE_CHECKSUM                                                        = 0x8009700C
	MSSIPOTF_E_FILE_CHECKSUM                                                         = 0x8009700D
	MSSIPOTF_E_FAILED_POLICY                                                         = 0x80097010
	MSSIPOTF_E_FAILED_HINTS_CHECK                                                    = 0x80097011
	MSSIPOTF_E_NOT_OPENTYPE                                                          = 0x80097012
	MSSIPOTF_E_FILE                                                                  = 0x80097013
	MSSIPOTF_E_CRYPT                                                                 = 0x80097014
	MSSIPOTF_E_BADVERSION                                                            = 0x80097015
	MSSIPOTF_E_DSIG_STRUCTURE                                                        = 0x80097016
	MSSIPOTF_E_PCONST_CHECK                                                          = 0x80097017
	MSSIPOTF_E_STRUCTURE                                                             = 0x80097018
	ERROR_CRED_REQUIRES_CONFIRMATION                                                 = 0x80097019
	NTE_OP_OK                                                                        = 0
	TRUST_E_PROVIDER_UNKNOWN                                                         = 0x800B0001
	TRUST_E_ACTION_UNKNOWN                                                           = 0x800B0002
	TRUST_E_SUBJECT_FORM_UNKNOWN                                                     = 0x800B0003
	TRUST_E_SUBJECT_NOT_TRUSTED                                                      = 0x800B0004
	DIGSIG_E_ENCODE                                                                  = 0x800B0005
	DIGSIG_E_DECODE                                                                  = 0x800B0006
	DIGSIG_E_EXTENSIBILITY                                                           = 0x800B0007
	DIGSIG_E_CRYPTO                                                                  = 0x800B0008
	PERSIST_E_SIZEDEFINITE                                                           = 0x800B0009
	PERSIST_E_SIZEINDEFINITE                                                         = 0x800B000A
	PERSIST_E_NOTSELFSIZING                                                          = 0x800B000B
	TRUST_E_NOSIGNATURE                                                              = 0x800B0100
	CERT_E_EXPIRED                                                                   = 0x800B0101
	CERT_E_VALIDITYPERIODNESTING                                                     = 0x800B0102
	CERT_E_ROLE                                                                      = 0x800B0103
	CERT_E_PATHLENCONST                                                              = 0x800B0104
	CERT_E_CRITICAL                                                                  = 0x800B0105
	CERT_E_PURPOSE                                                                   = 0x800B0106
	CERT_E_ISSUERCHAINING                                                            = 0x800B0107
	CERT_E_MALFORMED                                                                 = 0x800B0108
	CERT_E_UNTRUSTEDROOT                                                             = 0x800B0109
	CERT_E_CHAINING                                                                  = 0x800B010A
	TRUST_E_FAIL                                                                     = 0x800B010B
	CERT_E_REVOKED                                                                   = 0x800B010C
	CERT_E_UNTRUSTEDTESTROOT                                                         = 0x800B010D
	CERT_E_REVOCATION_FAILURE                                                        = 0x800B010E
	CERT_E_CN_NO_MATCH                                                               = 0x800B010F
	CERT_E_WRONG_USAGE                                                               = 0x800B0110
	TRUST_E_EXPLICIT_DISTRUST                                                        = 0x800B0111
	CERT_E_UNTRUSTEDCA                                                               = 0x800B0112
	CERT_E_INVALID_POLICY                                                            = 0x800B0113
	CERT_E_INVALID_NAME                                                              = 0x800B0114
	SPAPI_E_EXPECTED_SECTION_NAME                                                    = 0x800F0000
	SPAPI_E_BAD_SECTION_NAME_LINE                                                    = 0x800F0001
	SPAPI_E_SECTION_NAME_TOO_LONG                                                    = 0x800F0002
	SPAPI_E_GENERAL_SYNTAX                                                           = 0x800F0003
	SPAPI_E_WRONG_INF_STYLE                                                          = 0x800F0100
	SPAPI_E_SECTION_NOT_FOUND                                                        = 0x800F0101
	SPAPI_E_LINE_NOT_FOUND                                                           = 0x800F0102
	SPAPI_E_NO_BACKUP                                                                = 0x800F0103
	SPAPI_E_NO_ASSOCIATED_CLASS                                                      = 0x800F0200
	SPAPI_E_CLASS_MISMATCH                                                           = 0x800F0201
	SPAPI_E_DUPLICATE_FOUND                                                          = 0x800F0202
	SPAPI_E_NO_DRIVER_SELECTED                                                       = 0x800F0203
	SPAPI_E_KEY_DOES_NOT_EXIST                                                       = 0x800F0204
	SPAPI_E_INVALID_DEVINST_NAME                                                     = 0x800F0205
	SPAPI_E_INVALID_CLASS                                                            = 0x800F0206
	SPAPI_E_DEVINST_ALREADY_EXISTS                                                   = 0x800F0207
	SPAPI_E_DEVINFO_NOT_REGISTERED                                                   = 0x800F0208
	SPAPI_E_INVALID_REG_PROPERTY                                                     = 0x800F0209
	SPAPI_E_NO_INF                                                                   = 0x800F020A
	SPAPI_E_NO_SUCH_DEVINST                                                          = 0x800F020B
	SPAPI_E_CANT_LOAD_CLASS_ICON                                                     = 0x800F020C
	SPAPI_E_INVALID_CLASS_INSTALLER                                                  = 0x800F020D
	SPAPI_E_DI_DO_DEFAULT                                                            = 0x800F020E
	SPAPI_E_DI_NOFILECOPY                                                            = 0x800F020F
	SPAPI_E_INVALID_HWPROFILE                                                        = 0x800F0210
	SPAPI_E_NO_DEVICE_SELECTED                                                       = 0x800F0211
	SPAPI_E_DEVINFO_LIST_LOCKED                                                      = 0x800F0212
	SPAPI_E_DEVINFO_DATA_LOCKED                                                      = 0x800F0213
	SPAPI_E_DI_BAD_PATH                                                              = 0x800F0214
	SPAPI_E_NO_CLASSINSTALL_PARAMS                                                   = 0x800F0215
	SPAPI_E_FILEQUEUE_LOCKED                                                         = 0x800F0216
	SPAPI_E_BAD_SERVICE_INSTALLSECT                                                  = 0x800F0217
	SPAPI_E_NO_CLASS_DRIVER_LIST                                                     = 0x800F0218
	SPAPI_E_NO_ASSOCIATED_SERVICE                                                    = 0x800F0219
	SPAPI_E_NO_DEFAULT_DEVICE_INTERFACE                                              = 0x800F021A
	SPAPI_E_DEVICE_INTERFACE_ACTIVE                                                  = 0x800F021B
	SPAPI_E_DEVICE_INTERFACE_REMOVED                                                 = 0x800F021C
	SPAPI_E_BAD_INTERFACE_INSTALLSECT                                                = 0x800F021D
	SPAPI_E_NO_SUCH_INTERFACE_CLASS                                                  = 0x800F021E
	SPAPI_E_INVALID_REFERENCE_STRING                                                 = 0x800F021F
	SPAPI_E_INVALID_MACHINENAME                                                      = 0x800F0220
	SPAPI_E_REMOTE_COMM_FAILURE                                                      = 0x800F0221
	SPAPI_E_MACHINE_UNAVAILABLE                                                      = 0x800F0222
	SPAPI_E_NO_CONFIGMGR_SERVICES                                                    = 0x800F0223
	SPAPI_E_INVALID_PROPPAGE_PROVIDER                                                = 0x800F0224
	SPAPI_E_NO_SUCH_DEVICE_INTERFACE                                                 = 0x800F0225
	SPAPI_E_DI_POSTPROCESSING_REQUIRED                                               = 0x800F0226
	SPAPI_E_INVALID_COINSTALLER                                                      = 0x800F0227
	SPAPI_E_NO_COMPAT_DRIVERS                                                        = 0x800F0228
	SPAPI_E_NO_DEVICE_ICON                                                           = 0x800F0229
	SPAPI_E_INVALID_INF_LOGCONFIG                                                    = 0x800F022A
	SPAPI_E_DI_DONT_INSTALL                                                          = 0x800F022B
	SPAPI_E_INVALID_FILTER_DRIVER                                                    = 0x800F022C
	SPAPI_E_NON_WINDOWS_NT_DRIVER                                                    = 0x800F022D
	SPAPI_E_NON_WINDOWS_DRIVER                                                       = 0x800F022E
	SPAPI_E_NO_CATALOG_FOR_OEM_INF                                                   = 0x800F022F
	SPAPI_E_DEVINSTALL_QUEUE_NONNATIVE                                               = 0x800F0230
	SPAPI_E_NOT_DISABLEABLE                                                          = 0x800F0231
	SPAPI_E_CANT_REMOVE_DEVINST                                                      = 0x800F0232
	SPAPI_E_INVALID_TARGET                                                           = 0x800F0233
	SPAPI_E_DRIVER_NONNATIVE                                                         = 0x800F0234
	SPAPI_E_IN_WOW64                                                                 = 0x800F0235
	SPAPI_E_SET_SYSTEM_RESTORE_POINT                                                 = 0x800F0236
	SPAPI_E_INCORRECTLY_COPIED_INF                                                   = 0x800F0237
	SPAPI_E_SCE_DISABLED                                                             = 0x800F0238
	SPAPI_E_UNKNOWN_EXCEPTION                                                        = 0x800F0239
	SPAPI_E_PNP_REGISTRY_ERROR                                                       = 0x800F023A
	SPAPI_E_REMOTE_REQUEST_UNSUPPORTED                                               = 0x800F023B
	SPAPI_E_NOT_AN_INSTALLED_OEM_INF                                                 = 0x800F023C
	SPAPI_E_INF_IN_USE_BY_DEVICES                                                    = 0x800F023D
	SPAPI_E_DI_FUNCTION_OBSOLETE                                                     = 0x800F023E
	SPAPI_E_NO_AUTHENTICODE_CATALOG                                                  = 0x800F023F
	SPAPI_E_AUTHENTICODE_DISALLOWED                                                  = 0x800F0240
	SPAPI_E_AUTHENTICODE_TRUSTED_PUBLISHER                                           = 0x800F0241
	SPAPI_E_AUTHENTICODE_TRUST_NOT_ESTABLISHED                                       = 0x800F0242
	SPAPI_E_AUTHENTICODE_PUBLISHER_NOT_TRUSTED                                       = 0x800F0243
	SPAPI_E_SIGNATURE_OSATTRIBUTE_MISMATCH                                           = 0x800F0244
	SPAPI_E_ONLY_VALIDATE_VIA_AUTHENTICODE                                           = 0x800F0245
	SPAPI_E_DEVICE_INSTALLER_NOT_READY                                               = 0x800F0246
	SPAPI_E_DRIVER_STORE_ADD_FAILED                                                  = 0x800F0247
	SPAPI_E_DEVICE_INSTALL_BLOCKED                                                   = 0x800F0248
	SPAPI_E_DRIVER_INSTALL_BLOCKED                                                   = 0x800F0249
	SPAPI_E_WRONG_INF_TYPE                                                           = 0x800F024A
	SPAPI_E_FILE_HASH_NOT_IN_CATALOG                                                 = 0x800F024B
	SPAPI_E_DRIVER_STORE_DELETE_FAILED                                               = 0x800F024C
	SPAPI_E_UNRECOVERABLE_STACK_OVERFLOW                                             = 0x800F0300
	SPAPI_E_ERROR_NOT_INSTALLED                                                      = 0x800F1000
	SCARD_F_INTERNAL_ERROR                                                           = 0x80100001
	SCARD_E_CANCELLED                                                                = 0x80100002
	SCARD_E_INVALID_HANDLE                                                           = 0x80100003
	SCARD_E_INVALID_PARAMETER                                                        = 0x80100004
	SCARD_E_INVALID_TARGET                                                           = 0x80100005
	SCARD_E_NO_MEMORY                                                                = 0x80100006
	SCARD_F_WAITED_TOO_LONG                                                          = 0x80100007
	SCARD_E_INSUFFICIENT_BUFFER                                                      = 0x80100008
	SCARD_E_UNKNOWN_READER                                                           = 0x80100009
	SCARD_E_TIMEOUT                                                                  = 0x8010000A
	SCARD_E_SHARING_VIOLATION                                                        = 0x8010000B
	SCARD_E_NO_SMARTCARD                                                             = 0x8010000C
	SCARD_E_UNKNOWN_CARD                                                             = 0x8010000D
	SCARD_E_CANT_DISPOSE                                                             = 0x8010000E
	SCARD_E_PROTO_MISMATCH                                                           = 0x8010000F
	SCARD_E_NOT_READY                                                                = 0x80100010
	SCARD_E_INVALID_VALUE                                                            = 0x80100011
	SCARD_E_SYSTEM_CANCELLED                                                         = 0x80100012
	SCARD_F_COMM_ERROR                                                               = 0x80100013
	SCARD_F_UNKNOWN_ERROR                                                            = 0x80100014
	SCARD_E_INVALID_ATR                                                              = 0x80100015
	SCARD_E_NOT_TRANSACTED                                                           = 0x80100016
	SCARD_E_READER_UNAVAILABLE                                                       = 0x80100017
	SCARD_P_SHUTDOWN                                                                 = 0x80100018
	SCARD_E_PCI_TOO_SMALL                                                            = 0x80100019
	SCARD_E_READER_UNSUPPORTED                                                       = 0x8010001A
	SCARD_E_DUPLICATE_READER                                                         = 0x8010001B
	SCARD_E_CARD_UNSUPPORTED                                                         = 0x8010001C
	SCARD_E_NO_SERVICE                                                               = 0x8010001D
	SCARD_E_SERVICE_STOPPED                                                          = 0x8010001E
	SCARD_E_UNEXPECTED                                                               = 0x8010001F
	SCARD_E_ICC_INSTALLATION                                                         = 0x80100020
	SCARD_E_ICC_CREATEORDER                                                          = 0x80100021
	SCARD_E_UNSUPPORTED_FEATURE                                                      = 0x80100022
	SCARD_E_DIR_NOT_FOUND                                                            = 0x80100023
	SCARD_E_FILE_NOT_FOUND                                                           = 0x80100024
	SCARD_E_NO_DIR                                                                   = 0x80100025
	SCARD_E_NO_FILE                                                                  = 0x80100026
	SCARD_E_NO_ACCESS                                                                = 0x80100027
	SCARD_E_WRITE_TOO_MANY                                                           = 0x80100028
	SCARD_E_BAD_SEEK                                                                 = 0x80100029
	SCARD_E_INVALID_CHV                                                              = 0x8010002A
	SCARD_E_UNKNOWN_RES_MNG                                                          = 0x8010002B
	SCARD_E_NO_SUCH_CERTIFICATE                                                      = 0x8010002C
	SCARD_E_CERTIFICATE_UNAVAILABLE                                                  = 0x8010002D
	SCARD_E_NO_READERS_AVAILABLE                                                     = 0x8010002E
	SCARD_E_COMM_DATA_LOST                                                           = 0x8010002F
	SCARD_E_NO_KEY_CONTAINER                                                         = 0x80100030
	SCARD_E_SERVER_TOO_BUSY                                                          = 0x80100031
	SCARD_E_PIN_CACHE_EXPIRED                                                        = 0x80100032
	SCARD_E_NO_PIN_CACHE                                                             = 0x80100033
	SCARD_E_READ_ONLY_CARD                                                           = 0x80100034
	SCARD_W_UNSUPPORTED_CARD                                                         = 0x80100065
	SCARD_W_UNRESPONSIVE_CARD                                                        = 0x80100066
	SCARD_W_UNPOWERED_CARD                                                           = 0x80100067
	SCARD_W_RESET_CARD                                                               = 0x80100068
	SCARD_W_REMOVED_CARD                                                             = 0x80100069
	SCARD_W_SECURITY_VIOLATION                                                       = 0x8010006A
	SCARD_W_WRONG_CHV                                                                = 0x8010006B
	SCARD_W_CHV_BLOCKED                                                              = 0x8010006C
	SCARD_W_EOF                                                                      = 0x8010006D
	SCARD_W_CANCELLED_BY_USER                                                        = 0x8010006E
	SCARD_W_CARD_NOT_AUTHENTICATED                                                   = 0x8010006F
	SCARD_W_CACHE_ITEM_NOT_FOUND                                                     = 0x80100070
	SCARD_W_CACHE_ITEM_STALE                                                         = 0x80100071
	SCARD_W_CACHE_ITEM_TOO_BIG                                                       = 0x80100072
	COMADMIN_E_OBJECTERRORS                                                          = 0x80110401
	COMADMIN_E_OBJECTINVALID                                                         = 0x80110402
	COMADMIN_E_KEYMISSING                                                            = 0x80110403
	COMADMIN_E_ALREADYINSTALLED                                                      = 0x80110404
	COMADMIN_E_APP_FILE_WRITEFAIL                                                    = 0x80110407
	COMADMIN_E_APP_FILE_READFAIL                                                     = 0x80110408
	COMADMIN_E_APP_FILE_VERSION                                                      = 0x80110409
	COMADMIN_E_BADPATH                                                               = 0x8011040A
	COMADMIN_E_APPLICATIONEXISTS                                                     = 0x8011040B
	COMADMIN_E_ROLEEXISTS                                                            = 0x8011040C
	COMADMIN_E_CANTCOPYFILE                                                          = 0x8011040D
	COMADMIN_E_NOUSER                                                                = 0x8011040F
	COMADMIN_E_INVALIDUSERIDS                                                        = 0x80110410
	COMADMIN_E_NOREGISTRYCLSID                                                       = 0x80110411
	COMADMIN_E_BADREGISTRYPROGID                                                     = 0x80110412
	COMADMIN_E_AUTHENTICATIONLEVEL                                                   = 0x80110413
	COMADMIN_E_USERPASSWDNOTVALID                                                    = 0x80110414
	COMADMIN_E_CLSIDORIIDMISMATCH                                                    = 0x80110418
	COMADMIN_E_REMOTEINTERFACE                                                       = 0x80110419
	COMADMIN_E_DLLREGISTERSERVER                                                     = 0x8011041A
	COMADMIN_E_NOSERVERSHARE                                                         = 0x8011041B
	COMADMIN_E_DLLLOADFAILED                                                         = 0x8011041D
	COMADMIN_E_BADREGISTRYLIBID                                                      = 0x8011041E
	COMADMIN_E_APPDIRNOTFOUND                                                        = 0x8011041F
	COMADMIN_E_REGISTRARFAILED                                                       = 0x80110423
	COMADMIN_E_COMPFILE_DOESNOTEXIST                                                 = 0x80110424
	COMADMIN_E_COMPFILE_LOADDLLFAIL                                                  = 0x80110425
	COMADMIN_E_COMPFILE_GETCLASSOBJ                                                  = 0x80110426
	COMADMIN_E_COMPFILE_CLASSNOTAVAIL                                                = 0x80110427
	COMADMIN_E_COMPFILE_BADTLB                                                       = 0x80110428
	COMADMIN_E_COMPFILE_NOTINSTALLABLE                                               = 0x80110429
	COMADMIN_E_NOTCHANGEABLE                                                         = 0x8011042A
	COMADMIN_E_NOTDELETEABLE                                                         = 0x8011042B
	COMADMIN_E_SESSION                                                               = 0x8011042C
	COMADMIN_E_COMP_MOVE_LOCKED                                                      = 0x8011042D
	COMADMIN_E_COMP_MOVE_BAD_DEST                                                    = 0x8011042E
	COMADMIN_E_REGISTERTLB                                                           = 0x80110430
	COMADMIN_E_SYSTEMAPP                                                             = 0x80110433
	COMADMIN_E_COMPFILE_NOREGISTRAR                                                  = 0x80110434
	COMADMIN_E_COREQCOMPINSTALLED                                                    = 0x80110435
	COMADMIN_E_SERVICENOTINSTALLED                                                   = 0x80110436
	COMADMIN_E_PROPERTYSAVEFAILED                                                    = 0x80110437
	COMADMIN_E_OBJECTEXISTS                                                          = 0x80110438
	COMADMIN_E_COMPONENTEXISTS                                                       = 0x80110439
	COMADMIN_E_REGFILE_CORRUPT                                                       = 0x8011043B
	COMADMIN_E_PROPERTY_OVERFLOW                                                     = 0x8011043C
	COMADMIN_E_NOTINREGISTRY                                                         = 0x8011043E
	COMADMIN_E_OBJECTNOTPOOLABLE                                                     = 0x8011043F
	COMADMIN_E_APPLID_MATCHES_CLSID                                                  = 0x80110446
	COMADMIN_E_ROLE_DOES_NOT_EXIST                                                   = 0x80110447
	COMADMIN_E_START_APP_NEEDS_COMPONENTS                                            = 0x80110448
	COMADMIN_E_REQUIRES_DIFFERENT_PLATFORM                                           = 0x80110449
	COMADMIN_E_CAN_NOT_EXPORT_APP_PROXY                                              = 0x8011044A
	COMADMIN_E_CAN_NOT_START_APP                                                     = 0x8011044B
	COMADMIN_E_CAN_NOT_EXPORT_SYS_APP                                                = 0x8011044C
	COMADMIN_E_CANT_SUBSCRIBE_TO_COMPONENT                                           = 0x8011044D
	COMADMIN_E_EVENTCLASS_CANT_BE_SUBSCRIBER                                         = 0x8011044E
	COMADMIN_E_LIB_APP_PROXY_INCOMPATIBLE                                            = 0x8011044F
	COMADMIN_E_BASE_PARTITION_ONLY                                                   = 0x80110450
	COMADMIN_E_START_APP_DISABLED                                                    = 0x80110451
	COMADMIN_E_CAT_DUPLICATE_PARTITION_NAME                                          = 0x80110457
	COMADMIN_E_CAT_INVALID_PARTITION_NAME                                            = 0x80110458
	COMADMIN_E_CAT_PARTITION_IN_USE                                                  = 0x80110459
	COMADMIN_E_FILE_PARTITION_DUPLICATE_FILES                                        = 0x8011045A
	COMADMIN_E_CAT_IMPORTED_COMPONENTS_NOT_ALLOWED                                   = 0x8011045B
	COMADMIN_E_AMBIGUOUS_APPLICATION_NAME                                            = 0x8011045C
	COMADMIN_E_AMBIGUOUS_PARTITION_NAME                                              = 0x8011045D
	COMADMIN_E_REGDB_NOTINITIALIZED                                                  = 0x80110472
	COMADMIN_E_REGDB_NOTOPEN                                                         = 0x80110473
	COMADMIN_E_REGDB_SYSTEMERR                                                       = 0x80110474
	COMADMIN_E_REGDB_ALREADYRUNNING                                                  = 0x80110475
	COMADMIN_E_MIG_VERSIONNOTSUPPORTED                                               = 0x80110480
	COMADMIN_E_MIG_SCHEMANOTFOUND                                                    = 0x80110481
	COMADMIN_E_CAT_BITNESSMISMATCH                                                   = 0x80110482
	COMADMIN_E_CAT_UNACCEPTABLEBITNESS                                               = 0x80110483
	COMADMIN_E_CAT_WRONGAPPBITNESS                                                   = 0x80110484
	COMADMIN_E_CAT_PAUSE_RESUME_NOT_SUPPORTED                                        = 0x80110485
	COMADMIN_E_CAT_SERVERFAULT                                                       = 0x80110486
	COMQC_E_APPLICATION_NOT_QUEUED                                                   = 0x80110600
	COMQC_E_NO_QUEUEABLE_INTERFACES                                                  = 0x80110601
	COMQC_E_QUEUING_SERVICE_NOT_AVAILABLE                                            = 0x80110602
	COMQC_E_NO_IPERSISTSTREAM                                                        = 0x80110603
	COMQC_E_BAD_MESSAGE                                                              = 0x80110604
	COMQC_E_UNAUTHENTICATED                                                          = 0x80110605
	COMQC_E_UNTRUSTED_ENQUEUER                                                       = 0x80110606
	MSDTC_E_DUPLICATE_RESOURCE                                                       = 0x80110701
	COMADMIN_E_OBJECT_PARENT_MISSING                                                 = 0x80110808
	COMADMIN_E_OBJECT_DOES_NOT_EXIST                                                 = 0x80110809
	COMADMIN_E_APP_NOT_RUNNING                                                       = 0x8011080A
	COMADMIN_E_INVALID_PARTITION                                                     = 0x8011080B
	COMADMIN_E_SVCAPP_NOT_POOLABLE_OR_RECYCLABLE                                     = 0x8011080D
	COMADMIN_E_USER_IN_SET                                                           = 0x8011080E
	COMADMIN_E_CANTRECYCLELIBRARYAPPS                                                = 0x8011080F
	COMADMIN_E_CANTRECYCLESERVICEAPPS                                                = 0x80110811
	COMADMIN_E_PROCESSALREADYRECYCLED                                                = 0x80110812
	COMADMIN_E_PAUSEDPROCESSMAYNOTBERECYCLED                                         = 0x80110813
	COMADMIN_E_CANTMAKEINPROCSERVICE                                                 = 0x80110814
	COMADMIN_E_PROGIDINUSEBYCLSID                                                    = 0x80110815
	COMADMIN_E_DEFAULT_PARTITION_NOT_IN_SET                                          = 0x80110816
	COMADMIN_E_RECYCLEDPROCESSMAYNOTBEPAUSED                                         = 0x80110817
	COMADMIN_E_PARTITION_ACCESSDENIED                                                = 0x80110818
	COMADMIN_E_PARTITION_MSI_ONLY                                                    = 0x80110819
	COMADMIN_E_LEGACYCOMPS_NOT_ALLOWED_IN_1_0_FORMAT                                 = 0x8011081A
	COMADMIN_E_LEGACYCOMPS_NOT_ALLOWED_IN_NONBASE_PARTITIONS                         = 0x8011081B
	COMADMIN_E_COMP_MOVE_SOURCE                                                      = 0x8011081C
	COMADMIN_E_COMP_MOVE_DEST                                                        = 0x8011081D
	COMADMIN_E_COMP_MOVE_PRIVATE                                                     = 0x8011081E
	COMADMIN_E_BASEPARTITION_REQUIRED_IN_SET                                         = 0x8011081F
	COMADMIN_E_CANNOT_ALIAS_EVENTCLASS                                               = 0x80110820
	COMADMIN_E_PRIVATE_ACCESSDENIED                                                  = 0x80110821
	COMADMIN_E_SAFERINVALID                                                          = 0x80110822
	COMADMIN_E_REGISTRY_ACCESSDENIED                                                 = 0x80110823
	COMADMIN_E_PARTITIONS_DISABLED                                                   = 0x80110824
	ERROR_FLT_IO_COMPLETE                                                            = 0x001F0001
	ERROR_FLT_NO_HANDLER_DEFINED                                                     = 0x801F0001
	ERROR_FLT_CONTEXT_ALREADY_DEFINED                                                = 0x801F0002
	ERROR_FLT_INVALID_ASYNCHRONOUS_REQUEST                                           = 0x801F0003
	ERROR_FLT_DISALLOW_FAST_IO                                                       = 0x801F0004
	ERROR_FLT_INVALID_NAME_REQUEST                                                   = 0x801F0005
	ERROR_FLT_NOT_SAFE_TO_POST_OPERATION                                             = 0x801F0006
	ERROR_FLT_NOT_INITIALIZED                                                        = 0x801F0007
	ERROR_FLT_FILTER_NOT_READY                                                       = 0x801F0008
	ERROR_FLT_POST_OPERATION_CLEANUP                                                 = 0x801F0009
	ERROR_FLT_INTERNAL_ERROR                                                         = 0x801F000A
	ERROR_FLT_DELETING_OBJECT                                                        = 0x801F000B
	ERROR_FLT_MUST_BE_NONPAGED_POOL                                                  = 0x801F000C
	ERROR_FLT_DUPLICATE_ENTRY                                                        = 0x801F000D
	ERROR_FLT_CBDQ_DISABLED                                                          = 0x801F000E
	ERROR_FLT_DO_NOT_ATTACH                                                          = 0x801F000F
	ERROR_FLT_DO_NOT_DETACH                                                          = 0x801F0010
	ERROR_FLT_INSTANCE_ALTITUDE_COLLISION                                            = 0x801F0011
	ERROR_FLT_INSTANCE_NAME_COLLISION                                                = 0x801F0012
	ERROR_FLT_FILTER_NOT_FOUND                                                       = 0x801F0013
	ERROR_FLT_VOLUME_NOT_FOUND                                                       = 0x801F0014
	ERROR_FLT_INSTANCE_NOT_FOUND                                                     = 0x801F0015
	ERROR_FLT_CONTEXT_ALLOCATION_NOT_FOUND                                           = 0x801F0016
	ERROR_FLT_INVALID_CONTEXT_REGISTRATION                                           = 0x801F0017
	ERROR_FLT_NAME_CACHE_MISS                                                        = 0x801F0018
	ERROR_FLT_NO_DEVICE_OBJECT                                                       = 0x801F0019
	ERROR_FLT_VOLUME_ALREADY_MOUNTED                                                 = 0x801F001A
	ERROR_FLT_ALREADY_ENLISTED                                                       = 0x801F001B
	ERROR_FLT_CONTEXT_ALREADY_LINKED                                                 = 0x801F001C
	ERROR_FLT_NO_WAITER_FOR_REPLY                                                    = 0x801F0020
	ERROR_HUNG_DISPLAY_DRIVER_THREAD                                                 = 0x80260001
	DWM_E_COMPOSITIONDISABLED                                                        = 0x80263001
	DWM_E_REMOTING_NOT_SUPPORTED                                                     = 0x80263002
	DWM_E_NO_REDIRECTION_SURFACE_AVAILABLE                                           = 0x80263003
	DWM_E_NOT_QUEUING_PRESENTS                                                       = 0x80263004
	DWM_E_ADAPTER_NOT_FOUND                                                          = 0x80263005
	DWM_S_GDI_REDIRECTION_SURFACE                                                    = 0x00263005
	ERROR_MONITOR_NO_DESCRIPTOR                                                      = 0x00261001
	ERROR_MONITOR_UNKNOWN_DESCRIPTOR_FORMAT                                          = 0x00261002
	ERROR_MONITOR_INVALID_DESCRIPTOR_CHECKSUM                                        = 0xC0261003
	ERROR_MONITOR_INVALID_STANDARD_TIMING_BLOCK                                      = 0xC0261004
	ERROR_MONITOR_WMI_DATABLOCK_REGISTRATION_FAILED                                  = 0xC0261005
	ERROR_MONITOR_INVALID_SERIAL_NUMBER_MONDSC_BLOCK                                 = 0xC0261006
	ERROR_MONITOR_INVALID_USER_FRIENDLY_MONDSC_BLOCK                                 = 0xC0261007
	ERROR_MONITOR_NO_MORE_DESCRIPTOR_DATA                                            = 0xC0261008
	ERROR_MONITOR_INVALID_DETAILED_TIMING_BLOCK                                      = 0xC0261009
	ERROR_MONITOR_INVALID_MANUFACTURE_DATE                                           = 0xC026100A
	ERROR_GRAPHICS_NOT_EXCLUSIVE_MODE_OWNER                                          = 0xC0262000
	ERROR_GRAPHICS_INSUFFICIENT_DMA_BUFFER                                           = 0xC0262001
	ERROR_GRAPHICS_INVALID_DISPLAY_ADAPTER                                           = 0xC0262002
	ERROR_GRAPHICS_ADAPTER_WAS_RESET                                                 = 0xC0262003
	ERROR_GRAPHICS_INVALID_DRIVER_MODEL                                              = 0xC0262004
	ERROR_GRAPHICS_PRESENT_MODE_CHANGED                                              = 0xC0262005
	ERROR_GRAPHICS_PRESENT_OCCLUDED                                                  = 0xC0262006
	ERROR_GRAPHICS_PRESENT_DENIED                                                    = 0xC0262007
	ERROR_GRAPHICS_CANNOTCOLORCONVERT                                                = 0xC0262008
	ERROR_GRAPHICS_DRIVER_MISMATCH                                                   = 0xC0262009
	ERROR_GRAPHICS_PARTIAL_DATA_POPULATED                                            = 0x4026200A
	ERROR_GRAPHICS_PRESENT_REDIRECTION_DISABLED                                      = 0xC026200B
	ERROR_GRAPHICS_PRESENT_UNOCCLUDED                                                = 0xC026200C
	ERROR_GRAPHICS_NO_VIDEO_MEMORY                                                   = 0xC0262100
	ERROR_GRAPHICS_CANT_LOCK_MEMORY                                                  = 0xC0262101
	ERROR_GRAPHICS_ALLOCATION_BUSY                                                   = 0xC0262102
	ERROR_GRAPHICS_TOO_MANY_REFERENCES                                               = 0xC0262103
	ERROR_GRAPHICS_TRY_AGAIN_LATER                                                   = 0xC0262104
	ERROR_GRAPHICS_TRY_AGAIN_NOW                                                     = 0xC0262105
	ERROR_GRAPHICS_ALLOCATION_INVALID                                                = 0xC0262106
	ERROR_GRAPHICS_UNSWIZZLING_APERTURE_UNAVAILABLE                                  = 0xC0262107
	ERROR_GRAPHICS_UNSWIZZLING_APERTURE_UNSUPPORTED                                  = 0xC0262108
	ERROR_GRAPHICS_CANT_EVICT_PINNED_ALLOCATION                                      = 0xC0262109
	ERROR_GRAPHICS_INVALID_ALLOCATION_USAGE                                          = 0xC0262110
	ERROR_GRAPHICS_CANT_RENDER_LOCKED_ALLOCATION                                     = 0xC0262111
	ERROR_GRAPHICS_ALLOCATION_CLOSED                                                 = 0xC0262112
	ERROR_GRAPHICS_INVALID_ALLOCATION_INSTANCE                                       = 0xC0262113
	ERROR_GRAPHICS_INVALID_ALLOCATION_HANDLE                                         = 0xC0262114
	ERROR_GRAPHICS_WRONG_ALLOCATION_DEVICE                                           = 0xC0262115
	ERROR_GRAPHICS_ALLOCATION_CONTENT_LOST                                           = 0xC0262116
	ERROR_GRAPHICS_GPU_EXCEPTION_ON_DEVICE                                           = 0xC0262200
	ERROR_GRAPHICS_INVALID_VIDPN_TOPOLOGY                                            = 0xC0262300
	ERROR_GRAPHICS_VIDPN_TOPOLOGY_NOT_SUPPORTED                                      = 0xC0262301
	ERROR_GRAPHICS_VIDPN_TOPOLOGY_CURRENTLY_NOT_SUPPORTED                            = 0xC0262302
	ERROR_GRAPHICS_INVALID_VIDPN                                                     = 0xC0262303
	ERROR_GRAPHICS_INVALID_VIDEO_PRESENT_SOURCE                                      = 0xC0262304
	ERROR_GRAPHICS_INVALID_VIDEO_PRESENT_TARGET                                      = 0xC0262305
	ERROR_GRAPHICS_VIDPN_MODALITY_NOT_SUPPORTED                                      = 0xC0262306
	ERROR_GRAPHICS_MODE_NOT_PINNED                                                   = 0x00262307
	ERROR_GRAPHICS_INVALID_VIDPN_SOURCEMODESET                                       = 0xC0262308
	ERROR_GRAPHICS_INVALID_VIDPN_TARGETMODESET                                       = 0xC0262309
	ERROR_GRAPHICS_INVALID_FREQUENCY                                                 = 0xC026230A
	ERROR_GRAPHICS_INVALID_ACTIVE_REGION                                             = 0xC026230B
	ERROR_GRAPHICS_INVALID_TOTAL_REGION                                              = 0xC026230C
	ERROR_GRAPHICS_INVALID_VIDEO_PRESENT_SOURCE_MODE                                 = 0xC0262310
	ERROR_GRAPHICS_INVALID_VIDEO_PRESENT_TARGET_MODE                                 = 0xC0262311
	ERROR_GRAPHICS_PINNED_MODE_MUST_REMAIN_IN_SET                                    = 0xC0262312
	ERROR_GRAPHICS_PATH_ALREADY_IN_TOPOLOGY                                          = 0xC0262313
	ERROR_GRAPHICS_MODE_ALREADY_IN_MODESET                                           = 0xC0262314
	ERROR_GRAPHICS_INVALID_VIDEOPRESENTSOURCESET                                     = 0xC0262315
	ERROR_GRAPHICS_INVALID_VIDEOPRESENTTARGETSET                                     = 0xC0262316
	ERROR_GRAPHICS_SOURCE_ALREADY_IN_SET                                             = 0xC0262317
	ERROR_GRAPHICS_TARGET_ALREADY_IN_SET                                             = 0xC0262318
	ERROR_GRAPHICS_INVALID_VIDPN_PRESENT_PATH                                        = 0xC0262319
	ERROR_GRAPHICS_NO_RECOMMENDED_VIDPN_TOPOLOGY                                     = 0xC026231A
	ERROR_GRAPHICS_INVALID_MONITOR_FREQUENCYRANGESET                                 = 0xC026231B
	ERROR_GRAPHICS_INVALID_MONITOR_FREQUENCYRANGE                                    = 0xC026231C
	ERROR_GRAPHICS_FREQUENCYRANGE_NOT_IN_SET                                         = 0xC026231D
	ERROR_GRAPHICS_NO_PREFERRED_MODE                                                 = 0x0026231E
	ERROR_GRAPHICS_FREQUENCYRANGE_ALREADY_IN_SET                                     = 0xC026231F
	ERROR_GRAPHICS_STALE_MODESET                                                     = 0xC0262320
	ERROR_GRAPHICS_INVALID_MONITOR_SOURCEMODESET                                     = 0xC0262321
	ERROR_GRAPHICS_INVALID_MONITOR_SOURCE_MODE                                       = 0xC0262322
	ERROR_GRAPHICS_NO_RECOMMENDED_FUNCTIONAL_VIDPN                                   = 0xC0262323
	ERROR_GRAPHICS_MODE_ID_MUST_BE_UNIQUE                                            = 0xC0262324
	ERROR_GRAPHICS_EMPTY_ADAPTER_MONITOR_MODE_SUPPORT_INTERSECTION                   = 0xC0262325
	ERROR_GRAPHICS_VIDEO_PRESENT_TARGETS_LESS_THAN_SOURCES                           = 0xC0262326
	ERROR_GRAPHICS_PATH_NOT_IN_TOPOLOGY                                              = 0xC0262327
	ERROR_GRAPHICS_ADAPTER_MUST_HAVE_AT_LEAST_ONE_SOURCE                             = 0xC0262328
	ERROR_GRAPHICS_ADAPTER_MUST_HAVE_AT_LEAST_ONE_TARGET                             = 0xC0262329
	ERROR_GRAPHICS_INVALID_MONITORDESCRIPTORSET                                      = 0xC026232A
	ERROR_GRAPHICS_INVALID_MONITORDESCRIPTOR                                         = 0xC026232B
	ERROR_GRAPHICS_MONITORDESCRIPTOR_NOT_IN_SET                                      = 0xC026232C
	ERROR_GRAPHICS_MONITORDESCRIPTOR_ALREADY_IN_SET                                  = 0xC026232D
	ERROR_GRAPHICS_MONITORDESCRIPTOR_ID_MUST_BE_UNIQUE                               = 0xC026232E
	ERROR_GRAPHICS_INVALID_VIDPN_TARGET_SUBSET_TYPE                                  = 0xC026232F
	ERROR_GRAPHICS_RESOURCES_NOT_RELATED                                             = 0xC0262330
	ERROR_GRAPHICS_SOURCE_ID_MUST_BE_UNIQUE                                          = 0xC0262331
	ERROR_GRAPHICS_TARGET_ID_MUST_BE_UNIQUE                                          = 0xC0262332
	ERROR_GRAPHICS_NO_AVAILABLE_VIDPN_TARGET                                         = 0xC0262333
	ERROR_GRAPHICS_MONITOR_COULD_NOT_BE_ASSOCIATED_WITH_ADAPTER                      = 0xC0262334
	ERROR_GRAPHICS_NO_VIDPNMGR                                                       = 0xC0262335
	ERROR_GRAPHICS_NO_ACTIVE_VIDPN                                                   = 0xC0262336
	ERROR_GRAPHICS_STALE_VIDPN_TOPOLOGY                                              = 0xC0262337
	ERROR_GRAPHICS_MONITOR_NOT_CONNECTED                                             = 0xC0262338
	ERROR_GRAPHICS_SOURCE_NOT_IN_TOPOLOGY                                            = 0xC0262339
	ERROR_GRAPHICS_INVALID_PRIMARYSURFACE_SIZE                                       = 0xC026233A
	ERROR_GRAPHICS_INVALID_VISIBLEREGION_SIZE                                        = 0xC026233B
	ERROR_GRAPHICS_INVALID_STRIDE                                                    = 0xC026233C
	ERROR_GRAPHICS_INVALID_PIXELFORMAT                                               = 0xC026233D
	ERROR_GRAPHICS_INVALID_COLORBASIS                                                = 0xC026233E
	ERROR_GRAPHICS_INVALID_PIXELVALUEACCESSMODE                                      = 0xC026233F
	ERROR_GRAPHICS_TARGET_NOT_IN_TOPOLOGY                                            = 0xC0262340
	ERROR_GRAPHICS_NO_DISPLAY_MODE_MANAGEMENT_SUPPORT                                = 0xC0262341
	ERROR_GRAPHICS_VIDPN_SOURCE_IN_USE                                               = 0xC0262342
	ERROR_GRAPHICS_CANT_ACCESS_ACTIVE_VIDPN                                          = 0xC0262343
	ERROR_GRAPHICS_INVALID_PATH_IMPORTANCE_ORDINAL                                   = 0xC0262344
	ERROR_GRAPHICS_INVALID_PATH_CONTENT_GEOMETRY_TRANSFORMATION                      = 0xC0262345
	ERROR_GRAPHICS_PATH_CONTENT_GEOMETRY_TRANSFORMATION_NOT_SUPPORTED                = 0xC0262346
	ERROR_GRAPHICS_INVALID_GAMMA_RAMP                                                = 0xC0262347
	ERROR_GRAPHICS_GAMMA_RAMP_NOT_SUPPORTED                                          = 0xC0262348
	ERROR_GRAPHICS_MULTISAMPLING_NOT_SUPPORTED                                       = 0xC0262349
	ERROR_GRAPHICS_MODE_NOT_IN_MODESET                                               = 0xC026234A
	ERROR_GRAPHICS_DATASET_IS_EMPTY                                                  = 0x0026234B
	ERROR_GRAPHICS_NO_MORE_ELEMENTS_IN_DATASET                                       = 0x0026234C
	ERROR_GRAPHICS_INVALID_VIDPN_TOPOLOGY_RECOMMENDATION_REASON                      = 0xC026234D
	ERROR_GRAPHICS_INVALID_PATH_CONTENT_TYPE                                         = 0xC026234E
	ERROR_GRAPHICS_INVALID_COPYPROTECTION_TYPE                                       = 0xC026234F
	ERROR_GRAPHICS_UNASSIGNED_MODESET_ALREADY_EXISTS                                 = 0xC0262350
	ERROR_GRAPHICS_PATH_CONTENT_GEOMETRY_TRANSFORMATION_NOT_PINNED                   = 0x00262351
	ERROR_GRAPHICS_INVALID_SCANLINE_ORDERING                                         = 0xC0262352
	ERROR_GRAPHICS_TOPOLOGY_CHANGES_NOT_ALLOWED                                      = 0xC0262353
	ERROR_GRAPHICS_NO_AVAILABLE_IMPORTANCE_ORDINALS                                  = 0xC0262354
	ERROR_GRAPHICS_INCOMPATIBLE_PRIVATE_FORMAT                                       = 0xC0262355
	ERROR_GRAPHICS_INVALID_MODE_PRUNING_ALGORITHM                                    = 0xC0262356
	ERROR_GRAPHICS_INVALID_MONITOR_CAPABILITY_ORIGIN                                 = 0xC0262357
	ERROR_GRAPHICS_INVALID_MONITOR_FREQUENCYRANGE_CONSTRAINT                         = 0xC0262358
	ERROR_GRAPHICS_MAX_NUM_PATHS_REACHED                                             = 0xC0262359
	ERROR_GRAPHICS_CANCEL_VIDPN_TOPOLOGY_AUGMENTATION                                = 0xC026235A
	ERROR_GRAPHICS_INVALID_CLIENT_TYPE                                               = 0xC026235B
	ERROR_GRAPHICS_CLIENTVIDPN_NOT_SET                                               = 0xC026235C
	ERROR_GRAPHICS_SPECIFIED_CHILD_ALREADY_CONNECTED                                 = 0xC0262400
	ERROR_GRAPHICS_CHILD_DESCRIPTOR_NOT_SUPPORTED                                    = 0xC0262401
	ERROR_GRAPHICS_UNKNOWN_CHILD_STATUS                                              = 0x4026242F
	ERROR_GRAPHICS_NOT_A_LINKED_ADAPTER                                              = 0xC0262430
	ERROR_GRAPHICS_LEADLINK_NOT_ENUMERATED                                           = 0xC0262431
	ERROR_GRAPHICS_CHAINLINKS_NOT_ENUMERATED                                         = 0xC0262432
	ERROR_GRAPHICS_ADAPTER_CHAIN_NOT_READY                                           = 0xC0262433
	ERROR_GRAPHICS_CHAINLINKS_NOT_STARTED                                            = 0xC0262434
	ERROR_GRAPHICS_CHAINLINKS_NOT_POWERED_ON                                         = 0xC0262435
	ERROR_GRAPHICS_INCONSISTENT_DEVICE_LINK_STATE                                    = 0xC0262436
	ERROR_GRAPHICS_LEADLINK_START_DEFERRED                                           = 0x40262437
	ERROR_GRAPHICS_NOT_POST_DEVICE_DRIVER                                            = 0xC0262438
	ERROR_GRAPHICS_POLLING_TOO_FREQUENTLY                                            = 0x40262439
	ERROR_GRAPHICS_START_DEFERRED                                                    = 0x4026243A
	ERROR_GRAPHICS_ADAPTER_ACCESS_NOT_EXCLUDED                                       = 0xC026243B
	ERROR_GRAPHICS_OPM_NOT_SUPPORTED                                                 = 0xC0262500
	ERROR_GRAPHICS_COPP_NOT_SUPPORTED                                                = 0xC0262501
	ERROR_GRAPHICS_UAB_NOT_SUPPORTED                                                 = 0xC0262502
	ERROR_GRAPHICS_OPM_INVALID_ENCRYPTED_PARAMETERS                                  = 0xC0262503
	ERROR_GRAPHICS_OPM_NO_VIDEO_OUTPUTS_EXIST                                        = 0xC0262505
	ERROR_GRAPHICS_OPM_INTERNAL_ERROR                                                = 0xC026250B
	ERROR_GRAPHICS_OPM_INVALID_HANDLE                                                = 0xC026250C
	ERROR_GRAPHICS_PVP_INVALID_CERTIFICATE_LENGTH                                    = 0xC026250E
	ERROR_GRAPHICS_OPM_SPANNING_MODE_ENABLED                                         = 0xC026250F
	ERROR_GRAPHICS_OPM_THEATER_MODE_ENABLED                                          = 0xC0262510
	ERROR_GRAPHICS_PVP_HFS_FAILED                                                    = 0xC0262511
	ERROR_GRAPHICS_OPM_INVALID_SRM                                                   = 0xC0262512
	ERROR_GRAPHICS_OPM_OUTPUT_DOES_NOT_SUPPORT_HDCP                                  = 0xC0262513
	ERROR_GRAPHICS_OPM_OUTPUT_DOES_NOT_SUPPORT_ACP                                   = 0xC0262514
	ERROR_GRAPHICS_OPM_OUTPUT_DOES_NOT_SUPPORT_CGMSA                                 = 0xC0262515
	ERROR_GRAPHICS_OPM_HDCP_SRM_NEVER_SET                                            = 0xC0262516
	ERROR_GRAPHICS_OPM_RESOLUTION_TOO_HIGH                                           = 0xC0262517
	ERROR_GRAPHICS_OPM_ALL_HDCP_HARDWARE_ALREADY_IN_USE                              = 0xC0262518
	ERROR_GRAPHICS_OPM_VIDEO_OUTPUT_NO_LONGER_EXISTS                                 = 0xC026251A
	ERROR_GRAPHICS_OPM_SESSION_TYPE_CHANGE_IN_PROGRESS                               = 0xC026251B
	ERROR_GRAPHICS_OPM_VIDEO_OUTPUT_DOES_NOT_HAVE_COPP_SEMANTICS                     = 0xC026251C
	ERROR_GRAPHICS_OPM_INVALID_INFORMATION_REQUEST                                   = 0xC026251D
	ERROR_GRAPHICS_OPM_DRIVER_INTERNAL_ERROR                                         = 0xC026251E
	ERROR_GRAPHICS_OPM_VIDEO_OUTPUT_DOES_NOT_HAVE_OPM_SEMANTICS                      = 0xC026251F
	ERROR_GRAPHICS_OPM_SIGNALING_NOT_SUPPORTED                                       = 0xC0262520
	ERROR_GRAPHICS_OPM_INVALID_CONFIGURATION_REQUEST                                 = 0xC0262521
	ERROR_GRAPHICS_I2C_NOT_SUPPORTED                                                 = 0xC0262580
	ERROR_GRAPHICS_I2C_DEVICE_DOES_NOT_EXIST                                         = 0xC0262581
	ERROR_GRAPHICS_I2C_ERROR_TRANSMITTING_DATA                                       = 0xC0262582
	ERROR_GRAPHICS_I2C_ERROR_RECEIVING_DATA                                          = 0xC0262583
	ERROR_GRAPHICS_DDCCI_VCP_NOT_SUPPORTED                                           = 0xC0262584
	ERROR_GRAPHICS_DDCCI_INVALID_DATA                                                = 0xC0262585
	ERROR_GRAPHICS_DDCCI_MONITOR_RETURNED_INVALID_TIMING_STATUS_BYTE                 = 0xC0262586
	ERROR_GRAPHICS_MCA_INVALID_CAPABILITIES_STRING                                   = 0xC0262587
	ERROR_GRAPHICS_MCA_INTERNAL_ERROR                                                = 0xC0262588
	ERROR_GRAPHICS_DDCCI_INVALID_MESSAGE_COMMAND                                     = 0xC0262589
	ERROR_GRAPHICS_DDCCI_INVALID_MESSAGE_LENGTH                                      = 0xC026258A
	ERROR_GRAPHICS_DDCCI_INVALID_MESSAGE_CHECKSUM                                    = 0xC026258B
	ERROR_GRAPHICS_INVALID_PHYSICAL_MONITOR_HANDLE                                   = 0xC026258C
	ERROR_GRAPHICS_MONITOR_NO_LONGER_EXISTS                                          = 0xC026258D
	ERROR_GRAPHICS_DDCCI_CURRENT_CURRENT_VALUE_GREATER_THAN_MAXIMUM_VALUE            = 0xC02625D8
	ERROR_GRAPHICS_MCA_INVALID_VCP_VERSION                                           = 0xC02625D9
	ERROR_GRAPHICS_MCA_MONITOR_VIOLATES_MCCS_SPECIFICATION                           = 0xC02625DA
	ERROR_GRAPHICS_MCA_MCCS_VERSION_MISMATCH                                         = 0xC02625DB
	ERROR_GRAPHICS_MCA_UNSUPPORTED_MCCS_VERSION                                      = 0xC02625DC
	ERROR_GRAPHICS_MCA_INVALID_TECHNOLOGY_TYPE_RETURNED                              = 0xC02625DE
	ERROR_GRAPHICS_MCA_UNSUPPORTED_COLOR_TEMPERATURE                                 = 0xC02625DF
	ERROR_GRAPHICS_ONLY_CONSOLE_SESSION_SUPPORTED                                    = 0xC02625E0
	ERROR_GRAPHICS_NO_DISPLAY_DEVICE_CORRESPONDS_TO_NAME                             = 0xC02625E1
	ERROR_GRAPHICS_DISPLAY_DEVICE_NOT_ATTACHED_TO_DESKTOP                            = 0xC02625E2
	ERROR_GRAPHICS_MIRRORING_DEVICES_NOT_SUPPORTED                                   = 0xC02625E3
	ERROR_GRAPHICS_INVALID_POINTER                                                   = 0xC02625E4
	ERROR_GRAPHICS_NO_MONITORS_CORRESPOND_TO_DISPLAY_DEVICE                          = 0xC02625E5
	ERROR_GRAPHICS_PARAMETER_ARRAY_TOO_SMALL                                         = 0xC02625E6
	ERROR_GRAPHICS_INTERNAL_ERROR                                                    = 0xC02625E7
	ERROR_GRAPHICS_SESSION_TYPE_CHANGE_IN_PROGRESS                                   = 0xC02605E8
	TPM_E_ERROR_MASK                                                                 = 0x80280000
	TPM_E_AUTHFAIL                                                                   = 0x80280001
	TPM_E_BADINDEX                                                                   = 0x80280002
	TPM_E_BAD_PARAMETER                                                              = 0x80280003
	TPM_E_AUDITFAILURE                                                               = 0x80280004
	TPM_E_CLEAR_DISABLED                                                             = 0x80280005
	TPM_E_DEACTIVATED                                                                = 0x80280006
	TPM_E_DISABLED                                                                   = 0x80280007
	TPM_E_DISABLED_CMD                                                               = 0x80280008
	TPM_E_FAIL                                                                       = 0x80280009
	TPM_E_BAD_ORDINAL                                                                = 0x8028000A
	TPM_E_INSTALL_DISABLED                                                           = 0x8028000B
	TPM_E_INVALID_KEYHANDLE                                                          = 0x8028000C
	TPM_E_KEYNOTFOUND                                                                = 0x8028000D
	TPM_E_INAPPROPRIATE_ENC                                                          = 0x8028000E
	TPM_E_MIGRATEFAIL                                                                = 0x8028000F
	TPM_E_INVALID_PCR_INFO                                                           = 0x80280010
	TPM_E_NOSPACE                                                                    = 0x80280011
	TPM_E_NOSRK                                                                      = 0x80280012
	TPM_E_NOTSEALED_BLOB                                                             = 0x80280013
	TPM_E_OWNER_SET                                                                  = 0x80280014
	TPM_E_RESOURCES                                                                  = 0x80280015
	TPM_E_SHORTRANDOM                                                                = 0x80280016
	TPM_E_SIZE                                                                       = 0x80280017
	TPM_E_WRONGPCRVAL                                                                = 0x80280018
	TPM_E_BAD_PARAM_SIZE                                                             = 0x80280019
	TPM_E_SHA_THREAD                                                                 = 0x8028001A
	TPM_E_SHA_ERROR                                                                  = 0x8028001B
	TPM_E_FAILEDSELFTEST                                                             = 0x8028001C
	TPM_E_AUTH2FAIL                                                                  = 0x8028001D
	TPM_E_BADTAG                                                                     = 0x8028001E
	TPM_E_IOERROR                                                                    = 0x8028001F
	TPM_E_ENCRYPT_ERROR                                                              = 0x80280020
	TPM_E_DECRYPT_ERROR                                                              = 0x80280021
	TPM_E_INVALID_AUTHHANDLE                                                         = 0x80280022
	TPM_E_NO_ENDORSEMENT                                                             = 0x80280023
	TPM_E_INVALID_KEYUSAGE                                                           = 0x80280024
	TPM_E_WRONG_ENTITYTYPE                                                           = 0x80280025
	TPM_E_INVALID_POSTINIT                                                           = 0x80280026
	TPM_E_INAPPROPRIATE_SIG                                                          = 0x80280027
	TPM_E_BAD_KEY_PROPERTY                                                           = 0x80280028
	TPM_E_BAD_MIGRATION                                                              = 0x80280029
	TPM_E_BAD_SCHEME                                                                 = 0x8028002A
	TPM_E_BAD_DATASIZE                                                               = 0x8028002B
	TPM_E_BAD_MODE                                                                   = 0x8028002C
	TPM_E_BAD_PRESENCE                                                               = 0x8028002D
	TPM_E_BAD_VERSION                                                                = 0x8028002E
	TPM_E_NO_WRAP_TRANSPORT                                                          = 0x8028002F
	TPM_E_AUDITFAIL_UNSUCCESSFUL                                                     = 0x80280030
	TPM_E_AUDITFAIL_SUCCESSFUL                                                       = 0x80280031
	TPM_E_NOTRESETABLE                                                               = 0x80280032
	TPM_E_NOTLOCAL                                                                   = 0x80280033
	TPM_E_BAD_TYPE                                                                   = 0x80280034
	TPM_E_INVALID_RESOURCE                                                           = 0x80280035
	TPM_E_NOTFIPS                                                                    = 0x80280036
	TPM_E_INVALID_FAMILY                                                             = 0x80280037
	TPM_E_NO_NV_PERMISSION                                                           = 0x80280038
	TPM_E_REQUIRES_SIGN                                                              = 0x80280039
	TPM_E_KEY_NOTSUPPORTED                                                           = 0x8028003A
	TPM_E_AUTH_CONFLICT                                                              = 0x8028003B
	TPM_E_AREA_LOCKED                                                                = 0x8028003C
	TPM_E_BAD_LOCALITY                                                               = 0x8028003D
	TPM_E_READ_ONLY                                                                  = 0x8028003E
	TPM_E_PER_NOWRITE                                                                = 0x8028003F
	TPM_E_FAMILYCOUNT                                                                = 0x80280040
	TPM_E_WRITE_LOCKED                                                               = 0x80280041
	TPM_E_BAD_ATTRIBUTES                                                             = 0x80280042
	TPM_E_INVALID_STRUCTURE                                                          = 0x80280043
	TPM_E_KEY_OWNER_CONTROL                                                          = 0x80280044
	TPM_E_BAD_COUNTER                                                                = 0x80280045
	TPM_E_NOT_FULLWRITE                                                              = 0x80280046
	TPM_E_CONTEXT_GAP                                                                = 0x80280047
	TPM_E_MAXNVWRITES                                                                = 0x80280048
	TPM_E_NOOPERATOR                                                                 = 0x80280049
	TPM_E_RESOURCEMISSING                                                            = 0x8028004A
	TPM_E_DELEGATE_LOCK                                                              = 0x8028004B
	TPM_E_DELEGATE_FAMILY                                                            = 0x8028004C
	TPM_E_DELEGATE_ADMIN                                                             = 0x8028004D
	TPM_E_TRANSPORT_NOTEXCLUSIVE                                                     = 0x8028004E
	TPM_E_OWNER_CONTROL                                                              = 0x8028004F
	TPM_E_DAA_RESOURCES                                                              = 0x80280050
	TPM_E_DAA_INPUT_DATA0                                                            = 0x80280051
	TPM_E_DAA_INPUT_DATA1                                                            = 0x80280052
	TPM_E_DAA_ISSUER_SETTINGS                                                        = 0x80280053
	TPM_E_DAA_TPM_SETTINGS                                                           = 0x80280054
	TPM_E_DAA_STAGE                                                                  = 0x80280055
	TPM_E_DAA_ISSUER_VALIDITY                                                        = 0x80280056
	TPM_E_DAA_WRONG_W                                                                = 0x80280057
	TPM_E_BAD_HANDLE                                                                 = 0x80280058
	TPM_E_BAD_DELEGATE                                                               = 0x80280059
	TPM_E_BADCONTEXT                                                                 = 0x8028005A
	TPM_E_TOOMANYCONTEXTS                                                            = 0x8028005B
	TPM_E_MA_TICKET_SIGNATURE                                                        = 0x8028005C
	TPM_E_MA_DESTINATION                                                             = 0x8028005D
	TPM_E_MA_SOURCE                                                                  = 0x8028005E
	TPM_E_MA_AUTHORITY                                                               = 0x8028005F
	TPM_E_PERMANENTEK                                                                = 0x80280061
	TPM_E_BAD_SIGNATURE                                                              = 0x80280062
	TPM_E_NOCONTEXTSPACE                                                             = 0x80280063
	TPM_E_COMMAND_BLOCKED                                                            = 0x80280400
	TPM_E_INVALID_HANDLE                                                             = 0x80280401
	TPM_E_DUPLICATE_VHANDLE                                                          = 0x80280402
	TPM_E_EMBEDDED_COMMAND_BLOCKED                                                   = 0x80280403
	TPM_E_EMBEDDED_COMMAND_UNSUPPORTED                                               = 0x80280404
	TPM_E_RETRY                                                                      = 0x80280800
	TPM_E_NEEDS_SELFTEST                                                             = 0x80280801
	TPM_E_DOING_SELFTEST                                                             = 0x80280802
	TPM_E_DEFEND_LOCK_RUNNING                                                        = 0x80280803
	TBS_E_INTERNAL_ERROR                                                             = 0x80284001
	TBS_E_BAD_PARAMETER                                                              = 0x80284002
	TBS_E_INVALID_OUTPUT_POINTER                                                     = 0x80284003
	TBS_E_INVALID_CONTEXT                                                            = 0x80284004
	TBS_E_INSUFFICIENT_BUFFER                                                        = 0x80284005
	TBS_E_IOERROR                                                                    = 0x80284006
	TBS_E_INVALID_CONTEXT_PARAM                                                      = 0x80284007
	TBS_E_SERVICE_NOT_RUNNING                                                        = 0x80284008
	TBS_E_TOO_MANY_TBS_CONTEXTS                                                      = 0x80284009
	TBS_E_TOO_MANY_RESOURCES                                                         = 0x8028400A
	TBS_E_SERVICE_START_PENDING                                                      = 0x8028400B
	TBS_E_PPI_NOT_SUPPORTED                                                          = 0x8028400C
	TBS_E_COMMAND_CANCELED                                                           = 0x8028400D
	TBS_E_BUFFER_TOO_LARGE                                                           = 0x8028400E
	TBS_E_TPM_NOT_FOUND                                                              = 0x8028400F
	TBS_E_SERVICE_DISABLED                                                           = 0x80284010
	TBS_E_NO_EVENT_LOG                                                               = 0x80284011
	TPMAPI_E_INVALID_STATE                                                           = 0x80290100
	TPMAPI_E_NOT_ENOUGH_DATA                                                         = 0x80290101
	TPMAPI_E_TOO_MUCH_DATA                                                           = 0x80290102
	TPMAPI_E_INVALID_OUTPUT_POINTER                                                  = 0x80290103
	TPMAPI_E_INVALID_PARAMETER                                                       = 0x80290104
	TPMAPI_E_OUT_OF_MEMORY                                                           = 0x80290105
	TPMAPI_E_BUFFER_TOO_SMALL                                                        = 0x80290106
	TPMAPI_E_INTERNAL_ERROR                                                          = 0x80290107
	TPMAPI_E_ACCESS_DENIED                                                           = 0x80290108
	TPMAPI_E_AUTHORIZATION_FAILED                                                    = 0x80290109
	TPMAPI_E_INVALID_CONTEXT_HANDLE                                                  = 0x8029010A
	TPMAPI_E_TBS_COMMUNICATION_ERROR                                                 = 0x8029010B
	TPMAPI_E_TPM_COMMAND_ERROR                                                       = 0x8029010C
	TPMAPI_E_MESSAGE_TOO_LARGE                                                       = 0x8029010D
	TPMAPI_E_INVALID_ENCODING                                                        = 0x8029010E
	TPMAPI_E_INVALID_KEY_SIZE                                                        = 0x8029010F
	TPMAPI_E_ENCRYPTION_FAILED                                                       = 0x80290110
	TPMAPI_E_INVALID_KEY_PARAMS                                                      = 0x80290111
	TPMAPI_E_INVALID_MIGRATION_AUTHORIZATION_BLOB                                    = 0x80290112
	TPMAPI_E_INVALID_PCR_INDEX                                                       = 0x80290113
	TPMAPI_E_INVALID_DELEGATE_BLOB                                                   = 0x80290114
	TPMAPI_E_INVALID_CONTEXT_PARAMS                                                  = 0x80290115
	TPMAPI_E_INVALID_KEY_BLOB                                                        = 0x80290116
	TPMAPI_E_INVALID_PCR_DATA                                                        = 0x80290117
	TPMAPI_E_INVALID_OWNER_AUTH                                                      = 0x80290118
	TPMAPI_E_FIPS_RNG_CHECK_FAILED                                                   = 0x80290119
	TPMAPI_E_EMPTY_TCG_LOG                                                           = 0x8029011A
	TPMAPI_E_INVALID_TCG_LOG_ENTRY                                                   = 0x8029011B
	TPMAPI_E_TCG_SEPARATOR_ABSENT                                                    = 0x8029011C
	TPMAPI_E_TCG_INVALID_DIGEST_ENTRY                                                = 0x8029011D
	TBSIMP_E_BUFFER_TOO_SMALL                                                        = 0x80290200
	TBSIMP_E_CLEANUP_FAILED                                                          = 0x80290201
	TBSIMP_E_INVALID_CONTEXT_HANDLE                                                  = 0x80290202
	TBSIMP_E_INVALID_CONTEXT_PARAM                                                   = 0x80290203
	TBSIMP_E_TPM_ERROR                                                               = 0x80290204
	TBSIMP_E_HASH_BAD_KEY                                                            = 0x80290205
	TBSIMP_E_DUPLICATE_VHANDLE                                                       = 0x80290206
	TBSIMP_E_INVALID_OUTPUT_POINTER                                                  = 0x80290207
	TBSIMP_E_INVALID_PARAMETER                                                       = 0x80290208
	TBSIMP_E_RPC_INIT_FAILED                                                         = 0x80290209
	TBSIMP_E_SCHEDULER_NOT_RUNNING                                                   = 0x8029020A
	TBSIMP_E_COMMAND_CANCELED                                                        = 0x8029020B
	TBSIMP_E_OUT_OF_MEMORY                                                           = 0x8029020C
	TBSIMP_E_LIST_NO_MORE_ITEMS                                                      = 0x8029020D
	TBSIMP_E_LIST_NOT_FOUND                                                          = 0x8029020E
	TBSIMP_E_NOT_ENOUGH_SPACE                                                        = 0x8029020F
	TBSIMP_E_NOT_ENOUGH_TPM_CONTEXTS                                                 = 0x80290210
	TBSIMP_E_COMMAND_FAILED                                                          = 0x80290211
	TBSIMP_E_UNKNOWN_ORDINAL                                                         = 0x80290212
	TBSIMP_E_RESOURCE_EXPIRED                                                        = 0x80290213
	TBSIMP_E_INVALID_RESOURCE                                                        = 0x80290214
	TBSIMP_E_NOTHING_TO_UNLOAD                                                       = 0x80290215
	TBSIMP_E_HASH_TABLE_FULL                                                         = 0x80290216
	TBSIMP_E_TOO_MANY_TBS_CONTEXTS                                                   = 0x80290217
	TBSIMP_E_TOO_MANY_RESOURCES                                                      = 0x80290218
	TBSIMP_E_PPI_NOT_SUPPORTED                                                       = 0x80290219
	TBSIMP_E_TPM_INCOMPATIBLE                                                        = 0x8029021A
	TBSIMP_E_NO_EVENT_LOG                                                            = 0x8029021B
	TPM_E_PPI_ACPI_FAILURE                                                           = 0x80290300
	TPM_E_PPI_USER_ABORT                                                             = 0x80290301
	TPM_E_PPI_BIOS_FAILURE                                                           = 0x80290302
	TPM_E_PPI_NOT_SUPPORTED                                                          = 0x80290303
	PLA_E_DCS_NOT_FOUND                                                              = 0x80300002
	PLA_E_DCS_IN_USE                                                                 = 0x803000AA
	PLA_E_TOO_MANY_FOLDERS                                                           = 0x80300045
	PLA_E_NO_MIN_DISK                                                                = 0x80300070
	PLA_E_DCS_ALREADY_EXISTS                                                         = 0x803000B7
	PLA_S_PROPERTY_IGNORED                                                           = 0x00300100
	PLA_E_PROPERTY_CONFLICT                                                          = 0x80300101
	PLA_E_DCS_SINGLETON_REQUIRED                                                     = 0x80300102
	PLA_E_CREDENTIALS_REQUIRED                                                       = 0x80300103
	PLA_E_DCS_NOT_RUNNING                                                            = 0x80300104
	PLA_E_CONFLICT_INCL_EXCL_API                                                     = 0x80300105
	PLA_E_NETWORK_EXE_NOT_VALID                                                      = 0x80300106
	PLA_E_EXE_ALREADY_CONFIGURED                                                     = 0x80300107
	PLA_E_EXE_PATH_NOT_VALID                                                         = 0x80300108
	PLA_E_DC_ALREADY_EXISTS                                                          = 0x80300109
	PLA_E_DCS_START_WAIT_TIMEOUT                                                     = 0x8030010A
	PLA_E_DC_START_WAIT_TIMEOUT                                                      = 0x8030010B
	PLA_E_REPORT_WAIT_TIMEOUT                                                        = 0x8030010C
	PLA_E_NO_DUPLICATES                                                              = 0x8030010D
	PLA_E_EXE_FULL_PATH_REQUIRED                                                     = 0x8030010E
	PLA_E_INVALID_SESSION_NAME                                                       = 0x8030010F
	PLA_E_PLA_CHANNEL_NOT_ENABLED                                                    = 0x80300110
	PLA_E_TASKSCHED_CHANNEL_NOT_ENABLED                                              = 0x80300111
	PLA_E_RULES_MANAGER_FAILED                                                       = 0x80300112
	PLA_E_CABAPI_FAILURE                                                             = 0x80300113
	FVE_E_LOCKED_VOLUME                                                              = 0x80310000
	FVE_E_NOT_ENCRYPTED                                                              = 0x80310001
	FVE_E_NO_TPM_BIOS                                                                = 0x80310002
	FVE_E_NO_MBR_METRIC                                                              = 0x80310003
	FVE_E_NO_BOOTSECTOR_METRIC                                                       = 0x80310004
	FVE_E_NO_BOOTMGR_METRIC                                                          = 0x80310005
	FVE_E_WRONG_BOOTMGR                                                              = 0x80310006
	FVE_E_SECURE_KEY_REQUIRED                                                        = 0x80310007
	FVE_E_NOT_ACTIVATED                                                              = 0x80310008
	FVE_E_ACTION_NOT_ALLOWED                                                         = 0x80310009
	FVE_E_AD_SCHEMA_NOT_INSTALLED                                                    = 0x8031000A
	FVE_E_AD_INVALID_DATATYPE                                                        = 0x8031000B
	FVE_E_AD_INVALID_DATASIZE                                                        = 0x8031000C
	FVE_E_AD_NO_VALUES                                                               = 0x8031000D
	FVE_E_AD_ATTR_NOT_SET                                                            = 0x8031000E
	FVE_E_AD_GUID_NOT_FOUND                                                          = 0x8031000F
	FVE_E_BAD_INFORMATION                                                            = 0x80310010
	FVE_E_TOO_SMALL                                                                  = 0x80310011
	FVE_E_SYSTEM_VOLUME                                                              = 0x80310012
	FVE_E_FAILED_WRONG_FS                                                            = 0x80310013
	FVE_E_BAD_PARTITION_SIZE                                                         = 0x80310014
	FVE_E_NOT_SUPPORTED                                                              = 0x80310015
	FVE_E_BAD_DATA                                                                   = 0x80310016
	FVE_E_VOLUME_NOT_BOUND                                                           = 0x80310017
	FVE_E_TPM_NOT_OWNED                                                              = 0x80310018
	FVE_E_NOT_DATA_VOLUME                                                            = 0x80310019
	FVE_E_AD_INSUFFICIENT_BUFFER                                                     = 0x8031001A
	FVE_E_CONV_READ                                                                  = 0x8031001B
	FVE_E_CONV_WRITE                                                                 = 0x8031001C
	FVE_E_KEY_REQUIRED                                                               = 0x8031001D
	FVE_E_CLUSTERING_NOT_SUPPORTED                                                   = 0x8031001E
	FVE_E_VOLUME_BOUND_ALREADY                                                       = 0x8031001F
	FVE_E_OS_NOT_PROTECTED                                                           = 0x80310020
	FVE_E_PROTECTION_DISABLED                                                        = 0x80310021
	FVE_E_RECOVERY_KEY_REQUIRED                                                      = 0x80310022
	FVE_E_FOREIGN_VOLUME                                                             = 0x80310023
	FVE_E_OVERLAPPED_UPDATE                                                          = 0x80310024
	FVE_E_TPM_SRK_AUTH_NOT_ZERO                                                      = 0x80310025
	FVE_E_FAILED_SECTOR_SIZE                                                         = 0x80310026
	FVE_E_FAILED_AUTHENTICATION                                                      = 0x80310027
	FVE_E_NOT_OS_VOLUME                                                              = 0x80310028
	FVE_E_AUTOUNLOCK_ENABLED                                                         = 0x80310029
	FVE_E_WRONG_BOOTSECTOR                                                           = 0x8031002A
	FVE_E_WRONG_SYSTEM_FS                                                            = 0x8031002B
	FVE_E_POLICY_PASSWORD_REQUIRED                                                   = 0x8031002C
	FVE_E_CANNOT_SET_FVEK_ENCRYPTED                                                  = 0x8031002D
	FVE_E_CANNOT_ENCRYPT_NO_KEY                                                      = 0x8031002E
	FVE_E_BOOTABLE_CDDVD                                                             = 0x80310030
	FVE_E_PROTECTOR_EXISTS                                                           = 0x80310031
	FVE_E_RELATIVE_PATH                                                              = 0x80310032
	FVE_E_PROTECTOR_NOT_FOUND                                                        = 0x80310033
	FVE_E_INVALID_KEY_FORMAT                                                         = 0x80310034
	FVE_E_INVALID_PASSWORD_FORMAT                                                    = 0x80310035
	FVE_E_FIPS_RNG_CHECK_FAILED                                                      = 0x80310036
	FVE_E_FIPS_PREVENTS_RECOVERY_PASSWORD                                            = 0x80310037
	FVE_E_FIPS_PREVENTS_EXTERNAL_KEY_EXPORT                                          = 0x80310038
	FVE_E_NOT_DECRYPTED                                                              = 0x80310039
	FVE_E_INVALID_PROTECTOR_TYPE                                                     = 0x8031003A
	FVE_E_NO_PROTECTORS_TO_TEST                                                      = 0x8031003B
	FVE_E_KEYFILE_NOT_FOUND                                                          = 0x8031003C
	FVE_E_KEYFILE_INVALID                                                            = 0x8031003D
	FVE_E_KEYFILE_NO_VMK                                                             = 0x8031003E
	FVE_E_TPM_DISABLED                                                               = 0x8031003F
	FVE_E_NOT_ALLOWED_IN_SAFE_MODE                                                   = 0x80310040
	FVE_E_TPM_INVALID_PCR                                                            = 0x80310041
	FVE_E_TPM_NO_VMK                                                                 = 0x80310042
	FVE_E_PIN_INVALID                                                                = 0x80310043
	FVE_E_AUTH_INVALID_APPLICATION                                                   = 0x80310044
	FVE_E_AUTH_INVALID_CONFIG                                                        = 0x80310045
	FVE_E_FIPS_DISABLE_PROTECTION_NOT_ALLOWED                                        = 0x80310046
	FVE_E_FS_NOT_EXTENDED                                                            = 0x80310047
	FVE_E_FIRMWARE_TYPE_NOT_SUPPORTED                                                = 0x80310048
	FVE_E_NO_LICENSE                                                                 = 0x80310049
	FVE_E_NOT_ON_STACK                                                               = 0x8031004A
	FVE_E_FS_MOUNTED                                                                 = 0x8031004B
	FVE_E_TOKEN_NOT_IMPERSONATED                                                     = 0x8031004C
	FVE_E_DRY_RUN_FAILED                                                             = 0x8031004D
	FVE_E_REBOOT_REQUIRED                                                            = 0x8031004E
	FVE_E_DEBUGGER_ENABLED                                                           = 0x8031004F
	FVE_E_RAW_ACCESS                                                                 = 0x80310050
	FVE_E_RAW_BLOCKED                                                                = 0x80310051
	FVE_E_BCD_APPLICATIONS_PATH_INCORRECT                                            = 0x80310052
	FVE_E_NOT_ALLOWED_IN_VERSION                                                     = 0x80310053
	FVE_E_NO_AUTOUNLOCK_MASTER_KEY                                                   = 0x80310054
	FVE_E_MOR_FAILED                                                                 = 0x80310055
	FVE_E_HIDDEN_VOLUME                                                              = 0x80310056
	FVE_E_TRANSIENT_STATE                                                            = 0x80310057
	FVE_E_PUBKEY_NOT_ALLOWED                                                         = 0x80310058
	FVE_E_VOLUME_HANDLE_OPEN                                                         = 0x80310059
	FVE_E_NO_FEATURE_LICENSE                                                         = 0x8031005A
	FVE_E_INVALID_STARTUP_OPTIONS                                                    = 0x8031005B
	FVE_E_POLICY_RECOVERY_PASSWORD_NOT_ALLOWED                                       = 0x8031005C
	FVE_E_POLICY_RECOVERY_PASSWORD_REQUIRED                                          = 0x8031005D
	FVE_E_POLICY_RECOVERY_KEY_NOT_ALLOWED                                            = 0x8031005E
	FVE_E_POLICY_RECOVERY_KEY_REQUIRED                                               = 0x8031005F
	FVE_E_POLICY_STARTUP_PIN_NOT_ALLOWED                                             = 0x80310060
	FVE_E_POLICY_STARTUP_PIN_REQUIRED                                                = 0x80310061
	FVE_E_POLICY_STARTUP_KEY_NOT_ALLOWED                                             = 0x80310062
	FVE_E_POLICY_STARTUP_KEY_REQUIRED                                                = 0x80310063
	FVE_E_POLICY_STARTUP_PIN_KEY_NOT_ALLOWED                                         = 0x80310064
	FVE_E_POLICY_STARTUP_PIN_KEY_REQUIRED                                            = 0x80310065
	FVE_E_POLICY_STARTUP_TPM_NOT_ALLOWED                                             = 0x80310066
	FVE_E_POLICY_STARTUP_TPM_REQUIRED                                                = 0x80310067
	FVE_E_POLICY_INVALID_PIN_LENGTH                                                  = 0x80310068
	FVE_E_KEY_PROTECTOR_NOT_SUPPORTED                                                = 0x80310069
	FVE_E_POLICY_PASSPHRASE_NOT_ALLOWED                                              = 0x8031006A
	FVE_E_POLICY_PASSPHRASE_REQUIRED                                                 = 0x8031006B
	FVE_E_FIPS_PREVENTS_PASSPHRASE                                                   = 0x8031006C
	FVE_E_OS_VOLUME_PASSPHRASE_NOT_ALLOWED                                           = 0x8031006D
	FVE_E_INVALID_BITLOCKER_OID                                                      = 0x8031006E
	FVE_E_VOLUME_TOO_SMALL                                                           = 0x8031006F
	FVE_E_DV_NOT_SUPPORTED_ON_FS                                                     = 0x80310070
	FVE_E_DV_NOT_ALLOWED_BY_GP                                                       = 0x80310071
	FVE_E_POLICY_USER_CERTIFICATE_NOT_ALLOWED                                        = 0x80310072
	FVE_E_POLICY_USER_CERTIFICATE_REQUIRED                                           = 0x80310073
	FVE_E_POLICY_USER_CERT_MUST_BE_HW                                                = 0x80310074
	FVE_E_POLICY_USER_CONFIGURE_FDV_AUTOUNLOCK_NOT_ALLOWED                           = 0x80310075
	FVE_E_POLICY_USER_CONFIGURE_RDV_AUTOUNLOCK_NOT_ALLOWED                           = 0x80310076
	FVE_E_POLICY_USER_CONFIGURE_RDV_NOT_ALLOWED                                      = 0x80310077
	FVE_E_POLICY_USER_ENABLE_RDV_NOT_ALLOWED                                         = 0x80310078
	FVE_E_POLICY_USER_DISABLE_RDV_NOT_ALLOWED                                        = 0x80310079
	FVE_E_POLICY_INVALID_PASSPHRASE_LENGTH                                           = 0x80310080
	FVE_E_POLICY_PASSPHRASE_TOO_SIMPLE                                               = 0x80310081
	FVE_E_RECOVERY_PARTITION                                                         = 0x80310082
	FVE_E_POLICY_CONFLICT_FDV_RK_OFF_AUK_ON                                          = 0x80310083
	FVE_E_POLICY_CONFLICT_RDV_RK_OFF_AUK_ON                                          = 0x80310084
	FVE_E_NON_BITLOCKER_OID                                                          = 0x80310085
	FVE_E_POLICY_PROHIBITS_SELFSIGNED                                                = 0x80310086
	FVE_E_POLICY_CONFLICT_RO_AND_STARTUP_KEY_REQUIRED                                = 0x80310087
	FVE_E_CONV_RECOVERY_FAILED                                                       = 0x80310088
	FVE_E_VIRTUALIZED_SPACE_TOO_BIG                                                  = 0x80310089
	FVE_E_POLICY_CONFLICT_OSV_RP_OFF_ADB_ON                                          = 0x80310090
	FVE_E_POLICY_CONFLICT_FDV_RP_OFF_ADB_ON                                          = 0x80310091
	FVE_E_POLICY_CONFLICT_RDV_RP_OFF_ADB_ON                                          = 0x80310092
	FVE_E_NON_BITLOCKER_KU                                                           = 0x80310093
	FVE_E_PRIVATEKEY_AUTH_FAILED                                                     = 0x80310094
	FVE_E_REMOVAL_OF_DRA_FAILED                                                      = 0x80310095
	FVE_E_OPERATION_NOT_SUPPORTED_ON_VISTA_VOLUME                                    = 0x80310096
	FVE_E_CANT_LOCK_AUTOUNLOCK_ENABLED_VOLUME                                        = 0x80310097
	FVE_E_FIPS_HASH_KDF_NOT_ALLOWED                                                  = 0x80310098
	FVE_E_ENH_PIN_INVALID                                                            = 0x80310099
	FVE_E_INVALID_PIN_CHARS                                                          = 0x8031009A
	FVE_E_INVALID_DATUM_TYPE                                                         = 0x8031009B
	FWP_E_CALLOUT_NOT_FOUND                                                          = 0x80320001
	FWP_E_CONDITION_NOT_FOUND                                                        = 0x80320002
	FWP_E_FILTER_NOT_FOUND                                                           = 0x80320003
	FWP_E_LAYER_NOT_FOUND                                                            = 0x80320004
	FWP_E_PROVIDER_NOT_FOUND                                                         = 0x80320005
	FWP_E_PROVIDER_CONTEXT_NOT_FOUND                                                 = 0x80320006
	FWP_E_SUBLAYER_NOT_FOUND                                                         = 0x80320007
	FWP_E_NOT_FOUND                                                                  = 0x80320008
	FWP_E_ALREADY_EXISTS                                                             = 0x80320009
	FWP_E_IN_USE                                                                     = 0x8032000A
	FWP_E_DYNAMIC_SESSION_IN_PROGRESS                                                = 0x8032000B
	FWP_E_WRONG_SESSION                                                              = 0x8032000C
	FWP_E_NO_TXN_IN_PROGRESS                                                         = 0x8032000D
	FWP_E_TXN_IN_PROGRESS                                                            = 0x8032000E
	FWP_E_TXN_ABORTED                                                                = 0x8032000F
	FWP_E_SESSION_ABORTED                                                            = 0x80320010
	FWP_E_INCOMPATIBLE_TXN                                                           = 0x80320011
	FWP_E_TIMEOUT                                                                    = 0x80320012
	FWP_E_NET_EVENTS_DISABLED                                                        = 0x80320013
	FWP_E_INCOMPATIBLE_LAYER                                                         = 0x80320014
	FWP_E_KM_CLIENTS_ONLY                                                            = 0x80320015
	FWP_E_LIFETIME_MISMATCH                                                          = 0x80320016
	FWP_E_BUILTIN_OBJECT                                                             = 0x80320017
	FWP_E_TOO_MANY_CALLOUTS                                                          = 0x80320018
	FWP_E_NOTIFICATION_DROPPED                                                       = 0x80320019
	FWP_E_TRAFFIC_MISMATCH                                                           = 0x8032001A
	FWP_E_INCOMPATIBLE_SA_STATE                                                      = 0x8032001B
	FWP_E_NULL_POINTER                                                               = 0x8032001C
	FWP_E_INVALID_ENUMERATOR                                                         = 0x8032001D
	FWP_E_INVALID_FLAGS                                                              = 0x8032001E
	FWP_E_INVALID_NET_MASK                                                           = 0x8032001F
	FWP_E_INVALID_RANGE                                                              = 0x80320020
	FWP_E_INVALID_INTERVAL                                                           = 0x80320021
	FWP_E_ZERO_LENGTH_ARRAY                                                          = 0x80320022
	FWP_E_NULL_DISPLAY_NAME                                                          = 0x80320023
	FWP_E_INVALID_ACTION_TYPE                                                        = 0x80320024
	FWP_E_INVALID_WEIGHT                                                             = 0x80320025
	FWP_E_MATCH_TYPE_MISMATCH                                                        = 0x80320026
	FWP_E_TYPE_MISMATCH                                                              = 0x80320027
	FWP_E_OUT_OF_BOUNDS                                                              = 0x80320028
	FWP_E_RESERVED                                                                   = 0x80320029
	FWP_E_DUPLICATE_CONDITION                                                        = 0x8032002A
	FWP_E_DUPLICATE_KEYMOD                                                           = 0x8032002B
	FWP_E_ACTION_INCOMPATIBLE_WITH_LAYER                                             = 0x8032002C
	FWP_E_ACTION_INCOMPATIBLE_WITH_SUBLAYER                                          = 0x8032002D
	FWP_E_CONTEXT_INCOMPATIBLE_WITH_LAYER                                            = 0x8032002E
	FWP_E_CONTEXT_INCOMPATIBLE_WITH_CALLOUT                                          = 0x8032002F
	FWP_E_INCOMPATIBLE_AUTH_METHOD                                                   = 0x80320030
	FWP_E_INCOMPATIBLE_DH_GROUP                                                      = 0x80320031
	FWP_E_EM_NOT_SUPPORTED                                                           = 0x80320032
	FWP_E_NEVER_MATCH                                                                = 0x80320033
	FWP_E_PROVIDER_CONTEXT_MISMATCH                                                  = 0x80320034
	FWP_E_INVALID_PARAMETER                                                          = 0x80320035
	FWP_E_TOO_MANY_SUBLAYERS                                                         = 0x80320036
	FWP_E_CALLOUT_NOTIFICATION_FAILED                                                = 0x80320037
	FWP_E_INVALID_AUTH_TRANSFORM                                                     = 0x80320038
	FWP_E_INVALID_CIPHER_TRANSFORM                                                   = 0x80320039
	FWP_E_DROP_NOICMP                                                                = 0x80320104
	FWP_E_INCOMPATIBLE_CIPHER_TRANSFORM                                              = 0x8032003A
	FWP_E_INVALID_TRANSFORM_COMBINATION                                              = 0x8032003B
	FWP_E_DUPLICATE_AUTH_METHOD                                                      = 0x8032003C
	WS_S_ASYNC                                                                       = 0x003D0000
	WS_S_END                                                                         = 0x003D0001
	WS_E_INVALID_FORMAT                                                              = 0x803D0000
	WS_E_OBJECT_FAULTED                                                              = 0x803D0001
	WS_E_NUMERIC_OVERFLOW                                                            = 0x803D0002
	WS_E_INVALID_OPERATION                                                           = 0x803D0003
	WS_E_OPERATION_ABORTED                                                           = 0x803D0004
	WS_E_ENDPOINT_ACCESS_DENIED                                                      = 0x803D0005
	WS_E_OPERATION_TIMED_OUT                                                         = 0x803D0006
	WS_E_OPERATION_ABANDONED                                                         = 0x803D0007
	WS_E_QUOTA_EXCEEDED                                                              = 0x803D0008
	WS_E_NO_TRANSLATION_AVAILABLE                                                    = 0x803D0009
	WS_E_SECURITY_VERIFICATION_FAILURE                                               = 0x803D000A
	WS_E_ADDRESS_IN_USE                                                              = 0x803D000B
	WS_E_ADDRESS_NOT_AVAILABLE                                                       = 0x803D000C
	WS_E_ENDPOINT_NOT_FOUND                                                          = 0x803D000D
	WS_E_ENDPOINT_NOT_AVAILABLE                                                      = 0x803D000E
	WS_E_ENDPOINT_FAILURE                                                            = 0x803D000F
	WS_E_ENDPOINT_UNREACHABLE                                                        = 0x803D0010
	WS_E_ENDPOINT_ACTION_NOT_SUPPORTED                                               = 0x803D0011
	WS_E_ENDPOINT_TOO_BUSY                                                           = 0x803D0012
	WS_E_ENDPOINT_FAULT_RECEIVED                                                     = 0x803D0013
	WS_E_ENDPOINT_DISCONNECTED                                                       = 0x803D0014
	WS_E_PROXY_FAILURE                                                               = 0x803D0015
	WS_E_PROXY_ACCESS_DENIED                                                         = 0x803D0016
	WS_E_NOT_SUPPORTED                                                               = 0x803D0017
	WS_E_PROXY_REQUIRES_BASIC_AUTH                                                   = 0x803D0018
	WS_E_PROXY_REQUIRES_DIGEST_AUTH                                                  = 0x803D0019
	WS_E_PROXY_REQUIRES_NTLM_AUTH                                                    = 0x803D001A
	WS_E_PROXY_REQUIRES_NEGOTIATE_AUTH                                               = 0x803D001B
	WS_E_SERVER_REQUIRES_BASIC_AUTH                                                  = 0x803D001C
	WS_E_SERVER_REQUIRES_DIGEST_AUTH                                                 = 0x803D001D
	WS_E_SERVER_REQUIRES_NTLM_AUTH                                                   = 0x803D001E
	WS_E_SERVER_REQUIRES_NEGOTIATE_AUTH                                              = 0x803D001F
	WS_E_INVALID_ENDPOINT_URL                                                        = 0x803D0020
	WS_E_OTHER                                                                       = 0x803D0021
	WS_E_SECURITY_TOKEN_EXPIRED                                                      = 0x803D0022
	WS_E_SECURITY_SYSTEM_FAILURE                                                     = 0x803D0023
	ERROR_NDIS_INTERFACE_CLOSING                                                     = 0x80340002
	ERROR_NDIS_BAD_VERSION                                                           = 0x80340004
	ERROR_NDIS_BAD_CHARACTERISTICS                                                   = 0x80340005
	ERROR_NDIS_ADAPTER_NOT_FOUND                                                     = 0x80340006
	ERROR_NDIS_OPEN_FAILED                                                           = 0x80340007
	ERROR_NDIS_DEVICE_FAILED                                                         = 0x80340008
	ERROR_NDIS_MULTICAST_FULL                                                        = 0x80340009
	ERROR_NDIS_MULTICAST_EXISTS                                                      = 0x8034000A
	ERROR_NDIS_MULTICAST_NOT_FOUND                                                   = 0x8034000B
	ERROR_NDIS_REQUEST_ABORTED                                                       = 0x8034000C
	ERROR_NDIS_RESET_IN_PROGRESS                                                     = 0x8034000D
	ERROR_NDIS_NOT_SUPPORTED                                                         = 0x803400BB
	ERROR_NDIS_INVALID_PACKET                                                        = 0x8034000F
	ERROR_NDIS_ADAPTER_NOT_READY                                                     = 0x80340011
	ERROR_NDIS_INVALID_LENGTH                                                        = 0x80340014
	ERROR_NDIS_INVALID_DATA                                                          = 0x80340015
	ERROR_NDIS_BUFFER_TOO_SHORT                                                      = 0x80340016
	ERROR_NDIS_INVALID_OID                                                           = 0x80340017
	ERROR_NDIS_ADAPTER_REMOVED                                                       = 0x80340018
	ERROR_NDIS_UNSUPPORTED_MEDIA                                                     = 0x80340019
	ERROR_NDIS_GROUP_ADDRESS_IN_USE                                                  = 0x8034001A
	ERROR_NDIS_FILE_NOT_FOUND                                                        = 0x8034001B
	ERROR_NDIS_ERROR_READING_FILE                                                    = 0x8034001C
	ERROR_NDIS_ALREADY_MAPPED                                                        = 0x8034001D
	ERROR_NDIS_RESOURCE_CONFLICT                                                     = 0x8034001E
	ERROR_NDIS_MEDIA_DISCONNECTED                                                    = 0x8034001F
	ERROR_NDIS_INVALID_ADDRESS                                                       = 0x80340022
	ERROR_NDIS_INVALID_DEVICE_REQUEST                                                = 0x80340010
	ERROR_NDIS_PAUSED                                                                = 0x8034002A
	ERROR_NDIS_INTERFACE_NOT_FOUND                                                   = 0x8034002B
	ERROR_NDIS_UNSUPPORTED_REVISION                                                  = 0x8034002C
	ERROR_NDIS_INVALID_PORT                                                          = 0x8034002D
	ERROR_NDIS_INVALID_PORT_STATE                                                    = 0x8034002E
	ERROR_NDIS_LOW_POWER_STATE                                                       = 0x8034002F
	ERROR_NDIS_DOT11_AUTO_CONFIG_ENABLED                                             = 0x80342000
	ERROR_NDIS_DOT11_MEDIA_IN_USE                                                    = 0x80342001
	ERROR_NDIS_DOT11_POWER_STATE_INVALID                                             = 0x80342002
	ERROR_NDIS_PM_WOL_PATTERN_LIST_FULL                                              = 0x80342003
	ERROR_NDIS_PM_PROTOCOL_OFFLOAD_LIST_FULL                                         = 0x80342004
	ERROR_NDIS_INDICATION_REQUIRED                                                   = 0x00340001
	ERROR_NDIS_OFFLOAD_POLICY                                                        = 0xC034100F
	ERROR_NDIS_OFFLOAD_CONNECTION_REJECTED                                           = 0xC0341012
	ERROR_NDIS_OFFLOAD_PATH_REJECTED                                                 = 0xC0341013
	ERROR_HV_INVALID_HYPERCALL_CODE                                                  = 0xC0350002
	ERROR_HV_INVALID_HYPERCALL_INPUT                                                 = 0xC0350003
	ERROR_HV_INVALID_ALIGNMENT                                                       = 0xC0350004
	ERROR_HV_INVALID_PARAMETER                                                       = 0xC0350005
	ERROR_HV_ACCESS_DENIED                                                           = 0xC0350006
	ERROR_HV_INVALID_PARTITION_STATE                                                 = 0xC0350007
	ERROR_HV_OPERATION_DENIED                                                        = 0xC0350008
	ERROR_HV_UNKNOWN_PROPERTY                                                        = 0xC0350009
	ERROR_HV_PROPERTY_VALUE_OUT_OF_RANGE                                             = 0xC035000A
	ERROR_HV_INSUFFICIENT_MEMORY                                                     = 0xC035000B
	ERROR_HV_PARTITION_TOO_DEEP                                                      = 0xC035000C
	ERROR_HV_INVALID_PARTITION_ID                                                    = 0xC035000D
	ERROR_HV_INVALID_VP_INDEX                                                        = 0xC035000E
	ERROR_HV_INVALID_PORT_ID                                                         = 0xC0350011
	ERROR_HV_INVALID_CONNECTION_ID                                                   = 0xC0350012
	ERROR_HV_INSUFFICIENT_BUFFERS                                                    = 0xC0350013
	ERROR_HV_NOT_ACKNOWLEDGED                                                        = 0xC0350014
	ERROR_HV_ACKNOWLEDGED                                                            = 0xC0350016
	ERROR_HV_INVALID_SAVE_RESTORE_STATE                                              = 0xC0350017
	ERROR_HV_INVALID_SYNIC_STATE                                                     = 0xC0350018
	ERROR_HV_OBJECT_IN_USE                                                           = 0xC0350019
	ERROR_HV_INVALID_PROXIMITY_DOMAIN_INFO                                           = 0xC035001A
	ERROR_HV_NO_DATA                                                                 = 0xC035001B
	ERROR_HV_INACTIVE                                                                = 0xC035001C
	ERROR_HV_NO_RESOURCES                                                            = 0xC035001D
	ERROR_HV_FEATURE_UNAVAILABLE                                                     = 0xC035001E
	ERROR_HV_NOT_PRESENT                                                             = 0xC0351000
	ERROR_VID_DUPLICATE_HANDLER                                                      = 0xC0370001
	ERROR_VID_TOO_MANY_HANDLERS                                                      = 0xC0370002
	ERROR_VID_QUEUE_FULL                                                             = 0xC0370003
	ERROR_VID_HANDLER_NOT_PRESENT                                                    = 0xC0370004
	ERROR_VID_INVALID_OBJECT_NAME                                                    = 0xC0370005
	ERROR_VID_PARTITION_NAME_TOO_LONG                                                = 0xC0370006
	ERROR_VID_MESSAGE_QUEUE_NAME_TOO_LONG                                            = 0xC0370007
	ERROR_VID_PARTITION_ALREADY_EXISTS                                               = 0xC0370008
	ERROR_VID_PARTITION_DOES_NOT_EXIST                                               = 0xC0370009
	ERROR_VID_PARTITION_NAME_NOT_FOUND                                               = 0xC037000A
	ERROR_VID_MESSAGE_QUEUE_ALREADY_EXISTS                                           = 0xC037000B
	ERROR_VID_EXCEEDED_MBP_ENTRY_MAP_LIMIT                                           = 0xC037000C
	ERROR_VID_MB_STILL_REFERENCED                                                    = 0xC037000D
	ERROR_VID_CHILD_GPA_PAGE_SET_CORRUPTED                                           = 0xC037000E
	ERROR_VID_INVALID_NUMA_SETTINGS                                                  = 0xC037000F
	ERROR_VID_INVALID_NUMA_NODE_INDEX                                                = 0xC0370010
	ERROR_VID_NOTIFICATION_QUEUE_ALREADY_ASSOCIATED                                  = 0xC0370011
	ERROR_VID_INVALID_MEMORY_BLOCK_HANDLE                                            = 0xC0370012
	ERROR_VID_PAGE_RANGE_OVERFLOW                                                    = 0xC0370013
	ERROR_VID_INVALID_MESSAGE_QUEUE_HANDLE                                           = 0xC0370014
	ERROR_VID_INVALID_GPA_RANGE_HANDLE                                               = 0xC0370015
	ERROR_VID_NO_MEMORY_BLOCK_NOTIFICATION_QUEUE                                     = 0xC0370016
	ERROR_VID_MEMORY_BLOCK_LOCK_COUNT_EXCEEDED                                       = 0xC0370017
	ERROR_VID_INVALID_PPM_HANDLE                                                     = 0xC0370018
	ERROR_VID_MBPS_ARE_LOCKED                                                        = 0xC0370019
	ERROR_VID_MESSAGE_QUEUE_CLOSED                                                   = 0xC037001A
	ERROR_VID_VIRTUAL_PROCESSOR_LIMIT_EXCEEDED                                       = 0xC037001B
	ERROR_VID_STOP_PENDING                                                           = 0xC037001C
	ERROR_VID_INVALID_PROCESSOR_STATE                                                = 0xC037001D
	ERROR_VID_EXCEEDED_KM_CONTEXT_COUNT_LIMIT                                        = 0xC037001E
	ERROR_VID_KM_INTERFACE_ALREADY_INITIALIZED                                       = 0xC037001F
	ERROR_VID_MB_PROPERTY_ALREADY_SET_RESET                                          = 0xC0370020
	ERROR_VID_MMIO_RANGE_DESTROYED                                                   = 0xC0370021
	ERROR_VID_INVALID_CHILD_GPA_PAGE_SET                                             = 0xC0370022
	ERROR_VID_RESERVE_PAGE_SET_IS_BEING_USED                                         = 0xC0370023
	ERROR_VID_RESERVE_PAGE_SET_TOO_SMALL                                             = 0xC0370024
	ERROR_VID_MBP_ALREADY_LOCKED_USING_RESERVED_PAGE                                 = 0xC0370025
	ERROR_VID_MBP_COUNT_EXCEEDED_LIMIT                                               = 0xC0370026
	ERROR_VID_SAVED_STATE_CORRUPT                                                    = 0xC0370027
	ERROR_VID_SAVED_STATE_UNRECOGNIZED_ITEM                                          = 0xC0370028
	ERROR_VID_SAVED_STATE_INCOMPATIBLE                                               = 0xC0370029
	ERROR_VID_REMOTE_NODE_PARENT_GPA_PAGES_USED                                      = 0x80370001
	ERROR_VOLMGR_INCOMPLETE_REGENERATION                                             = 0x80380001
	ERROR_VOLMGR_INCOMPLETE_DISK_MIGRATION                                           = 0x80380002
	ERROR_VOLMGR_DATABASE_FULL                                                       = 0xC0380001
	ERROR_VOLMGR_DISK_CONFIGURATION_CORRUPTED                                        = 0xC0380002
	ERROR_VOLMGR_DISK_CONFIGURATION_NOT_IN_SYNC                                      = 0xC0380003
	ERROR_VOLMGR_PACK_CONFIG_UPDATE_FAILED                                           = 0xC0380004
	ERROR_VOLMGR_DISK_CONTAINS_NON_SIMPLE_VOLUME                                     = 0xC0380005
	ERROR_VOLMGR_DISK_DUPLICATE                                                      = 0xC0380006
	ERROR_VOLMGR_DISK_DYNAMIC                                                        = 0xC0380007
	ERROR_VOLMGR_DISK_ID_INVALID                                                     = 0xC0380008
	ERROR_VOLMGR_DISK_INVALID                                                        = 0xC0380009
	ERROR_VOLMGR_DISK_LAST_VOTER                                                     = 0xC038000A
	ERROR_VOLMGR_DISK_LAYOUT_INVALID                                                 = 0xC038000B
	ERROR_VOLMGR_DISK_LAYOUT_NON_BASIC_BETWEEN_BASIC_PARTITIONS                      = 0xC038000C
	ERROR_VOLMGR_DISK_LAYOUT_NOT_CYLINDER_ALIGNED                                    = 0xC038000D
	ERROR_VOLMGR_DISK_LAYOUT_PARTITIONS_TOO_SMALL                                    = 0xC038000E
	ERROR_VOLMGR_DISK_LAYOUT_PRIMARY_BETWEEN_LOGICAL_PARTITIONS                      = 0xC038000F
	ERROR_VOLMGR_DISK_LAYOUT_TOO_MANY_PARTITIONS                                     = 0xC0380010
	ERROR_VOLMGR_DISK_MISSING                                                        = 0xC0380011
	ERROR_VOLMGR_DISK_NOT_EMPTY                                                      = 0xC0380012
	ERROR_VOLMGR_DISK_NOT_ENOUGH_SPACE                                               = 0xC0380013
	ERROR_VOLMGR_DISK_REVECTORING_FAILED                                             = 0xC0380014
	ERROR_VOLMGR_DISK_SECTOR_SIZE_INVALID                                            = 0xC0380015
	ERROR_VOLMGR_DISK_SET_NOT_CONTAINED                                              = 0xC0380016
	ERROR_VOLMGR_DISK_USED_BY_MULTIPLE_MEMBERS                                       = 0xC0380017
	ERROR_VOLMGR_DISK_USED_BY_MULTIPLE_PLEXES                                        = 0xC0380018
	ERROR_VOLMGR_DYNAMIC_DISK_NOT_SUPPORTED                                          = 0xC0380019
	ERROR_VOLMGR_EXTENT_ALREADY_USED                                                 = 0xC038001A
	ERROR_VOLMGR_EXTENT_NOT_CONTIGUOUS                                               = 0xC038001B
	ERROR_VOLMGR_EXTENT_NOT_IN_PUBLIC_REGION                                         = 0xC038001C
	ERROR_VOLMGR_EXTENT_NOT_SECTOR_ALIGNED                                           = 0xC038001D
	ERROR_VOLMGR_EXTENT_OVERLAPS_EBR_PARTITION                                       = 0xC038001E
	ERROR_VOLMGR_EXTENT_VOLUME_LENGTHS_DO_NOT_MATCH                                  = 0xC038001F
	ERROR_VOLMGR_FAULT_TOLERANT_NOT_SUPPORTED                                        = 0xC0380020
	ERROR_VOLMGR_INTERLEAVE_LENGTH_INVALID                                           = 0xC0380021
	ERROR_VOLMGR_MAXIMUM_REGISTERED_USERS                                            = 0xC0380022
	ERROR_VOLMGR_MEMBER_IN_SYNC                                                      = 0xC0380023
	ERROR_VOLMGR_MEMBER_INDEX_DUPLICATE                                              = 0xC0380024
	ERROR_VOLMGR_MEMBER_INDEX_INVALID                                                = 0xC0380025
	ERROR_VOLMGR_MEMBER_MISSING                                                      = 0xC0380026
	ERROR_VOLMGR_MEMBER_NOT_DETACHED                                                 = 0xC0380027
	ERROR_VOLMGR_MEMBER_REGENERATING                                                 = 0xC0380028
	ERROR_VOLMGR_ALL_DISKS_FAILED                                                    = 0xC0380029
	ERROR_VOLMGR_NO_REGISTERED_USERS                                                 = 0xC038002A
	ERROR_VOLMGR_NO_SUCH_USER                                                        = 0xC038002B
	ERROR_VOLMGR_NOTIFICATION_RESET                                                  = 0xC038002C
	ERROR_VOLMGR_NUMBER_OF_MEMBERS_INVALID                                           = 0xC038002D
	ERROR_VOLMGR_NUMBER_OF_PLEXES_INVALID                                            = 0xC038002E
	ERROR_VOLMGR_PACK_DUPLICATE                                                      = 0xC038002F
	ERROR_VOLMGR_PACK_ID_INVALID                                                     = 0xC0380030
	ERROR_VOLMGR_PACK_INVALID                                                        = 0xC0380031
	ERROR_VOLMGR_PACK_NAME_INVALID                                                   = 0xC0380032
	ERROR_VOLMGR_PACK_OFFLINE                                                        = 0xC0380033
	ERROR_VOLMGR_PACK_HAS_QUORUM                                                     = 0xC0380034
	ERROR_VOLMGR_PACK_WITHOUT_QUORUM                                                 = 0xC0380035
	ERROR_VOLMGR_PARTITION_STYLE_INVALID                                             = 0xC0380036
	ERROR_VOLMGR_PARTITION_UPDATE_FAILED                                             = 0xC0380037
	ERROR_VOLMGR_PLEX_IN_SYNC                                                        = 0xC0380038
	ERROR_VOLMGR_PLEX_INDEX_DUPLICATE                                                = 0xC0380039
	ERROR_VOLMGR_PLEX_INDEX_INVALID                                                  = 0xC038003A
	ERROR_VOLMGR_PLEX_LAST_ACTIVE                                                    = 0xC038003B
	ERROR_VOLMGR_PLEX_MISSING                                                        = 0xC038003C
	ERROR_VOLMGR_PLEX_REGENERATING                                                   = 0xC038003D
	ERROR_VOLMGR_PLEX_TYPE_INVALID                                                   = 0xC038003E
	ERROR_VOLMGR_PLEX_NOT_RAID5                                                      = 0xC038003F
	ERROR_VOLMGR_PLEX_NOT_SIMPLE                                                     = 0xC0380040
	ERROR_VOLMGR_STRUCTURE_SIZE_INVALID                                              = 0xC0380041
	ERROR_VOLMGR_TOO_MANY_NOTIFICATION_REQUESTS                                      = 0xC0380042
	ERROR_VOLMGR_TRANSACTION_IN_PROGRESS                                             = 0xC0380043
	ERROR_VOLMGR_UNEXPECTED_DISK_LAYOUT_CHANGE                                       = 0xC0380044
	ERROR_VOLMGR_VOLUME_CONTAINS_MISSING_DISK                                        = 0xC0380045
	ERROR_VOLMGR_VOLUME_ID_INVALID                                                   = 0xC0380046
	ERROR_VOLMGR_VOLUME_LENGTH_INVALID                                               = 0xC0380047
	ERROR_VOLMGR_VOLUME_LENGTH_NOT_SECTOR_SIZE_MULTIPLE                              = 0xC0380048
	ERROR_VOLMGR_VOLUME_NOT_MIRRORED                                                 = 0xC0380049
	ERROR_VOLMGR_VOLUME_NOT_RETAINED                                                 = 0xC038004A
	ERROR_VOLMGR_VOLUME_OFFLINE                                                      = 0xC038004B
	ERROR_VOLMGR_VOLUME_RETAINED                                                     = 0xC038004C
	ERROR_VOLMGR_NUMBER_OF_EXTENTS_INVALID                                           = 0xC038004D
	ERROR_VOLMGR_DIFFERENT_SECTOR_SIZE                                               = 0xC038004E
	ERROR_VOLMGR_BAD_BOOT_DISK                                                       = 0xC038004F
	ERROR_VOLMGR_PACK_CONFIG_OFFLINE                                                 = 0xC0380050
	ERROR_VOLMGR_PACK_CONFIG_ONLINE                                                  = 0xC0380051
	ERROR_VOLMGR_NOT_PRIMARY_PACK                                                    = 0xC0380052
	ERROR_VOLMGR_PACK_LOG_UPDATE_FAILED                                              = 0xC0380053
	ERROR_VOLMGR_NUMBER_OF_DISKS_IN_PLEX_INVALID                                     = 0xC0380054
	ERROR_VOLMGR_NUMBER_OF_DISKS_IN_MEMBER_INVALID                                   = 0xC0380055
	ERROR_VOLMGR_VOLUME_MIRRORED                                                     = 0xC0380056
	ERROR_VOLMGR_PLEX_NOT_SIMPLE_SPANNED                                             = 0xC0380057
	ERROR_VOLMGR_NO_VALID_LOG_COPIES                                                 = 0xC0380058
	ERROR_VOLMGR_PRIMARY_PACK_PRESENT                                                = 0xC0380059
	ERROR_VOLMGR_NUMBER_OF_DISKS_INVALID                                             = 0xC038005A
	ERROR_VOLMGR_MIRROR_NOT_SUPPORTED                                                = 0xC038005B
	ERROR_VOLMGR_RAID5_NOT_SUPPORTED                                                 = 0xC038005C
	ERROR_BCD_NOT_ALL_ENTRIES_IMPORTED                                               = 0x80390001
	ERROR_BCD_TOO_MANY_ELEMENTS                                                      = 0xC0390002
	ERROR_BCD_NOT_ALL_ENTRIES_SYNCHRONIZED                                           = 0x80390003
	ERROR_VHD_DRIVE_FOOTER_MISSING                                                   = 0xC03A0001
	ERROR_VHD_DRIVE_FOOTER_CHECKSUM_MISMATCH                                         = 0xC03A0002
	ERROR_VHD_DRIVE_FOOTER_CORRUPT                                                   = 0xC03A0003
	ERROR_VHD_FORMAT_UNKNOWN                                                         = 0xC03A0004
	ERROR_VHD_FORMAT_UNSUPPORTED_VERSION                                             = 0xC03A0005
	ERROR_VHD_SPARSE_HEADER_CHECKSUM_MISMATCH                                        = 0xC03A0006
	ERROR_VHD_SPARSE_HEADER_UNSUPPORTED_VERSION                                      = 0xC03A0007
	ERROR_VHD_SPARSE_HEADER_CORRUPT                                                  = 0xC03A0008
	ERROR_VHD_BLOCK_ALLOCATION_FAILURE                                               = 0xC03A0009
	ERROR_VHD_BLOCK_ALLOCATION_TABLE_CORRUPT                                         = 0xC03A000A
	ERROR_VHD_INVALID_BLOCK_SIZE                                                     = 0xC03A000B
	ERROR_VHD_BITMAP_MISMATCH                                                        = 0xC03A000C
	ERROR_VHD_PARENT_VHD_NOT_FOUND                                                   = 0xC03A000D
	ERROR_VHD_CHILD_PARENT_ID_MISMATCH                                               = 0xC03A000E
	ERROR_VHD_CHILD_PARENT_TIMESTAMP_MISMATCH                                        = 0xC03A000F
	ERROR_VHD_METADATA_READ_FAILURE                                                  = 0xC03A0010
	ERROR_VHD_METADATA_WRITE_FAILURE                                                 = 0xC03A0011
	ERROR_VHD_INVALID_SIZE                                                           = 0xC03A0012
	ERROR_VHD_INVALID_FILE_SIZE                                                      = 0xC03A0013
	ERROR_VIRTDISK_PROVIDER_NOT_FOUND                                                = 0xC03A0014
	ERROR_VIRTDISK_NOT_VIRTUAL_DISK                                                  = 0xC03A0015
	ERROR_VHD_PARENT_VHD_ACCESS_DENIED                                               = 0xC03A0016
	ERROR_VHD_CHILD_PARENT_SIZE_MISMATCH                                             = 0xC03A0017
	ERROR_VHD_DIFFERENCING_CHAIN_CYCLE_DETECTED                                      = 0xC03A0018
	ERROR_VHD_DIFFERENCING_CHAIN_ERROR_IN_PARENT                                     = 0xC03A0019
	ERROR_VIRTUAL_DISK_LIMITATION                                                    = 0xC03A001A
	ERROR_VHD_INVALID_TYPE                                                           = 0xC03A001B
	ERROR_VHD_INVALID_STATE                                                          = 0xC03A001C
	ERROR_VIRTDISK_UNSUPPORTED_DISK_SECTOR_SIZE                                      = 0xC03A001D
	ERROR_QUERY_STORAGE_ERROR                                                        = 0x803A0001
	SDIAG_E_CANCELLED                                                                = 0x803C0100
	SDIAG_E_SCRIPT                                                                   = 0x803C0101
	SDIAG_E_POWERSHELL                                                               = 0x803C0102
	SDIAG_E_MANAGEDHOST                                                              = 0x803C0103
	SDIAG_E_NOVERIFIER                                                               = 0x803C0104
	SDIAG_S_CANNOTRUN                                                                = 0x003C0105
	SDIAG_E_DISABLED                                                                 = 0x803C0106
	SDIAG_E_TRUST                                                                    = 0x803C0107
	SDIAG_E_CANNOTRUN                                                                = 0x803C0108
	SDIAG_E_VERSION                                                                  = 0x803C0109
	SDIAG_E_RESOURCE                                                                 = 0x803C010A
	SDIAG_E_ROOTCAUSE                                                                = 0x803C010B
	E_MBN_CONTEXT_NOT_ACTIVATED                                                      = 0x80548201
	E_MBN_BAD_SIM                                                                    = 0x80548202
	E_MBN_DATA_CLASS_NOT_AVAILABLE                                                   = 0x80548203
	E_MBN_INVALID_ACCESS_STRING                                                      = 0x80548204
	E_MBN_MAX_ACTIVATED_CONTEXTS                                                     = 0x80548205
	E_MBN_PACKET_SVC_DETACHED                                                        = 0x80548206
	E_MBN_PROVIDER_NOT_VISIBLE                                                       = 0x80548207
	E_MBN_RADIO_POWER_OFF                                                            = 0x80548208
	E_MBN_SERVICE_NOT_ACTIVATED                                                      = 0x80548209
	E_MBN_SIM_NOT_INSERTED                                                           = 0x8054820A
	E_MBN_VOICE_CALL_IN_PROGRESS                                                     = 0x8054820B
	E_MBN_INVALID_CACHE                                                              = 0x8054820C
	E_MBN_NOT_REGISTERED                                                             = 0x8054820D
	E_MBN_PROVIDERS_NOT_FOUND                                                        = 0x8054820E
	E_MBN_PIN_NOT_SUPPORTED                                                          = 0x8054820F
	E_MBN_PIN_REQUIRED                                                               = 0x80548210
	E_MBN_PIN_DISABLED                                                               = 0x80548211
	E_MBN_FAILURE                                                                    = 0x80548212
	E_MBN_INVALID_PROFILE                                                            = 0x80548218
	E_MBN_DEFAULT_PROFILE_EXIST                                                      = 0x80548219
	E_MBN_SMS_ENCODING_NOT_SUPPORTED                                                 = 0x80548220
	E_MBN_SMS_FILTER_NOT_SUPPORTED                                                   = 0x80548221
	E_MBN_SMS_INVALID_MEMORY_INDEX                                                   = 0x80548222
	E_MBN_SMS_LANG_NOT_SUPPORTED                                                     = 0x80548223
	E_MBN_SMS_MEMORY_FAILURE                                                         = 0x80548224
	E_MBN_SMS_NETWORK_TIMEOUT                                                        = 0x80548225
	E_MBN_SMS_UNKNOWN_SMSC_ADDRESS                                                   = 0x80548226
	E_MBN_SMS_FORMAT_NOT_SUPPORTED                                                   = 0x80548227
	E_MBN_SMS_OPERATION_NOT_ALLOWED                                                  = 0x80548228
	E_MBN_SMS_MEMORY_FULL                                                            = 0x80548229
	UI_E_CREATE_FAILED                                                               = 0x802A0001
	UI_E_SHUTDOWN_CALLED                                                             = 0x802A0002
	UI_E_ILLEGAL_REENTRANCY                                                          = 0x802A0003
	UI_E_OBJECT_SEALED                                                               = 0x802A0004
	UI_E_VALUE_NOT_SET                                                               = 0x802A0005
	UI_E_VALUE_NOT_DETERMINED                                                        = 0x802A0006
	UI_E_INVALID_OUTPUT                                                              = 0x802A0007
	UI_E_BOOLEAN_EXPECTED                                                            = 0x802A0008
	UI_E_DIFFERENT_OWNER                                                             = 0x802A0009
	UI_E_AMBIGUOUS_MATCH                                                             = 0x802A000A
	UI_E_FP_OVERFLOW                                                                 = 0x802A000B
	UI_E_WRONG_THREAD                                                                = 0x802A000C
	UI_E_STORYBOARD_ACTIVE                                                           = 0x802A0101
	UI_E_STORYBOARD_NOT_PLAYING                                                      = 0x802A0102
	UI_E_START_KEYFRAME_AFTER_END                                                    = 0x802A0103
	UI_E_END_KEYFRAME_NOT_DETERMINED                                                 = 0x802A0104
	UI_E_LOOPS_OVERLAP                                                               = 0x802A0105
	UI_E_TRANSITION_ALREADY_USED                                                     = 0x802A0106
	UI_E_TRANSITION_NOT_IN_STORYBOARD                                                = 0x802A0107
	UI_E_TRANSITION_ECLIPSED                                                         = 0x802A0108
	UI_E_TIME_BEFORE_LAST_UPDATE                                                     = 0x802A0109
	UI_E_TIMER_CLIENT_ALREADY_CONNECTED                                              = 0x802A010A
)

func (k ErrorsKind) String() string {
	switch k {
	case ERROR_SUCCESS:
		return "0"
	case ERROR_INVALID_FUNCTION:
		return "1"
	case ERROR_FILE_NOT_FOUND:
		return "2"
	case ERROR_PATH_NOT_FOUND:
		return "3"
	case ERROR_TOO_MANY_OPEN_FILES:
		return "4"
	case ERROR_ACCESS_DENIED:
		return "5"
	case ERROR_INVALID_HANDLE:
		return "6"
	case ERROR_ARENA_TRASHED:
		return "7"
	case ERROR_NOT_ENOUGH_MEMORY:
		return "8"
	case ERROR_INVALID_BLOCK:
		return "9"
	case ERROR_BAD_ENVIRONMENT:
		return "10"
	case ERROR_BAD_FORMAT:
		return "11"
	case ERROR_INVALID_ACCESS:
		return "12"
	case ERROR_INVALID_DATA:
		return "13"
	case ERROR_OUTOFMEMORY:
		return "14"
	case ERROR_INVALID_DRIVE:
		return "15"
	case ERROR_CURRENT_DIRECTORY:
		return "16"
	case ERROR_NOT_SAME_DEVICE:
		return "17"
	case ERROR_NO_MORE_FILES:
		return "18"
	case ERROR_WRITE_PROTECT:
		return "19"
	case ERROR_BAD_UNIT:
		return "20"
	case ERROR_NOT_READY:
		return "21"
	case ERROR_BAD_COMMAND:
		return "22"
	case ERROR_CRC:
		return "23"
	case ERROR_BAD_LENGTH:
		return "24"
	case ERROR_SEEK:
		return "25"
	case ERROR_NOT_DOS_DISK:
		return "26"
	case ERROR_SECTOR_NOT_FOUND:
		return "27"
	case ERROR_OUT_OF_PAPER:
		return "28"
	case ERROR_WRITE_FAULT:
		return "29"
	case ERROR_READ_FAULT:
		return "30"
	case ERROR_GEN_FAILURE:
		return "31"
	case ERROR_SHARING_VIOLATION:
		return "32"
	case ERROR_LOCK_VIOLATION:
		return "33"
	case ERROR_WRONG_DISK:
		return "34"
	case ERROR_SHARING_BUFFER_EXCEEDED:
		return "36"
	case ERROR_HANDLE_EOF:
		return "38"
	case ERROR_HANDLE_DISK_FULL:
		return "39"
	case ERROR_NOT_SUPPORTED:
		return "50"
	case ERROR_REM_NOT_LIST:
		return "51"
	case ERROR_DUP_NAME:
		return "52"
	case ERROR_BAD_NETPATH:
		return "53"
	case ERROR_NETWORK_BUSY:
		return "54"
	case ERROR_DEV_NOT_EXIST:
		return "55"
	case ERROR_TOO_MANY_CMDS:
		return "56"
	case ERROR_ADAP_HDW_ERR:
		return "57"
	case ERROR_BAD_NET_RESP:
		return "58"
	case ERROR_UNEXP_NET_ERR:
		return "59"
	case ERROR_BAD_REM_ADAP:
		return "60"
	case ERROR_PRINTQ_FULL:
		return "61"
	case ERROR_NO_SPOOL_SPACE:
		return "62"
	case ERROR_PRINT_CANCELLED:
		return "63"
	case ERROR_NETNAME_DELETED:
		return "64"
	case ERROR_NETWORK_ACCESS_DENIED:
		return "65"
	case ERROR_BAD_DEV_TYPE:
		return "66"
	case ERROR_BAD_NET_NAME:
		return "67"
	case ERROR_TOO_MANY_NAMES:
		return "68"
	case ERROR_TOO_MANY_SESS:
		return "69"
	case ERROR_SHARING_PAUSED:
		return "70"
	case ERROR_REQ_NOT_ACCEP:
		return "71"
	case ERROR_REDIR_PAUSED:
		return "72"
	case ERROR_FILE_EXISTS:
		return "80"
	case ERROR_CANNOT_MAKE:
		return "82"
	case ERROR_FAIL_I24:
		return "83"
	case ERROR_OUT_OF_STRUCTURES:
		return "84"
	case ERROR_ALREADY_ASSIGNED:
		return "85"
	case ERROR_INVALID_PASSWORD:
		return "86"
	case ERROR_INVALID_PARAMETER:
		return "87"
	case ERROR_NET_WRITE_FAULT:
		return "88"
	case ERROR_NO_PROC_SLOTS:
		return "89"
	case ERROR_TOO_MANY_SEMAPHORES:
		return "100"
	case ERROR_EXCL_SEM_ALREADY_OWNED:
		return "101"
	case ERROR_SEM_IS_SET:
		return "102"
	case ERROR_TOO_MANY_SEM_REQUESTS:
		return "103"
	case ERROR_INVALID_AT_INTERRUPT_TIME:
		return "104"
	case ERROR_SEM_OWNER_DIED:
		return "105"
	case ERROR_SEM_USER_LIMIT:
		return "106"
	case ERROR_DISK_CHANGE:
		return "107"
	case ERROR_DRIVE_LOCKED:
		return "108"
	case ERROR_BROKEN_PIPE:
		return "109"
	case ERROR_OPEN_FAILED:
		return "110"
	case ERROR_BUFFER_OVERFLOW:
		return "111"
	case ERROR_DISK_FULL:
		return "112"
	case ERROR_NO_MORE_SEARCH_HANDLES:
		return "113"
	case ERROR_INVALID_TARGET_HANDLE:
		return "114"
	case ERROR_INVALID_CATEGORY:
		return "117"
	case ERROR_INVALID_VERIFY_SWITCH:
		return "118"
	case ERROR_BAD_DRIVER_LEVEL:
		return "119"
	case ERROR_CALL_NOT_IMPLEMENTED:
		return "120"
	case ERROR_SEM_TIMEOUT:
		return "121"
	case ERROR_INSUFFICIENT_BUFFER:
		return "122"
	case ERROR_INVALID_NAME:
		return "123"
	case ERROR_INVALID_LEVEL:
		return "124"
	case ERROR_NO_VOLUME_LABEL:
		return "125"
	case ERROR_MOD_NOT_FOUND:
		return "126"
	case ERROR_PROC_NOT_FOUND:
		return "127"
	case ERROR_WAIT_NO_CHILDREN:
		return "128"
	case ERROR_CHILD_NOT_COMPLETE:
		return "129"
	case ERROR_DIRECT_ACCESS_HANDLE:
		return "130"
	case ERROR_NEGATIVE_SEEK:
		return "131"
	case ERROR_SEEK_ON_DEVICE:
		return "132"
	case ERROR_IS_JOIN_TARGET:
		return "133"
	case ERROR_IS_JOINED:
		return "134"
	case ERROR_IS_SUBSTED:
		return "135"
	case ERROR_NOT_JOINED:
		return "136"
	case ERROR_NOT_SUBSTED:
		return "137"
	case ERROR_JOIN_TO_JOIN:
		return "138"
	case ERROR_SUBST_TO_SUBST:
		return "139"
	case ERROR_JOIN_TO_SUBST:
		return "140"
	case ERROR_SUBST_TO_JOIN:
		return "141"
	case ERROR_BUSY_DRIVE:
		return "142"
	case ERROR_SAME_DRIVE:
		return "143"
	case ERROR_DIR_NOT_ROOT:
		return "144"
	case ERROR_DIR_NOT_EMPTY:
		return "145"
	case ERROR_IS_SUBST_PATH:
		return "146"
	case ERROR_IS_JOIN_PATH:
		return "147"
	case ERROR_PATH_BUSY:
		return "148"
	case ERROR_IS_SUBST_TARGET:
		return "149"
	case ERROR_SYSTEM_TRACE:
		return "150"
	case ERROR_INVALID_EVENT_COUNT:
		return "151"
	case ERROR_TOO_MANY_MUXWAITERS:
		return "152"
	case ERROR_INVALID_LIST_FORMAT:
		return "153"
	case ERROR_LABEL_TOO_LONG:
		return "154"
	case ERROR_TOO_MANY_TCBS:
		return "155"
	case ERROR_SIGNAL_REFUSED:
		return "156"
	case ERROR_DISCARDED:
		return "157"
	case ERROR_NOT_LOCKED:
		return "158"
	case ERROR_BAD_THREADID_ADDR:
		return "159"
	case ERROR_BAD_ARGUMENTS:
		return "160"
	case ERROR_BAD_PATHNAME:
		return "161"
	case ERROR_SIGNAL_PENDING:
		return "162"
	case ERROR_MAX_THRDS_REACHED:
		return "164"
	case ERROR_LOCK_FAILED:
		return "167"
	case ERROR_BUSY:
		return "170"
	case ERROR_CANCEL_VIOLATION:
		return "173"
	case ERROR_ATOMIC_LOCKS_NOT_SUPPORTED:
		return "174"
	case ERROR_INVALID_SEGMENT_NUMBER:
		return "180"
	case ERROR_INVALID_ORDINAL:
		return "182"
	case ERROR_ALREADY_EXISTS:
		return "183"
	case ERROR_INVALID_FLAG_NUMBER:
		return "186"
	case ERROR_SEM_NOT_FOUND:
		return "187"
	case ERROR_INVALID_STARTING_CODESEG:
		return "188"
	case ERROR_INVALID_STACKSEG:
		return "189"
	case ERROR_INVALID_MODULETYPE:
		return "190"
	case ERROR_INVALID_EXE_SIGNATURE:
		return "191"
	case ERROR_EXE_MARKED_INVALID:
		return "192"
	case ERROR_BAD_EXE_FORMAT:
		return "193"
	case ERROR_ITERATED_DATA_EXCEEDS_64k:
		return "194"
	case ERROR_INVALID_MINALLOCSIZE:
		return "195"
	case ERROR_DYNLINK_FROM_INVALID_RING:
		return "196"
	case ERROR_IOPL_NOT_ENABLED:
		return "197"
	case ERROR_INVALID_SEGDPL:
		return "198"
	case ERROR_AUTODATASEG_EXCEEDS_64k:
		return "199"
	case ERROR_RING2SEG_MUST_BE_MOVABLE:
		return "200"
	case ERROR_RELOC_CHAIN_XEEDS_SEGLIM:
		return "201"
	case ERROR_INFLOOP_IN_RELOC_CHAIN:
		return "202"
	case ERROR_ENVVAR_NOT_FOUND:
		return "203"
	case ERROR_NO_SIGNAL_SENT:
		return "205"
	case ERROR_FILENAME_EXCED_RANGE:
		return "206"
	case ERROR_RING2_STACK_IN_USE:
		return "207"
	case ERROR_META_EXPANSION_TOO_LONG:
		return "208"
	case ERROR_INVALID_SIGNAL_NUMBER:
		return "209"
	case ERROR_THREAD_1_INACTIVE:
		return "210"
	case ERROR_LOCKED:
		return "212"
	case ERROR_TOO_MANY_MODULES:
		return "214"
	case ERROR_NESTING_NOT_ALLOWED:
		return "215"
	case ERROR_EXE_MACHINE_TYPE_MISMATCH:
		return "216"
	case ERROR_EXE_CANNOT_MODIFY_SIGNED_BINARY:
		return "217"
	case ERROR_EXE_CANNOT_MODIFY_STRONG_SIGNED_BINARY:
		return "218"
	case ERROR_FILE_CHECKED_OUT:
		return "220"
	case ERROR_CHECKOUT_REQUIRED:
		return "221"
	case ERROR_BAD_FILE_TYPE:
		return "222"
	case ERROR_FILE_TOO_LARGE:
		return "223"
	case ERROR_FORMS_AUTH_REQUIRED:
		return "224"
	case ERROR_VIRUS_INFECTED:
		return "225"
	case ERROR_VIRUS_DELETED:
		return "226"
	case ERROR_PIPE_LOCAL:
		return "229"
	case ERROR_BAD_PIPE:
		return "230"
	case ERROR_PIPE_BUSY:
		return "231"
	case ERROR_NO_DATA:
		return "232"
	case ERROR_PIPE_NOT_CONNECTED:
		return "233"
	case ERROR_MORE_DATA:
		return "234"
	case ERROR_VC_DISCONNECTED:
		return "240"
	case ERROR_INVALID_EA_NAME:
		return "254"
	case ERROR_EA_LIST_INCONSISTENT:
		return "255"
	case WAIT_TIMEOUT:
		return "258"
	case ERROR_NO_MORE_ITEMS:
		return "259"
	case ERROR_CANNOT_COPY:
		return "266"
	case ERROR_DIRECTORY:
		return "267"
	case ERROR_EAS_DIDNT_FIT:
		return "275"
	case ERROR_EA_FILE_CORRUPT:
		return "276"
	case ERROR_EA_TABLE_FULL:
		return "277"
	case ERROR_INVALID_EA_HANDLE:
		return "278"
	case ERROR_EAS_NOT_SUPPORTED:
		return "282"
	case ERROR_NOT_OWNER:
		return "288"
	case ERROR_TOO_MANY_POSTS:
		return "298"
	case ERROR_PARTIAL_COPY:
		return "299"
	case ERROR_OPLOCK_NOT_GRANTED:
		return "300"
	case ERROR_INVALID_OPLOCK_PROTOCOL:
		return "301"
	case ERROR_DISK_TOO_FRAGMENTED:
		return "302"
	case ERROR_DELETE_PENDING:
		return "303"
	case ERROR_INCOMPATIBLE_WITH_GLOBAL_SHORT_NAME_REGISTRY_SETTING:
		return "304"
	case ERROR_SHORT_NAMES_NOT_ENABLED_ON_VOLUME:
		return "305"
	case ERROR_SECURITY_STREAM_IS_INCONSISTENT:
		return "306"
	case ERROR_INVALID_LOCK_RANGE:
		return "307"
	case ERROR_IMAGE_SUBSYSTEM_NOT_PRESENT:
		return "308"
	case ERROR_NOTIFICATION_GUID_ALREADY_DEFINED:
		return "309"
	case ERROR_MR_MID_NOT_FOUND:
		return "317"
	case ERROR_SCOPE_NOT_FOUND:
		return "318"
	case ERROR_FAIL_NOACTION_REBOOT:
		return "350"
	case ERROR_FAIL_SHUTDOWN:
		return "351"
	case ERROR_FAIL_RESTART:
		return "352"
	case ERROR_MAX_SESSIONS_REACHED:
		return "353"
	case ERROR_THREAD_MODE_ALREADY_BACKGROUND:
		return "400"
	case ERROR_THREAD_MODE_NOT_BACKGROUND:
		return "401"
	case ERROR_PROCESS_MODE_ALREADY_BACKGROUND:
		return "402"
	case ERROR_PROCESS_MODE_NOT_BACKGROUND:
		return "403"
	case ERROR_INVALID_ADDRESS:
		return "487"
	case ERROR_USER_PROFILE_LOAD:
		return "500"
	case ERROR_ARITHMETIC_OVERFLOW:
		return "534"
	case ERROR_PIPE_CONNECTED:
		return "535"
	case ERROR_PIPE_LISTENING:
		return "536"
	case ERROR_VERIFIER_STOP:
		return "537"
	case ERROR_ABIOS_ERROR:
		return "538"
	case ERROR_WX86_WARNING:
		return "539"
	case ERROR_WX86_ERROR:
		return "540"
	case ERROR_TIMER_NOT_CANCELED:
		return "541"
	case ERROR_UNWIND:
		return "542"
	case ERROR_BAD_STACK:
		return "543"
	case ERROR_INVALID_UNWIND_TARGET:
		return "544"
	case ERROR_INVALID_PORT_ATTRIBUTES:
		return "545"
	case ERROR_PORT_MESSAGE_TOO_LONG:
		return "546"
	case ERROR_INVALID_QUOTA_LOWER:
		return "547"
	case ERROR_DEVICE_ALREADY_ATTACHED:
		return "548"
	case ERROR_INSTRUCTION_MISALIGNMENT:
		return "549"
	case ERROR_PROFILING_NOT_STARTED:
		return "550"
	case ERROR_PROFILING_NOT_STOPPED:
		return "551"
	case ERROR_COULD_NOT_INTERPRET:
		return "552"
	case ERROR_PROFILING_AT_LIMIT:
		return "553"
	case ERROR_CANT_WAIT:
		return "554"
	case ERROR_CANT_TERMINATE_SELF:
		return "555"
	case ERROR_UNEXPECTED_MM_CREATE_ERR:
		return "556"
	case ERROR_UNEXPECTED_MM_MAP_ERROR:
		return "557"
	case ERROR_UNEXPECTED_MM_EXTEND_ERR:
		return "558"
	case ERROR_BAD_FUNCTION_TABLE:
		return "559"
	case ERROR_NO_GUID_TRANSLATION:
		return "560"
	case ERROR_INVALID_LDT_SIZE:
		return "561"
	case ERROR_INVALID_LDT_OFFSET:
		return "563"
	case ERROR_INVALID_LDT_DESCRIPTOR:
		return "564"
	case ERROR_TOO_MANY_THREADS:
		return "565"
	case ERROR_THREAD_NOT_IN_PROCESS:
		return "566"
	case ERROR_PAGEFILE_QUOTA_EXCEEDED:
		return "567"
	case ERROR_LOGON_SERVER_CONFLICT:
		return "568"
	case ERROR_SYNCHRONIZATION_REQUIRED:
		return "569"
	case ERROR_NET_OPEN_FAILED:
		return "570"
	case ERROR_IO_PRIVILEGE_FAILED:
		return "571"
	case ERROR_CONTROL_C_EXIT:
		return "572"
	case ERROR_MISSING_SYSTEMFILE:
		return "573"
	case ERROR_UNHANDLED_EXCEPTION:
		return "574"
	case ERROR_APP_INIT_FAILURE:
		return "575"
	case ERROR_PAGEFILE_CREATE_FAILED:
		return "576"
	case ERROR_INVALID_IMAGE_HASH:
		return "577"
	case ERROR_NO_PAGEFILE:
		return "578"
	case ERROR_ILLEGAL_FLOAT_CONTEXT:
		return "579"
	case ERROR_NO_EVENT_PAIR:
		return "580"
	case ERROR_DOMAIN_CTRLR_CONFIG_ERROR:
		return "581"
	case ERROR_ILLEGAL_CHARACTER:
		return "582"
	case ERROR_UNDEFINED_CHARACTER:
		return "583"
	case ERROR_FLOPPY_VOLUME:
		return "584"
	case ERROR_BIOS_FAILED_TO_CONNECT_INTERRUPT:
		return "585"
	case ERROR_BACKUP_CONTROLLER:
		return "586"
	case ERROR_MUTANT_LIMIT_EXCEEDED:
		return "587"
	case ERROR_FS_DRIVER_REQUIRED:
		return "588"
	case ERROR_CANNOT_LOAD_REGISTRY_FILE:
		return "589"
	case ERROR_DEBUG_ATTACH_FAILED:
		return "590"
	case ERROR_SYSTEM_PROCESS_TERMINATED:
		return "591"
	case ERROR_DATA_NOT_ACCEPTED:
		return "592"
	case ERROR_VDM_HARD_ERROR:
		return "593"
	case ERROR_DRIVER_CANCEL_TIMEOUT:
		return "594"
	case ERROR_REPLY_MESSAGE_MISMATCH:
		return "595"
	case ERROR_LOST_WRITEBEHIND_DATA:
		return "596"
	case ERROR_CLIENT_SERVER_PARAMETERS_INVALID:
		return "597"
	case ERROR_NOT_TINY_STREAM:
		return "598"
	case ERROR_STACK_OVERFLOW_READ:
		return "599"
	case ERROR_CONVERT_TO_LARGE:
		return "600"
	case ERROR_FOUND_OUT_OF_SCOPE:
		return "601"
	case ERROR_ALLOCATE_BUCKET:
		return "602"
	case ERROR_MARSHALL_OVERFLOW:
		return "603"
	case ERROR_INVALID_VARIANT:
		return "604"
	case ERROR_BAD_COMPRESSION_BUFFER:
		return "605"
	case ERROR_AUDIT_FAILED:
		return "606"
	case ERROR_TIMER_RESOLUTION_NOT_SET:
		return "607"
	case ERROR_INSUFFICIENT_LOGON_INFO:
		return "608"
	case ERROR_BAD_DLL_ENTRYPOINT:
		return "609"
	case ERROR_BAD_SERVICE_ENTRYPOINT:
		return "610"
	case ERROR_IP_ADDRESS_CONFLICT1:
		return "611"
	case ERROR_IP_ADDRESS_CONFLICT2:
		return "612"
	case ERROR_REGISTRY_QUOTA_LIMIT:
		return "613"
	case ERROR_NO_CALLBACK_ACTIVE:
		return "614"
	case ERROR_PWD_TOO_SHORT:
		return "615"
	case ERROR_PWD_TOO_RECENT:
		return "616"
	case ERROR_PWD_HISTORY_CONFLICT:
		return "617"
	case ERROR_UNSUPPORTED_COMPRESSION:
		return "618"
	case ERROR_INVALID_HW_PROFILE:
		return "619"
	case ERROR_INVALID_PLUGPLAY_DEVICE_PATH:
		return "620"
	case ERROR_QUOTA_LIST_INCONSISTENT:
		return "621"
	case ERROR_EVALUATION_EXPIRATION:
		return "622"
	case ERROR_ILLEGAL_DLL_RELOCATION:
		return "623"
	case ERROR_DLL_INIT_FAILED_LOGOFF:
		return "624"
	case ERROR_VALIDATE_CONTINUE:
		return "625"
	case ERROR_NO_MORE_MATCHES:
		return "626"
	case ERROR_RANGE_LIST_CONFLICT:
		return "627"
	case ERROR_SERVER_SID_MISMATCH:
		return "628"
	case ERROR_CANT_ENABLE_DENY_ONLY:
		return "629"
	case ERROR_FLOAT_MULTIPLE_FAULTS:
		return "630"
	case ERROR_FLOAT_MULTIPLE_TRAPS:
		return "631"
	case ERROR_NOINTERFACE:
		return "632"
	case ERROR_DRIVER_FAILED_SLEEP:
		return "633"
	case ERROR_CORRUPT_SYSTEM_FILE:
		return "634"
	case ERROR_COMMITMENT_MINIMUM:
		return "635"
	case ERROR_PNP_RESTART_ENUMERATION:
		return "636"
	case ERROR_SYSTEM_IMAGE_BAD_SIGNATURE:
		return "637"
	case ERROR_PNP_REBOOT_REQUIRED:
		return "638"
	case ERROR_INSUFFICIENT_POWER:
		return "639"
	case ERROR_MULTIPLE_FAULT_VIOLATION:
		return "640"
	case ERROR_SYSTEM_SHUTDOWN:
		return "641"
	case ERROR_PORT_NOT_SET:
		return "642"
	case ERROR_DS_VERSION_CHECK_FAILURE:
		return "643"
	case ERROR_RANGE_NOT_FOUND:
		return "644"
	case ERROR_NOT_SAFE_MODE_DRIVER:
		return "646"
	case ERROR_FAILED_DRIVER_ENTRY:
		return "647"
	case ERROR_DEVICE_ENUMERATION_ERROR:
		return "648"
	case ERROR_MOUNT_POINT_NOT_RESOLVED:
		return "649"
	case ERROR_INVALID_DEVICE_OBJECT_PARAMETER:
		return "650"
	case ERROR_MCA_OCCURED:
		return "651"
	case ERROR_DRIVER_DATABASE_ERROR:
		return "652"
	case ERROR_SYSTEM_HIVE_TOO_LARGE:
		return "653"
	case ERROR_DRIVER_FAILED_PRIOR_UNLOAD:
		return "654"
	case ERROR_VOLSNAP_PREPARE_HIBERNATE:
		return "655"
	case ERROR_HIBERNATION_FAILURE:
		return "656"
	case ERROR_FILE_SYSTEM_LIMITATION:
		return "665"
	case ERROR_ASSERTION_FAILURE:
		return "668"
	case ERROR_ACPI_ERROR:
		return "669"
	case ERROR_WOW_ASSERTION:
		return "670"
	case ERROR_PNP_BAD_MPS_TABLE:
		return "671"
	case ERROR_PNP_TRANSLATION_FAILED:
		return "672"
	case ERROR_PNP_IRQ_TRANSLATION_FAILED:
		return "673"
	case ERROR_PNP_INVALID_ID:
		return "674"
	case ERROR_WAKE_SYSTEM_DEBUGGER:
		return "675"
	case ERROR_HANDLES_CLOSED:
		return "676"
	case ERROR_EXTRANEOUS_INFORMATION:
		return "677"
	case ERROR_RXACT_COMMIT_NECESSARY:
		return "678"
	case ERROR_MEDIA_CHECK:
		return "679"
	case ERROR_GUID_SUBSTITUTION_MADE:
		return "680"
	case ERROR_STOPPED_ON_SYMLINK:
		return "681"
	case ERROR_LONGJUMP:
		return "682"
	case ERROR_PLUGPLAY_QUERY_VETOED:
		return "683"
	case ERROR_UNWIND_CONSOLIDATE:
		return "684"
	case ERROR_REGISTRY_HIVE_RECOVERED:
		return "685"
	case ERROR_DLL_MIGHT_BE_INSECURE:
		return "686"
	case ERROR_DLL_MIGHT_BE_INCOMPATIBLE:
		return "687"
	case ERROR_DBG_EXCEPTION_NOT_HANDLED:
		return "688"
	case ERROR_DBG_REPLY_LATER:
		return "689"
	case ERROR_DBG_UNABLE_TO_PROVIDE_HANDLE:
		return "690"
	case ERROR_DBG_TERMINATE_THREAD:
		return "691"
	case ERROR_DBG_TERMINATE_PROCESS:
		return "692"
	case ERROR_DBG_CONTROL_C:
		return "693"
	case ERROR_DBG_PRINTEXCEPTION_C:
		return "694"
	case ERROR_DBG_RIPEXCEPTION:
		return "695"
	case ERROR_DBG_CONTROL_BREAK:
		return "696"
	case ERROR_DBG_COMMAND_EXCEPTION:
		return "697"
	case ERROR_OBJECT_NAME_EXISTS:
		return "698"
	case ERROR_THREAD_WAS_SUSPENDED:
		return "699"
	case ERROR_IMAGE_NOT_AT_BASE:
		return "700"
	case ERROR_RXACT_STATE_CREATED:
		return "701"
	case ERROR_SEGMENT_NOTIFICATION:
		return "702"
	case ERROR_BAD_CURRENT_DIRECTORY:
		return "703"
	case ERROR_FT_READ_RECOVERY_FROM_BACKUP:
		return "704"
	case ERROR_FT_WRITE_RECOVERY:
		return "705"
	case ERROR_IMAGE_MACHINE_TYPE_MISMATCH:
		return "706"
	case ERROR_RECEIVE_PARTIAL:
		return "707"
	case ERROR_RECEIVE_EXPEDITED:
		return "708"
	case ERROR_RECEIVE_PARTIAL_EXPEDITED:
		return "709"
	case ERROR_EVENT_DONE:
		return "710"
	case ERROR_EVENT_PENDING:
		return "711"
	case ERROR_CHECKING_FILE_SYSTEM:
		return "712"
	case ERROR_FATAL_APP_EXIT:
		return "713"
	case ERROR_PREDEFINED_HANDLE:
		return "714"
	case ERROR_WAS_UNLOCKED:
		return "715"
	case ERROR_SERVICE_NOTIFICATION:
		return "716"
	case ERROR_WAS_LOCKED:
		return "717"
	case ERROR_LOG_HARD_ERROR:
		return "718"
	case ERROR_ALREADY_WIN32:
		return "719"
	case ERROR_IMAGE_MACHINE_TYPE_MISMATCH_EXE:
		return "720"
	case ERROR_NO_YIELD_PERFORMED:
		return "721"
	case ERROR_TIMER_RESUME_IGNORED:
		return "722"
	case ERROR_ARBITRATION_UNHANDLED:
		return "723"
	case ERROR_CARDBUS_NOT_SUPPORTED:
		return "724"
	case ERROR_MP_PROCESSOR_MISMATCH:
		return "725"
	case ERROR_HIBERNATED:
		return "726"
	case ERROR_RESUME_HIBERNATION:
		return "727"
	case ERROR_FIRMWARE_UPDATED:
		return "728"
	case ERROR_DRIVERS_LEAKING_LOCKED_PAGES:
		return "729"
	case ERROR_WAKE_SYSTEM:
		return "730"
	case ERROR_WAIT_1:
		return "731"
	case ERROR_WAIT_2:
		return "732"
	case ERROR_WAIT_3:
		return "733"
	case ERROR_WAIT_63:
		return "734"
	case ERROR_ABANDONED_WAIT_0:
		return "735"
	case ERROR_ABANDONED_WAIT_63:
		return "736"
	case ERROR_USER_APC:
		return "737"
	case ERROR_KERNEL_APC:
		return "738"
	case ERROR_ALERTED:
		return "739"
	case ERROR_ELEVATION_REQUIRED:
		return "740"
	case ERROR_REPARSE:
		return "741"
	case ERROR_OPLOCK_BREAK_IN_PROGRESS:
		return "742"
	case ERROR_VOLUME_MOUNTED:
		return "743"
	case ERROR_RXACT_COMMITTED:
		return "744"
	case ERROR_NOTIFY_CLEANUP:
		return "745"
	case ERROR_PRIMARY_TRANSPORT_CONNECT_FAILED:
		return "746"
	case ERROR_PAGE_FAULT_TRANSITION:
		return "747"
	case ERROR_PAGE_FAULT_DEMAND_ZERO:
		return "748"
	case ERROR_PAGE_FAULT_COPY_ON_WRITE:
		return "749"
	case ERROR_PAGE_FAULT_GUARD_PAGE:
		return "750"
	case ERROR_PAGE_FAULT_PAGING_FILE:
		return "751"
	case ERROR_CACHE_PAGE_LOCKED:
		return "752"
	case ERROR_CRASH_DUMP:
		return "753"
	case ERROR_BUFFER_ALL_ZEROS:
		return "754"
	case ERROR_REPARSE_OBJECT:
		return "755"
	case ERROR_RESOURCE_REQUIREMENTS_CHANGED:
		return "756"
	case ERROR_TRANSLATION_COMPLETE:
		return "757"
	case ERROR_NOTHING_TO_TERMINATE:
		return "758"
	case ERROR_PROCESS_NOT_IN_JOB:
		return "759"
	case ERROR_PROCESS_IN_JOB:
		return "760"
	case ERROR_VOLSNAP_HIBERNATE_READY:
		return "761"
	case ERROR_FSFILTER_OP_COMPLETED_SUCCESSFULLY:
		return "762"
	case ERROR_INTERRUPT_VECTOR_ALREADY_CONNECTED:
		return "763"
	case ERROR_INTERRUPT_STILL_CONNECTED:
		return "764"
	case ERROR_WAIT_FOR_OPLOCK:
		return "765"
	case ERROR_DBG_EXCEPTION_HANDLED:
		return "766"
	case ERROR_DBG_CONTINUE:
		return "767"
	case ERROR_CALLBACK_POP_STACK:
		return "768"
	case ERROR_COMPRESSION_DISABLED:
		return "769"
	case ERROR_CANTFETCHBACKWARDS:
		return "770"
	case ERROR_CANTSCROLLBACKWARDS:
		return "771"
	case ERROR_ROWSNOTRELEASED:
		return "772"
	case ERROR_BAD_ACCESSOR_FLAGS:
		return "773"
	case ERROR_ERRORS_ENCOUNTERED:
		return "774"
	case ERROR_NOT_CAPABLE:
		return "775"
	case ERROR_REQUEST_OUT_OF_SEQUENCE:
		return "776"
	case ERROR_VERSION_PARSE_ERROR:
		return "777"
	case ERROR_BADSTARTPOSITION:
		return "778"
	case ERROR_MEMORY_HARDWARE:
		return "779"
	case ERROR_DISK_REPAIR_DISABLED:
		return "780"
	case ERROR_INSUFFICIENT_RESOURCE_FOR_SPECIFIED_SHARED_SECTION_SIZE:
		return "781"
	case ERROR_SYSTEM_POWERSTATE_TRANSITION:
		return "782"
	case ERROR_SYSTEM_POWERSTATE_COMPLEX_TRANSITION:
		return "783"
	case ERROR_MCA_EXCEPTION:
		return "784"
	case ERROR_ACCESS_AUDIT_BY_POLICY:
		return "785"
	case ERROR_ACCESS_DISABLED_NO_SAFER_UI_BY_POLICY:
		return "786"
	case ERROR_ABANDON_HIBERFILE:
		return "787"
	case ERROR_LOST_WRITEBEHIND_DATA_NETWORK_DISCONNECTED:
		return "788"
	case ERROR_LOST_WRITEBEHIND_DATA_NETWORK_SERVER_ERROR:
		return "789"
	case ERROR_LOST_WRITEBEHIND_DATA_LOCAL_DISK_ERROR:
		return "790"
	case ERROR_BAD_MCFG_TABLE:
		return "791"
	case ERROR_OPLOCK_SWITCHED_TO_NEW_HANDLE:
		return "800"
	case ERROR_CANNOT_GRANT_REQUESTED_OPLOCK:
		return "801"
	case ERROR_CANNOT_BREAK_OPLOCK:
		return "802"
	case ERROR_OPLOCK_HANDLE_CLOSED:
		return "803"
	case ERROR_NO_ACE_CONDITION:
		return "804"
	case ERROR_INVALID_ACE_CONDITION:
		return "805"
	case ERROR_EA_ACCESS_DENIED:
		return "994"
	case ERROR_OPERATION_ABORTED:
		return "995"
	case ERROR_IO_INCOMPLETE:
		return "996"
	case ERROR_IO_PENDING:
		return "997"
	case ERROR_NOACCESS:
		return "998"
	case ERROR_SWAPERROR:
		return "999"
	case ERROR_STACK_OVERFLOW:
		return "1001"
	case ERROR_INVALID_MESSAGE:
		return "1002"
	case ERROR_CAN_NOT_COMPLETE:
		return "1003"
	case ERROR_INVALID_FLAGS:
		return "1004"
	case ERROR_UNRECOGNIZED_VOLUME:
		return "1005"
	case ERROR_FILE_INVALID:
		return "1006"
	case ERROR_FULLSCREEN_MODE:
		return "1007"
	case ERROR_NO_TOKEN:
		return "1008"
	case ERROR_BADDB:
		return "1009"
	case ERROR_BADKEY:
		return "1010"
	case ERROR_CANTOPEN:
		return "1011"
	case ERROR_CANTREAD:
		return "1012"
	case ERROR_CANTWRITE:
		return "1013"
	case ERROR_REGISTRY_RECOVERED:
		return "1014"
	case ERROR_REGISTRY_CORRUPT:
		return "1015"
	case ERROR_REGISTRY_IO_FAILED:
		return "1016"
	case ERROR_NOT_REGISTRY_FILE:
		return "1017"
	case ERROR_KEY_DELETED:
		return "1018"
	case ERROR_NO_LOG_SPACE:
		return "1019"
	case ERROR_KEY_HAS_CHILDREN:
		return "1020"
	case ERROR_CHILD_MUST_BE_VOLATILE:
		return "1021"
	case ERROR_NOTIFY_ENUM_DIR:
		return "1022"
	case ERROR_DEPENDENT_SERVICES_RUNNING:
		return "1051"
	case ERROR_INVALID_SERVICE_CONTROL:
		return "1052"
	case ERROR_SERVICE_REQUEST_TIMEOUT:
		return "1053"
	case ERROR_SERVICE_NO_THREAD:
		return "1054"
	case ERROR_SERVICE_DATABASE_LOCKED:
		return "1055"
	case ERROR_SERVICE_ALREADY_RUNNING:
		return "1056"
	case ERROR_INVALID_SERVICE_ACCOUNT:
		return "1057"
	case ERROR_SERVICE_DISABLED:
		return "1058"
	case ERROR_CIRCULAR_DEPENDENCY:
		return "1059"
	case ERROR_SERVICE_DOES_NOT_EXIST:
		return "1060"
	case ERROR_SERVICE_CANNOT_ACCEPT_CTRL:
		return "1061"
	case ERROR_SERVICE_NOT_ACTIVE:
		return "1062"
	case ERROR_FAILED_SERVICE_CONTROLLER_CONNECT:
		return "1063"
	case ERROR_EXCEPTION_IN_SERVICE:
		return "1064"
	case ERROR_DATABASE_DOES_NOT_EXIST:
		return "1065"
	case ERROR_SERVICE_SPECIFIC_ERROR:
		return "1066"
	case ERROR_PROCESS_ABORTED:
		return "1067"
	case ERROR_SERVICE_DEPENDENCY_FAIL:
		return "1068"
	case ERROR_SERVICE_LOGON_FAILED:
		return "1069"
	case ERROR_SERVICE_START_HANG:
		return "1070"
	case ERROR_INVALID_SERVICE_LOCK:
		return "1071"
	case ERROR_SERVICE_MARKED_FOR_DELETE:
		return "1072"
	case ERROR_SERVICE_EXISTS:
		return "1073"
	case ERROR_ALREADY_RUNNING_LKG:
		return "1074"
	case ERROR_SERVICE_DEPENDENCY_DELETED:
		return "1075"
	case ERROR_BOOT_ALREADY_ACCEPTED:
		return "1076"
	case ERROR_SERVICE_NEVER_STARTED:
		return "1077"
	case ERROR_DUPLICATE_SERVICE_NAME:
		return "1078"
	case ERROR_DIFFERENT_SERVICE_ACCOUNT:
		return "1079"
	case ERROR_CANNOT_DETECT_DRIVER_FAILURE:
		return "1080"
	case ERROR_CANNOT_DETECT_PROCESS_ABORT:
		return "1081"
	case ERROR_NO_RECOVERY_PROGRAM:
		return "1082"
	case ERROR_SERVICE_NOT_IN_EXE:
		return "1083"
	case ERROR_NOT_SAFEBOOT_SERVICE:
		return "1084"
	case ERROR_END_OF_MEDIA:
		return "1100"
	case ERROR_FILEMARK_DETECTED:
		return "1101"
	case ERROR_BEGINNING_OF_MEDIA:
		return "1102"
	case ERROR_SETMARK_DETECTED:
		return "1103"
	case ERROR_NO_DATA_DETECTED:
		return "1104"
	case ERROR_PARTITION_FAILURE:
		return "1105"
	case ERROR_INVALID_BLOCK_LENGTH:
		return "1106"
	case ERROR_DEVICE_NOT_PARTITIONED:
		return "1107"
	case ERROR_UNABLE_TO_LOCK_MEDIA:
		return "1108"
	case ERROR_UNABLE_TO_UNLOAD_MEDIA:
		return "1109"
	case ERROR_MEDIA_CHANGED:
		return "1110"
	case ERROR_BUS_RESET:
		return "1111"
	case ERROR_NO_MEDIA_IN_DRIVE:
		return "1112"
	case ERROR_NO_UNICODE_TRANSLATION:
		return "1113"
	case ERROR_DLL_INIT_FAILED:
		return "1114"
	case ERROR_SHUTDOWN_IN_PROGRESS:
		return "1115"
	case ERROR_NO_SHUTDOWN_IN_PROGRESS:
		return "1116"
	case ERROR_IO_DEVICE:
		return "1117"
	case ERROR_SERIAL_NO_DEVICE:
		return "1118"
	case ERROR_IRQ_BUSY:
		return "1119"
	case ERROR_MORE_WRITES:
		return "1120"
	case ERROR_COUNTER_TIMEOUT:
		return "1121"
	case ERROR_FLOPPY_ID_MARK_NOT_FOUND:
		return "1122"
	case ERROR_FLOPPY_WRONG_CYLINDER:
		return "1123"
	case ERROR_FLOPPY_UNKNOWN_ERROR:
		return "1124"
	case ERROR_FLOPPY_BAD_REGISTERS:
		return "1125"
	case ERROR_DISK_RECALIBRATE_FAILED:
		return "1126"
	case ERROR_DISK_OPERATION_FAILED:
		return "1127"
	case ERROR_DISK_RESET_FAILED:
		return "1128"
	case ERROR_EOM_OVERFLOW:
		return "1129"
	case ERROR_NOT_ENOUGH_SERVER_MEMORY:
		return "1130"
	case ERROR_POSSIBLE_DEADLOCK:
		return "1131"
	case ERROR_MAPPED_ALIGNMENT:
		return "1132"
	case ERROR_SET_POWER_STATE_VETOED:
		return "1140"
	case ERROR_SET_POWER_STATE_FAILED:
		return "1141"
	case ERROR_TOO_MANY_LINKS:
		return "1142"
	case ERROR_OLD_WIN_VERSION:
		return "1150"
	case ERROR_APP_WRONG_OS:
		return "1151"
	case ERROR_SINGLE_INSTANCE_APP:
		return "1152"
	case ERROR_RMODE_APP:
		return "1153"
	case ERROR_INVALID_DLL:
		return "1154"
	case ERROR_NO_ASSOCIATION:
		return "1155"
	case ERROR_DDE_FAIL:
		return "1156"
	case ERROR_DLL_NOT_FOUND:
		return "1157"
	case ERROR_NO_MORE_USER_HANDLES:
		return "1158"
	case ERROR_MESSAGE_SYNC_ONLY:
		return "1159"
	case ERROR_SOURCE_ELEMENT_EMPTY:
		return "1160"
	case ERROR_DESTINATION_ELEMENT_FULL:
		return "1161"
	case ERROR_ILLEGAL_ELEMENT_ADDRESS:
		return "1162"
	case ERROR_MAGAZINE_NOT_PRESENT:
		return "1163"
	case ERROR_DEVICE_REINITIALIZATION_NEEDED:
		return "1164"
	case ERROR_DEVICE_REQUIRES_CLEANING:
		return "1165"
	case ERROR_DEVICE_DOOR_OPEN:
		return "1166"
	case ERROR_DEVICE_NOT_CONNECTED:
		return "1167"
	case ERROR_NOT_FOUND:
		return "1168"
	case ERROR_NO_MATCH:
		return "1169"
	case ERROR_SET_NOT_FOUND:
		return "1170"
	case ERROR_POINT_NOT_FOUND:
		return "1171"
	case ERROR_NO_TRACKING_SERVICE:
		return "1172"
	case ERROR_NO_VOLUME_ID:
		return "1173"
	case ERROR_UNABLE_TO_REMOVE_REPLACED:
		return "1175"
	case ERROR_UNABLE_TO_MOVE_REPLACEMENT:
		return "1176"
	case ERROR_UNABLE_TO_MOVE_REPLACEMENT_2:
		return "1177"
	case ERROR_JOURNAL_DELETE_IN_PROGRESS:
		return "1178"
	case ERROR_JOURNAL_NOT_ACTIVE:
		return "1179"
	case ERROR_POTENTIAL_FILE_FOUND:
		return "1180"
	case ERROR_JOURNAL_ENTRY_DELETED:
		return "1181"
	case ERROR_SHUTDOWN_IS_SCHEDULED:
		return "1190"
	case ERROR_SHUTDOWN_USERS_LOGGED_ON:
		return "1191"
	case ERROR_BAD_DEVICE:
		return "1200"
	case ERROR_CONNECTION_UNAVAIL:
		return "1201"
	case ERROR_DEVICE_ALREADY_REMEMBERED:
		return "1202"
	case ERROR_NO_NET_OR_BAD_PATH:
		return "1203"
	case ERROR_BAD_PROVIDER:
		return "1204"
	case ERROR_CANNOT_OPEN_PROFILE:
		return "1205"
	case ERROR_BAD_PROFILE:
		return "1206"
	case ERROR_NOT_CONTAINER:
		return "1207"
	case ERROR_EXTENDED_ERROR:
		return "1208"
	case ERROR_INVALID_GROUPNAME:
		return "1209"
	case ERROR_INVALID_COMPUTERNAME:
		return "1210"
	case ERROR_INVALID_EVENTNAME:
		return "1211"
	case ERROR_INVALID_DOMAINNAME:
		return "1212"
	case ERROR_INVALID_SERVICENAME:
		return "1213"
	case ERROR_INVALID_NETNAME:
		return "1214"
	case ERROR_INVALID_SHARENAME:
		return "1215"
	case ERROR_INVALID_PASSWORDNAME:
		return "1216"
	case ERROR_INVALID_MESSAGENAME:
		return "1217"
	case ERROR_INVALID_MESSAGEDEST:
		return "1218"
	case ERROR_SESSION_CREDENTIAL_CONFLICT:
		return "1219"
	case ERROR_REMOTE_SESSION_LIMIT_EXCEEDED:
		return "1220"
	case ERROR_DUP_DOMAINNAME:
		return "1221"
	case ERROR_NO_NETWORK:
		return "1222"
	case ERROR_CANCELLED:
		return "1223"
	case ERROR_USER_MAPPED_FILE:
		return "1224"
	case ERROR_CONNECTION_REFUSED:
		return "1225"
	case ERROR_GRACEFUL_DISCONNECT:
		return "1226"
	case ERROR_ADDRESS_ALREADY_ASSOCIATED:
		return "1227"
	case ERROR_ADDRESS_NOT_ASSOCIATED:
		return "1228"
	case ERROR_CONNECTION_INVALID:
		return "1229"
	case ERROR_CONNECTION_ACTIVE:
		return "1230"
	case ERROR_NETWORK_UNREACHABLE:
		return "1231"
	case ERROR_HOST_UNREACHABLE:
		return "1232"
	case ERROR_PROTOCOL_UNREACHABLE:
		return "1233"
	case ERROR_PORT_UNREACHABLE:
		return "1234"
	case ERROR_REQUEST_ABORTED:
		return "1235"
	case ERROR_CONNECTION_ABORTED:
		return "1236"
	case ERROR_RETRY:
		return "1237"
	case ERROR_CONNECTION_COUNT_LIMIT:
		return "1238"
	case ERROR_LOGIN_TIME_RESTRICTION:
		return "1239"
	case ERROR_LOGIN_WKSTA_RESTRICTION:
		return "1240"
	case ERROR_INCORRECT_ADDRESS:
		return "1241"
	case ERROR_ALREADY_REGISTERED:
		return "1242"
	case ERROR_SERVICE_NOT_FOUND:
		return "1243"
	case ERROR_NOT_AUTHENTICATED:
		return "1244"
	case ERROR_NOT_LOGGED_ON:
		return "1245"
	case ERROR_CONTINUE:
		return "1246"
	case ERROR_ALREADY_INITIALIZED:
		return "1247"
	case ERROR_NO_MORE_DEVICES:
		return "1248"
	case ERROR_NO_SUCH_SITE:
		return "1249"
	case ERROR_DOMAIN_CONTROLLER_EXISTS:
		return "1250"
	case ERROR_ONLY_IF_CONNECTED:
		return "1251"
	case ERROR_OVERRIDE_NOCHANGES:
		return "1252"
	case ERROR_BAD_USER_PROFILE:
		return "1253"
	case ERROR_NOT_SUPPORTED_ON_SBS:
		return "1254"
	case ERROR_SERVER_SHUTDOWN_IN_PROGRESS:
		return "1255"
	case ERROR_HOST_DOWN:
		return "1256"
	case ERROR_NON_ACCOUNT_SID:
		return "1257"
	case ERROR_NON_DOMAIN_SID:
		return "1258"
	case ERROR_APPHELP_BLOCK:
		return "1259"
	case ERROR_ACCESS_DISABLED_BY_POLICY:
		return "1260"
	case ERROR_REG_NAT_CONSUMPTION:
		return "1261"
	case ERROR_CSCSHARE_OFFLINE:
		return "1262"
	case ERROR_PKINIT_FAILURE:
		return "1263"
	case ERROR_SMARTCARD_SUBSYSTEM_FAILURE:
		return "1264"
	case ERROR_DOWNGRADE_DETECTED:
		return "1265"
	case ERROR_MACHINE_LOCKED:
		return "1271"
	case ERROR_CALLBACK_SUPPLIED_INVALID_DATA:
		return "1273"
	case ERROR_SYNC_FOREGROUND_REFRESH_REQUIRED:
		return "1274"
	case ERROR_DRIVER_BLOCKED:
		return "1275"
	case ERROR_INVALID_IMPORT_OF_NON_DLL:
		return "1276"
	case ERROR_ACCESS_DISABLED_WEBBLADE:
		return "1277"
	case ERROR_ACCESS_DISABLED_WEBBLADE_TAMPER:
		return "1278"
	case ERROR_RECOVERY_FAILURE:
		return "1279"
	case ERROR_ALREADY_FIBER:
		return "1280"
	case ERROR_ALREADY_THREAD:
		return "1281"
	case ERROR_STACK_BUFFER_OVERRUN:
		return "1282"
	case ERROR_PARAMETER_QUOTA_EXCEEDED:
		return "1283"
	case ERROR_DEBUGGER_INACTIVE:
		return "1284"
	case ERROR_DELAY_LOAD_FAILED:
		return "1285"
	case ERROR_VDM_DISALLOWED:
		return "1286"
	case ERROR_UNIDENTIFIED_ERROR:
		return "1287"
	case ERROR_INVALID_CRUNTIME_PARAMETER:
		return "1288"
	case ERROR_BEYOND_VDL:
		return "1289"
	case ERROR_INCOMPATIBLE_SERVICE_SID_TYPE:
		return "1290"
	case ERROR_DRIVER_PROCESS_TERMINATED:
		return "1291"
	case ERROR_IMPLEMENTATION_LIMIT:
		return "1292"
	case ERROR_PROCESS_IS_PROTECTED:
		return "1293"
	case ERROR_SERVICE_NOTIFY_CLIENT_LAGGING:
		return "1294"
	case ERROR_DISK_QUOTA_EXCEEDED:
		return "1295"
	case ERROR_CONTENT_BLOCKED:
		return "1296"
	case ERROR_INCOMPATIBLE_SERVICE_PRIVILEGE:
		return "1297"
	case ERROR_APP_HANG:
		return "1298"
	case ERROR_INVALID_LABEL:
		return "1299"
	case ERROR_NOT_ALL_ASSIGNED:
		return "1300"
	case ERROR_SOME_NOT_MAPPED:
		return "1301"
	case ERROR_NO_QUOTAS_FOR_ACCOUNT:
		return "1302"
	case ERROR_LOCAL_USER_SESSION_KEY:
		return "1303"
	case ERROR_NULL_LM_PASSWORD:
		return "1304"
	case ERROR_UNKNOWN_REVISION:
		return "1305"
	case ERROR_REVISION_MISMATCH:
		return "1306"
	case ERROR_INVALID_OWNER:
		return "1307"
	case ERROR_INVALID_PRIMARY_GROUP:
		return "1308"
	case ERROR_NO_IMPERSONATION_TOKEN:
		return "1309"
	case ERROR_CANT_DISABLE_MANDATORY:
		return "1310"
	case ERROR_NO_LOGON_SERVERS:
		return "1311"
	case ERROR_NO_SUCH_LOGON_SESSION:
		return "1312"
	case ERROR_NO_SUCH_PRIVILEGE:
		return "1313"
	case ERROR_PRIVILEGE_NOT_HELD:
		return "1314"
	case ERROR_INVALID_ACCOUNT_NAME:
		return "1315"
	case ERROR_USER_EXISTS:
		return "1316"
	case ERROR_NO_SUCH_USER:
		return "1317"
	case ERROR_GROUP_EXISTS:
		return "1318"
	case ERROR_NO_SUCH_GROUP:
		return "1319"
	case ERROR_MEMBER_IN_GROUP:
		return "1320"
	case ERROR_MEMBER_NOT_IN_GROUP:
		return "1321"
	case ERROR_LAST_ADMIN:
		return "1322"
	case ERROR_WRONG_PASSWORD:
		return "1323"
	case ERROR_ILL_FORMED_PASSWORD:
		return "1324"
	case ERROR_PASSWORD_RESTRICTION:
		return "1325"
	case ERROR_LOGON_FAILURE:
		return "1326"
	case ERROR_ACCOUNT_RESTRICTION:
		return "1327"
	case ERROR_INVALID_LOGON_HOURS:
		return "1328"
	case ERROR_INVALID_WORKSTATION:
		return "1329"
	case ERROR_PASSWORD_EXPIRED:
		return "1330"
	case ERROR_ACCOUNT_DISABLED:
		return "1331"
	case ERROR_NONE_MAPPED:
		return "1332"
	case ERROR_TOO_MANY_LUIDS_REQUESTED:
		return "1333"
	case ERROR_LUIDS_EXHAUSTED:
		return "1334"
	case ERROR_INVALID_SUB_AUTHORITY:
		return "1335"
	case ERROR_INVALID_ACL:
		return "1336"
	case ERROR_INVALID_SID:
		return "1337"
	case ERROR_INVALID_SECURITY_DESCR:
		return "1338"
	case ERROR_BAD_INHERITANCE_ACL:
		return "1340"
	case ERROR_SERVER_DISABLED:
		return "1341"
	case ERROR_SERVER_NOT_DISABLED:
		return "1342"
	case ERROR_INVALID_ID_AUTHORITY:
		return "1343"
	case ERROR_ALLOTTED_SPACE_EXCEEDED:
		return "1344"
	case ERROR_INVALID_GROUP_ATTRIBUTES:
		return "1345"
	case ERROR_BAD_IMPERSONATION_LEVEL:
		return "1346"
	case ERROR_CANT_OPEN_ANONYMOUS:
		return "1347"
	case ERROR_BAD_VALIDATION_CLASS:
		return "1348"
	case ERROR_BAD_TOKEN_TYPE:
		return "1349"
	case ERROR_NO_SECURITY_ON_OBJECT:
		return "1350"
	case ERROR_CANT_ACCESS_DOMAIN_INFO:
		return "1351"
	case ERROR_INVALID_SERVER_STATE:
		return "1352"
	case ERROR_INVALID_DOMAIN_STATE:
		return "1353"
	case ERROR_INVALID_DOMAIN_ROLE:
		return "1354"
	case ERROR_NO_SUCH_DOMAIN:
		return "1355"
	case ERROR_DOMAIN_EXISTS:
		return "1356"
	case ERROR_DOMAIN_LIMIT_EXCEEDED:
		return "1357"
	case ERROR_INTERNAL_DB_CORRUPTION:
		return "1358"
	case ERROR_INTERNAL_ERROR:
		return "1359"
	case ERROR_GENERIC_NOT_MAPPED:
		return "1360"
	case ERROR_BAD_DESCRIPTOR_FORMAT:
		return "1361"
	case ERROR_NOT_LOGON_PROCESS:
		return "1362"
	case ERROR_LOGON_SESSION_EXISTS:
		return "1363"
	case ERROR_NO_SUCH_PACKAGE:
		return "1364"
	case ERROR_BAD_LOGON_SESSION_STATE:
		return "1365"
	case ERROR_LOGON_SESSION_COLLISION:
		return "1366"
	case ERROR_INVALID_LOGON_TYPE:
		return "1367"
	case ERROR_CANNOT_IMPERSONATE:
		return "1368"
	case ERROR_RXACT_INVALID_STATE:
		return "1369"
	case ERROR_RXACT_COMMIT_FAILURE:
		return "1370"
	case ERROR_SPECIAL_ACCOUNT:
		return "1371"
	case ERROR_SPECIAL_GROUP:
		return "1372"
	case ERROR_SPECIAL_USER:
		return "1373"
	case ERROR_MEMBERS_PRIMARY_GROUP:
		return "1374"
	case ERROR_TOKEN_ALREADY_IN_USE:
		return "1375"
	case ERROR_NO_SUCH_ALIAS:
		return "1376"
	case ERROR_MEMBER_NOT_IN_ALIAS:
		return "1377"
	case ERROR_MEMBER_IN_ALIAS:
		return "1378"
	case ERROR_ALIAS_EXISTS:
		return "1379"
	case ERROR_LOGON_NOT_GRANTED:
		return "1380"
	case ERROR_TOO_MANY_SECRETS:
		return "1381"
	case ERROR_SECRET_TOO_LONG:
		return "1382"
	case ERROR_INTERNAL_DB_ERROR:
		return "1383"
	case ERROR_TOO_MANY_CONTEXT_IDS:
		return "1384"
	case ERROR_LOGON_TYPE_NOT_GRANTED:
		return "1385"
	case ERROR_NT_CROSS_ENCRYPTION_REQUIRED:
		return "1386"
	case ERROR_NO_SUCH_MEMBER:
		return "1387"
	case ERROR_INVALID_MEMBER:
		return "1388"
	case ERROR_TOO_MANY_SIDS:
		return "1389"
	case ERROR_LM_CROSS_ENCRYPTION_REQUIRED:
		return "1390"
	case ERROR_NO_INHERITANCE:
		return "1391"
	case ERROR_FILE_CORRUPT:
		return "1392"
	case ERROR_DISK_CORRUPT:
		return "1393"
	case ERROR_NO_USER_SESSION_KEY:
		return "1394"
	case ERROR_LICENSE_QUOTA_EXCEEDED:
		return "1395"
	case ERROR_WRONG_TARGET_NAME:
		return "1396"
	case ERROR_MUTUAL_AUTH_FAILED:
		return "1397"
	case ERROR_TIME_SKEW:
		return "1398"
	case ERROR_CURRENT_DOMAIN_NOT_ALLOWED:
		return "1399"
	case ERROR_INVALID_WINDOW_HANDLE:
		return "1400"
	case ERROR_INVALID_MENU_HANDLE:
		return "1401"
	case ERROR_INVALID_CURSOR_HANDLE:
		return "1402"
	case ERROR_INVALID_ACCEL_HANDLE:
		return "1403"
	case ERROR_INVALID_HOOK_HANDLE:
		return "1404"
	case ERROR_INVALID_DWP_HANDLE:
		return "1405"
	case ERROR_TLW_WITH_WSCHILD:
		return "1406"
	case ERROR_CANNOT_FIND_WND_CLASS:
		return "1407"
	case ERROR_WINDOW_OF_OTHER_THREAD:
		return "1408"
	case ERROR_HOTKEY_ALREADY_REGISTERED:
		return "1409"
	case ERROR_CLASS_ALREADY_EXISTS:
		return "1410"
	case ERROR_CLASS_DOES_NOT_EXIST:
		return "1411"
	case ERROR_CLASS_HAS_WINDOWS:
		return "1412"
	case ERROR_INVALID_INDEX:
		return "1413"
	case ERROR_INVALID_ICON_HANDLE:
		return "1414"
	case ERROR_PRIVATE_DIALOG_INDEX:
		return "1415"
	case ERROR_LISTBOX_ID_NOT_FOUND:
		return "1416"
	case ERROR_NO_WILDCARD_CHARACTERS:
		return "1417"
	case ERROR_CLIPBOARD_NOT_OPEN:
		return "1418"
	case ERROR_HOTKEY_NOT_REGISTERED:
		return "1419"
	case ERROR_WINDOW_NOT_DIALOG:
		return "1420"
	case ERROR_CONTROL_ID_NOT_FOUND:
		return "1421"
	case ERROR_INVALID_COMBOBOX_MESSAGE:
		return "1422"
	case ERROR_WINDOW_NOT_COMBOBOX:
		return "1423"
	case ERROR_INVALID_EDIT_HEIGHT:
		return "1424"
	case ERROR_DC_NOT_FOUND:
		return "1425"
	case ERROR_INVALID_HOOK_FILTER:
		return "1426"
	case ERROR_INVALID_FILTER_PROC:
		return "1427"
	case ERROR_HOOK_NEEDS_HMOD:
		return "1428"
	case ERROR_GLOBAL_ONLY_HOOK:
		return "1429"
	case ERROR_JOURNAL_HOOK_SET:
		return "1430"
	case ERROR_HOOK_NOT_INSTALLED:
		return "1431"
	case ERROR_INVALID_LB_MESSAGE:
		return "1432"
	case ERROR_SETCOUNT_ON_BAD_LB:
		return "1433"
	case ERROR_LB_WITHOUT_TABSTOPS:
		return "1434"
	case ERROR_DESTROY_OBJECT_OF_OTHER_THREAD:
		return "1435"
	case ERROR_CHILD_WINDOW_MENU:
		return "1436"
	case ERROR_NO_SYSTEM_MENU:
		return "1437"
	case ERROR_INVALID_MSGBOX_STYLE:
		return "1438"
	case ERROR_INVALID_SPI_VALUE:
		return "1439"
	case ERROR_SCREEN_ALREADY_LOCKED:
		return "1440"
	case ERROR_HWNDS_HAVE_DIFF_PARENT:
		return "1441"
	case ERROR_NOT_CHILD_WINDOW:
		return "1442"
	case ERROR_INVALID_GW_COMMAND:
		return "1443"
	case ERROR_INVALID_THREAD_ID:
		return "1444"
	case ERROR_NON_MDICHILD_WINDOW:
		return "1445"
	case ERROR_POPUP_ALREADY_ACTIVE:
		return "1446"
	case ERROR_NO_SCROLLBARS:
		return "1447"
	case ERROR_INVALID_SCROLLBAR_RANGE:
		return "1448"
	case ERROR_INVALID_SHOWWIN_COMMAND:
		return "1449"
	case ERROR_NO_SYSTEM_RESOURCES:
		return "1450"
	case ERROR_NONPAGED_SYSTEM_RESOURCES:
		return "1451"
	case ERROR_PAGED_SYSTEM_RESOURCES:
		return "1452"
	case ERROR_WORKING_SET_QUOTA:
		return "1453"
	case ERROR_PAGEFILE_QUOTA:
		return "1454"
	case ERROR_COMMITMENT_LIMIT:
		return "1455"
	case ERROR_MENU_ITEM_NOT_FOUND:
		return "1456"
	case ERROR_INVALID_KEYBOARD_HANDLE:
		return "1457"
	case ERROR_HOOK_TYPE_NOT_ALLOWED:
		return "1458"
	case ERROR_REQUIRES_INTERACTIVE_WINDOWSTATION:
		return "1459"
	case ERROR_TIMEOUT:
		return "1460"
	case ERROR_INVALID_MONITOR_HANDLE:
		return "1461"
	case ERROR_INCORRECT_SIZE:
		return "1462"
	case ERROR_SYMLINK_CLASS_DISABLED:
		return "1463"
	case ERROR_SYMLINK_NOT_SUPPORTED:
		return "1464"
	case ERROR_XML_PARSE_ERROR:
		return "1465"
	case ERROR_XMLDSIG_ERROR:
		return "1466"
	case ERROR_RESTART_APPLICATION:
		return "1467"
	case ERROR_WRONG_COMPARTMENT:
		return "1468"
	case ERROR_AUTHIP_FAILURE:
		return "1469"
	case ERROR_NO_NVRAM_RESOURCES:
		return "1470"
	case ERROR_EVENTLOG_FILE_CORRUPT:
		return "1500"
	case ERROR_EVENTLOG_CANT_START:
		return "1501"
	case ERROR_LOG_FILE_FULL:
		return "1502"
	case ERROR_EVENTLOG_FILE_CHANGED:
		return "1503"
	case ERROR_INVALID_TASK_NAME:
		return "1550"
	case ERROR_INVALID_TASK_INDEX:
		return "1551"
	case ERROR_THREAD_ALREADY_IN_TASK:
		return "1552"
	case ERROR_INSTALL_SERVICE_FAILURE:
		return "1601"
	case ERROR_INSTALL_USEREXIT:
		return "1602"
	case ERROR_INSTALL_FAILURE:
		return "1603"
	case ERROR_INSTALL_SUSPEND:
		return "1604"
	case ERROR_UNKNOWN_PRODUCT:
		return "1605"
	case ERROR_UNKNOWN_FEATURE:
		return "1606"
	case ERROR_UNKNOWN_COMPONENT:
		return "1607"
	case ERROR_UNKNOWN_PROPERTY:
		return "1608"
	case ERROR_INVALID_HANDLE_STATE:
		return "1609"
	case ERROR_BAD_CONFIGURATION:
		return "1610"
	case ERROR_INDEX_ABSENT:
		return "1611"
	case ERROR_INSTALL_SOURCE_ABSENT:
		return "1612"
	case ERROR_INSTALL_PACKAGE_VERSION:
		return "1613"
	case ERROR_PRODUCT_UNINSTALLED:
		return "1614"
	case ERROR_BAD_QUERY_SYNTAX:
		return "1615"
	case ERROR_INVALID_FIELD:
		return "1616"
	case ERROR_DEVICE_REMOVED:
		return "1617"
	case ERROR_INSTALL_ALREADY_RUNNING:
		return "1618"
	case ERROR_INSTALL_PACKAGE_OPEN_FAILED:
		return "1619"
	case ERROR_INSTALL_PACKAGE_INVALID:
		return "1620"
	case ERROR_INSTALL_UI_FAILURE:
		return "1621"
	case ERROR_INSTALL_LOG_FAILURE:
		return "1622"
	case ERROR_INSTALL_LANGUAGE_UNSUPPORTED:
		return "1623"
	case ERROR_INSTALL_TRANSFORM_FAILURE:
		return "1624"
	case ERROR_INSTALL_PACKAGE_REJECTED:
		return "1625"
	case ERROR_FUNCTION_NOT_CALLED:
		return "1626"
	case ERROR_FUNCTION_FAILED:
		return "1627"
	case ERROR_INVALID_TABLE:
		return "1628"
	case ERROR_DATATYPE_MISMATCH:
		return "1629"
	case ERROR_UNSUPPORTED_TYPE:
		return "1630"
	case ERROR_CREATE_FAILED:
		return "1631"
	case ERROR_INSTALL_TEMP_UNWRITABLE:
		return "1632"
	case ERROR_INSTALL_PLATFORM_UNSUPPORTED:
		return "1633"
	case ERROR_INSTALL_NOTUSED:
		return "1634"
	case ERROR_PATCH_PACKAGE_OPEN_FAILED:
		return "1635"
	case ERROR_PATCH_PACKAGE_INVALID:
		return "1636"
	case ERROR_PATCH_PACKAGE_UNSUPPORTED:
		return "1637"
	case ERROR_PRODUCT_VERSION:
		return "1638"
	case ERROR_INVALID_COMMAND_LINE:
		return "1639"
	case ERROR_INSTALL_REMOTE_DISALLOWED:
		return "1640"
	case ERROR_SUCCESS_REBOOT_INITIATED:
		return "1641"
	case ERROR_PATCH_TARGET_NOT_FOUND:
		return "1642"
	case ERROR_PATCH_PACKAGE_REJECTED:
		return "1643"
	case ERROR_INSTALL_TRANSFORM_REJECTED:
		return "1644"
	case ERROR_INSTALL_REMOTE_PROHIBITED:
		return "1645"
	case ERROR_PATCH_REMOVAL_UNSUPPORTED:
		return "1646"
	case ERROR_UNKNOWN_PATCH:
		return "1647"
	case ERROR_PATCH_NO_SEQUENCE:
		return "1648"
	case ERROR_PATCH_REMOVAL_DISALLOWED:
		return "1649"
	case ERROR_INVALID_PATCH_XML:
		return "1650"
	case ERROR_PATCH_MANAGED_ADVERTISED_PRODUCT:
		return "1651"
	case ERROR_INSTALL_SERVICE_SAFEBOOT:
		return "1652"
	case ERROR_FAIL_FAST_EXCEPTION:
		return "1653"
	case RPC_S_INVALID_STRING_BINDING:
		return "1700"
	case RPC_S_WRONG_KIND_OF_BINDING:
		return "1701"
	case RPC_S_INVALID_BINDING:
		return "1702"
	case RPC_S_PROTSEQ_NOT_SUPPORTED:
		return "1703"
	case RPC_S_INVALID_RPC_PROTSEQ:
		return "1704"
	case RPC_S_INVALID_STRING_UUID:
		return "1705"
	case RPC_S_INVALID_ENDPOINT_FORMAT:
		return "1706"
	case RPC_S_INVALID_NET_ADDR:
		return "1707"
	case RPC_S_NO_ENDPOINT_FOUND:
		return "1708"
	case RPC_S_INVALID_TIMEOUT:
		return "1709"
	case RPC_S_OBJECT_NOT_FOUND:
		return "1710"
	case RPC_S_ALREADY_REGISTERED:
		return "1711"
	case RPC_S_TYPE_ALREADY_REGISTERED:
		return "1712"
	case RPC_S_ALREADY_LISTENING:
		return "1713"
	case RPC_S_NO_PROTSEQS_REGISTERED:
		return "1714"
	case RPC_S_NOT_LISTENING:
		return "1715"
	case RPC_S_UNKNOWN_MGR_TYPE:
		return "1716"
	case RPC_S_UNKNOWN_IF:
		return "1717"
	case RPC_S_NO_BINDINGS:
		return "1718"
	case RPC_S_NO_PROTSEQS:
		return "1719"
	case RPC_S_CANT_CREATE_ENDPOINT:
		return "1720"
	case RPC_S_OUT_OF_RESOURCES:
		return "1721"
	case RPC_S_SERVER_UNAVAILABLE:
		return "1722"
	case RPC_S_SERVER_TOO_BUSY:
		return "1723"
	case RPC_S_INVALID_NETWORK_OPTIONS:
		return "1724"
	case RPC_S_NO_CALL_ACTIVE:
		return "1725"
	case RPC_S_CALL_FAILED:
		return "1726"
	case RPC_S_CALL_FAILED_DNE:
		return "1727"
	case RPC_S_PROTOCOL_ERROR:
		return "1728"
	case RPC_S_PROXY_ACCESS_DENIED:
		return "1729"
	case RPC_S_UNSUPPORTED_TRANS_SYN:
		return "1730"
	case RPC_S_UNSUPPORTED_TYPE:
		return "1732"
	case RPC_S_INVALID_TAG:
		return "1733"
	case RPC_S_INVALID_BOUND:
		return "1734"
	case RPC_S_NO_ENTRY_NAME:
		return "1735"
	case RPC_S_INVALID_NAME_SYNTAX:
		return "1736"
	case RPC_S_UNSUPPORTED_NAME_SYNTAX:
		return "1737"
	case RPC_S_UUID_NO_ADDRESS:
		return "1739"
	case RPC_S_DUPLICATE_ENDPOINT:
		return "1740"
	case RPC_S_UNKNOWN_AUTHN_TYPE:
		return "1741"
	case RPC_S_MAX_CALLS_TOO_SMALL:
		return "1742"
	case RPC_S_STRING_TOO_LONG:
		return "1743"
	case RPC_S_PROTSEQ_NOT_FOUND:
		return "1744"
	case RPC_S_PROCNUM_OUT_OF_RANGE:
		return "1745"
	case RPC_S_BINDING_HAS_NO_AUTH:
		return "1746"
	case RPC_S_UNKNOWN_AUTHN_SERVICE:
		return "1747"
	case RPC_S_UNKNOWN_AUTHN_LEVEL:
		return "1748"
	case RPC_S_INVALID_AUTH_IDENTITY:
		return "1749"
	case RPC_S_UNKNOWN_AUTHZ_SERVICE:
		return "1750"
	case EPT_S_INVALID_ENTRY:
		return "1751"
	case EPT_S_CANT_PERFORM_OP:
		return "1752"
	case EPT_S_NOT_REGISTERED:
		return "1753"
	case RPC_S_NOTHING_TO_EXPORT:
		return "1754"
	case RPC_S_INCOMPLETE_NAME:
		return "1755"
	case RPC_S_INVALID_VERS_OPTION:
		return "1756"
	case RPC_S_NO_MORE_MEMBERS:
		return "1757"
	case RPC_S_NOT_ALL_OBJS_UNEXPORTED:
		return "1758"
	case RPC_S_INTERFACE_NOT_FOUND:
		return "1759"
	case RPC_S_ENTRY_ALREADY_EXISTS:
		return "1760"
	case RPC_S_ENTRY_NOT_FOUND:
		return "1761"
	case RPC_S_NAME_SERVICE_UNAVAILABLE:
		return "1762"
	case RPC_S_INVALID_NAF_ID:
		return "1763"
	case RPC_S_CANNOT_SUPPORT:
		return "1764"
	case RPC_S_NO_CONTEXT_AVAILABLE:
		return "1765"
	case RPC_S_INTERNAL_ERROR:
		return "1766"
	case RPC_S_ZERO_DIVIDE:
		return "1767"
	case RPC_S_ADDRESS_ERROR:
		return "1768"
	case RPC_S_FP_DIV_ZERO:
		return "1769"
	case RPC_S_FP_UNDERFLOW:
		return "1770"
	case RPC_S_FP_OVERFLOW:
		return "1771"
	case RPC_X_NO_MORE_ENTRIES:
		return "1772"
	case RPC_X_SS_CHAR_TRANS_OPEN_FAIL:
		return "1773"
	case RPC_X_SS_CHAR_TRANS_SHORT_FILE:
		return "1774"
	case RPC_X_SS_IN_NULL_CONTEXT:
		return "1775"
	case RPC_X_SS_CONTEXT_DAMAGED:
		return "1777"
	case RPC_X_SS_HANDLES_MISMATCH:
		return "1778"
	case RPC_X_SS_CANNOT_GET_CALL_HANDLE:
		return "1779"
	case RPC_X_NULL_REF_POINTER:
		return "1780"
	case RPC_X_ENUM_VALUE_OUT_OF_RANGE:
		return "1781"
	case RPC_X_BYTE_COUNT_TOO_SMALL:
		return "1782"
	case RPC_X_BAD_STUB_DATA:
		return "1783"
	case ERROR_INVALID_USER_BUFFER:
		return "1784"
	case ERROR_UNRECOGNIZED_MEDIA:
		return "1785"
	case ERROR_NO_TRUST_LSA_SECRET:
		return "1786"
	case ERROR_NO_TRUST_SAM_ACCOUNT:
		return "1787"
	case ERROR_TRUSTED_DOMAIN_FAILURE:
		return "1788"
	case ERROR_TRUSTED_RELATIONSHIP_FAILURE:
		return "1789"
	case ERROR_TRUST_FAILURE:
		return "1790"
	case RPC_S_CALL_IN_PROGRESS:
		return "1791"
	case ERROR_NETLOGON_NOT_STARTED:
		return "1792"
	case ERROR_ACCOUNT_EXPIRED:
		return "1793"
	case ERROR_REDIRECTOR_HAS_OPEN_HANDLES:
		return "1794"
	case ERROR_PRINTER_DRIVER_ALREADY_INSTALLED:
		return "1795"
	case ERROR_UNKNOWN_PORT:
		return "1796"
	case ERROR_UNKNOWN_PRINTER_DRIVER:
		return "1797"
	case ERROR_UNKNOWN_PRINTPROCESSOR:
		return "1798"
	case ERROR_INVALID_SEPARATOR_FILE:
		return "1799"
	case ERROR_INVALID_PRIORITY:
		return "1800"
	case ERROR_INVALID_PRINTER_NAME:
		return "1801"
	case ERROR_PRINTER_ALREADY_EXISTS:
		return "1802"
	case ERROR_INVALID_PRINTER_COMMAND:
		return "1803"
	case ERROR_INVALID_DATATYPE:
		return "1804"
	case ERROR_INVALID_ENVIRONMENT:
		return "1805"
	case RPC_S_NO_MORE_BINDINGS:
		return "1806"
	case ERROR_NOLOGON_INTERDOMAIN_TRUST_ACCOUNT:
		return "1807"
	case ERROR_NOLOGON_WORKSTATION_TRUST_ACCOUNT:
		return "1808"
	case ERROR_NOLOGON_SERVER_TRUST_ACCOUNT:
		return "1809"
	case ERROR_DOMAIN_TRUST_INCONSISTENT:
		return "1810"
	case ERROR_SERVER_HAS_OPEN_HANDLES:
		return "1811"
	case ERROR_RESOURCE_DATA_NOT_FOUND:
		return "1812"
	case ERROR_RESOURCE_TYPE_NOT_FOUND:
		return "1813"
	case ERROR_RESOURCE_NAME_NOT_FOUND:
		return "1814"
	case ERROR_RESOURCE_LANG_NOT_FOUND:
		return "1815"
	case ERROR_NOT_ENOUGH_QUOTA:
		return "1816"
	case RPC_S_NO_INTERFACES:
		return "1817"
	case RPC_S_CALL_CANCELLED:
		return "1818"
	case RPC_S_BINDING_INCOMPLETE:
		return "1819"
	case RPC_S_COMM_FAILURE:
		return "1820"
	case RPC_S_UNSUPPORTED_AUTHN_LEVEL:
		return "1821"
	case RPC_S_NO_PRINC_NAME:
		return "1822"
	case RPC_S_NOT_RPC_ERROR:
		return "1823"
	case RPC_S_UUID_LOCAL_ONLY:
		return "1824"
	case RPC_S_SEC_PKG_ERROR:
		return "1825"
	case RPC_S_NOT_CANCELLED:
		return "1826"
	case RPC_X_INVALID_ES_ACTION:
		return "1827"
	case RPC_X_WRONG_ES_VERSION:
		return "1828"
	case RPC_X_WRONG_STUB_VERSION:
		return "1829"
	case RPC_X_INVALID_PIPE_OBJECT:
		return "1830"
	case RPC_X_WRONG_PIPE_ORDER:
		return "1831"
	case RPC_X_WRONG_PIPE_VERSION:
		return "1832"
	case RPC_S_COOKIE_AUTH_FAILED:
		return "1833"
	case RPC_S_GROUP_MEMBER_NOT_FOUND:
		return "1898"
	case EPT_S_CANT_CREATE:
		return "1899"
	case RPC_S_INVALID_OBJECT:
		return "1900"
	case ERROR_INVALID_TIME:
		return "1901"
	case ERROR_INVALID_FORM_NAME:
		return "1902"
	case ERROR_INVALID_FORM_SIZE:
		return "1903"
	case ERROR_ALREADY_WAITING:
		return "1904"
	case ERROR_PRINTER_DELETED:
		return "1905"
	case ERROR_INVALID_PRINTER_STATE:
		return "1906"
	case ERROR_PASSWORD_MUST_CHANGE:
		return "1907"
	case ERROR_DOMAIN_CONTROLLER_NOT_FOUND:
		return "1908"
	case ERROR_ACCOUNT_LOCKED_OUT:
		return "1909"
	case OR_INVALID_OXID:
		return "1910"
	case OR_INVALID_OID:
		return "1911"
	case OR_INVALID_SET:
		return "1912"
	case RPC_S_SEND_INCOMPLETE:
		return "1913"
	case RPC_S_INVALID_ASYNC_HANDLE:
		return "1914"
	case RPC_S_INVALID_ASYNC_CALL:
		return "1915"
	case RPC_X_PIPE_CLOSED:
		return "1916"
	case RPC_X_PIPE_DISCIPLINE_ERROR:
		return "1917"
	case RPC_X_PIPE_EMPTY:
		return "1918"
	case ERROR_NO_SITENAME:
		return "1919"
	case ERROR_CANT_ACCESS_FILE:
		return "1920"
	case ERROR_CANT_RESOLVE_FILENAME:
		return "1921"
	case RPC_S_ENTRY_TYPE_MISMATCH:
		return "1922"
	case RPC_S_NOT_ALL_OBJS_EXPORTED:
		return "1923"
	case RPC_S_INTERFACE_NOT_EXPORTED:
		return "1924"
	case RPC_S_PROFILE_NOT_ADDED:
		return "1925"
	case RPC_S_PRF_ELT_NOT_ADDED:
		return "1926"
	case RPC_S_PRF_ELT_NOT_REMOVED:
		return "1927"
	case RPC_S_GRP_ELT_NOT_ADDED:
		return "1928"
	case RPC_S_GRP_ELT_NOT_REMOVED:
		return "1929"
	case ERROR_KM_DRIVER_BLOCKED:
		return "1930"
	case ERROR_CONTEXT_EXPIRED:
		return "1931"
	case ERROR_PER_USER_TRUST_QUOTA_EXCEEDED:
		return "1932"
	case ERROR_ALL_USER_TRUST_QUOTA_EXCEEDED:
		return "1933"
	case ERROR_USER_DELETE_TRUST_QUOTA_EXCEEDED:
		return "1934"
	case ERROR_AUTHENTICATION_FIREWALL_FAILED:
		return "1935"
	case ERROR_REMOTE_PRINT_CONNECTIONS_BLOCKED:
		return "1936"
	case ERROR_NTLM_BLOCKED:
		return "1937"
	case ERROR_INVALID_PIXEL_FORMAT:
		return "2000"
	case ERROR_BAD_DRIVER:
		return "2001"
	case ERROR_INVALID_WINDOW_STYLE:
		return "2002"
	case ERROR_METAFILE_NOT_SUPPORTED:
		return "2003"
	case ERROR_TRANSFORM_NOT_SUPPORTED:
		return "2004"
	case ERROR_CLIPPING_NOT_SUPPORTED:
		return "2005"
	case ERROR_INVALID_CMM:
		return "2010"
	case ERROR_INVALID_PROFILE:
		return "2011"
	case ERROR_TAG_NOT_FOUND:
		return "2012"
	case ERROR_TAG_NOT_PRESENT:
		return "2013"
	case ERROR_DUPLICATE_TAG:
		return "2014"
	case ERROR_PROFILE_NOT_ASSOCIATED_WITH_DEVICE:
		return "2015"
	case ERROR_PROFILE_NOT_FOUND:
		return "2016"
	case ERROR_INVALID_COLORSPACE:
		return "2017"
	case ERROR_ICM_NOT_ENABLED:
		return "2018"
	case ERROR_DELETING_ICM_XFORM:
		return "2019"
	case ERROR_INVALID_TRANSFORM:
		return "2020"
	case ERROR_COLORSPACE_MISMATCH:
		return "2021"
	case ERROR_INVALID_COLORINDEX:
		return "2022"
	case ERROR_PROFILE_DOES_NOT_MATCH_DEVICE:
		return "2023"
	case ERROR_CONNECTED_OTHER_PASSWORD:
		return "2108"
	case ERROR_CONNECTED_OTHER_PASSWORD_DEFAULT:
		return "2109"
	case ERROR_BAD_USERNAME:
		return "2202"
	case ERROR_NOT_CONNECTED:
		return "2250"
	case ERROR_OPEN_FILES:
		return "2401"
	case ERROR_ACTIVE_CONNECTIONS:
		return "2402"
	case ERROR_DEVICE_IN_USE:
		return "2404"
	case ERROR_UNKNOWN_PRINT_MONITOR:
		return "3000"
	case ERROR_PRINTER_DRIVER_IN_USE:
		return "3001"
	case ERROR_SPOOL_FILE_NOT_FOUND:
		return "3002"
	case ERROR_SPL_NO_STARTDOC:
		return "3003"
	case ERROR_SPL_NO_ADDJOB:
		return "3004"
	case ERROR_PRINT_PROCESSOR_ALREADY_INSTALLED:
		return "3005"
	case ERROR_PRINT_MONITOR_ALREADY_INSTALLED:
		return "3006"
	case ERROR_INVALID_PRINT_MONITOR:
		return "3007"
	case ERROR_PRINT_MONITOR_IN_USE:
		return "3008"
	case ERROR_PRINTER_HAS_JOBS_QUEUED:
		return "3009"
	case ERROR_SUCCESS_REBOOT_REQUIRED:
		return "3010"
	case ERROR_SUCCESS_RESTART_REQUIRED:
		return "3011"
	case ERROR_PRINTER_NOT_FOUND:
		return "3012"
	case ERROR_PRINTER_DRIVER_WARNED:
		return "3013"
	case ERROR_PRINTER_DRIVER_BLOCKED:
		return "3014"
	case ERROR_PRINTER_DRIVER_PACKAGE_IN_USE:
		return "3015"
	case ERROR_CORE_DRIVER_PACKAGE_NOT_FOUND:
		return "3016"
	case ERROR_FAIL_REBOOT_REQUIRED:
		return "3017"
	case ERROR_FAIL_REBOOT_INITIATED:
		return "3018"
	case ERROR_PRINTER_DRIVER_DOWNLOAD_NEEDED:
		return "3019"
	case ERROR_PRINT_JOB_RESTART_REQUIRED:
		return "3020"
	case ERROR_IO_REISSUE_AS_CACHED:
		return "3950"
	case ERROR_WINS_INTERNAL:
		return "4000"
	case ERROR_CAN_NOT_DEL_LOCAL_WINS:
		return "4001"
	case ERROR_STATIC_INIT:
		return "4002"
	case ERROR_INC_BACKUP:
		return "4003"
	case ERROR_FULL_BACKUP:
		return "4004"
	case ERROR_REC_NON_EXISTENT:
		return "4005"
	case ERROR_RPL_NOT_ALLOWED:
		return "4006"
	case PEERDIST_ERROR_CONTENTINFO_VERSION_UNSUPPORTED:
		return "4050"
	case PEERDIST_ERROR_CANNOT_PARSE_CONTENTINFO:
		return "4051"
	case PEERDIST_ERROR_MISSING_DATA:
		return "4052"
	case PEERDIST_ERROR_NO_MORE:
		return "4053"
	case PEERDIST_ERROR_NOT_INITIALIZED:
		return "4054"
	case PEERDIST_ERROR_ALREADY_INITIALIZED:
		return "4055"
	case PEERDIST_ERROR_SHUTDOWN_IN_PROGRESS:
		return "4056"
	case PEERDIST_ERROR_INVALIDATED:
		return "4057"
	case PEERDIST_ERROR_ALREADY_EXISTS:
		return "4058"
	case PEERDIST_ERROR_OPERATION_NOTFOUND:
		return "4059"
	case PEERDIST_ERROR_ALREADY_COMPLETED:
		return "4060"
	case PEERDIST_ERROR_OUT_OF_BOUNDS:
		return "4061"
	case PEERDIST_ERROR_VERSION_UNSUPPORTED:
		return "4062"
	case PEERDIST_ERROR_INVALID_CONFIGURATION:
		return "4063"
	case PEERDIST_ERROR_NOT_LICENSED:
		return "4064"
	case PEERDIST_ERROR_SERVICE_UNAVAILABLE:
		return "4065"
	case ERROR_DHCP_ADDRESS_CONFLICT:
		return "4100"
	case ERROR_WMI_GUID_NOT_FOUND:
		return "4200"
	case ERROR_WMI_INSTANCE_NOT_FOUND:
		return "4201"
	case ERROR_WMI_ITEMID_NOT_FOUND:
		return "4202"
	case ERROR_WMI_TRY_AGAIN:
		return "4203"
	case ERROR_WMI_DP_NOT_FOUND:
		return "4204"
	case ERROR_WMI_UNRESOLVED_INSTANCE_REF:
		return "4205"
	case ERROR_WMI_ALREADY_ENABLED:
		return "4206"
	case ERROR_WMI_GUID_DISCONNECTED:
		return "4207"
	case ERROR_WMI_SERVER_UNAVAILABLE:
		return "4208"
	case ERROR_WMI_DP_FAILED:
		return "4209"
	case ERROR_WMI_INVALID_MOF:
		return "4210"
	case ERROR_WMI_INVALID_REGINFO:
		return "4211"
	case ERROR_WMI_ALREADY_DISABLED:
		return "4212"
	case ERROR_WMI_READ_ONLY:
		return "4213"
	case ERROR_WMI_SET_FAILURE:
		return "4214"
	case ERROR_INVALID_MEDIA:
		return "4300"
	case ERROR_INVALID_LIBRARY:
		return "4301"
	case ERROR_INVALID_MEDIA_POOL:
		return "4302"
	case ERROR_DRIVE_MEDIA_MISMATCH:
		return "4303"
	case ERROR_MEDIA_OFFLINE:
		return "4304"
	case ERROR_LIBRARY_OFFLINE:
		return "4305"
	case ERROR_EMPTY:
		return "4306"
	case ERROR_NOT_EMPTY:
		return "4307"
	case ERROR_MEDIA_UNAVAILABLE:
		return "4308"
	case ERROR_RESOURCE_DISABLED:
		return "4309"
	case ERROR_INVALID_CLEANER:
		return "4310"
	case ERROR_UNABLE_TO_CLEAN:
		return "4311"
	case ERROR_OBJECT_NOT_FOUND:
		return "4312"
	case ERROR_DATABASE_FAILURE:
		return "4313"
	case ERROR_DATABASE_FULL:
		return "4314"
	case ERROR_MEDIA_INCOMPATIBLE:
		return "4315"
	case ERROR_RESOURCE_NOT_PRESENT:
		return "4316"
	case ERROR_INVALID_OPERATION:
		return "4317"
	case ERROR_MEDIA_NOT_AVAILABLE:
		return "4318"
	case ERROR_DEVICE_NOT_AVAILABLE:
		return "4319"
	case ERROR_REQUEST_REFUSED:
		return "4320"
	case ERROR_INVALID_DRIVE_OBJECT:
		return "4321"
	case ERROR_LIBRARY_FULL:
		return "4322"
	case ERROR_MEDIUM_NOT_ACCESSIBLE:
		return "4323"
	case ERROR_UNABLE_TO_LOAD_MEDIUM:
		return "4324"
	case ERROR_UNABLE_TO_INVENTORY_DRIVE:
		return "4325"
	case ERROR_UNABLE_TO_INVENTORY_SLOT:
		return "4326"
	case ERROR_UNABLE_TO_INVENTORY_TRANSPORT:
		return "4327"
	case ERROR_TRANSPORT_FULL:
		return "4328"
	case ERROR_CONTROLLING_IEPORT:
		return "4329"
	case ERROR_UNABLE_TO_EJECT_MOUNTED_MEDIA:
		return "4330"
	case ERROR_CLEANER_SLOT_SET:
		return "4331"
	case ERROR_CLEANER_SLOT_NOT_SET:
		return "4332"
	case ERROR_CLEANER_CARTRIDGE_SPENT:
		return "4333"
	case ERROR_UNEXPECTED_OMID:
		return "4334"
	case ERROR_CANT_DELETE_LAST_ITEM:
		return "4335"
	case ERROR_MESSAGE_EXCEEDS_MAX_SIZE:
		return "4336"
	case ERROR_VOLUME_CONTAINS_SYS_FILES:
		return "4337"
	case ERROR_INDIGENOUS_TYPE:
		return "4338"
	case ERROR_NO_SUPPORTING_DRIVES:
		return "4339"
	case ERROR_CLEANER_CARTRIDGE_INSTALLED:
		return "4340"
	case ERROR_IEPORT_FULL:
		return "4341"
	case ERROR_FILE_OFFLINE:
		return "4350"
	case ERROR_REMOTE_STORAGE_NOT_ACTIVE:
		return "4351"
	case ERROR_REMOTE_STORAGE_MEDIA_ERROR:
		return "4352"
	case ERROR_NOT_A_REPARSE_POINT:
		return "4390"
	case ERROR_REPARSE_ATTRIBUTE_CONFLICT:
		return "4391"
	case ERROR_INVALID_REPARSE_DATA:
		return "4392"
	case ERROR_REPARSE_TAG_INVALID:
		return "4393"
	case ERROR_REPARSE_TAG_MISMATCH:
		return "4394"
	case ERROR_VOLUME_NOT_SIS_ENABLED:
		return "4500"
	case ERROR_DEPENDENT_RESOURCE_EXISTS:
		return "5001"
	case ERROR_DEPENDENCY_NOT_FOUND:
		return "5002"
	case ERROR_DEPENDENCY_ALREADY_EXISTS:
		return "5003"
	case ERROR_RESOURCE_NOT_ONLINE:
		return "5004"
	case ERROR_HOST_NODE_NOT_AVAILABLE:
		return "5005"
	case ERROR_RESOURCE_NOT_AVAILABLE:
		return "5006"
	case ERROR_RESOURCE_NOT_FOUND:
		return "5007"
	case ERROR_SHUTDOWN_CLUSTER:
		return "5008"
	case ERROR_CANT_EVICT_ACTIVE_NODE:
		return "5009"
	case ERROR_OBJECT_ALREADY_EXISTS:
		return "5010"
	case ERROR_OBJECT_IN_LIST:
		return "5011"
	case ERROR_GROUP_NOT_AVAILABLE:
		return "5012"
	case ERROR_GROUP_NOT_FOUND:
		return "5013"
	case ERROR_GROUP_NOT_ONLINE:
		return "5014"
	case ERROR_HOST_NODE_NOT_RESOURCE_OWNER:
		return "5015"
	case ERROR_HOST_NODE_NOT_GROUP_OWNER:
		return "5016"
	case ERROR_RESMON_CREATE_FAILED:
		return "5017"
	case ERROR_RESMON_ONLINE_FAILED:
		return "5018"
	case ERROR_RESOURCE_ONLINE:
		return "5019"
	case ERROR_QUORUM_RESOURCE:
		return "5020"
	case ERROR_NOT_QUORUM_CAPABLE:
		return "5021"
	case ERROR_CLUSTER_SHUTTING_DOWN:
		return "5022"
	case ERROR_INVALID_STATE:
		return "5023"
	case ERROR_RESOURCE_PROPERTIES_STORED:
		return "5024"
	case ERROR_NOT_QUORUM_CLASS:
		return "5025"
	case ERROR_CORE_RESOURCE:
		return "5026"
	case ERROR_QUORUM_RESOURCE_ONLINE_FAILED:
		return "5027"
	case ERROR_QUORUMLOG_OPEN_FAILED:
		return "5028"
	case ERROR_CLUSTERLOG_CORRUPT:
		return "5029"
	case ERROR_CLUSTERLOG_RECORD_EXCEEDS_MAXSIZE:
		return "5030"
	case ERROR_CLUSTERLOG_EXCEEDS_MAXSIZE:
		return "5031"
	case ERROR_CLUSTERLOG_CHKPOINT_NOT_FOUND:
		return "5032"
	case ERROR_CLUSTERLOG_NOT_ENOUGH_SPACE:
		return "5033"
	case ERROR_QUORUM_OWNER_ALIVE:
		return "5034"
	case ERROR_NETWORK_NOT_AVAILABLE:
		return "5035"
	case ERROR_NODE_NOT_AVAILABLE:
		return "5036"
	case ERROR_ALL_NODES_NOT_AVAILABLE:
		return "5037"
	case ERROR_RESOURCE_FAILED:
		return "5038"
	case ERROR_CLUSTER_INVALID_NODE:
		return "5039"
	case ERROR_CLUSTER_NODE_EXISTS:
		return "5040"
	case ERROR_CLUSTER_JOIN_IN_PROGRESS:
		return "5041"
	case ERROR_CLUSTER_NODE_NOT_FOUND:
		return "5042"
	case ERROR_CLUSTER_LOCAL_NODE_NOT_FOUND:
		return "5043"
	case ERROR_CLUSTER_NETWORK_EXISTS:
		return "5044"
	case ERROR_CLUSTER_NETWORK_NOT_FOUND:
		return "5045"
	case ERROR_CLUSTER_NETINTERFACE_EXISTS:
		return "5046"
	case ERROR_CLUSTER_NETINTERFACE_NOT_FOUND:
		return "5047"
	case ERROR_CLUSTER_INVALID_REQUEST:
		return "5048"
	case ERROR_CLUSTER_INVALID_NETWORK_PROVIDER:
		return "5049"
	case ERROR_CLUSTER_NODE_DOWN:
		return "5050"
	case ERROR_CLUSTER_NODE_UNREACHABLE:
		return "5051"
	case ERROR_CLUSTER_NODE_NOT_MEMBER:
		return "5052"
	case ERROR_CLUSTER_JOIN_NOT_IN_PROGRESS:
		return "5053"
	case ERROR_CLUSTER_INVALID_NETWORK:
		return "5054"
	case ERROR_CLUSTER_NODE_UP:
		return "5056"
	case ERROR_CLUSTER_IPADDR_IN_USE:
		return "5057"
	case ERROR_CLUSTER_NODE_NOT_PAUSED:
		return "5058"
	case ERROR_CLUSTER_NO_SECURITY_CONTEXT:
		return "5059"
	case ERROR_CLUSTER_NETWORK_NOT_INTERNAL:
		return "5060"
	case ERROR_CLUSTER_NODE_ALREADY_UP:
		return "5061"
	case ERROR_CLUSTER_NODE_ALREADY_DOWN:
		return "5062"
	case ERROR_CLUSTER_NETWORK_ALREADY_ONLINE:
		return "5063"
	case ERROR_CLUSTER_NETWORK_ALREADY_OFFLINE:
		return "5064"
	case ERROR_CLUSTER_NODE_ALREADY_MEMBER:
		return "5065"
	case ERROR_CLUSTER_LAST_INTERNAL_NETWORK:
		return "5066"
	case ERROR_CLUSTER_NETWORK_HAS_DEPENDENTS:
		return "5067"
	case ERROR_INVALID_OPERATION_ON_QUORUM:
		return "5068"
	case ERROR_DEPENDENCY_NOT_ALLOWED:
		return "5069"
	case ERROR_CLUSTER_NODE_PAUSED:
		return "5070"
	case ERROR_NODE_CANT_HOST_RESOURCE:
		return "5071"
	case ERROR_CLUSTER_NODE_NOT_READY:
		return "5072"
	case ERROR_CLUSTER_NODE_SHUTTING_DOWN:
		return "5073"
	case ERROR_CLUSTER_JOIN_ABORTED:
		return "5074"
	case ERROR_CLUSTER_INCOMPATIBLE_VERSIONS:
		return "5075"
	case ERROR_CLUSTER_MAXNUM_OF_RESOURCES_EXCEEDED:
		return "5076"
	case ERROR_CLUSTER_SYSTEM_CONFIG_CHANGED:
		return "5077"
	case ERROR_CLUSTER_RESOURCE_TYPE_NOT_FOUND:
		return "5078"
	case ERROR_CLUSTER_RESTYPE_NOT_SUPPORTED:
		return "5079"
	case ERROR_CLUSTER_RESNAME_NOT_FOUND:
		return "5080"
	case ERROR_CLUSTER_NO_RPC_PACKAGES_REGISTERED:
		return "5081"
	case ERROR_CLUSTER_OWNER_NOT_IN_PREFLIST:
		return "5082"
	case ERROR_CLUSTER_DATABASE_SEQMISMATCH:
		return "5083"
	case ERROR_RESMON_INVALID_STATE:
		return "5084"
	case ERROR_CLUSTER_GUM_NOT_LOCKER:
		return "5085"
	case ERROR_QUORUM_DISK_NOT_FOUND:
		return "5086"
	case ERROR_DATABASE_BACKUP_CORRUPT:
		return "5087"
	case ERROR_CLUSTER_NODE_ALREADY_HAS_DFS_ROOT:
		return "5088"
	case ERROR_RESOURCE_PROPERTY_UNCHANGEABLE:
		return "5089"
	case ERROR_CLUSTER_MEMBERSHIP_INVALID_STATE:
		return "5890"
	case ERROR_CLUSTER_QUORUMLOG_NOT_FOUND:
		return "5891"
	case ERROR_CLUSTER_MEMBERSHIP_HALT:
		return "5892"
	case ERROR_CLUSTER_INSTANCE_ID_MISMATCH:
		return "5893"
	case ERROR_CLUSTER_NETWORK_NOT_FOUND_FOR_IP:
		return "5894"
	case ERROR_CLUSTER_PROPERTY_DATA_TYPE_MISMATCH:
		return "5895"
	case ERROR_CLUSTER_EVICT_WITHOUT_CLEANUP:
		return "5896"
	case ERROR_CLUSTER_PARAMETER_MISMATCH:
		return "5897"
	case ERROR_NODE_CANNOT_BE_CLUSTERED:
		return "5898"
	case ERROR_CLUSTER_WRONG_OS_VERSION:
		return "5899"
	case ERROR_CLUSTER_CANT_CREATE_DUP_CLUSTER_NAME:
		return "5900"
	case ERROR_CLUSCFG_ALREADY_COMMITTED:
		return "5901"
	case ERROR_CLUSCFG_ROLLBACK_FAILED:
		return "5902"
	case ERROR_CLUSCFG_SYSTEM_DISK_DRIVE_LETTER_CONFLICT:
		return "5903"
	case ERROR_CLUSTER_OLD_VERSION:
		return "5904"
	case ERROR_CLUSTER_MISMATCHED_COMPUTER_ACCT_NAME:
		return "5905"
	case ERROR_CLUSTER_NO_NET_ADAPTERS:
		return "5906"
	case ERROR_CLUSTER_POISONED:
		return "5907"
	case ERROR_CLUSTER_GROUP_MOVING:
		return "5908"
	case ERROR_CLUSTER_RESOURCE_TYPE_BUSY:
		return "5909"
	case ERROR_RESOURCE_CALL_TIMED_OUT:
		return "5910"
	case ERROR_INVALID_CLUSTER_IPV6_ADDRESS:
		return "5911"
	case ERROR_CLUSTER_INTERNAL_INVALID_FUNCTION:
		return "5912"
	case ERROR_CLUSTER_PARAMETER_OUT_OF_BOUNDS:
		return "5913"
	case ERROR_CLUSTER_PARTIAL_SEND:
		return "5914"
	case ERROR_CLUSTER_REGISTRY_INVALID_FUNCTION:
		return "5915"
	case ERROR_CLUSTER_INVALID_STRING_TERMINATION:
		return "5916"
	case ERROR_CLUSTER_INVALID_STRING_FORMAT:
		return "5917"
	case ERROR_CLUSTER_DATABASE_TRANSACTION_IN_PROGRESS:
		return "5918"
	case ERROR_CLUSTER_DATABASE_TRANSACTION_NOT_IN_PROGRESS:
		return "5919"
	case ERROR_CLUSTER_NULL_DATA:
		return "5920"
	case ERROR_CLUSTER_PARTIAL_READ:
		return "5921"
	case ERROR_CLUSTER_PARTIAL_WRITE:
		return "5922"
	case ERROR_CLUSTER_CANT_DESERIALIZE_DATA:
		return "5923"
	case ERROR_DEPENDENT_RESOURCE_PROPERTY_CONFLICT:
		return "5924"
	case ERROR_CLUSTER_NO_QUORUM:
		return "5925"
	case ERROR_CLUSTER_INVALID_IPV6_NETWORK:
		return "5926"
	case ERROR_CLUSTER_INVALID_IPV6_TUNNEL_NETWORK:
		return "5927"
	case ERROR_QUORUM_NOT_ALLOWED_IN_THIS_GROUP:
		return "5928"
	case ERROR_DEPENDENCY_TREE_TOO_COMPLEX:
		return "5929"
	case ERROR_EXCEPTION_IN_RESOURCE_CALL:
		return "5930"
	case ERROR_CLUSTER_RHS_FAILED_INITIALIZATION:
		return "5931"
	case ERROR_CLUSTER_NOT_INSTALLED:
		return "5932"
	case ERROR_CLUSTER_RESOURCES_MUST_BE_ONLINE_ON_THE_SAME_NODE:
		return "5933"
	case ERROR_CLUSTER_MAX_NODES_IN_CLUSTER:
		return "5934"
	case ERROR_CLUSTER_TOO_MANY_NODES:
		return "5935"
	case ERROR_CLUSTER_OBJECT_ALREADY_USED:
		return "5936"
	case ERROR_NONCORE_GROUPS_FOUND:
		return "5937"
	case ERROR_FILE_SHARE_RESOURCE_CONFLICT:
		return "5938"
	case ERROR_CLUSTER_EVICT_INVALID_REQUEST:
		return "5939"
	case ERROR_CLUSTER_SINGLETON_RESOURCE:
		return "5940"
	case ERROR_CLUSTER_GROUP_SINGLETON_RESOURCE:
		return "5941"
	case ERROR_CLUSTER_RESOURCE_PROVIDER_FAILED:
		return "5942"
	case ERROR_CLUSTER_RESOURCE_CONFIGURATION_ERROR:
		return "5943"
	case ERROR_CLUSTER_GROUP_BUSY:
		return "5944"
	case ERROR_CLUSTER_NOT_SHARED_VOLUME:
		return "5945"
	case ERROR_CLUSTER_INVALID_SECURITY_DESCRIPTOR:
		return "5946"
	case ERROR_CLUSTER_SHARED_VOLUMES_IN_USE:
		return "5947"
	case ERROR_CLUSTER_USE_SHARED_VOLUMES_API:
		return "5948"
	case ERROR_CLUSTER_BACKUP_IN_PROGRESS:
		return "5949"
	case ERROR_NON_CSV_PATH:
		return "5950"
	case ERROR_CSV_VOLUME_NOT_LOCAL:
		return "5951"
	case ERROR_CLUSTER_WATCHDOG_TERMINATING:
		return "5952"
	case ERROR_ENCRYPTION_FAILED:
		return "6000"
	case ERROR_DECRYPTION_FAILED:
		return "6001"
	case ERROR_FILE_ENCRYPTED:
		return "6002"
	case ERROR_NO_RECOVERY_POLICY:
		return "6003"
	case ERROR_NO_EFS:
		return "6004"
	case ERROR_WRONG_EFS:
		return "6005"
	case ERROR_NO_USER_KEYS:
		return "6006"
	case ERROR_FILE_NOT_ENCRYPTED:
		return "6007"
	case ERROR_NOT_EXPORT_FORMAT:
		return "6008"
	case ERROR_FILE_READ_ONLY:
		return "6009"
	case ERROR_DIR_EFS_DISALLOWED:
		return "6010"
	case ERROR_EFS_SERVER_NOT_TRUSTED:
		return "6011"
	case ERROR_BAD_RECOVERY_POLICY:
		return "6012"
	case ERROR_EFS_ALG_BLOB_TOO_BIG:
		return "6013"
	case ERROR_VOLUME_NOT_SUPPORT_EFS:
		return "6014"
	case ERROR_EFS_DISABLED:
		return "6015"
	case ERROR_EFS_VERSION_NOT_SUPPORT:
		return "6016"
	case ERROR_CS_ENCRYPTION_INVALID_SERVER_RESPONSE:
		return "6017"
	case ERROR_CS_ENCRYPTION_UNSUPPORTED_SERVER:
		return "6018"
	case ERROR_CS_ENCRYPTION_EXISTING_ENCRYPTED_FILE:
		return "6019"
	case ERROR_CS_ENCRYPTION_NEW_ENCRYPTED_FILE:
		return "6020"
	case ERROR_CS_ENCRYPTION_FILE_NOT_CSE:
		return "6021"
	case ERROR_ENCRYPTION_POLICY_DENIES_OPERATION:
		return "6022"
	case ERROR_NO_BROWSER_SERVERS_FOUND:
		return "6118"
	case SCHED_E_SERVICE_NOT_LOCALSYSTEM:
		return "6200"
	case ERROR_LOG_SECTOR_INVALID:
		return "6600"
	case ERROR_LOG_SECTOR_PARITY_INVALID:
		return "6601"
	case ERROR_LOG_SECTOR_REMAPPED:
		return "6602"
	case ERROR_LOG_BLOCK_INCOMPLETE:
		return "6603"
	case ERROR_LOG_INVALID_RANGE:
		return "6604"
	case ERROR_LOG_BLOCKS_EXHAUSTED:
		return "6605"
	case ERROR_LOG_READ_CONTEXT_INVALID:
		return "6606"
	case ERROR_LOG_RESTART_INVALID:
		return "6607"
	case ERROR_LOG_BLOCK_VERSION:
		return "6608"
	case ERROR_LOG_BLOCK_INVALID:
		return "6609"
	case ERROR_LOG_READ_MODE_INVALID:
		return "6610"
	case ERROR_LOG_NO_RESTART:
		return "6611"
	case ERROR_LOG_METADATA_CORRUPT:
		return "6612"
	case ERROR_LOG_METADATA_INVALID:
		return "6613"
	case ERROR_LOG_METADATA_INCONSISTENT:
		return "6614"
	case ERROR_LOG_RESERVATION_INVALID:
		return "6615"
	case ERROR_LOG_CANT_DELETE:
		return "6616"
	case ERROR_LOG_CONTAINER_LIMIT_EXCEEDED:
		return "6617"
	case ERROR_LOG_START_OF_LOG:
		return "6618"
	case ERROR_LOG_POLICY_ALREADY_INSTALLED:
		return "6619"
	case ERROR_LOG_POLICY_NOT_INSTALLED:
		return "6620"
	case ERROR_LOG_POLICY_INVALID:
		return "6621"
	case ERROR_LOG_POLICY_CONFLICT:
		return "6622"
	case ERROR_LOG_PINNED_ARCHIVE_TAIL:
		return "6623"
	case ERROR_LOG_RECORD_NONEXISTENT:
		return "6624"
	case ERROR_LOG_RECORDS_RESERVED_INVALID:
		return "6625"
	case ERROR_LOG_SPACE_RESERVED_INVALID:
		return "6626"
	case ERROR_LOG_TAIL_INVALID:
		return "6627"
	case ERROR_LOG_FULL:
		return "6628"
	case ERROR_COULD_NOT_RESIZE_LOG:
		return "6629"
	case ERROR_LOG_MULTIPLEXED:
		return "6630"
	case ERROR_LOG_DEDICATED:
		return "6631"
	case ERROR_LOG_ARCHIVE_NOT_IN_PROGRESS:
		return "6632"
	case ERROR_LOG_ARCHIVE_IN_PROGRESS:
		return "6633"
	case ERROR_LOG_EPHEMERAL:
		return "6634"
	case ERROR_LOG_NOT_ENOUGH_CONTAINERS:
		return "6635"
	case ERROR_LOG_CLIENT_ALREADY_REGISTERED:
		return "6636"
	case ERROR_LOG_CLIENT_NOT_REGISTERED:
		return "6637"
	case ERROR_LOG_FULL_HANDLER_IN_PROGRESS:
		return "6638"
	case ERROR_LOG_CONTAINER_READ_FAILED:
		return "6639"
	case ERROR_LOG_CONTAINER_WRITE_FAILED:
		return "6640"
	case ERROR_LOG_CONTAINER_OPEN_FAILED:
		return "6641"
	case ERROR_LOG_CONTAINER_STATE_INVALID:
		return "6642"
	case ERROR_LOG_STATE_INVALID:
		return "6643"
	case ERROR_LOG_PINNED:
		return "6644"
	case ERROR_LOG_METADATA_FLUSH_FAILED:
		return "6645"
	case ERROR_LOG_INCONSISTENT_SECURITY:
		return "6646"
	case ERROR_LOG_APPENDED_FLUSH_FAILED:
		return "6647"
	case ERROR_LOG_PINNED_RESERVATION:
		return "6648"
	case ERROR_INVALID_TRANSACTION:
		return "6700"
	case ERROR_TRANSACTION_NOT_ACTIVE:
		return "6701"
	case ERROR_TRANSACTION_REQUEST_NOT_VALID:
		return "6702"
	case ERROR_TRANSACTION_NOT_REQUESTED:
		return "6703"
	case ERROR_TRANSACTION_ALREADY_ABORTED:
		return "6704"
	case ERROR_TRANSACTION_ALREADY_COMMITTED:
		return "6705"
	case ERROR_TM_INITIALIZATION_FAILED:
		return "6706"
	case ERROR_RESOURCEMANAGER_READ_ONLY:
		return "6707"
	case ERROR_TRANSACTION_NOT_JOINED:
		return "6708"
	case ERROR_TRANSACTION_SUPERIOR_EXISTS:
		return "6709"
	case ERROR_CRM_PROTOCOL_ALREADY_EXISTS:
		return "6710"
	case ERROR_TRANSACTION_PROPAGATION_FAILED:
		return "6711"
	case ERROR_CRM_PROTOCOL_NOT_FOUND:
		return "6712"
	case ERROR_TRANSACTION_INVALID_MARSHALL_BUFFER:
		return "6713"
	case ERROR_CURRENT_TRANSACTION_NOT_VALID:
		return "6714"
	case ERROR_TRANSACTION_NOT_FOUND:
		return "6715"
	case ERROR_RESOURCEMANAGER_NOT_FOUND:
		return "6716"
	case ERROR_ENLISTMENT_NOT_FOUND:
		return "6717"
	case ERROR_TRANSACTIONMANAGER_NOT_FOUND:
		return "6718"
	case ERROR_TRANSACTIONMANAGER_NOT_ONLINE:
		return "6719"
	case ERROR_TRANSACTIONMANAGER_RECOVERY_NAME_COLLISION:
		return "6720"
	case ERROR_TRANSACTION_NOT_ROOT:
		return "6721"
	case ERROR_TRANSACTION_OBJECT_EXPIRED:
		return "6722"
	case ERROR_TRANSACTION_RESPONSE_NOT_ENLISTED:
		return "6723"
	case ERROR_TRANSACTION_RECORD_TOO_LONG:
		return "6724"
	case ERROR_IMPLICIT_TRANSACTION_NOT_SUPPORTED:
		return "6725"
	case ERROR_TRANSACTION_INTEGRITY_VIOLATED:
		return "6726"
	case ERROR_TRANSACTIONMANAGER_IDENTITY_MISMATCH:
		return "6727"
	case ERROR_RM_CANNOT_BE_FROZEN_FOR_SNAPSHOT:
		return "6728"
	case ERROR_TRANSACTION_MUST_WRITETHROUGH:
		return "6729"
	case ERROR_TRANSACTION_NO_SUPERIOR:
		return "6730"
	case ERROR_HEURISTIC_DAMAGE_POSSIBLE:
		return "6731"
	case ERROR_TRANSACTIONAL_CONFLICT:
		return "6800"
	case ERROR_RM_NOT_ACTIVE:
		return "6801"
	case ERROR_RM_METADATA_CORRUPT:
		return "6802"
	case ERROR_DIRECTORY_NOT_RM:
		return "6803"
	case ERROR_TRANSACTIONS_UNSUPPORTED_REMOTE:
		return "6805"
	case ERROR_LOG_RESIZE_INVALID_SIZE:
		return "6806"
	case ERROR_OBJECT_NO_LONGER_EXISTS:
		return "6807"
	case ERROR_STREAM_MINIVERSION_NOT_FOUND:
		return "6808"
	case ERROR_STREAM_MINIVERSION_NOT_VALID:
		return "6809"
	case ERROR_MINIVERSION_INACCESSIBLE_FROM_SPECIFIED_TRANSACTION:
		return "6810"
	case ERROR_CANT_OPEN_MINIVERSION_WITH_MODIFY_INTENT:
		return "6811"
	case ERROR_CANT_CREATE_MORE_STREAM_MINIVERSIONS:
		return "6812"
	case ERROR_REMOTE_FILE_VERSION_MISMATCH:
		return "6814"
	case ERROR_HANDLE_NO_LONGER_VALID:
		return "6815"
	case ERROR_NO_TXF_METADATA:
		return "6816"
	case ERROR_LOG_CORRUPTION_DETECTED:
		return "6817"
	case ERROR_CANT_RECOVER_WITH_HANDLE_OPEN:
		return "6818"
	case ERROR_RM_DISCONNECTED:
		return "6819"
	case ERROR_ENLISTMENT_NOT_SUPERIOR:
		return "6820"
	case ERROR_RECOVERY_NOT_NEEDED:
		return "6821"
	case ERROR_RM_ALREADY_STARTED:
		return "6822"
	case ERROR_FILE_IDENTITY_NOT_PERSISTENT:
		return "6823"
	case ERROR_CANT_BREAK_TRANSACTIONAL_DEPENDENCY:
		return "6824"
	case ERROR_CANT_CROSS_RM_BOUNDARY:
		return "6825"
	case ERROR_TXF_DIR_NOT_EMPTY:
		return "6826"
	case ERROR_INDOUBT_TRANSACTIONS_EXIST:
		return "6827"
	case ERROR_TM_VOLATILE:
		return "6828"
	case ERROR_ROLLBACK_TIMER_EXPIRED:
		return "6829"
	case ERROR_TXF_ATTRIBUTE_CORRUPT:
		return "6830"
	case ERROR_EFS_NOT_ALLOWED_IN_TRANSACTION:
		return "6831"
	case ERROR_TRANSACTIONAL_OPEN_NOT_ALLOWED:
		return "6832"
	case ERROR_LOG_GROWTH_FAILED:
		return "6833"
	case ERROR_TRANSACTED_MAPPING_UNSUPPORTED_REMOTE:
		return "6834"
	case ERROR_TXF_METADATA_ALREADY_PRESENT:
		return "6835"
	case ERROR_TRANSACTION_SCOPE_CALLBACKS_NOT_SET:
		return "6836"
	case ERROR_TRANSACTION_REQUIRED_PROMOTION:
		return "6837"
	case ERROR_CANNOT_EXECUTE_FILE_IN_TRANSACTION:
		return "6838"
	case ERROR_TRANSACTIONS_NOT_FROZEN:
		return "6839"
	case ERROR_TRANSACTION_FREEZE_IN_PROGRESS:
		return "6840"
	case ERROR_NOT_SNAPSHOT_VOLUME:
		return "6841"
	case ERROR_NO_SAVEPOINT_WITH_OPEN_FILES:
		return "6842"
	case ERROR_DATA_LOST_REPAIR:
		return "6843"
	case ERROR_SPARSE_NOT_ALLOWED_IN_TRANSACTION:
		return "6844"
	case ERROR_TM_IDENTITY_MISMATCH:
		return "6845"
	case ERROR_FLOATED_SECTION:
		return "6846"
	case ERROR_CANNOT_ACCEPT_TRANSACTED_WORK:
		return "6847"
	case ERROR_CANNOT_ABORT_TRANSACTIONS:
		return "6848"
	case ERROR_BAD_CLUSTERS:
		return "6849"
	case ERROR_COMPRESSION_NOT_ALLOWED_IN_TRANSACTION:
		return "6850"
	case ERROR_VOLUME_DIRTY:
		return "6851"
	case ERROR_NO_LINK_TRACKING_IN_TRANSACTION:
		return "6852"
	case ERROR_OPERATION_NOT_SUPPORTED_IN_TRANSACTION:
		return "6853"
	case ERROR_EXPIRED_HANDLE:
		return "6854"
	case ERROR_TRANSACTION_NOT_ENLISTED:
		return "6855"
	case ERROR_CTX_WINSTATION_NAME_INVALID:
		return "7001"
	case ERROR_CTX_INVALID_PD:
		return "7002"
	case ERROR_CTX_PD_NOT_FOUND:
		return "7003"
	case ERROR_CTX_WD_NOT_FOUND:
		return "7004"
	case ERROR_CTX_CANNOT_MAKE_EVENTLOG_ENTRY:
		return "7005"
	case ERROR_CTX_SERVICE_NAME_COLLISION:
		return "7006"
	case ERROR_CTX_CLOSE_PENDING:
		return "7007"
	case ERROR_CTX_NO_OUTBUF:
		return "7008"
	case ERROR_CTX_MODEM_INF_NOT_FOUND:
		return "7009"
	case ERROR_CTX_INVALID_MODEMNAME:
		return "7010"
	case ERROR_CTX_MODEM_RESPONSE_ERROR:
		return "7011"
	case ERROR_CTX_MODEM_RESPONSE_TIMEOUT:
		return "7012"
	case ERROR_CTX_MODEM_RESPONSE_NO_CARRIER:
		return "7013"
	case ERROR_CTX_MODEM_RESPONSE_NO_DIALTONE:
		return "7014"
	case ERROR_CTX_MODEM_RESPONSE_BUSY:
		return "7015"
	case ERROR_CTX_MODEM_RESPONSE_VOICE:
		return "7016"
	case ERROR_CTX_TD_ERROR:
		return "7017"
	case ERROR_CTX_WINSTATION_NOT_FOUND:
		return "7022"
	case ERROR_CTX_WINSTATION_ALREADY_EXISTS:
		return "7023"
	case ERROR_CTX_WINSTATION_BUSY:
		return "7024"
	case ERROR_CTX_BAD_VIDEO_MODE:
		return "7025"
	case ERROR_CTX_GRAPHICS_INVALID:
		return "7035"
	case ERROR_CTX_LOGON_DISABLED:
		return "7037"
	case ERROR_CTX_NOT_CONSOLE:
		return "7038"
	case ERROR_CTX_CLIENT_QUERY_TIMEOUT:
		return "7040"
	case ERROR_CTX_CONSOLE_DISCONNECT:
		return "7041"
	case ERROR_CTX_CONSOLE_CONNECT:
		return "7042"
	case ERROR_CTX_SHADOW_DENIED:
		return "7044"
	case ERROR_CTX_WINSTATION_ACCESS_DENIED:
		return "7045"
	case ERROR_CTX_INVALID_WD:
		return "7049"
	case ERROR_CTX_SHADOW_INVALID:
		return "7050"
	case ERROR_CTX_SHADOW_DISABLED:
		return "7051"
	case ERROR_CTX_CLIENT_LICENSE_IN_USE:
		return "7052"
	case ERROR_CTX_CLIENT_LICENSE_NOT_SET:
		return "7053"
	case ERROR_CTX_LICENSE_NOT_AVAILABLE:
		return "7054"
	case ERROR_CTX_LICENSE_CLIENT_INVALID:
		return "7055"
	case ERROR_CTX_LICENSE_EXPIRED:
		return "7056"
	case ERROR_CTX_SHADOW_NOT_RUNNING:
		return "7057"
	case ERROR_CTX_SHADOW_ENDED_BY_MODE_CHANGE:
		return "7058"
	case ERROR_ACTIVATION_COUNT_EXCEEDED:
		return "7059"
	case ERROR_CTX_WINSTATIONS_DISABLED:
		return "7060"
	case ERROR_CTX_ENCRYPTION_LEVEL_REQUIRED:
		return "7061"
	case ERROR_CTX_SESSION_IN_USE:
		return "7062"
	case ERROR_CTX_NO_FORCE_LOGOFF:
		return "7063"
	case ERROR_CTX_ACCOUNT_RESTRICTION:
		return "7064"
	case ERROR_RDP_PROTOCOL_ERROR:
		return "7065"
	case ERROR_CTX_CDM_CONNECT:
		return "7066"
	case ERROR_CTX_CDM_DISCONNECT:
		return "7067"
	case ERROR_CTX_SECURITY_LAYER_ERROR:
		return "7068"
	case ERROR_TS_INCOMPATIBLE_SESSIONS:
		return "7069"
	case ERROR_TS_VIDEO_SUBSYSTEM_ERROR:
		return "7070"
	case FRS_ERR_INVALID_API_SEQUENCE:
		return "8001"
	case FRS_ERR_STARTING_SERVICE:
		return "8002"
	case FRS_ERR_STOPPING_SERVICE:
		return "8003"
	case FRS_ERR_INTERNAL_API:
		return "8004"
	case FRS_ERR_INTERNAL:
		return "8005"
	case FRS_ERR_SERVICE_COMM:
		return "8006"
	case FRS_ERR_INSUFFICIENT_PRIV:
		return "8007"
	case FRS_ERR_AUTHENTICATION:
		return "8008"
	case FRS_ERR_PARENT_INSUFFICIENT_PRIV:
		return "8009"
	case FRS_ERR_PARENT_AUTHENTICATION:
		return "8010"
	case FRS_ERR_CHILD_TO_PARENT_COMM:
		return "8011"
	case FRS_ERR_PARENT_TO_CHILD_COMM:
		return "8012"
	case FRS_ERR_SYSVOL_POPULATE:
		return "8013"
	case FRS_ERR_SYSVOL_POPULATE_TIMEOUT:
		return "8014"
	case FRS_ERR_SYSVOL_IS_BUSY:
		return "8015"
	case FRS_ERR_SYSVOL_DEMOTE:
		return "8016"
	case FRS_ERR_INVALID_SERVICE_PARAMETER:
		return "8017"
	case ERROR_DS_NOT_INSTALLED:
		return "8200"
	case ERROR_DS_MEMBERSHIP_EVALUATED_LOCALLY:
		return "8201"
	case ERROR_DS_NO_ATTRIBUTE_OR_VALUE:
		return "8202"
	case ERROR_DS_INVALID_ATTRIBUTE_SYNTAX:
		return "8203"
	case ERROR_DS_ATTRIBUTE_TYPE_UNDEFINED:
		return "8204"
	case ERROR_DS_ATTRIBUTE_OR_VALUE_EXISTS:
		return "8205"
	case ERROR_DS_BUSY:
		return "8206"
	case ERROR_DS_UNAVAILABLE:
		return "8207"
	case ERROR_DS_NO_RIDS_ALLOCATED:
		return "8208"
	case ERROR_DS_NO_MORE_RIDS:
		return "8209"
	case ERROR_DS_INCORRECT_ROLE_OWNER:
		return "8210"
	case ERROR_DS_RIDMGR_INIT_ERROR:
		return "8211"
	case ERROR_DS_OBJ_CLASS_VIOLATION:
		return "8212"
	case ERROR_DS_CANT_ON_NON_LEAF:
		return "8213"
	case ERROR_DS_CANT_ON_RDN:
		return "8214"
	case ERROR_DS_CANT_MOD_OBJ_CLASS:
		return "8215"
	case ERROR_DS_CROSS_DOM_MOVE_ERROR:
		return "8216"
	case ERROR_DS_GC_NOT_AVAILABLE:
		return "8217"
	case ERROR_SHARED_POLICY:
		return "8218"
	case ERROR_POLICY_OBJECT_NOT_FOUND:
		return "8219"
	case ERROR_POLICY_ONLY_IN_DS:
		return "8220"
	case ERROR_PROMOTION_ACTIVE:
		return "8221"
	case ERROR_NO_PROMOTION_ACTIVE:
		return "8222"
	case ERROR_DS_OPERATIONS_ERROR:
		return "8224"
	case ERROR_DS_PROTOCOL_ERROR:
		return "8225"
	case ERROR_DS_TIMELIMIT_EXCEEDED:
		return "8226"
	case ERROR_DS_SIZELIMIT_EXCEEDED:
		return "8227"
	case ERROR_DS_ADMIN_LIMIT_EXCEEDED:
		return "8228"
	case ERROR_DS_COMPARE_FALSE:
		return "8229"
	case ERROR_DS_COMPARE_TRUE:
		return "8230"
	case ERROR_DS_AUTH_METHOD_NOT_SUPPORTED:
		return "8231"
	case ERROR_DS_STRONG_AUTH_REQUIRED:
		return "8232"
	case ERROR_DS_INAPPROPRIATE_AUTH:
		return "8233"
	case ERROR_DS_AUTH_UNKNOWN:
		return "8234"
	case ERROR_DS_REFERRAL:
		return "8235"
	case ERROR_DS_UNAVAILABLE_CRIT_EXTENSION:
		return "8236"
	case ERROR_DS_CONFIDENTIALITY_REQUIRED:
		return "8237"
	case ERROR_DS_INAPPROPRIATE_MATCHING:
		return "8238"
	case ERROR_DS_CONSTRAINT_VIOLATION:
		return "8239"
	case ERROR_DS_NO_SUCH_OBJECT:
		return "8240"
	case ERROR_DS_ALIAS_PROBLEM:
		return "8241"
	case ERROR_DS_INVALID_DN_SYNTAX:
		return "8242"
	case ERROR_DS_IS_LEAF:
		return "8243"
	case ERROR_DS_ALIAS_DEREF_PROBLEM:
		return "8244"
	case ERROR_DS_UNWILLING_TO_PERFORM:
		return "8245"
	case ERROR_DS_LOOP_DETECT:
		return "8246"
	case ERROR_DS_NAMING_VIOLATION:
		return "8247"
	case ERROR_DS_OBJECT_RESULTS_TOO_LARGE:
		return "8248"
	case ERROR_DS_AFFECTS_MULTIPLE_DSAS:
		return "8249"
	case ERROR_DS_SERVER_DOWN:
		return "8250"
	case ERROR_DS_LOCAL_ERROR:
		return "8251"
	case ERROR_DS_ENCODING_ERROR:
		return "8252"
	case ERROR_DS_DECODING_ERROR:
		return "8253"
	case ERROR_DS_FILTER_UNKNOWN:
		return "8254"
	case ERROR_DS_PARAM_ERROR:
		return "8255"
	case ERROR_DS_NOT_SUPPORTED:
		return "8256"
	case ERROR_DS_NO_RESULTS_RETURNED:
		return "8257"
	case ERROR_DS_CONTROL_NOT_FOUND:
		return "8258"
	case ERROR_DS_CLIENT_LOOP:
		return "8259"
	case ERROR_DS_REFERRAL_LIMIT_EXCEEDED:
		return "8260"
	case ERROR_DS_SORT_CONTROL_MISSING:
		return "8261"
	case ERROR_DS_OFFSET_RANGE_ERROR:
		return "8262"
	case ERROR_DS_ROOT_MUST_BE_NC:
		return "8301"
	case ERROR_DS_ADD_REPLICA_INHIBITED:
		return "8302"
	case ERROR_DS_ATT_NOT_DEF_IN_SCHEMA:
		return "8303"
	case ERROR_DS_MAX_OBJ_SIZE_EXCEEDED:
		return "8304"
	case ERROR_DS_OBJ_STRING_NAME_EXISTS:
		return "8305"
	case ERROR_DS_NO_RDN_DEFINED_IN_SCHEMA:
		return "8306"
	case ERROR_DS_RDN_DOESNT_MATCH_SCHEMA:
		return "8307"
	case ERROR_DS_NO_REQUESTED_ATTS_FOUND:
		return "8308"
	case ERROR_DS_USER_BUFFER_TO_SMALL:
		return "8309"
	case ERROR_DS_ATT_IS_NOT_ON_OBJ:
		return "8310"
	case ERROR_DS_ILLEGAL_MOD_OPERATION:
		return "8311"
	case ERROR_DS_OBJ_TOO_LARGE:
		return "8312"
	case ERROR_DS_BAD_INSTANCE_TYPE:
		return "8313"
	case ERROR_DS_MASTERDSA_REQUIRED:
		return "8314"
	case ERROR_DS_OBJECT_CLASS_REQUIRED:
		return "8315"
	case ERROR_DS_MISSING_REQUIRED_ATT:
		return "8316"
	case ERROR_DS_ATT_NOT_DEF_FOR_CLASS:
		return "8317"
	case ERROR_DS_ATT_ALREADY_EXISTS:
		return "8318"
	case ERROR_DS_CANT_ADD_ATT_VALUES:
		return "8320"
	case ERROR_DS_SINGLE_VALUE_CONSTRAINT:
		return "8321"
	case ERROR_DS_RANGE_CONSTRAINT:
		return "8322"
	case ERROR_DS_ATT_VAL_ALREADY_EXISTS:
		return "8323"
	case ERROR_DS_CANT_REM_MISSING_ATT:
		return "8324"
	case ERROR_DS_CANT_REM_MISSING_ATT_VAL:
		return "8325"
	case ERROR_DS_ROOT_CANT_BE_SUBREF:
		return "8326"
	case ERROR_DS_NO_CHAINING:
		return "8327"
	case ERROR_DS_NO_CHAINED_EVAL:
		return "8328"
	case ERROR_DS_NO_PARENT_OBJECT:
		return "8329"
	case ERROR_DS_PARENT_IS_AN_ALIAS:
		return "8330"
	case ERROR_DS_CANT_MIX_MASTER_AND_REPS:
		return "8331"
	case ERROR_DS_CHILDREN_EXIST:
		return "8332"
	case ERROR_DS_OBJ_NOT_FOUND:
		return "8333"
	case ERROR_DS_ALIASED_OBJ_MISSING:
		return "8334"
	case ERROR_DS_BAD_NAME_SYNTAX:
		return "8335"
	case ERROR_DS_ALIAS_POINTS_TO_ALIAS:
		return "8336"
	case ERROR_DS_CANT_DEREF_ALIAS:
		return "8337"
	case ERROR_DS_OUT_OF_SCOPE:
		return "8338"
	case ERROR_DS_OBJECT_BEING_REMOVED:
		return "8339"
	case ERROR_DS_CANT_DELETE_DSA_OBJ:
		return "8340"
	case ERROR_DS_GENERIC_ERROR:
		return "8341"
	case ERROR_DS_DSA_MUST_BE_INT_MASTER:
		return "8342"
	case ERROR_DS_CLASS_NOT_DSA:
		return "8343"
	case ERROR_DS_INSUFF_ACCESS_RIGHTS:
		return "8344"
	case ERROR_DS_ILLEGAL_SUPERIOR:
		return "8345"
	case ERROR_DS_ATTRIBUTE_OWNED_BY_SAM:
		return "8346"
	case ERROR_DS_NAME_TOO_MANY_PARTS:
		return "8347"
	case ERROR_DS_NAME_TOO_LONG:
		return "8348"
	case ERROR_DS_NAME_VALUE_TOO_LONG:
		return "8349"
	case ERROR_DS_NAME_UNPARSEABLE:
		return "8350"
	case ERROR_DS_NAME_TYPE_UNKNOWN:
		return "8351"
	case ERROR_DS_NOT_AN_OBJECT:
		return "8352"
	case ERROR_DS_SEC_DESC_TOO_SHORT:
		return "8353"
	case ERROR_DS_SEC_DESC_INVALID:
		return "8354"
	case ERROR_DS_NO_DELETED_NAME:
		return "8355"
	case ERROR_DS_SUBREF_MUST_HAVE_PARENT:
		return "8356"
	case ERROR_DS_NCNAME_MUST_BE_NC:
		return "8357"
	case ERROR_DS_CANT_ADD_SYSTEM_ONLY:
		return "8358"
	case ERROR_DS_CLASS_MUST_BE_CONCRETE:
		return "8359"
	case ERROR_DS_INVALID_DMD:
		return "8360"
	case ERROR_DS_OBJ_GUID_EXISTS:
		return "8361"
	case ERROR_DS_NOT_ON_BACKLINK:
		return "8362"
	case ERROR_DS_NO_CROSSREF_FOR_NC:
		return "8363"
	case ERROR_DS_SHUTTING_DOWN:
		return "8364"
	case ERROR_DS_UNKNOWN_OPERATION:
		return "8365"
	case ERROR_DS_INVALID_ROLE_OWNER:
		return "8366"
	case ERROR_DS_COULDNT_CONTACT_FSMO:
		return "8367"
	case ERROR_DS_CROSS_NC_DN_RENAME:
		return "8368"
	case ERROR_DS_CANT_MOD_SYSTEM_ONLY:
		return "8369"
	case ERROR_DS_REPLICATOR_ONLY:
		return "8370"
	case ERROR_DS_OBJ_CLASS_NOT_DEFINED:
		return "8371"
	case ERROR_DS_OBJ_CLASS_NOT_SUBCLASS:
		return "8372"
	case ERROR_DS_NAME_REFERENCE_INVALID:
		return "8373"
	case ERROR_DS_CROSS_REF_EXISTS:
		return "8374"
	case ERROR_DS_CANT_DEL_MASTER_CROSSREF:
		return "8375"
	case ERROR_DS_SUBTREE_NOTIFY_NOT_NC_HEAD:
		return "8376"
	case ERROR_DS_NOTIFY_FILTER_TOO_COMPLEX:
		return "8377"
	case ERROR_DS_DUP_RDN:
		return "8378"
	case ERROR_DS_DUP_OID:
		return "8379"
	case ERROR_DS_DUP_MAPI_ID:
		return "8380"
	case ERROR_DS_DUP_SCHEMA_ID_GUID:
		return "8381"
	case ERROR_DS_DUP_LDAP_DISPLAY_NAME:
		return "8382"
	case ERROR_DS_SEMANTIC_ATT_TEST:
		return "8383"
	case ERROR_DS_SYNTAX_MISMATCH:
		return "8384"
	case ERROR_DS_EXISTS_IN_MUST_HAVE:
		return "8385"
	case ERROR_DS_EXISTS_IN_MAY_HAVE:
		return "8386"
	case ERROR_DS_NONEXISTENT_MAY_HAVE:
		return "8387"
	case ERROR_DS_NONEXISTENT_MUST_HAVE:
		return "8388"
	case ERROR_DS_AUX_CLS_TEST_FAIL:
		return "8389"
	case ERROR_DS_NONEXISTENT_POSS_SUP:
		return "8390"
	case ERROR_DS_SUB_CLS_TEST_FAIL:
		return "8391"
	case ERROR_DS_BAD_RDN_ATT_ID_SYNTAX:
		return "8392"
	case ERROR_DS_EXISTS_IN_AUX_CLS:
		return "8393"
	case ERROR_DS_EXISTS_IN_SUB_CLS:
		return "8394"
	case ERROR_DS_EXISTS_IN_POSS_SUP:
		return "8395"
	case ERROR_DS_RECALCSCHEMA_FAILED:
		return "8396"
	case ERROR_DS_TREE_DELETE_NOT_FINISHED:
		return "8397"
	case ERROR_DS_CANT_DELETE:
		return "8398"
	case ERROR_DS_ATT_SCHEMA_REQ_ID:
		return "8399"
	case ERROR_DS_BAD_ATT_SCHEMA_SYNTAX:
		return "8400"
	case ERROR_DS_CANT_CACHE_ATT:
		return "8401"
	case ERROR_DS_CANT_CACHE_CLASS:
		return "8402"
	case ERROR_DS_CANT_REMOVE_ATT_CACHE:
		return "8403"
	case ERROR_DS_CANT_REMOVE_CLASS_CACHE:
		return "8404"
	case ERROR_DS_CANT_RETRIEVE_DN:
		return "8405"
	case ERROR_DS_MISSING_SUPREF:
		return "8406"
	case ERROR_DS_CANT_RETRIEVE_INSTANCE:
		return "8407"
	case ERROR_DS_CODE_INCONSISTENCY:
		return "8408"
	case ERROR_DS_DATABASE_ERROR:
		return "8409"
	case ERROR_DS_GOVERNSID_MISSING:
		return "8410"
	case ERROR_DS_MISSING_EXPECTED_ATT:
		return "8411"
	case ERROR_DS_NCNAME_MISSING_CR_REF:
		return "8412"
	case ERROR_DS_SECURITY_CHECKING_ERROR:
		return "8413"
	case ERROR_DS_SCHEMA_NOT_LOADED:
		return "8414"
	case ERROR_DS_SCHEMA_ALLOC_FAILED:
		return "8415"
	case ERROR_DS_ATT_SCHEMA_REQ_SYNTAX:
		return "8416"
	case ERROR_DS_GCVERIFY_ERROR:
		return "8417"
	case ERROR_DS_DRA_SCHEMA_MISMATCH:
		return "8418"
	case ERROR_DS_CANT_FIND_DSA_OBJ:
		return "8419"
	case ERROR_DS_CANT_FIND_EXPECTED_NC:
		return "8420"
	case ERROR_DS_CANT_FIND_NC_IN_CACHE:
		return "8421"
	case ERROR_DS_CANT_RETRIEVE_CHILD:
		return "8422"
	case ERROR_DS_SECURITY_ILLEGAL_MODIFY:
		return "8423"
	case ERROR_DS_CANT_REPLACE_HIDDEN_REC:
		return "8424"
	case ERROR_DS_BAD_HIERARCHY_FILE:
		return "8425"
	case ERROR_DS_BUILD_HIERARCHY_TABLE_FAILED:
		return "8426"
	case ERROR_DS_CONFIG_PARAM_MISSING:
		return "8427"
	case ERROR_DS_COUNTING_AB_INDICES_FAILED:
		return "8428"
	case ERROR_DS_HIERARCHY_TABLE_MALLOC_FAILED:
		return "8429"
	case ERROR_DS_INTERNAL_FAILURE:
		return "8430"
	case ERROR_DS_UNKNOWN_ERROR:
		return "8431"
	case ERROR_DS_ROOT_REQUIRES_CLASS_TOP:
		return "8432"
	case ERROR_DS_REFUSING_FSMO_ROLES:
		return "8433"
	case ERROR_DS_MISSING_FSMO_SETTINGS:
		return "8434"
	case ERROR_DS_UNABLE_TO_SURRENDER_ROLES:
		return "8435"
	case ERROR_DS_DRA_GENERIC:
		return "8436"
	case ERROR_DS_DRA_INVALID_PARAMETER:
		return "8437"
	case ERROR_DS_DRA_BUSY:
		return "8438"
	case ERROR_DS_DRA_BAD_DN:
		return "8439"
	case ERROR_DS_DRA_BAD_NC:
		return "8440"
	case ERROR_DS_DRA_DN_EXISTS:
		return "8441"
	case ERROR_DS_DRA_INTERNAL_ERROR:
		return "8442"
	case ERROR_DS_DRA_INCONSISTENT_DIT:
		return "8443"
	case ERROR_DS_DRA_CONNECTION_FAILED:
		return "8444"
	case ERROR_DS_DRA_BAD_INSTANCE_TYPE:
		return "8445"
	case ERROR_DS_DRA_OUT_OF_MEM:
		return "8446"
	case ERROR_DS_DRA_MAIL_PROBLEM:
		return "8447"
	case ERROR_DS_DRA_REF_ALREADY_EXISTS:
		return "8448"
	case ERROR_DS_DRA_REF_NOT_FOUND:
		return "8449"
	case ERROR_DS_DRA_OBJ_IS_REP_SOURCE:
		return "8450"
	case ERROR_DS_DRA_DB_ERROR:
		return "8451"
	case ERROR_DS_DRA_NO_REPLICA:
		return "8452"
	case ERROR_DS_DRA_ACCESS_DENIED:
		return "8453"
	case ERROR_DS_DRA_NOT_SUPPORTED:
		return "8454"
	case ERROR_DS_DRA_RPC_CANCELLED:
		return "8455"
	case ERROR_DS_DRA_SOURCE_DISABLED:
		return "8456"
	case ERROR_DS_DRA_SINK_DISABLED:
		return "8457"
	case ERROR_DS_DRA_NAME_COLLISION:
		return "8458"
	case ERROR_DS_DRA_SOURCE_REINSTALLED:
		return "8459"
	case ERROR_DS_DRA_MISSING_PARENT:
		return "8460"
	case ERROR_DS_DRA_PREEMPTED:
		return "8461"
	case ERROR_DS_DRA_ABANDON_SYNC:
		return "8462"
	case ERROR_DS_DRA_SHUTDOWN:
		return "8463"
	case ERROR_DS_DRA_INCOMPATIBLE_PARTIAL_SET:
		return "8464"
	case ERROR_DS_DRA_SOURCE_IS_PARTIAL_REPLICA:
		return "8465"
	case ERROR_DS_DRA_EXTN_CONNECTION_FAILED:
		return "8466"
	case ERROR_DS_INSTALL_SCHEMA_MISMATCH:
		return "8467"
	case ERROR_DS_DUP_LINK_ID:
		return "8468"
	case ERROR_DS_NAME_ERROR_RESOLVING:
		return "8469"
	case ERROR_DS_NAME_ERROR_NOT_FOUND:
		return "8470"
	case ERROR_DS_NAME_ERROR_NOT_UNIQUE:
		return "8471"
	case ERROR_DS_NAME_ERROR_NO_MAPPING:
		return "8472"
	case ERROR_DS_NAME_ERROR_DOMAIN_ONLY:
		return "8473"
	case ERROR_DS_NAME_ERROR_NO_SYNTACTICAL_MAPPING:
		return "8474"
	case ERROR_DS_CONSTRUCTED_ATT_MOD:
		return "8475"
	case ERROR_DS_WRONG_OM_OBJ_CLASS:
		return "8476"
	case ERROR_DS_DRA_REPL_PENDING:
		return "8477"
	case ERROR_DS_DS_REQUIRED:
		return "8478"
	case ERROR_DS_INVALID_LDAP_DISPLAY_NAME:
		return "8479"
	case ERROR_DS_NON_BASE_SEARCH:
		return "8480"
	case ERROR_DS_CANT_RETRIEVE_ATTS:
		return "8481"
	case ERROR_DS_BACKLINK_WITHOUT_LINK:
		return "8482"
	case ERROR_DS_EPOCH_MISMATCH:
		return "8483"
	case ERROR_DS_SRC_NAME_MISMATCH:
		return "8484"
	case ERROR_DS_SRC_AND_DST_NC_IDENTICAL:
		return "8485"
	case ERROR_DS_DST_NC_MISMATCH:
		return "8486"
	case ERROR_DS_NOT_AUTHORITIVE_FOR_DST_NC:
		return "8487"
	case ERROR_DS_SRC_GUID_MISMATCH:
		return "8488"
	case ERROR_DS_CANT_MOVE_DELETED_OBJECT:
		return "8489"
	case ERROR_DS_PDC_OPERATION_IN_PROGRESS:
		return "8490"
	case ERROR_DS_CROSS_DOMAIN_CLEANUP_REQD:
		return "8491"
	case ERROR_DS_ILLEGAL_XDOM_MOVE_OPERATION:
		return "8492"
	case ERROR_DS_CANT_WITH_ACCT_GROUP_MEMBERSHPS:
		return "8493"
	case ERROR_DS_NC_MUST_HAVE_NC_PARENT:
		return "8494"
	case ERROR_DS_CR_IMPOSSIBLE_TO_VALIDATE:
		return "8495"
	case ERROR_DS_DST_DOMAIN_NOT_NATIVE:
		return "8496"
	case ERROR_DS_MISSING_INFRASTRUCTURE_CONTAINER:
		return "8497"
	case ERROR_DS_CANT_MOVE_ACCOUNT_GROUP:
		return "8498"
	case ERROR_DS_CANT_MOVE_RESOURCE_GROUP:
		return "8499"
	case ERROR_DS_INVALID_SEARCH_FLAG:
		return "8500"
	case ERROR_DS_NO_TREE_DELETE_ABOVE_NC:
		return "8501"
	case ERROR_DS_COULDNT_LOCK_TREE_FOR_DELETE:
		return "8502"
	case ERROR_DS_COULDNT_IDENTIFY_OBJECTS_FOR_TREE_DELETE:
		return "8503"
	case ERROR_DS_SAM_INIT_FAILURE:
		return "8504"
	case ERROR_DS_SENSITIVE_GROUP_VIOLATION:
		return "8505"
	case ERROR_DS_CANT_MOD_PRIMARYGROUPID:
		return "8506"
	case ERROR_DS_ILLEGAL_BASE_SCHEMA_MOD:
		return "8507"
	case ERROR_DS_NONSAFE_SCHEMA_CHANGE:
		return "8508"
	case ERROR_DS_SCHEMA_UPDATE_DISALLOWED:
		return "8509"
	case ERROR_DS_CANT_CREATE_UNDER_SCHEMA:
		return "8510"
	case ERROR_DS_INSTALL_NO_SRC_SCH_VERSION:
		return "8511"
	case ERROR_DS_INSTALL_NO_SCH_VERSION_IN_INIFILE:
		return "8512"
	case ERROR_DS_INVALID_GROUP_TYPE:
		return "8513"
	case ERROR_DS_NO_NEST_GLOBALGROUP_IN_MIXEDDOMAIN:
		return "8514"
	case ERROR_DS_NO_NEST_LOCALGROUP_IN_MIXEDDOMAIN:
		return "8515"
	case ERROR_DS_GLOBAL_CANT_HAVE_LOCAL_MEMBER:
		return "8516"
	case ERROR_DS_GLOBAL_CANT_HAVE_UNIVERSAL_MEMBER:
		return "8517"
	case ERROR_DS_UNIVERSAL_CANT_HAVE_LOCAL_MEMBER:
		return "8518"
	case ERROR_DS_GLOBAL_CANT_HAVE_CROSSDOMAIN_MEMBER:
		return "8519"
	case ERROR_DS_LOCAL_CANT_HAVE_CROSSDOMAIN_LOCAL_MEMBER:
		return "8520"
	case ERROR_DS_HAVE_PRIMARY_MEMBERS:
		return "8521"
	case ERROR_DS_STRING_SD_CONVERSION_FAILED:
		return "8522"
	case ERROR_DS_NAMING_MASTER_GC:
		return "8523"
	case ERROR_DS_DNS_LOOKUP_FAILURE:
		return "8524"
	case ERROR_DS_COULDNT_UPDATE_SPNS:
		return "8525"
	case ERROR_DS_CANT_RETRIEVE_SD:
		return "8526"
	case ERROR_DS_KEY_NOT_UNIQUE:
		return "8527"
	case ERROR_DS_WRONG_LINKED_ATT_SYNTAX:
		return "8528"
	case ERROR_DS_SAM_NEED_BOOTKEY_PASSWORD:
		return "8529"
	case ERROR_DS_SAM_NEED_BOOTKEY_FLOPPY:
		return "8530"
	case ERROR_DS_CANT_START:
		return "8531"
	case ERROR_DS_INIT_FAILURE:
		return "8532"
	case ERROR_DS_NO_PKT_PRIVACY_ON_CONNECTION:
		return "8533"
	case ERROR_DS_SOURCE_DOMAIN_IN_FOREST:
		return "8534"
	case ERROR_DS_DESTINATION_DOMAIN_NOT_IN_FOREST:
		return "8535"
	case ERROR_DS_DESTINATION_AUDITING_NOT_ENABLED:
		return "8536"
	case ERROR_DS_CANT_FIND_DC_FOR_SRC_DOMAIN:
		return "8537"
	case ERROR_DS_SRC_OBJ_NOT_GROUP_OR_USER:
		return "8538"
	case ERROR_DS_SRC_SID_EXISTS_IN_FOREST:
		return "8539"
	case ERROR_DS_SRC_AND_DST_OBJECT_CLASS_MISMATCH:
		return "8540"
	case ERROR_SAM_INIT_FAILURE:
		return "8541"
	case ERROR_DS_DRA_SCHEMA_INFO_SHIP:
		return "8542"
	case ERROR_DS_DRA_SCHEMA_CONFLICT:
		return "8543"
	case ERROR_DS_DRA_EARLIER_SCHEMA_CONFLICT:
		return "8544"
	case ERROR_DS_DRA_OBJ_NC_MISMATCH:
		return "8545"
	case ERROR_DS_NC_STILL_HAS_DSAS:
		return "8546"
	case ERROR_DS_GC_REQUIRED:
		return "8547"
	case ERROR_DS_LOCAL_MEMBER_OF_LOCAL_ONLY:
		return "8548"
	case ERROR_DS_NO_FPO_IN_UNIVERSAL_GROUPS:
		return "8549"
	case ERROR_DS_CANT_ADD_TO_GC:
		return "8550"
	case ERROR_DS_NO_CHECKPOINT_WITH_PDC:
		return "8551"
	case ERROR_DS_SOURCE_AUDITING_NOT_ENABLED:
		return "8552"
	case ERROR_DS_CANT_CREATE_IN_NONDOMAIN_NC:
		return "8553"
	case ERROR_DS_INVALID_NAME_FOR_SPN:
		return "8554"
	case ERROR_DS_FILTER_USES_CONTRUCTED_ATTRS:
		return "8555"
	case ERROR_DS_UNICODEPWD_NOT_IN_QUOTES:
		return "8556"
	case ERROR_DS_MACHINE_ACCOUNT_QUOTA_EXCEEDED:
		return "8557"
	case ERROR_DS_MUST_BE_RUN_ON_DST_DC:
		return "8558"
	case ERROR_DS_SRC_DC_MUST_BE_SP4_OR_GREATER:
		return "8559"
	case ERROR_DS_CANT_TREE_DELETE_CRITICAL_OBJ:
		return "8560"
	case ERROR_DS_INIT_FAILURE_CONSOLE:
		return "8561"
	case ERROR_DS_SAM_INIT_FAILURE_CONSOLE:
		return "8562"
	case ERROR_DS_FOREST_VERSION_TOO_HIGH:
		return "8563"
	case ERROR_DS_DOMAIN_VERSION_TOO_HIGH:
		return "8564"
	case ERROR_DS_FOREST_VERSION_TOO_LOW:
		return "8565"
	case ERROR_DS_DOMAIN_VERSION_TOO_LOW:
		return "8566"
	case ERROR_DS_INCOMPATIBLE_VERSION:
		return "8567"
	case ERROR_DS_LOW_DSA_VERSION:
		return "8568"
	case ERROR_DS_NO_BEHAVIOR_VERSION_IN_MIXEDDOMAIN:
		return "8569"
	case ERROR_DS_NOT_SUPPORTED_SORT_ORDER:
		return "8570"
	case ERROR_DS_NAME_NOT_UNIQUE:
		return "8571"
	case ERROR_DS_MACHINE_ACCOUNT_CREATED_PRENT4:
		return "8572"
	case ERROR_DS_OUT_OF_VERSION_STORE:
		return "8573"
	case ERROR_DS_INCOMPATIBLE_CONTROLS_USED:
		return "8574"
	case ERROR_DS_NO_REF_DOMAIN:
		return "8575"
	case ERROR_DS_RESERVED_LINK_ID:
		return "8576"
	case ERROR_DS_LINK_ID_NOT_AVAILABLE:
		return "8577"
	case ERROR_DS_AG_CANT_HAVE_UNIVERSAL_MEMBER:
		return "8578"
	case ERROR_DS_MODIFYDN_DISALLOWED_BY_INSTANCE_TYPE:
		return "8579"
	case ERROR_DS_NO_OBJECT_MOVE_IN_SCHEMA_NC:
		return "8580"
	case ERROR_DS_MODIFYDN_DISALLOWED_BY_FLAG:
		return "8581"
	case ERROR_DS_MODIFYDN_WRONG_GRANDPARENT:
		return "8582"
	case ERROR_DS_NAME_ERROR_TRUST_REFERRAL:
		return "8583"
	case ERROR_NOT_SUPPORTED_ON_STANDARD_SERVER:
		return "8584"
	case ERROR_DS_CANT_ACCESS_REMOTE_PART_OF_AD:
		return "8585"
	case ERROR_DS_CR_IMPOSSIBLE_TO_VALIDATE_V2:
		return "8586"
	case ERROR_DS_THREAD_LIMIT_EXCEEDED:
		return "8587"
	case ERROR_DS_NOT_CLOSEST:
		return "8588"
	case ERROR_DS_CANT_DERIVE_SPN_WITHOUT_SERVER_REF:
		return "8589"
	case ERROR_DS_SINGLE_USER_MODE_FAILED:
		return "8590"
	case ERROR_DS_NTDSCRIPT_SYNTAX_ERROR:
		return "8591"
	case ERROR_DS_NTDSCRIPT_PROCESS_ERROR:
		return "8592"
	case ERROR_DS_DIFFERENT_REPL_EPOCHS:
		return "8593"
	case ERROR_DS_DRS_EXTENSIONS_CHANGED:
		return "8594"
	case ERROR_DS_REPLICA_SET_CHANGE_NOT_ALLOWED_ON_DISABLED_CR:
		return "8595"
	case ERROR_DS_NO_MSDS_INTID:
		return "8596"
	case ERROR_DS_DUP_MSDS_INTID:
		return "8597"
	case ERROR_DS_EXISTS_IN_RDNATTID:
		return "8598"
	case ERROR_DS_AUTHORIZATION_FAILED:
		return "8599"
	case ERROR_DS_INVALID_SCRIPT:
		return "8600"
	case ERROR_DS_REMOTE_CROSSREF_OP_FAILED:
		return "8601"
	case ERROR_DS_CROSS_REF_BUSY:
		return "8602"
	case ERROR_DS_CANT_DERIVE_SPN_FOR_DELETED_DOMAIN:
		return "8603"
	case ERROR_DS_CANT_DEMOTE_WITH_WRITEABLE_NC:
		return "8604"
	case ERROR_DS_DUPLICATE_ID_FOUND:
		return "8605"
	case ERROR_DS_INSUFFICIENT_ATTR_TO_CREATE_OBJECT:
		return "8606"
	case ERROR_DS_GROUP_CONVERSION_ERROR:
		return "8607"
	case ERROR_DS_CANT_MOVE_APP_BASIC_GROUP:
		return "8608"
	case ERROR_DS_CANT_MOVE_APP_QUERY_GROUP:
		return "8609"
	case ERROR_DS_ROLE_NOT_VERIFIED:
		return "8610"
	case ERROR_DS_WKO_CONTAINER_CANNOT_BE_SPECIAL:
		return "8611"
	case ERROR_DS_DOMAIN_RENAME_IN_PROGRESS:
		return "8612"
	case ERROR_DS_EXISTING_AD_CHILD_NC:
		return "8613"
	case ERROR_DS_REPL_LIFETIME_EXCEEDED:
		return "8614"
	case ERROR_DS_DISALLOWED_IN_SYSTEM_CONTAINER:
		return "8615"
	case ERROR_DS_LDAP_SEND_QUEUE_FULL:
		return "8616"
	case ERROR_DS_DRA_OUT_SCHEDULE_WINDOW:
		return "8617"
	case ERROR_DS_POLICY_NOT_KNOWN:
		return "8618"
	case ERROR_NO_SITE_SETTINGS_OBJECT:
		return "8619"
	case ERROR_NO_SECRETS:
		return "8620"
	case ERROR_NO_WRITABLE_DC_FOUND:
		return "8621"
	case ERROR_DS_NO_SERVER_OBJECT:
		return "8622"
	case ERROR_DS_NO_NTDSA_OBJECT:
		return "8623"
	case ERROR_DS_NON_ASQ_SEARCH:
		return "8624"
	case ERROR_DS_AUDIT_FAILURE:
		return "8625"
	case ERROR_DS_INVALID_SEARCH_FLAG_SUBTREE:
		return "8626"
	case ERROR_DS_INVALID_SEARCH_FLAG_TUPLE:
		return "8627"
	case ERROR_DS_HIERARCHY_TABLE_TOO_DEEP:
		return "8628"
	case ERROR_DS_DRA_CORRUPT_UTD_VECTOR:
		return "8629"
	case ERROR_DS_DRA_SECRETS_DENIED:
		return "8630"
	case ERROR_DS_RESERVED_MAPI_ID:
		return "8631"
	case ERROR_DS_MAPI_ID_NOT_AVAILABLE:
		return "8632"
	case ERROR_DS_DRA_MISSING_KRBTGT_SECRET:
		return "8633"
	case ERROR_DS_DOMAIN_NAME_EXISTS_IN_FOREST:
		return "8634"
	case ERROR_DS_FLAT_NAME_EXISTS_IN_FOREST:
		return "8635"
	case ERROR_INVALID_USER_PRINCIPAL_NAME:
		return "8636"
	case ERROR_DS_OID_MAPPED_GROUP_CANT_HAVE_MEMBERS:
		return "8637"
	case ERROR_DS_OID_NOT_FOUND:
		return "8638"
	case ERROR_DS_DRA_RECYCLED_TARGET:
		return "8639"
	case DNS_ERROR_RESPONSE_CODES_BASE:
		return "9000"
	case DNS_ERROR_MASK:
		return "0x00002328"
	case DNS_ERROR_RCODE_FORMAT_ERROR:
		return "9001"
	case DNS_ERROR_RCODE_SERVER_FAILURE:
		return "9002"
	case DNS_ERROR_RCODE_NAME_ERROR:
		return "9003"
	case DNS_ERROR_RCODE_NOT_IMPLEMENTED:
		return "9004"
	case DNS_ERROR_RCODE_REFUSED:
		return "9005"
	case DNS_ERROR_RCODE_YXDOMAIN:
		return "9006"
	case DNS_ERROR_RCODE_YXRRSET:
		return "9007"
	case DNS_ERROR_RCODE_NXRRSET:
		return "9008"
	case DNS_ERROR_RCODE_NOTAUTH:
		return "9009"
	case DNS_ERROR_RCODE_NOTZONE:
		return "9010"
	case DNS_ERROR_RCODE_BADSIG:
		return "9016"
	case DNS_ERROR_RCODE_BADKEY:
		return "9017"
	case DNS_ERROR_RCODE_BADTIME:
		return "9018"
	case DNS_ERROR_PACKET_FMT_BASE:
		return "9500"
	case DNS_INFO_NO_RECORDS:
		return "9501"
	case DNS_ERROR_BAD_PACKET:
		return "9502"
	case DNS_ERROR_NO_PACKET:
		return "9503"
	case DNS_ERROR_RCODE:
		return "9504"
	case DNS_ERROR_UNSECURE_PACKET:
		return "9505"
	case DNS_ERROR_GENERAL_API_BASE:
		return "9550"
	case DNS_ERROR_INVALID_TYPE:
		return "9551"
	case DNS_ERROR_INVALID_IP_ADDRESS:
		return "9552"
	case DNS_ERROR_INVALID_PROPERTY:
		return "9553"
	case DNS_ERROR_TRY_AGAIN_LATER:
		return "9554"
	case DNS_ERROR_NOT_UNIQUE:
		return "9555"
	case DNS_ERROR_NON_RFC_NAME:
		return "9556"
	case DNS_STATUS_FQDN:
		return "9557"
	case DNS_STATUS_DOTTED_NAME:
		return "9558"
	case DNS_STATUS_SINGLE_PART_NAME:
		return "9559"
	case DNS_ERROR_INVALID_NAME_CHAR:
		return "9560"
	case DNS_ERROR_NUMERIC_NAME:
		return "9561"
	case DNS_ERROR_NOT_ALLOWED_ON_ROOT_SERVER:
		return "9562"
	case DNS_ERROR_NOT_ALLOWED_UNDER_DELEGATION:
		return "9563"
	case DNS_ERROR_CANNOT_FIND_ROOT_HINTS:
		return "9564"
	case DNS_ERROR_INCONSISTENT_ROOT_HINTS:
		return "9565"
	case DNS_ERROR_DWORD_VALUE_TOO_SMALL:
		return "9566"
	case DNS_ERROR_DWORD_VALUE_TOO_LARGE:
		return "9567"
	case DNS_ERROR_BACKGROUND_LOADING:
		return "9568"
	case DNS_ERROR_NOT_ALLOWED_ON_RODC:
		return "9569"
	case DNS_ERROR_NOT_ALLOWED_UNDER_DNAME:
		return "9570"
	case DNS_ERROR_DELEGATION_REQUIRED:
		return "9571"
	case DNS_ERROR_INVALID_POLICY_TABLE:
		return "9572"
	case DNS_ERROR_ZONE_BASE:
		return "9600"
	case DNS_ERROR_ZONE_DOES_NOT_EXIST:
		return "9601"
	case DNS_ERROR_NO_ZONE_INFO:
		return "9602"
	case DNS_ERROR_INVALID_ZONE_OPERATION:
		return "9603"
	case DNS_ERROR_ZONE_CONFIGURATION_ERROR:
		return "9604"
	case DNS_ERROR_ZONE_HAS_NO_SOA_RECORD:
		return "9605"
	case DNS_ERROR_ZONE_HAS_NO_NS_RECORDS:
		return "9606"
	case DNS_ERROR_ZONE_LOCKED:
		return "9607"
	case DNS_ERROR_ZONE_CREATION_FAILED:
		return "9608"
	case DNS_ERROR_ZONE_ALREADY_EXISTS:
		return "9609"
	case DNS_ERROR_AUTOZONE_ALREADY_EXISTS:
		return "9610"
	case DNS_ERROR_INVALID_ZONE_TYPE:
		return "9611"
	case DNS_ERROR_SECONDARY_REQUIRES_MASTER_IP:
		return "9612"
	case DNS_ERROR_ZONE_NOT_SECONDARY:
		return "9613"
	case DNS_ERROR_NEED_SECONDARY_ADDRESSES:
		return "9614"
	case DNS_ERROR_WINS_INIT_FAILED:
		return "9615"
	case DNS_ERROR_NEED_WINS_SERVERS:
		return "9616"
	case DNS_ERROR_NBSTAT_INIT_FAILED:
		return "9617"
	case DNS_ERROR_SOA_DELETE_INVALID:
		return "9618"
	case DNS_ERROR_FORWARDER_ALREADY_EXISTS:
		return "9619"
	case DNS_ERROR_ZONE_REQUIRES_MASTER_IP:
		return "9620"
	case DNS_ERROR_ZONE_IS_SHUTDOWN:
		return "9621"
	case DNS_ERROR_DATAFILE_BASE:
		return "9650"
	case DNS_ERROR_PRIMARY_REQUIRES_DATAFILE:
		return "9651"
	case DNS_ERROR_INVALID_DATAFILE_NAME:
		return "9652"
	case DNS_ERROR_DATAFILE_OPEN_FAILURE:
		return "9653"
	case DNS_ERROR_FILE_WRITEBACK_FAILED:
		return "9654"
	case DNS_ERROR_DATAFILE_PARSING:
		return "9655"
	case DNS_ERROR_DATABASE_BASE:
		return "9700"
	case DNS_ERROR_RECORD_DOES_NOT_EXIST:
		return "9701"
	case DNS_ERROR_RECORD_FORMAT:
		return "9702"
	case DNS_ERROR_NODE_CREATION_FAILED:
		return "9703"
	case DNS_ERROR_UNKNOWN_RECORD_TYPE:
		return "9704"
	case DNS_ERROR_RECORD_TIMED_OUT:
		return "9705"
	case DNS_ERROR_NAME_NOT_IN_ZONE:
		return "9706"
	case DNS_ERROR_CNAME_LOOP:
		return "9707"
	case DNS_ERROR_NODE_IS_CNAME:
		return "9708"
	case DNS_ERROR_CNAME_COLLISION:
		return "9709"
	case DNS_ERROR_RECORD_ONLY_AT_ZONE_ROOT:
		return "9710"
	case DNS_ERROR_RECORD_ALREADY_EXISTS:
		return "9711"
	case DNS_ERROR_SECONDARY_DATA:
		return "9712"
	case DNS_ERROR_NO_CREATE_CACHE_DATA:
		return "9713"
	case DNS_ERROR_NAME_DOES_NOT_EXIST:
		return "9714"
	case DNS_WARNING_PTR_CREATE_FAILED:
		return "9715"
	case DNS_WARNING_DOMAIN_UNDELETED:
		return "9716"
	case DNS_ERROR_DS_UNAVAILABLE:
		return "9717"
	case DNS_ERROR_DS_ZONE_ALREADY_EXISTS:
		return "9718"
	case DNS_ERROR_NO_BOOTFILE_IF_DS_ZONE:
		return "9719"
	case DNS_ERROR_NODE_IS_DNAME:
		return "9720"
	case DNS_ERROR_DNAME_COLLISION:
		return "9721"
	case DNS_ERROR_ALIAS_LOOP:
		return "9722"
	case DNS_ERROR_OPERATION_BASE:
		return "9750"
	case DNS_INFO_AXFR_COMPLETE:
		return "9751"
	case DNS_ERROR_AXFR:
		return "9752"
	case DNS_INFO_ADDED_LOCAL_WINS:
		return "9753"
	case DNS_ERROR_SECURE_BASE:
		return "9800"
	case DNS_STATUS_CONTINUE_NEEDED:
		return "9801"
	case DNS_ERROR_SETUP_BASE:
		return "9850"
	case DNS_ERROR_NO_TCPIP:
		return "9851"
	case DNS_ERROR_NO_DNS_SERVERS:
		return "9852"
	case DNS_ERROR_DP_BASE:
		return "9900"
	case DNS_ERROR_DP_DOES_NOT_EXIST:
		return "9901"
	case DNS_ERROR_DP_ALREADY_EXISTS:
		return "9902"
	case DNS_ERROR_DP_NOT_ENLISTED:
		return "9903"
	case DNS_ERROR_DP_ALREADY_ENLISTED:
		return "9904"
	case DNS_ERROR_DP_NOT_AVAILABLE:
		return "9905"
	case DNS_ERROR_DP_FSMO_ERROR:
		return "9906"
	case WSABASEERR:
		return "10000"
	case WSAEINTR:
		return "10004"
	case WSAEBADF:
		return "10009"
	case WSAEACCES:
		return "10013"
	case WSAEFAULT:
		return "10014"
	case WSAEINVAL:
		return "10022"
	case WSAEMFILE:
		return "10024"
	case WSAEWOULDBLOCK:
		return "10035"
	case WSAEINPROGRESS:
		return "10036"
	case WSAEALREADY:
		return "10037"
	case WSAENOTSOCK:
		return "10038"
	case WSAEDESTADDRREQ:
		return "10039"
	case WSAEMSGSIZE:
		return "10040"
	case WSAEPROTOTYPE:
		return "10041"
	case WSAENOPROTOOPT:
		return "10042"
	case WSAEPROTONOSUPPORT:
		return "10043"
	case WSAESOCKTNOSUPPORT:
		return "10044"
	case WSAEOPNOTSUPP:
		return "10045"
	case WSAEPFNOSUPPORT:
		return "10046"
	case WSAEAFNOSUPPORT:
		return "10047"
	case WSAEADDRINUSE:
		return "10048"
	case WSAEADDRNOTAVAIL:
		return "10049"
	case WSAENETDOWN:
		return "10050"
	case WSAENETUNREACH:
		return "10051"
	case WSAENETRESET:
		return "10052"
	case WSAECONNABORTED:
		return "10053"
	case WSAECONNRESET:
		return "10054"
	case WSAENOBUFS:
		return "10055"
	case WSAEISCONN:
		return "10056"
	case WSAENOTCONN:
		return "10057"
	case WSAESHUTDOWN:
		return "10058"
	case WSAETOOMANYREFS:
		return "10059"
	case WSAETIMEDOUT:
		return "10060"
	case WSAECONNREFUSED:
		return "10061"
	case WSAELOOP:
		return "10062"
	case WSAENAMETOOLONG:
		return "10063"
	case WSAEHOSTDOWN:
		return "10064"
	case WSAEHOSTUNREACH:
		return "10065"
	case WSAENOTEMPTY:
		return "10066"
	case WSAEPROCLIM:
		return "10067"
	case WSAEUSERS:
		return "10068"
	case WSAEDQUOT:
		return "10069"
	case WSAESTALE:
		return "10070"
	case WSAEREMOTE:
		return "10071"
	case WSASYSNOTREADY:
		return "10091"
	case WSAVERNOTSUPPORTED:
		return "10092"
	case WSANOTINITIALISED:
		return "10093"
	case WSAEDISCON:
		return "10101"
	case WSAENOMORE:
		return "10102"
	case WSAECANCELLED:
		return "10103"
	case WSAEINVALIDPROCTABLE:
		return "10104"
	case WSAEINVALIDPROVIDER:
		return "10105"
	case WSAEPROVIDERFAILEDINIT:
		return "10106"
	case WSASYSCALLFAILURE:
		return "10107"
	case WSASERVICE_NOT_FOUND:
		return "10108"
	case WSATYPE_NOT_FOUND:
		return "10109"
	case WSA_E_NO_MORE:
		return "10110"
	case WSA_E_CANCELLED:
		return "10111"
	case WSAEREFUSED:
		return "10112"
	case WSAHOST_NOT_FOUND:
		return "11001"
	case WSATRY_AGAIN:
		return "11002"
	case WSANO_RECOVERY:
		return "11003"
	case WSANO_DATA:
		return "11004"
	case WSA_QOS_RECEIVERS:
		return "11005"
	case WSA_QOS_SENDERS:
		return "11006"
	case WSA_QOS_NO_SENDERS:
		return "11007"
	case WSA_QOS_NO_RECEIVERS:
		return "11008"
	case WSA_QOS_REQUEST_CONFIRMED:
		return "11009"
	case WSA_QOS_ADMISSION_FAILURE:
		return "11010"
	case WSA_QOS_POLICY_FAILURE:
		return "11011"
	case WSA_QOS_BAD_STYLE:
		return "11012"
	case WSA_QOS_BAD_OBJECT:
		return "11013"
	case WSA_QOS_TRAFFIC_CTRL_ERROR:
		return "11014"
	case WSA_QOS_GENERIC_ERROR:
		return "11015"
	case WSA_QOS_ESERVICETYPE:
		return "11016"
	case WSA_QOS_EFLOWSPEC:
		return "11017"
	case WSA_QOS_EPROVSPECBUF:
		return "11018"
	case WSA_QOS_EFILTERSTYLE:
		return "11019"
	case WSA_QOS_EFILTERTYPE:
		return "11020"
	case WSA_QOS_EFILTERCOUNT:
		return "11021"
	case WSA_QOS_EOBJLENGTH:
		return "11022"
	case WSA_QOS_EFLOWCOUNT:
		return "11023"
	case WSA_QOS_EUNKOWNPSOBJ:
		return "11024"
	case WSA_QOS_EPOLICYOBJ:
		return "11025"
	case WSA_QOS_EFLOWDESC:
		return "11026"
	case WSA_QOS_EPSFLOWSPEC:
		return "11027"
	case WSA_QOS_EPSFILTERSPEC:
		return "11028"
	case WSA_QOS_ESDMODEOBJ:
		return "11029"
	case WSA_QOS_ESHAPERATEOBJ:
		return "11030"
	case WSA_QOS_RESERVED_PETYPE:
		return "11031"
	case WSA_SECURE_HOST_NOT_FOUND:
		return "11032"
	case WSA_IPSEC_NAME_POLICY_ERROR:
		return "11033"
	case ERROR_IPSEC_QM_POLICY_EXISTS:
		return "13000"
	case ERROR_IPSEC_QM_POLICY_NOT_FOUND:
		return "13001"
	case ERROR_IPSEC_QM_POLICY_IN_USE:
		return "13002"
	case ERROR_IPSEC_MM_POLICY_EXISTS:
		return "13003"
	case ERROR_IPSEC_MM_POLICY_NOT_FOUND:
		return "13004"
	case ERROR_IPSEC_MM_POLICY_IN_USE:
		return "13005"
	case ERROR_IPSEC_MM_FILTER_EXISTS:
		return "13006"
	case ERROR_IPSEC_MM_FILTER_NOT_FOUND:
		return "13007"
	case ERROR_IPSEC_TRANSPORT_FILTER_EXISTS:
		return "13008"
	case ERROR_IPSEC_TRANSPORT_FILTER_NOT_FOUND:
		return "13009"
	case ERROR_IPSEC_MM_AUTH_EXISTS:
		return "13010"
	case ERROR_IPSEC_MM_AUTH_NOT_FOUND:
		return "13011"
	case ERROR_IPSEC_MM_AUTH_IN_USE:
		return "13012"
	case ERROR_IPSEC_DEFAULT_MM_POLICY_NOT_FOUND:
		return "13013"
	case ERROR_IPSEC_DEFAULT_MM_AUTH_NOT_FOUND:
		return "13014"
	case ERROR_IPSEC_DEFAULT_QM_POLICY_NOT_FOUND:
		return "13015"
	case ERROR_IPSEC_TUNNEL_FILTER_EXISTS:
		return "13016"
	case ERROR_IPSEC_TUNNEL_FILTER_NOT_FOUND:
		return "13017"
	case ERROR_IPSEC_MM_FILTER_PENDING_DELETION:
		return "13018"
	case ERROR_IPSEC_TRANSPORT_FILTER_PENDING_DELETION:
		return "13019"
	case ERROR_IPSEC_TUNNEL_FILTER_PENDING_DELETION:
		return "13020"
	case ERROR_IPSEC_MM_POLICY_PENDING_DELETION:
		return "13021"
	case ERROR_IPSEC_MM_AUTH_PENDING_DELETION:
		return "13022"
	case ERROR_IPSEC_QM_POLICY_PENDING_DELETION:
		return "13023"
	case WARNING_IPSEC_MM_POLICY_PRUNED:
		return "13024"
	case WARNING_IPSEC_QM_POLICY_PRUNED:
		return "13025"
	case ERROR_IPSEC_IKE_NEG_STATUS_BEGIN:
		return "13800"
	case ERROR_IPSEC_IKE_AUTH_FAIL:
		return "13801"
	case ERROR_IPSEC_IKE_ATTRIB_FAIL:
		return "13802"
	case ERROR_IPSEC_IKE_NEGOTIATION_PENDING:
		return "13803"
	case ERROR_IPSEC_IKE_GENERAL_PROCESSING_ERROR:
		return "13804"
	case ERROR_IPSEC_IKE_TIMED_OUT:
		return "13805"
	case ERROR_IPSEC_IKE_NO_CERT:
		return "13806"
	case ERROR_IPSEC_IKE_SA_DELETED:
		return "13807"
	case ERROR_IPSEC_IKE_SA_REAPED:
		return "13808"
	case ERROR_IPSEC_IKE_MM_ACQUIRE_DROP:
		return "13809"
	case ERROR_IPSEC_IKE_QM_ACQUIRE_DROP:
		return "13810"
	case ERROR_IPSEC_IKE_QUEUE_DROP_MM:
		return "13811"
	case ERROR_IPSEC_IKE_QUEUE_DROP_NO_MM:
		return "13812"
	case ERROR_IPSEC_IKE_DROP_NO_RESPONSE:
		return "13813"
	case ERROR_IPSEC_IKE_MM_DELAY_DROP:
		return "13814"
	case ERROR_IPSEC_IKE_QM_DELAY_DROP:
		return "13815"
	case ERROR_IPSEC_IKE_ERROR:
		return "13816"
	case ERROR_IPSEC_IKE_CRL_FAILED:
		return "13817"
	case ERROR_IPSEC_IKE_INVALID_KEY_USAGE:
		return "13818"
	case ERROR_IPSEC_IKE_INVALID_CERT_TYPE:
		return "13819"
	case ERROR_IPSEC_IKE_NO_PRIVATE_KEY:
		return "13820"
	case ERROR_IPSEC_IKE_SIMULTANEOUS_REKEY:
		return "13821"
	case ERROR_IPSEC_IKE_DH_FAIL:
		return "13822"
	case ERROR_IPSEC_IKE_CRITICAL_PAYLOAD_NOT_RECOGNIZED:
		return "13823"
	case ERROR_IPSEC_IKE_INVALID_HEADER:
		return "13824"
	case ERROR_IPSEC_IKE_NO_POLICY:
		return "13825"
	case ERROR_IPSEC_IKE_INVALID_SIGNATURE:
		return "13826"
	case ERROR_IPSEC_IKE_KERBEROS_ERROR:
		return "13827"
	case ERROR_IPSEC_IKE_NO_PUBLIC_KEY:
		return "13828"
	case ERROR_IPSEC_IKE_PROCESS_ERR:
		return "13829"
	case ERROR_IPSEC_IKE_PROCESS_ERR_SA:
		return "13830"
	case ERROR_IPSEC_IKE_PROCESS_ERR_PROP:
		return "13831"
	case ERROR_IPSEC_IKE_PROCESS_ERR_TRANS:
		return "13832"
	case ERROR_IPSEC_IKE_PROCESS_ERR_KE:
		return "13833"
	case ERROR_IPSEC_IKE_PROCESS_ERR_ID:
		return "13834"
	case ERROR_IPSEC_IKE_PROCESS_ERR_CERT:
		return "13835"
	case ERROR_IPSEC_IKE_PROCESS_ERR_CERT_REQ:
		return "13836"
	case ERROR_IPSEC_IKE_PROCESS_ERR_HASH:
		return "13837"
	case ERROR_IPSEC_IKE_PROCESS_ERR_SIG:
		return "13838"
	case ERROR_IPSEC_IKE_PROCESS_ERR_NONCE:
		return "13839"
	case ERROR_IPSEC_IKE_PROCESS_ERR_NOTIFY:
		return "13840"
	case ERROR_IPSEC_IKE_PROCESS_ERR_DELETE:
		return "13841"
	case ERROR_IPSEC_IKE_PROCESS_ERR_VENDOR:
		return "13842"
	case ERROR_IPSEC_IKE_INVALID_PAYLOAD:
		return "13843"
	case ERROR_IPSEC_IKE_LOAD_SOFT_SA:
		return "13844"
	case ERROR_IPSEC_IKE_SOFT_SA_TORN_DOWN:
		return "13845"
	case ERROR_IPSEC_IKE_INVALID_COOKIE:
		return "13846"
	case ERROR_IPSEC_IKE_NO_PEER_CERT:
		return "13847"
	case ERROR_IPSEC_IKE_PEER_CRL_FAILED:
		return "13848"
	case ERROR_IPSEC_IKE_POLICY_CHANGE:
		return "13849"
	case ERROR_IPSEC_IKE_NO_MM_POLICY:
		return "13850"
	case ERROR_IPSEC_IKE_NOTCBPRIV:
		return "13851"
	case ERROR_IPSEC_IKE_SECLOADFAIL:
		return "13852"
	case ERROR_IPSEC_IKE_FAILSSPINIT:
		return "13853"
	case ERROR_IPSEC_IKE_FAILQUERYSSP:
		return "13854"
	case ERROR_IPSEC_IKE_SRVACQFAIL:
		return "13855"
	case ERROR_IPSEC_IKE_SRVQUERYCRED:
		return "13856"
	case ERROR_IPSEC_IKE_GETSPIFAIL:
		return "13857"
	case ERROR_IPSEC_IKE_INVALID_FILTER:
		return "13858"
	case ERROR_IPSEC_IKE_OUT_OF_MEMORY:
		return "13859"
	case ERROR_IPSEC_IKE_ADD_UPDATE_KEY_FAILED:
		return "13860"
	case ERROR_IPSEC_IKE_INVALID_POLICY:
		return "13861"
	case ERROR_IPSEC_IKE_UNKNOWN_DOI:
		return "13862"
	case ERROR_IPSEC_IKE_INVALID_SITUATION:
		return "13863"
	case ERROR_IPSEC_IKE_DH_FAILURE:
		return "13864"
	case ERROR_IPSEC_IKE_INVALID_GROUP:
		return "13865"
	case ERROR_IPSEC_IKE_ENCRYPT:
		return "13866"
	case ERROR_IPSEC_IKE_DECRYPT:
		return "13867"
	case ERROR_IPSEC_IKE_POLICY_MATCH:
		return "13868"
	case ERROR_IPSEC_IKE_UNSUPPORTED_ID:
		return "13869"
	case ERROR_IPSEC_IKE_INVALID_HASH:
		return "13870"
	case ERROR_IPSEC_IKE_INVALID_HASH_ALG:
		return "13871"
	case ERROR_IPSEC_IKE_INVALID_HASH_SIZE:
		return "13872"
	case ERROR_IPSEC_IKE_INVALID_ENCRYPT_ALG:
		return "13873"
	case ERROR_IPSEC_IKE_INVALID_AUTH_ALG:
		return "13874"
	case ERROR_IPSEC_IKE_INVALID_SIG:
		return "13875"
	case ERROR_IPSEC_IKE_LOAD_FAILED:
		return "13876"
	case ERROR_IPSEC_IKE_RPC_DELETE:
		return "13877"
	case ERROR_IPSEC_IKE_BENIGN_REINIT:
		return "13878"
	case ERROR_IPSEC_IKE_INVALID_RESPONDER_LIFETIME_NOTIFY:
		return "13879"
	case ERROR_IPSEC_IKE_INVALID_MAJOR_VERSION:
		return "13880"
	case ERROR_IPSEC_IKE_INVALID_CERT_KEYLEN:
		return "13881"
	case ERROR_IPSEC_IKE_MM_LIMIT:
		return "13882"
	case ERROR_IPSEC_IKE_NEGOTIATION_DISABLED:
		return "13883"
	case ERROR_IPSEC_IKE_QM_LIMIT:
		return "13884"
	case ERROR_IPSEC_IKE_MM_EXPIRED:
		return "13885"
	case ERROR_IPSEC_IKE_PEER_MM_ASSUMED_INVALID:
		return "13886"
	case ERROR_IPSEC_IKE_CERT_CHAIN_POLICY_MISMATCH:
		return "13887"
	case ERROR_IPSEC_IKE_UNEXPECTED_MESSAGE_ID:
		return "13888"
	case ERROR_IPSEC_IKE_INVALID_AUTH_PAYLOAD:
		return "13889"
	case ERROR_IPSEC_IKE_DOS_COOKIE_SENT:
		return "13890"
	case ERROR_IPSEC_IKE_SHUTTING_DOWN:
		return "13891"
	case ERROR_IPSEC_IKE_CGA_AUTH_FAILED:
		return "13892"
	case ERROR_IPSEC_IKE_PROCESS_ERR_NATOA:
		return "13893"
	case ERROR_IPSEC_IKE_INVALID_MM_FOR_QM:
		return "13894"
	case ERROR_IPSEC_IKE_QM_EXPIRED:
		return "13895"
	case ERROR_IPSEC_IKE_TOO_MANY_FILTERS:
		return "13896"
	case ERROR_IPSEC_IKE_NEG_STATUS_END:
		return "13897"
	case ERROR_IPSEC_IKE_KILL_DUMMY_NAP_TUNNEL:
		return "13898"
	case ERROR_IPSEC_IKE_INNER_IP_ASSIGNMENT_FAILURE:
		return "13899"
	case ERROR_IPSEC_IKE_REQUIRE_CP_PAYLOAD_MISSING:
		return "13900"
	case ERROR_IPSEC_KEY_MODULE_IMPERSONATION_NEGOTIATION_PENDING:
		return "13901"
	case ERROR_IPSEC_IKE_COEXISTENCE_SUPPRESS:
		return "13902"
	case ERROR_IPSEC_IKE_RATELIMIT_DROP:
		return "13903"
	case ERROR_IPSEC_IKE_PEER_DOESNT_SUPPORT_MOBIKE:
		return "13904"
	case ERROR_IPSEC_IKE_AUTHORIZATION_FAILURE:
		return "13905"
	case ERROR_IPSEC_IKE_STRONG_CRED_AUTHORIZATION_FAILURE:
		return "13906"
	case ERROR_IPSEC_IKE_AUTHORIZATION_FAILURE_WITH_OPTIONAL_RETRY:
		return "13907"
	case ERROR_IPSEC_IKE_STRONG_CRED_AUTHORIZATION_AND_CERTMAP_FAILURE:
		return "13908"
	case ERROR_IPSEC_IKE_NEG_STATUS_EXTENDED_END:
		return "13909"
	case ERROR_IPSEC_BAD_SPI:
		return "13910"
	case ERROR_IPSEC_SA_LIFETIME_EXPIRED:
		return "13911"
	case ERROR_IPSEC_WRONG_SA:
		return "13912"
	case ERROR_IPSEC_REPLAY_CHECK_FAILED:
		return "13913"
	case ERROR_IPSEC_INVALID_PACKET:
		return "13914"
	case ERROR_IPSEC_INTEGRITY_CHECK_FAILED:
		return "13915"
	case ERROR_IPSEC_CLEAR_TEXT_DROP:
		return "13916"
	case ERROR_IPSEC_AUTH_FIREWALL_DROP:
		return "13917"
	case ERROR_IPSEC_THROTTLE_DROP:
		return "13918"
	case ERROR_IPSEC_DOSP_BLOCK:
		return "13925"
	case ERROR_IPSEC_DOSP_RECEIVED_MULTICAST:
		return "13926"
	case ERROR_IPSEC_DOSP_INVALID_PACKET:
		return "13927"
	case ERROR_IPSEC_DOSP_STATE_LOOKUP_FAILED:
		return "13928"
	case ERROR_IPSEC_DOSP_MAX_ENTRIES:
		return "13929"
	case ERROR_IPSEC_DOSP_KEYMOD_NOT_ALLOWED:
		return "13930"
	case ERROR_IPSEC_DOSP_NOT_INSTALLED:
		return "13931"
	case ERROR_IPSEC_DOSP_MAX_PER_IP_RATELIMIT_QUEUES:
		return "13932"
	case ERROR_SXS_SECTION_NOT_FOUND:
		return "14000"
	case ERROR_SXS_CANT_GEN_ACTCTX:
		return "14001"
	case ERROR_SXS_INVALID_ACTCTXDATA_FORMAT:
		return "14002"
	case ERROR_SXS_ASSEMBLY_NOT_FOUND:
		return "14003"
	case ERROR_SXS_MANIFEST_FORMAT_ERROR:
		return "14004"
	case ERROR_SXS_MANIFEST_PARSE_ERROR:
		return "14005"
	case ERROR_SXS_ACTIVATION_CONTEXT_DISABLED:
		return "14006"
	case ERROR_SXS_KEY_NOT_FOUND:
		return "14007"
	case ERROR_SXS_VERSION_CONFLICT:
		return "14008"
	case ERROR_SXS_WRONG_SECTION_TYPE:
		return "14009"
	case ERROR_SXS_THREAD_QUERIES_DISABLED:
		return "14010"
	case ERROR_SXS_PROCESS_DEFAULT_ALREADY_SET:
		return "14011"
	case ERROR_SXS_UNKNOWN_ENCODING_GROUP:
		return "14012"
	case ERROR_SXS_UNKNOWN_ENCODING:
		return "14013"
	case ERROR_SXS_INVALID_XML_NAMESPACE_URI:
		return "14014"
	case ERROR_SXS_ROOT_MANIFEST_DEPENDENCY_NOT_INSTALLED:
		return "14015"
	case ERROR_SXS_LEAF_MANIFEST_DEPENDENCY_NOT_INSTALLED:
		return "14016"
	case ERROR_SXS_INVALID_ASSEMBLY_IDENTITY_ATTRIBUTE:
		return "14017"
	case ERROR_SXS_MANIFEST_MISSING_REQUIRED_DEFAULT_NAMESPACE:
		return "14018"
	case ERROR_SXS_MANIFEST_INVALID_REQUIRED_DEFAULT_NAMESPACE:
		return "14019"
	case ERROR_SXS_PRIVATE_MANIFEST_CROSS_PATH_WITH_REPARSE_POINT:
		return "14020"
	case ERROR_SXS_DUPLICATE_DLL_NAME:
		return "14021"
	case ERROR_SXS_DUPLICATE_WINDOWCLASS_NAME:
		return "14022"
	case ERROR_SXS_DUPLICATE_CLSID:
		return "14023"
	case ERROR_SXS_DUPLICATE_IID:
		return "14024"
	case ERROR_SXS_DUPLICATE_TLBID:
		return "14025"
	case ERROR_SXS_DUPLICATE_PROGID:
		return "14026"
	case ERROR_SXS_DUPLICATE_ASSEMBLY_NAME:
		return "14027"
	case ERROR_SXS_FILE_HASH_MISMATCH:
		return "14028"
	case ERROR_SXS_POLICY_PARSE_ERROR:
		return "14029"
	case ERROR_SXS_XML_E_MISSINGQUOTE:
		return "14030"
	case ERROR_SXS_XML_E_COMMENTSYNTAX:
		return "14031"
	case ERROR_SXS_XML_E_BADSTARTNAMECHAR:
		return "14032"
	case ERROR_SXS_XML_E_BADNAMECHAR:
		return "14033"
	case ERROR_SXS_XML_E_BADCHARINSTRING:
		return "14034"
	case ERROR_SXS_XML_E_XMLDECLSYNTAX:
		return "14035"
	case ERROR_SXS_XML_E_BADCHARDATA:
		return "14036"
	case ERROR_SXS_XML_E_MISSINGWHITESPACE:
		return "14037"
	case ERROR_SXS_XML_E_EXPECTINGTAGEND:
		return "14038"
	case ERROR_SXS_XML_E_MISSINGSEMICOLON:
		return "14039"
	case ERROR_SXS_XML_E_UNBALANCEDPAREN:
		return "14040"
	case ERROR_SXS_XML_E_INTERNALERROR:
		return "14041"
	case ERROR_SXS_XML_E_UNEXPECTED_WHITESPACE:
		return "14042"
	case ERROR_SXS_XML_E_INCOMPLETE_ENCODING:
		return "14043"
	case ERROR_SXS_XML_E_MISSING_PAREN:
		return "14044"
	case ERROR_SXS_XML_E_EXPECTINGCLOSEQUOTE:
		return "14045"
	case ERROR_SXS_XML_E_MULTIPLE_COLONS:
		return "14046"
	case ERROR_SXS_XML_E_INVALID_DECIMAL:
		return "14047"
	case ERROR_SXS_XML_E_INVALID_HEXIDECIMAL:
		return "14048"
	case ERROR_SXS_XML_E_INVALID_UNICODE:
		return "14049"
	case ERROR_SXS_XML_E_WHITESPACEORQUESTIONMARK:
		return "14050"
	case ERROR_SXS_XML_E_UNEXPECTEDENDTAG:
		return "14051"
	case ERROR_SXS_XML_E_UNCLOSEDTAG:
		return "14052"
	case ERROR_SXS_XML_E_DUPLICATEATTRIBUTE:
		return "14053"
	case ERROR_SXS_XML_E_MULTIPLEROOTS:
		return "14054"
	case ERROR_SXS_XML_E_INVALIDATROOTLEVEL:
		return "14055"
	case ERROR_SXS_XML_E_BADXMLDECL:
		return "14056"
	case ERROR_SXS_XML_E_MISSINGROOT:
		return "14057"
	case ERROR_SXS_XML_E_UNEXPECTEDEOF:
		return "14058"
	case ERROR_SXS_XML_E_BADPEREFINSUBSET:
		return "14059"
	case ERROR_SXS_XML_E_UNCLOSEDSTARTTAG:
		return "14060"
	case ERROR_SXS_XML_E_UNCLOSEDENDTAG:
		return "14061"
	case ERROR_SXS_XML_E_UNCLOSEDSTRING:
		return "14062"
	case ERROR_SXS_XML_E_UNCLOSEDCOMMENT:
		return "14063"
	case ERROR_SXS_XML_E_UNCLOSEDDECL:
		return "14064"
	case ERROR_SXS_XML_E_UNCLOSEDCDATA:
		return "14065"
	case ERROR_SXS_XML_E_RESERVEDNAMESPACE:
		return "14066"
	case ERROR_SXS_XML_E_INVALIDENCODING:
		return "14067"
	case ERROR_SXS_XML_E_INVALIDSWITCH:
		return "14068"
	case ERROR_SXS_XML_E_BADXMLCASE:
		return "14069"
	case ERROR_SXS_XML_E_INVALID_STANDALONE:
		return "14070"
	case ERROR_SXS_XML_E_UNEXPECTED_STANDALONE:
		return "14071"
	case ERROR_SXS_XML_E_INVALID_VERSION:
		return "14072"
	case ERROR_SXS_XML_E_MISSINGEQUALS:
		return "14073"
	case ERROR_SXS_PROTECTION_RECOVERY_FAILED:
		return "14074"
	case ERROR_SXS_PROTECTION_PUBLIC_KEY_TOO_SHORT:
		return "14075"
	case ERROR_SXS_PROTECTION_CATALOG_NOT_VALID:
		return "14076"
	case ERROR_SXS_UNTRANSLATABLE_HRESULT:
		return "14077"
	case ERROR_SXS_PROTECTION_CATALOG_FILE_MISSING:
		return "14078"
	case ERROR_SXS_MISSING_ASSEMBLY_IDENTITY_ATTRIBUTE:
		return "14079"
	case ERROR_SXS_INVALID_ASSEMBLY_IDENTITY_ATTRIBUTE_NAME:
		return "14080"
	case ERROR_SXS_ASSEMBLY_MISSING:
		return "14081"
	case ERROR_SXS_CORRUPT_ACTIVATION_STACK:
		return "14082"
	case ERROR_SXS_CORRUPTION:
		return "14083"
	case ERROR_SXS_EARLY_DEACTIVATION:
		return "14084"
	case ERROR_SXS_INVALID_DEACTIVATION:
		return "14085"
	case ERROR_SXS_MULTIPLE_DEACTIVATION:
		return "14086"
	case ERROR_SXS_PROCESS_TERMINATION_REQUESTED:
		return "14087"
	case ERROR_SXS_RELEASE_ACTIVATION_CONTEXT:
		return "14088"
	case ERROR_SXS_SYSTEM_DEFAULT_ACTIVATION_CONTEXT_EMPTY:
		return "14089"
	case ERROR_SXS_INVALID_IDENTITY_ATTRIBUTE_VALUE:
		return "14090"
	case ERROR_SXS_INVALID_IDENTITY_ATTRIBUTE_NAME:
		return "14091"
	case ERROR_SXS_IDENTITY_DUPLICATE_ATTRIBUTE:
		return "14092"
	case ERROR_SXS_IDENTITY_PARSE_ERROR:
		return "14093"
	case ERROR_MALFORMED_SUBSTITUTION_STRING:
		return "14094"
	case ERROR_SXS_INCORRECT_PUBLIC_KEY_TOKEN:
		return "14095"
	case ERROR_UNMAPPED_SUBSTITUTION_STRING:
		return "14096"
	case ERROR_SXS_ASSEMBLY_NOT_LOCKED:
		return "14097"
	case ERROR_SXS_COMPONENT_STORE_CORRUPT:
		return "14098"
	case ERROR_ADVANCED_INSTALLER_FAILED:
		return "14099"
	case ERROR_XML_ENCODING_MISMATCH:
		return "14100"
	case ERROR_SXS_MANIFEST_IDENTITY_SAME_BUT_CONTENTS_DIFFERENT:
		return "14101"
	case ERROR_SXS_IDENTITIES_DIFFERENT:
		return "14102"
	case ERROR_SXS_ASSEMBLY_IS_NOT_A_DEPLOYMENT:
		return "14103"
	case ERROR_SXS_FILE_NOT_PART_OF_ASSEMBLY:
		return "14104"
	case ERROR_SXS_MANIFEST_TOO_BIG:
		return "14105"
	case ERROR_SXS_SETTING_NOT_REGISTERED:
		return "14106"
	case ERROR_SXS_TRANSACTION_CLOSURE_INCOMPLETE:
		return "14107"
	case ERROR_SMI_PRIMITIVE_INSTALLER_FAILED:
		return "14108"
	case ERROR_GENERIC_COMMAND_FAILED:
		return "14109"
	case ERROR_SXS_FILE_HASH_MISSING:
		return "14110"
	case ERROR_EVT_INVALID_CHANNEL_PATH:
		return "15000"
	case ERROR_EVT_INVALID_QUERY:
		return "15001"
	case ERROR_EVT_PUBLISHER_METADATA_NOT_FOUND:
		return "15002"
	case ERROR_EVT_EVENT_TEMPLATE_NOT_FOUND:
		return "15003"
	case ERROR_EVT_INVALID_PUBLISHER_NAME:
		return "15004"
	case ERROR_EVT_INVALID_EVENT_DATA:
		return "15005"
	case ERROR_EVT_CHANNEL_NOT_FOUND:
		return "15007"
	case ERROR_EVT_MALFORMED_XML_TEXT:
		return "15008"
	case ERROR_EVT_SUBSCRIPTION_TO_DIRECT_CHANNEL:
		return "15009"
	case ERROR_EVT_CONFIGURATION_ERROR:
		return "15010"
	case ERROR_EVT_QUERY_RESULT_STALE:
		return "15011"
	case ERROR_EVT_QUERY_RESULT_INVALID_POSITION:
		return "15012"
	case ERROR_EVT_NON_VALIDATING_MSXML:
		return "15013"
	case ERROR_EVT_FILTER_ALREADYSCOPED:
		return "15014"
	case ERROR_EVT_FILTER_NOTELTSET:
		return "15015"
	case ERROR_EVT_FILTER_INVARG:
		return "15016"
	case ERROR_EVT_FILTER_INVTEST:
		return "15017"
	case ERROR_EVT_FILTER_INVTYPE:
		return "15018"
	case ERROR_EVT_FILTER_PARSEERR:
		return "15019"
	case ERROR_EVT_FILTER_UNSUPPORTEDOP:
		return "15020"
	case ERROR_EVT_FILTER_UNEXPECTEDTOKEN:
		return "15021"
	case ERROR_EVT_INVALID_OPERATION_OVER_ENABLED_DIRECT_CHANNEL:
		return "15022"
	case ERROR_EVT_INVALID_CHANNEL_PROPERTY_VALUE:
		return "15023"
	case ERROR_EVT_INVALID_PUBLISHER_PROPERTY_VALUE:
		return "15024"
	case ERROR_EVT_CHANNEL_CANNOT_ACTIVATE:
		return "15025"
	case ERROR_EVT_FILTER_TOO_COMPLEX:
		return "15026"
	case ERROR_EVT_MESSAGE_NOT_FOUND:
		return "15027"
	case ERROR_EVT_MESSAGE_ID_NOT_FOUND:
		return "15028"
	case ERROR_EVT_UNRESOLVED_VALUE_INSERT:
		return "15029"
	case ERROR_EVT_UNRESOLVED_PARAMETER_INSERT:
		return "15030"
	case ERROR_EVT_MAX_INSERTS_REACHED:
		return "15031"
	case ERROR_EVT_EVENT_DEFINITION_NOT_FOUND:
		return "15032"
	case ERROR_EVT_MESSAGE_LOCALE_NOT_FOUND:
		return "15033"
	case ERROR_EVT_VERSION_TOO_OLD:
		return "15034"
	case ERROR_EVT_VERSION_TOO_NEW:
		return "15035"
	case ERROR_EVT_CANNOT_OPEN_CHANNEL_OF_QUERY:
		return "15036"
	case ERROR_EVT_PUBLISHER_DISABLED:
		return "15037"
	case ERROR_EVT_FILTER_OUT_OF_RANGE:
		return "15038"
	case ERROR_EC_SUBSCRIPTION_CANNOT_ACTIVATE:
		return "15080"
	case ERROR_EC_LOG_DISABLED:
		return "15081"
	case ERROR_EC_CIRCULAR_FORWARDING:
		return "15082"
	case ERROR_EC_CREDSTORE_FULL:
		return "15083"
	case ERROR_EC_CRED_NOT_FOUND:
		return "15084"
	case ERROR_EC_NO_ACTIVE_CHANNEL:
		return "15085"
	case ERROR_MUI_FILE_NOT_FOUND:
		return "15100"
	case ERROR_MUI_INVALID_FILE:
		return "15101"
	case ERROR_MUI_INVALID_RC_CONFIG:
		return "15102"
	case ERROR_MUI_INVALID_LOCALE_NAME:
		return "15103"
	case ERROR_MUI_INVALID_ULTIMATEFALLBACK_NAME:
		return "15104"
	case ERROR_MUI_FILE_NOT_LOADED:
		return "15105"
	case ERROR_RESOURCE_ENUM_USER_STOP:
		return "15106"
	case ERROR_MUI_INTLSETTINGS_UILANG_NOT_INSTALLED:
		return "15107"
	case ERROR_MUI_INTLSETTINGS_INVALID_LOCALE_NAME:
		return "15108"
	case ERROR_MCA_INVALID_CAPABILITIES_STRING:
		return "15200"
	case ERROR_MCA_INVALID_VCP_VERSION:
		return "15201"
	case ERROR_MCA_MONITOR_VIOLATES_MCCS_SPECIFICATION:
		return "15202"
	case ERROR_MCA_MCCS_VERSION_MISMATCH:
		return "15203"
	case ERROR_MCA_UNSUPPORTED_MCCS_VERSION:
		return "15204"
	case ERROR_MCA_INTERNAL_ERROR:
		return "15205"
	case ERROR_MCA_INVALID_TECHNOLOGY_TYPE_RETURNED:
		return "15206"
	case ERROR_MCA_UNSUPPORTED_COLOR_TEMPERATURE:
		return "15207"
	case ERROR_AMBIGUOUS_SYSTEM_DEVICE:
		return "15250"
	case ERROR_SYSTEM_DEVICE_NOT_FOUND:
		return "15299"
	case ERROR_HASH_NOT_SUPPORTED:
		return "15300"
	case ERROR_HASH_NOT_PRESENT:
		return "15301"
	case SEVERITY_SUCCESS:
		return "0"
	case SEVERITY_ERROR:
		return "1"
	case FACILITY_NT_BIT:
		return "0x10000000"
	case E_UNEXPECTED:
		return "0x8000FFFF"
	case E_NOTIMPL:
		return "0x80004001"
	case E_OUTOFMEMORY:
		return "0x8007000E"
	case E_INVALIDARG:
		return "0x80070057"
	case E_NOINTERFACE:
		return "0x80004002"
	case E_POINTER:
		return "0x80004003"
	case E_HANDLE:
		return "0x80070006"
	case E_ABORT:
		return "0x80004004"
	case E_FAIL:
		return "0x80004005"
	case E_ACCESSDENIED:
		return "0x80070005"
	case E_PENDING:
		return "0x8000000A"
	case CO_E_INIT_TLS:
		return "0x80004006"
	case CO_E_INIT_SHARED_ALLOCATOR:
		return "0x80004007"
	case CO_E_INIT_MEMORY_ALLOCATOR:
		return "0x80004008"
	case CO_E_INIT_CLASS_CACHE:
		return "0x80004009"
	case CO_E_INIT_RPC_CHANNEL:
		return "0x8000400A"
	case CO_E_INIT_TLS_SET_CHANNEL_CONTROL:
		return "0x8000400B"
	case CO_E_INIT_TLS_CHANNEL_CONTROL:
		return "0x8000400C"
	case CO_E_INIT_UNACCEPTED_USER_ALLOCATOR:
		return "0x8000400D"
	case CO_E_INIT_SCM_MUTEX_EXISTS:
		return "0x8000400E"
	case CO_E_INIT_SCM_FILE_MAPPING_EXISTS:
		return "0x8000400F"
	case CO_E_INIT_SCM_MAP_VIEW_OF_FILE:
		return "0x80004010"
	case CO_E_INIT_SCM_EXEC_FAILURE:
		return "0x80004011"
	case CO_E_INIT_ONLY_SINGLE_THREADED:
		return "0x80004012"
	case CO_E_CANT_REMOTE:
		return "0x80004013"
	case CO_E_BAD_SERVER_NAME:
		return "0x80004014"
	case CO_E_WRONG_SERVER_IDENTITY:
		return "0x80004015"
	case CO_E_OLE1DDE_DISABLED:
		return "0x80004016"
	case CO_E_RUNAS_SYNTAX:
		return "0x80004017"
	case CO_E_CREATEPROCESS_FAILURE:
		return "0x80004018"
	case CO_E_RUNAS_CREATEPROCESS_FAILURE:
		return "0x80004019"
	case CO_E_RUNAS_LOGON_FAILURE:
		return "0x8000401A"
	case CO_E_LAUNCH_PERMSSION_DENIED:
		return "0x8000401B"
	case CO_E_START_SERVICE_FAILURE:
		return "0x8000401C"
	case CO_E_REMOTE_COMMUNICATION_FAILURE:
		return "0x8000401D"
	case CO_E_SERVER_START_TIMEOUT:
		return "0x8000401E"
	case CO_E_CLSREG_INCONSISTENT:
		return "0x8000401F"
	case CO_E_IIDREG_INCONSISTENT:
		return "0x80004020"
	case CO_E_NOT_SUPPORTED:
		return "0x80004021"
	case CO_E_RELOAD_DLL:
		return "0x80004022"
	case CO_E_MSI_ERROR:
		return "0x80004023"
	case CO_E_ATTEMPT_TO_CREATE_OUTSIDE_CLIENT_CONTEXT:
		return "0x80004024"
	case CO_E_SERVER_PAUSED:
		return "0x80004025"
	case CO_E_SERVER_NOT_PAUSED:
		return "0x80004026"
	case CO_E_CLASS_DISABLED:
		return "0x80004027"
	case CO_E_CLRNOTAVAILABLE:
		return "0x80004028"
	case CO_E_ASYNC_WORK_REJECTED:
		return "0x80004029"
	case CO_E_SERVER_INIT_TIMEOUT:
		return "0x8000402A"
	case CO_E_NO_SECCTX_IN_ACTIVATE:
		return "0x8000402B"
	case CO_E_TRACKER_CONFIG:
		return "0x80004030"
	case CO_E_THREADPOOL_CONFIG:
		return "0x80004031"
	case CO_E_SXS_CONFIG:
		return "0x80004032"
	case CO_E_MALFORMED_SPN:
		return "0x80004033"
	case S_OK:
		return "0"
	case S_FALSE:
		return "1"
	case OLE_E_FIRST:
		return "0x80040000"
	case OLE_E_LAST:
		return "0x800400FF"
	case OLE_S_FIRST:
		return "0x00040000"
	case OLE_S_LAST:
		return "0x000400FF"
	case OLE_E_OLEVERB:
		return "0x80040000"
	case OLE_E_ADVF:
		return "0x80040001"
	case OLE_E_ENUM_NOMORE:
		return "0x80040002"
	case OLE_E_ADVISENOTSUPPORTED:
		return "0x80040003"
	case OLE_E_NOCONNECTION:
		return "0x80040004"
	case OLE_E_NOTRUNNING:
		return "0x80040005"
	case OLE_E_NOCACHE:
		return "0x80040006"
	case OLE_E_BLANK:
		return "0x80040007"
	case OLE_E_CLASSDIFF:
		return "0x80040008"
	case OLE_E_CANT_GETMONIKER:
		return "0x80040009"
	case OLE_E_CANT_BINDTOSOURCE:
		return "0x8004000A"
	case OLE_E_STATIC:
		return "0x8004000B"
	case OLE_E_PROMPTSAVECANCELLED:
		return "0x8004000C"
	case OLE_E_INVALIDRECT:
		return "0x8004000D"
	case OLE_E_WRONGCOMPOBJ:
		return "0x8004000E"
	case OLE_E_INVALIDHWND:
		return "0x8004000F"
	case OLE_E_NOT_INPLACEACTIVE:
		return "0x80040010"
	case OLE_E_CANTCONVERT:
		return "0x80040011"
	case OLE_E_NOSTORAGE:
		return "0x80040012"
	case DV_E_FORMATETC:
		return "0x80040064"
	case DV_E_DVTARGETDEVICE:
		return "0x80040065"
	case DV_E_STGMEDIUM:
		return "0x80040066"
	case DV_E_STATDATA:
		return "0x80040067"
	case DV_E_LINDEX:
		return "0x80040068"
	case DV_E_TYMED:
		return "0x80040069"
	case DV_E_CLIPFORMAT:
		return "0x8004006A"
	case DV_E_DVASPECT:
		return "0x8004006B"
	case DV_E_DVTARGETDEVICE_SIZE:
		return "0x8004006C"
	case DV_E_NOIVIEWOBJECT:
		return "0x8004006D"
	case DRAGDROP_E_FIRST:
		return "0x80040100"
	case DRAGDROP_E_LAST:
		return "0x8004010F"
	case DRAGDROP_S_FIRST:
		return "0x00040100"
	case DRAGDROP_S_LAST:
		return "0x0004010F"
	case DRAGDROP_E_NOTREGISTERED:
		return "0x80040100"
	case DRAGDROP_E_ALREADYREGISTERED:
		return "0x80040101"
	case DRAGDROP_E_INVALIDHWND:
		return "0x80040102"
	case CLASSFACTORY_E_FIRST:
		return "0x80040110"
	case CLASSFACTORY_E_LAST:
		return "0x8004011F"
	case CLASSFACTORY_S_FIRST:
		return "0x00040110"
	case CLASSFACTORY_S_LAST:
		return "0x0004011F"
	case CLASS_E_NOAGGREGATION:
		return "0x80040110"
	case CLASS_E_CLASSNOTAVAILABLE:
		return "0x80040111"
	case CLASS_E_NOTLICENSED:
		return "0x80040112"
	case MARSHAL_E_FIRST:
		return "0x80040120"
	case MARSHAL_E_LAST:
		return "0x8004012F"
	case MARSHAL_S_FIRST:
		return "0x00040120"
	case MARSHAL_S_LAST:
		return "0x0004012F"
	case DATA_E_FIRST:
		return "0x80040130"
	case DATA_E_LAST:
		return "0x8004013F"
	case DATA_S_FIRST:
		return "0x00040130"
	case DATA_S_LAST:
		return "0x0004013F"
	case VIEW_E_FIRST:
		return "0x80040140"
	case VIEW_E_LAST:
		return "0x8004014F"
	case VIEW_S_FIRST:
		return "0x00040140"
	case VIEW_S_LAST:
		return "0x0004014F"
	case VIEW_E_DRAW:
		return "0x80040140"
	case REGDB_E_FIRST:
		return "0x80040150"
	case REGDB_E_LAST:
		return "0x8004015F"
	case REGDB_S_FIRST:
		return "0x00040150"
	case REGDB_S_LAST:
		return "0x0004015F"
	case REGDB_E_READREGDB:
		return "0x80040150"
	case REGDB_E_WRITEREGDB:
		return "0x80040151"
	case REGDB_E_KEYMISSING:
		return "0x80040152"
	case REGDB_E_INVALIDVALUE:
		return "0x80040153"
	case REGDB_E_CLASSNOTREG:
		return "0x80040154"
	case REGDB_E_IIDNOTREG:
		return "0x80040155"
	case REGDB_E_BADTHREADINGMODEL:
		return "0x80040156"
	case CAT_E_FIRST:
		return "0x80040160"
	case CAT_E_LAST:
		return "0x80040161"
	case CAT_E_CATIDNOEXIST:
		return "0x80040160"
	case CAT_E_NODESCRIPTION:
		return "0x80040161"
	case CS_E_FIRST:
		return "0x80040164"
	case CS_E_LAST:
		return "0x8004016F"
	case CS_E_PACKAGE_NOTFOUND:
		return "0x80040164"
	case CS_E_NOT_DELETABLE:
		return "0x80040165"
	case CS_E_CLASS_NOTFOUND:
		return "0x80040166"
	case CS_E_INVALID_VERSION:
		return "0x80040167"
	case CS_E_NO_CLASSSTORE:
		return "0x80040168"
	case CS_E_OBJECT_NOTFOUND:
		return "0x80040169"
	case CS_E_OBJECT_ALREADY_EXISTS:
		return "0x8004016A"
	case CS_E_INVALID_PATH:
		return "0x8004016B"
	case CS_E_NETWORK_ERROR:
		return "0x8004016C"
	case CS_E_ADMIN_LIMIT_EXCEEDED:
		return "0x8004016D"
	case CS_E_SCHEMA_MISMATCH:
		return "0x8004016E"
	case CS_E_INTERNAL_ERROR:
		return "0x8004016F"
	case CACHE_E_FIRST:
		return "0x80040170"
	case CACHE_E_LAST:
		return "0x8004017F"
	case CACHE_S_FIRST:
		return "0x00040170"
	case CACHE_S_LAST:
		return "0x0004017F"
	case CACHE_E_NOCACHE_UPDATED:
		return "0x80040170"
	case OLEOBJ_E_FIRST:
		return "0x80040180"
	case OLEOBJ_E_LAST:
		return "0x8004018F"
	case OLEOBJ_S_FIRST:
		return "0x00040180"
	case OLEOBJ_S_LAST:
		return "0x0004018F"
	case OLEOBJ_E_NOVERBS:
		return "0x80040180"
	case OLEOBJ_E_INVALIDVERB:
		return "0x80040181"
	case CLIENTSITE_E_FIRST:
		return "0x80040190"
	case CLIENTSITE_E_LAST:
		return "0x8004019F"
	case CLIENTSITE_S_FIRST:
		return "0x00040190"
	case CLIENTSITE_S_LAST:
		return "0x0004019F"
	case INPLACE_E_NOTUNDOABLE:
		return "0x800401A0"
	case INPLACE_E_NOTOOLSPACE:
		return "0x800401A1"
	case INPLACE_E_FIRST:
		return "0x800401A0"
	case INPLACE_E_LAST:
		return "0x800401AF"
	case INPLACE_S_FIRST:
		return "0x000401A0"
	case INPLACE_S_LAST:
		return "0x000401AF"
	case ENUM_E_FIRST:
		return "0x800401B0"
	case ENUM_E_LAST:
		return "0x800401BF"
	case ENUM_S_FIRST:
		return "0x000401B0"
	case ENUM_S_LAST:
		return "0x000401BF"
	case CONVERT10_E_FIRST:
		return "0x800401C0"
	case CONVERT10_E_LAST:
		return "0x800401CF"
	case CONVERT10_S_FIRST:
		return "0x000401C0"
	case CONVERT10_S_LAST:
		return "0x000401CF"
	case CONVERT10_E_OLESTREAM_GET:
		return "0x800401C0"
	case CONVERT10_E_OLESTREAM_PUT:
		return "0x800401C1"
	case CONVERT10_E_OLESTREAM_FMT:
		return "0x800401C2"
	case CONVERT10_E_OLESTREAM_BITMAP_TO_DIB:
		return "0x800401C3"
	case CONVERT10_E_STG_FMT:
		return "0x800401C4"
	case CONVERT10_E_STG_NO_STD_STREAM:
		return "0x800401C5"
	case CONVERT10_E_STG_DIB_TO_BITMAP:
		return "0x800401C6"
	case CLIPBRD_E_FIRST:
		return "0x800401D0"
	case CLIPBRD_E_LAST:
		return "0x800401DF"
	case CLIPBRD_S_FIRST:
		return "0x000401D0"
	case CLIPBRD_S_LAST:
		return "0x000401DF"
	case CLIPBRD_E_CANT_OPEN:
		return "0x800401D0"
	case CLIPBRD_E_CANT_EMPTY:
		return "0x800401D1"
	case CLIPBRD_E_CANT_SET:
		return "0x800401D2"
	case CLIPBRD_E_BAD_DATA:
		return "0x800401D3"
	case CLIPBRD_E_CANT_CLOSE:
		return "0x800401D4"
	case MK_E_FIRST:
		return "0x800401E0"
	case MK_E_LAST:
		return "0x800401EF"
	case MK_S_FIRST:
		return "0x000401E0"
	case MK_S_LAST:
		return "0x000401EF"
	case MK_E_CONNECTMANUALLY:
		return "0x800401E0"
	case MK_E_EXCEEDEDDEADLINE:
		return "0x800401E1"
	case MK_E_NEEDGENERIC:
		return "0x800401E2"
	case MK_E_UNAVAILABLE:
		return "0x800401E3"
	case MK_E_SYNTAX:
		return "0x800401E4"
	case MK_E_NOOBJECT:
		return "0x800401E5"
	case MK_E_INVALIDEXTENSION:
		return "0x800401E6"
	case MK_E_INTERMEDIATEINTERFACENOTSUPPORTED:
		return "0x800401E7"
	case MK_E_NOTBINDABLE:
		return "0x800401E8"
	case MK_E_NOTBOUND:
		return "0x800401E9"
	case MK_E_CANTOPENFILE:
		return "0x800401EA"
	case MK_E_MUSTBOTHERUSER:
		return "0x800401EB"
	case MK_E_NOINVERSE:
		return "0x800401EC"
	case MK_E_NOSTORAGE:
		return "0x800401ED"
	case MK_E_NOPREFIX:
		return "0x800401EE"
	case MK_E_ENUMERATION_FAILED:
		return "0x800401EF"
	case CO_E_FIRST:
		return "0x800401F0"
	case CO_E_LAST:
		return "0x800401FF"
	case CO_S_FIRST:
		return "0x000401F0"
	case CO_S_LAST:
		return "0x000401FF"
	case CO_E_NOTINITIALIZED:
		return "0x800401F0"
	case CO_E_ALREADYINITIALIZED:
		return "0x800401F1"
	case CO_E_CANTDETERMINECLASS:
		return "0x800401F2"
	case CO_E_CLASSSTRING:
		return "0x800401F3"
	case CO_E_IIDSTRING:
		return "0x800401F4"
	case CO_E_APPNOTFOUND:
		return "0x800401F5"
	case CO_E_APPSINGLEUSE:
		return "0x800401F6"
	case CO_E_ERRORINAPP:
		return "0x800401F7"
	case CO_E_DLLNOTFOUND:
		return "0x800401F8"
	case CO_E_ERRORINDLL:
		return "0x800401F9"
	case CO_E_WRONGOSFORAPP:
		return "0x800401FA"
	case CO_E_OBJNOTREG:
		return "0x800401FB"
	case CO_E_OBJISREG:
		return "0x800401FC"
	case CO_E_OBJNOTCONNECTED:
		return "0x800401FD"
	case CO_E_APPDIDNTREG:
		return "0x800401FE"
	case CO_E_RELEASED:
		return "0x800401FF"
	case EVENT_E_FIRST:
		return "0x80040200"
	case EVENT_E_LAST:
		return "0x8004021F"
	case EVENT_S_FIRST:
		return "0x00040200"
	case EVENT_S_LAST:
		return "0x0004021F"
	case EVENT_S_SOME_SUBSCRIBERS_FAILED:
		return "0x00040200"
	case EVENT_E_ALL_SUBSCRIBERS_FAILED:
		return "0x80040201"
	case EVENT_S_NOSUBSCRIBERS:
		return "0x00040202"
	case EVENT_E_QUERYSYNTAX:
		return "0x80040203"
	case EVENT_E_QUERYFIELD:
		return "0x80040204"
	case EVENT_E_INTERNALEXCEPTION:
		return "0x80040205"
	case EVENT_E_INTERNALERROR:
		return "0x80040206"
	case EVENT_E_INVALID_PER_USER_SID:
		return "0x80040207"
	case EVENT_E_USER_EXCEPTION:
		return "0x80040208"
	case EVENT_E_TOO_MANY_METHODS:
		return "0x80040209"
	case EVENT_E_MISSING_EVENTCLASS:
		return "0x8004020A"
	case EVENT_E_NOT_ALL_REMOVED:
		return "0x8004020B"
	case EVENT_E_COMPLUS_NOT_INSTALLED:
		return "0x8004020C"
	case EVENT_E_CANT_MODIFY_OR_DELETE_UNCONFIGURED_OBJECT:
		return "0x8004020D"
	case EVENT_E_CANT_MODIFY_OR_DELETE_CONFIGURED_OBJECT:
		return "0x8004020E"
	case EVENT_E_INVALID_EVENT_CLASS_PARTITION:
		return "0x8004020F"
	case EVENT_E_PER_USER_SID_NOT_LOGGED_ON:
		return "0x80040210"
	case XACT_E_FIRST:
		return "0x8004D000"
	case XACT_E_LAST:
		return "0x8004D02B"
	case XACT_S_FIRST:
		return "0x0004D000"
	case XACT_S_LAST:
		return "0x0004D010"
	case XACT_E_ALREADYOTHERSINGLEPHASE:
		return "0x8004D000"
	case XACT_E_CANTRETAIN:
		return "0x8004D001"
	case XACT_E_COMMITFAILED:
		return "0x8004D002"
	case XACT_E_COMMITPREVENTED:
		return "0x8004D003"
	case XACT_E_HEURISTICABORT:
		return "0x8004D004"
	case XACT_E_HEURISTICCOMMIT:
		return "0x8004D005"
	case XACT_E_HEURISTICDAMAGE:
		return "0x8004D006"
	case XACT_E_HEURISTICDANGER:
		return "0x8004D007"
	case XACT_E_ISOLATIONLEVEL:
		return "0x8004D008"
	case XACT_E_NOASYNC:
		return "0x8004D009"
	case XACT_E_NOENLIST:
		return "0x8004D00A"
	case XACT_E_NOISORETAIN:
		return "0x8004D00B"
	case XACT_E_NORESOURCE:
		return "0x8004D00C"
	case XACT_E_NOTCURRENT:
		return "0x8004D00D"
	case XACT_E_NOTRANSACTION:
		return "0x8004D00E"
	case XACT_E_NOTSUPPORTED:
		return "0x8004D00F"
	case XACT_E_UNKNOWNRMGRID:
		return "0x8004D010"
	case XACT_E_WRONGSTATE:
		return "0x8004D011"
	case XACT_E_WRONGUOW:
		return "0x8004D012"
	case XACT_E_XTIONEXISTS:
		return "0x8004D013"
	case XACT_E_NOIMPORTOBJECT:
		return "0x8004D014"
	case XACT_E_INVALIDCOOKIE:
		return "0x8004D015"
	case XACT_E_INDOUBT:
		return "0x8004D016"
	case XACT_E_NOTIMEOUT:
		return "0x8004D017"
	case XACT_E_ALREADYINPROGRESS:
		return "0x8004D018"
	case XACT_E_ABORTED:
		return "0x8004D019"
	case XACT_E_LOGFULL:
		return "0x8004D01A"
	case XACT_E_TMNOTAVAILABLE:
		return "0x8004D01B"
	case XACT_E_CONNECTION_DOWN:
		return "0x8004D01C"
	case XACT_E_CONNECTION_DENIED:
		return "0x8004D01D"
	case XACT_E_REENLISTTIMEOUT:
		return "0x8004D01E"
	case XACT_E_TIP_CONNECT_FAILED:
		return "0x8004D01F"
	case XACT_E_TIP_PROTOCOL_ERROR:
		return "0x8004D020"
	case XACT_E_TIP_PULL_FAILED:
		return "0x8004D021"
	case XACT_E_DEST_TMNOTAVAILABLE:
		return "0x8004D022"
	case XACT_E_TIP_DISABLED:
		return "0x8004D023"
	case XACT_E_NETWORK_TX_DISABLED:
		return "0x8004D024"
	case XACT_E_PARTNER_NETWORK_TX_DISABLED:
		return "0x8004D025"
	case XACT_E_XA_TX_DISABLED:
		return "0x8004D026"
	case XACT_E_UNABLE_TO_READ_DTC_CONFIG:
		return "0x8004D027"
	case XACT_E_UNABLE_TO_LOAD_DTC_PROXY:
		return "0x8004D028"
	case XACT_E_ABORTING:
		return "0x8004D029"
	case XACT_E_PUSH_COMM_FAILURE:
		return "0x8004D02A"
	case XACT_E_PULL_COMM_FAILURE:
		return "0x8004D02B"
	case XACT_E_LU_TX_DISABLED:
		return "0x8004D02C"
	case XACT_E_CLERKNOTFOUND:
		return "0x8004D080"
	case XACT_E_CLERKEXISTS:
		return "0x8004D081"
	case XACT_E_RECOVERYINPROGRESS:
		return "0x8004D082"
	case XACT_E_TRANSACTIONCLOSED:
		return "0x8004D083"
	case XACT_E_INVALIDLSN:
		return "0x8004D084"
	case XACT_E_REPLAYREQUEST:
		return "0x8004D085"
	case XACT_S_ASYNC:
		return "0x0004D000"
	case XACT_S_DEFECT:
		return "0x0004D001"
	case XACT_S_READONLY:
		return "0x0004D002"
	case XACT_S_SOMENORETAIN:
		return "0x0004D003"
	case XACT_S_OKINFORM:
		return "0x0004D004"
	case XACT_S_MADECHANGESCONTENT:
		return "0x0004D005"
	case XACT_S_MADECHANGESINFORM:
		return "0x0004D006"
	case XACT_S_ALLNORETAIN:
		return "0x0004D007"
	case XACT_S_ABORTING:
		return "0x0004D008"
	case XACT_S_SINGLEPHASE:
		return "0x0004D009"
	case XACT_S_LOCALLY_OK:
		return "0x0004D00A"
	case XACT_S_LASTRESOURCEMANAGER:
		return "0x0004D010"
	case CONTEXT_E_FIRST:
		return "0x8004E000"
	case CONTEXT_E_LAST:
		return "0x8004E02F"
	case CONTEXT_S_FIRST:
		return "0x0004E000"
	case CONTEXT_S_LAST:
		return "0x0004E02F"
	case CONTEXT_E_ABORTED:
		return "0x8004E002"
	case CONTEXT_E_ABORTING:
		return "0x8004E003"
	case CONTEXT_E_NOCONTEXT:
		return "0x8004E004"
	case CONTEXT_E_WOULD_DEADLOCK:
		return "0x8004E005"
	case CONTEXT_E_SYNCH_TIMEOUT:
		return "0x8004E006"
	case CONTEXT_E_OLDREF:
		return "0x8004E007"
	case CONTEXT_E_ROLENOTFOUND:
		return "0x8004E00C"
	case CONTEXT_E_TMNOTAVAILABLE:
		return "0x8004E00F"
	case CO_E_ACTIVATIONFAILED:
		return "0x8004E021"
	case CO_E_ACTIVATIONFAILED_EVENTLOGGED:
		return "0x8004E022"
	case CO_E_ACTIVATIONFAILED_CATALOGERROR:
		return "0x8004E023"
	case CO_E_ACTIVATIONFAILED_TIMEOUT:
		return "0x8004E024"
	case CO_E_INITIALIZATIONFAILED:
		return "0x8004E025"
	case CONTEXT_E_NOJIT:
		return "0x8004E026"
	case CONTEXT_E_NOTRANSACTION:
		return "0x8004E027"
	case CO_E_THREADINGMODEL_CHANGED:
		return "0x8004E028"
	case CO_E_NOIISINTRINSICS:
		return "0x8004E029"
	case CO_E_NOCOOKIES:
		return "0x8004E02A"
	case CO_E_DBERROR:
		return "0x8004E02B"
	case CO_E_NOTPOOLED:
		return "0x8004E02C"
	case CO_E_NOTCONSTRUCTED:
		return "0x8004E02D"
	case CO_E_NOSYNCHRONIZATION:
		return "0x8004E02E"
	case CO_E_ISOLEVELMISMATCH:
		return "0x8004E02F"
	case CO_E_CALL_OUT_OF_TX_SCOPE_NOT_ALLOWED:
		return "0x8004E030"
	case CO_E_EXIT_TRANSACTION_SCOPE_NOT_CALLED:
		return "0x8004E031"
	case OLE_S_USEREG:
		return "0x00040000"
	case OLE_S_STATIC:
		return "0x00040001"
	case OLE_S_MAC_CLIPFORMAT:
		return "0x00040002"
	case DRAGDROP_S_DROP:
		return "0x00040100"
	case DRAGDROP_S_CANCEL:
		return "0x00040101"
	case DRAGDROP_S_USEDEFAULTCURSORS:
		return "0x00040102"
	case DATA_S_SAMEFORMATETC:
		return "0x00040130"
	case VIEW_S_ALREADY_FROZEN:
		return "0x00040140"
	case CACHE_S_FORMATETC_NOTSUPPORTED:
		return "0x00040170"
	case CACHE_S_SAMECACHE:
		return "0x00040171"
	case CACHE_S_SOMECACHES_NOTUPDATED:
		return "0x00040172"
	case OLEOBJ_S_INVALIDVERB:
		return "0x00040180"
	case OLEOBJ_S_CANNOT_DOVERB_NOW:
		return "0x00040181"
	case OLEOBJ_S_INVALIDHWND:
		return "0x00040182"
	case INPLACE_S_TRUNCATED:
		return "0x000401A0"
	case CONVERT10_S_NO_PRESENTATION:
		return "0x000401C0"
	case MK_S_REDUCED_TO_SELF:
		return "0x000401E2"
	case MK_S_ME:
		return "0x000401E4"
	case MK_S_HIM:
		return "0x000401E5"
	case MK_S_US:
		return "0x000401E6"
	case MK_S_MONIKERALREADYREGISTERED:
		return "0x000401E7"
	case SCHED_S_TASK_READY:
		return "0x00041300"
	case SCHED_S_TASK_RUNNING:
		return "0x00041301"
	case SCHED_S_TASK_DISABLED:
		return "0x00041302"
	case SCHED_S_TASK_HAS_NOT_RUN:
		return "0x00041303"
	case SCHED_S_TASK_NO_MORE_RUNS:
		return "0x00041304"
	case SCHED_S_TASK_NOT_SCHEDULED:
		return "0x00041305"
	case SCHED_S_TASK_TERMINATED:
		return "0x00041306"
	case SCHED_S_TASK_NO_VALID_TRIGGERS:
		return "0x00041307"
	case SCHED_S_EVENT_TRIGGER:
		return "0x00041308"
	case SCHED_E_TRIGGER_NOT_FOUND:
		return "0x80041309"
	case SCHED_E_TASK_NOT_READY:
		return "0x8004130A"
	case SCHED_E_TASK_NOT_RUNNING:
		return "0x8004130B"
	case SCHED_E_SERVICE_NOT_INSTALLED:
		return "0x8004130C"
	case SCHED_E_CANNOT_OPEN_TASK:
		return "0x8004130D"
	case SCHED_E_INVALID_TASK:
		return "0x8004130E"
	case SCHED_E_ACCOUNT_INFORMATION_NOT_SET:
		return "0x8004130F"
	case SCHED_E_ACCOUNT_NAME_NOT_FOUND:
		return "0x80041310"
	case SCHED_E_ACCOUNT_DBASE_CORRUPT:
		return "0x80041311"
	case SCHED_E_NO_SECURITY_SERVICES:
		return "0x80041312"
	case SCHED_E_UNKNOWN_OBJECT_VERSION:
		return "0x80041313"
	case SCHED_E_UNSUPPORTED_ACCOUNT_OPTION:
		return "0x80041314"
	case SCHED_E_SERVICE_NOT_RUNNING:
		return "0x80041315"
	case SCHED_E_UNEXPECTEDNODE:
		return "0x80041316"
	case SCHED_E_NAMESPACE:
		return "0x80041317"
	case SCHED_E_INVALIDVALUE:
		return "0x80041318"
	case SCHED_E_MISSINGNODE:
		return "0x80041319"
	case SCHED_E_MALFORMEDXML:
		return "0x8004131A"
	case SCHED_S_SOME_TRIGGERS_FAILED:
		return "0x0004131B"
	case SCHED_S_BATCH_LOGON_PROBLEM:
		return "0x0004131C"
	case SCHED_E_TOO_MANY_NODES:
		return "0x8004131D"
	case SCHED_E_PAST_END_BOUNDARY:
		return "0x8004131E"
	case SCHED_E_ALREADY_RUNNING:
		return "0x8004131F"
	case SCHED_E_USER_NOT_LOGGED_ON:
		return "0x80041320"
	case SCHED_E_INVALID_TASK_HASH:
		return "0x80041321"
	case SCHED_E_SERVICE_NOT_AVAILABLE:
		return "0x80041322"
	case SCHED_E_SERVICE_TOO_BUSY:
		return "0x80041323"
	case SCHED_E_TASK_ATTEMPTED:
		return "0x80041324"
	case SCHED_S_TASK_QUEUED:
		return "0x00041325"
	case SCHED_E_TASK_DISABLED:
		return "0x80041326"
	case SCHED_E_TASK_NOT_V1_COMPAT:
		return "0x80041327"
	case SCHED_E_START_ON_DEMAND:
		return "0x80041328"
	case CO_E_CLASS_CREATE_FAILED:
		return "0x80080001"
	case CO_E_SCM_ERROR:
		return "0x80080002"
	case CO_E_SCM_RPC_FAILURE:
		return "0x80080003"
	case CO_E_BAD_PATH:
		return "0x80080004"
	case CO_E_SERVER_EXEC_FAILURE:
		return "0x80080005"
	case CO_E_OBJSRV_RPC_FAILURE:
		return "0x80080006"
	case MK_E_NO_NORMALIZED:
		return "0x80080007"
	case CO_E_SERVER_STOPPING:
		return "0x80080008"
	case MEM_E_INVALID_ROOT:
		return "0x80080009"
	case MEM_E_INVALID_LINK:
		return "0x80080010"
	case MEM_E_INVALID_SIZE:
		return "0x80080011"
	case CO_S_NOTALLINTERFACES:
		return "0x00080012"
	case CO_S_MACHINENAMENOTFOUND:
		return "0x00080013"
	case CO_E_MISSING_DISPLAYNAME:
		return "0x80080015"
	case CO_E_RUNAS_VALUE_MUST_BE_AAA:
		return "0x80080016"
	case CO_E_ELEVATION_DISABLED:
		return "0x80080017"
	case DISP_E_UNKNOWNINTERFACE:
		return "0x80020001"
	case DISP_E_MEMBERNOTFOUND:
		return "0x80020003"
	case DISP_E_PARAMNOTFOUND:
		return "0x80020004"
	case DISP_E_TYPEMISMATCH:
		return "0x80020005"
	case DISP_E_UNKNOWNNAME:
		return "0x80020006"
	case DISP_E_NONAMEDARGS:
		return "0x80020007"
	case DISP_E_BADVARTYPE:
		return "0x80020008"
	case DISP_E_EXCEPTION:
		return "0x80020009"
	case DISP_E_OVERFLOW:
		return "0x8002000A"
	case DISP_E_BADINDEX:
		return "0x8002000B"
	case DISP_E_UNKNOWNLCID:
		return "0x8002000C"
	case DISP_E_ARRAYISLOCKED:
		return "0x8002000D"
	case DISP_E_BADPARAMCOUNT:
		return "0x8002000E"
	case DISP_E_PARAMNOTOPTIONAL:
		return "0x8002000F"
	case DISP_E_BADCALLEE:
		return "0x80020010"
	case DISP_E_NOTACOLLECTION:
		return "0x80020011"
	case DISP_E_DIVBYZERO:
		return "0x80020012"
	case DISP_E_BUFFERTOOSMALL:
		return "0x80020013"
	case TYPE_E_BUFFERTOOSMALL:
		return "0x80028016"
	case TYPE_E_FIELDNOTFOUND:
		return "0x80028017"
	case TYPE_E_INVDATAREAD:
		return "0x80028018"
	case TYPE_E_UNSUPFORMAT:
		return "0x80028019"
	case TYPE_E_REGISTRYACCESS:
		return "0x8002801C"
	case TYPE_E_LIBNOTREGISTERED:
		return "0x8002801D"
	case TYPE_E_UNDEFINEDTYPE:
		return "0x80028027"
	case TYPE_E_QUALIFIEDNAMEDISALLOWED:
		return "0x80028028"
	case TYPE_E_INVALIDSTATE:
		return "0x80028029"
	case TYPE_E_WRONGTYPEKIND:
		return "0x8002802A"
	case TYPE_E_ELEMENTNOTFOUND:
		return "0x8002802B"
	case TYPE_E_AMBIGUOUSNAME:
		return "0x8002802C"
	case TYPE_E_NAMECONFLICT:
		return "0x8002802D"
	case TYPE_E_UNKNOWNLCID:
		return "0x8002802E"
	case TYPE_E_DLLFUNCTIONNOTFOUND:
		return "0x8002802F"
	case TYPE_E_BADMODULEKIND:
		return "0x800288BD"
	case TYPE_E_SIZETOOBIG:
		return "0x800288C5"
	case TYPE_E_DUPLICATEID:
		return "0x800288C6"
	case TYPE_E_INVALIDID:
		return "0x800288CF"
	case TYPE_E_TYPEMISMATCH:
		return "0x80028CA0"
	case TYPE_E_OUTOFBOUNDS:
		return "0x80028CA1"
	case TYPE_E_IOERROR:
		return "0x80028CA2"
	case TYPE_E_CANTCREATETMPFILE:
		return "0x80028CA3"
	case TYPE_E_CANTLOADLIBRARY:
		return "0x80029C4A"
	case TYPE_E_INCONSISTENTPROPFUNCS:
		return "0x80029C83"
	case TYPE_E_CIRCULARTYPE:
		return "0x80029C84"
	case STG_E_INVALIDFUNCTION:
		return "0x80030001"
	case STG_E_FILENOTFOUND:
		return "0x80030002"
	case STG_E_PATHNOTFOUND:
		return "0x80030003"
	case STG_E_TOOMANYOPENFILES:
		return "0x80030004"
	case STG_E_ACCESSDENIED:
		return "0x80030005"
	case STG_E_INVALIDHANDLE:
		return "0x80030006"
	case STG_E_INSUFFICIENTMEMORY:
		return "0x80030008"
	case STG_E_INVALIDPOINTER:
		return "0x80030009"
	case STG_E_NOMOREFILES:
		return "0x80030012"
	case STG_E_DISKISWRITEPROTECTED:
		return "0x80030013"
	case STG_E_SEEKERROR:
		return "0x80030019"
	case STG_E_WRITEFAULT:
		return "0x8003001D"
	case STG_E_READFAULT:
		return "0x8003001E"
	case STG_E_SHAREVIOLATION:
		return "0x80030020"
	case STG_E_LOCKVIOLATION:
		return "0x80030021"
	case STG_E_FILEALREADYEXISTS:
		return "0x80030050"
	case STG_E_INVALIDPARAMETER:
		return "0x80030057"
	case STG_E_MEDIUMFULL:
		return "0x80030070"
	case STG_E_PROPSETMISMATCHED:
		return "0x800300F0"
	case STG_E_ABNORMALAPIEXIT:
		return "0x800300FA"
	case STG_E_INVALIDHEADER:
		return "0x800300FB"
	case STG_E_INVALIDNAME:
		return "0x800300FC"
	case STG_E_UNKNOWN:
		return "0x800300FD"
	case STG_E_UNIMPLEMENTEDFUNCTION:
		return "0x800300FE"
	case STG_E_INVALIDFLAG:
		return "0x800300FF"
	case STG_E_INUSE:
		return "0x80030100"
	case STG_E_NOTCURRENT:
		return "0x80030101"
	case STG_E_REVERTED:
		return "0x80030102"
	case STG_E_CANTSAVE:
		return "0x80030103"
	case STG_E_OLDFORMAT:
		return "0x80030104"
	case STG_E_OLDDLL:
		return "0x80030105"
	case STG_E_SHAREREQUIRED:
		return "0x80030106"
	case STG_E_NOTFILEBASEDSTORAGE:
		return "0x80030107"
	case STG_E_EXTANTMARSHALLINGS:
		return "0x80030108"
	case STG_E_DOCFILECORRUPT:
		return "0x80030109"
	case STG_E_BADBASEADDRESS:
		return "0x80030110"
	case STG_E_DOCFILETOOLARGE:
		return "0x80030111"
	case STG_E_NOTSIMPLEFORMAT:
		return "0x80030112"
	case STG_E_INCOMPLETE:
		return "0x80030201"
	case STG_E_TERMINATED:
		return "0x80030202"
	case STG_S_CONVERTED:
		return "0x00030200"
	case STG_S_BLOCK:
		return "0x00030201"
	case STG_S_RETRYNOW:
		return "0x00030202"
	case STG_S_MONITORING:
		return "0x00030203"
	case STG_S_MULTIPLEOPENS:
		return "0x00030204"
	case STG_S_CONSOLIDATIONFAILED:
		return "0x00030205"
	case STG_S_CANNOTCONSOLIDATE:
		return "0x00030206"
	case STG_E_STATUS_COPY_PROTECTION_FAILURE:
		return "0x80030305"
	case STG_E_CSS_AUTHENTICATION_FAILURE:
		return "0x80030306"
	case STG_E_CSS_KEY_NOT_PRESENT:
		return "0x80030307"
	case STG_E_CSS_KEY_NOT_ESTABLISHED:
		return "0x80030308"
	case STG_E_CSS_SCRAMBLED_SECTOR:
		return "0x80030309"
	case STG_E_CSS_REGION_MISMATCH:
		return "0x8003030A"
	case STG_E_RESETS_EXHAUSTED:
		return "0x8003030B"
	case RPC_E_CALL_REJECTED:
		return "0x80010001"
	case RPC_E_CALL_CANCELED:
		return "0x80010002"
	case RPC_E_CANTPOST_INSENDCALL:
		return "0x80010003"
	case RPC_E_CANTCALLOUT_INASYNCCALL:
		return "0x80010004"
	case RPC_E_CANTCALLOUT_INEXTERNALCALL:
		return "0x80010005"
	case RPC_E_CONNECTION_TERMINATED:
		return "0x80010006"
	case RPC_E_SERVER_DIED:
		return "0x80010007"
	case RPC_E_CLIENT_DIED:
		return "0x80010008"
	case RPC_E_INVALID_DATAPACKET:
		return "0x80010009"
	case RPC_E_CANTTRANSMIT_CALL:
		return "0x8001000A"
	case RPC_E_CLIENT_CANTMARSHAL_DATA:
		return "0x8001000B"
	case RPC_E_CLIENT_CANTUNMARSHAL_DATA:
		return "0x8001000C"
	case RPC_E_SERVER_CANTMARSHAL_DATA:
		return "0x8001000D"
	case RPC_E_SERVER_CANTUNMARSHAL_DATA:
		return "0x8001000E"
	case RPC_E_INVALID_DATA:
		return "0x8001000F"
	case RPC_E_INVALID_PARAMETER:
		return "0x80010010"
	case RPC_E_CANTCALLOUT_AGAIN:
		return "0x80010011"
	case RPC_E_SERVER_DIED_DNE:
		return "0x80010012"
	case RPC_E_SYS_CALL_FAILED:
		return "0x80010100"
	case RPC_E_OUT_OF_RESOURCES:
		return "0x80010101"
	case RPC_E_ATTEMPTED_MULTITHREAD:
		return "0x80010102"
	case RPC_E_NOT_REGISTERED:
		return "0x80010103"
	case RPC_E_FAULT:
		return "0x80010104"
	case RPC_E_SERVERFAULT:
		return "0x80010105"
	case RPC_E_CHANGED_MODE:
		return "0x80010106"
	case RPC_E_INVALIDMETHOD:
		return "0x80010107"
	case RPC_E_DISCONNECTED:
		return "0x80010108"
	case RPC_E_RETRY:
		return "0x80010109"
	case RPC_E_SERVERCALL_RETRYLATER:
		return "0x8001010A"
	case RPC_E_SERVERCALL_REJECTED:
		return "0x8001010B"
	case RPC_E_INVALID_CALLDATA:
		return "0x8001010C"
	case RPC_E_CANTCALLOUT_ININPUTSYNCCALL:
		return "0x8001010D"
	case RPC_E_WRONG_THREAD:
		return "0x8001010E"
	case RPC_E_THREAD_NOT_INIT:
		return "0x8001010F"
	case RPC_E_VERSION_MISMATCH:
		return "0x80010110"
	case RPC_E_INVALID_HEADER:
		return "0x80010111"
	case RPC_E_INVALID_EXTENSION:
		return "0x80010112"
	case RPC_E_INVALID_IPID:
		return "0x80010113"
	case RPC_E_INVALID_OBJECT:
		return "0x80010114"
	case RPC_S_CALLPENDING:
		return "0x80010115"
	case RPC_S_WAITONTIMER:
		return "0x80010116"
	case RPC_E_CALL_COMPLETE:
		return "0x80010117"
	case RPC_E_UNSECURE_CALL:
		return "0x80010118"
	case RPC_E_TOO_LATE:
		return "0x80010119"
	case RPC_E_NO_GOOD_SECURITY_PACKAGES:
		return "0x8001011A"
	case RPC_E_ACCESS_DENIED:
		return "0x8001011B"
	case RPC_E_REMOTE_DISABLED:
		return "0x8001011C"
	case RPC_E_INVALID_OBJREF:
		return "0x8001011D"
	case RPC_E_NO_CONTEXT:
		return "0x8001011E"
	case RPC_E_TIMEOUT:
		return "0x8001011F"
	case RPC_E_NO_SYNC:
		return "0x80010120"
	case RPC_E_FULLSIC_REQUIRED:
		return "0x80010121"
	case RPC_E_INVALID_STD_NAME:
		return "0x80010122"
	case CO_E_FAILEDTOIMPERSONATE:
		return "0x80010123"
	case CO_E_FAILEDTOGETSECCTX:
		return "0x80010124"
	case CO_E_FAILEDTOOPENTHREADTOKEN:
		return "0x80010125"
	case CO_E_FAILEDTOGETTOKENINFO:
		return "0x80010126"
	case CO_E_TRUSTEEDOESNTMATCHCLIENT:
		return "0x80010127"
	case CO_E_FAILEDTOQUERYCLIENTBLANKET:
		return "0x80010128"
	case CO_E_FAILEDTOSETDACL:
		return "0x80010129"
	case CO_E_ACCESSCHECKFAILED:
		return "0x8001012A"
	case CO_E_NETACCESSAPIFAILED:
		return "0x8001012B"
	case CO_E_WRONGTRUSTEENAMESYNTAX:
		return "0x8001012C"
	case CO_E_INVALIDSID:
		return "0x8001012D"
	case CO_E_CONVERSIONFAILED:
		return "0x8001012E"
	case CO_E_NOMATCHINGSIDFOUND:
		return "0x8001012F"
	case CO_E_LOOKUPACCSIDFAILED:
		return "0x80010130"
	case CO_E_NOMATCHINGNAMEFOUND:
		return "0x80010131"
	case CO_E_LOOKUPACCNAMEFAILED:
		return "0x80010132"
	case CO_E_SETSERLHNDLFAILED:
		return "0x80010133"
	case CO_E_FAILEDTOGETWINDIR:
		return "0x80010134"
	case CO_E_PATHTOOLONG:
		return "0x80010135"
	case CO_E_FAILEDTOGENUUID:
		return "0x80010136"
	case CO_E_FAILEDTOCREATEFILE:
		return "0x80010137"
	case CO_E_FAILEDTOCLOSEHANDLE:
		return "0x80010138"
	case CO_E_EXCEEDSYSACLLIMIT:
		return "0x80010139"
	case CO_E_ACESINWRONGORDER:
		return "0x8001013A"
	case CO_E_INCOMPATIBLESTREAMVERSION:
		return "0x8001013B"
	case CO_E_FAILEDTOOPENPROCESSTOKEN:
		return "0x8001013C"
	case CO_E_DECODEFAILED:
		return "0x8001013D"
	case CO_E_ACNOTINITIALIZED:
		return "0x8001013F"
	case CO_E_CANCEL_DISABLED:
		return "0x80010140"
	case RPC_E_UNEXPECTED:
		return "0x8001FFFF"
	case ERROR_AUDITING_DISABLED:
		return "0xC0090001"
	case ERROR_ALL_SIDS_FILTERED:
		return "0xC0090002"
	case ERROR_BIZRULES_NOT_ENABLED:
		return "0xC0090003"
	case NTE_BAD_UID:
		return "0x80090001"
	case NTE_BAD_HASH:
		return "0x80090002"
	case NTE_BAD_KEY:
		return "0x80090003"
	case NTE_BAD_LEN:
		return "0x80090004"
	case NTE_BAD_DATA:
		return "0x80090005"
	case NTE_BAD_SIGNATURE:
		return "0x80090006"
	case NTE_BAD_VER:
		return "0x80090007"
	case NTE_BAD_ALGID:
		return "0x80090008"
	case NTE_BAD_FLAGS:
		return "0x80090009"
	case NTE_BAD_TYPE:
		return "0x8009000A"
	case NTE_BAD_KEY_STATE:
		return "0x8009000B"
	case NTE_BAD_HASH_STATE:
		return "0x8009000C"
	case NTE_NO_KEY:
		return "0x8009000D"
	case NTE_NO_MEMORY:
		return "0x8009000E"
	case NTE_EXISTS:
		return "0x8009000F"
	case NTE_PERM:
		return "0x80090010"
	case NTE_NOT_FOUND:
		return "0x80090011"
	case NTE_DOUBLE_ENCRYPT:
		return "0x80090012"
	case NTE_BAD_PROVIDER:
		return "0x80090013"
	case NTE_BAD_PROV_TYPE:
		return "0x80090014"
	case NTE_BAD_PUBLIC_KEY:
		return "0x80090015"
	case NTE_BAD_KEYSET:
		return "0x80090016"
	case NTE_PROV_TYPE_NOT_DEF:
		return "0x80090017"
	case NTE_PROV_TYPE_ENTRY_BAD:
		return "0x80090018"
	case NTE_KEYSET_NOT_DEF:
		return "0x80090019"
	case NTE_KEYSET_ENTRY_BAD:
		return "0x8009001A"
	case NTE_PROV_TYPE_NO_MATCH:
		return "0x8009001B"
	case NTE_SIGNATURE_FILE_BAD:
		return "0x8009001C"
	case NTE_PROVIDER_DLL_FAIL:
		return "0x8009001D"
	case NTE_PROV_DLL_NOT_FOUND:
		return "0x8009001E"
	case NTE_BAD_KEYSET_PARAM:
		return "0x8009001F"
	case NTE_FAIL:
		return "0x80090020"
	case NTE_SYS_ERR:
		return "0x80090021"
	case NTE_SILENT_CONTEXT:
		return "0x80090022"
	case NTE_TOKEN_KEYSET_STORAGE_FULL:
		return "0x80090023"
	case NTE_TEMPORARY_PROFILE:
		return "0x80090024"
	case NTE_FIXEDPARAMETER:
		return "0x80090025"
	case NTE_INVALID_HANDLE:
		return "0x80090026"
	case NTE_INVALID_PARAMETER:
		return "0x80090027"
	case NTE_BUFFER_TOO_SMALL:
		return "0x80090028"
	case NTE_NOT_SUPPORTED:
		return "0x80090029"
	case NTE_NO_MORE_ITEMS:
		return "0x8009002A"
	case NTE_BUFFERS_OVERLAP:
		return "0x8009002B"
	case NTE_DECRYPTION_FAILURE:
		return "0x8009002C"
	case NTE_INTERNAL_ERROR:
		return "0x8009002D"
	case NTE_UI_REQUIRED:
		return "0x8009002E"
	case NTE_HMAC_NOT_SUPPORTED:
		return "0x8009002F"
	case SEC_E_INSUFFICIENT_MEMORY:
		return "0x80090300"
	case SEC_E_INVALID_HANDLE:
		return "0x80090301"
	case SEC_E_UNSUPPORTED_FUNCTION:
		return "0x80090302"
	case SEC_E_TARGET_UNKNOWN:
		return "0x80090303"
	case SEC_E_INTERNAL_ERROR:
		return "0x80090304"
	case SEC_E_SECPKG_NOT_FOUND:
		return "0x80090305"
	case SEC_E_NOT_OWNER:
		return "0x80090306"
	case SEC_E_CANNOT_INSTALL:
		return "0x80090307"
	case SEC_E_INVALID_TOKEN:
		return "0x80090308"
	case SEC_E_CANNOT_PACK:
		return "0x80090309"
	case SEC_E_QOP_NOT_SUPPORTED:
		return "0x8009030A"
	case SEC_E_NO_IMPERSONATION:
		return "0x8009030B"
	case SEC_E_LOGON_DENIED:
		return "0x8009030C"
	case SEC_E_UNKNOWN_CREDENTIALS:
		return "0x8009030D"
	case SEC_E_NO_CREDENTIALS:
		return "0x8009030E"
	case SEC_E_MESSAGE_ALTERED:
		return "0x8009030F"
	case SEC_E_OUT_OF_SEQUENCE:
		return "0x80090310"
	case SEC_E_NO_AUTHENTICATING_AUTHORITY:
		return "0x80090311"
	case SEC_I_CONTINUE_NEEDED:
		return "0x00090312"
	case SEC_I_COMPLETE_NEEDED:
		return "0x00090313"
	case SEC_I_COMPLETE_AND_CONTINUE:
		return "0x00090314"
	case SEC_I_LOCAL_LOGON:
		return "0x00090315"
	case SEC_E_BAD_PKGID:
		return "0x80090316"
	case SEC_E_CONTEXT_EXPIRED:
		return "0x80090317"
	case SEC_I_CONTEXT_EXPIRED:
		return "0x00090317"
	case SEC_E_INCOMPLETE_MESSAGE:
		return "0x80090318"
	case SEC_E_INCOMPLETE_CREDENTIALS:
		return "0x80090320"
	case SEC_E_BUFFER_TOO_SMALL:
		return "0x80090321"
	case SEC_I_INCOMPLETE_CREDENTIALS:
		return "0x00090320"
	case SEC_I_RENEGOTIATE:
		return "0x00090321"
	case SEC_E_WRONG_PRINCIPAL:
		return "0x80090322"
	case SEC_I_NO_LSA_CONTEXT:
		return "0x00090323"
	case SEC_E_TIME_SKEW:
		return "0x80090324"
	case SEC_E_UNTRUSTED_ROOT:
		return "0x80090325"
	case SEC_E_ILLEGAL_MESSAGE:
		return "0x80090326"
	case SEC_E_CERT_UNKNOWN:
		return "0x80090327"
	case SEC_E_CERT_EXPIRED:
		return "0x80090328"
	case SEC_E_ENCRYPT_FAILURE:
		return "0x80090329"
	case SEC_E_DECRYPT_FAILURE:
		return "0x80090330"
	case SEC_E_ALGORITHM_MISMATCH:
		return "0x80090331"
	case SEC_E_SECURITY_QOS_FAILED:
		return "0x80090332"
	case SEC_E_UNFINISHED_CONTEXT_DELETED:
		return "0x80090333"
	case SEC_E_NO_TGT_REPLY:
		return "0x80090334"
	case SEC_E_NO_IP_ADDRESSES:
		return "0x80090335"
	case SEC_E_WRONG_CREDENTIAL_HANDLE:
		return "0x80090336"
	case SEC_E_CRYPTO_SYSTEM_INVALID:
		return "0x80090337"
	case SEC_E_MAX_REFERRALS_EXCEEDED:
		return "0x80090338"
	case SEC_E_MUST_BE_KDC:
		return "0x80090339"
	case SEC_E_STRONG_CRYPTO_NOT_SUPPORTED:
		return "0x8009033A"
	case SEC_E_TOO_MANY_PRINCIPALS:
		return "0x8009033B"
	case SEC_E_NO_PA_DATA:
		return "0x8009033C"
	case SEC_E_PKINIT_NAME_MISMATCH:
		return "0x8009033D"
	case SEC_E_SMARTCARD_LOGON_REQUIRED:
		return "0x8009033E"
	case SEC_E_SHUTDOWN_IN_PROGRESS:
		return "0x8009033F"
	case SEC_E_KDC_INVALID_REQUEST:
		return "0x80090340"
	case SEC_E_KDC_UNABLE_TO_REFER:
		return "0x80090341"
	case SEC_E_KDC_UNKNOWN_ETYPE:
		return "0x80090342"
	case SEC_E_UNSUPPORTED_PREAUTH:
		return "0x80090343"
	case SEC_E_DELEGATION_REQUIRED:
		return "0x80090345"
	case SEC_E_BAD_BINDINGS:
		return "0x80090346"
	case SEC_E_MULTIPLE_ACCOUNTS:
		return "0x80090347"
	case SEC_E_NO_KERB_KEY:
		return "0x80090348"
	case SEC_E_CERT_WRONG_USAGE:
		return "0x80090349"
	case SEC_E_DOWNGRADE_DETECTED:
		return "0x80090350"
	case SEC_E_SMARTCARD_CERT_REVOKED:
		return "0x80090351"
	case SEC_E_ISSUING_CA_UNTRUSTED:
		return "0x80090352"
	case SEC_E_REVOCATION_OFFLINE_C:
		return "0x80090353"
	case SEC_E_PKINIT_CLIENT_FAILURE:
		return "0x80090354"
	case SEC_E_SMARTCARD_CERT_EXPIRED:
		return "0x80090355"
	case SEC_E_NO_S4U_PROT_SUPPORT:
		return "0x80090356"
	case SEC_E_CROSSREALM_DELEGATION_FAILURE:
		return "0x80090357"
	case SEC_E_REVOCATION_OFFLINE_KDC:
		return "0x80090358"
	case SEC_E_ISSUING_CA_UNTRUSTED_KDC:
		return "0x80090359"
	case SEC_E_KDC_CERT_EXPIRED:
		return "0x8009035A"
	case SEC_E_KDC_CERT_REVOKED:
		return "0x8009035B"
	case SEC_I_SIGNATURE_NEEDED:
		return "0x0009035C"
	case SEC_E_INVALID_PARAMETER:
		return "0x8009035D"
	case SEC_E_DELEGATION_POLICY:
		return "0x8009035E"
	case SEC_E_POLICY_NLTM_ONLY:
		return "0x8009035F"
	case SEC_I_NO_RENEGOTIATION:
		return "0x00090360"
	case SEC_E_NO_CONTEXT:
		return "0x80090361"
	case SEC_E_PKU2U_CERT_FAILURE:
		return "0x80090362"
	case SEC_E_MUTUAL_AUTH_FAILED:
		return "0x80090363"
	case CRYPT_E_MSG_ERROR:
		return "0x80091001"
	case CRYPT_E_UNKNOWN_ALGO:
		return "0x80091002"
	case CRYPT_E_OID_FORMAT:
		return "0x80091003"
	case CRYPT_E_INVALID_MSG_TYPE:
		return "0x80091004"
	case CRYPT_E_UNEXPECTED_ENCODING:
		return "0x80091005"
	case CRYPT_E_AUTH_ATTR_MISSING:
		return "0x80091006"
	case CRYPT_E_HASH_VALUE:
		return "0x80091007"
	case CRYPT_E_INVALID_INDEX:
		return "0x80091008"
	case CRYPT_E_ALREADY_DECRYPTED:
		return "0x80091009"
	case CRYPT_E_NOT_DECRYPTED:
		return "0x8009100A"
	case CRYPT_E_RECIPIENT_NOT_FOUND:
		return "0x8009100B"
	case CRYPT_E_CONTROL_TYPE:
		return "0x8009100C"
	case CRYPT_E_ISSUER_SERIALNUMBER:
		return "0x8009100D"
	case CRYPT_E_SIGNER_NOT_FOUND:
		return "0x8009100E"
	case CRYPT_E_ATTRIBUTES_MISSING:
		return "0x8009100F"
	case CRYPT_E_STREAM_MSG_NOT_READY:
		return "0x80091010"
	case CRYPT_E_STREAM_INSUFFICIENT_DATA:
		return "0x80091011"
	case CRYPT_I_NEW_PROTECTION_REQUIRED:
		return "0x00091012"
	case CRYPT_E_BAD_LEN:
		return "0x80092001"
	case CRYPT_E_BAD_ENCODE:
		return "0x80092002"
	case CRYPT_E_FILE_ERROR:
		return "0x80092003"
	case CRYPT_E_NOT_FOUND:
		return "0x80092004"
	case CRYPT_E_EXISTS:
		return "0x80092005"
	case CRYPT_E_NO_PROVIDER:
		return "0x80092006"
	case CRYPT_E_SELF_SIGNED:
		return "0x80092007"
	case CRYPT_E_DELETED_PREV:
		return "0x80092008"
	case CRYPT_E_NO_MATCH:
		return "0x80092009"
	case CRYPT_E_UNEXPECTED_MSG_TYPE:
		return "0x8009200A"
	case CRYPT_E_NO_KEY_PROPERTY:
		return "0x8009200B"
	case CRYPT_E_NO_DECRYPT_CERT:
		return "0x8009200C"
	case CRYPT_E_BAD_MSG:
		return "0x8009200D"
	case CRYPT_E_NO_SIGNER:
		return "0x8009200E"
	case CRYPT_E_PENDING_CLOSE:
		return "0x8009200F"
	case CRYPT_E_REVOKED:
		return "0x80092010"
	case CRYPT_E_NO_REVOCATION_DLL:
		return "0x80092011"
	case CRYPT_E_NO_REVOCATION_CHECK:
		return "0x80092012"
	case CRYPT_E_REVOCATION_OFFLINE:
		return "0x80092013"
	case CRYPT_E_NOT_IN_REVOCATION_DATABASE:
		return "0x80092014"
	case CRYPT_E_INVALID_NUMERIC_STRING:
		return "0x80092020"
	case CRYPT_E_INVALID_PRINTABLE_STRING:
		return "0x80092021"
	case CRYPT_E_INVALID_IA5_STRING:
		return "0x80092022"
	case CRYPT_E_INVALID_X500_STRING:
		return "0x80092023"
	case CRYPT_E_NOT_CHAR_STRING:
		return "0x80092024"
	case CRYPT_E_FILERESIZED:
		return "0x80092025"
	case CRYPT_E_SECURITY_SETTINGS:
		return "0x80092026"
	case CRYPT_E_NO_VERIFY_USAGE_DLL:
		return "0x80092027"
	case CRYPT_E_NO_VERIFY_USAGE_CHECK:
		return "0x80092028"
	case CRYPT_E_VERIFY_USAGE_OFFLINE:
		return "0x80092029"
	case CRYPT_E_NOT_IN_CTL:
		return "0x8009202A"
	case CRYPT_E_NO_TRUSTED_SIGNER:
		return "0x8009202B"
	case CRYPT_E_MISSING_PUBKEY_PARA:
		return "0x8009202C"
	case CRYPT_E_OSS_ERROR:
		return "0x80093000"
	case OSS_MORE_BUF:
		return "0x80093001"
	case OSS_NEGATIVE_UINTEGER:
		return "0x80093002"
	case OSS_PDU_RANGE:
		return "0x80093003"
	case OSS_MORE_INPUT:
		return "0x80093004"
	case OSS_DATA_ERROR:
		return "0x80093005"
	case OSS_BAD_ARG:
		return "0x80093006"
	case OSS_BAD_VERSION:
		return "0x80093007"
	case OSS_OUT_MEMORY:
		return "0x80093008"
	case OSS_PDU_MISMATCH:
		return "0x80093009"
	case OSS_LIMITED:
		return "0x8009300A"
	case OSS_BAD_PTR:
		return "0x8009300B"
	case OSS_BAD_TIME:
		return "0x8009300C"
	case OSS_INDEFINITE_NOT_SUPPORTED:
		return "0x8009300D"
	case OSS_MEM_ERROR:
		return "0x8009300E"
	case OSS_BAD_TABLE:
		return "0x8009300F"
	case OSS_TOO_LONG:
		return "0x80093010"
	case OSS_CONSTRAINT_VIOLATED:
		return "0x80093011"
	case OSS_FATAL_ERROR:
		return "0x80093012"
	case OSS_ACCESS_SERIALIZATION_ERROR:
		return "0x80093013"
	case OSS_NULL_TBL:
		return "0x80093014"
	case OSS_NULL_FCN:
		return "0x80093015"
	case OSS_BAD_ENCRULES:
		return "0x80093016"
	case OSS_UNAVAIL_ENCRULES:
		return "0x80093017"
	case OSS_CANT_OPEN_TRACE_WINDOW:
		return "0x80093018"
	case OSS_UNIMPLEMENTED:
		return "0x80093019"
	case OSS_OID_DLL_NOT_LINKED:
		return "0x8009301A"
	case OSS_CANT_OPEN_TRACE_FILE:
		return "0x8009301B"
	case OSS_TRACE_FILE_ALREADY_OPEN:
		return "0x8009301C"
	case OSS_TABLE_MISMATCH:
		return "0x8009301D"
	case OSS_TYPE_NOT_SUPPORTED:
		return "0x8009301E"
	case OSS_REAL_DLL_NOT_LINKED:
		return "0x8009301F"
	case OSS_REAL_CODE_NOT_LINKED:
		return "0x80093020"
	case OSS_OUT_OF_RANGE:
		return "0x80093021"
	case OSS_COPIER_DLL_NOT_LINKED:
		return "0x80093022"
	case OSS_CONSTRAINT_DLL_NOT_LINKED:
		return "0x80093023"
	case OSS_COMPARATOR_DLL_NOT_LINKED:
		return "0x80093024"
	case OSS_COMPARATOR_CODE_NOT_LINKED:
		return "0x80093025"
	case OSS_MEM_MGR_DLL_NOT_LINKED:
		return "0x80093026"
	case OSS_PDV_DLL_NOT_LINKED:
		return "0x80093027"
	case OSS_PDV_CODE_NOT_LINKED:
		return "0x80093028"
	case OSS_API_DLL_NOT_LINKED:
		return "0x80093029"
	case OSS_BERDER_DLL_NOT_LINKED:
		return "0x8009302A"
	case OSS_PER_DLL_NOT_LINKED:
		return "0x8009302B"
	case OSS_OPEN_TYPE_ERROR:
		return "0x8009302C"
	case OSS_MUTEX_NOT_CREATED:
		return "0x8009302D"
	case OSS_CANT_CLOSE_TRACE_FILE:
		return "0x8009302E"
	case CRYPT_E_ASN1_ERROR:
		return "0x80093100"
	case CRYPT_E_ASN1_INTERNAL:
		return "0x80093101"
	case CRYPT_E_ASN1_EOD:
		return "0x80093102"
	case CRYPT_E_ASN1_CORRUPT:
		return "0x80093103"
	case CRYPT_E_ASN1_LARGE:
		return "0x80093104"
	case CRYPT_E_ASN1_CONSTRAINT:
		return "0x80093105"
	case CRYPT_E_ASN1_MEMORY:
		return "0x80093106"
	case CRYPT_E_ASN1_OVERFLOW:
		return "0x80093107"
	case CRYPT_E_ASN1_BADPDU:
		return "0x80093108"
	case CRYPT_E_ASN1_BADARGS:
		return "0x80093109"
	case CRYPT_E_ASN1_BADREAL:
		return "0x8009310A"
	case CRYPT_E_ASN1_BADTAG:
		return "0x8009310B"
	case CRYPT_E_ASN1_CHOICE:
		return "0x8009310C"
	case CRYPT_E_ASN1_RULE:
		return "0x8009310D"
	case CRYPT_E_ASN1_UTF8:
		return "0x8009310E"
	case CRYPT_E_ASN1_PDU_TYPE:
		return "0x80093133"
	case CRYPT_E_ASN1_NYI:
		return "0x80093134"
	case CRYPT_E_ASN1_EXTENDED:
		return "0x80093201"
	case CRYPT_E_ASN1_NOEOD:
		return "0x80093202"
	case CERTSRV_E_BAD_REQUESTSUBJECT:
		return "0x80094001"
	case CERTSRV_E_NO_REQUEST:
		return "0x80094002"
	case CERTSRV_E_BAD_REQUESTSTATUS:
		return "0x80094003"
	case CERTSRV_E_PROPERTY_EMPTY:
		return "0x80094004"
	case CERTSRV_E_INVALID_CA_CERTIFICATE:
		return "0x80094005"
	case CERTSRV_E_SERVER_SUSPENDED:
		return "0x80094006"
	case CERTSRV_E_ENCODING_LENGTH:
		return "0x80094007"
	case CERTSRV_E_ROLECONFLICT:
		return "0x80094008"
	case CERTSRV_E_RESTRICTEDOFFICER:
		return "0x80094009"
	case CERTSRV_E_KEY_ARCHIVAL_NOT_CONFIGURED:
		return "0x8009400A"
	case CERTSRV_E_NO_VALID_KRA:
		return "0x8009400B"
	case CERTSRV_E_BAD_REQUEST_KEY_ARCHIVAL:
		return "0x8009400C"
	case CERTSRV_E_NO_CAADMIN_DEFINED:
		return "0x8009400D"
	case CERTSRV_E_BAD_RENEWAL_CERT_ATTRIBUTE:
		return "0x8009400E"
	case CERTSRV_E_NO_DB_SESSIONS:
		return "0x8009400F"
	case CERTSRV_E_ALIGNMENT_FAULT:
		return "0x80094010"
	case CERTSRV_E_ENROLL_DENIED:
		return "0x80094011"
	case CERTSRV_E_TEMPLATE_DENIED:
		return "0x80094012"
	case CERTSRV_E_DOWNLEVEL_DC_SSL_OR_UPGRADE:
		return "0x80094013"
	case CERTSRV_E_ADMIN_DENIED_REQUEST:
		return "0x80094014"
	case CERTSRV_E_NO_POLICY_SERVER:
		return "0x80094015"
	case CERTSRV_E_UNSUPPORTED_CERT_TYPE:
		return "0x80094800"
	case CERTSRV_E_NO_CERT_TYPE:
		return "0x80094801"
	case CERTSRV_E_TEMPLATE_CONFLICT:
		return "0x80094802"
	case CERTSRV_E_SUBJECT_ALT_NAME_REQUIRED:
		return "0x80094803"
	case CERTSRV_E_ARCHIVED_KEY_REQUIRED:
		return "0x80094804"
	case CERTSRV_E_SMIME_REQUIRED:
		return "0x80094805"
	case CERTSRV_E_BAD_RENEWAL_SUBJECT:
		return "0x80094806"
	case CERTSRV_E_BAD_TEMPLATE_VERSION:
		return "0x80094807"
	case CERTSRV_E_TEMPLATE_POLICY_REQUIRED:
		return "0x80094808"
	case CERTSRV_E_SIGNATURE_POLICY_REQUIRED:
		return "0x80094809"
	case CERTSRV_E_SIGNATURE_COUNT:
		return "0x8009480A"
	case CERTSRV_E_SIGNATURE_REJECTED:
		return "0x8009480B"
	case CERTSRV_E_ISSUANCE_POLICY_REQUIRED:
		return "0x8009480C"
	case CERTSRV_E_SUBJECT_UPN_REQUIRED:
		return "0x8009480D"
	case CERTSRV_E_SUBJECT_DIRECTORY_GUID_REQUIRED:
		return "0x8009480E"
	case CERTSRV_E_SUBJECT_DNS_REQUIRED:
		return "0x8009480F"
	case CERTSRV_E_ARCHIVED_KEY_UNEXPECTED:
		return "0x80094810"
	case CERTSRV_E_KEY_LENGTH:
		return "0x80094811"
	case CERTSRV_E_SUBJECT_EMAIL_REQUIRED:
		return "0x80094812"
	case CERTSRV_E_UNKNOWN_CERT_TYPE:
		return "0x80094813"
	case CERTSRV_E_CERT_TYPE_OVERLAP:
		return "0x80094814"
	case CERTSRV_E_TOO_MANY_SIGNATURES:
		return "0x80094815"
	case XENROLL_E_KEY_NOT_EXPORTABLE:
		return "0x80095000"
	case XENROLL_E_CANNOT_ADD_ROOT_CERT:
		return "0x80095001"
	case XENROLL_E_RESPONSE_KA_HASH_NOT_FOUND:
		return "0x80095002"
	case XENROLL_E_RESPONSE_UNEXPECTED_KA_HASH:
		return "0x80095003"
	case XENROLL_E_RESPONSE_KA_HASH_MISMATCH:
		return "0x80095004"
	case XENROLL_E_KEYSPEC_SMIME_MISMATCH:
		return "0x80095005"
	case TRUST_E_SYSTEM_ERROR:
		return "0x80096001"
	case TRUST_E_NO_SIGNER_CERT:
		return "0x80096002"
	case TRUST_E_COUNTER_SIGNER:
		return "0x80096003"
	case TRUST_E_CERT_SIGNATURE:
		return "0x80096004"
	case TRUST_E_TIME_STAMP:
		return "0x80096005"
	case TRUST_E_BAD_DIGEST:
		return "0x80096010"
	case TRUST_E_BASIC_CONSTRAINTS:
		return "0x80096019"
	case TRUST_E_FINANCIAL_CRITERIA:
		return "0x8009601E"
	case MSSIPOTF_E_OUTOFMEMRANGE:
		return "0x80097001"
	case MSSIPOTF_E_CANTGETOBJECT:
		return "0x80097002"
	case MSSIPOTF_E_NOHEADTABLE:
		return "0x80097003"
	case MSSIPOTF_E_BAD_MAGICNUMBER:
		return "0x80097004"
	case MSSIPOTF_E_BAD_OFFSET_TABLE:
		return "0x80097005"
	case MSSIPOTF_E_TABLE_TAGORDER:
		return "0x80097006"
	case MSSIPOTF_E_TABLE_LONGWORD:
		return "0x80097007"
	case MSSIPOTF_E_BAD_FIRST_TABLE_PLACEMENT:
		return "0x80097008"
	case MSSIPOTF_E_TABLES_OVERLAP:
		return "0x80097009"
	case MSSIPOTF_E_TABLE_PADBYTES:
		return "0x8009700A"
	case MSSIPOTF_E_FILETOOSMALL:
		return "0x8009700B"
	case MSSIPOTF_E_TABLE_CHECKSUM:
		return "0x8009700C"
	case MSSIPOTF_E_FILE_CHECKSUM:
		return "0x8009700D"
	case MSSIPOTF_E_FAILED_POLICY:
		return "0x80097010"
	case MSSIPOTF_E_FAILED_HINTS_CHECK:
		return "0x80097011"
	case MSSIPOTF_E_NOT_OPENTYPE:
		return "0x80097012"
	case MSSIPOTF_E_FILE:
		return "0x80097013"
	case MSSIPOTF_E_CRYPT:
		return "0x80097014"
	case MSSIPOTF_E_BADVERSION:
		return "0x80097015"
	case MSSIPOTF_E_DSIG_STRUCTURE:
		return "0x80097016"
	case MSSIPOTF_E_PCONST_CHECK:
		return "0x80097017"
	case MSSIPOTF_E_STRUCTURE:
		return "0x80097018"
	case ERROR_CRED_REQUIRES_CONFIRMATION:
		return "0x80097019"
	case NTE_OP_OK:
		return "0"
	case TRUST_E_PROVIDER_UNKNOWN:
		return "0x800B0001"
	case TRUST_E_ACTION_UNKNOWN:
		return "0x800B0002"
	case TRUST_E_SUBJECT_FORM_UNKNOWN:
		return "0x800B0003"
	case TRUST_E_SUBJECT_NOT_TRUSTED:
		return "0x800B0004"
	case DIGSIG_E_ENCODE:
		return "0x800B0005"
	case DIGSIG_E_DECODE:
		return "0x800B0006"
	case DIGSIG_E_EXTENSIBILITY:
		return "0x800B0007"
	case DIGSIG_E_CRYPTO:
		return "0x800B0008"
	case PERSIST_E_SIZEDEFINITE:
		return "0x800B0009"
	case PERSIST_E_SIZEINDEFINITE:
		return "0x800B000A"
	case PERSIST_E_NOTSELFSIZING:
		return "0x800B000B"
	case TRUST_E_NOSIGNATURE:
		return "0x800B0100"
	case CERT_E_EXPIRED:
		return "0x800B0101"
	case CERT_E_VALIDITYPERIODNESTING:
		return "0x800B0102"
	case CERT_E_ROLE:
		return "0x800B0103"
	case CERT_E_PATHLENCONST:
		return "0x800B0104"
	case CERT_E_CRITICAL:
		return "0x800B0105"
	case CERT_E_PURPOSE:
		return "0x800B0106"
	case CERT_E_ISSUERCHAINING:
		return "0x800B0107"
	case CERT_E_MALFORMED:
		return "0x800B0108"
	case CERT_E_UNTRUSTEDROOT:
		return "0x800B0109"
	case CERT_E_CHAINING:
		return "0x800B010A"
	case TRUST_E_FAIL:
		return "0x800B010B"
	case CERT_E_REVOKED:
		return "0x800B010C"
	case CERT_E_UNTRUSTEDTESTROOT:
		return "0x800B010D"
	case CERT_E_REVOCATION_FAILURE:
		return "0x800B010E"
	case CERT_E_CN_NO_MATCH:
		return "0x800B010F"
	case CERT_E_WRONG_USAGE:
		return "0x800B0110"
	case TRUST_E_EXPLICIT_DISTRUST:
		return "0x800B0111"
	case CERT_E_UNTRUSTEDCA:
		return "0x800B0112"
	case CERT_E_INVALID_POLICY:
		return "0x800B0113"
	case CERT_E_INVALID_NAME:
		return "0x800B0114"
	case SPAPI_E_EXPECTED_SECTION_NAME:
		return "0x800F0000"
	case SPAPI_E_BAD_SECTION_NAME_LINE:
		return "0x800F0001"
	case SPAPI_E_SECTION_NAME_TOO_LONG:
		return "0x800F0002"
	case SPAPI_E_GENERAL_SYNTAX:
		return "0x800F0003"
	case SPAPI_E_WRONG_INF_STYLE:
		return "0x800F0100"
	case SPAPI_E_SECTION_NOT_FOUND:
		return "0x800F0101"
	case SPAPI_E_LINE_NOT_FOUND:
		return "0x800F0102"
	case SPAPI_E_NO_BACKUP:
		return "0x800F0103"
	case SPAPI_E_NO_ASSOCIATED_CLASS:
		return "0x800F0200"
	case SPAPI_E_CLASS_MISMATCH:
		return "0x800F0201"
	case SPAPI_E_DUPLICATE_FOUND:
		return "0x800F0202"
	case SPAPI_E_NO_DRIVER_SELECTED:
		return "0x800F0203"
	case SPAPI_E_KEY_DOES_NOT_EXIST:
		return "0x800F0204"
	case SPAPI_E_INVALID_DEVINST_NAME:
		return "0x800F0205"
	case SPAPI_E_INVALID_CLASS:
		return "0x800F0206"
	case SPAPI_E_DEVINST_ALREADY_EXISTS:
		return "0x800F0207"
	case SPAPI_E_DEVINFO_NOT_REGISTERED:
		return "0x800F0208"
	case SPAPI_E_INVALID_REG_PROPERTY:
		return "0x800F0209"
	case SPAPI_E_NO_INF:
		return "0x800F020A"
	case SPAPI_E_NO_SUCH_DEVINST:
		return "0x800F020B"
	case SPAPI_E_CANT_LOAD_CLASS_ICON:
		return "0x800F020C"
	case SPAPI_E_INVALID_CLASS_INSTALLER:
		return "0x800F020D"
	case SPAPI_E_DI_DO_DEFAULT:
		return "0x800F020E"
	case SPAPI_E_DI_NOFILECOPY:
		return "0x800F020F"
	case SPAPI_E_INVALID_HWPROFILE:
		return "0x800F0210"
	case SPAPI_E_NO_DEVICE_SELECTED:
		return "0x800F0211"
	case SPAPI_E_DEVINFO_LIST_LOCKED:
		return "0x800F0212"
	case SPAPI_E_DEVINFO_DATA_LOCKED:
		return "0x800F0213"
	case SPAPI_E_DI_BAD_PATH:
		return "0x800F0214"
	case SPAPI_E_NO_CLASSINSTALL_PARAMS:
		return "0x800F0215"
	case SPAPI_E_FILEQUEUE_LOCKED:
		return "0x800F0216"
	case SPAPI_E_BAD_SERVICE_INSTALLSECT:
		return "0x800F0217"
	case SPAPI_E_NO_CLASS_DRIVER_LIST:
		return "0x800F0218"
	case SPAPI_E_NO_ASSOCIATED_SERVICE:
		return "0x800F0219"
	case SPAPI_E_NO_DEFAULT_DEVICE_INTERFACE:
		return "0x800F021A"
	case SPAPI_E_DEVICE_INTERFACE_ACTIVE:
		return "0x800F021B"
	case SPAPI_E_DEVICE_INTERFACE_REMOVED:
		return "0x800F021C"
	case SPAPI_E_BAD_INTERFACE_INSTALLSECT:
		return "0x800F021D"
	case SPAPI_E_NO_SUCH_INTERFACE_CLASS:
		return "0x800F021E"
	case SPAPI_E_INVALID_REFERENCE_STRING:
		return "0x800F021F"
	case SPAPI_E_INVALID_MACHINENAME:
		return "0x800F0220"
	case SPAPI_E_REMOTE_COMM_FAILURE:
		return "0x800F0221"
	case SPAPI_E_MACHINE_UNAVAILABLE:
		return "0x800F0222"
	case SPAPI_E_NO_CONFIGMGR_SERVICES:
		return "0x800F0223"
	case SPAPI_E_INVALID_PROPPAGE_PROVIDER:
		return "0x800F0224"
	case SPAPI_E_NO_SUCH_DEVICE_INTERFACE:
		return "0x800F0225"
	case SPAPI_E_DI_POSTPROCESSING_REQUIRED:
		return "0x800F0226"
	case SPAPI_E_INVALID_COINSTALLER:
		return "0x800F0227"
	case SPAPI_E_NO_COMPAT_DRIVERS:
		return "0x800F0228"
	case SPAPI_E_NO_DEVICE_ICON:
		return "0x800F0229"
	case SPAPI_E_INVALID_INF_LOGCONFIG:
		return "0x800F022A"
	case SPAPI_E_DI_DONT_INSTALL:
		return "0x800F022B"
	case SPAPI_E_INVALID_FILTER_DRIVER:
		return "0x800F022C"
	case SPAPI_E_NON_WINDOWS_NT_DRIVER:
		return "0x800F022D"
	case SPAPI_E_NON_WINDOWS_DRIVER:
		return "0x800F022E"
	case SPAPI_E_NO_CATALOG_FOR_OEM_INF:
		return "0x800F022F"
	case SPAPI_E_DEVINSTALL_QUEUE_NONNATIVE:
		return "0x800F0230"
	case SPAPI_E_NOT_DISABLEABLE:
		return "0x800F0231"
	case SPAPI_E_CANT_REMOVE_DEVINST:
		return "0x800F0232"
	case SPAPI_E_INVALID_TARGET:
		return "0x800F0233"
	case SPAPI_E_DRIVER_NONNATIVE:
		return "0x800F0234"
	case SPAPI_E_IN_WOW64:
		return "0x800F0235"
	case SPAPI_E_SET_SYSTEM_RESTORE_POINT:
		return "0x800F0236"
	case SPAPI_E_INCORRECTLY_COPIED_INF:
		return "0x800F0237"
	case SPAPI_E_SCE_DISABLED:
		return "0x800F0238"
	case SPAPI_E_UNKNOWN_EXCEPTION:
		return "0x800F0239"
	case SPAPI_E_PNP_REGISTRY_ERROR:
		return "0x800F023A"
	case SPAPI_E_REMOTE_REQUEST_UNSUPPORTED:
		return "0x800F023B"
	case SPAPI_E_NOT_AN_INSTALLED_OEM_INF:
		return "0x800F023C"
	case SPAPI_E_INF_IN_USE_BY_DEVICES:
		return "0x800F023D"
	case SPAPI_E_DI_FUNCTION_OBSOLETE:
		return "0x800F023E"
	case SPAPI_E_NO_AUTHENTICODE_CATALOG:
		return "0x800F023F"
	case SPAPI_E_AUTHENTICODE_DISALLOWED:
		return "0x800F0240"
	case SPAPI_E_AUTHENTICODE_TRUSTED_PUBLISHER:
		return "0x800F0241"
	case SPAPI_E_AUTHENTICODE_TRUST_NOT_ESTABLISHED:
		return "0x800F0242"
	case SPAPI_E_AUTHENTICODE_PUBLISHER_NOT_TRUSTED:
		return "0x800F0243"
	case SPAPI_E_SIGNATURE_OSATTRIBUTE_MISMATCH:
		return "0x800F0244"
	case SPAPI_E_ONLY_VALIDATE_VIA_AUTHENTICODE:
		return "0x800F0245"
	case SPAPI_E_DEVICE_INSTALLER_NOT_READY:
		return "0x800F0246"
	case SPAPI_E_DRIVER_STORE_ADD_FAILED:
		return "0x800F0247"
	case SPAPI_E_DEVICE_INSTALL_BLOCKED:
		return "0x800F0248"
	case SPAPI_E_DRIVER_INSTALL_BLOCKED:
		return "0x800F0249"
	case SPAPI_E_WRONG_INF_TYPE:
		return "0x800F024A"
	case SPAPI_E_FILE_HASH_NOT_IN_CATALOG:
		return "0x800F024B"
	case SPAPI_E_DRIVER_STORE_DELETE_FAILED:
		return "0x800F024C"
	case SPAPI_E_UNRECOVERABLE_STACK_OVERFLOW:
		return "0x800F0300"
	case SPAPI_E_ERROR_NOT_INSTALLED:
		return "0x800F1000"
	case SCARD_F_INTERNAL_ERROR:
		return "0x80100001"
	case SCARD_E_CANCELLED:
		return "0x80100002"
	case SCARD_E_INVALID_HANDLE:
		return "0x80100003"
	case SCARD_E_INVALID_PARAMETER:
		return "0x80100004"
	case SCARD_E_INVALID_TARGET:
		return "0x80100005"
	case SCARD_E_NO_MEMORY:
		return "0x80100006"
	case SCARD_F_WAITED_TOO_LONG:
		return "0x80100007"
	case SCARD_E_INSUFFICIENT_BUFFER:
		return "0x80100008"
	case SCARD_E_UNKNOWN_READER:
		return "0x80100009"
	case SCARD_E_TIMEOUT:
		return "0x8010000A"
	case SCARD_E_SHARING_VIOLATION:
		return "0x8010000B"
	case SCARD_E_NO_SMARTCARD:
		return "0x8010000C"
	case SCARD_E_UNKNOWN_CARD:
		return "0x8010000D"
	case SCARD_E_CANT_DISPOSE:
		return "0x8010000E"
	case SCARD_E_PROTO_MISMATCH:
		return "0x8010000F"
	case SCARD_E_NOT_READY:
		return "0x80100010"
	case SCARD_E_INVALID_VALUE:
		return "0x80100011"
	case SCARD_E_SYSTEM_CANCELLED:
		return "0x80100012"
	case SCARD_F_COMM_ERROR:
		return "0x80100013"
	case SCARD_F_UNKNOWN_ERROR:
		return "0x80100014"
	case SCARD_E_INVALID_ATR:
		return "0x80100015"
	case SCARD_E_NOT_TRANSACTED:
		return "0x80100016"
	case SCARD_E_READER_UNAVAILABLE:
		return "0x80100017"
	case SCARD_P_SHUTDOWN:
		return "0x80100018"
	case SCARD_E_PCI_TOO_SMALL:
		return "0x80100019"
	case SCARD_E_READER_UNSUPPORTED:
		return "0x8010001A"
	case SCARD_E_DUPLICATE_READER:
		return "0x8010001B"
	case SCARD_E_CARD_UNSUPPORTED:
		return "0x8010001C"
	case SCARD_E_NO_SERVICE:
		return "0x8010001D"
	case SCARD_E_SERVICE_STOPPED:
		return "0x8010001E"
	case SCARD_E_UNEXPECTED:
		return "0x8010001F"
	case SCARD_E_ICC_INSTALLATION:
		return "0x80100020"
	case SCARD_E_ICC_CREATEORDER:
		return "0x80100021"
	case SCARD_E_UNSUPPORTED_FEATURE:
		return "0x80100022"
	case SCARD_E_DIR_NOT_FOUND:
		return "0x80100023"
	case SCARD_E_FILE_NOT_FOUND:
		return "0x80100024"
	case SCARD_E_NO_DIR:
		return "0x80100025"
	case SCARD_E_NO_FILE:
		return "0x80100026"
	case SCARD_E_NO_ACCESS:
		return "0x80100027"
	case SCARD_E_WRITE_TOO_MANY:
		return "0x80100028"
	case SCARD_E_BAD_SEEK:
		return "0x80100029"
	case SCARD_E_INVALID_CHV:
		return "0x8010002A"
	case SCARD_E_UNKNOWN_RES_MNG:
		return "0x8010002B"
	case SCARD_E_NO_SUCH_CERTIFICATE:
		return "0x8010002C"
	case SCARD_E_CERTIFICATE_UNAVAILABLE:
		return "0x8010002D"
	case SCARD_E_NO_READERS_AVAILABLE:
		return "0x8010002E"
	case SCARD_E_COMM_DATA_LOST:
		return "0x8010002F"
	case SCARD_E_NO_KEY_CONTAINER:
		return "0x80100030"
	case SCARD_E_SERVER_TOO_BUSY:
		return "0x80100031"
	case SCARD_E_PIN_CACHE_EXPIRED:
		return "0x80100032"
	case SCARD_E_NO_PIN_CACHE:
		return "0x80100033"
	case SCARD_E_READ_ONLY_CARD:
		return "0x80100034"
	case SCARD_W_UNSUPPORTED_CARD:
		return "0x80100065"
	case SCARD_W_UNRESPONSIVE_CARD:
		return "0x80100066"
	case SCARD_W_UNPOWERED_CARD:
		return "0x80100067"
	case SCARD_W_RESET_CARD:
		return "0x80100068"
	case SCARD_W_REMOVED_CARD:
		return "0x80100069"
	case SCARD_W_SECURITY_VIOLATION:
		return "0x8010006A"
	case SCARD_W_WRONG_CHV:
		return "0x8010006B"
	case SCARD_W_CHV_BLOCKED:
		return "0x8010006C"
	case SCARD_W_EOF:
		return "0x8010006D"
	case SCARD_W_CANCELLED_BY_USER:
		return "0x8010006E"
	case SCARD_W_CARD_NOT_AUTHENTICATED:
		return "0x8010006F"
	case SCARD_W_CACHE_ITEM_NOT_FOUND:
		return "0x80100070"
	case SCARD_W_CACHE_ITEM_STALE:
		return "0x80100071"
	case SCARD_W_CACHE_ITEM_TOO_BIG:
		return "0x80100072"
	case COMADMIN_E_OBJECTERRORS:
		return "0x80110401"
	case COMADMIN_E_OBJECTINVALID:
		return "0x80110402"
	case COMADMIN_E_KEYMISSING:
		return "0x80110403"
	case COMADMIN_E_ALREADYINSTALLED:
		return "0x80110404"
	case COMADMIN_E_APP_FILE_WRITEFAIL:
		return "0x80110407"
	case COMADMIN_E_APP_FILE_READFAIL:
		return "0x80110408"
	case COMADMIN_E_APP_FILE_VERSION:
		return "0x80110409"
	case COMADMIN_E_BADPATH:
		return "0x8011040A"
	case COMADMIN_E_APPLICATIONEXISTS:
		return "0x8011040B"
	case COMADMIN_E_ROLEEXISTS:
		return "0x8011040C"
	case COMADMIN_E_CANTCOPYFILE:
		return "0x8011040D"
	case COMADMIN_E_NOUSER:
		return "0x8011040F"
	case COMADMIN_E_INVALIDUSERIDS:
		return "0x80110410"
	case COMADMIN_E_NOREGISTRYCLSID:
		return "0x80110411"
	case COMADMIN_E_BADREGISTRYPROGID:
		return "0x80110412"
	case COMADMIN_E_AUTHENTICATIONLEVEL:
		return "0x80110413"
	case COMADMIN_E_USERPASSWDNOTVALID:
		return "0x80110414"
	case COMADMIN_E_CLSIDORIIDMISMATCH:
		return "0x80110418"
	case COMADMIN_E_REMOTEINTERFACE:
		return "0x80110419"
	case COMADMIN_E_DLLREGISTERSERVER:
		return "0x8011041A"
	case COMADMIN_E_NOSERVERSHARE:
		return "0x8011041B"
	case COMADMIN_E_DLLLOADFAILED:
		return "0x8011041D"
	case COMADMIN_E_BADREGISTRYLIBID:
		return "0x8011041E"
	case COMADMIN_E_APPDIRNOTFOUND:
		return "0x8011041F"
	case COMADMIN_E_REGISTRARFAILED:
		return "0x80110423"
	case COMADMIN_E_COMPFILE_DOESNOTEXIST:
		return "0x80110424"
	case COMADMIN_E_COMPFILE_LOADDLLFAIL:
		return "0x80110425"
	case COMADMIN_E_COMPFILE_GETCLASSOBJ:
		return "0x80110426"
	case COMADMIN_E_COMPFILE_CLASSNOTAVAIL:
		return "0x80110427"
	case COMADMIN_E_COMPFILE_BADTLB:
		return "0x80110428"
	case COMADMIN_E_COMPFILE_NOTINSTALLABLE:
		return "0x80110429"
	case COMADMIN_E_NOTCHANGEABLE:
		return "0x8011042A"
	case COMADMIN_E_NOTDELETEABLE:
		return "0x8011042B"
	case COMADMIN_E_SESSION:
		return "0x8011042C"
	case COMADMIN_E_COMP_MOVE_LOCKED:
		return "0x8011042D"
	case COMADMIN_E_COMP_MOVE_BAD_DEST:
		return "0x8011042E"
	case COMADMIN_E_REGISTERTLB:
		return "0x80110430"
	case COMADMIN_E_SYSTEMAPP:
		return "0x80110433"
	case COMADMIN_E_COMPFILE_NOREGISTRAR:
		return "0x80110434"
	case COMADMIN_E_COREQCOMPINSTALLED:
		return "0x80110435"
	case COMADMIN_E_SERVICENOTINSTALLED:
		return "0x80110436"
	case COMADMIN_E_PROPERTYSAVEFAILED:
		return "0x80110437"
	case COMADMIN_E_OBJECTEXISTS:
		return "0x80110438"
	case COMADMIN_E_COMPONENTEXISTS:
		return "0x80110439"
	case COMADMIN_E_REGFILE_CORRUPT:
		return "0x8011043B"
	case COMADMIN_E_PROPERTY_OVERFLOW:
		return "0x8011043C"
	case COMADMIN_E_NOTINREGISTRY:
		return "0x8011043E"
	case COMADMIN_E_OBJECTNOTPOOLABLE:
		return "0x8011043F"
	case COMADMIN_E_APPLID_MATCHES_CLSID:
		return "0x80110446"
	case COMADMIN_E_ROLE_DOES_NOT_EXIST:
		return "0x80110447"
	case COMADMIN_E_START_APP_NEEDS_COMPONENTS:
		return "0x80110448"
	case COMADMIN_E_REQUIRES_DIFFERENT_PLATFORM:
		return "0x80110449"
	case COMADMIN_E_CAN_NOT_EXPORT_APP_PROXY:
		return "0x8011044A"
	case COMADMIN_E_CAN_NOT_START_APP:
		return "0x8011044B"
	case COMADMIN_E_CAN_NOT_EXPORT_SYS_APP:
		return "0x8011044C"
	case COMADMIN_E_CANT_SUBSCRIBE_TO_COMPONENT:
		return "0x8011044D"
	case COMADMIN_E_EVENTCLASS_CANT_BE_SUBSCRIBER:
		return "0x8011044E"
	case COMADMIN_E_LIB_APP_PROXY_INCOMPATIBLE:
		return "0x8011044F"
	case COMADMIN_E_BASE_PARTITION_ONLY:
		return "0x80110450"
	case COMADMIN_E_START_APP_DISABLED:
		return "0x80110451"
	case COMADMIN_E_CAT_DUPLICATE_PARTITION_NAME:
		return "0x80110457"
	case COMADMIN_E_CAT_INVALID_PARTITION_NAME:
		return "0x80110458"
	case COMADMIN_E_CAT_PARTITION_IN_USE:
		return "0x80110459"
	case COMADMIN_E_FILE_PARTITION_DUPLICATE_FILES:
		return "0x8011045A"
	case COMADMIN_E_CAT_IMPORTED_COMPONENTS_NOT_ALLOWED:
		return "0x8011045B"
	case COMADMIN_E_AMBIGUOUS_APPLICATION_NAME:
		return "0x8011045C"
	case COMADMIN_E_AMBIGUOUS_PARTITION_NAME:
		return "0x8011045D"
	case COMADMIN_E_REGDB_NOTINITIALIZED:
		return "0x80110472"
	case COMADMIN_E_REGDB_NOTOPEN:
		return "0x80110473"
	case COMADMIN_E_REGDB_SYSTEMERR:
		return "0x80110474"
	case COMADMIN_E_REGDB_ALREADYRUNNING:
		return "0x80110475"
	case COMADMIN_E_MIG_VERSIONNOTSUPPORTED:
		return "0x80110480"
	case COMADMIN_E_MIG_SCHEMANOTFOUND:
		return "0x80110481"
	case COMADMIN_E_CAT_BITNESSMISMATCH:
		return "0x80110482"
	case COMADMIN_E_CAT_UNACCEPTABLEBITNESS:
		return "0x80110483"
	case COMADMIN_E_CAT_WRONGAPPBITNESS:
		return "0x80110484"
	case COMADMIN_E_CAT_PAUSE_RESUME_NOT_SUPPORTED:
		return "0x80110485"
	case COMADMIN_E_CAT_SERVERFAULT:
		return "0x80110486"
	case COMQC_E_APPLICATION_NOT_QUEUED:
		return "0x80110600"
	case COMQC_E_NO_QUEUEABLE_INTERFACES:
		return "0x80110601"
	case COMQC_E_QUEUING_SERVICE_NOT_AVAILABLE:
		return "0x80110602"
	case COMQC_E_NO_IPERSISTSTREAM:
		return "0x80110603"
	case COMQC_E_BAD_MESSAGE:
		return "0x80110604"
	case COMQC_E_UNAUTHENTICATED:
		return "0x80110605"
	case COMQC_E_UNTRUSTED_ENQUEUER:
		return "0x80110606"
	case MSDTC_E_DUPLICATE_RESOURCE:
		return "0x80110701"
	case COMADMIN_E_OBJECT_PARENT_MISSING:
		return "0x80110808"
	case COMADMIN_E_OBJECT_DOES_NOT_EXIST:
		return "0x80110809"
	case COMADMIN_E_APP_NOT_RUNNING:
		return "0x8011080A"
	case COMADMIN_E_INVALID_PARTITION:
		return "0x8011080B"
	case COMADMIN_E_SVCAPP_NOT_POOLABLE_OR_RECYCLABLE:
		return "0x8011080D"
	case COMADMIN_E_USER_IN_SET:
		return "0x8011080E"
	case COMADMIN_E_CANTRECYCLELIBRARYAPPS:
		return "0x8011080F"
	case COMADMIN_E_CANTRECYCLESERVICEAPPS:
		return "0x80110811"
	case COMADMIN_E_PROCESSALREADYRECYCLED:
		return "0x80110812"
	case COMADMIN_E_PAUSEDPROCESSMAYNOTBERECYCLED:
		return "0x80110813"
	case COMADMIN_E_CANTMAKEINPROCSERVICE:
		return "0x80110814"
	case COMADMIN_E_PROGIDINUSEBYCLSID:
		return "0x80110815"
	case COMADMIN_E_DEFAULT_PARTITION_NOT_IN_SET:
		return "0x80110816"
	case COMADMIN_E_RECYCLEDPROCESSMAYNOTBEPAUSED:
		return "0x80110817"
	case COMADMIN_E_PARTITION_ACCESSDENIED:
		return "0x80110818"
	case COMADMIN_E_PARTITION_MSI_ONLY:
		return "0x80110819"
	case COMADMIN_E_LEGACYCOMPS_NOT_ALLOWED_IN_1_0_FORMAT:
		return "0x8011081A"
	case COMADMIN_E_LEGACYCOMPS_NOT_ALLOWED_IN_NONBASE_PARTITIONS:
		return "0x8011081B"
	case COMADMIN_E_COMP_MOVE_SOURCE:
		return "0x8011081C"
	case COMADMIN_E_COMP_MOVE_DEST:
		return "0x8011081D"
	case COMADMIN_E_COMP_MOVE_PRIVATE:
		return "0x8011081E"
	case COMADMIN_E_BASEPARTITION_REQUIRED_IN_SET:
		return "0x8011081F"
	case COMADMIN_E_CANNOT_ALIAS_EVENTCLASS:
		return "0x80110820"
	case COMADMIN_E_PRIVATE_ACCESSDENIED:
		return "0x80110821"
	case COMADMIN_E_SAFERINVALID:
		return "0x80110822"
	case COMADMIN_E_REGISTRY_ACCESSDENIED:
		return "0x80110823"
	case COMADMIN_E_PARTITIONS_DISABLED:
		return "0x80110824"
	case ERROR_FLT_IO_COMPLETE:
		return "0x001F0001"
	case ERROR_FLT_NO_HANDLER_DEFINED:
		return "0x801F0001"
	case ERROR_FLT_CONTEXT_ALREADY_DEFINED:
		return "0x801F0002"
	case ERROR_FLT_INVALID_ASYNCHRONOUS_REQUEST:
		return "0x801F0003"
	case ERROR_FLT_DISALLOW_FAST_IO:
		return "0x801F0004"
	case ERROR_FLT_INVALID_NAME_REQUEST:
		return "0x801F0005"
	case ERROR_FLT_NOT_SAFE_TO_POST_OPERATION:
		return "0x801F0006"
	case ERROR_FLT_NOT_INITIALIZED:
		return "0x801F0007"
	case ERROR_FLT_FILTER_NOT_READY:
		return "0x801F0008"
	case ERROR_FLT_POST_OPERATION_CLEANUP:
		return "0x801F0009"
	case ERROR_FLT_INTERNAL_ERROR:
		return "0x801F000A"
	case ERROR_FLT_DELETING_OBJECT:
		return "0x801F000B"
	case ERROR_FLT_MUST_BE_NONPAGED_POOL:
		return "0x801F000C"
	case ERROR_FLT_DUPLICATE_ENTRY:
		return "0x801F000D"
	case ERROR_FLT_CBDQ_DISABLED:
		return "0x801F000E"
	case ERROR_FLT_DO_NOT_ATTACH:
		return "0x801F000F"
	case ERROR_FLT_DO_NOT_DETACH:
		return "0x801F0010"
	case ERROR_FLT_INSTANCE_ALTITUDE_COLLISION:
		return "0x801F0011"
	case ERROR_FLT_INSTANCE_NAME_COLLISION:
		return "0x801F0012"
	case ERROR_FLT_FILTER_NOT_FOUND:
		return "0x801F0013"
	case ERROR_FLT_VOLUME_NOT_FOUND:
		return "0x801F0014"
	case ERROR_FLT_INSTANCE_NOT_FOUND:
		return "0x801F0015"
	case ERROR_FLT_CONTEXT_ALLOCATION_NOT_FOUND:
		return "0x801F0016"
	case ERROR_FLT_INVALID_CONTEXT_REGISTRATION:
		return "0x801F0017"
	case ERROR_FLT_NAME_CACHE_MISS:
		return "0x801F0018"
	case ERROR_FLT_NO_DEVICE_OBJECT:
		return "0x801F0019"
	case ERROR_FLT_VOLUME_ALREADY_MOUNTED:
		return "0x801F001A"
	case ERROR_FLT_ALREADY_ENLISTED:
		return "0x801F001B"
	case ERROR_FLT_CONTEXT_ALREADY_LINKED:
		return "0x801F001C"
	case ERROR_FLT_NO_WAITER_FOR_REPLY:
		return "0x801F0020"
	case ERROR_HUNG_DISPLAY_DRIVER_THREAD:
		return "0x80260001"
	case DWM_E_COMPOSITIONDISABLED:
		return "0x80263001"
	case DWM_E_REMOTING_NOT_SUPPORTED:
		return "0x80263002"
	case DWM_E_NO_REDIRECTION_SURFACE_AVAILABLE:
		return "0x80263003"
	case DWM_E_NOT_QUEUING_PRESENTS:
		return "0x80263004"
	case DWM_E_ADAPTER_NOT_FOUND:
		return "0x80263005"
	case DWM_S_GDI_REDIRECTION_SURFACE:
		return "0x00263005"
	case ERROR_MONITOR_NO_DESCRIPTOR:
		return "0x00261001"
	case ERROR_MONITOR_UNKNOWN_DESCRIPTOR_FORMAT:
		return "0x00261002"
	case ERROR_MONITOR_INVALID_DESCRIPTOR_CHECKSUM:
		return "0xC0261003"
	case ERROR_MONITOR_INVALID_STANDARD_TIMING_BLOCK:
		return "0xC0261004"
	case ERROR_MONITOR_WMI_DATABLOCK_REGISTRATION_FAILED:
		return "0xC0261005"
	case ERROR_MONITOR_INVALID_SERIAL_NUMBER_MONDSC_BLOCK:
		return "0xC0261006"
	case ERROR_MONITOR_INVALID_USER_FRIENDLY_MONDSC_BLOCK:
		return "0xC0261007"
	case ERROR_MONITOR_NO_MORE_DESCRIPTOR_DATA:
		return "0xC0261008"
	case ERROR_MONITOR_INVALID_DETAILED_TIMING_BLOCK:
		return "0xC0261009"
	case ERROR_MONITOR_INVALID_MANUFACTURE_DATE:
		return "0xC026100A"
	case ERROR_GRAPHICS_NOT_EXCLUSIVE_MODE_OWNER:
		return "0xC0262000"
	case ERROR_GRAPHICS_INSUFFICIENT_DMA_BUFFER:
		return "0xC0262001"
	case ERROR_GRAPHICS_INVALID_DISPLAY_ADAPTER:
		return "0xC0262002"
	case ERROR_GRAPHICS_ADAPTER_WAS_RESET:
		return "0xC0262003"
	case ERROR_GRAPHICS_INVALID_DRIVER_MODEL:
		return "0xC0262004"
	case ERROR_GRAPHICS_PRESENT_MODE_CHANGED:
		return "0xC0262005"
	case ERROR_GRAPHICS_PRESENT_OCCLUDED:
		return "0xC0262006"
	case ERROR_GRAPHICS_PRESENT_DENIED:
		return "0xC0262007"
	case ERROR_GRAPHICS_CANNOTCOLORCONVERT:
		return "0xC0262008"
	case ERROR_GRAPHICS_DRIVER_MISMATCH:
		return "0xC0262009"
	case ERROR_GRAPHICS_PARTIAL_DATA_POPULATED:
		return "0x4026200A"
	case ERROR_GRAPHICS_PRESENT_REDIRECTION_DISABLED:
		return "0xC026200B"
	case ERROR_GRAPHICS_PRESENT_UNOCCLUDED:
		return "0xC026200C"
	case ERROR_GRAPHICS_NO_VIDEO_MEMORY:
		return "0xC0262100"
	case ERROR_GRAPHICS_CANT_LOCK_MEMORY:
		return "0xC0262101"
	case ERROR_GRAPHICS_ALLOCATION_BUSY:
		return "0xC0262102"
	case ERROR_GRAPHICS_TOO_MANY_REFERENCES:
		return "0xC0262103"
	case ERROR_GRAPHICS_TRY_AGAIN_LATER:
		return "0xC0262104"
	case ERROR_GRAPHICS_TRY_AGAIN_NOW:
		return "0xC0262105"
	case ERROR_GRAPHICS_ALLOCATION_INVALID:
		return "0xC0262106"
	case ERROR_GRAPHICS_UNSWIZZLING_APERTURE_UNAVAILABLE:
		return "0xC0262107"
	case ERROR_GRAPHICS_UNSWIZZLING_APERTURE_UNSUPPORTED:
		return "0xC0262108"
	case ERROR_GRAPHICS_CANT_EVICT_PINNED_ALLOCATION:
		return "0xC0262109"
	case ERROR_GRAPHICS_INVALID_ALLOCATION_USAGE:
		return "0xC0262110"
	case ERROR_GRAPHICS_CANT_RENDER_LOCKED_ALLOCATION:
		return "0xC0262111"
	case ERROR_GRAPHICS_ALLOCATION_CLOSED:
		return "0xC0262112"
	case ERROR_GRAPHICS_INVALID_ALLOCATION_INSTANCE:
		return "0xC0262113"
	case ERROR_GRAPHICS_INVALID_ALLOCATION_HANDLE:
		return "0xC0262114"
	case ERROR_GRAPHICS_WRONG_ALLOCATION_DEVICE:
		return "0xC0262115"
	case ERROR_GRAPHICS_ALLOCATION_CONTENT_LOST:
		return "0xC0262116"
	case ERROR_GRAPHICS_GPU_EXCEPTION_ON_DEVICE:
		return "0xC0262200"
	case ERROR_GRAPHICS_INVALID_VIDPN_TOPOLOGY:
		return "0xC0262300"
	case ERROR_GRAPHICS_VIDPN_TOPOLOGY_NOT_SUPPORTED:
		return "0xC0262301"
	case ERROR_GRAPHICS_VIDPN_TOPOLOGY_CURRENTLY_NOT_SUPPORTED:
		return "0xC0262302"
	case ERROR_GRAPHICS_INVALID_VIDPN:
		return "0xC0262303"
	case ERROR_GRAPHICS_INVALID_VIDEO_PRESENT_SOURCE:
		return "0xC0262304"
	case ERROR_GRAPHICS_INVALID_VIDEO_PRESENT_TARGET:
		return "0xC0262305"
	case ERROR_GRAPHICS_VIDPN_MODALITY_NOT_SUPPORTED:
		return "0xC0262306"
	case ERROR_GRAPHICS_MODE_NOT_PINNED:
		return "0x00262307"
	case ERROR_GRAPHICS_INVALID_VIDPN_SOURCEMODESET:
		return "0xC0262308"
	case ERROR_GRAPHICS_INVALID_VIDPN_TARGETMODESET:
		return "0xC0262309"
	case ERROR_GRAPHICS_INVALID_FREQUENCY:
		return "0xC026230A"
	case ERROR_GRAPHICS_INVALID_ACTIVE_REGION:
		return "0xC026230B"
	case ERROR_GRAPHICS_INVALID_TOTAL_REGION:
		return "0xC026230C"
	case ERROR_GRAPHICS_INVALID_VIDEO_PRESENT_SOURCE_MODE:
		return "0xC0262310"
	case ERROR_GRAPHICS_INVALID_VIDEO_PRESENT_TARGET_MODE:
		return "0xC0262311"
	case ERROR_GRAPHICS_PINNED_MODE_MUST_REMAIN_IN_SET:
		return "0xC0262312"
	case ERROR_GRAPHICS_PATH_ALREADY_IN_TOPOLOGY:
		return "0xC0262313"
	case ERROR_GRAPHICS_MODE_ALREADY_IN_MODESET:
		return "0xC0262314"
	case ERROR_GRAPHICS_INVALID_VIDEOPRESENTSOURCESET:
		return "0xC0262315"
	case ERROR_GRAPHICS_INVALID_VIDEOPRESENTTARGETSET:
		return "0xC0262316"
	case ERROR_GRAPHICS_SOURCE_ALREADY_IN_SET:
		return "0xC0262317"
	case ERROR_GRAPHICS_TARGET_ALREADY_IN_SET:
		return "0xC0262318"
	case ERROR_GRAPHICS_INVALID_VIDPN_PRESENT_PATH:
		return "0xC0262319"
	case ERROR_GRAPHICS_NO_RECOMMENDED_VIDPN_TOPOLOGY:
		return "0xC026231A"
	case ERROR_GRAPHICS_INVALID_MONITOR_FREQUENCYRANGESET:
		return "0xC026231B"
	case ERROR_GRAPHICS_INVALID_MONITOR_FREQUENCYRANGE:
		return "0xC026231C"
	case ERROR_GRAPHICS_FREQUENCYRANGE_NOT_IN_SET:
		return "0xC026231D"
	case ERROR_GRAPHICS_NO_PREFERRED_MODE:
		return "0x0026231E"
	case ERROR_GRAPHICS_FREQUENCYRANGE_ALREADY_IN_SET:
		return "0xC026231F"
	case ERROR_GRAPHICS_STALE_MODESET:
		return "0xC0262320"
	case ERROR_GRAPHICS_INVALID_MONITOR_SOURCEMODESET:
		return "0xC0262321"
	case ERROR_GRAPHICS_INVALID_MONITOR_SOURCE_MODE:
		return "0xC0262322"
	case ERROR_GRAPHICS_NO_RECOMMENDED_FUNCTIONAL_VIDPN:
		return "0xC0262323"
	case ERROR_GRAPHICS_MODE_ID_MUST_BE_UNIQUE:
		return "0xC0262324"
	case ERROR_GRAPHICS_EMPTY_ADAPTER_MONITOR_MODE_SUPPORT_INTERSECTION:
		return "0xC0262325"
	case ERROR_GRAPHICS_VIDEO_PRESENT_TARGETS_LESS_THAN_SOURCES:
		return "0xC0262326"
	case ERROR_GRAPHICS_PATH_NOT_IN_TOPOLOGY:
		return "0xC0262327"
	case ERROR_GRAPHICS_ADAPTER_MUST_HAVE_AT_LEAST_ONE_SOURCE:
		return "0xC0262328"
	case ERROR_GRAPHICS_ADAPTER_MUST_HAVE_AT_LEAST_ONE_TARGET:
		return "0xC0262329"
	case ERROR_GRAPHICS_INVALID_MONITORDESCRIPTORSET:
		return "0xC026232A"
	case ERROR_GRAPHICS_INVALID_MONITORDESCRIPTOR:
		return "0xC026232B"
	case ERROR_GRAPHICS_MONITORDESCRIPTOR_NOT_IN_SET:
		return "0xC026232C"
	case ERROR_GRAPHICS_MONITORDESCRIPTOR_ALREADY_IN_SET:
		return "0xC026232D"
	case ERROR_GRAPHICS_MONITORDESCRIPTOR_ID_MUST_BE_UNIQUE:
		return "0xC026232E"
	case ERROR_GRAPHICS_INVALID_VIDPN_TARGET_SUBSET_TYPE:
		return "0xC026232F"
	case ERROR_GRAPHICS_RESOURCES_NOT_RELATED:
		return "0xC0262330"
	case ERROR_GRAPHICS_SOURCE_ID_MUST_BE_UNIQUE:
		return "0xC0262331"
	case ERROR_GRAPHICS_TARGET_ID_MUST_BE_UNIQUE:
		return "0xC0262332"
	case ERROR_GRAPHICS_NO_AVAILABLE_VIDPN_TARGET:
		return "0xC0262333"
	case ERROR_GRAPHICS_MONITOR_COULD_NOT_BE_ASSOCIATED_WITH_ADAPTER:
		return "0xC0262334"
	case ERROR_GRAPHICS_NO_VIDPNMGR:
		return "0xC0262335"
	case ERROR_GRAPHICS_NO_ACTIVE_VIDPN:
		return "0xC0262336"
	case ERROR_GRAPHICS_STALE_VIDPN_TOPOLOGY:
		return "0xC0262337"
	case ERROR_GRAPHICS_MONITOR_NOT_CONNECTED:
		return "0xC0262338"
	case ERROR_GRAPHICS_SOURCE_NOT_IN_TOPOLOGY:
		return "0xC0262339"
	case ERROR_GRAPHICS_INVALID_PRIMARYSURFACE_SIZE:
		return "0xC026233A"
	case ERROR_GRAPHICS_INVALID_VISIBLEREGION_SIZE:
		return "0xC026233B"
	case ERROR_GRAPHICS_INVALID_STRIDE:
		return "0xC026233C"
	case ERROR_GRAPHICS_INVALID_PIXELFORMAT:
		return "0xC026233D"
	case ERROR_GRAPHICS_INVALID_COLORBASIS:
		return "0xC026233E"
	case ERROR_GRAPHICS_INVALID_PIXELVALUEACCESSMODE:
		return "0xC026233F"
	case ERROR_GRAPHICS_TARGET_NOT_IN_TOPOLOGY:
		return "0xC0262340"
	case ERROR_GRAPHICS_NO_DISPLAY_MODE_MANAGEMENT_SUPPORT:
		return "0xC0262341"
	case ERROR_GRAPHICS_VIDPN_SOURCE_IN_USE:
		return "0xC0262342"
	case ERROR_GRAPHICS_CANT_ACCESS_ACTIVE_VIDPN:
		return "0xC0262343"
	case ERROR_GRAPHICS_INVALID_PATH_IMPORTANCE_ORDINAL:
		return "0xC0262344"
	case ERROR_GRAPHICS_INVALID_PATH_CONTENT_GEOMETRY_TRANSFORMATION:
		return "0xC0262345"
	case ERROR_GRAPHICS_PATH_CONTENT_GEOMETRY_TRANSFORMATION_NOT_SUPPORTED:
		return "0xC0262346"
	case ERROR_GRAPHICS_INVALID_GAMMA_RAMP:
		return "0xC0262347"
	case ERROR_GRAPHICS_GAMMA_RAMP_NOT_SUPPORTED:
		return "0xC0262348"
	case ERROR_GRAPHICS_MULTISAMPLING_NOT_SUPPORTED:
		return "0xC0262349"
	case ERROR_GRAPHICS_MODE_NOT_IN_MODESET:
		return "0xC026234A"
	case ERROR_GRAPHICS_DATASET_IS_EMPTY:
		return "0x0026234B"
	case ERROR_GRAPHICS_NO_MORE_ELEMENTS_IN_DATASET:
		return "0x0026234C"
	case ERROR_GRAPHICS_INVALID_VIDPN_TOPOLOGY_RECOMMENDATION_REASON:
		return "0xC026234D"
	case ERROR_GRAPHICS_INVALID_PATH_CONTENT_TYPE:
		return "0xC026234E"
	case ERROR_GRAPHICS_INVALID_COPYPROTECTION_TYPE:
		return "0xC026234F"
	case ERROR_GRAPHICS_UNASSIGNED_MODESET_ALREADY_EXISTS:
		return "0xC0262350"
	case ERROR_GRAPHICS_PATH_CONTENT_GEOMETRY_TRANSFORMATION_NOT_PINNED:
		return "0x00262351"
	case ERROR_GRAPHICS_INVALID_SCANLINE_ORDERING:
		return "0xC0262352"
	case ERROR_GRAPHICS_TOPOLOGY_CHANGES_NOT_ALLOWED:
		return "0xC0262353"
	case ERROR_GRAPHICS_NO_AVAILABLE_IMPORTANCE_ORDINALS:
		return "0xC0262354"
	case ERROR_GRAPHICS_INCOMPATIBLE_PRIVATE_FORMAT:
		return "0xC0262355"
	case ERROR_GRAPHICS_INVALID_MODE_PRUNING_ALGORITHM:
		return "0xC0262356"
	case ERROR_GRAPHICS_INVALID_MONITOR_CAPABILITY_ORIGIN:
		return "0xC0262357"
	case ERROR_GRAPHICS_INVALID_MONITOR_FREQUENCYRANGE_CONSTRAINT:
		return "0xC0262358"
	case ERROR_GRAPHICS_MAX_NUM_PATHS_REACHED:
		return "0xC0262359"
	case ERROR_GRAPHICS_CANCEL_VIDPN_TOPOLOGY_AUGMENTATION:
		return "0xC026235A"
	case ERROR_GRAPHICS_INVALID_CLIENT_TYPE:
		return "0xC026235B"
	case ERROR_GRAPHICS_CLIENTVIDPN_NOT_SET:
		return "0xC026235C"
	case ERROR_GRAPHICS_SPECIFIED_CHILD_ALREADY_CONNECTED:
		return "0xC0262400"
	case ERROR_GRAPHICS_CHILD_DESCRIPTOR_NOT_SUPPORTED:
		return "0xC0262401"
	case ERROR_GRAPHICS_UNKNOWN_CHILD_STATUS:
		return "0x4026242F"
	case ERROR_GRAPHICS_NOT_A_LINKED_ADAPTER:
		return "0xC0262430"
	case ERROR_GRAPHICS_LEADLINK_NOT_ENUMERATED:
		return "0xC0262431"
	case ERROR_GRAPHICS_CHAINLINKS_NOT_ENUMERATED:
		return "0xC0262432"
	case ERROR_GRAPHICS_ADAPTER_CHAIN_NOT_READY:
		return "0xC0262433"
	case ERROR_GRAPHICS_CHAINLINKS_NOT_STARTED:
		return "0xC0262434"
	case ERROR_GRAPHICS_CHAINLINKS_NOT_POWERED_ON:
		return "0xC0262435"
	case ERROR_GRAPHICS_INCONSISTENT_DEVICE_LINK_STATE:
		return "0xC0262436"
	case ERROR_GRAPHICS_LEADLINK_START_DEFERRED:
		return "0x40262437"
	case ERROR_GRAPHICS_NOT_POST_DEVICE_DRIVER:
		return "0xC0262438"
	case ERROR_GRAPHICS_POLLING_TOO_FREQUENTLY:
		return "0x40262439"
	case ERROR_GRAPHICS_START_DEFERRED:
		return "0x4026243A"
	case ERROR_GRAPHICS_ADAPTER_ACCESS_NOT_EXCLUDED:
		return "0xC026243B"
	case ERROR_GRAPHICS_OPM_NOT_SUPPORTED:
		return "0xC0262500"
	case ERROR_GRAPHICS_COPP_NOT_SUPPORTED:
		return "0xC0262501"
	case ERROR_GRAPHICS_UAB_NOT_SUPPORTED:
		return "0xC0262502"
	case ERROR_GRAPHICS_OPM_INVALID_ENCRYPTED_PARAMETERS:
		return "0xC0262503"
	case ERROR_GRAPHICS_OPM_NO_VIDEO_OUTPUTS_EXIST:
		return "0xC0262505"
	case ERROR_GRAPHICS_OPM_INTERNAL_ERROR:
		return "0xC026250B"
	case ERROR_GRAPHICS_OPM_INVALID_HANDLE:
		return "0xC026250C"
	case ERROR_GRAPHICS_PVP_INVALID_CERTIFICATE_LENGTH:
		return "0xC026250E"
	case ERROR_GRAPHICS_OPM_SPANNING_MODE_ENABLED:
		return "0xC026250F"
	case ERROR_GRAPHICS_OPM_THEATER_MODE_ENABLED:
		return "0xC0262510"
	case ERROR_GRAPHICS_PVP_HFS_FAILED:
		return "0xC0262511"
	case ERROR_GRAPHICS_OPM_INVALID_SRM:
		return "0xC0262512"
	case ERROR_GRAPHICS_OPM_OUTPUT_DOES_NOT_SUPPORT_HDCP:
		return "0xC0262513"
	case ERROR_GRAPHICS_OPM_OUTPUT_DOES_NOT_SUPPORT_ACP:
		return "0xC0262514"
	case ERROR_GRAPHICS_OPM_OUTPUT_DOES_NOT_SUPPORT_CGMSA:
		return "0xC0262515"
	case ERROR_GRAPHICS_OPM_HDCP_SRM_NEVER_SET:
		return "0xC0262516"
	case ERROR_GRAPHICS_OPM_RESOLUTION_TOO_HIGH:
		return "0xC0262517"
	case ERROR_GRAPHICS_OPM_ALL_HDCP_HARDWARE_ALREADY_IN_USE:
		return "0xC0262518"
	case ERROR_GRAPHICS_OPM_VIDEO_OUTPUT_NO_LONGER_EXISTS:
		return "0xC026251A"
	case ERROR_GRAPHICS_OPM_SESSION_TYPE_CHANGE_IN_PROGRESS:
		return "0xC026251B"
	case ERROR_GRAPHICS_OPM_VIDEO_OUTPUT_DOES_NOT_HAVE_COPP_SEMANTICS:
		return "0xC026251C"
	case ERROR_GRAPHICS_OPM_INVALID_INFORMATION_REQUEST:
		return "0xC026251D"
	case ERROR_GRAPHICS_OPM_DRIVER_INTERNAL_ERROR:
		return "0xC026251E"
	case ERROR_GRAPHICS_OPM_VIDEO_OUTPUT_DOES_NOT_HAVE_OPM_SEMANTICS:
		return "0xC026251F"
	case ERROR_GRAPHICS_OPM_SIGNALING_NOT_SUPPORTED:
		return "0xC0262520"
	case ERROR_GRAPHICS_OPM_INVALID_CONFIGURATION_REQUEST:
		return "0xC0262521"
	case ERROR_GRAPHICS_I2C_NOT_SUPPORTED:
		return "0xC0262580"
	case ERROR_GRAPHICS_I2C_DEVICE_DOES_NOT_EXIST:
		return "0xC0262581"
	case ERROR_GRAPHICS_I2C_ERROR_TRANSMITTING_DATA:
		return "0xC0262582"
	case ERROR_GRAPHICS_I2C_ERROR_RECEIVING_DATA:
		return "0xC0262583"
	case ERROR_GRAPHICS_DDCCI_VCP_NOT_SUPPORTED:
		return "0xC0262584"
	case ERROR_GRAPHICS_DDCCI_INVALID_DATA:
		return "0xC0262585"
	case ERROR_GRAPHICS_DDCCI_MONITOR_RETURNED_INVALID_TIMING_STATUS_BYTE:
		return "0xC0262586"
	case ERROR_GRAPHICS_MCA_INVALID_CAPABILITIES_STRING:
		return "0xC0262587"
	case ERROR_GRAPHICS_MCA_INTERNAL_ERROR:
		return "0xC0262588"
	case ERROR_GRAPHICS_DDCCI_INVALID_MESSAGE_COMMAND:
		return "0xC0262589"
	case ERROR_GRAPHICS_DDCCI_INVALID_MESSAGE_LENGTH:
		return "0xC026258A"
	case ERROR_GRAPHICS_DDCCI_INVALID_MESSAGE_CHECKSUM:
		return "0xC026258B"
	case ERROR_GRAPHICS_INVALID_PHYSICAL_MONITOR_HANDLE:
		return "0xC026258C"
	case ERROR_GRAPHICS_MONITOR_NO_LONGER_EXISTS:
		return "0xC026258D"
	case ERROR_GRAPHICS_DDCCI_CURRENT_CURRENT_VALUE_GREATER_THAN_MAXIMUM_VALUE:
		return "0xC02625D8"
	case ERROR_GRAPHICS_MCA_INVALID_VCP_VERSION:
		return "0xC02625D9"
	case ERROR_GRAPHICS_MCA_MONITOR_VIOLATES_MCCS_SPECIFICATION:
		return "0xC02625DA"
	case ERROR_GRAPHICS_MCA_MCCS_VERSION_MISMATCH:
		return "0xC02625DB"
	case ERROR_GRAPHICS_MCA_UNSUPPORTED_MCCS_VERSION:
		return "0xC02625DC"
	case ERROR_GRAPHICS_MCA_INVALID_TECHNOLOGY_TYPE_RETURNED:
		return "0xC02625DE"
	case ERROR_GRAPHICS_MCA_UNSUPPORTED_COLOR_TEMPERATURE:
		return "0xC02625DF"
	case ERROR_GRAPHICS_ONLY_CONSOLE_SESSION_SUPPORTED:
		return "0xC02625E0"
	case ERROR_GRAPHICS_NO_DISPLAY_DEVICE_CORRESPONDS_TO_NAME:
		return "0xC02625E1"
	case ERROR_GRAPHICS_DISPLAY_DEVICE_NOT_ATTACHED_TO_DESKTOP:
		return "0xC02625E2"
	case ERROR_GRAPHICS_MIRRORING_DEVICES_NOT_SUPPORTED:
		return "0xC02625E3"
	case ERROR_GRAPHICS_INVALID_POINTER:
		return "0xC02625E4"
	case ERROR_GRAPHICS_NO_MONITORS_CORRESPOND_TO_DISPLAY_DEVICE:
		return "0xC02625E5"
	case ERROR_GRAPHICS_PARAMETER_ARRAY_TOO_SMALL:
		return "0xC02625E6"
	case ERROR_GRAPHICS_INTERNAL_ERROR:
		return "0xC02625E7"
	case ERROR_GRAPHICS_SESSION_TYPE_CHANGE_IN_PROGRESS:
		return "0xC02605E8"
	case TPM_E_ERROR_MASK:
		return "0x80280000"
	case TPM_E_AUTHFAIL:
		return "0x80280001"
	case TPM_E_BADINDEX:
		return "0x80280002"
	case TPM_E_BAD_PARAMETER:
		return "0x80280003"
	case TPM_E_AUDITFAILURE:
		return "0x80280004"
	case TPM_E_CLEAR_DISABLED:
		return "0x80280005"
	case TPM_E_DEACTIVATED:
		return "0x80280006"
	case TPM_E_DISABLED:
		return "0x80280007"
	case TPM_E_DISABLED_CMD:
		return "0x80280008"
	case TPM_E_FAIL:
		return "0x80280009"
	case TPM_E_BAD_ORDINAL:
		return "0x8028000A"
	case TPM_E_INSTALL_DISABLED:
		return "0x8028000B"
	case TPM_E_INVALID_KEYHANDLE:
		return "0x8028000C"
	case TPM_E_KEYNOTFOUND:
		return "0x8028000D"
	case TPM_E_INAPPROPRIATE_ENC:
		return "0x8028000E"
	case TPM_E_MIGRATEFAIL:
		return "0x8028000F"
	case TPM_E_INVALID_PCR_INFO:
		return "0x80280010"
	case TPM_E_NOSPACE:
		return "0x80280011"
	case TPM_E_NOSRK:
		return "0x80280012"
	case TPM_E_NOTSEALED_BLOB:
		return "0x80280013"
	case TPM_E_OWNER_SET:
		return "0x80280014"
	case TPM_E_RESOURCES:
		return "0x80280015"
	case TPM_E_SHORTRANDOM:
		return "0x80280016"
	case TPM_E_SIZE:
		return "0x80280017"
	case TPM_E_WRONGPCRVAL:
		return "0x80280018"
	case TPM_E_BAD_PARAM_SIZE:
		return "0x80280019"
	case TPM_E_SHA_THREAD:
		return "0x8028001A"
	case TPM_E_SHA_ERROR:
		return "0x8028001B"
	case TPM_E_FAILEDSELFTEST:
		return "0x8028001C"
	case TPM_E_AUTH2FAIL:
		return "0x8028001D"
	case TPM_E_BADTAG:
		return "0x8028001E"
	case TPM_E_IOERROR:
		return "0x8028001F"
	case TPM_E_ENCRYPT_ERROR:
		return "0x80280020"
	case TPM_E_DECRYPT_ERROR:
		return "0x80280021"
	case TPM_E_INVALID_AUTHHANDLE:
		return "0x80280022"
	case TPM_E_NO_ENDORSEMENT:
		return "0x80280023"
	case TPM_E_INVALID_KEYUSAGE:
		return "0x80280024"
	case TPM_E_WRONG_ENTITYTYPE:
		return "0x80280025"
	case TPM_E_INVALID_POSTINIT:
		return "0x80280026"
	case TPM_E_INAPPROPRIATE_SIG:
		return "0x80280027"
	case TPM_E_BAD_KEY_PROPERTY:
		return "0x80280028"
	case TPM_E_BAD_MIGRATION:
		return "0x80280029"
	case TPM_E_BAD_SCHEME:
		return "0x8028002A"
	case TPM_E_BAD_DATASIZE:
		return "0x8028002B"
	case TPM_E_BAD_MODE:
		return "0x8028002C"
	case TPM_E_BAD_PRESENCE:
		return "0x8028002D"
	case TPM_E_BAD_VERSION:
		return "0x8028002E"
	case TPM_E_NO_WRAP_TRANSPORT:
		return "0x8028002F"
	case TPM_E_AUDITFAIL_UNSUCCESSFUL:
		return "0x80280030"
	case TPM_E_AUDITFAIL_SUCCESSFUL:
		return "0x80280031"
	case TPM_E_NOTRESETABLE:
		return "0x80280032"
	case TPM_E_NOTLOCAL:
		return "0x80280033"
	case TPM_E_BAD_TYPE:
		return "0x80280034"
	case TPM_E_INVALID_RESOURCE:
		return "0x80280035"
	case TPM_E_NOTFIPS:
		return "0x80280036"
	case TPM_E_INVALID_FAMILY:
		return "0x80280037"
	case TPM_E_NO_NV_PERMISSION:
		return "0x80280038"
	case TPM_E_REQUIRES_SIGN:
		return "0x80280039"
	case TPM_E_KEY_NOTSUPPORTED:
		return "0x8028003A"
	case TPM_E_AUTH_CONFLICT:
		return "0x8028003B"
	case TPM_E_AREA_LOCKED:
		return "0x8028003C"
	case TPM_E_BAD_LOCALITY:
		return "0x8028003D"
	case TPM_E_READ_ONLY:
		return "0x8028003E"
	case TPM_E_PER_NOWRITE:
		return "0x8028003F"
	case TPM_E_FAMILYCOUNT:
		return "0x80280040"
	case TPM_E_WRITE_LOCKED:
		return "0x80280041"
	case TPM_E_BAD_ATTRIBUTES:
		return "0x80280042"
	case TPM_E_INVALID_STRUCTURE:
		return "0x80280043"
	case TPM_E_KEY_OWNER_CONTROL:
		return "0x80280044"
	case TPM_E_BAD_COUNTER:
		return "0x80280045"
	case TPM_E_NOT_FULLWRITE:
		return "0x80280046"
	case TPM_E_CONTEXT_GAP:
		return "0x80280047"
	case TPM_E_MAXNVWRITES:
		return "0x80280048"
	case TPM_E_NOOPERATOR:
		return "0x80280049"
	case TPM_E_RESOURCEMISSING:
		return "0x8028004A"
	case TPM_E_DELEGATE_LOCK:
		return "0x8028004B"
	case TPM_E_DELEGATE_FAMILY:
		return "0x8028004C"
	case TPM_E_DELEGATE_ADMIN:
		return "0x8028004D"
	case TPM_E_TRANSPORT_NOTEXCLUSIVE:
		return "0x8028004E"
	case TPM_E_OWNER_CONTROL:
		return "0x8028004F"
	case TPM_E_DAA_RESOURCES:
		return "0x80280050"
	case TPM_E_DAA_INPUT_DATA0:
		return "0x80280051"
	case TPM_E_DAA_INPUT_DATA1:
		return "0x80280052"
	case TPM_E_DAA_ISSUER_SETTINGS:
		return "0x80280053"
	case TPM_E_DAA_TPM_SETTINGS:
		return "0x80280054"
	case TPM_E_DAA_STAGE:
		return "0x80280055"
	case TPM_E_DAA_ISSUER_VALIDITY:
		return "0x80280056"
	case TPM_E_DAA_WRONG_W:
		return "0x80280057"
	case TPM_E_BAD_HANDLE:
		return "0x80280058"
	case TPM_E_BAD_DELEGATE:
		return "0x80280059"
	case TPM_E_BADCONTEXT:
		return "0x8028005A"
	case TPM_E_TOOMANYCONTEXTS:
		return "0x8028005B"
	case TPM_E_MA_TICKET_SIGNATURE:
		return "0x8028005C"
	case TPM_E_MA_DESTINATION:
		return "0x8028005D"
	case TPM_E_MA_SOURCE:
		return "0x8028005E"
	case TPM_E_MA_AUTHORITY:
		return "0x8028005F"
	case TPM_E_PERMANENTEK:
		return "0x80280061"
	case TPM_E_BAD_SIGNATURE:
		return "0x80280062"
	case TPM_E_NOCONTEXTSPACE:
		return "0x80280063"
	case TPM_E_COMMAND_BLOCKED:
		return "0x80280400"
	case TPM_E_INVALID_HANDLE:
		return "0x80280401"
	case TPM_E_DUPLICATE_VHANDLE:
		return "0x80280402"
	case TPM_E_EMBEDDED_COMMAND_BLOCKED:
		return "0x80280403"
	case TPM_E_EMBEDDED_COMMAND_UNSUPPORTED:
		return "0x80280404"
	case TPM_E_RETRY:
		return "0x80280800"
	case TPM_E_NEEDS_SELFTEST:
		return "0x80280801"
	case TPM_E_DOING_SELFTEST:
		return "0x80280802"
	case TPM_E_DEFEND_LOCK_RUNNING:
		return "0x80280803"
	case TBS_E_INTERNAL_ERROR:
		return "0x80284001"
	case TBS_E_BAD_PARAMETER:
		return "0x80284002"
	case TBS_E_INVALID_OUTPUT_POINTER:
		return "0x80284003"
	case TBS_E_INVALID_CONTEXT:
		return "0x80284004"
	case TBS_E_INSUFFICIENT_BUFFER:
		return "0x80284005"
	case TBS_E_IOERROR:
		return "0x80284006"
	case TBS_E_INVALID_CONTEXT_PARAM:
		return "0x80284007"
	case TBS_E_SERVICE_NOT_RUNNING:
		return "0x80284008"
	case TBS_E_TOO_MANY_TBS_CONTEXTS:
		return "0x80284009"
	case TBS_E_TOO_MANY_RESOURCES:
		return "0x8028400A"
	case TBS_E_SERVICE_START_PENDING:
		return "0x8028400B"
	case TBS_E_PPI_NOT_SUPPORTED:
		return "0x8028400C"
	case TBS_E_COMMAND_CANCELED:
		return "0x8028400D"
	case TBS_E_BUFFER_TOO_LARGE:
		return "0x8028400E"
	case TBS_E_TPM_NOT_FOUND:
		return "0x8028400F"
	case TBS_E_SERVICE_DISABLED:
		return "0x80284010"
	case TBS_E_NO_EVENT_LOG:
		return "0x80284011"
	case TPMAPI_E_INVALID_STATE:
		return "0x80290100"
	case TPMAPI_E_NOT_ENOUGH_DATA:
		return "0x80290101"
	case TPMAPI_E_TOO_MUCH_DATA:
		return "0x80290102"
	case TPMAPI_E_INVALID_OUTPUT_POINTER:
		return "0x80290103"
	case TPMAPI_E_INVALID_PARAMETER:
		return "0x80290104"
	case TPMAPI_E_OUT_OF_MEMORY:
		return "0x80290105"
	case TPMAPI_E_BUFFER_TOO_SMALL:
		return "0x80290106"
	case TPMAPI_E_INTERNAL_ERROR:
		return "0x80290107"
	case TPMAPI_E_ACCESS_DENIED:
		return "0x80290108"
	case TPMAPI_E_AUTHORIZATION_FAILED:
		return "0x80290109"
	case TPMAPI_E_INVALID_CONTEXT_HANDLE:
		return "0x8029010A"
	case TPMAPI_E_TBS_COMMUNICATION_ERROR:
		return "0x8029010B"
	case TPMAPI_E_TPM_COMMAND_ERROR:
		return "0x8029010C"
	case TPMAPI_E_MESSAGE_TOO_LARGE:
		return "0x8029010D"
	case TPMAPI_E_INVALID_ENCODING:
		return "0x8029010E"
	case TPMAPI_E_INVALID_KEY_SIZE:
		return "0x8029010F"
	case TPMAPI_E_ENCRYPTION_FAILED:
		return "0x80290110"
	case TPMAPI_E_INVALID_KEY_PARAMS:
		return "0x80290111"
	case TPMAPI_E_INVALID_MIGRATION_AUTHORIZATION_BLOB:
		return "0x80290112"
	case TPMAPI_E_INVALID_PCR_INDEX:
		return "0x80290113"
	case TPMAPI_E_INVALID_DELEGATE_BLOB:
		return "0x80290114"
	case TPMAPI_E_INVALID_CONTEXT_PARAMS:
		return "0x80290115"
	case TPMAPI_E_INVALID_KEY_BLOB:
		return "0x80290116"
	case TPMAPI_E_INVALID_PCR_DATA:
		return "0x80290117"
	case TPMAPI_E_INVALID_OWNER_AUTH:
		return "0x80290118"
	case TPMAPI_E_FIPS_RNG_CHECK_FAILED:
		return "0x80290119"
	case TPMAPI_E_EMPTY_TCG_LOG:
		return "0x8029011A"
	case TPMAPI_E_INVALID_TCG_LOG_ENTRY:
		return "0x8029011B"
	case TPMAPI_E_TCG_SEPARATOR_ABSENT:
		return "0x8029011C"
	case TPMAPI_E_TCG_INVALID_DIGEST_ENTRY:
		return "0x8029011D"
	case TBSIMP_E_BUFFER_TOO_SMALL:
		return "0x80290200"
	case TBSIMP_E_CLEANUP_FAILED:
		return "0x80290201"
	case TBSIMP_E_INVALID_CONTEXT_HANDLE:
		return "0x80290202"
	case TBSIMP_E_INVALID_CONTEXT_PARAM:
		return "0x80290203"
	case TBSIMP_E_TPM_ERROR:
		return "0x80290204"
	case TBSIMP_E_HASH_BAD_KEY:
		return "0x80290205"
	case TBSIMP_E_DUPLICATE_VHANDLE:
		return "0x80290206"
	case TBSIMP_E_INVALID_OUTPUT_POINTER:
		return "0x80290207"
	case TBSIMP_E_INVALID_PARAMETER:
		return "0x80290208"
	case TBSIMP_E_RPC_INIT_FAILED:
		return "0x80290209"
	case TBSIMP_E_SCHEDULER_NOT_RUNNING:
		return "0x8029020A"
	case TBSIMP_E_COMMAND_CANCELED:
		return "0x8029020B"
	case TBSIMP_E_OUT_OF_MEMORY:
		return "0x8029020C"
	case TBSIMP_E_LIST_NO_MORE_ITEMS:
		return "0x8029020D"
	case TBSIMP_E_LIST_NOT_FOUND:
		return "0x8029020E"
	case TBSIMP_E_NOT_ENOUGH_SPACE:
		return "0x8029020F"
	case TBSIMP_E_NOT_ENOUGH_TPM_CONTEXTS:
		return "0x80290210"
	case TBSIMP_E_COMMAND_FAILED:
		return "0x80290211"
	case TBSIMP_E_UNKNOWN_ORDINAL:
		return "0x80290212"
	case TBSIMP_E_RESOURCE_EXPIRED:
		return "0x80290213"
	case TBSIMP_E_INVALID_RESOURCE:
		return "0x80290214"
	case TBSIMP_E_NOTHING_TO_UNLOAD:
		return "0x80290215"
	case TBSIMP_E_HASH_TABLE_FULL:
		return "0x80290216"
	case TBSIMP_E_TOO_MANY_TBS_CONTEXTS:
		return "0x80290217"
	case TBSIMP_E_TOO_MANY_RESOURCES:
		return "0x80290218"
	case TBSIMP_E_PPI_NOT_SUPPORTED:
		return "0x80290219"
	case TBSIMP_E_TPM_INCOMPATIBLE:
		return "0x8029021A"
	case TBSIMP_E_NO_EVENT_LOG:
		return "0x8029021B"
	case TPM_E_PPI_ACPI_FAILURE:
		return "0x80290300"
	case TPM_E_PPI_USER_ABORT:
		return "0x80290301"
	case TPM_E_PPI_BIOS_FAILURE:
		return "0x80290302"
	case TPM_E_PPI_NOT_SUPPORTED:
		return "0x80290303"
	case PLA_E_DCS_NOT_FOUND:
		return "0x80300002"
	case PLA_E_DCS_IN_USE:
		return "0x803000AA"
	case PLA_E_TOO_MANY_FOLDERS:
		return "0x80300045"
	case PLA_E_NO_MIN_DISK:
		return "0x80300070"
	case PLA_E_DCS_ALREADY_EXISTS:
		return "0x803000B7"
	case PLA_S_PROPERTY_IGNORED:
		return "0x00300100"
	case PLA_E_PROPERTY_CONFLICT:
		return "0x80300101"
	case PLA_E_DCS_SINGLETON_REQUIRED:
		return "0x80300102"
	case PLA_E_CREDENTIALS_REQUIRED:
		return "0x80300103"
	case PLA_E_DCS_NOT_RUNNING:
		return "0x80300104"
	case PLA_E_CONFLICT_INCL_EXCL_API:
		return "0x80300105"
	case PLA_E_NETWORK_EXE_NOT_VALID:
		return "0x80300106"
	case PLA_E_EXE_ALREADY_CONFIGURED:
		return "0x80300107"
	case PLA_E_EXE_PATH_NOT_VALID:
		return "0x80300108"
	case PLA_E_DC_ALREADY_EXISTS:
		return "0x80300109"
	case PLA_E_DCS_START_WAIT_TIMEOUT:
		return "0x8030010A"
	case PLA_E_DC_START_WAIT_TIMEOUT:
		return "0x8030010B"
	case PLA_E_REPORT_WAIT_TIMEOUT:
		return "0x8030010C"
	case PLA_E_NO_DUPLICATES:
		return "0x8030010D"
	case PLA_E_EXE_FULL_PATH_REQUIRED:
		return "0x8030010E"
	case PLA_E_INVALID_SESSION_NAME:
		return "0x8030010F"
	case PLA_E_PLA_CHANNEL_NOT_ENABLED:
		return "0x80300110"
	case PLA_E_TASKSCHED_CHANNEL_NOT_ENABLED:
		return "0x80300111"
	case PLA_E_RULES_MANAGER_FAILED:
		return "0x80300112"
	case PLA_E_CABAPI_FAILURE:
		return "0x80300113"
	case FVE_E_LOCKED_VOLUME:
		return "0x80310000"
	case FVE_E_NOT_ENCRYPTED:
		return "0x80310001"
	case FVE_E_NO_TPM_BIOS:
		return "0x80310002"
	case FVE_E_NO_MBR_METRIC:
		return "0x80310003"
	case FVE_E_NO_BOOTSECTOR_METRIC:
		return "0x80310004"
	case FVE_E_NO_BOOTMGR_METRIC:
		return "0x80310005"
	case FVE_E_WRONG_BOOTMGR:
		return "0x80310006"
	case FVE_E_SECURE_KEY_REQUIRED:
		return "0x80310007"
	case FVE_E_NOT_ACTIVATED:
		return "0x80310008"
	case FVE_E_ACTION_NOT_ALLOWED:
		return "0x80310009"
	case FVE_E_AD_SCHEMA_NOT_INSTALLED:
		return "0x8031000A"
	case FVE_E_AD_INVALID_DATATYPE:
		return "0x8031000B"
	case FVE_E_AD_INVALID_DATASIZE:
		return "0x8031000C"
	case FVE_E_AD_NO_VALUES:
		return "0x8031000D"
	case FVE_E_AD_ATTR_NOT_SET:
		return "0x8031000E"
	case FVE_E_AD_GUID_NOT_FOUND:
		return "0x8031000F"
	case FVE_E_BAD_INFORMATION:
		return "0x80310010"
	case FVE_E_TOO_SMALL:
		return "0x80310011"
	case FVE_E_SYSTEM_VOLUME:
		return "0x80310012"
	case FVE_E_FAILED_WRONG_FS:
		return "0x80310013"
	case FVE_E_BAD_PARTITION_SIZE:
		return "0x80310014"
	case FVE_E_NOT_SUPPORTED:
		return "0x80310015"
	case FVE_E_BAD_DATA:
		return "0x80310016"
	case FVE_E_VOLUME_NOT_BOUND:
		return "0x80310017"
	case FVE_E_TPM_NOT_OWNED:
		return "0x80310018"
	case FVE_E_NOT_DATA_VOLUME:
		return "0x80310019"
	case FVE_E_AD_INSUFFICIENT_BUFFER:
		return "0x8031001A"
	case FVE_E_CONV_READ:
		return "0x8031001B"
	case FVE_E_CONV_WRITE:
		return "0x8031001C"
	case FVE_E_KEY_REQUIRED:
		return "0x8031001D"
	case FVE_E_CLUSTERING_NOT_SUPPORTED:
		return "0x8031001E"
	case FVE_E_VOLUME_BOUND_ALREADY:
		return "0x8031001F"
	case FVE_E_OS_NOT_PROTECTED:
		return "0x80310020"
	case FVE_E_PROTECTION_DISABLED:
		return "0x80310021"
	case FVE_E_RECOVERY_KEY_REQUIRED:
		return "0x80310022"
	case FVE_E_FOREIGN_VOLUME:
		return "0x80310023"
	case FVE_E_OVERLAPPED_UPDATE:
		return "0x80310024"
	case FVE_E_TPM_SRK_AUTH_NOT_ZERO:
		return "0x80310025"
	case FVE_E_FAILED_SECTOR_SIZE:
		return "0x80310026"
	case FVE_E_FAILED_AUTHENTICATION:
		return "0x80310027"
	case FVE_E_NOT_OS_VOLUME:
		return "0x80310028"
	case FVE_E_AUTOUNLOCK_ENABLED:
		return "0x80310029"
	case FVE_E_WRONG_BOOTSECTOR:
		return "0x8031002A"
	case FVE_E_WRONG_SYSTEM_FS:
		return "0x8031002B"
	case FVE_E_POLICY_PASSWORD_REQUIRED:
		return "0x8031002C"
	case FVE_E_CANNOT_SET_FVEK_ENCRYPTED:
		return "0x8031002D"
	case FVE_E_CANNOT_ENCRYPT_NO_KEY:
		return "0x8031002E"
	case FVE_E_BOOTABLE_CDDVD:
		return "0x80310030"
	case FVE_E_PROTECTOR_EXISTS:
		return "0x80310031"
	case FVE_E_RELATIVE_PATH:
		return "0x80310032"
	case FVE_E_PROTECTOR_NOT_FOUND:
		return "0x80310033"
	case FVE_E_INVALID_KEY_FORMAT:
		return "0x80310034"
	case FVE_E_INVALID_PASSWORD_FORMAT:
		return "0x80310035"
	case FVE_E_FIPS_RNG_CHECK_FAILED:
		return "0x80310036"
	case FVE_E_FIPS_PREVENTS_RECOVERY_PASSWORD:
		return "0x80310037"
	case FVE_E_FIPS_PREVENTS_EXTERNAL_KEY_EXPORT:
		return "0x80310038"
	case FVE_E_NOT_DECRYPTED:
		return "0x80310039"
	case FVE_E_INVALID_PROTECTOR_TYPE:
		return "0x8031003A"
	case FVE_E_NO_PROTECTORS_TO_TEST:
		return "0x8031003B"
	case FVE_E_KEYFILE_NOT_FOUND:
		return "0x8031003C"
	case FVE_E_KEYFILE_INVALID:
		return "0x8031003D"
	case FVE_E_KEYFILE_NO_VMK:
		return "0x8031003E"
	case FVE_E_TPM_DISABLED:
		return "0x8031003F"
	case FVE_E_NOT_ALLOWED_IN_SAFE_MODE:
		return "0x80310040"
	case FVE_E_TPM_INVALID_PCR:
		return "0x80310041"
	case FVE_E_TPM_NO_VMK:
		return "0x80310042"
	case FVE_E_PIN_INVALID:
		return "0x80310043"
	case FVE_E_AUTH_INVALID_APPLICATION:
		return "0x80310044"
	case FVE_E_AUTH_INVALID_CONFIG:
		return "0x80310045"
	case FVE_E_FIPS_DISABLE_PROTECTION_NOT_ALLOWED:
		return "0x80310046"
	case FVE_E_FS_NOT_EXTENDED:
		return "0x80310047"
	case FVE_E_FIRMWARE_TYPE_NOT_SUPPORTED:
		return "0x80310048"
	case FVE_E_NO_LICENSE:
		return "0x80310049"
	case FVE_E_NOT_ON_STACK:
		return "0x8031004A"
	case FVE_E_FS_MOUNTED:
		return "0x8031004B"
	case FVE_E_TOKEN_NOT_IMPERSONATED:
		return "0x8031004C"
	case FVE_E_DRY_RUN_FAILED:
		return "0x8031004D"
	case FVE_E_REBOOT_REQUIRED:
		return "0x8031004E"
	case FVE_E_DEBUGGER_ENABLED:
		return "0x8031004F"
	case FVE_E_RAW_ACCESS:
		return "0x80310050"
	case FVE_E_RAW_BLOCKED:
		return "0x80310051"
	case FVE_E_BCD_APPLICATIONS_PATH_INCORRECT:
		return "0x80310052"
	case FVE_E_NOT_ALLOWED_IN_VERSION:
		return "0x80310053"
	case FVE_E_NO_AUTOUNLOCK_MASTER_KEY:
		return "0x80310054"
	case FVE_E_MOR_FAILED:
		return "0x80310055"
	case FVE_E_HIDDEN_VOLUME:
		return "0x80310056"
	case FVE_E_TRANSIENT_STATE:
		return "0x80310057"
	case FVE_E_PUBKEY_NOT_ALLOWED:
		return "0x80310058"
	case FVE_E_VOLUME_HANDLE_OPEN:
		return "0x80310059"
	case FVE_E_NO_FEATURE_LICENSE:
		return "0x8031005A"
	case FVE_E_INVALID_STARTUP_OPTIONS:
		return "0x8031005B"
	case FVE_E_POLICY_RECOVERY_PASSWORD_NOT_ALLOWED:
		return "0x8031005C"
	case FVE_E_POLICY_RECOVERY_PASSWORD_REQUIRED:
		return "0x8031005D"
	case FVE_E_POLICY_RECOVERY_KEY_NOT_ALLOWED:
		return "0x8031005E"
	case FVE_E_POLICY_RECOVERY_KEY_REQUIRED:
		return "0x8031005F"
	case FVE_E_POLICY_STARTUP_PIN_NOT_ALLOWED:
		return "0x80310060"
	case FVE_E_POLICY_STARTUP_PIN_REQUIRED:
		return "0x80310061"
	case FVE_E_POLICY_STARTUP_KEY_NOT_ALLOWED:
		return "0x80310062"
	case FVE_E_POLICY_STARTUP_KEY_REQUIRED:
		return "0x80310063"
	case FVE_E_POLICY_STARTUP_PIN_KEY_NOT_ALLOWED:
		return "0x80310064"
	case FVE_E_POLICY_STARTUP_PIN_KEY_REQUIRED:
		return "0x80310065"
	case FVE_E_POLICY_STARTUP_TPM_NOT_ALLOWED:
		return "0x80310066"
	case FVE_E_POLICY_STARTUP_TPM_REQUIRED:
		return "0x80310067"
	case FVE_E_POLICY_INVALID_PIN_LENGTH:
		return "0x80310068"
	case FVE_E_KEY_PROTECTOR_NOT_SUPPORTED:
		return "0x80310069"
	case FVE_E_POLICY_PASSPHRASE_NOT_ALLOWED:
		return "0x8031006A"
	case FVE_E_POLICY_PASSPHRASE_REQUIRED:
		return "0x8031006B"
	case FVE_E_FIPS_PREVENTS_PASSPHRASE:
		return "0x8031006C"
	case FVE_E_OS_VOLUME_PASSPHRASE_NOT_ALLOWED:
		return "0x8031006D"
	case FVE_E_INVALID_BITLOCKER_OID:
		return "0x8031006E"
	case FVE_E_VOLUME_TOO_SMALL:
		return "0x8031006F"
	case FVE_E_DV_NOT_SUPPORTED_ON_FS:
		return "0x80310070"
	case FVE_E_DV_NOT_ALLOWED_BY_GP:
		return "0x80310071"
	case FVE_E_POLICY_USER_CERTIFICATE_NOT_ALLOWED:
		return "0x80310072"
	case FVE_E_POLICY_USER_CERTIFICATE_REQUIRED:
		return "0x80310073"
	case FVE_E_POLICY_USER_CERT_MUST_BE_HW:
		return "0x80310074"
	case FVE_E_POLICY_USER_CONFIGURE_FDV_AUTOUNLOCK_NOT_ALLOWED:
		return "0x80310075"
	case FVE_E_POLICY_USER_CONFIGURE_RDV_AUTOUNLOCK_NOT_ALLOWED:
		return "0x80310076"
	case FVE_E_POLICY_USER_CONFIGURE_RDV_NOT_ALLOWED:
		return "0x80310077"
	case FVE_E_POLICY_USER_ENABLE_RDV_NOT_ALLOWED:
		return "0x80310078"
	case FVE_E_POLICY_USER_DISABLE_RDV_NOT_ALLOWED:
		return "0x80310079"
	case FVE_E_POLICY_INVALID_PASSPHRASE_LENGTH:
		return "0x80310080"
	case FVE_E_POLICY_PASSPHRASE_TOO_SIMPLE:
		return "0x80310081"
	case FVE_E_RECOVERY_PARTITION:
		return "0x80310082"
	case FVE_E_POLICY_CONFLICT_FDV_RK_OFF_AUK_ON:
		return "0x80310083"
	case FVE_E_POLICY_CONFLICT_RDV_RK_OFF_AUK_ON:
		return "0x80310084"
	case FVE_E_NON_BITLOCKER_OID:
		return "0x80310085"
	case FVE_E_POLICY_PROHIBITS_SELFSIGNED:
		return "0x80310086"
	case FVE_E_POLICY_CONFLICT_RO_AND_STARTUP_KEY_REQUIRED:
		return "0x80310087"
	case FVE_E_CONV_RECOVERY_FAILED:
		return "0x80310088"
	case FVE_E_VIRTUALIZED_SPACE_TOO_BIG:
		return "0x80310089"
	case FVE_E_POLICY_CONFLICT_OSV_RP_OFF_ADB_ON:
		return "0x80310090"
	case FVE_E_POLICY_CONFLICT_FDV_RP_OFF_ADB_ON:
		return "0x80310091"
	case FVE_E_POLICY_CONFLICT_RDV_RP_OFF_ADB_ON:
		return "0x80310092"
	case FVE_E_NON_BITLOCKER_KU:
		return "0x80310093"
	case FVE_E_PRIVATEKEY_AUTH_FAILED:
		return "0x80310094"
	case FVE_E_REMOVAL_OF_DRA_FAILED:
		return "0x80310095"
	case FVE_E_OPERATION_NOT_SUPPORTED_ON_VISTA_VOLUME:
		return "0x80310096"
	case FVE_E_CANT_LOCK_AUTOUNLOCK_ENABLED_VOLUME:
		return "0x80310097"
	case FVE_E_FIPS_HASH_KDF_NOT_ALLOWED:
		return "0x80310098"
	case FVE_E_ENH_PIN_INVALID:
		return "0x80310099"
	case FVE_E_INVALID_PIN_CHARS:
		return "0x8031009A"
	case FVE_E_INVALID_DATUM_TYPE:
		return "0x8031009B"
	case FWP_E_CALLOUT_NOT_FOUND:
		return "0x80320001"
	case FWP_E_CONDITION_NOT_FOUND:
		return "0x80320002"
	case FWP_E_FILTER_NOT_FOUND:
		return "0x80320003"
	case FWP_E_LAYER_NOT_FOUND:
		return "0x80320004"
	case FWP_E_PROVIDER_NOT_FOUND:
		return "0x80320005"
	case FWP_E_PROVIDER_CONTEXT_NOT_FOUND:
		return "0x80320006"
	case FWP_E_SUBLAYER_NOT_FOUND:
		return "0x80320007"
	case FWP_E_NOT_FOUND:
		return "0x80320008"
	case FWP_E_ALREADY_EXISTS:
		return "0x80320009"
	case FWP_E_IN_USE:
		return "0x8032000A"
	case FWP_E_DYNAMIC_SESSION_IN_PROGRESS:
		return "0x8032000B"
	case FWP_E_WRONG_SESSION:
		return "0x8032000C"
	case FWP_E_NO_TXN_IN_PROGRESS:
		return "0x8032000D"
	case FWP_E_TXN_IN_PROGRESS:
		return "0x8032000E"
	case FWP_E_TXN_ABORTED:
		return "0x8032000F"
	case FWP_E_SESSION_ABORTED:
		return "0x80320010"
	case FWP_E_INCOMPATIBLE_TXN:
		return "0x80320011"
	case FWP_E_TIMEOUT:
		return "0x80320012"
	case FWP_E_NET_EVENTS_DISABLED:
		return "0x80320013"
	case FWP_E_INCOMPATIBLE_LAYER:
		return "0x80320014"
	case FWP_E_KM_CLIENTS_ONLY:
		return "0x80320015"
	case FWP_E_LIFETIME_MISMATCH:
		return "0x80320016"
	case FWP_E_BUILTIN_OBJECT:
		return "0x80320017"
	case FWP_E_TOO_MANY_CALLOUTS:
		return "0x80320018"
	case FWP_E_NOTIFICATION_DROPPED:
		return "0x80320019"
	case FWP_E_TRAFFIC_MISMATCH:
		return "0x8032001A"
	case FWP_E_INCOMPATIBLE_SA_STATE:
		return "0x8032001B"
	case FWP_E_NULL_POINTER:
		return "0x8032001C"
	case FWP_E_INVALID_ENUMERATOR:
		return "0x8032001D"
	case FWP_E_INVALID_FLAGS:
		return "0x8032001E"
	case FWP_E_INVALID_NET_MASK:
		return "0x8032001F"
	case FWP_E_INVALID_RANGE:
		return "0x80320020"
	case FWP_E_INVALID_INTERVAL:
		return "0x80320021"
	case FWP_E_ZERO_LENGTH_ARRAY:
		return "0x80320022"
	case FWP_E_NULL_DISPLAY_NAME:
		return "0x80320023"
	case FWP_E_INVALID_ACTION_TYPE:
		return "0x80320024"
	case FWP_E_INVALID_WEIGHT:
		return "0x80320025"
	case FWP_E_MATCH_TYPE_MISMATCH:
		return "0x80320026"
	case FWP_E_TYPE_MISMATCH:
		return "0x80320027"
	case FWP_E_OUT_OF_BOUNDS:
		return "0x80320028"
	case FWP_E_RESERVED:
		return "0x80320029"
	case FWP_E_DUPLICATE_CONDITION:
		return "0x8032002A"
	case FWP_E_DUPLICATE_KEYMOD:
		return "0x8032002B"
	case FWP_E_ACTION_INCOMPATIBLE_WITH_LAYER:
		return "0x8032002C"
	case FWP_E_ACTION_INCOMPATIBLE_WITH_SUBLAYER:
		return "0x8032002D"
	case FWP_E_CONTEXT_INCOMPATIBLE_WITH_LAYER:
		return "0x8032002E"
	case FWP_E_CONTEXT_INCOMPATIBLE_WITH_CALLOUT:
		return "0x8032002F"
	case FWP_E_INCOMPATIBLE_AUTH_METHOD:
		return "0x80320030"
	case FWP_E_INCOMPATIBLE_DH_GROUP:
		return "0x80320031"
	case FWP_E_EM_NOT_SUPPORTED:
		return "0x80320032"
	case FWP_E_NEVER_MATCH:
		return "0x80320033"
	case FWP_E_PROVIDER_CONTEXT_MISMATCH:
		return "0x80320034"
	case FWP_E_INVALID_PARAMETER:
		return "0x80320035"
	case FWP_E_TOO_MANY_SUBLAYERS:
		return "0x80320036"
	case FWP_E_CALLOUT_NOTIFICATION_FAILED:
		return "0x80320037"
	case FWP_E_INVALID_AUTH_TRANSFORM:
		return "0x80320038"
	case FWP_E_INVALID_CIPHER_TRANSFORM:
		return "0x80320039"
	case FWP_E_DROP_NOICMP:
		return "0x80320104"
	case FWP_E_INCOMPATIBLE_CIPHER_TRANSFORM:
		return "0x8032003A"
	case FWP_E_INVALID_TRANSFORM_COMBINATION:
		return "0x8032003B"
	case FWP_E_DUPLICATE_AUTH_METHOD:
		return "0x8032003C"
	case WS_S_ASYNC:
		return "0x003D0000"
	case WS_S_END:
		return "0x003D0001"
	case WS_E_INVALID_FORMAT:
		return "0x803D0000"
	case WS_E_OBJECT_FAULTED:
		return "0x803D0001"
	case WS_E_NUMERIC_OVERFLOW:
		return "0x803D0002"
	case WS_E_INVALID_OPERATION:
		return "0x803D0003"
	case WS_E_OPERATION_ABORTED:
		return "0x803D0004"
	case WS_E_ENDPOINT_ACCESS_DENIED:
		return "0x803D0005"
	case WS_E_OPERATION_TIMED_OUT:
		return "0x803D0006"
	case WS_E_OPERATION_ABANDONED:
		return "0x803D0007"
	case WS_E_QUOTA_EXCEEDED:
		return "0x803D0008"
	case WS_E_NO_TRANSLATION_AVAILABLE:
		return "0x803D0009"
	case WS_E_SECURITY_VERIFICATION_FAILURE:
		return "0x803D000A"
	case WS_E_ADDRESS_IN_USE:
		return "0x803D000B"
	case WS_E_ADDRESS_NOT_AVAILABLE:
		return "0x803D000C"
	case WS_E_ENDPOINT_NOT_FOUND:
		return "0x803D000D"
	case WS_E_ENDPOINT_NOT_AVAILABLE:
		return "0x803D000E"
	case WS_E_ENDPOINT_FAILURE:
		return "0x803D000F"
	case WS_E_ENDPOINT_UNREACHABLE:
		return "0x803D0010"
	case WS_E_ENDPOINT_ACTION_NOT_SUPPORTED:
		return "0x803D0011"
	case WS_E_ENDPOINT_TOO_BUSY:
		return "0x803D0012"
	case WS_E_ENDPOINT_FAULT_RECEIVED:
		return "0x803D0013"
	case WS_E_ENDPOINT_DISCONNECTED:
		return "0x803D0014"
	case WS_E_PROXY_FAILURE:
		return "0x803D0015"
	case WS_E_PROXY_ACCESS_DENIED:
		return "0x803D0016"
	case WS_E_NOT_SUPPORTED:
		return "0x803D0017"
	case WS_E_PROXY_REQUIRES_BASIC_AUTH:
		return "0x803D0018"
	case WS_E_PROXY_REQUIRES_DIGEST_AUTH:
		return "0x803D0019"
	case WS_E_PROXY_REQUIRES_NTLM_AUTH:
		return "0x803D001A"
	case WS_E_PROXY_REQUIRES_NEGOTIATE_AUTH:
		return "0x803D001B"
	case WS_E_SERVER_REQUIRES_BASIC_AUTH:
		return "0x803D001C"
	case WS_E_SERVER_REQUIRES_DIGEST_AUTH:
		return "0x803D001D"
	case WS_E_SERVER_REQUIRES_NTLM_AUTH:
		return "0x803D001E"
	case WS_E_SERVER_REQUIRES_NEGOTIATE_AUTH:
		return "0x803D001F"
	case WS_E_INVALID_ENDPOINT_URL:
		return "0x803D0020"
	case WS_E_OTHER:
		return "0x803D0021"
	case WS_E_SECURITY_TOKEN_EXPIRED:
		return "0x803D0022"
	case WS_E_SECURITY_SYSTEM_FAILURE:
		return "0x803D0023"
	case ERROR_NDIS_INTERFACE_CLOSING:
		return "0x80340002"
	case ERROR_NDIS_BAD_VERSION:
		return "0x80340004"
	case ERROR_NDIS_BAD_CHARACTERISTICS:
		return "0x80340005"
	case ERROR_NDIS_ADAPTER_NOT_FOUND:
		return "0x80340006"
	case ERROR_NDIS_OPEN_FAILED:
		return "0x80340007"
	case ERROR_NDIS_DEVICE_FAILED:
		return "0x80340008"
	case ERROR_NDIS_MULTICAST_FULL:
		return "0x80340009"
	case ERROR_NDIS_MULTICAST_EXISTS:
		return "0x8034000A"
	case ERROR_NDIS_MULTICAST_NOT_FOUND:
		return "0x8034000B"
	case ERROR_NDIS_REQUEST_ABORTED:
		return "0x8034000C"
	case ERROR_NDIS_RESET_IN_PROGRESS:
		return "0x8034000D"
	case ERROR_NDIS_NOT_SUPPORTED:
		return "0x803400BB"
	case ERROR_NDIS_INVALID_PACKET:
		return "0x8034000F"
	case ERROR_NDIS_ADAPTER_NOT_READY:
		return "0x80340011"
	case ERROR_NDIS_INVALID_LENGTH:
		return "0x80340014"
	case ERROR_NDIS_INVALID_DATA:
		return "0x80340015"
	case ERROR_NDIS_BUFFER_TOO_SHORT:
		return "0x80340016"
	case ERROR_NDIS_INVALID_OID:
		return "0x80340017"
	case ERROR_NDIS_ADAPTER_REMOVED:
		return "0x80340018"
	case ERROR_NDIS_UNSUPPORTED_MEDIA:
		return "0x80340019"
	case ERROR_NDIS_GROUP_ADDRESS_IN_USE:
		return "0x8034001A"
	case ERROR_NDIS_FILE_NOT_FOUND:
		return "0x8034001B"
	case ERROR_NDIS_ERROR_READING_FILE:
		return "0x8034001C"
	case ERROR_NDIS_ALREADY_MAPPED:
		return "0x8034001D"
	case ERROR_NDIS_RESOURCE_CONFLICT:
		return "0x8034001E"
	case ERROR_NDIS_MEDIA_DISCONNECTED:
		return "0x8034001F"
	case ERROR_NDIS_INVALID_ADDRESS:
		return "0x80340022"
	case ERROR_NDIS_INVALID_DEVICE_REQUEST:
		return "0x80340010"
	case ERROR_NDIS_PAUSED:
		return "0x8034002A"
	case ERROR_NDIS_INTERFACE_NOT_FOUND:
		return "0x8034002B"
	case ERROR_NDIS_UNSUPPORTED_REVISION:
		return "0x8034002C"
	case ERROR_NDIS_INVALID_PORT:
		return "0x8034002D"
	case ERROR_NDIS_INVALID_PORT_STATE:
		return "0x8034002E"
	case ERROR_NDIS_LOW_POWER_STATE:
		return "0x8034002F"
	case ERROR_NDIS_DOT11_AUTO_CONFIG_ENABLED:
		return "0x80342000"
	case ERROR_NDIS_DOT11_MEDIA_IN_USE:
		return "0x80342001"
	case ERROR_NDIS_DOT11_POWER_STATE_INVALID:
		return "0x80342002"
	case ERROR_NDIS_PM_WOL_PATTERN_LIST_FULL:
		return "0x80342003"
	case ERROR_NDIS_PM_PROTOCOL_OFFLOAD_LIST_FULL:
		return "0x80342004"
	case ERROR_NDIS_INDICATION_REQUIRED:
		return "0x00340001"
	case ERROR_NDIS_OFFLOAD_POLICY:
		return "0xC034100F"
	case ERROR_NDIS_OFFLOAD_CONNECTION_REJECTED:
		return "0xC0341012"
	case ERROR_NDIS_OFFLOAD_PATH_REJECTED:
		return "0xC0341013"
	case ERROR_HV_INVALID_HYPERCALL_CODE:
		return "0xC0350002"
	case ERROR_HV_INVALID_HYPERCALL_INPUT:
		return "0xC0350003"
	case ERROR_HV_INVALID_ALIGNMENT:
		return "0xC0350004"
	case ERROR_HV_INVALID_PARAMETER:
		return "0xC0350005"
	case ERROR_HV_ACCESS_DENIED:
		return "0xC0350006"
	case ERROR_HV_INVALID_PARTITION_STATE:
		return "0xC0350007"
	case ERROR_HV_OPERATION_DENIED:
		return "0xC0350008"
	case ERROR_HV_UNKNOWN_PROPERTY:
		return "0xC0350009"
	case ERROR_HV_PROPERTY_VALUE_OUT_OF_RANGE:
		return "0xC035000A"
	case ERROR_HV_INSUFFICIENT_MEMORY:
		return "0xC035000B"
	case ERROR_HV_PARTITION_TOO_DEEP:
		return "0xC035000C"
	case ERROR_HV_INVALID_PARTITION_ID:
		return "0xC035000D"
	case ERROR_HV_INVALID_VP_INDEX:
		return "0xC035000E"
	case ERROR_HV_INVALID_PORT_ID:
		return "0xC0350011"
	case ERROR_HV_INVALID_CONNECTION_ID:
		return "0xC0350012"
	case ERROR_HV_INSUFFICIENT_BUFFERS:
		return "0xC0350013"
	case ERROR_HV_NOT_ACKNOWLEDGED:
		return "0xC0350014"
	case ERROR_HV_ACKNOWLEDGED:
		return "0xC0350016"
	case ERROR_HV_INVALID_SAVE_RESTORE_STATE:
		return "0xC0350017"
	case ERROR_HV_INVALID_SYNIC_STATE:
		return "0xC0350018"
	case ERROR_HV_OBJECT_IN_USE:
		return "0xC0350019"
	case ERROR_HV_INVALID_PROXIMITY_DOMAIN_INFO:
		return "0xC035001A"
	case ERROR_HV_NO_DATA:
		return "0xC035001B"
	case ERROR_HV_INACTIVE:
		return "0xC035001C"
	case ERROR_HV_NO_RESOURCES:
		return "0xC035001D"
	case ERROR_HV_FEATURE_UNAVAILABLE:
		return "0xC035001E"
	case ERROR_HV_NOT_PRESENT:
		return "0xC0351000"
	case ERROR_VID_DUPLICATE_HANDLER:
		return "0xC0370001"
	case ERROR_VID_TOO_MANY_HANDLERS:
		return "0xC0370002"
	case ERROR_VID_QUEUE_FULL:
		return "0xC0370003"
	case ERROR_VID_HANDLER_NOT_PRESENT:
		return "0xC0370004"
	case ERROR_VID_INVALID_OBJECT_NAME:
		return "0xC0370005"
	case ERROR_VID_PARTITION_NAME_TOO_LONG:
		return "0xC0370006"
	case ERROR_VID_MESSAGE_QUEUE_NAME_TOO_LONG:
		return "0xC0370007"
	case ERROR_VID_PARTITION_ALREADY_EXISTS:
		return "0xC0370008"
	case ERROR_VID_PARTITION_DOES_NOT_EXIST:
		return "0xC0370009"
	case ERROR_VID_PARTITION_NAME_NOT_FOUND:
		return "0xC037000A"
	case ERROR_VID_MESSAGE_QUEUE_ALREADY_EXISTS:
		return "0xC037000B"
	case ERROR_VID_EXCEEDED_MBP_ENTRY_MAP_LIMIT:
		return "0xC037000C"
	case ERROR_VID_MB_STILL_REFERENCED:
		return "0xC037000D"
	case ERROR_VID_CHILD_GPA_PAGE_SET_CORRUPTED:
		return "0xC037000E"
	case ERROR_VID_INVALID_NUMA_SETTINGS:
		return "0xC037000F"
	case ERROR_VID_INVALID_NUMA_NODE_INDEX:
		return "0xC0370010"
	case ERROR_VID_NOTIFICATION_QUEUE_ALREADY_ASSOCIATED:
		return "0xC0370011"
	case ERROR_VID_INVALID_MEMORY_BLOCK_HANDLE:
		return "0xC0370012"
	case ERROR_VID_PAGE_RANGE_OVERFLOW:
		return "0xC0370013"
	case ERROR_VID_INVALID_MESSAGE_QUEUE_HANDLE:
		return "0xC0370014"
	case ERROR_VID_INVALID_GPA_RANGE_HANDLE:
		return "0xC0370015"
	case ERROR_VID_NO_MEMORY_BLOCK_NOTIFICATION_QUEUE:
		return "0xC0370016"
	case ERROR_VID_MEMORY_BLOCK_LOCK_COUNT_EXCEEDED:
		return "0xC0370017"
	case ERROR_VID_INVALID_PPM_HANDLE:
		return "0xC0370018"
	case ERROR_VID_MBPS_ARE_LOCKED:
		return "0xC0370019"
	case ERROR_VID_MESSAGE_QUEUE_CLOSED:
		return "0xC037001A"
	case ERROR_VID_VIRTUAL_PROCESSOR_LIMIT_EXCEEDED:
		return "0xC037001B"
	case ERROR_VID_STOP_PENDING:
		return "0xC037001C"
	case ERROR_VID_INVALID_PROCESSOR_STATE:
		return "0xC037001D"
	case ERROR_VID_EXCEEDED_KM_CONTEXT_COUNT_LIMIT:
		return "0xC037001E"
	case ERROR_VID_KM_INTERFACE_ALREADY_INITIALIZED:
		return "0xC037001F"
	case ERROR_VID_MB_PROPERTY_ALREADY_SET_RESET:
		return "0xC0370020"
	case ERROR_VID_MMIO_RANGE_DESTROYED:
		return "0xC0370021"
	case ERROR_VID_INVALID_CHILD_GPA_PAGE_SET:
		return "0xC0370022"
	case ERROR_VID_RESERVE_PAGE_SET_IS_BEING_USED:
		return "0xC0370023"
	case ERROR_VID_RESERVE_PAGE_SET_TOO_SMALL:
		return "0xC0370024"
	case ERROR_VID_MBP_ALREADY_LOCKED_USING_RESERVED_PAGE:
		return "0xC0370025"
	case ERROR_VID_MBP_COUNT_EXCEEDED_LIMIT:
		return "0xC0370026"
	case ERROR_VID_SAVED_STATE_CORRUPT:
		return "0xC0370027"
	case ERROR_VID_SAVED_STATE_UNRECOGNIZED_ITEM:
		return "0xC0370028"
	case ERROR_VID_SAVED_STATE_INCOMPATIBLE:
		return "0xC0370029"
	case ERROR_VID_REMOTE_NODE_PARENT_GPA_PAGES_USED:
		return "0x80370001"
	case ERROR_VOLMGR_INCOMPLETE_REGENERATION:
		return "0x80380001"
	case ERROR_VOLMGR_INCOMPLETE_DISK_MIGRATION:
		return "0x80380002"
	case ERROR_VOLMGR_DATABASE_FULL:
		return "0xC0380001"
	case ERROR_VOLMGR_DISK_CONFIGURATION_CORRUPTED:
		return "0xC0380002"
	case ERROR_VOLMGR_DISK_CONFIGURATION_NOT_IN_SYNC:
		return "0xC0380003"
	case ERROR_VOLMGR_PACK_CONFIG_UPDATE_FAILED:
		return "0xC0380004"
	case ERROR_VOLMGR_DISK_CONTAINS_NON_SIMPLE_VOLUME:
		return "0xC0380005"
	case ERROR_VOLMGR_DISK_DUPLICATE:
		return "0xC0380006"
	case ERROR_VOLMGR_DISK_DYNAMIC:
		return "0xC0380007"
	case ERROR_VOLMGR_DISK_ID_INVALID:
		return "0xC0380008"
	case ERROR_VOLMGR_DISK_INVALID:
		return "0xC0380009"
	case ERROR_VOLMGR_DISK_LAST_VOTER:
		return "0xC038000A"
	case ERROR_VOLMGR_DISK_LAYOUT_INVALID:
		return "0xC038000B"
	case ERROR_VOLMGR_DISK_LAYOUT_NON_BASIC_BETWEEN_BASIC_PARTITIONS:
		return "0xC038000C"
	case ERROR_VOLMGR_DISK_LAYOUT_NOT_CYLINDER_ALIGNED:
		return "0xC038000D"
	case ERROR_VOLMGR_DISK_LAYOUT_PARTITIONS_TOO_SMALL:
		return "0xC038000E"
	case ERROR_VOLMGR_DISK_LAYOUT_PRIMARY_BETWEEN_LOGICAL_PARTITIONS:
		return "0xC038000F"
	case ERROR_VOLMGR_DISK_LAYOUT_TOO_MANY_PARTITIONS:
		return "0xC0380010"
	case ERROR_VOLMGR_DISK_MISSING:
		return "0xC0380011"
	case ERROR_VOLMGR_DISK_NOT_EMPTY:
		return "0xC0380012"
	case ERROR_VOLMGR_DISK_NOT_ENOUGH_SPACE:
		return "0xC0380013"
	case ERROR_VOLMGR_DISK_REVECTORING_FAILED:
		return "0xC0380014"
	case ERROR_VOLMGR_DISK_SECTOR_SIZE_INVALID:
		return "0xC0380015"
	case ERROR_VOLMGR_DISK_SET_NOT_CONTAINED:
		return "0xC0380016"
	case ERROR_VOLMGR_DISK_USED_BY_MULTIPLE_MEMBERS:
		return "0xC0380017"
	case ERROR_VOLMGR_DISK_USED_BY_MULTIPLE_PLEXES:
		return "0xC0380018"
	case ERROR_VOLMGR_DYNAMIC_DISK_NOT_SUPPORTED:
		return "0xC0380019"
	case ERROR_VOLMGR_EXTENT_ALREADY_USED:
		return "0xC038001A"
	case ERROR_VOLMGR_EXTENT_NOT_CONTIGUOUS:
		return "0xC038001B"
	case ERROR_VOLMGR_EXTENT_NOT_IN_PUBLIC_REGION:
		return "0xC038001C"
	case ERROR_VOLMGR_EXTENT_NOT_SECTOR_ALIGNED:
		return "0xC038001D"
	case ERROR_VOLMGR_EXTENT_OVERLAPS_EBR_PARTITION:
		return "0xC038001E"
	case ERROR_VOLMGR_EXTENT_VOLUME_LENGTHS_DO_NOT_MATCH:
		return "0xC038001F"
	case ERROR_VOLMGR_FAULT_TOLERANT_NOT_SUPPORTED:
		return "0xC0380020"
	case ERROR_VOLMGR_INTERLEAVE_LENGTH_INVALID:
		return "0xC0380021"
	case ERROR_VOLMGR_MAXIMUM_REGISTERED_USERS:
		return "0xC0380022"
	case ERROR_VOLMGR_MEMBER_IN_SYNC:
		return "0xC0380023"
	case ERROR_VOLMGR_MEMBER_INDEX_DUPLICATE:
		return "0xC0380024"
	case ERROR_VOLMGR_MEMBER_INDEX_INVALID:
		return "0xC0380025"
	case ERROR_VOLMGR_MEMBER_MISSING:
		return "0xC0380026"
	case ERROR_VOLMGR_MEMBER_NOT_DETACHED:
		return "0xC0380027"
	case ERROR_VOLMGR_MEMBER_REGENERATING:
		return "0xC0380028"
	case ERROR_VOLMGR_ALL_DISKS_FAILED:
		return "0xC0380029"
	case ERROR_VOLMGR_NO_REGISTERED_USERS:
		return "0xC038002A"
	case ERROR_VOLMGR_NO_SUCH_USER:
		return "0xC038002B"
	case ERROR_VOLMGR_NOTIFICATION_RESET:
		return "0xC038002C"
	case ERROR_VOLMGR_NUMBER_OF_MEMBERS_INVALID:
		return "0xC038002D"
	case ERROR_VOLMGR_NUMBER_OF_PLEXES_INVALID:
		return "0xC038002E"
	case ERROR_VOLMGR_PACK_DUPLICATE:
		return "0xC038002F"
	case ERROR_VOLMGR_PACK_ID_INVALID:
		return "0xC0380030"
	case ERROR_VOLMGR_PACK_INVALID:
		return "0xC0380031"
	case ERROR_VOLMGR_PACK_NAME_INVALID:
		return "0xC0380032"
	case ERROR_VOLMGR_PACK_OFFLINE:
		return "0xC0380033"
	case ERROR_VOLMGR_PACK_HAS_QUORUM:
		return "0xC0380034"
	case ERROR_VOLMGR_PACK_WITHOUT_QUORUM:
		return "0xC0380035"
	case ERROR_VOLMGR_PARTITION_STYLE_INVALID:
		return "0xC0380036"
	case ERROR_VOLMGR_PARTITION_UPDATE_FAILED:
		return "0xC0380037"
	case ERROR_VOLMGR_PLEX_IN_SYNC:
		return "0xC0380038"
	case ERROR_VOLMGR_PLEX_INDEX_DUPLICATE:
		return "0xC0380039"
	case ERROR_VOLMGR_PLEX_INDEX_INVALID:
		return "0xC038003A"
	case ERROR_VOLMGR_PLEX_LAST_ACTIVE:
		return "0xC038003B"
	case ERROR_VOLMGR_PLEX_MISSING:
		return "0xC038003C"
	case ERROR_VOLMGR_PLEX_REGENERATING:
		return "0xC038003D"
	case ERROR_VOLMGR_PLEX_TYPE_INVALID:
		return "0xC038003E"
	case ERROR_VOLMGR_PLEX_NOT_RAID5:
		return "0xC038003F"
	case ERROR_VOLMGR_PLEX_NOT_SIMPLE:
		return "0xC0380040"
	case ERROR_VOLMGR_STRUCTURE_SIZE_INVALID:
		return "0xC0380041"
	case ERROR_VOLMGR_TOO_MANY_NOTIFICATION_REQUESTS:
		return "0xC0380042"
	case ERROR_VOLMGR_TRANSACTION_IN_PROGRESS:
		return "0xC0380043"
	case ERROR_VOLMGR_UNEXPECTED_DISK_LAYOUT_CHANGE:
		return "0xC0380044"
	case ERROR_VOLMGR_VOLUME_CONTAINS_MISSING_DISK:
		return "0xC0380045"
	case ERROR_VOLMGR_VOLUME_ID_INVALID:
		return "0xC0380046"
	case ERROR_VOLMGR_VOLUME_LENGTH_INVALID:
		return "0xC0380047"
	case ERROR_VOLMGR_VOLUME_LENGTH_NOT_SECTOR_SIZE_MULTIPLE:
		return "0xC0380048"
	case ERROR_VOLMGR_VOLUME_NOT_MIRRORED:
		return "0xC0380049"
	case ERROR_VOLMGR_VOLUME_NOT_RETAINED:
		return "0xC038004A"
	case ERROR_VOLMGR_VOLUME_OFFLINE:
		return "0xC038004B"
	case ERROR_VOLMGR_VOLUME_RETAINED:
		return "0xC038004C"
	case ERROR_VOLMGR_NUMBER_OF_EXTENTS_INVALID:
		return "0xC038004D"
	case ERROR_VOLMGR_DIFFERENT_SECTOR_SIZE:
		return "0xC038004E"
	case ERROR_VOLMGR_BAD_BOOT_DISK:
		return "0xC038004F"
	case ERROR_VOLMGR_PACK_CONFIG_OFFLINE:
		return "0xC0380050"
	case ERROR_VOLMGR_PACK_CONFIG_ONLINE:
		return "0xC0380051"
	case ERROR_VOLMGR_NOT_PRIMARY_PACK:
		return "0xC0380052"
	case ERROR_VOLMGR_PACK_LOG_UPDATE_FAILED:
		return "0xC0380053"
	case ERROR_VOLMGR_NUMBER_OF_DISKS_IN_PLEX_INVALID:
		return "0xC0380054"
	case ERROR_VOLMGR_NUMBER_OF_DISKS_IN_MEMBER_INVALID:
		return "0xC0380055"
	case ERROR_VOLMGR_VOLUME_MIRRORED:
		return "0xC0380056"
	case ERROR_VOLMGR_PLEX_NOT_SIMPLE_SPANNED:
		return "0xC0380057"
	case ERROR_VOLMGR_NO_VALID_LOG_COPIES:
		return "0xC0380058"
	case ERROR_VOLMGR_PRIMARY_PACK_PRESENT:
		return "0xC0380059"
	case ERROR_VOLMGR_NUMBER_OF_DISKS_INVALID:
		return "0xC038005A"
	case ERROR_VOLMGR_MIRROR_NOT_SUPPORTED:
		return "0xC038005B"
	case ERROR_VOLMGR_RAID5_NOT_SUPPORTED:
		return "0xC038005C"
	case ERROR_BCD_NOT_ALL_ENTRIES_IMPORTED:
		return "0x80390001"
	case ERROR_BCD_TOO_MANY_ELEMENTS:
		return "0xC0390002"
	case ERROR_BCD_NOT_ALL_ENTRIES_SYNCHRONIZED:
		return "0x80390003"
	case ERROR_VHD_DRIVE_FOOTER_MISSING:
		return "0xC03A0001"
	case ERROR_VHD_DRIVE_FOOTER_CHECKSUM_MISMATCH:
		return "0xC03A0002"
	case ERROR_VHD_DRIVE_FOOTER_CORRUPT:
		return "0xC03A0003"
	case ERROR_VHD_FORMAT_UNKNOWN:
		return "0xC03A0004"
	case ERROR_VHD_FORMAT_UNSUPPORTED_VERSION:
		return "0xC03A0005"
	case ERROR_VHD_SPARSE_HEADER_CHECKSUM_MISMATCH:
		return "0xC03A0006"
	case ERROR_VHD_SPARSE_HEADER_UNSUPPORTED_VERSION:
		return "0xC03A0007"
	case ERROR_VHD_SPARSE_HEADER_CORRUPT:
		return "0xC03A0008"
	case ERROR_VHD_BLOCK_ALLOCATION_FAILURE:
		return "0xC03A0009"
	case ERROR_VHD_BLOCK_ALLOCATION_TABLE_CORRUPT:
		return "0xC03A000A"
	case ERROR_VHD_INVALID_BLOCK_SIZE:
		return "0xC03A000B"
	case ERROR_VHD_BITMAP_MISMATCH:
		return "0xC03A000C"
	case ERROR_VHD_PARENT_VHD_NOT_FOUND:
		return "0xC03A000D"
	case ERROR_VHD_CHILD_PARENT_ID_MISMATCH:
		return "0xC03A000E"
	case ERROR_VHD_CHILD_PARENT_TIMESTAMP_MISMATCH:
		return "0xC03A000F"
	case ERROR_VHD_METADATA_READ_FAILURE:
		return "0xC03A0010"
	case ERROR_VHD_METADATA_WRITE_FAILURE:
		return "0xC03A0011"
	case ERROR_VHD_INVALID_SIZE:
		return "0xC03A0012"
	case ERROR_VHD_INVALID_FILE_SIZE:
		return "0xC03A0013"
	case ERROR_VIRTDISK_PROVIDER_NOT_FOUND:
		return "0xC03A0014"
	case ERROR_VIRTDISK_NOT_VIRTUAL_DISK:
		return "0xC03A0015"
	case ERROR_VHD_PARENT_VHD_ACCESS_DENIED:
		return "0xC03A0016"
	case ERROR_VHD_CHILD_PARENT_SIZE_MISMATCH:
		return "0xC03A0017"
	case ERROR_VHD_DIFFERENCING_CHAIN_CYCLE_DETECTED:
		return "0xC03A0018"
	case ERROR_VHD_DIFFERENCING_CHAIN_ERROR_IN_PARENT:
		return "0xC03A0019"
	case ERROR_VIRTUAL_DISK_LIMITATION:
		return "0xC03A001A"
	case ERROR_VHD_INVALID_TYPE:
		return "0xC03A001B"
	case ERROR_VHD_INVALID_STATE:
		return "0xC03A001C"
	case ERROR_VIRTDISK_UNSUPPORTED_DISK_SECTOR_SIZE:
		return "0xC03A001D"
	case ERROR_QUERY_STORAGE_ERROR:
		return "0x803A0001"
	case SDIAG_E_CANCELLED:
		return "0x803C0100"
	case SDIAG_E_SCRIPT:
		return "0x803C0101"
	case SDIAG_E_POWERSHELL:
		return "0x803C0102"
	case SDIAG_E_MANAGEDHOST:
		return "0x803C0103"
	case SDIAG_E_NOVERIFIER:
		return "0x803C0104"
	case SDIAG_S_CANNOTRUN:
		return "0x003C0105"
	case SDIAG_E_DISABLED:
		return "0x803C0106"
	case SDIAG_E_TRUST:
		return "0x803C0107"
	case SDIAG_E_CANNOTRUN:
		return "0x803C0108"
	case SDIAG_E_VERSION:
		return "0x803C0109"
	case SDIAG_E_RESOURCE:
		return "0x803C010A"
	case SDIAG_E_ROOTCAUSE:
		return "0x803C010B"
	case E_MBN_CONTEXT_NOT_ACTIVATED:
		return "0x80548201"
	case E_MBN_BAD_SIM:
		return "0x80548202"
	case E_MBN_DATA_CLASS_NOT_AVAILABLE:
		return "0x80548203"
	case E_MBN_INVALID_ACCESS_STRING:
		return "0x80548204"
	case E_MBN_MAX_ACTIVATED_CONTEXTS:
		return "0x80548205"
	case E_MBN_PACKET_SVC_DETACHED:
		return "0x80548206"
	case E_MBN_PROVIDER_NOT_VISIBLE:
		return "0x80548207"
	case E_MBN_RADIO_POWER_OFF:
		return "0x80548208"
	case E_MBN_SERVICE_NOT_ACTIVATED:
		return "0x80548209"
	case E_MBN_SIM_NOT_INSERTED:
		return "0x8054820A"
	case E_MBN_VOICE_CALL_IN_PROGRESS:
		return "0x8054820B"
	case E_MBN_INVALID_CACHE:
		return "0x8054820C"
	case E_MBN_NOT_REGISTERED:
		return "0x8054820D"
	case E_MBN_PROVIDERS_NOT_FOUND:
		return "0x8054820E"
	case E_MBN_PIN_NOT_SUPPORTED:
		return "0x8054820F"
	case E_MBN_PIN_REQUIRED:
		return "0x80548210"
	case E_MBN_PIN_DISABLED:
		return "0x80548211"
	case E_MBN_FAILURE:
		return "0x80548212"
	case E_MBN_INVALID_PROFILE:
		return "0x80548218"
	case E_MBN_DEFAULT_PROFILE_EXIST:
		return "0x80548219"
	case E_MBN_SMS_ENCODING_NOT_SUPPORTED:
		return "0x80548220"
	case E_MBN_SMS_FILTER_NOT_SUPPORTED:
		return "0x80548221"
	case E_MBN_SMS_INVALID_MEMORY_INDEX:
		return "0x80548222"
	case E_MBN_SMS_LANG_NOT_SUPPORTED:
		return "0x80548223"
	case E_MBN_SMS_MEMORY_FAILURE:
		return "0x80548224"
	case E_MBN_SMS_NETWORK_TIMEOUT:
		return "0x80548225"
	case E_MBN_SMS_UNKNOWN_SMSC_ADDRESS:
		return "0x80548226"
	case E_MBN_SMS_FORMAT_NOT_SUPPORTED:
		return "0x80548227"
	case E_MBN_SMS_OPERATION_NOT_ALLOWED:
		return "0x80548228"
	case E_MBN_SMS_MEMORY_FULL:
		return "0x80548229"
	case UI_E_CREATE_FAILED:
		return "0x802A0001"
	case UI_E_SHUTDOWN_CALLED:
		return "0x802A0002"
	case UI_E_ILLEGAL_REENTRANCY:
		return "0x802A0003"
	case UI_E_OBJECT_SEALED:
		return "0x802A0004"
	case UI_E_VALUE_NOT_SET:
		return "0x802A0005"
	case UI_E_VALUE_NOT_DETERMINED:
		return "0x802A0006"
	case UI_E_INVALID_OUTPUT:
		return "0x802A0007"
	case UI_E_BOOLEAN_EXPECTED:
		return "0x802A0008"
	case UI_E_DIFFERENT_OWNER:
		return "0x802A0009"
	case UI_E_AMBIGUOUS_MATCH:
		return "0x802A000A"
	case UI_E_FP_OVERFLOW:
		return "0x802A000B"
	case UI_E_WRONG_THREAD:
		return "0x802A000C"
	case UI_E_STORYBOARD_ACTIVE:
		return "0x802A0101"
	case UI_E_STORYBOARD_NOT_PLAYING:
		return "0x802A0102"
	case UI_E_START_KEYFRAME_AFTER_END:
		return "0x802A0103"
	case UI_E_END_KEYFRAME_NOT_DETERMINED:
		return "0x802A0104"
	case UI_E_LOOPS_OVERLAP:
		return "0x802A0105"
	case UI_E_TRANSITION_ALREADY_USED:
		return "0x802A0106"
	case UI_E_TRANSITION_NOT_IN_STORYBOARD:
		return "0x802A0107"
	case UI_E_TRANSITION_ECLIPSED:
		return "0x802A0108"
	case UI_E_TIME_BEFORE_LAST_UPDATE:
		return "0x802A0109"
	case UI_E_TIMER_CLIENT_ALREADY_CONNECTED:
		return "0x802A010A"
	default:
		return "unknown"
	}
}

func (k ErrorsKind) Elements() []ErrorsKind {
	return []ErrorsKind{
		ERROR_SUCCESS,
		ERROR_INVALID_FUNCTION,
		ERROR_FILE_NOT_FOUND,
		ERROR_PATH_NOT_FOUND,
		ERROR_TOO_MANY_OPEN_FILES,
		ERROR_ACCESS_DENIED,
		ERROR_INVALID_HANDLE,
		ERROR_ARENA_TRASHED,
		ERROR_NOT_ENOUGH_MEMORY,
		ERROR_INVALID_BLOCK,
		ERROR_BAD_ENVIRONMENT,
		ERROR_BAD_FORMAT,
		ERROR_INVALID_ACCESS,
		ERROR_INVALID_DATA,
		ERROR_OUTOFMEMORY,
		ERROR_INVALID_DRIVE,
		ERROR_CURRENT_DIRECTORY,
		ERROR_NOT_SAME_DEVICE,
		ERROR_NO_MORE_FILES,
		ERROR_WRITE_PROTECT,
		ERROR_BAD_UNIT,
		ERROR_NOT_READY,
		ERROR_BAD_COMMAND,
		ERROR_CRC,
		ERROR_BAD_LENGTH,
		ERROR_SEEK,
		ERROR_NOT_DOS_DISK,
		ERROR_SECTOR_NOT_FOUND,
		ERROR_OUT_OF_PAPER,
		ERROR_WRITE_FAULT,
		ERROR_READ_FAULT,
		ERROR_GEN_FAILURE,
		ERROR_SHARING_VIOLATION,
		ERROR_LOCK_VIOLATION,
		ERROR_WRONG_DISK,
		ERROR_SHARING_BUFFER_EXCEEDED,
		ERROR_HANDLE_EOF,
		ERROR_HANDLE_DISK_FULL,
		ERROR_NOT_SUPPORTED,
		ERROR_REM_NOT_LIST,
		ERROR_DUP_NAME,
		ERROR_BAD_NETPATH,
		ERROR_NETWORK_BUSY,
		ERROR_DEV_NOT_EXIST,
		ERROR_TOO_MANY_CMDS,
		ERROR_ADAP_HDW_ERR,
		ERROR_BAD_NET_RESP,
		ERROR_UNEXP_NET_ERR,
		ERROR_BAD_REM_ADAP,
		ERROR_PRINTQ_FULL,
		ERROR_NO_SPOOL_SPACE,
		ERROR_PRINT_CANCELLED,
		ERROR_NETNAME_DELETED,
		ERROR_NETWORK_ACCESS_DENIED,
		ERROR_BAD_DEV_TYPE,
		ERROR_BAD_NET_NAME,
		ERROR_TOO_MANY_NAMES,
		ERROR_TOO_MANY_SESS,
		ERROR_SHARING_PAUSED,
		ERROR_REQ_NOT_ACCEP,
		ERROR_REDIR_PAUSED,
		ERROR_FILE_EXISTS,
		ERROR_CANNOT_MAKE,
		ERROR_FAIL_I24,
		ERROR_OUT_OF_STRUCTURES,
		ERROR_ALREADY_ASSIGNED,
		ERROR_INVALID_PASSWORD,
		ERROR_INVALID_PARAMETER,
		ERROR_NET_WRITE_FAULT,
		ERROR_NO_PROC_SLOTS,
		ERROR_TOO_MANY_SEMAPHORES,
		ERROR_EXCL_SEM_ALREADY_OWNED,
		ERROR_SEM_IS_SET,
		ERROR_TOO_MANY_SEM_REQUESTS,
		ERROR_INVALID_AT_INTERRUPT_TIME,
		ERROR_SEM_OWNER_DIED,
		ERROR_SEM_USER_LIMIT,
		ERROR_DISK_CHANGE,
		ERROR_DRIVE_LOCKED,
		ERROR_BROKEN_PIPE,
		ERROR_OPEN_FAILED,
		ERROR_BUFFER_OVERFLOW,
		ERROR_DISK_FULL,
		ERROR_NO_MORE_SEARCH_HANDLES,
		ERROR_INVALID_TARGET_HANDLE,
		ERROR_INVALID_CATEGORY,
		ERROR_INVALID_VERIFY_SWITCH,
		ERROR_BAD_DRIVER_LEVEL,
		ERROR_CALL_NOT_IMPLEMENTED,
		ERROR_SEM_TIMEOUT,
		ERROR_INSUFFICIENT_BUFFER,
		ERROR_INVALID_NAME,
		ERROR_INVALID_LEVEL,
		ERROR_NO_VOLUME_LABEL,
		ERROR_MOD_NOT_FOUND,
		ERROR_PROC_NOT_FOUND,
		ERROR_WAIT_NO_CHILDREN,
		ERROR_CHILD_NOT_COMPLETE,
		ERROR_DIRECT_ACCESS_HANDLE,
		ERROR_NEGATIVE_SEEK,
		ERROR_SEEK_ON_DEVICE,
		ERROR_IS_JOIN_TARGET,
		ERROR_IS_JOINED,
		ERROR_IS_SUBSTED,
		ERROR_NOT_JOINED,
		ERROR_NOT_SUBSTED,
		ERROR_JOIN_TO_JOIN,
		ERROR_SUBST_TO_SUBST,
		ERROR_JOIN_TO_SUBST,
		ERROR_SUBST_TO_JOIN,
		ERROR_BUSY_DRIVE,
		ERROR_SAME_DRIVE,
		ERROR_DIR_NOT_ROOT,
		ERROR_DIR_NOT_EMPTY,
		ERROR_IS_SUBST_PATH,
		ERROR_IS_JOIN_PATH,
		ERROR_PATH_BUSY,
		ERROR_IS_SUBST_TARGET,
		ERROR_SYSTEM_TRACE,
		ERROR_INVALID_EVENT_COUNT,
		ERROR_TOO_MANY_MUXWAITERS,
		ERROR_INVALID_LIST_FORMAT,
		ERROR_LABEL_TOO_LONG,
		ERROR_TOO_MANY_TCBS,
		ERROR_SIGNAL_REFUSED,
		ERROR_DISCARDED,
		ERROR_NOT_LOCKED,
		ERROR_BAD_THREADID_ADDR,
		ERROR_BAD_ARGUMENTS,
		ERROR_BAD_PATHNAME,
		ERROR_SIGNAL_PENDING,
		ERROR_MAX_THRDS_REACHED,
		ERROR_LOCK_FAILED,
		ERROR_BUSY,
		ERROR_CANCEL_VIOLATION,
		ERROR_ATOMIC_LOCKS_NOT_SUPPORTED,
		ERROR_INVALID_SEGMENT_NUMBER,
		ERROR_INVALID_ORDINAL,
		ERROR_ALREADY_EXISTS,
		ERROR_INVALID_FLAG_NUMBER,
		ERROR_SEM_NOT_FOUND,
		ERROR_INVALID_STARTING_CODESEG,
		ERROR_INVALID_STACKSEG,
		ERROR_INVALID_MODULETYPE,
		ERROR_INVALID_EXE_SIGNATURE,
		ERROR_EXE_MARKED_INVALID,
		ERROR_BAD_EXE_FORMAT,
		ERROR_ITERATED_DATA_EXCEEDS_64k,
		ERROR_INVALID_MINALLOCSIZE,
		ERROR_DYNLINK_FROM_INVALID_RING,
		ERROR_IOPL_NOT_ENABLED,
		ERROR_INVALID_SEGDPL,
		ERROR_AUTODATASEG_EXCEEDS_64k,
		ERROR_RING2SEG_MUST_BE_MOVABLE,
		ERROR_RELOC_CHAIN_XEEDS_SEGLIM,
		ERROR_INFLOOP_IN_RELOC_CHAIN,
		ERROR_ENVVAR_NOT_FOUND,
		ERROR_NO_SIGNAL_SENT,
		ERROR_FILENAME_EXCED_RANGE,
		ERROR_RING2_STACK_IN_USE,
		ERROR_META_EXPANSION_TOO_LONG,
		ERROR_INVALID_SIGNAL_NUMBER,
		ERROR_THREAD_1_INACTIVE,
		ERROR_LOCKED,
		ERROR_TOO_MANY_MODULES,
		ERROR_NESTING_NOT_ALLOWED,
		ERROR_EXE_MACHINE_TYPE_MISMATCH,
		ERROR_EXE_CANNOT_MODIFY_SIGNED_BINARY,
		ERROR_EXE_CANNOT_MODIFY_STRONG_SIGNED_BINARY,
		ERROR_FILE_CHECKED_OUT,
		ERROR_CHECKOUT_REQUIRED,
		ERROR_BAD_FILE_TYPE,
		ERROR_FILE_TOO_LARGE,
		ERROR_FORMS_AUTH_REQUIRED,
		ERROR_VIRUS_INFECTED,
		ERROR_VIRUS_DELETED,
		ERROR_PIPE_LOCAL,
		ERROR_BAD_PIPE,
		ERROR_PIPE_BUSY,
		ERROR_NO_DATA,
		ERROR_PIPE_NOT_CONNECTED,
		ERROR_MORE_DATA,
		ERROR_VC_DISCONNECTED,
		ERROR_INVALID_EA_NAME,
		ERROR_EA_LIST_INCONSISTENT,
		WAIT_TIMEOUT,
		ERROR_NO_MORE_ITEMS,
		ERROR_CANNOT_COPY,
		ERROR_DIRECTORY,
		ERROR_EAS_DIDNT_FIT,
		ERROR_EA_FILE_CORRUPT,
		ERROR_EA_TABLE_FULL,
		ERROR_INVALID_EA_HANDLE,
		ERROR_EAS_NOT_SUPPORTED,
		ERROR_NOT_OWNER,
		ERROR_TOO_MANY_POSTS,
		ERROR_PARTIAL_COPY,
		ERROR_OPLOCK_NOT_GRANTED,
		ERROR_INVALID_OPLOCK_PROTOCOL,
		ERROR_DISK_TOO_FRAGMENTED,
		ERROR_DELETE_PENDING,
		ERROR_INCOMPATIBLE_WITH_GLOBAL_SHORT_NAME_REGISTRY_SETTING,
		ERROR_SHORT_NAMES_NOT_ENABLED_ON_VOLUME,
		ERROR_SECURITY_STREAM_IS_INCONSISTENT,
		ERROR_INVALID_LOCK_RANGE,
		ERROR_IMAGE_SUBSYSTEM_NOT_PRESENT,
		ERROR_NOTIFICATION_GUID_ALREADY_DEFINED,
		ERROR_MR_MID_NOT_FOUND,
		ERROR_SCOPE_NOT_FOUND,
		ERROR_FAIL_NOACTION_REBOOT,
		ERROR_FAIL_SHUTDOWN,
		ERROR_FAIL_RESTART,
		ERROR_MAX_SESSIONS_REACHED,
		ERROR_THREAD_MODE_ALREADY_BACKGROUND,
		ERROR_THREAD_MODE_NOT_BACKGROUND,
		ERROR_PROCESS_MODE_ALREADY_BACKGROUND,
		ERROR_PROCESS_MODE_NOT_BACKGROUND,
		ERROR_INVALID_ADDRESS,
		ERROR_USER_PROFILE_LOAD,
		ERROR_ARITHMETIC_OVERFLOW,
		ERROR_PIPE_CONNECTED,
		ERROR_PIPE_LISTENING,
		ERROR_VERIFIER_STOP,
		ERROR_ABIOS_ERROR,
		ERROR_WX86_WARNING,
		ERROR_WX86_ERROR,
		ERROR_TIMER_NOT_CANCELED,
		ERROR_UNWIND,
		ERROR_BAD_STACK,
		ERROR_INVALID_UNWIND_TARGET,
		ERROR_INVALID_PORT_ATTRIBUTES,
		ERROR_PORT_MESSAGE_TOO_LONG,
		ERROR_INVALID_QUOTA_LOWER,
		ERROR_DEVICE_ALREADY_ATTACHED,
		ERROR_INSTRUCTION_MISALIGNMENT,
		ERROR_PROFILING_NOT_STARTED,
		ERROR_PROFILING_NOT_STOPPED,
		ERROR_COULD_NOT_INTERPRET,
		ERROR_PROFILING_AT_LIMIT,
		ERROR_CANT_WAIT,
		ERROR_CANT_TERMINATE_SELF,
		ERROR_UNEXPECTED_MM_CREATE_ERR,
		ERROR_UNEXPECTED_MM_MAP_ERROR,
		ERROR_UNEXPECTED_MM_EXTEND_ERR,
		ERROR_BAD_FUNCTION_TABLE,
		ERROR_NO_GUID_TRANSLATION,
		ERROR_INVALID_LDT_SIZE,
		ERROR_INVALID_LDT_OFFSET,
		ERROR_INVALID_LDT_DESCRIPTOR,
		ERROR_TOO_MANY_THREADS,
		ERROR_THREAD_NOT_IN_PROCESS,
		ERROR_PAGEFILE_QUOTA_EXCEEDED,
		ERROR_LOGON_SERVER_CONFLICT,
		ERROR_SYNCHRONIZATION_REQUIRED,
		ERROR_NET_OPEN_FAILED,
		ERROR_IO_PRIVILEGE_FAILED,
		ERROR_CONTROL_C_EXIT,
		ERROR_MISSING_SYSTEMFILE,
		ERROR_UNHANDLED_EXCEPTION,
		ERROR_APP_INIT_FAILURE,
		ERROR_PAGEFILE_CREATE_FAILED,
		ERROR_INVALID_IMAGE_HASH,
		ERROR_NO_PAGEFILE,
		ERROR_ILLEGAL_FLOAT_CONTEXT,
		ERROR_NO_EVENT_PAIR,
		ERROR_DOMAIN_CTRLR_CONFIG_ERROR,
		ERROR_ILLEGAL_CHARACTER,
		ERROR_UNDEFINED_CHARACTER,
		ERROR_FLOPPY_VOLUME,
		ERROR_BIOS_FAILED_TO_CONNECT_INTERRUPT,
		ERROR_BACKUP_CONTROLLER,
		ERROR_MUTANT_LIMIT_EXCEEDED,
		ERROR_FS_DRIVER_REQUIRED,
		ERROR_CANNOT_LOAD_REGISTRY_FILE,
		ERROR_DEBUG_ATTACH_FAILED,
		ERROR_SYSTEM_PROCESS_TERMINATED,
		ERROR_DATA_NOT_ACCEPTED,
		ERROR_VDM_HARD_ERROR,
		ERROR_DRIVER_CANCEL_TIMEOUT,
		ERROR_REPLY_MESSAGE_MISMATCH,
		ERROR_LOST_WRITEBEHIND_DATA,
		ERROR_CLIENT_SERVER_PARAMETERS_INVALID,
		ERROR_NOT_TINY_STREAM,
		ERROR_STACK_OVERFLOW_READ,
		ERROR_CONVERT_TO_LARGE,
		ERROR_FOUND_OUT_OF_SCOPE,
		ERROR_ALLOCATE_BUCKET,
		ERROR_MARSHALL_OVERFLOW,
		ERROR_INVALID_VARIANT,
		ERROR_BAD_COMPRESSION_BUFFER,
		ERROR_AUDIT_FAILED,
		ERROR_TIMER_RESOLUTION_NOT_SET,
		ERROR_INSUFFICIENT_LOGON_INFO,
		ERROR_BAD_DLL_ENTRYPOINT,
		ERROR_BAD_SERVICE_ENTRYPOINT,
		ERROR_IP_ADDRESS_CONFLICT1,
		ERROR_IP_ADDRESS_CONFLICT2,
		ERROR_REGISTRY_QUOTA_LIMIT,
		ERROR_NO_CALLBACK_ACTIVE,
		ERROR_PWD_TOO_SHORT,
		ERROR_PWD_TOO_RECENT,
		ERROR_PWD_HISTORY_CONFLICT,
		ERROR_UNSUPPORTED_COMPRESSION,
		ERROR_INVALID_HW_PROFILE,
		ERROR_INVALID_PLUGPLAY_DEVICE_PATH,
		ERROR_QUOTA_LIST_INCONSISTENT,
		ERROR_EVALUATION_EXPIRATION,
		ERROR_ILLEGAL_DLL_RELOCATION,
		ERROR_DLL_INIT_FAILED_LOGOFF,
		ERROR_VALIDATE_CONTINUE,
		ERROR_NO_MORE_MATCHES,
		ERROR_RANGE_LIST_CONFLICT,
		ERROR_SERVER_SID_MISMATCH,
		ERROR_CANT_ENABLE_DENY_ONLY,
		ERROR_FLOAT_MULTIPLE_FAULTS,
		ERROR_FLOAT_MULTIPLE_TRAPS,
		ERROR_NOINTERFACE,
		ERROR_DRIVER_FAILED_SLEEP,
		ERROR_CORRUPT_SYSTEM_FILE,
		ERROR_COMMITMENT_MINIMUM,
		ERROR_PNP_RESTART_ENUMERATION,
		ERROR_SYSTEM_IMAGE_BAD_SIGNATURE,
		ERROR_PNP_REBOOT_REQUIRED,
		ERROR_INSUFFICIENT_POWER,
		ERROR_MULTIPLE_FAULT_VIOLATION,
		ERROR_SYSTEM_SHUTDOWN,
		ERROR_PORT_NOT_SET,
		ERROR_DS_VERSION_CHECK_FAILURE,
		ERROR_RANGE_NOT_FOUND,
		ERROR_NOT_SAFE_MODE_DRIVER,
		ERROR_FAILED_DRIVER_ENTRY,
		ERROR_DEVICE_ENUMERATION_ERROR,
		ERROR_MOUNT_POINT_NOT_RESOLVED,
		ERROR_INVALID_DEVICE_OBJECT_PARAMETER,
		ERROR_MCA_OCCURED,
		ERROR_DRIVER_DATABASE_ERROR,
		ERROR_SYSTEM_HIVE_TOO_LARGE,
		ERROR_DRIVER_FAILED_PRIOR_UNLOAD,
		ERROR_VOLSNAP_PREPARE_HIBERNATE,
		ERROR_HIBERNATION_FAILURE,
		ERROR_FILE_SYSTEM_LIMITATION,
		ERROR_ASSERTION_FAILURE,
		ERROR_ACPI_ERROR,
		ERROR_WOW_ASSERTION,
		ERROR_PNP_BAD_MPS_TABLE,
		ERROR_PNP_TRANSLATION_FAILED,
		ERROR_PNP_IRQ_TRANSLATION_FAILED,
		ERROR_PNP_INVALID_ID,
		ERROR_WAKE_SYSTEM_DEBUGGER,
		ERROR_HANDLES_CLOSED,
		ERROR_EXTRANEOUS_INFORMATION,
		ERROR_RXACT_COMMIT_NECESSARY,
		ERROR_MEDIA_CHECK,
		ERROR_GUID_SUBSTITUTION_MADE,
		ERROR_STOPPED_ON_SYMLINK,
		ERROR_LONGJUMP,
		ERROR_PLUGPLAY_QUERY_VETOED,
		ERROR_UNWIND_CONSOLIDATE,
		ERROR_REGISTRY_HIVE_RECOVERED,
		ERROR_DLL_MIGHT_BE_INSECURE,
		ERROR_DLL_MIGHT_BE_INCOMPATIBLE,
		ERROR_DBG_EXCEPTION_NOT_HANDLED,
		ERROR_DBG_REPLY_LATER,
		ERROR_DBG_UNABLE_TO_PROVIDE_HANDLE,
		ERROR_DBG_TERMINATE_THREAD,
		ERROR_DBG_TERMINATE_PROCESS,
		ERROR_DBG_CONTROL_C,
		ERROR_DBG_PRINTEXCEPTION_C,
		ERROR_DBG_RIPEXCEPTION,
		ERROR_DBG_CONTROL_BREAK,
		ERROR_DBG_COMMAND_EXCEPTION,
		ERROR_OBJECT_NAME_EXISTS,
		ERROR_THREAD_WAS_SUSPENDED,
		ERROR_IMAGE_NOT_AT_BASE,
		ERROR_RXACT_STATE_CREATED,
		ERROR_SEGMENT_NOTIFICATION,
		ERROR_BAD_CURRENT_DIRECTORY,
		ERROR_FT_READ_RECOVERY_FROM_BACKUP,
		ERROR_FT_WRITE_RECOVERY,
		ERROR_IMAGE_MACHINE_TYPE_MISMATCH,
		ERROR_RECEIVE_PARTIAL,
		ERROR_RECEIVE_EXPEDITED,
		ERROR_RECEIVE_PARTIAL_EXPEDITED,
		ERROR_EVENT_DONE,
		ERROR_EVENT_PENDING,
		ERROR_CHECKING_FILE_SYSTEM,
		ERROR_FATAL_APP_EXIT,
		ERROR_PREDEFINED_HANDLE,
		ERROR_WAS_UNLOCKED,
		ERROR_SERVICE_NOTIFICATION,
		ERROR_WAS_LOCKED,
		ERROR_LOG_HARD_ERROR,
		ERROR_ALREADY_WIN32,
		ERROR_IMAGE_MACHINE_TYPE_MISMATCH_EXE,
		ERROR_NO_YIELD_PERFORMED,
		ERROR_TIMER_RESUME_IGNORED,
		ERROR_ARBITRATION_UNHANDLED,
		ERROR_CARDBUS_NOT_SUPPORTED,
		ERROR_MP_PROCESSOR_MISMATCH,
		ERROR_HIBERNATED,
		ERROR_RESUME_HIBERNATION,
		ERROR_FIRMWARE_UPDATED,
		ERROR_DRIVERS_LEAKING_LOCKED_PAGES,
		ERROR_WAKE_SYSTEM,
		ERROR_WAIT_1,
		ERROR_WAIT_2,
		ERROR_WAIT_3,
		ERROR_WAIT_63,
		ERROR_ABANDONED_WAIT_0,
		ERROR_ABANDONED_WAIT_63,
		ERROR_USER_APC,
		ERROR_KERNEL_APC,
		ERROR_ALERTED,
		ERROR_ELEVATION_REQUIRED,
		ERROR_REPARSE,
		ERROR_OPLOCK_BREAK_IN_PROGRESS,
		ERROR_VOLUME_MOUNTED,
		ERROR_RXACT_COMMITTED,
		ERROR_NOTIFY_CLEANUP,
		ERROR_PRIMARY_TRANSPORT_CONNECT_FAILED,
		ERROR_PAGE_FAULT_TRANSITION,
		ERROR_PAGE_FAULT_DEMAND_ZERO,
		ERROR_PAGE_FAULT_COPY_ON_WRITE,
		ERROR_PAGE_FAULT_GUARD_PAGE,
		ERROR_PAGE_FAULT_PAGING_FILE,
		ERROR_CACHE_PAGE_LOCKED,
		ERROR_CRASH_DUMP,
		ERROR_BUFFER_ALL_ZEROS,
		ERROR_REPARSE_OBJECT,
		ERROR_RESOURCE_REQUIREMENTS_CHANGED,
		ERROR_TRANSLATION_COMPLETE,
		ERROR_NOTHING_TO_TERMINATE,
		ERROR_PROCESS_NOT_IN_JOB,
		ERROR_PROCESS_IN_JOB,
		ERROR_VOLSNAP_HIBERNATE_READY,
		ERROR_FSFILTER_OP_COMPLETED_SUCCESSFULLY,
		ERROR_INTERRUPT_VECTOR_ALREADY_CONNECTED,
		ERROR_INTERRUPT_STILL_CONNECTED,
		ERROR_WAIT_FOR_OPLOCK,
		ERROR_DBG_EXCEPTION_HANDLED,
		ERROR_DBG_CONTINUE,
		ERROR_CALLBACK_POP_STACK,
		ERROR_COMPRESSION_DISABLED,
		ERROR_CANTFETCHBACKWARDS,
		ERROR_CANTSCROLLBACKWARDS,
		ERROR_ROWSNOTRELEASED,
		ERROR_BAD_ACCESSOR_FLAGS,
		ERROR_ERRORS_ENCOUNTERED,
		ERROR_NOT_CAPABLE,
		ERROR_REQUEST_OUT_OF_SEQUENCE,
		ERROR_VERSION_PARSE_ERROR,
		ERROR_BADSTARTPOSITION,
		ERROR_MEMORY_HARDWARE,
		ERROR_DISK_REPAIR_DISABLED,
		ERROR_INSUFFICIENT_RESOURCE_FOR_SPECIFIED_SHARED_SECTION_SIZE,
		ERROR_SYSTEM_POWERSTATE_TRANSITION,
		ERROR_SYSTEM_POWERSTATE_COMPLEX_TRANSITION,
		ERROR_MCA_EXCEPTION,
		ERROR_ACCESS_AUDIT_BY_POLICY,
		ERROR_ACCESS_DISABLED_NO_SAFER_UI_BY_POLICY,
		ERROR_ABANDON_HIBERFILE,
		ERROR_LOST_WRITEBEHIND_DATA_NETWORK_DISCONNECTED,
		ERROR_LOST_WRITEBEHIND_DATA_NETWORK_SERVER_ERROR,
		ERROR_LOST_WRITEBEHIND_DATA_LOCAL_DISK_ERROR,
		ERROR_BAD_MCFG_TABLE,
		ERROR_OPLOCK_SWITCHED_TO_NEW_HANDLE,
		ERROR_CANNOT_GRANT_REQUESTED_OPLOCK,
		ERROR_CANNOT_BREAK_OPLOCK,
		ERROR_OPLOCK_HANDLE_CLOSED,
		ERROR_NO_ACE_CONDITION,
		ERROR_INVALID_ACE_CONDITION,
		ERROR_EA_ACCESS_DENIED,
		ERROR_OPERATION_ABORTED,
		ERROR_IO_INCOMPLETE,
		ERROR_IO_PENDING,
		ERROR_NOACCESS,
		ERROR_SWAPERROR,
		ERROR_STACK_OVERFLOW,
		ERROR_INVALID_MESSAGE,
		ERROR_CAN_NOT_COMPLETE,
		ERROR_INVALID_FLAGS,
		ERROR_UNRECOGNIZED_VOLUME,
		ERROR_FILE_INVALID,
		ERROR_FULLSCREEN_MODE,
		ERROR_NO_TOKEN,
		ERROR_BADDB,
		ERROR_BADKEY,
		ERROR_CANTOPEN,
		ERROR_CANTREAD,
		ERROR_CANTWRITE,
		ERROR_REGISTRY_RECOVERED,
		ERROR_REGISTRY_CORRUPT,
		ERROR_REGISTRY_IO_FAILED,
		ERROR_NOT_REGISTRY_FILE,
		ERROR_KEY_DELETED,
		ERROR_NO_LOG_SPACE,
		ERROR_KEY_HAS_CHILDREN,
		ERROR_CHILD_MUST_BE_VOLATILE,
		ERROR_NOTIFY_ENUM_DIR,
		ERROR_DEPENDENT_SERVICES_RUNNING,
		ERROR_INVALID_SERVICE_CONTROL,
		ERROR_SERVICE_REQUEST_TIMEOUT,
		ERROR_SERVICE_NO_THREAD,
		ERROR_SERVICE_DATABASE_LOCKED,
		ERROR_SERVICE_ALREADY_RUNNING,
		ERROR_INVALID_SERVICE_ACCOUNT,
		ERROR_SERVICE_DISABLED,
		ERROR_CIRCULAR_DEPENDENCY,
		ERROR_SERVICE_DOES_NOT_EXIST,
		ERROR_SERVICE_CANNOT_ACCEPT_CTRL,
		ERROR_SERVICE_NOT_ACTIVE,
		ERROR_FAILED_SERVICE_CONTROLLER_CONNECT,
		ERROR_EXCEPTION_IN_SERVICE,
		ERROR_DATABASE_DOES_NOT_EXIST,
		ERROR_SERVICE_SPECIFIC_ERROR,
		ERROR_PROCESS_ABORTED,
		ERROR_SERVICE_DEPENDENCY_FAIL,
		ERROR_SERVICE_LOGON_FAILED,
		ERROR_SERVICE_START_HANG,
		ERROR_INVALID_SERVICE_LOCK,
		ERROR_SERVICE_MARKED_FOR_DELETE,
		ERROR_SERVICE_EXISTS,
		ERROR_ALREADY_RUNNING_LKG,
		ERROR_SERVICE_DEPENDENCY_DELETED,
		ERROR_BOOT_ALREADY_ACCEPTED,
		ERROR_SERVICE_NEVER_STARTED,
		ERROR_DUPLICATE_SERVICE_NAME,
		ERROR_DIFFERENT_SERVICE_ACCOUNT,
		ERROR_CANNOT_DETECT_DRIVER_FAILURE,
		ERROR_CANNOT_DETECT_PROCESS_ABORT,
		ERROR_NO_RECOVERY_PROGRAM,
		ERROR_SERVICE_NOT_IN_EXE,
		ERROR_NOT_SAFEBOOT_SERVICE,
		ERROR_END_OF_MEDIA,
		ERROR_FILEMARK_DETECTED,
		ERROR_BEGINNING_OF_MEDIA,
		ERROR_SETMARK_DETECTED,
		ERROR_NO_DATA_DETECTED,
		ERROR_PARTITION_FAILURE,
		ERROR_INVALID_BLOCK_LENGTH,
		ERROR_DEVICE_NOT_PARTITIONED,
		ERROR_UNABLE_TO_LOCK_MEDIA,
		ERROR_UNABLE_TO_UNLOAD_MEDIA,
		ERROR_MEDIA_CHANGED,
		ERROR_BUS_RESET,
		ERROR_NO_MEDIA_IN_DRIVE,
		ERROR_NO_UNICODE_TRANSLATION,
		ERROR_DLL_INIT_FAILED,
		ERROR_SHUTDOWN_IN_PROGRESS,
		ERROR_NO_SHUTDOWN_IN_PROGRESS,
		ERROR_IO_DEVICE,
		ERROR_SERIAL_NO_DEVICE,
		ERROR_IRQ_BUSY,
		ERROR_MORE_WRITES,
		ERROR_COUNTER_TIMEOUT,
		ERROR_FLOPPY_ID_MARK_NOT_FOUND,
		ERROR_FLOPPY_WRONG_CYLINDER,
		ERROR_FLOPPY_UNKNOWN_ERROR,
		ERROR_FLOPPY_BAD_REGISTERS,
		ERROR_DISK_RECALIBRATE_FAILED,
		ERROR_DISK_OPERATION_FAILED,
		ERROR_DISK_RESET_FAILED,
		ERROR_EOM_OVERFLOW,
		ERROR_NOT_ENOUGH_SERVER_MEMORY,
		ERROR_POSSIBLE_DEADLOCK,
		ERROR_MAPPED_ALIGNMENT,
		ERROR_SET_POWER_STATE_VETOED,
		ERROR_SET_POWER_STATE_FAILED,
		ERROR_TOO_MANY_LINKS,
		ERROR_OLD_WIN_VERSION,
		ERROR_APP_WRONG_OS,
		ERROR_SINGLE_INSTANCE_APP,
		ERROR_RMODE_APP,
		ERROR_INVALID_DLL,
		ERROR_NO_ASSOCIATION,
		ERROR_DDE_FAIL,
		ERROR_DLL_NOT_FOUND,
		ERROR_NO_MORE_USER_HANDLES,
		ERROR_MESSAGE_SYNC_ONLY,
		ERROR_SOURCE_ELEMENT_EMPTY,
		ERROR_DESTINATION_ELEMENT_FULL,
		ERROR_ILLEGAL_ELEMENT_ADDRESS,
		ERROR_MAGAZINE_NOT_PRESENT,
		ERROR_DEVICE_REINITIALIZATION_NEEDED,
		ERROR_DEVICE_REQUIRES_CLEANING,
		ERROR_DEVICE_DOOR_OPEN,
		ERROR_DEVICE_NOT_CONNECTED,
		ERROR_NOT_FOUND,
		ERROR_NO_MATCH,
		ERROR_SET_NOT_FOUND,
		ERROR_POINT_NOT_FOUND,
		ERROR_NO_TRACKING_SERVICE,
		ERROR_NO_VOLUME_ID,
		ERROR_UNABLE_TO_REMOVE_REPLACED,
		ERROR_UNABLE_TO_MOVE_REPLACEMENT,
		ERROR_UNABLE_TO_MOVE_REPLACEMENT_2,
		ERROR_JOURNAL_DELETE_IN_PROGRESS,
		ERROR_JOURNAL_NOT_ACTIVE,
		ERROR_POTENTIAL_FILE_FOUND,
		ERROR_JOURNAL_ENTRY_DELETED,
		ERROR_SHUTDOWN_IS_SCHEDULED,
		ERROR_SHUTDOWN_USERS_LOGGED_ON,
		ERROR_BAD_DEVICE,
		ERROR_CONNECTION_UNAVAIL,
		ERROR_DEVICE_ALREADY_REMEMBERED,
		ERROR_NO_NET_OR_BAD_PATH,
		ERROR_BAD_PROVIDER,
		ERROR_CANNOT_OPEN_PROFILE,
		ERROR_BAD_PROFILE,
		ERROR_NOT_CONTAINER,
		ERROR_EXTENDED_ERROR,
		ERROR_INVALID_GROUPNAME,
		ERROR_INVALID_COMPUTERNAME,
		ERROR_INVALID_EVENTNAME,
		ERROR_INVALID_DOMAINNAME,
		ERROR_INVALID_SERVICENAME,
		ERROR_INVALID_NETNAME,
		ERROR_INVALID_SHARENAME,
		ERROR_INVALID_PASSWORDNAME,
		ERROR_INVALID_MESSAGENAME,
		ERROR_INVALID_MESSAGEDEST,
		ERROR_SESSION_CREDENTIAL_CONFLICT,
		ERROR_REMOTE_SESSION_LIMIT_EXCEEDED,
		ERROR_DUP_DOMAINNAME,
		ERROR_NO_NETWORK,
		ERROR_CANCELLED,
		ERROR_USER_MAPPED_FILE,
		ERROR_CONNECTION_REFUSED,
		ERROR_GRACEFUL_DISCONNECT,
		ERROR_ADDRESS_ALREADY_ASSOCIATED,
		ERROR_ADDRESS_NOT_ASSOCIATED,
		ERROR_CONNECTION_INVALID,
		ERROR_CONNECTION_ACTIVE,
		ERROR_NETWORK_UNREACHABLE,
		ERROR_HOST_UNREACHABLE,
		ERROR_PROTOCOL_UNREACHABLE,
		ERROR_PORT_UNREACHABLE,
		ERROR_REQUEST_ABORTED,
		ERROR_CONNECTION_ABORTED,
		ERROR_RETRY,
		ERROR_CONNECTION_COUNT_LIMIT,
		ERROR_LOGIN_TIME_RESTRICTION,
		ERROR_LOGIN_WKSTA_RESTRICTION,
		ERROR_INCORRECT_ADDRESS,
		ERROR_ALREADY_REGISTERED,
		ERROR_SERVICE_NOT_FOUND,
		ERROR_NOT_AUTHENTICATED,
		ERROR_NOT_LOGGED_ON,
		ERROR_CONTINUE,
		ERROR_ALREADY_INITIALIZED,
		ERROR_NO_MORE_DEVICES,
		ERROR_NO_SUCH_SITE,
		ERROR_DOMAIN_CONTROLLER_EXISTS,
		ERROR_ONLY_IF_CONNECTED,
		ERROR_OVERRIDE_NOCHANGES,
		ERROR_BAD_USER_PROFILE,
		ERROR_NOT_SUPPORTED_ON_SBS,
		ERROR_SERVER_SHUTDOWN_IN_PROGRESS,
		ERROR_HOST_DOWN,
		ERROR_NON_ACCOUNT_SID,
		ERROR_NON_DOMAIN_SID,
		ERROR_APPHELP_BLOCK,
		ERROR_ACCESS_DISABLED_BY_POLICY,
		ERROR_REG_NAT_CONSUMPTION,
		ERROR_CSCSHARE_OFFLINE,
		ERROR_PKINIT_FAILURE,
		ERROR_SMARTCARD_SUBSYSTEM_FAILURE,
		ERROR_DOWNGRADE_DETECTED,
		ERROR_MACHINE_LOCKED,
		ERROR_CALLBACK_SUPPLIED_INVALID_DATA,
		ERROR_SYNC_FOREGROUND_REFRESH_REQUIRED,
		ERROR_DRIVER_BLOCKED,
		ERROR_INVALID_IMPORT_OF_NON_DLL,
		ERROR_ACCESS_DISABLED_WEBBLADE,
		ERROR_ACCESS_DISABLED_WEBBLADE_TAMPER,
		ERROR_RECOVERY_FAILURE,
		ERROR_ALREADY_FIBER,
		ERROR_ALREADY_THREAD,
		ERROR_STACK_BUFFER_OVERRUN,
		ERROR_PARAMETER_QUOTA_EXCEEDED,
		ERROR_DEBUGGER_INACTIVE,
		ERROR_DELAY_LOAD_FAILED,
		ERROR_VDM_DISALLOWED,
		ERROR_UNIDENTIFIED_ERROR,
		ERROR_INVALID_CRUNTIME_PARAMETER,
		ERROR_BEYOND_VDL,
		ERROR_INCOMPATIBLE_SERVICE_SID_TYPE,
		ERROR_DRIVER_PROCESS_TERMINATED,
		ERROR_IMPLEMENTATION_LIMIT,
		ERROR_PROCESS_IS_PROTECTED,
		ERROR_SERVICE_NOTIFY_CLIENT_LAGGING,
		ERROR_DISK_QUOTA_EXCEEDED,
		ERROR_CONTENT_BLOCKED,
		ERROR_INCOMPATIBLE_SERVICE_PRIVILEGE,
		ERROR_APP_HANG,
		ERROR_INVALID_LABEL,
		ERROR_NOT_ALL_ASSIGNED,
		ERROR_SOME_NOT_MAPPED,
		ERROR_NO_QUOTAS_FOR_ACCOUNT,
		ERROR_LOCAL_USER_SESSION_KEY,
		ERROR_NULL_LM_PASSWORD,
		ERROR_UNKNOWN_REVISION,
		ERROR_REVISION_MISMATCH,
		ERROR_INVALID_OWNER,
		ERROR_INVALID_PRIMARY_GROUP,
		ERROR_NO_IMPERSONATION_TOKEN,
		ERROR_CANT_DISABLE_MANDATORY,
		ERROR_NO_LOGON_SERVERS,
		ERROR_NO_SUCH_LOGON_SESSION,
		ERROR_NO_SUCH_PRIVILEGE,
		ERROR_PRIVILEGE_NOT_HELD,
		ERROR_INVALID_ACCOUNT_NAME,
		ERROR_USER_EXISTS,
		ERROR_NO_SUCH_USER,
		ERROR_GROUP_EXISTS,
		ERROR_NO_SUCH_GROUP,
		ERROR_MEMBER_IN_GROUP,
		ERROR_MEMBER_NOT_IN_GROUP,
		ERROR_LAST_ADMIN,
		ERROR_WRONG_PASSWORD,
		ERROR_ILL_FORMED_PASSWORD,
		ERROR_PASSWORD_RESTRICTION,
		ERROR_LOGON_FAILURE,
		ERROR_ACCOUNT_RESTRICTION,
		ERROR_INVALID_LOGON_HOURS,
		ERROR_INVALID_WORKSTATION,
		ERROR_PASSWORD_EXPIRED,
		ERROR_ACCOUNT_DISABLED,
		ERROR_NONE_MAPPED,
		ERROR_TOO_MANY_LUIDS_REQUESTED,
		ERROR_LUIDS_EXHAUSTED,
		ERROR_INVALID_SUB_AUTHORITY,
		ERROR_INVALID_ACL,
		ERROR_INVALID_SID,
		ERROR_INVALID_SECURITY_DESCR,
		ERROR_BAD_INHERITANCE_ACL,
		ERROR_SERVER_DISABLED,
		ERROR_SERVER_NOT_DISABLED,
		ERROR_INVALID_ID_AUTHORITY,
		ERROR_ALLOTTED_SPACE_EXCEEDED,
		ERROR_INVALID_GROUP_ATTRIBUTES,
		ERROR_BAD_IMPERSONATION_LEVEL,
		ERROR_CANT_OPEN_ANONYMOUS,
		ERROR_BAD_VALIDATION_CLASS,
		ERROR_BAD_TOKEN_TYPE,
		ERROR_NO_SECURITY_ON_OBJECT,
		ERROR_CANT_ACCESS_DOMAIN_INFO,
		ERROR_INVALID_SERVER_STATE,
		ERROR_INVALID_DOMAIN_STATE,
		ERROR_INVALID_DOMAIN_ROLE,
		ERROR_NO_SUCH_DOMAIN,
		ERROR_DOMAIN_EXISTS,
		ERROR_DOMAIN_LIMIT_EXCEEDED,
		ERROR_INTERNAL_DB_CORRUPTION,
		ERROR_INTERNAL_ERROR,
		ERROR_GENERIC_NOT_MAPPED,
		ERROR_BAD_DESCRIPTOR_FORMAT,
		ERROR_NOT_LOGON_PROCESS,
		ERROR_LOGON_SESSION_EXISTS,
		ERROR_NO_SUCH_PACKAGE,
		ERROR_BAD_LOGON_SESSION_STATE,
		ERROR_LOGON_SESSION_COLLISION,
		ERROR_INVALID_LOGON_TYPE,
		ERROR_CANNOT_IMPERSONATE,
		ERROR_RXACT_INVALID_STATE,
		ERROR_RXACT_COMMIT_FAILURE,
		ERROR_SPECIAL_ACCOUNT,
		ERROR_SPECIAL_GROUP,
		ERROR_SPECIAL_USER,
		ERROR_MEMBERS_PRIMARY_GROUP,
		ERROR_TOKEN_ALREADY_IN_USE,
		ERROR_NO_SUCH_ALIAS,
		ERROR_MEMBER_NOT_IN_ALIAS,
		ERROR_MEMBER_IN_ALIAS,
		ERROR_ALIAS_EXISTS,
		ERROR_LOGON_NOT_GRANTED,
		ERROR_TOO_MANY_SECRETS,
		ERROR_SECRET_TOO_LONG,
		ERROR_INTERNAL_DB_ERROR,
		ERROR_TOO_MANY_CONTEXT_IDS,
		ERROR_LOGON_TYPE_NOT_GRANTED,
		ERROR_NT_CROSS_ENCRYPTION_REQUIRED,
		ERROR_NO_SUCH_MEMBER,
		ERROR_INVALID_MEMBER,
		ERROR_TOO_MANY_SIDS,
		ERROR_LM_CROSS_ENCRYPTION_REQUIRED,
		ERROR_NO_INHERITANCE,
		ERROR_FILE_CORRUPT,
		ERROR_DISK_CORRUPT,
		ERROR_NO_USER_SESSION_KEY,
		ERROR_LICENSE_QUOTA_EXCEEDED,
		ERROR_WRONG_TARGET_NAME,
		ERROR_MUTUAL_AUTH_FAILED,
		ERROR_TIME_SKEW,
		ERROR_CURRENT_DOMAIN_NOT_ALLOWED,
		ERROR_INVALID_WINDOW_HANDLE,
		ERROR_INVALID_MENU_HANDLE,
		ERROR_INVALID_CURSOR_HANDLE,
		ERROR_INVALID_ACCEL_HANDLE,
		ERROR_INVALID_HOOK_HANDLE,
		ERROR_INVALID_DWP_HANDLE,
		ERROR_TLW_WITH_WSCHILD,
		ERROR_CANNOT_FIND_WND_CLASS,
		ERROR_WINDOW_OF_OTHER_THREAD,
		ERROR_HOTKEY_ALREADY_REGISTERED,
		ERROR_CLASS_ALREADY_EXISTS,
		ERROR_CLASS_DOES_NOT_EXIST,
		ERROR_CLASS_HAS_WINDOWS,
		ERROR_INVALID_INDEX,
		ERROR_INVALID_ICON_HANDLE,
		ERROR_PRIVATE_DIALOG_INDEX,
		ERROR_LISTBOX_ID_NOT_FOUND,
		ERROR_NO_WILDCARD_CHARACTERS,
		ERROR_CLIPBOARD_NOT_OPEN,
		ERROR_HOTKEY_NOT_REGISTERED,
		ERROR_WINDOW_NOT_DIALOG,
		ERROR_CONTROL_ID_NOT_FOUND,
		ERROR_INVALID_COMBOBOX_MESSAGE,
		ERROR_WINDOW_NOT_COMBOBOX,
		ERROR_INVALID_EDIT_HEIGHT,
		ERROR_DC_NOT_FOUND,
		ERROR_INVALID_HOOK_FILTER,
		ERROR_INVALID_FILTER_PROC,
		ERROR_HOOK_NEEDS_HMOD,
		ERROR_GLOBAL_ONLY_HOOK,
		ERROR_JOURNAL_HOOK_SET,
		ERROR_HOOK_NOT_INSTALLED,
		ERROR_INVALID_LB_MESSAGE,
		ERROR_SETCOUNT_ON_BAD_LB,
		ERROR_LB_WITHOUT_TABSTOPS,
		ERROR_DESTROY_OBJECT_OF_OTHER_THREAD,
		ERROR_CHILD_WINDOW_MENU,
		ERROR_NO_SYSTEM_MENU,
		ERROR_INVALID_MSGBOX_STYLE,
		ERROR_INVALID_SPI_VALUE,
		ERROR_SCREEN_ALREADY_LOCKED,
		ERROR_HWNDS_HAVE_DIFF_PARENT,
		ERROR_NOT_CHILD_WINDOW,
		ERROR_INVALID_GW_COMMAND,
		ERROR_INVALID_THREAD_ID,
		ERROR_NON_MDICHILD_WINDOW,
		ERROR_POPUP_ALREADY_ACTIVE,
		ERROR_NO_SCROLLBARS,
		ERROR_INVALID_SCROLLBAR_RANGE,
		ERROR_INVALID_SHOWWIN_COMMAND,
		ERROR_NO_SYSTEM_RESOURCES,
		ERROR_NONPAGED_SYSTEM_RESOURCES,
		ERROR_PAGED_SYSTEM_RESOURCES,
		ERROR_WORKING_SET_QUOTA,
		ERROR_PAGEFILE_QUOTA,
		ERROR_COMMITMENT_LIMIT,
		ERROR_MENU_ITEM_NOT_FOUND,
		ERROR_INVALID_KEYBOARD_HANDLE,
		ERROR_HOOK_TYPE_NOT_ALLOWED,
		ERROR_REQUIRES_INTERACTIVE_WINDOWSTATION,
		ERROR_TIMEOUT,
		ERROR_INVALID_MONITOR_HANDLE,
		ERROR_INCORRECT_SIZE,
		ERROR_SYMLINK_CLASS_DISABLED,
		ERROR_SYMLINK_NOT_SUPPORTED,
		ERROR_XML_PARSE_ERROR,
		ERROR_XMLDSIG_ERROR,
		ERROR_RESTART_APPLICATION,
		ERROR_WRONG_COMPARTMENT,
		ERROR_AUTHIP_FAILURE,
		ERROR_NO_NVRAM_RESOURCES,
		ERROR_EVENTLOG_FILE_CORRUPT,
		ERROR_EVENTLOG_CANT_START,
		ERROR_LOG_FILE_FULL,
		ERROR_EVENTLOG_FILE_CHANGED,
		ERROR_INVALID_TASK_NAME,
		ERROR_INVALID_TASK_INDEX,
		ERROR_THREAD_ALREADY_IN_TASK,
		ERROR_INSTALL_SERVICE_FAILURE,
		ERROR_INSTALL_USEREXIT,
		ERROR_INSTALL_FAILURE,
		ERROR_INSTALL_SUSPEND,
		ERROR_UNKNOWN_PRODUCT,
		ERROR_UNKNOWN_FEATURE,
		ERROR_UNKNOWN_COMPONENT,
		ERROR_UNKNOWN_PROPERTY,
		ERROR_INVALID_HANDLE_STATE,
		ERROR_BAD_CONFIGURATION,
		ERROR_INDEX_ABSENT,
		ERROR_INSTALL_SOURCE_ABSENT,
		ERROR_INSTALL_PACKAGE_VERSION,
		ERROR_PRODUCT_UNINSTALLED,
		ERROR_BAD_QUERY_SYNTAX,
		ERROR_INVALID_FIELD,
		ERROR_DEVICE_REMOVED,
		ERROR_INSTALL_ALREADY_RUNNING,
		ERROR_INSTALL_PACKAGE_OPEN_FAILED,
		ERROR_INSTALL_PACKAGE_INVALID,
		ERROR_INSTALL_UI_FAILURE,
		ERROR_INSTALL_LOG_FAILURE,
		ERROR_INSTALL_LANGUAGE_UNSUPPORTED,
		ERROR_INSTALL_TRANSFORM_FAILURE,
		ERROR_INSTALL_PACKAGE_REJECTED,
		ERROR_FUNCTION_NOT_CALLED,
		ERROR_FUNCTION_FAILED,
		ERROR_INVALID_TABLE,
		ERROR_DATATYPE_MISMATCH,
		ERROR_UNSUPPORTED_TYPE,
		ERROR_CREATE_FAILED,
		ERROR_INSTALL_TEMP_UNWRITABLE,
		ERROR_INSTALL_PLATFORM_UNSUPPORTED,
		ERROR_INSTALL_NOTUSED,
		ERROR_PATCH_PACKAGE_OPEN_FAILED,
		ERROR_PATCH_PACKAGE_INVALID,
		ERROR_PATCH_PACKAGE_UNSUPPORTED,
		ERROR_PRODUCT_VERSION,
		ERROR_INVALID_COMMAND_LINE,
		ERROR_INSTALL_REMOTE_DISALLOWED,
		ERROR_SUCCESS_REBOOT_INITIATED,
		ERROR_PATCH_TARGET_NOT_FOUND,
		ERROR_PATCH_PACKAGE_REJECTED,
		ERROR_INSTALL_TRANSFORM_REJECTED,
		ERROR_INSTALL_REMOTE_PROHIBITED,
		ERROR_PATCH_REMOVAL_UNSUPPORTED,
		ERROR_UNKNOWN_PATCH,
		ERROR_PATCH_NO_SEQUENCE,
		ERROR_PATCH_REMOVAL_DISALLOWED,
		ERROR_INVALID_PATCH_XML,
		ERROR_PATCH_MANAGED_ADVERTISED_PRODUCT,
		ERROR_INSTALL_SERVICE_SAFEBOOT,
		ERROR_FAIL_FAST_EXCEPTION,
		RPC_S_INVALID_STRING_BINDING,
		RPC_S_WRONG_KIND_OF_BINDING,
		RPC_S_INVALID_BINDING,
		RPC_S_PROTSEQ_NOT_SUPPORTED,
		RPC_S_INVALID_RPC_PROTSEQ,
		RPC_S_INVALID_STRING_UUID,
		RPC_S_INVALID_ENDPOINT_FORMAT,
		RPC_S_INVALID_NET_ADDR,
		RPC_S_NO_ENDPOINT_FOUND,
		RPC_S_INVALID_TIMEOUT,
		RPC_S_OBJECT_NOT_FOUND,
		RPC_S_ALREADY_REGISTERED,
		RPC_S_TYPE_ALREADY_REGISTERED,
		RPC_S_ALREADY_LISTENING,
		RPC_S_NO_PROTSEQS_REGISTERED,
		RPC_S_NOT_LISTENING,
		RPC_S_UNKNOWN_MGR_TYPE,
		RPC_S_UNKNOWN_IF,
		RPC_S_NO_BINDINGS,
		RPC_S_NO_PROTSEQS,
		RPC_S_CANT_CREATE_ENDPOINT,
		RPC_S_OUT_OF_RESOURCES,
		RPC_S_SERVER_UNAVAILABLE,
		RPC_S_SERVER_TOO_BUSY,
		RPC_S_INVALID_NETWORK_OPTIONS,
		RPC_S_NO_CALL_ACTIVE,
		RPC_S_CALL_FAILED,
		RPC_S_CALL_FAILED_DNE,
		RPC_S_PROTOCOL_ERROR,
		RPC_S_PROXY_ACCESS_DENIED,
		RPC_S_UNSUPPORTED_TRANS_SYN,
		RPC_S_UNSUPPORTED_TYPE,
		RPC_S_INVALID_TAG,
		RPC_S_INVALID_BOUND,
		RPC_S_NO_ENTRY_NAME,
		RPC_S_INVALID_NAME_SYNTAX,
		RPC_S_UNSUPPORTED_NAME_SYNTAX,
		RPC_S_UUID_NO_ADDRESS,
		RPC_S_DUPLICATE_ENDPOINT,
		RPC_S_UNKNOWN_AUTHN_TYPE,
		RPC_S_MAX_CALLS_TOO_SMALL,
		RPC_S_STRING_TOO_LONG,
		RPC_S_PROTSEQ_NOT_FOUND,
		RPC_S_PROCNUM_OUT_OF_RANGE,
		RPC_S_BINDING_HAS_NO_AUTH,
		RPC_S_UNKNOWN_AUTHN_SERVICE,
		RPC_S_UNKNOWN_AUTHN_LEVEL,
		RPC_S_INVALID_AUTH_IDENTITY,
		RPC_S_UNKNOWN_AUTHZ_SERVICE,
		EPT_S_INVALID_ENTRY,
		EPT_S_CANT_PERFORM_OP,
		EPT_S_NOT_REGISTERED,
		RPC_S_NOTHING_TO_EXPORT,
		RPC_S_INCOMPLETE_NAME,
		RPC_S_INVALID_VERS_OPTION,
		RPC_S_NO_MORE_MEMBERS,
		RPC_S_NOT_ALL_OBJS_UNEXPORTED,
		RPC_S_INTERFACE_NOT_FOUND,
		RPC_S_ENTRY_ALREADY_EXISTS,
		RPC_S_ENTRY_NOT_FOUND,
		RPC_S_NAME_SERVICE_UNAVAILABLE,
		RPC_S_INVALID_NAF_ID,
		RPC_S_CANNOT_SUPPORT,
		RPC_S_NO_CONTEXT_AVAILABLE,
		RPC_S_INTERNAL_ERROR,
		RPC_S_ZERO_DIVIDE,
		RPC_S_ADDRESS_ERROR,
		RPC_S_FP_DIV_ZERO,
		RPC_S_FP_UNDERFLOW,
		RPC_S_FP_OVERFLOW,
		RPC_X_NO_MORE_ENTRIES,
		RPC_X_SS_CHAR_TRANS_OPEN_FAIL,
		RPC_X_SS_CHAR_TRANS_SHORT_FILE,
		RPC_X_SS_IN_NULL_CONTEXT,
		RPC_X_SS_CONTEXT_DAMAGED,
		RPC_X_SS_HANDLES_MISMATCH,
		RPC_X_SS_CANNOT_GET_CALL_HANDLE,
		RPC_X_NULL_REF_POINTER,
		RPC_X_ENUM_VALUE_OUT_OF_RANGE,
		RPC_X_BYTE_COUNT_TOO_SMALL,
		RPC_X_BAD_STUB_DATA,
		ERROR_INVALID_USER_BUFFER,
		ERROR_UNRECOGNIZED_MEDIA,
		ERROR_NO_TRUST_LSA_SECRET,
		ERROR_NO_TRUST_SAM_ACCOUNT,
		ERROR_TRUSTED_DOMAIN_FAILURE,
		ERROR_TRUSTED_RELATIONSHIP_FAILURE,
		ERROR_TRUST_FAILURE,
		RPC_S_CALL_IN_PROGRESS,
		ERROR_NETLOGON_NOT_STARTED,
		ERROR_ACCOUNT_EXPIRED,
		ERROR_REDIRECTOR_HAS_OPEN_HANDLES,
		ERROR_PRINTER_DRIVER_ALREADY_INSTALLED,
		ERROR_UNKNOWN_PORT,
		ERROR_UNKNOWN_PRINTER_DRIVER,
		ERROR_UNKNOWN_PRINTPROCESSOR,
		ERROR_INVALID_SEPARATOR_FILE,
		ERROR_INVALID_PRIORITY,
		ERROR_INVALID_PRINTER_NAME,
		ERROR_PRINTER_ALREADY_EXISTS,
		ERROR_INVALID_PRINTER_COMMAND,
		ERROR_INVALID_DATATYPE,
		ERROR_INVALID_ENVIRONMENT,
		RPC_S_NO_MORE_BINDINGS,
		ERROR_NOLOGON_INTERDOMAIN_TRUST_ACCOUNT,
		ERROR_NOLOGON_WORKSTATION_TRUST_ACCOUNT,
		ERROR_NOLOGON_SERVER_TRUST_ACCOUNT,
		ERROR_DOMAIN_TRUST_INCONSISTENT,
		ERROR_SERVER_HAS_OPEN_HANDLES,
		ERROR_RESOURCE_DATA_NOT_FOUND,
		ERROR_RESOURCE_TYPE_NOT_FOUND,
		ERROR_RESOURCE_NAME_NOT_FOUND,
		ERROR_RESOURCE_LANG_NOT_FOUND,
		ERROR_NOT_ENOUGH_QUOTA,
		RPC_S_NO_INTERFACES,
		RPC_S_CALL_CANCELLED,
		RPC_S_BINDING_INCOMPLETE,
		RPC_S_COMM_FAILURE,
		RPC_S_UNSUPPORTED_AUTHN_LEVEL,
		RPC_S_NO_PRINC_NAME,
		RPC_S_NOT_RPC_ERROR,
		RPC_S_UUID_LOCAL_ONLY,
		RPC_S_SEC_PKG_ERROR,
		RPC_S_NOT_CANCELLED,
		RPC_X_INVALID_ES_ACTION,
		RPC_X_WRONG_ES_VERSION,
		RPC_X_WRONG_STUB_VERSION,
		RPC_X_INVALID_PIPE_OBJECT,
		RPC_X_WRONG_PIPE_ORDER,
		RPC_X_WRONG_PIPE_VERSION,
		RPC_S_COOKIE_AUTH_FAILED,
		RPC_S_GROUP_MEMBER_NOT_FOUND,
		EPT_S_CANT_CREATE,
		RPC_S_INVALID_OBJECT,
		ERROR_INVALID_TIME,
		ERROR_INVALID_FORM_NAME,
		ERROR_INVALID_FORM_SIZE,
		ERROR_ALREADY_WAITING,
		ERROR_PRINTER_DELETED,
		ERROR_INVALID_PRINTER_STATE,
		ERROR_PASSWORD_MUST_CHANGE,
		ERROR_DOMAIN_CONTROLLER_NOT_FOUND,
		ERROR_ACCOUNT_LOCKED_OUT,
		OR_INVALID_OXID,
		OR_INVALID_OID,
		OR_INVALID_SET,
		RPC_S_SEND_INCOMPLETE,
		RPC_S_INVALID_ASYNC_HANDLE,
		RPC_S_INVALID_ASYNC_CALL,
		RPC_X_PIPE_CLOSED,
		RPC_X_PIPE_DISCIPLINE_ERROR,
		RPC_X_PIPE_EMPTY,
		ERROR_NO_SITENAME,
		ERROR_CANT_ACCESS_FILE,
		ERROR_CANT_RESOLVE_FILENAME,
		RPC_S_ENTRY_TYPE_MISMATCH,
		RPC_S_NOT_ALL_OBJS_EXPORTED,
		RPC_S_INTERFACE_NOT_EXPORTED,
		RPC_S_PROFILE_NOT_ADDED,
		RPC_S_PRF_ELT_NOT_ADDED,
		RPC_S_PRF_ELT_NOT_REMOVED,
		RPC_S_GRP_ELT_NOT_ADDED,
		RPC_S_GRP_ELT_NOT_REMOVED,
		ERROR_KM_DRIVER_BLOCKED,
		ERROR_CONTEXT_EXPIRED,
		ERROR_PER_USER_TRUST_QUOTA_EXCEEDED,
		ERROR_ALL_USER_TRUST_QUOTA_EXCEEDED,
		ERROR_USER_DELETE_TRUST_QUOTA_EXCEEDED,
		ERROR_AUTHENTICATION_FIREWALL_FAILED,
		ERROR_REMOTE_PRINT_CONNECTIONS_BLOCKED,
		ERROR_NTLM_BLOCKED,
		ERROR_INVALID_PIXEL_FORMAT,
		ERROR_BAD_DRIVER,
		ERROR_INVALID_WINDOW_STYLE,
		ERROR_METAFILE_NOT_SUPPORTED,
		ERROR_TRANSFORM_NOT_SUPPORTED,
		ERROR_CLIPPING_NOT_SUPPORTED,
		ERROR_INVALID_CMM,
		ERROR_INVALID_PROFILE,
		ERROR_TAG_NOT_FOUND,
		ERROR_TAG_NOT_PRESENT,
		ERROR_DUPLICATE_TAG,
		ERROR_PROFILE_NOT_ASSOCIATED_WITH_DEVICE,
		ERROR_PROFILE_NOT_FOUND,
		ERROR_INVALID_COLORSPACE,
		ERROR_ICM_NOT_ENABLED,
		ERROR_DELETING_ICM_XFORM,
		ERROR_INVALID_TRANSFORM,
		ERROR_COLORSPACE_MISMATCH,
		ERROR_INVALID_COLORINDEX,
		ERROR_PROFILE_DOES_NOT_MATCH_DEVICE,
		ERROR_CONNECTED_OTHER_PASSWORD,
		ERROR_CONNECTED_OTHER_PASSWORD_DEFAULT,
		ERROR_BAD_USERNAME,
		ERROR_NOT_CONNECTED,
		ERROR_OPEN_FILES,
		ERROR_ACTIVE_CONNECTIONS,
		ERROR_DEVICE_IN_USE,
		ERROR_UNKNOWN_PRINT_MONITOR,
		ERROR_PRINTER_DRIVER_IN_USE,
		ERROR_SPOOL_FILE_NOT_FOUND,
		ERROR_SPL_NO_STARTDOC,
		ERROR_SPL_NO_ADDJOB,
		ERROR_PRINT_PROCESSOR_ALREADY_INSTALLED,
		ERROR_PRINT_MONITOR_ALREADY_INSTALLED,
		ERROR_INVALID_PRINT_MONITOR,
		ERROR_PRINT_MONITOR_IN_USE,
		ERROR_PRINTER_HAS_JOBS_QUEUED,
		ERROR_SUCCESS_REBOOT_REQUIRED,
		ERROR_SUCCESS_RESTART_REQUIRED,
		ERROR_PRINTER_NOT_FOUND,
		ERROR_PRINTER_DRIVER_WARNED,
		ERROR_PRINTER_DRIVER_BLOCKED,
		ERROR_PRINTER_DRIVER_PACKAGE_IN_USE,
		ERROR_CORE_DRIVER_PACKAGE_NOT_FOUND,
		ERROR_FAIL_REBOOT_REQUIRED,
		ERROR_FAIL_REBOOT_INITIATED,
		ERROR_PRINTER_DRIVER_DOWNLOAD_NEEDED,
		ERROR_PRINT_JOB_RESTART_REQUIRED,
		ERROR_IO_REISSUE_AS_CACHED,
		ERROR_WINS_INTERNAL,
		ERROR_CAN_NOT_DEL_LOCAL_WINS,
		ERROR_STATIC_INIT,
		ERROR_INC_BACKUP,
		ERROR_FULL_BACKUP,
		ERROR_REC_NON_EXISTENT,
		ERROR_RPL_NOT_ALLOWED,
		PEERDIST_ERROR_CONTENTINFO_VERSION_UNSUPPORTED,
		PEERDIST_ERROR_CANNOT_PARSE_CONTENTINFO,
		PEERDIST_ERROR_MISSING_DATA,
		PEERDIST_ERROR_NO_MORE,
		PEERDIST_ERROR_NOT_INITIALIZED,
		PEERDIST_ERROR_ALREADY_INITIALIZED,
		PEERDIST_ERROR_SHUTDOWN_IN_PROGRESS,
		PEERDIST_ERROR_INVALIDATED,
		PEERDIST_ERROR_ALREADY_EXISTS,
		PEERDIST_ERROR_OPERATION_NOTFOUND,
		PEERDIST_ERROR_ALREADY_COMPLETED,
		PEERDIST_ERROR_OUT_OF_BOUNDS,
		PEERDIST_ERROR_VERSION_UNSUPPORTED,
		PEERDIST_ERROR_INVALID_CONFIGURATION,
		PEERDIST_ERROR_NOT_LICENSED,
		PEERDIST_ERROR_SERVICE_UNAVAILABLE,
		ERROR_DHCP_ADDRESS_CONFLICT,
		ERROR_WMI_GUID_NOT_FOUND,
		ERROR_WMI_INSTANCE_NOT_FOUND,
		ERROR_WMI_ITEMID_NOT_FOUND,
		ERROR_WMI_TRY_AGAIN,
		ERROR_WMI_DP_NOT_FOUND,
		ERROR_WMI_UNRESOLVED_INSTANCE_REF,
		ERROR_WMI_ALREADY_ENABLED,
		ERROR_WMI_GUID_DISCONNECTED,
		ERROR_WMI_SERVER_UNAVAILABLE,
		ERROR_WMI_DP_FAILED,
		ERROR_WMI_INVALID_MOF,
		ERROR_WMI_INVALID_REGINFO,
		ERROR_WMI_ALREADY_DISABLED,
		ERROR_WMI_READ_ONLY,
		ERROR_WMI_SET_FAILURE,
		ERROR_INVALID_MEDIA,
		ERROR_INVALID_LIBRARY,
		ERROR_INVALID_MEDIA_POOL,
		ERROR_DRIVE_MEDIA_MISMATCH,
		ERROR_MEDIA_OFFLINE,
		ERROR_LIBRARY_OFFLINE,
		ERROR_EMPTY,
		ERROR_NOT_EMPTY,
		ERROR_MEDIA_UNAVAILABLE,
		ERROR_RESOURCE_DISABLED,
		ERROR_INVALID_CLEANER,
		ERROR_UNABLE_TO_CLEAN,
		ERROR_OBJECT_NOT_FOUND,
		ERROR_DATABASE_FAILURE,
		ERROR_DATABASE_FULL,
		ERROR_MEDIA_INCOMPATIBLE,
		ERROR_RESOURCE_NOT_PRESENT,
		ERROR_INVALID_OPERATION,
		ERROR_MEDIA_NOT_AVAILABLE,
		ERROR_DEVICE_NOT_AVAILABLE,
		ERROR_REQUEST_REFUSED,
		ERROR_INVALID_DRIVE_OBJECT,
		ERROR_LIBRARY_FULL,
		ERROR_MEDIUM_NOT_ACCESSIBLE,
		ERROR_UNABLE_TO_LOAD_MEDIUM,
		ERROR_UNABLE_TO_INVENTORY_DRIVE,
		ERROR_UNABLE_TO_INVENTORY_SLOT,
		ERROR_UNABLE_TO_INVENTORY_TRANSPORT,
		ERROR_TRANSPORT_FULL,
		ERROR_CONTROLLING_IEPORT,
		ERROR_UNABLE_TO_EJECT_MOUNTED_MEDIA,
		ERROR_CLEANER_SLOT_SET,
		ERROR_CLEANER_SLOT_NOT_SET,
		ERROR_CLEANER_CARTRIDGE_SPENT,
		ERROR_UNEXPECTED_OMID,
		ERROR_CANT_DELETE_LAST_ITEM,
		ERROR_MESSAGE_EXCEEDS_MAX_SIZE,
		ERROR_VOLUME_CONTAINS_SYS_FILES,
		ERROR_INDIGENOUS_TYPE,
		ERROR_NO_SUPPORTING_DRIVES,
		ERROR_CLEANER_CARTRIDGE_INSTALLED,
		ERROR_IEPORT_FULL,
		ERROR_FILE_OFFLINE,
		ERROR_REMOTE_STORAGE_NOT_ACTIVE,
		ERROR_REMOTE_STORAGE_MEDIA_ERROR,
		ERROR_NOT_A_REPARSE_POINT,
		ERROR_REPARSE_ATTRIBUTE_CONFLICT,
		ERROR_INVALID_REPARSE_DATA,
		ERROR_REPARSE_TAG_INVALID,
		ERROR_REPARSE_TAG_MISMATCH,
		ERROR_VOLUME_NOT_SIS_ENABLED,
		ERROR_DEPENDENT_RESOURCE_EXISTS,
		ERROR_DEPENDENCY_NOT_FOUND,
		ERROR_DEPENDENCY_ALREADY_EXISTS,
		ERROR_RESOURCE_NOT_ONLINE,
		ERROR_HOST_NODE_NOT_AVAILABLE,
		ERROR_RESOURCE_NOT_AVAILABLE,
		ERROR_RESOURCE_NOT_FOUND,
		ERROR_SHUTDOWN_CLUSTER,
		ERROR_CANT_EVICT_ACTIVE_NODE,
		ERROR_OBJECT_ALREADY_EXISTS,
		ERROR_OBJECT_IN_LIST,
		ERROR_GROUP_NOT_AVAILABLE,
		ERROR_GROUP_NOT_FOUND,
		ERROR_GROUP_NOT_ONLINE,
		ERROR_HOST_NODE_NOT_RESOURCE_OWNER,
		ERROR_HOST_NODE_NOT_GROUP_OWNER,
		ERROR_RESMON_CREATE_FAILED,
		ERROR_RESMON_ONLINE_FAILED,
		ERROR_RESOURCE_ONLINE,
		ERROR_QUORUM_RESOURCE,
		ERROR_NOT_QUORUM_CAPABLE,
		ERROR_CLUSTER_SHUTTING_DOWN,
		ERROR_INVALID_STATE,
		ERROR_RESOURCE_PROPERTIES_STORED,
		ERROR_NOT_QUORUM_CLASS,
		ERROR_CORE_RESOURCE,
		ERROR_QUORUM_RESOURCE_ONLINE_FAILED,
		ERROR_QUORUMLOG_OPEN_FAILED,
		ERROR_CLUSTERLOG_CORRUPT,
		ERROR_CLUSTERLOG_RECORD_EXCEEDS_MAXSIZE,
		ERROR_CLUSTERLOG_EXCEEDS_MAXSIZE,
		ERROR_CLUSTERLOG_CHKPOINT_NOT_FOUND,
		ERROR_CLUSTERLOG_NOT_ENOUGH_SPACE,
		ERROR_QUORUM_OWNER_ALIVE,
		ERROR_NETWORK_NOT_AVAILABLE,
		ERROR_NODE_NOT_AVAILABLE,
		ERROR_ALL_NODES_NOT_AVAILABLE,
		ERROR_RESOURCE_FAILED,
		ERROR_CLUSTER_INVALID_NODE,
		ERROR_CLUSTER_NODE_EXISTS,
		ERROR_CLUSTER_JOIN_IN_PROGRESS,
		ERROR_CLUSTER_NODE_NOT_FOUND,
		ERROR_CLUSTER_LOCAL_NODE_NOT_FOUND,
		ERROR_CLUSTER_NETWORK_EXISTS,
		ERROR_CLUSTER_NETWORK_NOT_FOUND,
		ERROR_CLUSTER_NETINTERFACE_EXISTS,
		ERROR_CLUSTER_NETINTERFACE_NOT_FOUND,
		ERROR_CLUSTER_INVALID_REQUEST,
		ERROR_CLUSTER_INVALID_NETWORK_PROVIDER,
		ERROR_CLUSTER_NODE_DOWN,
		ERROR_CLUSTER_NODE_UNREACHABLE,
		ERROR_CLUSTER_NODE_NOT_MEMBER,
		ERROR_CLUSTER_JOIN_NOT_IN_PROGRESS,
		ERROR_CLUSTER_INVALID_NETWORK,
		ERROR_CLUSTER_NODE_UP,
		ERROR_CLUSTER_IPADDR_IN_USE,
		ERROR_CLUSTER_NODE_NOT_PAUSED,
		ERROR_CLUSTER_NO_SECURITY_CONTEXT,
		ERROR_CLUSTER_NETWORK_NOT_INTERNAL,
		ERROR_CLUSTER_NODE_ALREADY_UP,
		ERROR_CLUSTER_NODE_ALREADY_DOWN,
		ERROR_CLUSTER_NETWORK_ALREADY_ONLINE,
		ERROR_CLUSTER_NETWORK_ALREADY_OFFLINE,
		ERROR_CLUSTER_NODE_ALREADY_MEMBER,
		ERROR_CLUSTER_LAST_INTERNAL_NETWORK,
		ERROR_CLUSTER_NETWORK_HAS_DEPENDENTS,
		ERROR_INVALID_OPERATION_ON_QUORUM,
		ERROR_DEPENDENCY_NOT_ALLOWED,
		ERROR_CLUSTER_NODE_PAUSED,
		ERROR_NODE_CANT_HOST_RESOURCE,
		ERROR_CLUSTER_NODE_NOT_READY,
		ERROR_CLUSTER_NODE_SHUTTING_DOWN,
		ERROR_CLUSTER_JOIN_ABORTED,
		ERROR_CLUSTER_INCOMPATIBLE_VERSIONS,
		ERROR_CLUSTER_MAXNUM_OF_RESOURCES_EXCEEDED,
		ERROR_CLUSTER_SYSTEM_CONFIG_CHANGED,
		ERROR_CLUSTER_RESOURCE_TYPE_NOT_FOUND,
		ERROR_CLUSTER_RESTYPE_NOT_SUPPORTED,
		ERROR_CLUSTER_RESNAME_NOT_FOUND,
		ERROR_CLUSTER_NO_RPC_PACKAGES_REGISTERED,
		ERROR_CLUSTER_OWNER_NOT_IN_PREFLIST,
		ERROR_CLUSTER_DATABASE_SEQMISMATCH,
		ERROR_RESMON_INVALID_STATE,
		ERROR_CLUSTER_GUM_NOT_LOCKER,
		ERROR_QUORUM_DISK_NOT_FOUND,
		ERROR_DATABASE_BACKUP_CORRUPT,
		ERROR_CLUSTER_NODE_ALREADY_HAS_DFS_ROOT,
		ERROR_RESOURCE_PROPERTY_UNCHANGEABLE,
		ERROR_CLUSTER_MEMBERSHIP_INVALID_STATE,
		ERROR_CLUSTER_QUORUMLOG_NOT_FOUND,
		ERROR_CLUSTER_MEMBERSHIP_HALT,
		ERROR_CLUSTER_INSTANCE_ID_MISMATCH,
		ERROR_CLUSTER_NETWORK_NOT_FOUND_FOR_IP,
		ERROR_CLUSTER_PROPERTY_DATA_TYPE_MISMATCH,
		ERROR_CLUSTER_EVICT_WITHOUT_CLEANUP,
		ERROR_CLUSTER_PARAMETER_MISMATCH,
		ERROR_NODE_CANNOT_BE_CLUSTERED,
		ERROR_CLUSTER_WRONG_OS_VERSION,
		ERROR_CLUSTER_CANT_CREATE_DUP_CLUSTER_NAME,
		ERROR_CLUSCFG_ALREADY_COMMITTED,
		ERROR_CLUSCFG_ROLLBACK_FAILED,
		ERROR_CLUSCFG_SYSTEM_DISK_DRIVE_LETTER_CONFLICT,
		ERROR_CLUSTER_OLD_VERSION,
		ERROR_CLUSTER_MISMATCHED_COMPUTER_ACCT_NAME,
		ERROR_CLUSTER_NO_NET_ADAPTERS,
		ERROR_CLUSTER_POISONED,
		ERROR_CLUSTER_GROUP_MOVING,
		ERROR_CLUSTER_RESOURCE_TYPE_BUSY,
		ERROR_RESOURCE_CALL_TIMED_OUT,
		ERROR_INVALID_CLUSTER_IPV6_ADDRESS,
		ERROR_CLUSTER_INTERNAL_INVALID_FUNCTION,
		ERROR_CLUSTER_PARAMETER_OUT_OF_BOUNDS,
		ERROR_CLUSTER_PARTIAL_SEND,
		ERROR_CLUSTER_REGISTRY_INVALID_FUNCTION,
		ERROR_CLUSTER_INVALID_STRING_TERMINATION,
		ERROR_CLUSTER_INVALID_STRING_FORMAT,
		ERROR_CLUSTER_DATABASE_TRANSACTION_IN_PROGRESS,
		ERROR_CLUSTER_DATABASE_TRANSACTION_NOT_IN_PROGRESS,
		ERROR_CLUSTER_NULL_DATA,
		ERROR_CLUSTER_PARTIAL_READ,
		ERROR_CLUSTER_PARTIAL_WRITE,
		ERROR_CLUSTER_CANT_DESERIALIZE_DATA,
		ERROR_DEPENDENT_RESOURCE_PROPERTY_CONFLICT,
		ERROR_CLUSTER_NO_QUORUM,
		ERROR_CLUSTER_INVALID_IPV6_NETWORK,
		ERROR_CLUSTER_INVALID_IPV6_TUNNEL_NETWORK,
		ERROR_QUORUM_NOT_ALLOWED_IN_THIS_GROUP,
		ERROR_DEPENDENCY_TREE_TOO_COMPLEX,
		ERROR_EXCEPTION_IN_RESOURCE_CALL,
		ERROR_CLUSTER_RHS_FAILED_INITIALIZATION,
		ERROR_CLUSTER_NOT_INSTALLED,
		ERROR_CLUSTER_RESOURCES_MUST_BE_ONLINE_ON_THE_SAME_NODE,
		ERROR_CLUSTER_MAX_NODES_IN_CLUSTER,
		ERROR_CLUSTER_TOO_MANY_NODES,
		ERROR_CLUSTER_OBJECT_ALREADY_USED,
		ERROR_NONCORE_GROUPS_FOUND,
		ERROR_FILE_SHARE_RESOURCE_CONFLICT,
		ERROR_CLUSTER_EVICT_INVALID_REQUEST,
		ERROR_CLUSTER_SINGLETON_RESOURCE,
		ERROR_CLUSTER_GROUP_SINGLETON_RESOURCE,
		ERROR_CLUSTER_RESOURCE_PROVIDER_FAILED,
		ERROR_CLUSTER_RESOURCE_CONFIGURATION_ERROR,
		ERROR_CLUSTER_GROUP_BUSY,
		ERROR_CLUSTER_NOT_SHARED_VOLUME,
		ERROR_CLUSTER_INVALID_SECURITY_DESCRIPTOR,
		ERROR_CLUSTER_SHARED_VOLUMES_IN_USE,
		ERROR_CLUSTER_USE_SHARED_VOLUMES_API,
		ERROR_CLUSTER_BACKUP_IN_PROGRESS,
		ERROR_NON_CSV_PATH,
		ERROR_CSV_VOLUME_NOT_LOCAL,
		ERROR_CLUSTER_WATCHDOG_TERMINATING,
		ERROR_ENCRYPTION_FAILED,
		ERROR_DECRYPTION_FAILED,
		ERROR_FILE_ENCRYPTED,
		ERROR_NO_RECOVERY_POLICY,
		ERROR_NO_EFS,
		ERROR_WRONG_EFS,
		ERROR_NO_USER_KEYS,
		ERROR_FILE_NOT_ENCRYPTED,
		ERROR_NOT_EXPORT_FORMAT,
		ERROR_FILE_READ_ONLY,
		ERROR_DIR_EFS_DISALLOWED,
		ERROR_EFS_SERVER_NOT_TRUSTED,
		ERROR_BAD_RECOVERY_POLICY,
		ERROR_EFS_ALG_BLOB_TOO_BIG,
		ERROR_VOLUME_NOT_SUPPORT_EFS,
		ERROR_EFS_DISABLED,
		ERROR_EFS_VERSION_NOT_SUPPORT,
		ERROR_CS_ENCRYPTION_INVALID_SERVER_RESPONSE,
		ERROR_CS_ENCRYPTION_UNSUPPORTED_SERVER,
		ERROR_CS_ENCRYPTION_EXISTING_ENCRYPTED_FILE,
		ERROR_CS_ENCRYPTION_NEW_ENCRYPTED_FILE,
		ERROR_CS_ENCRYPTION_FILE_NOT_CSE,
		ERROR_ENCRYPTION_POLICY_DENIES_OPERATION,
		ERROR_NO_BROWSER_SERVERS_FOUND,
		SCHED_E_SERVICE_NOT_LOCALSYSTEM,
		ERROR_LOG_SECTOR_INVALID,
		ERROR_LOG_SECTOR_PARITY_INVALID,
		ERROR_LOG_SECTOR_REMAPPED,
		ERROR_LOG_BLOCK_INCOMPLETE,
		ERROR_LOG_INVALID_RANGE,
		ERROR_LOG_BLOCKS_EXHAUSTED,
		ERROR_LOG_READ_CONTEXT_INVALID,
		ERROR_LOG_RESTART_INVALID,
		ERROR_LOG_BLOCK_VERSION,
		ERROR_LOG_BLOCK_INVALID,
		ERROR_LOG_READ_MODE_INVALID,
		ERROR_LOG_NO_RESTART,
		ERROR_LOG_METADATA_CORRUPT,
		ERROR_LOG_METADATA_INVALID,
		ERROR_LOG_METADATA_INCONSISTENT,
		ERROR_LOG_RESERVATION_INVALID,
		ERROR_LOG_CANT_DELETE,
		ERROR_LOG_CONTAINER_LIMIT_EXCEEDED,
		ERROR_LOG_START_OF_LOG,
		ERROR_LOG_POLICY_ALREADY_INSTALLED,
		ERROR_LOG_POLICY_NOT_INSTALLED,
		ERROR_LOG_POLICY_INVALID,
		ERROR_LOG_POLICY_CONFLICT,
		ERROR_LOG_PINNED_ARCHIVE_TAIL,
		ERROR_LOG_RECORD_NONEXISTENT,
		ERROR_LOG_RECORDS_RESERVED_INVALID,
		ERROR_LOG_SPACE_RESERVED_INVALID,
		ERROR_LOG_TAIL_INVALID,
		ERROR_LOG_FULL,
		ERROR_COULD_NOT_RESIZE_LOG,
		ERROR_LOG_MULTIPLEXED,
		ERROR_LOG_DEDICATED,
		ERROR_LOG_ARCHIVE_NOT_IN_PROGRESS,
		ERROR_LOG_ARCHIVE_IN_PROGRESS,
		ERROR_LOG_EPHEMERAL,
		ERROR_LOG_NOT_ENOUGH_CONTAINERS,
		ERROR_LOG_CLIENT_ALREADY_REGISTERED,
		ERROR_LOG_CLIENT_NOT_REGISTERED,
		ERROR_LOG_FULL_HANDLER_IN_PROGRESS,
		ERROR_LOG_CONTAINER_READ_FAILED,
		ERROR_LOG_CONTAINER_WRITE_FAILED,
		ERROR_LOG_CONTAINER_OPEN_FAILED,
		ERROR_LOG_CONTAINER_STATE_INVALID,
		ERROR_LOG_STATE_INVALID,
		ERROR_LOG_PINNED,
		ERROR_LOG_METADATA_FLUSH_FAILED,
		ERROR_LOG_INCONSISTENT_SECURITY,
		ERROR_LOG_APPENDED_FLUSH_FAILED,
		ERROR_LOG_PINNED_RESERVATION,
		ERROR_INVALID_TRANSACTION,
		ERROR_TRANSACTION_NOT_ACTIVE,
		ERROR_TRANSACTION_REQUEST_NOT_VALID,
		ERROR_TRANSACTION_NOT_REQUESTED,
		ERROR_TRANSACTION_ALREADY_ABORTED,
		ERROR_TRANSACTION_ALREADY_COMMITTED,
		ERROR_TM_INITIALIZATION_FAILED,
		ERROR_RESOURCEMANAGER_READ_ONLY,
		ERROR_TRANSACTION_NOT_JOINED,
		ERROR_TRANSACTION_SUPERIOR_EXISTS,
		ERROR_CRM_PROTOCOL_ALREADY_EXISTS,
		ERROR_TRANSACTION_PROPAGATION_FAILED,
		ERROR_CRM_PROTOCOL_NOT_FOUND,
		ERROR_TRANSACTION_INVALID_MARSHALL_BUFFER,
		ERROR_CURRENT_TRANSACTION_NOT_VALID,
		ERROR_TRANSACTION_NOT_FOUND,
		ERROR_RESOURCEMANAGER_NOT_FOUND,
		ERROR_ENLISTMENT_NOT_FOUND,
		ERROR_TRANSACTIONMANAGER_NOT_FOUND,
		ERROR_TRANSACTIONMANAGER_NOT_ONLINE,
		ERROR_TRANSACTIONMANAGER_RECOVERY_NAME_COLLISION,
		ERROR_TRANSACTION_NOT_ROOT,
		ERROR_TRANSACTION_OBJECT_EXPIRED,
		ERROR_TRANSACTION_RESPONSE_NOT_ENLISTED,
		ERROR_TRANSACTION_RECORD_TOO_LONG,
		ERROR_IMPLICIT_TRANSACTION_NOT_SUPPORTED,
		ERROR_TRANSACTION_INTEGRITY_VIOLATED,
		ERROR_TRANSACTIONMANAGER_IDENTITY_MISMATCH,
		ERROR_RM_CANNOT_BE_FROZEN_FOR_SNAPSHOT,
		ERROR_TRANSACTION_MUST_WRITETHROUGH,
		ERROR_TRANSACTION_NO_SUPERIOR,
		ERROR_HEURISTIC_DAMAGE_POSSIBLE,
		ERROR_TRANSACTIONAL_CONFLICT,
		ERROR_RM_NOT_ACTIVE,
		ERROR_RM_METADATA_CORRUPT,
		ERROR_DIRECTORY_NOT_RM,
		ERROR_TRANSACTIONS_UNSUPPORTED_REMOTE,
		ERROR_LOG_RESIZE_INVALID_SIZE,
		ERROR_OBJECT_NO_LONGER_EXISTS,
		ERROR_STREAM_MINIVERSION_NOT_FOUND,
		ERROR_STREAM_MINIVERSION_NOT_VALID,
		ERROR_MINIVERSION_INACCESSIBLE_FROM_SPECIFIED_TRANSACTION,
		ERROR_CANT_OPEN_MINIVERSION_WITH_MODIFY_INTENT,
		ERROR_CANT_CREATE_MORE_STREAM_MINIVERSIONS,
		ERROR_REMOTE_FILE_VERSION_MISMATCH,
		ERROR_HANDLE_NO_LONGER_VALID,
		ERROR_NO_TXF_METADATA,
		ERROR_LOG_CORRUPTION_DETECTED,
		ERROR_CANT_RECOVER_WITH_HANDLE_OPEN,
		ERROR_RM_DISCONNECTED,
		ERROR_ENLISTMENT_NOT_SUPERIOR,
		ERROR_RECOVERY_NOT_NEEDED,
		ERROR_RM_ALREADY_STARTED,
		ERROR_FILE_IDENTITY_NOT_PERSISTENT,
		ERROR_CANT_BREAK_TRANSACTIONAL_DEPENDENCY,
		ERROR_CANT_CROSS_RM_BOUNDARY,
		ERROR_TXF_DIR_NOT_EMPTY,
		ERROR_INDOUBT_TRANSACTIONS_EXIST,
		ERROR_TM_VOLATILE,
		ERROR_ROLLBACK_TIMER_EXPIRED,
		ERROR_TXF_ATTRIBUTE_CORRUPT,
		ERROR_EFS_NOT_ALLOWED_IN_TRANSACTION,
		ERROR_TRANSACTIONAL_OPEN_NOT_ALLOWED,
		ERROR_LOG_GROWTH_FAILED,
		ERROR_TRANSACTED_MAPPING_UNSUPPORTED_REMOTE,
		ERROR_TXF_METADATA_ALREADY_PRESENT,
		ERROR_TRANSACTION_SCOPE_CALLBACKS_NOT_SET,
		ERROR_TRANSACTION_REQUIRED_PROMOTION,
		ERROR_CANNOT_EXECUTE_FILE_IN_TRANSACTION,
		ERROR_TRANSACTIONS_NOT_FROZEN,
		ERROR_TRANSACTION_FREEZE_IN_PROGRESS,
		ERROR_NOT_SNAPSHOT_VOLUME,
		ERROR_NO_SAVEPOINT_WITH_OPEN_FILES,
		ERROR_DATA_LOST_REPAIR,
		ERROR_SPARSE_NOT_ALLOWED_IN_TRANSACTION,
		ERROR_TM_IDENTITY_MISMATCH,
		ERROR_FLOATED_SECTION,
		ERROR_CANNOT_ACCEPT_TRANSACTED_WORK,
		ERROR_CANNOT_ABORT_TRANSACTIONS,
		ERROR_BAD_CLUSTERS,
		ERROR_COMPRESSION_NOT_ALLOWED_IN_TRANSACTION,
		ERROR_VOLUME_DIRTY,
		ERROR_NO_LINK_TRACKING_IN_TRANSACTION,
		ERROR_OPERATION_NOT_SUPPORTED_IN_TRANSACTION,
		ERROR_EXPIRED_HANDLE,
		ERROR_TRANSACTION_NOT_ENLISTED,
		ERROR_CTX_WINSTATION_NAME_INVALID,
		ERROR_CTX_INVALID_PD,
		ERROR_CTX_PD_NOT_FOUND,
		ERROR_CTX_WD_NOT_FOUND,
		ERROR_CTX_CANNOT_MAKE_EVENTLOG_ENTRY,
		ERROR_CTX_SERVICE_NAME_COLLISION,
		ERROR_CTX_CLOSE_PENDING,
		ERROR_CTX_NO_OUTBUF,
		ERROR_CTX_MODEM_INF_NOT_FOUND,
		ERROR_CTX_INVALID_MODEMNAME,
		ERROR_CTX_MODEM_RESPONSE_ERROR,
		ERROR_CTX_MODEM_RESPONSE_TIMEOUT,
		ERROR_CTX_MODEM_RESPONSE_NO_CARRIER,
		ERROR_CTX_MODEM_RESPONSE_NO_DIALTONE,
		ERROR_CTX_MODEM_RESPONSE_BUSY,
		ERROR_CTX_MODEM_RESPONSE_VOICE,
		ERROR_CTX_TD_ERROR,
		ERROR_CTX_WINSTATION_NOT_FOUND,
		ERROR_CTX_WINSTATION_ALREADY_EXISTS,
		ERROR_CTX_WINSTATION_BUSY,
		ERROR_CTX_BAD_VIDEO_MODE,
		ERROR_CTX_GRAPHICS_INVALID,
		ERROR_CTX_LOGON_DISABLED,
		ERROR_CTX_NOT_CONSOLE,
		ERROR_CTX_CLIENT_QUERY_TIMEOUT,
		ERROR_CTX_CONSOLE_DISCONNECT,
		ERROR_CTX_CONSOLE_CONNECT,
		ERROR_CTX_SHADOW_DENIED,
		ERROR_CTX_WINSTATION_ACCESS_DENIED,
		ERROR_CTX_INVALID_WD,
		ERROR_CTX_SHADOW_INVALID,
		ERROR_CTX_SHADOW_DISABLED,
		ERROR_CTX_CLIENT_LICENSE_IN_USE,
		ERROR_CTX_CLIENT_LICENSE_NOT_SET,
		ERROR_CTX_LICENSE_NOT_AVAILABLE,
		ERROR_CTX_LICENSE_CLIENT_INVALID,
		ERROR_CTX_LICENSE_EXPIRED,
		ERROR_CTX_SHADOW_NOT_RUNNING,
		ERROR_CTX_SHADOW_ENDED_BY_MODE_CHANGE,
		ERROR_ACTIVATION_COUNT_EXCEEDED,
		ERROR_CTX_WINSTATIONS_DISABLED,
		ERROR_CTX_ENCRYPTION_LEVEL_REQUIRED,
		ERROR_CTX_SESSION_IN_USE,
		ERROR_CTX_NO_FORCE_LOGOFF,
		ERROR_CTX_ACCOUNT_RESTRICTION,
		ERROR_RDP_PROTOCOL_ERROR,
		ERROR_CTX_CDM_CONNECT,
		ERROR_CTX_CDM_DISCONNECT,
		ERROR_CTX_SECURITY_LAYER_ERROR,
		ERROR_TS_INCOMPATIBLE_SESSIONS,
		ERROR_TS_VIDEO_SUBSYSTEM_ERROR,
		FRS_ERR_INVALID_API_SEQUENCE,
		FRS_ERR_STARTING_SERVICE,
		FRS_ERR_STOPPING_SERVICE,
		FRS_ERR_INTERNAL_API,
		FRS_ERR_INTERNAL,
		FRS_ERR_SERVICE_COMM,
		FRS_ERR_INSUFFICIENT_PRIV,
		FRS_ERR_AUTHENTICATION,
		FRS_ERR_PARENT_INSUFFICIENT_PRIV,
		FRS_ERR_PARENT_AUTHENTICATION,
		FRS_ERR_CHILD_TO_PARENT_COMM,
		FRS_ERR_PARENT_TO_CHILD_COMM,
		FRS_ERR_SYSVOL_POPULATE,
		FRS_ERR_SYSVOL_POPULATE_TIMEOUT,
		FRS_ERR_SYSVOL_IS_BUSY,
		FRS_ERR_SYSVOL_DEMOTE,
		FRS_ERR_INVALID_SERVICE_PARAMETER,
		ERROR_DS_NOT_INSTALLED,
		ERROR_DS_MEMBERSHIP_EVALUATED_LOCALLY,
		ERROR_DS_NO_ATTRIBUTE_OR_VALUE,
		ERROR_DS_INVALID_ATTRIBUTE_SYNTAX,
		ERROR_DS_ATTRIBUTE_TYPE_UNDEFINED,
		ERROR_DS_ATTRIBUTE_OR_VALUE_EXISTS,
		ERROR_DS_BUSY,
		ERROR_DS_UNAVAILABLE,
		ERROR_DS_NO_RIDS_ALLOCATED,
		ERROR_DS_NO_MORE_RIDS,
		ERROR_DS_INCORRECT_ROLE_OWNER,
		ERROR_DS_RIDMGR_INIT_ERROR,
		ERROR_DS_OBJ_CLASS_VIOLATION,
		ERROR_DS_CANT_ON_NON_LEAF,
		ERROR_DS_CANT_ON_RDN,
		ERROR_DS_CANT_MOD_OBJ_CLASS,
		ERROR_DS_CROSS_DOM_MOVE_ERROR,
		ERROR_DS_GC_NOT_AVAILABLE,
		ERROR_SHARED_POLICY,
		ERROR_POLICY_OBJECT_NOT_FOUND,
		ERROR_POLICY_ONLY_IN_DS,
		ERROR_PROMOTION_ACTIVE,
		ERROR_NO_PROMOTION_ACTIVE,
		ERROR_DS_OPERATIONS_ERROR,
		ERROR_DS_PROTOCOL_ERROR,
		ERROR_DS_TIMELIMIT_EXCEEDED,
		ERROR_DS_SIZELIMIT_EXCEEDED,
		ERROR_DS_ADMIN_LIMIT_EXCEEDED,
		ERROR_DS_COMPARE_FALSE,
		ERROR_DS_COMPARE_TRUE,
		ERROR_DS_AUTH_METHOD_NOT_SUPPORTED,
		ERROR_DS_STRONG_AUTH_REQUIRED,
		ERROR_DS_INAPPROPRIATE_AUTH,
		ERROR_DS_AUTH_UNKNOWN,
		ERROR_DS_REFERRAL,
		ERROR_DS_UNAVAILABLE_CRIT_EXTENSION,
		ERROR_DS_CONFIDENTIALITY_REQUIRED,
		ERROR_DS_INAPPROPRIATE_MATCHING,
		ERROR_DS_CONSTRAINT_VIOLATION,
		ERROR_DS_NO_SUCH_OBJECT,
		ERROR_DS_ALIAS_PROBLEM,
		ERROR_DS_INVALID_DN_SYNTAX,
		ERROR_DS_IS_LEAF,
		ERROR_DS_ALIAS_DEREF_PROBLEM,
		ERROR_DS_UNWILLING_TO_PERFORM,
		ERROR_DS_LOOP_DETECT,
		ERROR_DS_NAMING_VIOLATION,
		ERROR_DS_OBJECT_RESULTS_TOO_LARGE,
		ERROR_DS_AFFECTS_MULTIPLE_DSAS,
		ERROR_DS_SERVER_DOWN,
		ERROR_DS_LOCAL_ERROR,
		ERROR_DS_ENCODING_ERROR,
		ERROR_DS_DECODING_ERROR,
		ERROR_DS_FILTER_UNKNOWN,
		ERROR_DS_PARAM_ERROR,
		ERROR_DS_NOT_SUPPORTED,
		ERROR_DS_NO_RESULTS_RETURNED,
		ERROR_DS_CONTROL_NOT_FOUND,
		ERROR_DS_CLIENT_LOOP,
		ERROR_DS_REFERRAL_LIMIT_EXCEEDED,
		ERROR_DS_SORT_CONTROL_MISSING,
		ERROR_DS_OFFSET_RANGE_ERROR,
		ERROR_DS_ROOT_MUST_BE_NC,
		ERROR_DS_ADD_REPLICA_INHIBITED,
		ERROR_DS_ATT_NOT_DEF_IN_SCHEMA,
		ERROR_DS_MAX_OBJ_SIZE_EXCEEDED,
		ERROR_DS_OBJ_STRING_NAME_EXISTS,
		ERROR_DS_NO_RDN_DEFINED_IN_SCHEMA,
		ERROR_DS_RDN_DOESNT_MATCH_SCHEMA,
		ERROR_DS_NO_REQUESTED_ATTS_FOUND,
		ERROR_DS_USER_BUFFER_TO_SMALL,
		ERROR_DS_ATT_IS_NOT_ON_OBJ,
		ERROR_DS_ILLEGAL_MOD_OPERATION,
		ERROR_DS_OBJ_TOO_LARGE,
		ERROR_DS_BAD_INSTANCE_TYPE,
		ERROR_DS_MASTERDSA_REQUIRED,
		ERROR_DS_OBJECT_CLASS_REQUIRED,
		ERROR_DS_MISSING_REQUIRED_ATT,
		ERROR_DS_ATT_NOT_DEF_FOR_CLASS,
		ERROR_DS_ATT_ALREADY_EXISTS,
		ERROR_DS_CANT_ADD_ATT_VALUES,
		ERROR_DS_SINGLE_VALUE_CONSTRAINT,
		ERROR_DS_RANGE_CONSTRAINT,
		ERROR_DS_ATT_VAL_ALREADY_EXISTS,
		ERROR_DS_CANT_REM_MISSING_ATT,
		ERROR_DS_CANT_REM_MISSING_ATT_VAL,
		ERROR_DS_ROOT_CANT_BE_SUBREF,
		ERROR_DS_NO_CHAINING,
		ERROR_DS_NO_CHAINED_EVAL,
		ERROR_DS_NO_PARENT_OBJECT,
		ERROR_DS_PARENT_IS_AN_ALIAS,
		ERROR_DS_CANT_MIX_MASTER_AND_REPS,
		ERROR_DS_CHILDREN_EXIST,
		ERROR_DS_OBJ_NOT_FOUND,
		ERROR_DS_ALIASED_OBJ_MISSING,
		ERROR_DS_BAD_NAME_SYNTAX,
		ERROR_DS_ALIAS_POINTS_TO_ALIAS,
		ERROR_DS_CANT_DEREF_ALIAS,
		ERROR_DS_OUT_OF_SCOPE,
		ERROR_DS_OBJECT_BEING_REMOVED,
		ERROR_DS_CANT_DELETE_DSA_OBJ,
		ERROR_DS_GENERIC_ERROR,
		ERROR_DS_DSA_MUST_BE_INT_MASTER,
		ERROR_DS_CLASS_NOT_DSA,
		ERROR_DS_INSUFF_ACCESS_RIGHTS,
		ERROR_DS_ILLEGAL_SUPERIOR,
		ERROR_DS_ATTRIBUTE_OWNED_BY_SAM,
		ERROR_DS_NAME_TOO_MANY_PARTS,
		ERROR_DS_NAME_TOO_LONG,
		ERROR_DS_NAME_VALUE_TOO_LONG,
		ERROR_DS_NAME_UNPARSEABLE,
		ERROR_DS_NAME_TYPE_UNKNOWN,
		ERROR_DS_NOT_AN_OBJECT,
		ERROR_DS_SEC_DESC_TOO_SHORT,
		ERROR_DS_SEC_DESC_INVALID,
		ERROR_DS_NO_DELETED_NAME,
		ERROR_DS_SUBREF_MUST_HAVE_PARENT,
		ERROR_DS_NCNAME_MUST_BE_NC,
		ERROR_DS_CANT_ADD_SYSTEM_ONLY,
		ERROR_DS_CLASS_MUST_BE_CONCRETE,
		ERROR_DS_INVALID_DMD,
		ERROR_DS_OBJ_GUID_EXISTS,
		ERROR_DS_NOT_ON_BACKLINK,
		ERROR_DS_NO_CROSSREF_FOR_NC,
		ERROR_DS_SHUTTING_DOWN,
		ERROR_DS_UNKNOWN_OPERATION,
		ERROR_DS_INVALID_ROLE_OWNER,
		ERROR_DS_COULDNT_CONTACT_FSMO,
		ERROR_DS_CROSS_NC_DN_RENAME,
		ERROR_DS_CANT_MOD_SYSTEM_ONLY,
		ERROR_DS_REPLICATOR_ONLY,
		ERROR_DS_OBJ_CLASS_NOT_DEFINED,
		ERROR_DS_OBJ_CLASS_NOT_SUBCLASS,
		ERROR_DS_NAME_REFERENCE_INVALID,
		ERROR_DS_CROSS_REF_EXISTS,
		ERROR_DS_CANT_DEL_MASTER_CROSSREF,
		ERROR_DS_SUBTREE_NOTIFY_NOT_NC_HEAD,
		ERROR_DS_NOTIFY_FILTER_TOO_COMPLEX,
		ERROR_DS_DUP_RDN,
		ERROR_DS_DUP_OID,
		ERROR_DS_DUP_MAPI_ID,
		ERROR_DS_DUP_SCHEMA_ID_GUID,
		ERROR_DS_DUP_LDAP_DISPLAY_NAME,
		ERROR_DS_SEMANTIC_ATT_TEST,
		ERROR_DS_SYNTAX_MISMATCH,
		ERROR_DS_EXISTS_IN_MUST_HAVE,
		ERROR_DS_EXISTS_IN_MAY_HAVE,
		ERROR_DS_NONEXISTENT_MAY_HAVE,
		ERROR_DS_NONEXISTENT_MUST_HAVE,
		ERROR_DS_AUX_CLS_TEST_FAIL,
		ERROR_DS_NONEXISTENT_POSS_SUP,
		ERROR_DS_SUB_CLS_TEST_FAIL,
		ERROR_DS_BAD_RDN_ATT_ID_SYNTAX,
		ERROR_DS_EXISTS_IN_AUX_CLS,
		ERROR_DS_EXISTS_IN_SUB_CLS,
		ERROR_DS_EXISTS_IN_POSS_SUP,
		ERROR_DS_RECALCSCHEMA_FAILED,
		ERROR_DS_TREE_DELETE_NOT_FINISHED,
		ERROR_DS_CANT_DELETE,
		ERROR_DS_ATT_SCHEMA_REQ_ID,
		ERROR_DS_BAD_ATT_SCHEMA_SYNTAX,
		ERROR_DS_CANT_CACHE_ATT,
		ERROR_DS_CANT_CACHE_CLASS,
		ERROR_DS_CANT_REMOVE_ATT_CACHE,
		ERROR_DS_CANT_REMOVE_CLASS_CACHE,
		ERROR_DS_CANT_RETRIEVE_DN,
		ERROR_DS_MISSING_SUPREF,
		ERROR_DS_CANT_RETRIEVE_INSTANCE,
		ERROR_DS_CODE_INCONSISTENCY,
		ERROR_DS_DATABASE_ERROR,
		ERROR_DS_GOVERNSID_MISSING,
		ERROR_DS_MISSING_EXPECTED_ATT,
		ERROR_DS_NCNAME_MISSING_CR_REF,
		ERROR_DS_SECURITY_CHECKING_ERROR,
		ERROR_DS_SCHEMA_NOT_LOADED,
		ERROR_DS_SCHEMA_ALLOC_FAILED,
		ERROR_DS_ATT_SCHEMA_REQ_SYNTAX,
		ERROR_DS_GCVERIFY_ERROR,
		ERROR_DS_DRA_SCHEMA_MISMATCH,
		ERROR_DS_CANT_FIND_DSA_OBJ,
		ERROR_DS_CANT_FIND_EXPECTED_NC,
		ERROR_DS_CANT_FIND_NC_IN_CACHE,
		ERROR_DS_CANT_RETRIEVE_CHILD,
		ERROR_DS_SECURITY_ILLEGAL_MODIFY,
		ERROR_DS_CANT_REPLACE_HIDDEN_REC,
		ERROR_DS_BAD_HIERARCHY_FILE,
		ERROR_DS_BUILD_HIERARCHY_TABLE_FAILED,
		ERROR_DS_CONFIG_PARAM_MISSING,
		ERROR_DS_COUNTING_AB_INDICES_FAILED,
		ERROR_DS_HIERARCHY_TABLE_MALLOC_FAILED,
		ERROR_DS_INTERNAL_FAILURE,
		ERROR_DS_UNKNOWN_ERROR,
		ERROR_DS_ROOT_REQUIRES_CLASS_TOP,
		ERROR_DS_REFUSING_FSMO_ROLES,
		ERROR_DS_MISSING_FSMO_SETTINGS,
		ERROR_DS_UNABLE_TO_SURRENDER_ROLES,
		ERROR_DS_DRA_GENERIC,
		ERROR_DS_DRA_INVALID_PARAMETER,
		ERROR_DS_DRA_BUSY,
		ERROR_DS_DRA_BAD_DN,
		ERROR_DS_DRA_BAD_NC,
		ERROR_DS_DRA_DN_EXISTS,
		ERROR_DS_DRA_INTERNAL_ERROR,
		ERROR_DS_DRA_INCONSISTENT_DIT,
		ERROR_DS_DRA_CONNECTION_FAILED,
		ERROR_DS_DRA_BAD_INSTANCE_TYPE,
		ERROR_DS_DRA_OUT_OF_MEM,
		ERROR_DS_DRA_MAIL_PROBLEM,
		ERROR_DS_DRA_REF_ALREADY_EXISTS,
		ERROR_DS_DRA_REF_NOT_FOUND,
		ERROR_DS_DRA_OBJ_IS_REP_SOURCE,
		ERROR_DS_DRA_DB_ERROR,
		ERROR_DS_DRA_NO_REPLICA,
		ERROR_DS_DRA_ACCESS_DENIED,
		ERROR_DS_DRA_NOT_SUPPORTED,
		ERROR_DS_DRA_RPC_CANCELLED,
		ERROR_DS_DRA_SOURCE_DISABLED,
		ERROR_DS_DRA_SINK_DISABLED,
		ERROR_DS_DRA_NAME_COLLISION,
		ERROR_DS_DRA_SOURCE_REINSTALLED,
		ERROR_DS_DRA_MISSING_PARENT,
		ERROR_DS_DRA_PREEMPTED,
		ERROR_DS_DRA_ABANDON_SYNC,
		ERROR_DS_DRA_SHUTDOWN,
		ERROR_DS_DRA_INCOMPATIBLE_PARTIAL_SET,
		ERROR_DS_DRA_SOURCE_IS_PARTIAL_REPLICA,
		ERROR_DS_DRA_EXTN_CONNECTION_FAILED,
		ERROR_DS_INSTALL_SCHEMA_MISMATCH,
		ERROR_DS_DUP_LINK_ID,
		ERROR_DS_NAME_ERROR_RESOLVING,
		ERROR_DS_NAME_ERROR_NOT_FOUND,
		ERROR_DS_NAME_ERROR_NOT_UNIQUE,
		ERROR_DS_NAME_ERROR_NO_MAPPING,
		ERROR_DS_NAME_ERROR_DOMAIN_ONLY,
		ERROR_DS_NAME_ERROR_NO_SYNTACTICAL_MAPPING,
		ERROR_DS_CONSTRUCTED_ATT_MOD,
		ERROR_DS_WRONG_OM_OBJ_CLASS,
		ERROR_DS_DRA_REPL_PENDING,
		ERROR_DS_DS_REQUIRED,
		ERROR_DS_INVALID_LDAP_DISPLAY_NAME,
		ERROR_DS_NON_BASE_SEARCH,
		ERROR_DS_CANT_RETRIEVE_ATTS,
		ERROR_DS_BACKLINK_WITHOUT_LINK,
		ERROR_DS_EPOCH_MISMATCH,
		ERROR_DS_SRC_NAME_MISMATCH,
		ERROR_DS_SRC_AND_DST_NC_IDENTICAL,
		ERROR_DS_DST_NC_MISMATCH,
		ERROR_DS_NOT_AUTHORITIVE_FOR_DST_NC,
		ERROR_DS_SRC_GUID_MISMATCH,
		ERROR_DS_CANT_MOVE_DELETED_OBJECT,
		ERROR_DS_PDC_OPERATION_IN_PROGRESS,
		ERROR_DS_CROSS_DOMAIN_CLEANUP_REQD,
		ERROR_DS_ILLEGAL_XDOM_MOVE_OPERATION,
		ERROR_DS_CANT_WITH_ACCT_GROUP_MEMBERSHPS,
		ERROR_DS_NC_MUST_HAVE_NC_PARENT,
		ERROR_DS_CR_IMPOSSIBLE_TO_VALIDATE,
		ERROR_DS_DST_DOMAIN_NOT_NATIVE,
		ERROR_DS_MISSING_INFRASTRUCTURE_CONTAINER,
		ERROR_DS_CANT_MOVE_ACCOUNT_GROUP,
		ERROR_DS_CANT_MOVE_RESOURCE_GROUP,
		ERROR_DS_INVALID_SEARCH_FLAG,
		ERROR_DS_NO_TREE_DELETE_ABOVE_NC,
		ERROR_DS_COULDNT_LOCK_TREE_FOR_DELETE,
		ERROR_DS_COULDNT_IDENTIFY_OBJECTS_FOR_TREE_DELETE,
		ERROR_DS_SAM_INIT_FAILURE,
		ERROR_DS_SENSITIVE_GROUP_VIOLATION,
		ERROR_DS_CANT_MOD_PRIMARYGROUPID,
		ERROR_DS_ILLEGAL_BASE_SCHEMA_MOD,
		ERROR_DS_NONSAFE_SCHEMA_CHANGE,
		ERROR_DS_SCHEMA_UPDATE_DISALLOWED,
		ERROR_DS_CANT_CREATE_UNDER_SCHEMA,
		ERROR_DS_INSTALL_NO_SRC_SCH_VERSION,
		ERROR_DS_INSTALL_NO_SCH_VERSION_IN_INIFILE,
		ERROR_DS_INVALID_GROUP_TYPE,
		ERROR_DS_NO_NEST_GLOBALGROUP_IN_MIXEDDOMAIN,
		ERROR_DS_NO_NEST_LOCALGROUP_IN_MIXEDDOMAIN,
		ERROR_DS_GLOBAL_CANT_HAVE_LOCAL_MEMBER,
		ERROR_DS_GLOBAL_CANT_HAVE_UNIVERSAL_MEMBER,
		ERROR_DS_UNIVERSAL_CANT_HAVE_LOCAL_MEMBER,
		ERROR_DS_GLOBAL_CANT_HAVE_CROSSDOMAIN_MEMBER,
		ERROR_DS_LOCAL_CANT_HAVE_CROSSDOMAIN_LOCAL_MEMBER,
		ERROR_DS_HAVE_PRIMARY_MEMBERS,
		ERROR_DS_STRING_SD_CONVERSION_FAILED,
		ERROR_DS_NAMING_MASTER_GC,
		ERROR_DS_DNS_LOOKUP_FAILURE,
		ERROR_DS_COULDNT_UPDATE_SPNS,
		ERROR_DS_CANT_RETRIEVE_SD,
		ERROR_DS_KEY_NOT_UNIQUE,
		ERROR_DS_WRONG_LINKED_ATT_SYNTAX,
		ERROR_DS_SAM_NEED_BOOTKEY_PASSWORD,
		ERROR_DS_SAM_NEED_BOOTKEY_FLOPPY,
		ERROR_DS_CANT_START,
		ERROR_DS_INIT_FAILURE,
		ERROR_DS_NO_PKT_PRIVACY_ON_CONNECTION,
		ERROR_DS_SOURCE_DOMAIN_IN_FOREST,
		ERROR_DS_DESTINATION_DOMAIN_NOT_IN_FOREST,
		ERROR_DS_DESTINATION_AUDITING_NOT_ENABLED,
		ERROR_DS_CANT_FIND_DC_FOR_SRC_DOMAIN,
		ERROR_DS_SRC_OBJ_NOT_GROUP_OR_USER,
		ERROR_DS_SRC_SID_EXISTS_IN_FOREST,
		ERROR_DS_SRC_AND_DST_OBJECT_CLASS_MISMATCH,
		ERROR_SAM_INIT_FAILURE,
		ERROR_DS_DRA_SCHEMA_INFO_SHIP,
		ERROR_DS_DRA_SCHEMA_CONFLICT,
		ERROR_DS_DRA_EARLIER_SCHEMA_CONFLICT,
		ERROR_DS_DRA_OBJ_NC_MISMATCH,
		ERROR_DS_NC_STILL_HAS_DSAS,
		ERROR_DS_GC_REQUIRED,
		ERROR_DS_LOCAL_MEMBER_OF_LOCAL_ONLY,
		ERROR_DS_NO_FPO_IN_UNIVERSAL_GROUPS,
		ERROR_DS_CANT_ADD_TO_GC,
		ERROR_DS_NO_CHECKPOINT_WITH_PDC,
		ERROR_DS_SOURCE_AUDITING_NOT_ENABLED,
		ERROR_DS_CANT_CREATE_IN_NONDOMAIN_NC,
		ERROR_DS_INVALID_NAME_FOR_SPN,
		ERROR_DS_FILTER_USES_CONTRUCTED_ATTRS,
		ERROR_DS_UNICODEPWD_NOT_IN_QUOTES,
		ERROR_DS_MACHINE_ACCOUNT_QUOTA_EXCEEDED,
		ERROR_DS_MUST_BE_RUN_ON_DST_DC,
		ERROR_DS_SRC_DC_MUST_BE_SP4_OR_GREATER,
		ERROR_DS_CANT_TREE_DELETE_CRITICAL_OBJ,
		ERROR_DS_INIT_FAILURE_CONSOLE,
		ERROR_DS_SAM_INIT_FAILURE_CONSOLE,
		ERROR_DS_FOREST_VERSION_TOO_HIGH,
		ERROR_DS_DOMAIN_VERSION_TOO_HIGH,
		ERROR_DS_FOREST_VERSION_TOO_LOW,
		ERROR_DS_DOMAIN_VERSION_TOO_LOW,
		ERROR_DS_INCOMPATIBLE_VERSION,
		ERROR_DS_LOW_DSA_VERSION,
		ERROR_DS_NO_BEHAVIOR_VERSION_IN_MIXEDDOMAIN,
		ERROR_DS_NOT_SUPPORTED_SORT_ORDER,
		ERROR_DS_NAME_NOT_UNIQUE,
		ERROR_DS_MACHINE_ACCOUNT_CREATED_PRENT4,
		ERROR_DS_OUT_OF_VERSION_STORE,
		ERROR_DS_INCOMPATIBLE_CONTROLS_USED,
		ERROR_DS_NO_REF_DOMAIN,
		ERROR_DS_RESERVED_LINK_ID,
		ERROR_DS_LINK_ID_NOT_AVAILABLE,
		ERROR_DS_AG_CANT_HAVE_UNIVERSAL_MEMBER,
		ERROR_DS_MODIFYDN_DISALLOWED_BY_INSTANCE_TYPE,
		ERROR_DS_NO_OBJECT_MOVE_IN_SCHEMA_NC,
		ERROR_DS_MODIFYDN_DISALLOWED_BY_FLAG,
		ERROR_DS_MODIFYDN_WRONG_GRANDPARENT,
		ERROR_DS_NAME_ERROR_TRUST_REFERRAL,
		ERROR_NOT_SUPPORTED_ON_STANDARD_SERVER,
		ERROR_DS_CANT_ACCESS_REMOTE_PART_OF_AD,
		ERROR_DS_CR_IMPOSSIBLE_TO_VALIDATE_V2,
		ERROR_DS_THREAD_LIMIT_EXCEEDED,
		ERROR_DS_NOT_CLOSEST,
		ERROR_DS_CANT_DERIVE_SPN_WITHOUT_SERVER_REF,
		ERROR_DS_SINGLE_USER_MODE_FAILED,
		ERROR_DS_NTDSCRIPT_SYNTAX_ERROR,
		ERROR_DS_NTDSCRIPT_PROCESS_ERROR,
		ERROR_DS_DIFFERENT_REPL_EPOCHS,
		ERROR_DS_DRS_EXTENSIONS_CHANGED,
		ERROR_DS_REPLICA_SET_CHANGE_NOT_ALLOWED_ON_DISABLED_CR,
		ERROR_DS_NO_MSDS_INTID,
		ERROR_DS_DUP_MSDS_INTID,
		ERROR_DS_EXISTS_IN_RDNATTID,
		ERROR_DS_AUTHORIZATION_FAILED,
		ERROR_DS_INVALID_SCRIPT,
		ERROR_DS_REMOTE_CROSSREF_OP_FAILED,
		ERROR_DS_CROSS_REF_BUSY,
		ERROR_DS_CANT_DERIVE_SPN_FOR_DELETED_DOMAIN,
		ERROR_DS_CANT_DEMOTE_WITH_WRITEABLE_NC,
		ERROR_DS_DUPLICATE_ID_FOUND,
		ERROR_DS_INSUFFICIENT_ATTR_TO_CREATE_OBJECT,
		ERROR_DS_GROUP_CONVERSION_ERROR,
		ERROR_DS_CANT_MOVE_APP_BASIC_GROUP,
		ERROR_DS_CANT_MOVE_APP_QUERY_GROUP,
		ERROR_DS_ROLE_NOT_VERIFIED,
		ERROR_DS_WKO_CONTAINER_CANNOT_BE_SPECIAL,
		ERROR_DS_DOMAIN_RENAME_IN_PROGRESS,
		ERROR_DS_EXISTING_AD_CHILD_NC,
		ERROR_DS_REPL_LIFETIME_EXCEEDED,
		ERROR_DS_DISALLOWED_IN_SYSTEM_CONTAINER,
		ERROR_DS_LDAP_SEND_QUEUE_FULL,
		ERROR_DS_DRA_OUT_SCHEDULE_WINDOW,
		ERROR_DS_POLICY_NOT_KNOWN,
		ERROR_NO_SITE_SETTINGS_OBJECT,
		ERROR_NO_SECRETS,
		ERROR_NO_WRITABLE_DC_FOUND,
		ERROR_DS_NO_SERVER_OBJECT,
		ERROR_DS_NO_NTDSA_OBJECT,
		ERROR_DS_NON_ASQ_SEARCH,
		ERROR_DS_AUDIT_FAILURE,
		ERROR_DS_INVALID_SEARCH_FLAG_SUBTREE,
		ERROR_DS_INVALID_SEARCH_FLAG_TUPLE,
		ERROR_DS_HIERARCHY_TABLE_TOO_DEEP,
		ERROR_DS_DRA_CORRUPT_UTD_VECTOR,
		ERROR_DS_DRA_SECRETS_DENIED,
		ERROR_DS_RESERVED_MAPI_ID,
		ERROR_DS_MAPI_ID_NOT_AVAILABLE,
		ERROR_DS_DRA_MISSING_KRBTGT_SECRET,
		ERROR_DS_DOMAIN_NAME_EXISTS_IN_FOREST,
		ERROR_DS_FLAT_NAME_EXISTS_IN_FOREST,
		ERROR_INVALID_USER_PRINCIPAL_NAME,
		ERROR_DS_OID_MAPPED_GROUP_CANT_HAVE_MEMBERS,
		ERROR_DS_OID_NOT_FOUND,
		ERROR_DS_DRA_RECYCLED_TARGET,
		DNS_ERROR_RESPONSE_CODES_BASE,
		DNS_ERROR_MASK,
		DNS_ERROR_RCODE_FORMAT_ERROR,
		DNS_ERROR_RCODE_SERVER_FAILURE,
		DNS_ERROR_RCODE_NAME_ERROR,
		DNS_ERROR_RCODE_NOT_IMPLEMENTED,
		DNS_ERROR_RCODE_REFUSED,
		DNS_ERROR_RCODE_YXDOMAIN,
		DNS_ERROR_RCODE_YXRRSET,
		DNS_ERROR_RCODE_NXRRSET,
		DNS_ERROR_RCODE_NOTAUTH,
		DNS_ERROR_RCODE_NOTZONE,
		DNS_ERROR_RCODE_BADSIG,
		DNS_ERROR_RCODE_BADKEY,
		DNS_ERROR_RCODE_BADTIME,
		DNS_ERROR_PACKET_FMT_BASE,
		DNS_INFO_NO_RECORDS,
		DNS_ERROR_BAD_PACKET,
		DNS_ERROR_NO_PACKET,
		DNS_ERROR_RCODE,
		DNS_ERROR_UNSECURE_PACKET,
		DNS_ERROR_GENERAL_API_BASE,
		DNS_ERROR_INVALID_TYPE,
		DNS_ERROR_INVALID_IP_ADDRESS,
		DNS_ERROR_INVALID_PROPERTY,
		DNS_ERROR_TRY_AGAIN_LATER,
		DNS_ERROR_NOT_UNIQUE,
		DNS_ERROR_NON_RFC_NAME,
		DNS_STATUS_FQDN,
		DNS_STATUS_DOTTED_NAME,
		DNS_STATUS_SINGLE_PART_NAME,
		DNS_ERROR_INVALID_NAME_CHAR,
		DNS_ERROR_NUMERIC_NAME,
		DNS_ERROR_NOT_ALLOWED_ON_ROOT_SERVER,
		DNS_ERROR_NOT_ALLOWED_UNDER_DELEGATION,
		DNS_ERROR_CANNOT_FIND_ROOT_HINTS,
		DNS_ERROR_INCONSISTENT_ROOT_HINTS,
		DNS_ERROR_DWORD_VALUE_TOO_SMALL,
		DNS_ERROR_DWORD_VALUE_TOO_LARGE,
		DNS_ERROR_BACKGROUND_LOADING,
		DNS_ERROR_NOT_ALLOWED_ON_RODC,
		DNS_ERROR_NOT_ALLOWED_UNDER_DNAME,
		DNS_ERROR_DELEGATION_REQUIRED,
		DNS_ERROR_INVALID_POLICY_TABLE,
		DNS_ERROR_ZONE_BASE,
		DNS_ERROR_ZONE_DOES_NOT_EXIST,
		DNS_ERROR_NO_ZONE_INFO,
		DNS_ERROR_INVALID_ZONE_OPERATION,
		DNS_ERROR_ZONE_CONFIGURATION_ERROR,
		DNS_ERROR_ZONE_HAS_NO_SOA_RECORD,
		DNS_ERROR_ZONE_HAS_NO_NS_RECORDS,
		DNS_ERROR_ZONE_LOCKED,
		DNS_ERROR_ZONE_CREATION_FAILED,
		DNS_ERROR_ZONE_ALREADY_EXISTS,
		DNS_ERROR_AUTOZONE_ALREADY_EXISTS,
		DNS_ERROR_INVALID_ZONE_TYPE,
		DNS_ERROR_SECONDARY_REQUIRES_MASTER_IP,
		DNS_ERROR_ZONE_NOT_SECONDARY,
		DNS_ERROR_NEED_SECONDARY_ADDRESSES,
		DNS_ERROR_WINS_INIT_FAILED,
		DNS_ERROR_NEED_WINS_SERVERS,
		DNS_ERROR_NBSTAT_INIT_FAILED,
		DNS_ERROR_SOA_DELETE_INVALID,
		DNS_ERROR_FORWARDER_ALREADY_EXISTS,
		DNS_ERROR_ZONE_REQUIRES_MASTER_IP,
		DNS_ERROR_ZONE_IS_SHUTDOWN,
		DNS_ERROR_DATAFILE_BASE,
		DNS_ERROR_PRIMARY_REQUIRES_DATAFILE,
		DNS_ERROR_INVALID_DATAFILE_NAME,
		DNS_ERROR_DATAFILE_OPEN_FAILURE,
		DNS_ERROR_FILE_WRITEBACK_FAILED,
		DNS_ERROR_DATAFILE_PARSING,
		DNS_ERROR_DATABASE_BASE,
		DNS_ERROR_RECORD_DOES_NOT_EXIST,
		DNS_ERROR_RECORD_FORMAT,
		DNS_ERROR_NODE_CREATION_FAILED,
		DNS_ERROR_UNKNOWN_RECORD_TYPE,
		DNS_ERROR_RECORD_TIMED_OUT,
		DNS_ERROR_NAME_NOT_IN_ZONE,
		DNS_ERROR_CNAME_LOOP,
		DNS_ERROR_NODE_IS_CNAME,
		DNS_ERROR_CNAME_COLLISION,
		DNS_ERROR_RECORD_ONLY_AT_ZONE_ROOT,
		DNS_ERROR_RECORD_ALREADY_EXISTS,
		DNS_ERROR_SECONDARY_DATA,
		DNS_ERROR_NO_CREATE_CACHE_DATA,
		DNS_ERROR_NAME_DOES_NOT_EXIST,
		DNS_WARNING_PTR_CREATE_FAILED,
		DNS_WARNING_DOMAIN_UNDELETED,
		DNS_ERROR_DS_UNAVAILABLE,
		DNS_ERROR_DS_ZONE_ALREADY_EXISTS,
		DNS_ERROR_NO_BOOTFILE_IF_DS_ZONE,
		DNS_ERROR_NODE_IS_DNAME,
		DNS_ERROR_DNAME_COLLISION,
		DNS_ERROR_ALIAS_LOOP,
		DNS_ERROR_OPERATION_BASE,
		DNS_INFO_AXFR_COMPLETE,
		DNS_ERROR_AXFR,
		DNS_INFO_ADDED_LOCAL_WINS,
		DNS_ERROR_SECURE_BASE,
		DNS_STATUS_CONTINUE_NEEDED,
		DNS_ERROR_SETUP_BASE,
		DNS_ERROR_NO_TCPIP,
		DNS_ERROR_NO_DNS_SERVERS,
		DNS_ERROR_DP_BASE,
		DNS_ERROR_DP_DOES_NOT_EXIST,
		DNS_ERROR_DP_ALREADY_EXISTS,
		DNS_ERROR_DP_NOT_ENLISTED,
		DNS_ERROR_DP_ALREADY_ENLISTED,
		DNS_ERROR_DP_NOT_AVAILABLE,
		DNS_ERROR_DP_FSMO_ERROR,
		WSABASEERR,
		WSAEINTR,
		WSAEBADF,
		WSAEACCES,
		WSAEFAULT,
		WSAEINVAL,
		WSAEMFILE,
		WSAEWOULDBLOCK,
		WSAEINPROGRESS,
		WSAEALREADY,
		WSAENOTSOCK,
		WSAEDESTADDRREQ,
		WSAEMSGSIZE,
		WSAEPROTOTYPE,
		WSAENOPROTOOPT,
		WSAEPROTONOSUPPORT,
		WSAESOCKTNOSUPPORT,
		WSAEOPNOTSUPP,
		WSAEPFNOSUPPORT,
		WSAEAFNOSUPPORT,
		WSAEADDRINUSE,
		WSAEADDRNOTAVAIL,
		WSAENETDOWN,
		WSAENETUNREACH,
		WSAENETRESET,
		WSAECONNABORTED,
		WSAECONNRESET,
		WSAENOBUFS,
		WSAEISCONN,
		WSAENOTCONN,
		WSAESHUTDOWN,
		WSAETOOMANYREFS,
		WSAETIMEDOUT,
		WSAECONNREFUSED,
		WSAELOOP,
		WSAENAMETOOLONG,
		WSAEHOSTDOWN,
		WSAEHOSTUNREACH,
		WSAENOTEMPTY,
		WSAEPROCLIM,
		WSAEUSERS,
		WSAEDQUOT,
		WSAESTALE,
		WSAEREMOTE,
		WSASYSNOTREADY,
		WSAVERNOTSUPPORTED,
		WSANOTINITIALISED,
		WSAEDISCON,
		WSAENOMORE,
		WSAECANCELLED,
		WSAEINVALIDPROCTABLE,
		WSAEINVALIDPROVIDER,
		WSAEPROVIDERFAILEDINIT,
		WSASYSCALLFAILURE,
		WSASERVICE_NOT_FOUND,
		WSATYPE_NOT_FOUND,
		WSA_E_NO_MORE,
		WSA_E_CANCELLED,
		WSAEREFUSED,
		WSAHOST_NOT_FOUND,
		WSATRY_AGAIN,
		WSANO_RECOVERY,
		WSANO_DATA,
		WSA_QOS_RECEIVERS,
		WSA_QOS_SENDERS,
		WSA_QOS_NO_SENDERS,
		WSA_QOS_NO_RECEIVERS,
		WSA_QOS_REQUEST_CONFIRMED,
		WSA_QOS_ADMISSION_FAILURE,
		WSA_QOS_POLICY_FAILURE,
		WSA_QOS_BAD_STYLE,
		WSA_QOS_BAD_OBJECT,
		WSA_QOS_TRAFFIC_CTRL_ERROR,
		WSA_QOS_GENERIC_ERROR,
		WSA_QOS_ESERVICETYPE,
		WSA_QOS_EFLOWSPEC,
		WSA_QOS_EPROVSPECBUF,
		WSA_QOS_EFILTERSTYLE,
		WSA_QOS_EFILTERTYPE,
		WSA_QOS_EFILTERCOUNT,
		WSA_QOS_EOBJLENGTH,
		WSA_QOS_EFLOWCOUNT,
		WSA_QOS_EUNKOWNPSOBJ,
		WSA_QOS_EPOLICYOBJ,
		WSA_QOS_EFLOWDESC,
		WSA_QOS_EPSFLOWSPEC,
		WSA_QOS_EPSFILTERSPEC,
		WSA_QOS_ESDMODEOBJ,
		WSA_QOS_ESHAPERATEOBJ,
		WSA_QOS_RESERVED_PETYPE,
		WSA_SECURE_HOST_NOT_FOUND,
		WSA_IPSEC_NAME_POLICY_ERROR,
		ERROR_IPSEC_QM_POLICY_EXISTS,
		ERROR_IPSEC_QM_POLICY_NOT_FOUND,
		ERROR_IPSEC_QM_POLICY_IN_USE,
		ERROR_IPSEC_MM_POLICY_EXISTS,
		ERROR_IPSEC_MM_POLICY_NOT_FOUND,
		ERROR_IPSEC_MM_POLICY_IN_USE,
		ERROR_IPSEC_MM_FILTER_EXISTS,
		ERROR_IPSEC_MM_FILTER_NOT_FOUND,
		ERROR_IPSEC_TRANSPORT_FILTER_EXISTS,
		ERROR_IPSEC_TRANSPORT_FILTER_NOT_FOUND,
		ERROR_IPSEC_MM_AUTH_EXISTS,
		ERROR_IPSEC_MM_AUTH_NOT_FOUND,
		ERROR_IPSEC_MM_AUTH_IN_USE,
		ERROR_IPSEC_DEFAULT_MM_POLICY_NOT_FOUND,
		ERROR_IPSEC_DEFAULT_MM_AUTH_NOT_FOUND,
		ERROR_IPSEC_DEFAULT_QM_POLICY_NOT_FOUND,
		ERROR_IPSEC_TUNNEL_FILTER_EXISTS,
		ERROR_IPSEC_TUNNEL_FILTER_NOT_FOUND,
		ERROR_IPSEC_MM_FILTER_PENDING_DELETION,
		ERROR_IPSEC_TRANSPORT_FILTER_PENDING_DELETION,
		ERROR_IPSEC_TUNNEL_FILTER_PENDING_DELETION,
		ERROR_IPSEC_MM_POLICY_PENDING_DELETION,
		ERROR_IPSEC_MM_AUTH_PENDING_DELETION,
		ERROR_IPSEC_QM_POLICY_PENDING_DELETION,
		WARNING_IPSEC_MM_POLICY_PRUNED,
		WARNING_IPSEC_QM_POLICY_PRUNED,
		ERROR_IPSEC_IKE_NEG_STATUS_BEGIN,
		ERROR_IPSEC_IKE_AUTH_FAIL,
		ERROR_IPSEC_IKE_ATTRIB_FAIL,
		ERROR_IPSEC_IKE_NEGOTIATION_PENDING,
		ERROR_IPSEC_IKE_GENERAL_PROCESSING_ERROR,
		ERROR_IPSEC_IKE_TIMED_OUT,
		ERROR_IPSEC_IKE_NO_CERT,
		ERROR_IPSEC_IKE_SA_DELETED,
		ERROR_IPSEC_IKE_SA_REAPED,
		ERROR_IPSEC_IKE_MM_ACQUIRE_DROP,
		ERROR_IPSEC_IKE_QM_ACQUIRE_DROP,
		ERROR_IPSEC_IKE_QUEUE_DROP_MM,
		ERROR_IPSEC_IKE_QUEUE_DROP_NO_MM,
		ERROR_IPSEC_IKE_DROP_NO_RESPONSE,
		ERROR_IPSEC_IKE_MM_DELAY_DROP,
		ERROR_IPSEC_IKE_QM_DELAY_DROP,
		ERROR_IPSEC_IKE_ERROR,
		ERROR_IPSEC_IKE_CRL_FAILED,
		ERROR_IPSEC_IKE_INVALID_KEY_USAGE,
		ERROR_IPSEC_IKE_INVALID_CERT_TYPE,
		ERROR_IPSEC_IKE_NO_PRIVATE_KEY,
		ERROR_IPSEC_IKE_SIMULTANEOUS_REKEY,
		ERROR_IPSEC_IKE_DH_FAIL,
		ERROR_IPSEC_IKE_CRITICAL_PAYLOAD_NOT_RECOGNIZED,
		ERROR_IPSEC_IKE_INVALID_HEADER,
		ERROR_IPSEC_IKE_NO_POLICY,
		ERROR_IPSEC_IKE_INVALID_SIGNATURE,
		ERROR_IPSEC_IKE_KERBEROS_ERROR,
		ERROR_IPSEC_IKE_NO_PUBLIC_KEY,
		ERROR_IPSEC_IKE_PROCESS_ERR,
		ERROR_IPSEC_IKE_PROCESS_ERR_SA,
		ERROR_IPSEC_IKE_PROCESS_ERR_PROP,
		ERROR_IPSEC_IKE_PROCESS_ERR_TRANS,
		ERROR_IPSEC_IKE_PROCESS_ERR_KE,
		ERROR_IPSEC_IKE_PROCESS_ERR_ID,
		ERROR_IPSEC_IKE_PROCESS_ERR_CERT,
		ERROR_IPSEC_IKE_PROCESS_ERR_CERT_REQ,
		ERROR_IPSEC_IKE_PROCESS_ERR_HASH,
		ERROR_IPSEC_IKE_PROCESS_ERR_SIG,
		ERROR_IPSEC_IKE_PROCESS_ERR_NONCE,
		ERROR_IPSEC_IKE_PROCESS_ERR_NOTIFY,
		ERROR_IPSEC_IKE_PROCESS_ERR_DELETE,
		ERROR_IPSEC_IKE_PROCESS_ERR_VENDOR,
		ERROR_IPSEC_IKE_INVALID_PAYLOAD,
		ERROR_IPSEC_IKE_LOAD_SOFT_SA,
		ERROR_IPSEC_IKE_SOFT_SA_TORN_DOWN,
		ERROR_IPSEC_IKE_INVALID_COOKIE,
		ERROR_IPSEC_IKE_NO_PEER_CERT,
		ERROR_IPSEC_IKE_PEER_CRL_FAILED,
		ERROR_IPSEC_IKE_POLICY_CHANGE,
		ERROR_IPSEC_IKE_NO_MM_POLICY,
		ERROR_IPSEC_IKE_NOTCBPRIV,
		ERROR_IPSEC_IKE_SECLOADFAIL,
		ERROR_IPSEC_IKE_FAILSSPINIT,
		ERROR_IPSEC_IKE_FAILQUERYSSP,
		ERROR_IPSEC_IKE_SRVACQFAIL,
		ERROR_IPSEC_IKE_SRVQUERYCRED,
		ERROR_IPSEC_IKE_GETSPIFAIL,
		ERROR_IPSEC_IKE_INVALID_FILTER,
		ERROR_IPSEC_IKE_OUT_OF_MEMORY,
		ERROR_IPSEC_IKE_ADD_UPDATE_KEY_FAILED,
		ERROR_IPSEC_IKE_INVALID_POLICY,
		ERROR_IPSEC_IKE_UNKNOWN_DOI,
		ERROR_IPSEC_IKE_INVALID_SITUATION,
		ERROR_IPSEC_IKE_DH_FAILURE,
		ERROR_IPSEC_IKE_INVALID_GROUP,
		ERROR_IPSEC_IKE_ENCRYPT,
		ERROR_IPSEC_IKE_DECRYPT,
		ERROR_IPSEC_IKE_POLICY_MATCH,
		ERROR_IPSEC_IKE_UNSUPPORTED_ID,
		ERROR_IPSEC_IKE_INVALID_HASH,
		ERROR_IPSEC_IKE_INVALID_HASH_ALG,
		ERROR_IPSEC_IKE_INVALID_HASH_SIZE,
		ERROR_IPSEC_IKE_INVALID_ENCRYPT_ALG,
		ERROR_IPSEC_IKE_INVALID_AUTH_ALG,
		ERROR_IPSEC_IKE_INVALID_SIG,
		ERROR_IPSEC_IKE_LOAD_FAILED,
		ERROR_IPSEC_IKE_RPC_DELETE,
		ERROR_IPSEC_IKE_BENIGN_REINIT,
		ERROR_IPSEC_IKE_INVALID_RESPONDER_LIFETIME_NOTIFY,
		ERROR_IPSEC_IKE_INVALID_MAJOR_VERSION,
		ERROR_IPSEC_IKE_INVALID_CERT_KEYLEN,
		ERROR_IPSEC_IKE_MM_LIMIT,
		ERROR_IPSEC_IKE_NEGOTIATION_DISABLED,
		ERROR_IPSEC_IKE_QM_LIMIT,
		ERROR_IPSEC_IKE_MM_EXPIRED,
		ERROR_IPSEC_IKE_PEER_MM_ASSUMED_INVALID,
		ERROR_IPSEC_IKE_CERT_CHAIN_POLICY_MISMATCH,
		ERROR_IPSEC_IKE_UNEXPECTED_MESSAGE_ID,
		ERROR_IPSEC_IKE_INVALID_AUTH_PAYLOAD,
		ERROR_IPSEC_IKE_DOS_COOKIE_SENT,
		ERROR_IPSEC_IKE_SHUTTING_DOWN,
		ERROR_IPSEC_IKE_CGA_AUTH_FAILED,
		ERROR_IPSEC_IKE_PROCESS_ERR_NATOA,
		ERROR_IPSEC_IKE_INVALID_MM_FOR_QM,
		ERROR_IPSEC_IKE_QM_EXPIRED,
		ERROR_IPSEC_IKE_TOO_MANY_FILTERS,
		ERROR_IPSEC_IKE_NEG_STATUS_END,
		ERROR_IPSEC_IKE_KILL_DUMMY_NAP_TUNNEL,
		ERROR_IPSEC_IKE_INNER_IP_ASSIGNMENT_FAILURE,
		ERROR_IPSEC_IKE_REQUIRE_CP_PAYLOAD_MISSING,
		ERROR_IPSEC_KEY_MODULE_IMPERSONATION_NEGOTIATION_PENDING,
		ERROR_IPSEC_IKE_COEXISTENCE_SUPPRESS,
		ERROR_IPSEC_IKE_RATELIMIT_DROP,
		ERROR_IPSEC_IKE_PEER_DOESNT_SUPPORT_MOBIKE,
		ERROR_IPSEC_IKE_AUTHORIZATION_FAILURE,
		ERROR_IPSEC_IKE_STRONG_CRED_AUTHORIZATION_FAILURE,
		ERROR_IPSEC_IKE_AUTHORIZATION_FAILURE_WITH_OPTIONAL_RETRY,
		ERROR_IPSEC_IKE_STRONG_CRED_AUTHORIZATION_AND_CERTMAP_FAILURE,
		ERROR_IPSEC_IKE_NEG_STATUS_EXTENDED_END,
		ERROR_IPSEC_BAD_SPI,
		ERROR_IPSEC_SA_LIFETIME_EXPIRED,
		ERROR_IPSEC_WRONG_SA,
		ERROR_IPSEC_REPLAY_CHECK_FAILED,
		ERROR_IPSEC_INVALID_PACKET,
		ERROR_IPSEC_INTEGRITY_CHECK_FAILED,
		ERROR_IPSEC_CLEAR_TEXT_DROP,
		ERROR_IPSEC_AUTH_FIREWALL_DROP,
		ERROR_IPSEC_THROTTLE_DROP,
		ERROR_IPSEC_DOSP_BLOCK,
		ERROR_IPSEC_DOSP_RECEIVED_MULTICAST,
		ERROR_IPSEC_DOSP_INVALID_PACKET,
		ERROR_IPSEC_DOSP_STATE_LOOKUP_FAILED,
		ERROR_IPSEC_DOSP_MAX_ENTRIES,
		ERROR_IPSEC_DOSP_KEYMOD_NOT_ALLOWED,
		ERROR_IPSEC_DOSP_NOT_INSTALLED,
		ERROR_IPSEC_DOSP_MAX_PER_IP_RATELIMIT_QUEUES,
		ERROR_SXS_SECTION_NOT_FOUND,
		ERROR_SXS_CANT_GEN_ACTCTX,
		ERROR_SXS_INVALID_ACTCTXDATA_FORMAT,
		ERROR_SXS_ASSEMBLY_NOT_FOUND,
		ERROR_SXS_MANIFEST_FORMAT_ERROR,
		ERROR_SXS_MANIFEST_PARSE_ERROR,
		ERROR_SXS_ACTIVATION_CONTEXT_DISABLED,
		ERROR_SXS_KEY_NOT_FOUND,
		ERROR_SXS_VERSION_CONFLICT,
		ERROR_SXS_WRONG_SECTION_TYPE,
		ERROR_SXS_THREAD_QUERIES_DISABLED,
		ERROR_SXS_PROCESS_DEFAULT_ALREADY_SET,
		ERROR_SXS_UNKNOWN_ENCODING_GROUP,
		ERROR_SXS_UNKNOWN_ENCODING,
		ERROR_SXS_INVALID_XML_NAMESPACE_URI,
		ERROR_SXS_ROOT_MANIFEST_DEPENDENCY_NOT_INSTALLED,
		ERROR_SXS_LEAF_MANIFEST_DEPENDENCY_NOT_INSTALLED,
		ERROR_SXS_INVALID_ASSEMBLY_IDENTITY_ATTRIBUTE,
		ERROR_SXS_MANIFEST_MISSING_REQUIRED_DEFAULT_NAMESPACE,
		ERROR_SXS_MANIFEST_INVALID_REQUIRED_DEFAULT_NAMESPACE,
		ERROR_SXS_PRIVATE_MANIFEST_CROSS_PATH_WITH_REPARSE_POINT,
		ERROR_SXS_DUPLICATE_DLL_NAME,
		ERROR_SXS_DUPLICATE_WINDOWCLASS_NAME,
		ERROR_SXS_DUPLICATE_CLSID,
		ERROR_SXS_DUPLICATE_IID,
		ERROR_SXS_DUPLICATE_TLBID,
		ERROR_SXS_DUPLICATE_PROGID,
		ERROR_SXS_DUPLICATE_ASSEMBLY_NAME,
		ERROR_SXS_FILE_HASH_MISMATCH,
		ERROR_SXS_POLICY_PARSE_ERROR,
		ERROR_SXS_XML_E_MISSINGQUOTE,
		ERROR_SXS_XML_E_COMMENTSYNTAX,
		ERROR_SXS_XML_E_BADSTARTNAMECHAR,
		ERROR_SXS_XML_E_BADNAMECHAR,
		ERROR_SXS_XML_E_BADCHARINSTRING,
		ERROR_SXS_XML_E_XMLDECLSYNTAX,
		ERROR_SXS_XML_E_BADCHARDATA,
		ERROR_SXS_XML_E_MISSINGWHITESPACE,
		ERROR_SXS_XML_E_EXPECTINGTAGEND,
		ERROR_SXS_XML_E_MISSINGSEMICOLON,
		ERROR_SXS_XML_E_UNBALANCEDPAREN,
		ERROR_SXS_XML_E_INTERNALERROR,
		ERROR_SXS_XML_E_UNEXPECTED_WHITESPACE,
		ERROR_SXS_XML_E_INCOMPLETE_ENCODING,
		ERROR_SXS_XML_E_MISSING_PAREN,
		ERROR_SXS_XML_E_EXPECTINGCLOSEQUOTE,
		ERROR_SXS_XML_E_MULTIPLE_COLONS,
		ERROR_SXS_XML_E_INVALID_DECIMAL,
		ERROR_SXS_XML_E_INVALID_HEXIDECIMAL,
		ERROR_SXS_XML_E_INVALID_UNICODE,
		ERROR_SXS_XML_E_WHITESPACEORQUESTIONMARK,
		ERROR_SXS_XML_E_UNEXPECTEDENDTAG,
		ERROR_SXS_XML_E_UNCLOSEDTAG,
		ERROR_SXS_XML_E_DUPLICATEATTRIBUTE,
		ERROR_SXS_XML_E_MULTIPLEROOTS,
		ERROR_SXS_XML_E_INVALIDATROOTLEVEL,
		ERROR_SXS_XML_E_BADXMLDECL,
		ERROR_SXS_XML_E_MISSINGROOT,
		ERROR_SXS_XML_E_UNEXPECTEDEOF,
		ERROR_SXS_XML_E_BADPEREFINSUBSET,
		ERROR_SXS_XML_E_UNCLOSEDSTARTTAG,
		ERROR_SXS_XML_E_UNCLOSEDENDTAG,
		ERROR_SXS_XML_E_UNCLOSEDSTRING,
		ERROR_SXS_XML_E_UNCLOSEDCOMMENT,
		ERROR_SXS_XML_E_UNCLOSEDDECL,
		ERROR_SXS_XML_E_UNCLOSEDCDATA,
		ERROR_SXS_XML_E_RESERVEDNAMESPACE,
		ERROR_SXS_XML_E_INVALIDENCODING,
		ERROR_SXS_XML_E_INVALIDSWITCH,
		ERROR_SXS_XML_E_BADXMLCASE,
		ERROR_SXS_XML_E_INVALID_STANDALONE,
		ERROR_SXS_XML_E_UNEXPECTED_STANDALONE,
		ERROR_SXS_XML_E_INVALID_VERSION,
		ERROR_SXS_XML_E_MISSINGEQUALS,
		ERROR_SXS_PROTECTION_RECOVERY_FAILED,
		ERROR_SXS_PROTECTION_PUBLIC_KEY_TOO_SHORT,
		ERROR_SXS_PROTECTION_CATALOG_NOT_VALID,
		ERROR_SXS_UNTRANSLATABLE_HRESULT,
		ERROR_SXS_PROTECTION_CATALOG_FILE_MISSING,
		ERROR_SXS_MISSING_ASSEMBLY_IDENTITY_ATTRIBUTE,
		ERROR_SXS_INVALID_ASSEMBLY_IDENTITY_ATTRIBUTE_NAME,
		ERROR_SXS_ASSEMBLY_MISSING,
		ERROR_SXS_CORRUPT_ACTIVATION_STACK,
		ERROR_SXS_CORRUPTION,
		ERROR_SXS_EARLY_DEACTIVATION,
		ERROR_SXS_INVALID_DEACTIVATION,
		ERROR_SXS_MULTIPLE_DEACTIVATION,
		ERROR_SXS_PROCESS_TERMINATION_REQUESTED,
		ERROR_SXS_RELEASE_ACTIVATION_CONTEXT,
		ERROR_SXS_SYSTEM_DEFAULT_ACTIVATION_CONTEXT_EMPTY,
		ERROR_SXS_INVALID_IDENTITY_ATTRIBUTE_VALUE,
		ERROR_SXS_INVALID_IDENTITY_ATTRIBUTE_NAME,
		ERROR_SXS_IDENTITY_DUPLICATE_ATTRIBUTE,
		ERROR_SXS_IDENTITY_PARSE_ERROR,
		ERROR_MALFORMED_SUBSTITUTION_STRING,
		ERROR_SXS_INCORRECT_PUBLIC_KEY_TOKEN,
		ERROR_UNMAPPED_SUBSTITUTION_STRING,
		ERROR_SXS_ASSEMBLY_NOT_LOCKED,
		ERROR_SXS_COMPONENT_STORE_CORRUPT,
		ERROR_ADVANCED_INSTALLER_FAILED,
		ERROR_XML_ENCODING_MISMATCH,
		ERROR_SXS_MANIFEST_IDENTITY_SAME_BUT_CONTENTS_DIFFERENT,
		ERROR_SXS_IDENTITIES_DIFFERENT,
		ERROR_SXS_ASSEMBLY_IS_NOT_A_DEPLOYMENT,
		ERROR_SXS_FILE_NOT_PART_OF_ASSEMBLY,
		ERROR_SXS_MANIFEST_TOO_BIG,
		ERROR_SXS_SETTING_NOT_REGISTERED,
		ERROR_SXS_TRANSACTION_CLOSURE_INCOMPLETE,
		ERROR_SMI_PRIMITIVE_INSTALLER_FAILED,
		ERROR_GENERIC_COMMAND_FAILED,
		ERROR_SXS_FILE_HASH_MISSING,
		ERROR_EVT_INVALID_CHANNEL_PATH,
		ERROR_EVT_INVALID_QUERY,
		ERROR_EVT_PUBLISHER_METADATA_NOT_FOUND,
		ERROR_EVT_EVENT_TEMPLATE_NOT_FOUND,
		ERROR_EVT_INVALID_PUBLISHER_NAME,
		ERROR_EVT_INVALID_EVENT_DATA,
		ERROR_EVT_CHANNEL_NOT_FOUND,
		ERROR_EVT_MALFORMED_XML_TEXT,
		ERROR_EVT_SUBSCRIPTION_TO_DIRECT_CHANNEL,
		ERROR_EVT_CONFIGURATION_ERROR,
		ERROR_EVT_QUERY_RESULT_STALE,
		ERROR_EVT_QUERY_RESULT_INVALID_POSITION,
		ERROR_EVT_NON_VALIDATING_MSXML,
		ERROR_EVT_FILTER_ALREADYSCOPED,
		ERROR_EVT_FILTER_NOTELTSET,
		ERROR_EVT_FILTER_INVARG,
		ERROR_EVT_FILTER_INVTEST,
		ERROR_EVT_FILTER_INVTYPE,
		ERROR_EVT_FILTER_PARSEERR,
		ERROR_EVT_FILTER_UNSUPPORTEDOP,
		ERROR_EVT_FILTER_UNEXPECTEDTOKEN,
		ERROR_EVT_INVALID_OPERATION_OVER_ENABLED_DIRECT_CHANNEL,
		ERROR_EVT_INVALID_CHANNEL_PROPERTY_VALUE,
		ERROR_EVT_INVALID_PUBLISHER_PROPERTY_VALUE,
		ERROR_EVT_CHANNEL_CANNOT_ACTIVATE,
		ERROR_EVT_FILTER_TOO_COMPLEX,
		ERROR_EVT_MESSAGE_NOT_FOUND,
		ERROR_EVT_MESSAGE_ID_NOT_FOUND,
		ERROR_EVT_UNRESOLVED_VALUE_INSERT,
		ERROR_EVT_UNRESOLVED_PARAMETER_INSERT,
		ERROR_EVT_MAX_INSERTS_REACHED,
		ERROR_EVT_EVENT_DEFINITION_NOT_FOUND,
		ERROR_EVT_MESSAGE_LOCALE_NOT_FOUND,
		ERROR_EVT_VERSION_TOO_OLD,
		ERROR_EVT_VERSION_TOO_NEW,
		ERROR_EVT_CANNOT_OPEN_CHANNEL_OF_QUERY,
		ERROR_EVT_PUBLISHER_DISABLED,
		ERROR_EVT_FILTER_OUT_OF_RANGE,
		ERROR_EC_SUBSCRIPTION_CANNOT_ACTIVATE,
		ERROR_EC_LOG_DISABLED,
		ERROR_EC_CIRCULAR_FORWARDING,
		ERROR_EC_CREDSTORE_FULL,
		ERROR_EC_CRED_NOT_FOUND,
		ERROR_EC_NO_ACTIVE_CHANNEL,
		ERROR_MUI_FILE_NOT_FOUND,
		ERROR_MUI_INVALID_FILE,
		ERROR_MUI_INVALID_RC_CONFIG,
		ERROR_MUI_INVALID_LOCALE_NAME,
		ERROR_MUI_INVALID_ULTIMATEFALLBACK_NAME,
		ERROR_MUI_FILE_NOT_LOADED,
		ERROR_RESOURCE_ENUM_USER_STOP,
		ERROR_MUI_INTLSETTINGS_UILANG_NOT_INSTALLED,
		ERROR_MUI_INTLSETTINGS_INVALID_LOCALE_NAME,
		ERROR_MCA_INVALID_CAPABILITIES_STRING,
		ERROR_MCA_INVALID_VCP_VERSION,
		ERROR_MCA_MONITOR_VIOLATES_MCCS_SPECIFICATION,
		ERROR_MCA_MCCS_VERSION_MISMATCH,
		ERROR_MCA_UNSUPPORTED_MCCS_VERSION,
		ERROR_MCA_INTERNAL_ERROR,
		ERROR_MCA_INVALID_TECHNOLOGY_TYPE_RETURNED,
		ERROR_MCA_UNSUPPORTED_COLOR_TEMPERATURE,
		ERROR_AMBIGUOUS_SYSTEM_DEVICE,
		ERROR_SYSTEM_DEVICE_NOT_FOUND,
		ERROR_HASH_NOT_SUPPORTED,
		ERROR_HASH_NOT_PRESENT,
		SEVERITY_SUCCESS,
		SEVERITY_ERROR,
		FACILITY_NT_BIT,
		E_UNEXPECTED,
		E_NOTIMPL,
		E_OUTOFMEMORY,
		E_INVALIDARG,
		E_NOINTERFACE,
		E_POINTER,
		E_HANDLE,
		E_ABORT,
		E_FAIL,
		E_ACCESSDENIED,
		E_PENDING,
		CO_E_INIT_TLS,
		CO_E_INIT_SHARED_ALLOCATOR,
		CO_E_INIT_MEMORY_ALLOCATOR,
		CO_E_INIT_CLASS_CACHE,
		CO_E_INIT_RPC_CHANNEL,
		CO_E_INIT_TLS_SET_CHANNEL_CONTROL,
		CO_E_INIT_TLS_CHANNEL_CONTROL,
		CO_E_INIT_UNACCEPTED_USER_ALLOCATOR,
		CO_E_INIT_SCM_MUTEX_EXISTS,
		CO_E_INIT_SCM_FILE_MAPPING_EXISTS,
		CO_E_INIT_SCM_MAP_VIEW_OF_FILE,
		CO_E_INIT_SCM_EXEC_FAILURE,
		CO_E_INIT_ONLY_SINGLE_THREADED,
		CO_E_CANT_REMOTE,
		CO_E_BAD_SERVER_NAME,
		CO_E_WRONG_SERVER_IDENTITY,
		CO_E_OLE1DDE_DISABLED,
		CO_E_RUNAS_SYNTAX,
		CO_E_CREATEPROCESS_FAILURE,
		CO_E_RUNAS_CREATEPROCESS_FAILURE,
		CO_E_RUNAS_LOGON_FAILURE,
		CO_E_LAUNCH_PERMSSION_DENIED,
		CO_E_START_SERVICE_FAILURE,
		CO_E_REMOTE_COMMUNICATION_FAILURE,
		CO_E_SERVER_START_TIMEOUT,
		CO_E_CLSREG_INCONSISTENT,
		CO_E_IIDREG_INCONSISTENT,
		CO_E_NOT_SUPPORTED,
		CO_E_RELOAD_DLL,
		CO_E_MSI_ERROR,
		CO_E_ATTEMPT_TO_CREATE_OUTSIDE_CLIENT_CONTEXT,
		CO_E_SERVER_PAUSED,
		CO_E_SERVER_NOT_PAUSED,
		CO_E_CLASS_DISABLED,
		CO_E_CLRNOTAVAILABLE,
		CO_E_ASYNC_WORK_REJECTED,
		CO_E_SERVER_INIT_TIMEOUT,
		CO_E_NO_SECCTX_IN_ACTIVATE,
		CO_E_TRACKER_CONFIG,
		CO_E_THREADPOOL_CONFIG,
		CO_E_SXS_CONFIG,
		CO_E_MALFORMED_SPN,
		S_OK,
		S_FALSE,
		OLE_E_FIRST,
		OLE_E_LAST,
		OLE_S_FIRST,
		OLE_S_LAST,
		OLE_E_OLEVERB,
		OLE_E_ADVF,
		OLE_E_ENUM_NOMORE,
		OLE_E_ADVISENOTSUPPORTED,
		OLE_E_NOCONNECTION,
		OLE_E_NOTRUNNING,
		OLE_E_NOCACHE,
		OLE_E_BLANK,
		OLE_E_CLASSDIFF,
		OLE_E_CANT_GETMONIKER,
		OLE_E_CANT_BINDTOSOURCE,
		OLE_E_STATIC,
		OLE_E_PROMPTSAVECANCELLED,
		OLE_E_INVALIDRECT,
		OLE_E_WRONGCOMPOBJ,
		OLE_E_INVALIDHWND,
		OLE_E_NOT_INPLACEACTIVE,
		OLE_E_CANTCONVERT,
		OLE_E_NOSTORAGE,
		DV_E_FORMATETC,
		DV_E_DVTARGETDEVICE,
		DV_E_STGMEDIUM,
		DV_E_STATDATA,
		DV_E_LINDEX,
		DV_E_TYMED,
		DV_E_CLIPFORMAT,
		DV_E_DVASPECT,
		DV_E_DVTARGETDEVICE_SIZE,
		DV_E_NOIVIEWOBJECT,
		DRAGDROP_E_FIRST,
		DRAGDROP_E_LAST,
		DRAGDROP_S_FIRST,
		DRAGDROP_S_LAST,
		DRAGDROP_E_NOTREGISTERED,
		DRAGDROP_E_ALREADYREGISTERED,
		DRAGDROP_E_INVALIDHWND,
		CLASSFACTORY_E_FIRST,
		CLASSFACTORY_E_LAST,
		CLASSFACTORY_S_FIRST,
		CLASSFACTORY_S_LAST,
		CLASS_E_NOAGGREGATION,
		CLASS_E_CLASSNOTAVAILABLE,
		CLASS_E_NOTLICENSED,
		MARSHAL_E_FIRST,
		MARSHAL_E_LAST,
		MARSHAL_S_FIRST,
		MARSHAL_S_LAST,
		DATA_E_FIRST,
		DATA_E_LAST,
		DATA_S_FIRST,
		DATA_S_LAST,
		VIEW_E_FIRST,
		VIEW_E_LAST,
		VIEW_S_FIRST,
		VIEW_S_LAST,
		VIEW_E_DRAW,
		REGDB_E_FIRST,
		REGDB_E_LAST,
		REGDB_S_FIRST,
		REGDB_S_LAST,
		REGDB_E_READREGDB,
		REGDB_E_WRITEREGDB,
		REGDB_E_KEYMISSING,
		REGDB_E_INVALIDVALUE,
		REGDB_E_CLASSNOTREG,
		REGDB_E_IIDNOTREG,
		REGDB_E_BADTHREADINGMODEL,
		CAT_E_FIRST,
		CAT_E_LAST,
		CAT_E_CATIDNOEXIST,
		CAT_E_NODESCRIPTION,
		CS_E_FIRST,
		CS_E_LAST,
		CS_E_PACKAGE_NOTFOUND,
		CS_E_NOT_DELETABLE,
		CS_E_CLASS_NOTFOUND,
		CS_E_INVALID_VERSION,
		CS_E_NO_CLASSSTORE,
		CS_E_OBJECT_NOTFOUND,
		CS_E_OBJECT_ALREADY_EXISTS,
		CS_E_INVALID_PATH,
		CS_E_NETWORK_ERROR,
		CS_E_ADMIN_LIMIT_EXCEEDED,
		CS_E_SCHEMA_MISMATCH,
		CS_E_INTERNAL_ERROR,
		CACHE_E_FIRST,
		CACHE_E_LAST,
		CACHE_S_FIRST,
		CACHE_S_LAST,
		CACHE_E_NOCACHE_UPDATED,
		OLEOBJ_E_FIRST,
		OLEOBJ_E_LAST,
		OLEOBJ_S_FIRST,
		OLEOBJ_S_LAST,
		OLEOBJ_E_NOVERBS,
		OLEOBJ_E_INVALIDVERB,
		CLIENTSITE_E_FIRST,
		CLIENTSITE_E_LAST,
		CLIENTSITE_S_FIRST,
		CLIENTSITE_S_LAST,
		INPLACE_E_NOTUNDOABLE,
		INPLACE_E_NOTOOLSPACE,
		INPLACE_E_FIRST,
		INPLACE_E_LAST,
		INPLACE_S_FIRST,
		INPLACE_S_LAST,
		ENUM_E_FIRST,
		ENUM_E_LAST,
		ENUM_S_FIRST,
		ENUM_S_LAST,
		CONVERT10_E_FIRST,
		CONVERT10_E_LAST,
		CONVERT10_S_FIRST,
		CONVERT10_S_LAST,
		CONVERT10_E_OLESTREAM_GET,
		CONVERT10_E_OLESTREAM_PUT,
		CONVERT10_E_OLESTREAM_FMT,
		CONVERT10_E_OLESTREAM_BITMAP_TO_DIB,
		CONVERT10_E_STG_FMT,
		CONVERT10_E_STG_NO_STD_STREAM,
		CONVERT10_E_STG_DIB_TO_BITMAP,
		CLIPBRD_E_FIRST,
		CLIPBRD_E_LAST,
		CLIPBRD_S_FIRST,
		CLIPBRD_S_LAST,
		CLIPBRD_E_CANT_OPEN,
		CLIPBRD_E_CANT_EMPTY,
		CLIPBRD_E_CANT_SET,
		CLIPBRD_E_BAD_DATA,
		CLIPBRD_E_CANT_CLOSE,
		MK_E_FIRST,
		MK_E_LAST,
		MK_S_FIRST,
		MK_S_LAST,
		MK_E_CONNECTMANUALLY,
		MK_E_EXCEEDEDDEADLINE,
		MK_E_NEEDGENERIC,
		MK_E_UNAVAILABLE,
		MK_E_SYNTAX,
		MK_E_NOOBJECT,
		MK_E_INVALIDEXTENSION,
		MK_E_INTERMEDIATEINTERFACENOTSUPPORTED,
		MK_E_NOTBINDABLE,
		MK_E_NOTBOUND,
		MK_E_CANTOPENFILE,
		MK_E_MUSTBOTHERUSER,
		MK_E_NOINVERSE,
		MK_E_NOSTORAGE,
		MK_E_NOPREFIX,
		MK_E_ENUMERATION_FAILED,
		CO_E_FIRST,
		CO_E_LAST,
		CO_S_FIRST,
		CO_S_LAST,
		CO_E_NOTINITIALIZED,
		CO_E_ALREADYINITIALIZED,
		CO_E_CANTDETERMINECLASS,
		CO_E_CLASSSTRING,
		CO_E_IIDSTRING,
		CO_E_APPNOTFOUND,
		CO_E_APPSINGLEUSE,
		CO_E_ERRORINAPP,
		CO_E_DLLNOTFOUND,
		CO_E_ERRORINDLL,
		CO_E_WRONGOSFORAPP,
		CO_E_OBJNOTREG,
		CO_E_OBJISREG,
		CO_E_OBJNOTCONNECTED,
		CO_E_APPDIDNTREG,
		CO_E_RELEASED,
		EVENT_E_FIRST,
		EVENT_E_LAST,
		EVENT_S_FIRST,
		EVENT_S_LAST,
		EVENT_S_SOME_SUBSCRIBERS_FAILED,
		EVENT_E_ALL_SUBSCRIBERS_FAILED,
		EVENT_S_NOSUBSCRIBERS,
		EVENT_E_QUERYSYNTAX,
		EVENT_E_QUERYFIELD,
		EVENT_E_INTERNALEXCEPTION,
		EVENT_E_INTERNALERROR,
		EVENT_E_INVALID_PER_USER_SID,
		EVENT_E_USER_EXCEPTION,
		EVENT_E_TOO_MANY_METHODS,
		EVENT_E_MISSING_EVENTCLASS,
		EVENT_E_NOT_ALL_REMOVED,
		EVENT_E_COMPLUS_NOT_INSTALLED,
		EVENT_E_CANT_MODIFY_OR_DELETE_UNCONFIGURED_OBJECT,
		EVENT_E_CANT_MODIFY_OR_DELETE_CONFIGURED_OBJECT,
		EVENT_E_INVALID_EVENT_CLASS_PARTITION,
		EVENT_E_PER_USER_SID_NOT_LOGGED_ON,
		XACT_E_FIRST,
		XACT_E_LAST,
		XACT_S_FIRST,
		XACT_S_LAST,
		XACT_E_ALREADYOTHERSINGLEPHASE,
		XACT_E_CANTRETAIN,
		XACT_E_COMMITFAILED,
		XACT_E_COMMITPREVENTED,
		XACT_E_HEURISTICABORT,
		XACT_E_HEURISTICCOMMIT,
		XACT_E_HEURISTICDAMAGE,
		XACT_E_HEURISTICDANGER,
		XACT_E_ISOLATIONLEVEL,
		XACT_E_NOASYNC,
		XACT_E_NOENLIST,
		XACT_E_NOISORETAIN,
		XACT_E_NORESOURCE,
		XACT_E_NOTCURRENT,
		XACT_E_NOTRANSACTION,
		XACT_E_NOTSUPPORTED,
		XACT_E_UNKNOWNRMGRID,
		XACT_E_WRONGSTATE,
		XACT_E_WRONGUOW,
		XACT_E_XTIONEXISTS,
		XACT_E_NOIMPORTOBJECT,
		XACT_E_INVALIDCOOKIE,
		XACT_E_INDOUBT,
		XACT_E_NOTIMEOUT,
		XACT_E_ALREADYINPROGRESS,
		XACT_E_ABORTED,
		XACT_E_LOGFULL,
		XACT_E_TMNOTAVAILABLE,
		XACT_E_CONNECTION_DOWN,
		XACT_E_CONNECTION_DENIED,
		XACT_E_REENLISTTIMEOUT,
		XACT_E_TIP_CONNECT_FAILED,
		XACT_E_TIP_PROTOCOL_ERROR,
		XACT_E_TIP_PULL_FAILED,
		XACT_E_DEST_TMNOTAVAILABLE,
		XACT_E_TIP_DISABLED,
		XACT_E_NETWORK_TX_DISABLED,
		XACT_E_PARTNER_NETWORK_TX_DISABLED,
		XACT_E_XA_TX_DISABLED,
		XACT_E_UNABLE_TO_READ_DTC_CONFIG,
		XACT_E_UNABLE_TO_LOAD_DTC_PROXY,
		XACT_E_ABORTING,
		XACT_E_PUSH_COMM_FAILURE,
		XACT_E_PULL_COMM_FAILURE,
		XACT_E_LU_TX_DISABLED,
		XACT_E_CLERKNOTFOUND,
		XACT_E_CLERKEXISTS,
		XACT_E_RECOVERYINPROGRESS,
		XACT_E_TRANSACTIONCLOSED,
		XACT_E_INVALIDLSN,
		XACT_E_REPLAYREQUEST,
		XACT_S_ASYNC,
		XACT_S_DEFECT,
		XACT_S_READONLY,
		XACT_S_SOMENORETAIN,
		XACT_S_OKINFORM,
		XACT_S_MADECHANGESCONTENT,
		XACT_S_MADECHANGESINFORM,
		XACT_S_ALLNORETAIN,
		XACT_S_ABORTING,
		XACT_S_SINGLEPHASE,
		XACT_S_LOCALLY_OK,
		XACT_S_LASTRESOURCEMANAGER,
		CONTEXT_E_FIRST,
		CONTEXT_E_LAST,
		CONTEXT_S_FIRST,
		CONTEXT_S_LAST,
		CONTEXT_E_ABORTED,
		CONTEXT_E_ABORTING,
		CONTEXT_E_NOCONTEXT,
		CONTEXT_E_WOULD_DEADLOCK,
		CONTEXT_E_SYNCH_TIMEOUT,
		CONTEXT_E_OLDREF,
		CONTEXT_E_ROLENOTFOUND,
		CONTEXT_E_TMNOTAVAILABLE,
		CO_E_ACTIVATIONFAILED,
		CO_E_ACTIVATIONFAILED_EVENTLOGGED,
		CO_E_ACTIVATIONFAILED_CATALOGERROR,
		CO_E_ACTIVATIONFAILED_TIMEOUT,
		CO_E_INITIALIZATIONFAILED,
		CONTEXT_E_NOJIT,
		CONTEXT_E_NOTRANSACTION,
		CO_E_THREADINGMODEL_CHANGED,
		CO_E_NOIISINTRINSICS,
		CO_E_NOCOOKIES,
		CO_E_DBERROR,
		CO_E_NOTPOOLED,
		CO_E_NOTCONSTRUCTED,
		CO_E_NOSYNCHRONIZATION,
		CO_E_ISOLEVELMISMATCH,
		CO_E_CALL_OUT_OF_TX_SCOPE_NOT_ALLOWED,
		CO_E_EXIT_TRANSACTION_SCOPE_NOT_CALLED,
		OLE_S_USEREG,
		OLE_S_STATIC,
		OLE_S_MAC_CLIPFORMAT,
		DRAGDROP_S_DROP,
		DRAGDROP_S_CANCEL,
		DRAGDROP_S_USEDEFAULTCURSORS,
		DATA_S_SAMEFORMATETC,
		VIEW_S_ALREADY_FROZEN,
		CACHE_S_FORMATETC_NOTSUPPORTED,
		CACHE_S_SAMECACHE,
		CACHE_S_SOMECACHES_NOTUPDATED,
		OLEOBJ_S_INVALIDVERB,
		OLEOBJ_S_CANNOT_DOVERB_NOW,
		OLEOBJ_S_INVALIDHWND,
		INPLACE_S_TRUNCATED,
		CONVERT10_S_NO_PRESENTATION,
		MK_S_REDUCED_TO_SELF,
		MK_S_ME,
		MK_S_HIM,
		MK_S_US,
		MK_S_MONIKERALREADYREGISTERED,
		SCHED_S_TASK_READY,
		SCHED_S_TASK_RUNNING,
		SCHED_S_TASK_DISABLED,
		SCHED_S_TASK_HAS_NOT_RUN,
		SCHED_S_TASK_NO_MORE_RUNS,
		SCHED_S_TASK_NOT_SCHEDULED,
		SCHED_S_TASK_TERMINATED,
		SCHED_S_TASK_NO_VALID_TRIGGERS,
		SCHED_S_EVENT_TRIGGER,
		SCHED_E_TRIGGER_NOT_FOUND,
		SCHED_E_TASK_NOT_READY,
		SCHED_E_TASK_NOT_RUNNING,
		SCHED_E_SERVICE_NOT_INSTALLED,
		SCHED_E_CANNOT_OPEN_TASK,
		SCHED_E_INVALID_TASK,
		SCHED_E_ACCOUNT_INFORMATION_NOT_SET,
		SCHED_E_ACCOUNT_NAME_NOT_FOUND,
		SCHED_E_ACCOUNT_DBASE_CORRUPT,
		SCHED_E_NO_SECURITY_SERVICES,
		SCHED_E_UNKNOWN_OBJECT_VERSION,
		SCHED_E_UNSUPPORTED_ACCOUNT_OPTION,
		SCHED_E_SERVICE_NOT_RUNNING,
		SCHED_E_UNEXPECTEDNODE,
		SCHED_E_NAMESPACE,
		SCHED_E_INVALIDVALUE,
		SCHED_E_MISSINGNODE,
		SCHED_E_MALFORMEDXML,
		SCHED_S_SOME_TRIGGERS_FAILED,
		SCHED_S_BATCH_LOGON_PROBLEM,
		SCHED_E_TOO_MANY_NODES,
		SCHED_E_PAST_END_BOUNDARY,
		SCHED_E_ALREADY_RUNNING,
		SCHED_E_USER_NOT_LOGGED_ON,
		SCHED_E_INVALID_TASK_HASH,
		SCHED_E_SERVICE_NOT_AVAILABLE,
		SCHED_E_SERVICE_TOO_BUSY,
		SCHED_E_TASK_ATTEMPTED,
		SCHED_S_TASK_QUEUED,
		SCHED_E_TASK_DISABLED,
		SCHED_E_TASK_NOT_V1_COMPAT,
		SCHED_E_START_ON_DEMAND,
		CO_E_CLASS_CREATE_FAILED,
		CO_E_SCM_ERROR,
		CO_E_SCM_RPC_FAILURE,
		CO_E_BAD_PATH,
		CO_E_SERVER_EXEC_FAILURE,
		CO_E_OBJSRV_RPC_FAILURE,
		MK_E_NO_NORMALIZED,
		CO_E_SERVER_STOPPING,
		MEM_E_INVALID_ROOT,
		MEM_E_INVALID_LINK,
		MEM_E_INVALID_SIZE,
		CO_S_NOTALLINTERFACES,
		CO_S_MACHINENAMENOTFOUND,
		CO_E_MISSING_DISPLAYNAME,
		CO_E_RUNAS_VALUE_MUST_BE_AAA,
		CO_E_ELEVATION_DISABLED,
		DISP_E_UNKNOWNINTERFACE,
		DISP_E_MEMBERNOTFOUND,
		DISP_E_PARAMNOTFOUND,
		DISP_E_TYPEMISMATCH,
		DISP_E_UNKNOWNNAME,
		DISP_E_NONAMEDARGS,
		DISP_E_BADVARTYPE,
		DISP_E_EXCEPTION,
		DISP_E_OVERFLOW,
		DISP_E_BADINDEX,
		DISP_E_UNKNOWNLCID,
		DISP_E_ARRAYISLOCKED,
		DISP_E_BADPARAMCOUNT,
		DISP_E_PARAMNOTOPTIONAL,
		DISP_E_BADCALLEE,
		DISP_E_NOTACOLLECTION,
		DISP_E_DIVBYZERO,
		DISP_E_BUFFERTOOSMALL,
		TYPE_E_BUFFERTOOSMALL,
		TYPE_E_FIELDNOTFOUND,
		TYPE_E_INVDATAREAD,
		TYPE_E_UNSUPFORMAT,
		TYPE_E_REGISTRYACCESS,
		TYPE_E_LIBNOTREGISTERED,
		TYPE_E_UNDEFINEDTYPE,
		TYPE_E_QUALIFIEDNAMEDISALLOWED,
		TYPE_E_INVALIDSTATE,
		TYPE_E_WRONGTYPEKIND,
		TYPE_E_ELEMENTNOTFOUND,
		TYPE_E_AMBIGUOUSNAME,
		TYPE_E_NAMECONFLICT,
		TYPE_E_UNKNOWNLCID,
		TYPE_E_DLLFUNCTIONNOTFOUND,
		TYPE_E_BADMODULEKIND,
		TYPE_E_SIZETOOBIG,
		TYPE_E_DUPLICATEID,
		TYPE_E_INVALIDID,
		TYPE_E_TYPEMISMATCH,
		TYPE_E_OUTOFBOUNDS,
		TYPE_E_IOERROR,
		TYPE_E_CANTCREATETMPFILE,
		TYPE_E_CANTLOADLIBRARY,
		TYPE_E_INCONSISTENTPROPFUNCS,
		TYPE_E_CIRCULARTYPE,
		STG_E_INVALIDFUNCTION,
		STG_E_FILENOTFOUND,
		STG_E_PATHNOTFOUND,
		STG_E_TOOMANYOPENFILES,
		STG_E_ACCESSDENIED,
		STG_E_INVALIDHANDLE,
		STG_E_INSUFFICIENTMEMORY,
		STG_E_INVALIDPOINTER,
		STG_E_NOMOREFILES,
		STG_E_DISKISWRITEPROTECTED,
		STG_E_SEEKERROR,
		STG_E_WRITEFAULT,
		STG_E_READFAULT,
		STG_E_SHAREVIOLATION,
		STG_E_LOCKVIOLATION,
		STG_E_FILEALREADYEXISTS,
		STG_E_INVALIDPARAMETER,
		STG_E_MEDIUMFULL,
		STG_E_PROPSETMISMATCHED,
		STG_E_ABNORMALAPIEXIT,
		STG_E_INVALIDHEADER,
		STG_E_INVALIDNAME,
		STG_E_UNKNOWN,
		STG_E_UNIMPLEMENTEDFUNCTION,
		STG_E_INVALIDFLAG,
		STG_E_INUSE,
		STG_E_NOTCURRENT,
		STG_E_REVERTED,
		STG_E_CANTSAVE,
		STG_E_OLDFORMAT,
		STG_E_OLDDLL,
		STG_E_SHAREREQUIRED,
		STG_E_NOTFILEBASEDSTORAGE,
		STG_E_EXTANTMARSHALLINGS,
		STG_E_DOCFILECORRUPT,
		STG_E_BADBASEADDRESS,
		STG_E_DOCFILETOOLARGE,
		STG_E_NOTSIMPLEFORMAT,
		STG_E_INCOMPLETE,
		STG_E_TERMINATED,
		STG_S_CONVERTED,
		STG_S_BLOCK,
		STG_S_RETRYNOW,
		STG_S_MONITORING,
		STG_S_MULTIPLEOPENS,
		STG_S_CONSOLIDATIONFAILED,
		STG_S_CANNOTCONSOLIDATE,
		STG_E_STATUS_COPY_PROTECTION_FAILURE,
		STG_E_CSS_AUTHENTICATION_FAILURE,
		STG_E_CSS_KEY_NOT_PRESENT,
		STG_E_CSS_KEY_NOT_ESTABLISHED,
		STG_E_CSS_SCRAMBLED_SECTOR,
		STG_E_CSS_REGION_MISMATCH,
		STG_E_RESETS_EXHAUSTED,
		RPC_E_CALL_REJECTED,
		RPC_E_CALL_CANCELED,
		RPC_E_CANTPOST_INSENDCALL,
		RPC_E_CANTCALLOUT_INASYNCCALL,
		RPC_E_CANTCALLOUT_INEXTERNALCALL,
		RPC_E_CONNECTION_TERMINATED,
		RPC_E_SERVER_DIED,
		RPC_E_CLIENT_DIED,
		RPC_E_INVALID_DATAPACKET,
		RPC_E_CANTTRANSMIT_CALL,
		RPC_E_CLIENT_CANTMARSHAL_DATA,
		RPC_E_CLIENT_CANTUNMARSHAL_DATA,
		RPC_E_SERVER_CANTMARSHAL_DATA,
		RPC_E_SERVER_CANTUNMARSHAL_DATA,
		RPC_E_INVALID_DATA,
		RPC_E_INVALID_PARAMETER,
		RPC_E_CANTCALLOUT_AGAIN,
		RPC_E_SERVER_DIED_DNE,
		RPC_E_SYS_CALL_FAILED,
		RPC_E_OUT_OF_RESOURCES,
		RPC_E_ATTEMPTED_MULTITHREAD,
		RPC_E_NOT_REGISTERED,
		RPC_E_FAULT,
		RPC_E_SERVERFAULT,
		RPC_E_CHANGED_MODE,
		RPC_E_INVALIDMETHOD,
		RPC_E_DISCONNECTED,
		RPC_E_RETRY,
		RPC_E_SERVERCALL_RETRYLATER,
		RPC_E_SERVERCALL_REJECTED,
		RPC_E_INVALID_CALLDATA,
		RPC_E_CANTCALLOUT_ININPUTSYNCCALL,
		RPC_E_WRONG_THREAD,
		RPC_E_THREAD_NOT_INIT,
		RPC_E_VERSION_MISMATCH,
		RPC_E_INVALID_HEADER,
		RPC_E_INVALID_EXTENSION,
		RPC_E_INVALID_IPID,
		RPC_E_INVALID_OBJECT,
		RPC_S_CALLPENDING,
		RPC_S_WAITONTIMER,
		RPC_E_CALL_COMPLETE,
		RPC_E_UNSECURE_CALL,
		RPC_E_TOO_LATE,
		RPC_E_NO_GOOD_SECURITY_PACKAGES,
		RPC_E_ACCESS_DENIED,
		RPC_E_REMOTE_DISABLED,
		RPC_E_INVALID_OBJREF,
		RPC_E_NO_CONTEXT,
		RPC_E_TIMEOUT,
		RPC_E_NO_SYNC,
		RPC_E_FULLSIC_REQUIRED,
		RPC_E_INVALID_STD_NAME,
		CO_E_FAILEDTOIMPERSONATE,
		CO_E_FAILEDTOGETSECCTX,
		CO_E_FAILEDTOOPENTHREADTOKEN,
		CO_E_FAILEDTOGETTOKENINFO,
		CO_E_TRUSTEEDOESNTMATCHCLIENT,
		CO_E_FAILEDTOQUERYCLIENTBLANKET,
		CO_E_FAILEDTOSETDACL,
		CO_E_ACCESSCHECKFAILED,
		CO_E_NETACCESSAPIFAILED,
		CO_E_WRONGTRUSTEENAMESYNTAX,
		CO_E_INVALIDSID,
		CO_E_CONVERSIONFAILED,
		CO_E_NOMATCHINGSIDFOUND,
		CO_E_LOOKUPACCSIDFAILED,
		CO_E_NOMATCHINGNAMEFOUND,
		CO_E_LOOKUPACCNAMEFAILED,
		CO_E_SETSERLHNDLFAILED,
		CO_E_FAILEDTOGETWINDIR,
		CO_E_PATHTOOLONG,
		CO_E_FAILEDTOGENUUID,
		CO_E_FAILEDTOCREATEFILE,
		CO_E_FAILEDTOCLOSEHANDLE,
		CO_E_EXCEEDSYSACLLIMIT,
		CO_E_ACESINWRONGORDER,
		CO_E_INCOMPATIBLESTREAMVERSION,
		CO_E_FAILEDTOOPENPROCESSTOKEN,
		CO_E_DECODEFAILED,
		CO_E_ACNOTINITIALIZED,
		CO_E_CANCEL_DISABLED,
		RPC_E_UNEXPECTED,
		ERROR_AUDITING_DISABLED,
		ERROR_ALL_SIDS_FILTERED,
		ERROR_BIZRULES_NOT_ENABLED,
		NTE_BAD_UID,
		NTE_BAD_HASH,
		NTE_BAD_KEY,
		NTE_BAD_LEN,
		NTE_BAD_DATA,
		NTE_BAD_SIGNATURE,
		NTE_BAD_VER,
		NTE_BAD_ALGID,
		NTE_BAD_FLAGS,
		NTE_BAD_TYPE,
		NTE_BAD_KEY_STATE,
		NTE_BAD_HASH_STATE,
		NTE_NO_KEY,
		NTE_NO_MEMORY,
		NTE_EXISTS,
		NTE_PERM,
		NTE_NOT_FOUND,
		NTE_DOUBLE_ENCRYPT,
		NTE_BAD_PROVIDER,
		NTE_BAD_PROV_TYPE,
		NTE_BAD_PUBLIC_KEY,
		NTE_BAD_KEYSET,
		NTE_PROV_TYPE_NOT_DEF,
		NTE_PROV_TYPE_ENTRY_BAD,
		NTE_KEYSET_NOT_DEF,
		NTE_KEYSET_ENTRY_BAD,
		NTE_PROV_TYPE_NO_MATCH,
		NTE_SIGNATURE_FILE_BAD,
		NTE_PROVIDER_DLL_FAIL,
		NTE_PROV_DLL_NOT_FOUND,
		NTE_BAD_KEYSET_PARAM,
		NTE_FAIL,
		NTE_SYS_ERR,
		NTE_SILENT_CONTEXT,
		NTE_TOKEN_KEYSET_STORAGE_FULL,
		NTE_TEMPORARY_PROFILE,
		NTE_FIXEDPARAMETER,
		NTE_INVALID_HANDLE,
		NTE_INVALID_PARAMETER,
		NTE_BUFFER_TOO_SMALL,
		NTE_NOT_SUPPORTED,
		NTE_NO_MORE_ITEMS,
		NTE_BUFFERS_OVERLAP,
		NTE_DECRYPTION_FAILURE,
		NTE_INTERNAL_ERROR,
		NTE_UI_REQUIRED,
		NTE_HMAC_NOT_SUPPORTED,
		SEC_E_INSUFFICIENT_MEMORY,
		SEC_E_INVALID_HANDLE,
		SEC_E_UNSUPPORTED_FUNCTION,
		SEC_E_TARGET_UNKNOWN,
		SEC_E_INTERNAL_ERROR,
		SEC_E_SECPKG_NOT_FOUND,
		SEC_E_NOT_OWNER,
		SEC_E_CANNOT_INSTALL,
		SEC_E_INVALID_TOKEN,
		SEC_E_CANNOT_PACK,
		SEC_E_QOP_NOT_SUPPORTED,
		SEC_E_NO_IMPERSONATION,
		SEC_E_LOGON_DENIED,
		SEC_E_UNKNOWN_CREDENTIALS,
		SEC_E_NO_CREDENTIALS,
		SEC_E_MESSAGE_ALTERED,
		SEC_E_OUT_OF_SEQUENCE,
		SEC_E_NO_AUTHENTICATING_AUTHORITY,
		SEC_I_CONTINUE_NEEDED,
		SEC_I_COMPLETE_NEEDED,
		SEC_I_COMPLETE_AND_CONTINUE,
		SEC_I_LOCAL_LOGON,
		SEC_E_BAD_PKGID,
		SEC_E_CONTEXT_EXPIRED,
		SEC_I_CONTEXT_EXPIRED,
		SEC_E_INCOMPLETE_MESSAGE,
		SEC_E_INCOMPLETE_CREDENTIALS,
		SEC_E_BUFFER_TOO_SMALL,
		SEC_I_INCOMPLETE_CREDENTIALS,
		SEC_I_RENEGOTIATE,
		SEC_E_WRONG_PRINCIPAL,
		SEC_I_NO_LSA_CONTEXT,
		SEC_E_TIME_SKEW,
		SEC_E_UNTRUSTED_ROOT,
		SEC_E_ILLEGAL_MESSAGE,
		SEC_E_CERT_UNKNOWN,
		SEC_E_CERT_EXPIRED,
		SEC_E_ENCRYPT_FAILURE,
		SEC_E_DECRYPT_FAILURE,
		SEC_E_ALGORITHM_MISMATCH,
		SEC_E_SECURITY_QOS_FAILED,
		SEC_E_UNFINISHED_CONTEXT_DELETED,
		SEC_E_NO_TGT_REPLY,
		SEC_E_NO_IP_ADDRESSES,
		SEC_E_WRONG_CREDENTIAL_HANDLE,
		SEC_E_CRYPTO_SYSTEM_INVALID,
		SEC_E_MAX_REFERRALS_EXCEEDED,
		SEC_E_MUST_BE_KDC,
		SEC_E_STRONG_CRYPTO_NOT_SUPPORTED,
		SEC_E_TOO_MANY_PRINCIPALS,
		SEC_E_NO_PA_DATA,
		SEC_E_PKINIT_NAME_MISMATCH,
		SEC_E_SMARTCARD_LOGON_REQUIRED,
		SEC_E_SHUTDOWN_IN_PROGRESS,
		SEC_E_KDC_INVALID_REQUEST,
		SEC_E_KDC_UNABLE_TO_REFER,
		SEC_E_KDC_UNKNOWN_ETYPE,
		SEC_E_UNSUPPORTED_PREAUTH,
		SEC_E_DELEGATION_REQUIRED,
		SEC_E_BAD_BINDINGS,
		SEC_E_MULTIPLE_ACCOUNTS,
		SEC_E_NO_KERB_KEY,
		SEC_E_CERT_WRONG_USAGE,
		SEC_E_DOWNGRADE_DETECTED,
		SEC_E_SMARTCARD_CERT_REVOKED,
		SEC_E_ISSUING_CA_UNTRUSTED,
		SEC_E_REVOCATION_OFFLINE_C,
		SEC_E_PKINIT_CLIENT_FAILURE,
		SEC_E_SMARTCARD_CERT_EXPIRED,
		SEC_E_NO_S4U_PROT_SUPPORT,
		SEC_E_CROSSREALM_DELEGATION_FAILURE,
		SEC_E_REVOCATION_OFFLINE_KDC,
		SEC_E_ISSUING_CA_UNTRUSTED_KDC,
		SEC_E_KDC_CERT_EXPIRED,
		SEC_E_KDC_CERT_REVOKED,
		SEC_I_SIGNATURE_NEEDED,
		SEC_E_INVALID_PARAMETER,
		SEC_E_DELEGATION_POLICY,
		SEC_E_POLICY_NLTM_ONLY,
		SEC_I_NO_RENEGOTIATION,
		SEC_E_NO_CONTEXT,
		SEC_E_PKU2U_CERT_FAILURE,
		SEC_E_MUTUAL_AUTH_FAILED,
		CRYPT_E_MSG_ERROR,
		CRYPT_E_UNKNOWN_ALGO,
		CRYPT_E_OID_FORMAT,
		CRYPT_E_INVALID_MSG_TYPE,
		CRYPT_E_UNEXPECTED_ENCODING,
		CRYPT_E_AUTH_ATTR_MISSING,
		CRYPT_E_HASH_VALUE,
		CRYPT_E_INVALID_INDEX,
		CRYPT_E_ALREADY_DECRYPTED,
		CRYPT_E_NOT_DECRYPTED,
		CRYPT_E_RECIPIENT_NOT_FOUND,
		CRYPT_E_CONTROL_TYPE,
		CRYPT_E_ISSUER_SERIALNUMBER,
		CRYPT_E_SIGNER_NOT_FOUND,
		CRYPT_E_ATTRIBUTES_MISSING,
		CRYPT_E_STREAM_MSG_NOT_READY,
		CRYPT_E_STREAM_INSUFFICIENT_DATA,
		CRYPT_I_NEW_PROTECTION_REQUIRED,
		CRYPT_E_BAD_LEN,
		CRYPT_E_BAD_ENCODE,
		CRYPT_E_FILE_ERROR,
		CRYPT_E_NOT_FOUND,
		CRYPT_E_EXISTS,
		CRYPT_E_NO_PROVIDER,
		CRYPT_E_SELF_SIGNED,
		CRYPT_E_DELETED_PREV,
		CRYPT_E_NO_MATCH,
		CRYPT_E_UNEXPECTED_MSG_TYPE,
		CRYPT_E_NO_KEY_PROPERTY,
		CRYPT_E_NO_DECRYPT_CERT,
		CRYPT_E_BAD_MSG,
		CRYPT_E_NO_SIGNER,
		CRYPT_E_PENDING_CLOSE,
		CRYPT_E_REVOKED,
		CRYPT_E_NO_REVOCATION_DLL,
		CRYPT_E_NO_REVOCATION_CHECK,
		CRYPT_E_REVOCATION_OFFLINE,
		CRYPT_E_NOT_IN_REVOCATION_DATABASE,
		CRYPT_E_INVALID_NUMERIC_STRING,
		CRYPT_E_INVALID_PRINTABLE_STRING,
		CRYPT_E_INVALID_IA5_STRING,
		CRYPT_E_INVALID_X500_STRING,
		CRYPT_E_NOT_CHAR_STRING,
		CRYPT_E_FILERESIZED,
		CRYPT_E_SECURITY_SETTINGS,
		CRYPT_E_NO_VERIFY_USAGE_DLL,
		CRYPT_E_NO_VERIFY_USAGE_CHECK,
		CRYPT_E_VERIFY_USAGE_OFFLINE,
		CRYPT_E_NOT_IN_CTL,
		CRYPT_E_NO_TRUSTED_SIGNER,
		CRYPT_E_MISSING_PUBKEY_PARA,
		CRYPT_E_OSS_ERROR,
		OSS_MORE_BUF,
		OSS_NEGATIVE_UINTEGER,
		OSS_PDU_RANGE,
		OSS_MORE_INPUT,
		OSS_DATA_ERROR,
		OSS_BAD_ARG,
		OSS_BAD_VERSION,
		OSS_OUT_MEMORY,
		OSS_PDU_MISMATCH,
		OSS_LIMITED,
		OSS_BAD_PTR,
		OSS_BAD_TIME,
		OSS_INDEFINITE_NOT_SUPPORTED,
		OSS_MEM_ERROR,
		OSS_BAD_TABLE,
		OSS_TOO_LONG,
		OSS_CONSTRAINT_VIOLATED,
		OSS_FATAL_ERROR,
		OSS_ACCESS_SERIALIZATION_ERROR,
		OSS_NULL_TBL,
		OSS_NULL_FCN,
		OSS_BAD_ENCRULES,
		OSS_UNAVAIL_ENCRULES,
		OSS_CANT_OPEN_TRACE_WINDOW,
		OSS_UNIMPLEMENTED,
		OSS_OID_DLL_NOT_LINKED,
		OSS_CANT_OPEN_TRACE_FILE,
		OSS_TRACE_FILE_ALREADY_OPEN,
		OSS_TABLE_MISMATCH,
		OSS_TYPE_NOT_SUPPORTED,
		OSS_REAL_DLL_NOT_LINKED,
		OSS_REAL_CODE_NOT_LINKED,
		OSS_OUT_OF_RANGE,
		OSS_COPIER_DLL_NOT_LINKED,
		OSS_CONSTRAINT_DLL_NOT_LINKED,
		OSS_COMPARATOR_DLL_NOT_LINKED,
		OSS_COMPARATOR_CODE_NOT_LINKED,
		OSS_MEM_MGR_DLL_NOT_LINKED,
		OSS_PDV_DLL_NOT_LINKED,
		OSS_PDV_CODE_NOT_LINKED,
		OSS_API_DLL_NOT_LINKED,
		OSS_BERDER_DLL_NOT_LINKED,
		OSS_PER_DLL_NOT_LINKED,
		OSS_OPEN_TYPE_ERROR,
		OSS_MUTEX_NOT_CREATED,
		OSS_CANT_CLOSE_TRACE_FILE,
		CRYPT_E_ASN1_ERROR,
		CRYPT_E_ASN1_INTERNAL,
		CRYPT_E_ASN1_EOD,
		CRYPT_E_ASN1_CORRUPT,
		CRYPT_E_ASN1_LARGE,
		CRYPT_E_ASN1_CONSTRAINT,
		CRYPT_E_ASN1_MEMORY,
		CRYPT_E_ASN1_OVERFLOW,
		CRYPT_E_ASN1_BADPDU,
		CRYPT_E_ASN1_BADARGS,
		CRYPT_E_ASN1_BADREAL,
		CRYPT_E_ASN1_BADTAG,
		CRYPT_E_ASN1_CHOICE,
		CRYPT_E_ASN1_RULE,
		CRYPT_E_ASN1_UTF8,
		CRYPT_E_ASN1_PDU_TYPE,
		CRYPT_E_ASN1_NYI,
		CRYPT_E_ASN1_EXTENDED,
		CRYPT_E_ASN1_NOEOD,
		CERTSRV_E_BAD_REQUESTSUBJECT,
		CERTSRV_E_NO_REQUEST,
		CERTSRV_E_BAD_REQUESTSTATUS,
		CERTSRV_E_PROPERTY_EMPTY,
		CERTSRV_E_INVALID_CA_CERTIFICATE,
		CERTSRV_E_SERVER_SUSPENDED,
		CERTSRV_E_ENCODING_LENGTH,
		CERTSRV_E_ROLECONFLICT,
		CERTSRV_E_RESTRICTEDOFFICER,
		CERTSRV_E_KEY_ARCHIVAL_NOT_CONFIGURED,
		CERTSRV_E_NO_VALID_KRA,
		CERTSRV_E_BAD_REQUEST_KEY_ARCHIVAL,
		CERTSRV_E_NO_CAADMIN_DEFINED,
		CERTSRV_E_BAD_RENEWAL_CERT_ATTRIBUTE,
		CERTSRV_E_NO_DB_SESSIONS,
		CERTSRV_E_ALIGNMENT_FAULT,
		CERTSRV_E_ENROLL_DENIED,
		CERTSRV_E_TEMPLATE_DENIED,
		CERTSRV_E_DOWNLEVEL_DC_SSL_OR_UPGRADE,
		CERTSRV_E_ADMIN_DENIED_REQUEST,
		CERTSRV_E_NO_POLICY_SERVER,
		CERTSRV_E_UNSUPPORTED_CERT_TYPE,
		CERTSRV_E_NO_CERT_TYPE,
		CERTSRV_E_TEMPLATE_CONFLICT,
		CERTSRV_E_SUBJECT_ALT_NAME_REQUIRED,
		CERTSRV_E_ARCHIVED_KEY_REQUIRED,
		CERTSRV_E_SMIME_REQUIRED,
		CERTSRV_E_BAD_RENEWAL_SUBJECT,
		CERTSRV_E_BAD_TEMPLATE_VERSION,
		CERTSRV_E_TEMPLATE_POLICY_REQUIRED,
		CERTSRV_E_SIGNATURE_POLICY_REQUIRED,
		CERTSRV_E_SIGNATURE_COUNT,
		CERTSRV_E_SIGNATURE_REJECTED,
		CERTSRV_E_ISSUANCE_POLICY_REQUIRED,
		CERTSRV_E_SUBJECT_UPN_REQUIRED,
		CERTSRV_E_SUBJECT_DIRECTORY_GUID_REQUIRED,
		CERTSRV_E_SUBJECT_DNS_REQUIRED,
		CERTSRV_E_ARCHIVED_KEY_UNEXPECTED,
		CERTSRV_E_KEY_LENGTH,
		CERTSRV_E_SUBJECT_EMAIL_REQUIRED,
		CERTSRV_E_UNKNOWN_CERT_TYPE,
		CERTSRV_E_CERT_TYPE_OVERLAP,
		CERTSRV_E_TOO_MANY_SIGNATURES,
		XENROLL_E_KEY_NOT_EXPORTABLE,
		XENROLL_E_CANNOT_ADD_ROOT_CERT,
		XENROLL_E_RESPONSE_KA_HASH_NOT_FOUND,
		XENROLL_E_RESPONSE_UNEXPECTED_KA_HASH,
		XENROLL_E_RESPONSE_KA_HASH_MISMATCH,
		XENROLL_E_KEYSPEC_SMIME_MISMATCH,
		TRUST_E_SYSTEM_ERROR,
		TRUST_E_NO_SIGNER_CERT,
		TRUST_E_COUNTER_SIGNER,
		TRUST_E_CERT_SIGNATURE,
		TRUST_E_TIME_STAMP,
		TRUST_E_BAD_DIGEST,
		TRUST_E_BASIC_CONSTRAINTS,
		TRUST_E_FINANCIAL_CRITERIA,
		MSSIPOTF_E_OUTOFMEMRANGE,
		MSSIPOTF_E_CANTGETOBJECT,
		MSSIPOTF_E_NOHEADTABLE,
		MSSIPOTF_E_BAD_MAGICNUMBER,
		MSSIPOTF_E_BAD_OFFSET_TABLE,
		MSSIPOTF_E_TABLE_TAGORDER,
		MSSIPOTF_E_TABLE_LONGWORD,
		MSSIPOTF_E_BAD_FIRST_TABLE_PLACEMENT,
		MSSIPOTF_E_TABLES_OVERLAP,
		MSSIPOTF_E_TABLE_PADBYTES,
		MSSIPOTF_E_FILETOOSMALL,
		MSSIPOTF_E_TABLE_CHECKSUM,
		MSSIPOTF_E_FILE_CHECKSUM,
		MSSIPOTF_E_FAILED_POLICY,
		MSSIPOTF_E_FAILED_HINTS_CHECK,
		MSSIPOTF_E_NOT_OPENTYPE,
		MSSIPOTF_E_FILE,
		MSSIPOTF_E_CRYPT,
		MSSIPOTF_E_BADVERSION,
		MSSIPOTF_E_DSIG_STRUCTURE,
		MSSIPOTF_E_PCONST_CHECK,
		MSSIPOTF_E_STRUCTURE,
		ERROR_CRED_REQUIRES_CONFIRMATION,
		NTE_OP_OK,
		TRUST_E_PROVIDER_UNKNOWN,
		TRUST_E_ACTION_UNKNOWN,
		TRUST_E_SUBJECT_FORM_UNKNOWN,
		TRUST_E_SUBJECT_NOT_TRUSTED,
		DIGSIG_E_ENCODE,
		DIGSIG_E_DECODE,
		DIGSIG_E_EXTENSIBILITY,
		DIGSIG_E_CRYPTO,
		PERSIST_E_SIZEDEFINITE,
		PERSIST_E_SIZEINDEFINITE,
		PERSIST_E_NOTSELFSIZING,
		TRUST_E_NOSIGNATURE,
		CERT_E_EXPIRED,
		CERT_E_VALIDITYPERIODNESTING,
		CERT_E_ROLE,
		CERT_E_PATHLENCONST,
		CERT_E_CRITICAL,
		CERT_E_PURPOSE,
		CERT_E_ISSUERCHAINING,
		CERT_E_MALFORMED,
		CERT_E_UNTRUSTEDROOT,
		CERT_E_CHAINING,
		TRUST_E_FAIL,
		CERT_E_REVOKED,
		CERT_E_UNTRUSTEDTESTROOT,
		CERT_E_REVOCATION_FAILURE,
		CERT_E_CN_NO_MATCH,
		CERT_E_WRONG_USAGE,
		TRUST_E_EXPLICIT_DISTRUST,
		CERT_E_UNTRUSTEDCA,
		CERT_E_INVALID_POLICY,
		CERT_E_INVALID_NAME,
		SPAPI_E_EXPECTED_SECTION_NAME,
		SPAPI_E_BAD_SECTION_NAME_LINE,
		SPAPI_E_SECTION_NAME_TOO_LONG,
		SPAPI_E_GENERAL_SYNTAX,
		SPAPI_E_WRONG_INF_STYLE,
		SPAPI_E_SECTION_NOT_FOUND,
		SPAPI_E_LINE_NOT_FOUND,
		SPAPI_E_NO_BACKUP,
		SPAPI_E_NO_ASSOCIATED_CLASS,
		SPAPI_E_CLASS_MISMATCH,
		SPAPI_E_DUPLICATE_FOUND,
		SPAPI_E_NO_DRIVER_SELECTED,
		SPAPI_E_KEY_DOES_NOT_EXIST,
		SPAPI_E_INVALID_DEVINST_NAME,
		SPAPI_E_INVALID_CLASS,
		SPAPI_E_DEVINST_ALREADY_EXISTS,
		SPAPI_E_DEVINFO_NOT_REGISTERED,
		SPAPI_E_INVALID_REG_PROPERTY,
		SPAPI_E_NO_INF,
		SPAPI_E_NO_SUCH_DEVINST,
		SPAPI_E_CANT_LOAD_CLASS_ICON,
		SPAPI_E_INVALID_CLASS_INSTALLER,
		SPAPI_E_DI_DO_DEFAULT,
		SPAPI_E_DI_NOFILECOPY,
		SPAPI_E_INVALID_HWPROFILE,
		SPAPI_E_NO_DEVICE_SELECTED,
		SPAPI_E_DEVINFO_LIST_LOCKED,
		SPAPI_E_DEVINFO_DATA_LOCKED,
		SPAPI_E_DI_BAD_PATH,
		SPAPI_E_NO_CLASSINSTALL_PARAMS,
		SPAPI_E_FILEQUEUE_LOCKED,
		SPAPI_E_BAD_SERVICE_INSTALLSECT,
		SPAPI_E_NO_CLASS_DRIVER_LIST,
		SPAPI_E_NO_ASSOCIATED_SERVICE,
		SPAPI_E_NO_DEFAULT_DEVICE_INTERFACE,
		SPAPI_E_DEVICE_INTERFACE_ACTIVE,
		SPAPI_E_DEVICE_INTERFACE_REMOVED,
		SPAPI_E_BAD_INTERFACE_INSTALLSECT,
		SPAPI_E_NO_SUCH_INTERFACE_CLASS,
		SPAPI_E_INVALID_REFERENCE_STRING,
		SPAPI_E_INVALID_MACHINENAME,
		SPAPI_E_REMOTE_COMM_FAILURE,
		SPAPI_E_MACHINE_UNAVAILABLE,
		SPAPI_E_NO_CONFIGMGR_SERVICES,
		SPAPI_E_INVALID_PROPPAGE_PROVIDER,
		SPAPI_E_NO_SUCH_DEVICE_INTERFACE,
		SPAPI_E_DI_POSTPROCESSING_REQUIRED,
		SPAPI_E_INVALID_COINSTALLER,
		SPAPI_E_NO_COMPAT_DRIVERS,
		SPAPI_E_NO_DEVICE_ICON,
		SPAPI_E_INVALID_INF_LOGCONFIG,
		SPAPI_E_DI_DONT_INSTALL,
		SPAPI_E_INVALID_FILTER_DRIVER,
		SPAPI_E_NON_WINDOWS_NT_DRIVER,
		SPAPI_E_NON_WINDOWS_DRIVER,
		SPAPI_E_NO_CATALOG_FOR_OEM_INF,
		SPAPI_E_DEVINSTALL_QUEUE_NONNATIVE,
		SPAPI_E_NOT_DISABLEABLE,
		SPAPI_E_CANT_REMOVE_DEVINST,
		SPAPI_E_INVALID_TARGET,
		SPAPI_E_DRIVER_NONNATIVE,
		SPAPI_E_IN_WOW64,
		SPAPI_E_SET_SYSTEM_RESTORE_POINT,
		SPAPI_E_INCORRECTLY_COPIED_INF,
		SPAPI_E_SCE_DISABLED,
		SPAPI_E_UNKNOWN_EXCEPTION,
		SPAPI_E_PNP_REGISTRY_ERROR,
		SPAPI_E_REMOTE_REQUEST_UNSUPPORTED,
		SPAPI_E_NOT_AN_INSTALLED_OEM_INF,
		SPAPI_E_INF_IN_USE_BY_DEVICES,
		SPAPI_E_DI_FUNCTION_OBSOLETE,
		SPAPI_E_NO_AUTHENTICODE_CATALOG,
		SPAPI_E_AUTHENTICODE_DISALLOWED,
		SPAPI_E_AUTHENTICODE_TRUSTED_PUBLISHER,
		SPAPI_E_AUTHENTICODE_TRUST_NOT_ESTABLISHED,
		SPAPI_E_AUTHENTICODE_PUBLISHER_NOT_TRUSTED,
		SPAPI_E_SIGNATURE_OSATTRIBUTE_MISMATCH,
		SPAPI_E_ONLY_VALIDATE_VIA_AUTHENTICODE,
		SPAPI_E_DEVICE_INSTALLER_NOT_READY,
		SPAPI_E_DRIVER_STORE_ADD_FAILED,
		SPAPI_E_DEVICE_INSTALL_BLOCKED,
		SPAPI_E_DRIVER_INSTALL_BLOCKED,
		SPAPI_E_WRONG_INF_TYPE,
		SPAPI_E_FILE_HASH_NOT_IN_CATALOG,
		SPAPI_E_DRIVER_STORE_DELETE_FAILED,
		SPAPI_E_UNRECOVERABLE_STACK_OVERFLOW,
		SPAPI_E_ERROR_NOT_INSTALLED,
		SCARD_F_INTERNAL_ERROR,
		SCARD_E_CANCELLED,
		SCARD_E_INVALID_HANDLE,
		SCARD_E_INVALID_PARAMETER,
		SCARD_E_INVALID_TARGET,
		SCARD_E_NO_MEMORY,
		SCARD_F_WAITED_TOO_LONG,
		SCARD_E_INSUFFICIENT_BUFFER,
		SCARD_E_UNKNOWN_READER,
		SCARD_E_TIMEOUT,
		SCARD_E_SHARING_VIOLATION,
		SCARD_E_NO_SMARTCARD,
		SCARD_E_UNKNOWN_CARD,
		SCARD_E_CANT_DISPOSE,
		SCARD_E_PROTO_MISMATCH,
		SCARD_E_NOT_READY,
		SCARD_E_INVALID_VALUE,
		SCARD_E_SYSTEM_CANCELLED,
		SCARD_F_COMM_ERROR,
		SCARD_F_UNKNOWN_ERROR,
		SCARD_E_INVALID_ATR,
		SCARD_E_NOT_TRANSACTED,
		SCARD_E_READER_UNAVAILABLE,
		SCARD_P_SHUTDOWN,
		SCARD_E_PCI_TOO_SMALL,
		SCARD_E_READER_UNSUPPORTED,
		SCARD_E_DUPLICATE_READER,
		SCARD_E_CARD_UNSUPPORTED,
		SCARD_E_NO_SERVICE,
		SCARD_E_SERVICE_STOPPED,
		SCARD_E_UNEXPECTED,
		SCARD_E_ICC_INSTALLATION,
		SCARD_E_ICC_CREATEORDER,
		SCARD_E_UNSUPPORTED_FEATURE,
		SCARD_E_DIR_NOT_FOUND,
		SCARD_E_FILE_NOT_FOUND,
		SCARD_E_NO_DIR,
		SCARD_E_NO_FILE,
		SCARD_E_NO_ACCESS,
		SCARD_E_WRITE_TOO_MANY,
		SCARD_E_BAD_SEEK,
		SCARD_E_INVALID_CHV,
		SCARD_E_UNKNOWN_RES_MNG,
		SCARD_E_NO_SUCH_CERTIFICATE,
		SCARD_E_CERTIFICATE_UNAVAILABLE,
		SCARD_E_NO_READERS_AVAILABLE,
		SCARD_E_COMM_DATA_LOST,
		SCARD_E_NO_KEY_CONTAINER,
		SCARD_E_SERVER_TOO_BUSY,
		SCARD_E_PIN_CACHE_EXPIRED,
		SCARD_E_NO_PIN_CACHE,
		SCARD_E_READ_ONLY_CARD,
		SCARD_W_UNSUPPORTED_CARD,
		SCARD_W_UNRESPONSIVE_CARD,
		SCARD_W_UNPOWERED_CARD,
		SCARD_W_RESET_CARD,
		SCARD_W_REMOVED_CARD,
		SCARD_W_SECURITY_VIOLATION,
		SCARD_W_WRONG_CHV,
		SCARD_W_CHV_BLOCKED,
		SCARD_W_EOF,
		SCARD_W_CANCELLED_BY_USER,
		SCARD_W_CARD_NOT_AUTHENTICATED,
		SCARD_W_CACHE_ITEM_NOT_FOUND,
		SCARD_W_CACHE_ITEM_STALE,
		SCARD_W_CACHE_ITEM_TOO_BIG,
		COMADMIN_E_OBJECTERRORS,
		COMADMIN_E_OBJECTINVALID,
		COMADMIN_E_KEYMISSING,
		COMADMIN_E_ALREADYINSTALLED,
		COMADMIN_E_APP_FILE_WRITEFAIL,
		COMADMIN_E_APP_FILE_READFAIL,
		COMADMIN_E_APP_FILE_VERSION,
		COMADMIN_E_BADPATH,
		COMADMIN_E_APPLICATIONEXISTS,
		COMADMIN_E_ROLEEXISTS,
		COMADMIN_E_CANTCOPYFILE,
		COMADMIN_E_NOUSER,
		COMADMIN_E_INVALIDUSERIDS,
		COMADMIN_E_NOREGISTRYCLSID,
		COMADMIN_E_BADREGISTRYPROGID,
		COMADMIN_E_AUTHENTICATIONLEVEL,
		COMADMIN_E_USERPASSWDNOTVALID,
		COMADMIN_E_CLSIDORIIDMISMATCH,
		COMADMIN_E_REMOTEINTERFACE,
		COMADMIN_E_DLLREGISTERSERVER,
		COMADMIN_E_NOSERVERSHARE,
		COMADMIN_E_DLLLOADFAILED,
		COMADMIN_E_BADREGISTRYLIBID,
		COMADMIN_E_APPDIRNOTFOUND,
		COMADMIN_E_REGISTRARFAILED,
		COMADMIN_E_COMPFILE_DOESNOTEXIST,
		COMADMIN_E_COMPFILE_LOADDLLFAIL,
		COMADMIN_E_COMPFILE_GETCLASSOBJ,
		COMADMIN_E_COMPFILE_CLASSNOTAVAIL,
		COMADMIN_E_COMPFILE_BADTLB,
		COMADMIN_E_COMPFILE_NOTINSTALLABLE,
		COMADMIN_E_NOTCHANGEABLE,
		COMADMIN_E_NOTDELETEABLE,
		COMADMIN_E_SESSION,
		COMADMIN_E_COMP_MOVE_LOCKED,
		COMADMIN_E_COMP_MOVE_BAD_DEST,
		COMADMIN_E_REGISTERTLB,
		COMADMIN_E_SYSTEMAPP,
		COMADMIN_E_COMPFILE_NOREGISTRAR,
		COMADMIN_E_COREQCOMPINSTALLED,
		COMADMIN_E_SERVICENOTINSTALLED,
		COMADMIN_E_PROPERTYSAVEFAILED,
		COMADMIN_E_OBJECTEXISTS,
		COMADMIN_E_COMPONENTEXISTS,
		COMADMIN_E_REGFILE_CORRUPT,
		COMADMIN_E_PROPERTY_OVERFLOW,
		COMADMIN_E_NOTINREGISTRY,
		COMADMIN_E_OBJECTNOTPOOLABLE,
		COMADMIN_E_APPLID_MATCHES_CLSID,
		COMADMIN_E_ROLE_DOES_NOT_EXIST,
		COMADMIN_E_START_APP_NEEDS_COMPONENTS,
		COMADMIN_E_REQUIRES_DIFFERENT_PLATFORM,
		COMADMIN_E_CAN_NOT_EXPORT_APP_PROXY,
		COMADMIN_E_CAN_NOT_START_APP,
		COMADMIN_E_CAN_NOT_EXPORT_SYS_APP,
		COMADMIN_E_CANT_SUBSCRIBE_TO_COMPONENT,
		COMADMIN_E_EVENTCLASS_CANT_BE_SUBSCRIBER,
		COMADMIN_E_LIB_APP_PROXY_INCOMPATIBLE,
		COMADMIN_E_BASE_PARTITION_ONLY,
		COMADMIN_E_START_APP_DISABLED,
		COMADMIN_E_CAT_DUPLICATE_PARTITION_NAME,
		COMADMIN_E_CAT_INVALID_PARTITION_NAME,
		COMADMIN_E_CAT_PARTITION_IN_USE,
		COMADMIN_E_FILE_PARTITION_DUPLICATE_FILES,
		COMADMIN_E_CAT_IMPORTED_COMPONENTS_NOT_ALLOWED,
		COMADMIN_E_AMBIGUOUS_APPLICATION_NAME,
		COMADMIN_E_AMBIGUOUS_PARTITION_NAME,
		COMADMIN_E_REGDB_NOTINITIALIZED,
		COMADMIN_E_REGDB_NOTOPEN,
		COMADMIN_E_REGDB_SYSTEMERR,
		COMADMIN_E_REGDB_ALREADYRUNNING,
		COMADMIN_E_MIG_VERSIONNOTSUPPORTED,
		COMADMIN_E_MIG_SCHEMANOTFOUND,
		COMADMIN_E_CAT_BITNESSMISMATCH,
		COMADMIN_E_CAT_UNACCEPTABLEBITNESS,
		COMADMIN_E_CAT_WRONGAPPBITNESS,
		COMADMIN_E_CAT_PAUSE_RESUME_NOT_SUPPORTED,
		COMADMIN_E_CAT_SERVERFAULT,
		COMQC_E_APPLICATION_NOT_QUEUED,
		COMQC_E_NO_QUEUEABLE_INTERFACES,
		COMQC_E_QUEUING_SERVICE_NOT_AVAILABLE,
		COMQC_E_NO_IPERSISTSTREAM,
		COMQC_E_BAD_MESSAGE,
		COMQC_E_UNAUTHENTICATED,
		COMQC_E_UNTRUSTED_ENQUEUER,
		MSDTC_E_DUPLICATE_RESOURCE,
		COMADMIN_E_OBJECT_PARENT_MISSING,
		COMADMIN_E_OBJECT_DOES_NOT_EXIST,
		COMADMIN_E_APP_NOT_RUNNING,
		COMADMIN_E_INVALID_PARTITION,
		COMADMIN_E_SVCAPP_NOT_POOLABLE_OR_RECYCLABLE,
		COMADMIN_E_USER_IN_SET,
		COMADMIN_E_CANTRECYCLELIBRARYAPPS,
		COMADMIN_E_CANTRECYCLESERVICEAPPS,
		COMADMIN_E_PROCESSALREADYRECYCLED,
		COMADMIN_E_PAUSEDPROCESSMAYNOTBERECYCLED,
		COMADMIN_E_CANTMAKEINPROCSERVICE,
		COMADMIN_E_PROGIDINUSEBYCLSID,
		COMADMIN_E_DEFAULT_PARTITION_NOT_IN_SET,
		COMADMIN_E_RECYCLEDPROCESSMAYNOTBEPAUSED,
		COMADMIN_E_PARTITION_ACCESSDENIED,
		COMADMIN_E_PARTITION_MSI_ONLY,
		COMADMIN_E_LEGACYCOMPS_NOT_ALLOWED_IN_1_0_FORMAT,
		COMADMIN_E_LEGACYCOMPS_NOT_ALLOWED_IN_NONBASE_PARTITIONS,
		COMADMIN_E_COMP_MOVE_SOURCE,
		COMADMIN_E_COMP_MOVE_DEST,
		COMADMIN_E_COMP_MOVE_PRIVATE,
		COMADMIN_E_BASEPARTITION_REQUIRED_IN_SET,
		COMADMIN_E_CANNOT_ALIAS_EVENTCLASS,
		COMADMIN_E_PRIVATE_ACCESSDENIED,
		COMADMIN_E_SAFERINVALID,
		COMADMIN_E_REGISTRY_ACCESSDENIED,
		COMADMIN_E_PARTITIONS_DISABLED,
		ERROR_FLT_IO_COMPLETE,
		ERROR_FLT_NO_HANDLER_DEFINED,
		ERROR_FLT_CONTEXT_ALREADY_DEFINED,
		ERROR_FLT_INVALID_ASYNCHRONOUS_REQUEST,
		ERROR_FLT_DISALLOW_FAST_IO,
		ERROR_FLT_INVALID_NAME_REQUEST,
		ERROR_FLT_NOT_SAFE_TO_POST_OPERATION,
		ERROR_FLT_NOT_INITIALIZED,
		ERROR_FLT_FILTER_NOT_READY,
		ERROR_FLT_POST_OPERATION_CLEANUP,
		ERROR_FLT_INTERNAL_ERROR,
		ERROR_FLT_DELETING_OBJECT,
		ERROR_FLT_MUST_BE_NONPAGED_POOL,
		ERROR_FLT_DUPLICATE_ENTRY,
		ERROR_FLT_CBDQ_DISABLED,
		ERROR_FLT_DO_NOT_ATTACH,
		ERROR_FLT_DO_NOT_DETACH,
		ERROR_FLT_INSTANCE_ALTITUDE_COLLISION,
		ERROR_FLT_INSTANCE_NAME_COLLISION,
		ERROR_FLT_FILTER_NOT_FOUND,
		ERROR_FLT_VOLUME_NOT_FOUND,
		ERROR_FLT_INSTANCE_NOT_FOUND,
		ERROR_FLT_CONTEXT_ALLOCATION_NOT_FOUND,
		ERROR_FLT_INVALID_CONTEXT_REGISTRATION,
		ERROR_FLT_NAME_CACHE_MISS,
		ERROR_FLT_NO_DEVICE_OBJECT,
		ERROR_FLT_VOLUME_ALREADY_MOUNTED,
		ERROR_FLT_ALREADY_ENLISTED,
		ERROR_FLT_CONTEXT_ALREADY_LINKED,
		ERROR_FLT_NO_WAITER_FOR_REPLY,
		ERROR_HUNG_DISPLAY_DRIVER_THREAD,
		DWM_E_COMPOSITIONDISABLED,
		DWM_E_REMOTING_NOT_SUPPORTED,
		DWM_E_NO_REDIRECTION_SURFACE_AVAILABLE,
		DWM_E_NOT_QUEUING_PRESENTS,
		DWM_E_ADAPTER_NOT_FOUND,
		DWM_S_GDI_REDIRECTION_SURFACE,
		ERROR_MONITOR_NO_DESCRIPTOR,
		ERROR_MONITOR_UNKNOWN_DESCRIPTOR_FORMAT,
		ERROR_MONITOR_INVALID_DESCRIPTOR_CHECKSUM,
		ERROR_MONITOR_INVALID_STANDARD_TIMING_BLOCK,
		ERROR_MONITOR_WMI_DATABLOCK_REGISTRATION_FAILED,
		ERROR_MONITOR_INVALID_SERIAL_NUMBER_MONDSC_BLOCK,
		ERROR_MONITOR_INVALID_USER_FRIENDLY_MONDSC_BLOCK,
		ERROR_MONITOR_NO_MORE_DESCRIPTOR_DATA,
		ERROR_MONITOR_INVALID_DETAILED_TIMING_BLOCK,
		ERROR_MONITOR_INVALID_MANUFACTURE_DATE,
		ERROR_GRAPHICS_NOT_EXCLUSIVE_MODE_OWNER,
		ERROR_GRAPHICS_INSUFFICIENT_DMA_BUFFER,
		ERROR_GRAPHICS_INVALID_DISPLAY_ADAPTER,
		ERROR_GRAPHICS_ADAPTER_WAS_RESET,
		ERROR_GRAPHICS_INVALID_DRIVER_MODEL,
		ERROR_GRAPHICS_PRESENT_MODE_CHANGED,
		ERROR_GRAPHICS_PRESENT_OCCLUDED,
		ERROR_GRAPHICS_PRESENT_DENIED,
		ERROR_GRAPHICS_CANNOTCOLORCONVERT,
		ERROR_GRAPHICS_DRIVER_MISMATCH,
		ERROR_GRAPHICS_PARTIAL_DATA_POPULATED,
		ERROR_GRAPHICS_PRESENT_REDIRECTION_DISABLED,
		ERROR_GRAPHICS_PRESENT_UNOCCLUDED,
		ERROR_GRAPHICS_NO_VIDEO_MEMORY,
		ERROR_GRAPHICS_CANT_LOCK_MEMORY,
		ERROR_GRAPHICS_ALLOCATION_BUSY,
		ERROR_GRAPHICS_TOO_MANY_REFERENCES,
		ERROR_GRAPHICS_TRY_AGAIN_LATER,
		ERROR_GRAPHICS_TRY_AGAIN_NOW,
		ERROR_GRAPHICS_ALLOCATION_INVALID,
		ERROR_GRAPHICS_UNSWIZZLING_APERTURE_UNAVAILABLE,
		ERROR_GRAPHICS_UNSWIZZLING_APERTURE_UNSUPPORTED,
		ERROR_GRAPHICS_CANT_EVICT_PINNED_ALLOCATION,
		ERROR_GRAPHICS_INVALID_ALLOCATION_USAGE,
		ERROR_GRAPHICS_CANT_RENDER_LOCKED_ALLOCATION,
		ERROR_GRAPHICS_ALLOCATION_CLOSED,
		ERROR_GRAPHICS_INVALID_ALLOCATION_INSTANCE,
		ERROR_GRAPHICS_INVALID_ALLOCATION_HANDLE,
		ERROR_GRAPHICS_WRONG_ALLOCATION_DEVICE,
		ERROR_GRAPHICS_ALLOCATION_CONTENT_LOST,
		ERROR_GRAPHICS_GPU_EXCEPTION_ON_DEVICE,
		ERROR_GRAPHICS_INVALID_VIDPN_TOPOLOGY,
		ERROR_GRAPHICS_VIDPN_TOPOLOGY_NOT_SUPPORTED,
		ERROR_GRAPHICS_VIDPN_TOPOLOGY_CURRENTLY_NOT_SUPPORTED,
		ERROR_GRAPHICS_INVALID_VIDPN,
		ERROR_GRAPHICS_INVALID_VIDEO_PRESENT_SOURCE,
		ERROR_GRAPHICS_INVALID_VIDEO_PRESENT_TARGET,
		ERROR_GRAPHICS_VIDPN_MODALITY_NOT_SUPPORTED,
		ERROR_GRAPHICS_MODE_NOT_PINNED,
		ERROR_GRAPHICS_INVALID_VIDPN_SOURCEMODESET,
		ERROR_GRAPHICS_INVALID_VIDPN_TARGETMODESET,
		ERROR_GRAPHICS_INVALID_FREQUENCY,
		ERROR_GRAPHICS_INVALID_ACTIVE_REGION,
		ERROR_GRAPHICS_INVALID_TOTAL_REGION,
		ERROR_GRAPHICS_INVALID_VIDEO_PRESENT_SOURCE_MODE,
		ERROR_GRAPHICS_INVALID_VIDEO_PRESENT_TARGET_MODE,
		ERROR_GRAPHICS_PINNED_MODE_MUST_REMAIN_IN_SET,
		ERROR_GRAPHICS_PATH_ALREADY_IN_TOPOLOGY,
		ERROR_GRAPHICS_MODE_ALREADY_IN_MODESET,
		ERROR_GRAPHICS_INVALID_VIDEOPRESENTSOURCESET,
		ERROR_GRAPHICS_INVALID_VIDEOPRESENTTARGETSET,
		ERROR_GRAPHICS_SOURCE_ALREADY_IN_SET,
		ERROR_GRAPHICS_TARGET_ALREADY_IN_SET,
		ERROR_GRAPHICS_INVALID_VIDPN_PRESENT_PATH,
		ERROR_GRAPHICS_NO_RECOMMENDED_VIDPN_TOPOLOGY,
		ERROR_GRAPHICS_INVALID_MONITOR_FREQUENCYRANGESET,
		ERROR_GRAPHICS_INVALID_MONITOR_FREQUENCYRANGE,
		ERROR_GRAPHICS_FREQUENCYRANGE_NOT_IN_SET,
		ERROR_GRAPHICS_NO_PREFERRED_MODE,
		ERROR_GRAPHICS_FREQUENCYRANGE_ALREADY_IN_SET,
		ERROR_GRAPHICS_STALE_MODESET,
		ERROR_GRAPHICS_INVALID_MONITOR_SOURCEMODESET,
		ERROR_GRAPHICS_INVALID_MONITOR_SOURCE_MODE,
		ERROR_GRAPHICS_NO_RECOMMENDED_FUNCTIONAL_VIDPN,
		ERROR_GRAPHICS_MODE_ID_MUST_BE_UNIQUE,
		ERROR_GRAPHICS_EMPTY_ADAPTER_MONITOR_MODE_SUPPORT_INTERSECTION,
		ERROR_GRAPHICS_VIDEO_PRESENT_TARGETS_LESS_THAN_SOURCES,
		ERROR_GRAPHICS_PATH_NOT_IN_TOPOLOGY,
		ERROR_GRAPHICS_ADAPTER_MUST_HAVE_AT_LEAST_ONE_SOURCE,
		ERROR_GRAPHICS_ADAPTER_MUST_HAVE_AT_LEAST_ONE_TARGET,
		ERROR_GRAPHICS_INVALID_MONITORDESCRIPTORSET,
		ERROR_GRAPHICS_INVALID_MONITORDESCRIPTOR,
		ERROR_GRAPHICS_MONITORDESCRIPTOR_NOT_IN_SET,
		ERROR_GRAPHICS_MONITORDESCRIPTOR_ALREADY_IN_SET,
		ERROR_GRAPHICS_MONITORDESCRIPTOR_ID_MUST_BE_UNIQUE,
		ERROR_GRAPHICS_INVALID_VIDPN_TARGET_SUBSET_TYPE,
		ERROR_GRAPHICS_RESOURCES_NOT_RELATED,
		ERROR_GRAPHICS_SOURCE_ID_MUST_BE_UNIQUE,
		ERROR_GRAPHICS_TARGET_ID_MUST_BE_UNIQUE,
		ERROR_GRAPHICS_NO_AVAILABLE_VIDPN_TARGET,
		ERROR_GRAPHICS_MONITOR_COULD_NOT_BE_ASSOCIATED_WITH_ADAPTER,
		ERROR_GRAPHICS_NO_VIDPNMGR,
		ERROR_GRAPHICS_NO_ACTIVE_VIDPN,
		ERROR_GRAPHICS_STALE_VIDPN_TOPOLOGY,
		ERROR_GRAPHICS_MONITOR_NOT_CONNECTED,
		ERROR_GRAPHICS_SOURCE_NOT_IN_TOPOLOGY,
		ERROR_GRAPHICS_INVALID_PRIMARYSURFACE_SIZE,
		ERROR_GRAPHICS_INVALID_VISIBLEREGION_SIZE,
		ERROR_GRAPHICS_INVALID_STRIDE,
		ERROR_GRAPHICS_INVALID_PIXELFORMAT,
		ERROR_GRAPHICS_INVALID_COLORBASIS,
		ERROR_GRAPHICS_INVALID_PIXELVALUEACCESSMODE,
		ERROR_GRAPHICS_TARGET_NOT_IN_TOPOLOGY,
		ERROR_GRAPHICS_NO_DISPLAY_MODE_MANAGEMENT_SUPPORT,
		ERROR_GRAPHICS_VIDPN_SOURCE_IN_USE,
		ERROR_GRAPHICS_CANT_ACCESS_ACTIVE_VIDPN,
		ERROR_GRAPHICS_INVALID_PATH_IMPORTANCE_ORDINAL,
		ERROR_GRAPHICS_INVALID_PATH_CONTENT_GEOMETRY_TRANSFORMATION,
		ERROR_GRAPHICS_PATH_CONTENT_GEOMETRY_TRANSFORMATION_NOT_SUPPORTED,
		ERROR_GRAPHICS_INVALID_GAMMA_RAMP,
		ERROR_GRAPHICS_GAMMA_RAMP_NOT_SUPPORTED,
		ERROR_GRAPHICS_MULTISAMPLING_NOT_SUPPORTED,
		ERROR_GRAPHICS_MODE_NOT_IN_MODESET,
		ERROR_GRAPHICS_DATASET_IS_EMPTY,
		ERROR_GRAPHICS_NO_MORE_ELEMENTS_IN_DATASET,
		ERROR_GRAPHICS_INVALID_VIDPN_TOPOLOGY_RECOMMENDATION_REASON,
		ERROR_GRAPHICS_INVALID_PATH_CONTENT_TYPE,
		ERROR_GRAPHICS_INVALID_COPYPROTECTION_TYPE,
		ERROR_GRAPHICS_UNASSIGNED_MODESET_ALREADY_EXISTS,
		ERROR_GRAPHICS_PATH_CONTENT_GEOMETRY_TRANSFORMATION_NOT_PINNED,
		ERROR_GRAPHICS_INVALID_SCANLINE_ORDERING,
		ERROR_GRAPHICS_TOPOLOGY_CHANGES_NOT_ALLOWED,
		ERROR_GRAPHICS_NO_AVAILABLE_IMPORTANCE_ORDINALS,
		ERROR_GRAPHICS_INCOMPATIBLE_PRIVATE_FORMAT,
		ERROR_GRAPHICS_INVALID_MODE_PRUNING_ALGORITHM,
		ERROR_GRAPHICS_INVALID_MONITOR_CAPABILITY_ORIGIN,
		ERROR_GRAPHICS_INVALID_MONITOR_FREQUENCYRANGE_CONSTRAINT,
		ERROR_GRAPHICS_MAX_NUM_PATHS_REACHED,
		ERROR_GRAPHICS_CANCEL_VIDPN_TOPOLOGY_AUGMENTATION,
		ERROR_GRAPHICS_INVALID_CLIENT_TYPE,
		ERROR_GRAPHICS_CLIENTVIDPN_NOT_SET,
		ERROR_GRAPHICS_SPECIFIED_CHILD_ALREADY_CONNECTED,
		ERROR_GRAPHICS_CHILD_DESCRIPTOR_NOT_SUPPORTED,
		ERROR_GRAPHICS_UNKNOWN_CHILD_STATUS,
		ERROR_GRAPHICS_NOT_A_LINKED_ADAPTER,
		ERROR_GRAPHICS_LEADLINK_NOT_ENUMERATED,
		ERROR_GRAPHICS_CHAINLINKS_NOT_ENUMERATED,
		ERROR_GRAPHICS_ADAPTER_CHAIN_NOT_READY,
		ERROR_GRAPHICS_CHAINLINKS_NOT_STARTED,
		ERROR_GRAPHICS_CHAINLINKS_NOT_POWERED_ON,
		ERROR_GRAPHICS_INCONSISTENT_DEVICE_LINK_STATE,
		ERROR_GRAPHICS_LEADLINK_START_DEFERRED,
		ERROR_GRAPHICS_NOT_POST_DEVICE_DRIVER,
		ERROR_GRAPHICS_POLLING_TOO_FREQUENTLY,
		ERROR_GRAPHICS_START_DEFERRED,
		ERROR_GRAPHICS_ADAPTER_ACCESS_NOT_EXCLUDED,
		ERROR_GRAPHICS_OPM_NOT_SUPPORTED,
		ERROR_GRAPHICS_COPP_NOT_SUPPORTED,
		ERROR_GRAPHICS_UAB_NOT_SUPPORTED,
		ERROR_GRAPHICS_OPM_INVALID_ENCRYPTED_PARAMETERS,
		ERROR_GRAPHICS_OPM_NO_VIDEO_OUTPUTS_EXIST,
		ERROR_GRAPHICS_OPM_INTERNAL_ERROR,
		ERROR_GRAPHICS_OPM_INVALID_HANDLE,
		ERROR_GRAPHICS_PVP_INVALID_CERTIFICATE_LENGTH,
		ERROR_GRAPHICS_OPM_SPANNING_MODE_ENABLED,
		ERROR_GRAPHICS_OPM_THEATER_MODE_ENABLED,
		ERROR_GRAPHICS_PVP_HFS_FAILED,
		ERROR_GRAPHICS_OPM_INVALID_SRM,
		ERROR_GRAPHICS_OPM_OUTPUT_DOES_NOT_SUPPORT_HDCP,
		ERROR_GRAPHICS_OPM_OUTPUT_DOES_NOT_SUPPORT_ACP,
		ERROR_GRAPHICS_OPM_OUTPUT_DOES_NOT_SUPPORT_CGMSA,
		ERROR_GRAPHICS_OPM_HDCP_SRM_NEVER_SET,
		ERROR_GRAPHICS_OPM_RESOLUTION_TOO_HIGH,
		ERROR_GRAPHICS_OPM_ALL_HDCP_HARDWARE_ALREADY_IN_USE,
		ERROR_GRAPHICS_OPM_VIDEO_OUTPUT_NO_LONGER_EXISTS,
		ERROR_GRAPHICS_OPM_SESSION_TYPE_CHANGE_IN_PROGRESS,
		ERROR_GRAPHICS_OPM_VIDEO_OUTPUT_DOES_NOT_HAVE_COPP_SEMANTICS,
		ERROR_GRAPHICS_OPM_INVALID_INFORMATION_REQUEST,
		ERROR_GRAPHICS_OPM_DRIVER_INTERNAL_ERROR,
		ERROR_GRAPHICS_OPM_VIDEO_OUTPUT_DOES_NOT_HAVE_OPM_SEMANTICS,
		ERROR_GRAPHICS_OPM_SIGNALING_NOT_SUPPORTED,
		ERROR_GRAPHICS_OPM_INVALID_CONFIGURATION_REQUEST,
		ERROR_GRAPHICS_I2C_NOT_SUPPORTED,
		ERROR_GRAPHICS_I2C_DEVICE_DOES_NOT_EXIST,
		ERROR_GRAPHICS_I2C_ERROR_TRANSMITTING_DATA,
		ERROR_GRAPHICS_I2C_ERROR_RECEIVING_DATA,
		ERROR_GRAPHICS_DDCCI_VCP_NOT_SUPPORTED,
		ERROR_GRAPHICS_DDCCI_INVALID_DATA,
		ERROR_GRAPHICS_DDCCI_MONITOR_RETURNED_INVALID_TIMING_STATUS_BYTE,
		ERROR_GRAPHICS_MCA_INVALID_CAPABILITIES_STRING,
		ERROR_GRAPHICS_MCA_INTERNAL_ERROR,
		ERROR_GRAPHICS_DDCCI_INVALID_MESSAGE_COMMAND,
		ERROR_GRAPHICS_DDCCI_INVALID_MESSAGE_LENGTH,
		ERROR_GRAPHICS_DDCCI_INVALID_MESSAGE_CHECKSUM,
		ERROR_GRAPHICS_INVALID_PHYSICAL_MONITOR_HANDLE,
		ERROR_GRAPHICS_MONITOR_NO_LONGER_EXISTS,
		ERROR_GRAPHICS_DDCCI_CURRENT_CURRENT_VALUE_GREATER_THAN_MAXIMUM_VALUE,
		ERROR_GRAPHICS_MCA_INVALID_VCP_VERSION,
		ERROR_GRAPHICS_MCA_MONITOR_VIOLATES_MCCS_SPECIFICATION,
		ERROR_GRAPHICS_MCA_MCCS_VERSION_MISMATCH,
		ERROR_GRAPHICS_MCA_UNSUPPORTED_MCCS_VERSION,
		ERROR_GRAPHICS_MCA_INVALID_TECHNOLOGY_TYPE_RETURNED,
		ERROR_GRAPHICS_MCA_UNSUPPORTED_COLOR_TEMPERATURE,
		ERROR_GRAPHICS_ONLY_CONSOLE_SESSION_SUPPORTED,
		ERROR_GRAPHICS_NO_DISPLAY_DEVICE_CORRESPONDS_TO_NAME,
		ERROR_GRAPHICS_DISPLAY_DEVICE_NOT_ATTACHED_TO_DESKTOP,
		ERROR_GRAPHICS_MIRRORING_DEVICES_NOT_SUPPORTED,
		ERROR_GRAPHICS_INVALID_POINTER,
		ERROR_GRAPHICS_NO_MONITORS_CORRESPOND_TO_DISPLAY_DEVICE,
		ERROR_GRAPHICS_PARAMETER_ARRAY_TOO_SMALL,
		ERROR_GRAPHICS_INTERNAL_ERROR,
		ERROR_GRAPHICS_SESSION_TYPE_CHANGE_IN_PROGRESS,
		TPM_E_ERROR_MASK,
		TPM_E_AUTHFAIL,
		TPM_E_BADINDEX,
		TPM_E_BAD_PARAMETER,
		TPM_E_AUDITFAILURE,
		TPM_E_CLEAR_DISABLED,
		TPM_E_DEACTIVATED,
		TPM_E_DISABLED,
		TPM_E_DISABLED_CMD,
		TPM_E_FAIL,
		TPM_E_BAD_ORDINAL,
		TPM_E_INSTALL_DISABLED,
		TPM_E_INVALID_KEYHANDLE,
		TPM_E_KEYNOTFOUND,
		TPM_E_INAPPROPRIATE_ENC,
		TPM_E_MIGRATEFAIL,
		TPM_E_INVALID_PCR_INFO,
		TPM_E_NOSPACE,
		TPM_E_NOSRK,
		TPM_E_NOTSEALED_BLOB,
		TPM_E_OWNER_SET,
		TPM_E_RESOURCES,
		TPM_E_SHORTRANDOM,
		TPM_E_SIZE,
		TPM_E_WRONGPCRVAL,
		TPM_E_BAD_PARAM_SIZE,
		TPM_E_SHA_THREAD,
		TPM_E_SHA_ERROR,
		TPM_E_FAILEDSELFTEST,
		TPM_E_AUTH2FAIL,
		TPM_E_BADTAG,
		TPM_E_IOERROR,
		TPM_E_ENCRYPT_ERROR,
		TPM_E_DECRYPT_ERROR,
		TPM_E_INVALID_AUTHHANDLE,
		TPM_E_NO_ENDORSEMENT,
		TPM_E_INVALID_KEYUSAGE,
		TPM_E_WRONG_ENTITYTYPE,
		TPM_E_INVALID_POSTINIT,
		TPM_E_INAPPROPRIATE_SIG,
		TPM_E_BAD_KEY_PROPERTY,
		TPM_E_BAD_MIGRATION,
		TPM_E_BAD_SCHEME,
		TPM_E_BAD_DATASIZE,
		TPM_E_BAD_MODE,
		TPM_E_BAD_PRESENCE,
		TPM_E_BAD_VERSION,
		TPM_E_NO_WRAP_TRANSPORT,
		TPM_E_AUDITFAIL_UNSUCCESSFUL,
		TPM_E_AUDITFAIL_SUCCESSFUL,
		TPM_E_NOTRESETABLE,
		TPM_E_NOTLOCAL,
		TPM_E_BAD_TYPE,
		TPM_E_INVALID_RESOURCE,
		TPM_E_NOTFIPS,
		TPM_E_INVALID_FAMILY,
		TPM_E_NO_NV_PERMISSION,
		TPM_E_REQUIRES_SIGN,
		TPM_E_KEY_NOTSUPPORTED,
		TPM_E_AUTH_CONFLICT,
		TPM_E_AREA_LOCKED,
		TPM_E_BAD_LOCALITY,
		TPM_E_READ_ONLY,
		TPM_E_PER_NOWRITE,
		TPM_E_FAMILYCOUNT,
		TPM_E_WRITE_LOCKED,
		TPM_E_BAD_ATTRIBUTES,
		TPM_E_INVALID_STRUCTURE,
		TPM_E_KEY_OWNER_CONTROL,
		TPM_E_BAD_COUNTER,
		TPM_E_NOT_FULLWRITE,
		TPM_E_CONTEXT_GAP,
		TPM_E_MAXNVWRITES,
		TPM_E_NOOPERATOR,
		TPM_E_RESOURCEMISSING,
		TPM_E_DELEGATE_LOCK,
		TPM_E_DELEGATE_FAMILY,
		TPM_E_DELEGATE_ADMIN,
		TPM_E_TRANSPORT_NOTEXCLUSIVE,
		TPM_E_OWNER_CONTROL,
		TPM_E_DAA_RESOURCES,
		TPM_E_DAA_INPUT_DATA0,
		TPM_E_DAA_INPUT_DATA1,
		TPM_E_DAA_ISSUER_SETTINGS,
		TPM_E_DAA_TPM_SETTINGS,
		TPM_E_DAA_STAGE,
		TPM_E_DAA_ISSUER_VALIDITY,
		TPM_E_DAA_WRONG_W,
		TPM_E_BAD_HANDLE,
		TPM_E_BAD_DELEGATE,
		TPM_E_BADCONTEXT,
		TPM_E_TOOMANYCONTEXTS,
		TPM_E_MA_TICKET_SIGNATURE,
		TPM_E_MA_DESTINATION,
		TPM_E_MA_SOURCE,
		TPM_E_MA_AUTHORITY,
		TPM_E_PERMANENTEK,
		TPM_E_BAD_SIGNATURE,
		TPM_E_NOCONTEXTSPACE,
		TPM_E_COMMAND_BLOCKED,
		TPM_E_INVALID_HANDLE,
		TPM_E_DUPLICATE_VHANDLE,
		TPM_E_EMBEDDED_COMMAND_BLOCKED,
		TPM_E_EMBEDDED_COMMAND_UNSUPPORTED,
		TPM_E_RETRY,
		TPM_E_NEEDS_SELFTEST,
		TPM_E_DOING_SELFTEST,
		TPM_E_DEFEND_LOCK_RUNNING,
		TBS_E_INTERNAL_ERROR,
		TBS_E_BAD_PARAMETER,
		TBS_E_INVALID_OUTPUT_POINTER,
		TBS_E_INVALID_CONTEXT,
		TBS_E_INSUFFICIENT_BUFFER,
		TBS_E_IOERROR,
		TBS_E_INVALID_CONTEXT_PARAM,
		TBS_E_SERVICE_NOT_RUNNING,
		TBS_E_TOO_MANY_TBS_CONTEXTS,
		TBS_E_TOO_MANY_RESOURCES,
		TBS_E_SERVICE_START_PENDING,
		TBS_E_PPI_NOT_SUPPORTED,
		TBS_E_COMMAND_CANCELED,
		TBS_E_BUFFER_TOO_LARGE,
		TBS_E_TPM_NOT_FOUND,
		TBS_E_SERVICE_DISABLED,
		TBS_E_NO_EVENT_LOG,
		TPMAPI_E_INVALID_STATE,
		TPMAPI_E_NOT_ENOUGH_DATA,
		TPMAPI_E_TOO_MUCH_DATA,
		TPMAPI_E_INVALID_OUTPUT_POINTER,
		TPMAPI_E_INVALID_PARAMETER,
		TPMAPI_E_OUT_OF_MEMORY,
		TPMAPI_E_BUFFER_TOO_SMALL,
		TPMAPI_E_INTERNAL_ERROR,
		TPMAPI_E_ACCESS_DENIED,
		TPMAPI_E_AUTHORIZATION_FAILED,
		TPMAPI_E_INVALID_CONTEXT_HANDLE,
		TPMAPI_E_TBS_COMMUNICATION_ERROR,
		TPMAPI_E_TPM_COMMAND_ERROR,
		TPMAPI_E_MESSAGE_TOO_LARGE,
		TPMAPI_E_INVALID_ENCODING,
		TPMAPI_E_INVALID_KEY_SIZE,
		TPMAPI_E_ENCRYPTION_FAILED,
		TPMAPI_E_INVALID_KEY_PARAMS,
		TPMAPI_E_INVALID_MIGRATION_AUTHORIZATION_BLOB,
		TPMAPI_E_INVALID_PCR_INDEX,
		TPMAPI_E_INVALID_DELEGATE_BLOB,
		TPMAPI_E_INVALID_CONTEXT_PARAMS,
		TPMAPI_E_INVALID_KEY_BLOB,
		TPMAPI_E_INVALID_PCR_DATA,
		TPMAPI_E_INVALID_OWNER_AUTH,
		TPMAPI_E_FIPS_RNG_CHECK_FAILED,
		TPMAPI_E_EMPTY_TCG_LOG,
		TPMAPI_E_INVALID_TCG_LOG_ENTRY,
		TPMAPI_E_TCG_SEPARATOR_ABSENT,
		TPMAPI_E_TCG_INVALID_DIGEST_ENTRY,
		TBSIMP_E_BUFFER_TOO_SMALL,
		TBSIMP_E_CLEANUP_FAILED,
		TBSIMP_E_INVALID_CONTEXT_HANDLE,
		TBSIMP_E_INVALID_CONTEXT_PARAM,
		TBSIMP_E_TPM_ERROR,
		TBSIMP_E_HASH_BAD_KEY,
		TBSIMP_E_DUPLICATE_VHANDLE,
		TBSIMP_E_INVALID_OUTPUT_POINTER,
		TBSIMP_E_INVALID_PARAMETER,
		TBSIMP_E_RPC_INIT_FAILED,
		TBSIMP_E_SCHEDULER_NOT_RUNNING,
		TBSIMP_E_COMMAND_CANCELED,
		TBSIMP_E_OUT_OF_MEMORY,
		TBSIMP_E_LIST_NO_MORE_ITEMS,
		TBSIMP_E_LIST_NOT_FOUND,
		TBSIMP_E_NOT_ENOUGH_SPACE,
		TBSIMP_E_NOT_ENOUGH_TPM_CONTEXTS,
		TBSIMP_E_COMMAND_FAILED,
		TBSIMP_E_UNKNOWN_ORDINAL,
		TBSIMP_E_RESOURCE_EXPIRED,
		TBSIMP_E_INVALID_RESOURCE,
		TBSIMP_E_NOTHING_TO_UNLOAD,
		TBSIMP_E_HASH_TABLE_FULL,
		TBSIMP_E_TOO_MANY_TBS_CONTEXTS,
		TBSIMP_E_TOO_MANY_RESOURCES,
		TBSIMP_E_PPI_NOT_SUPPORTED,
		TBSIMP_E_TPM_INCOMPATIBLE,
		TBSIMP_E_NO_EVENT_LOG,
		TPM_E_PPI_ACPI_FAILURE,
		TPM_E_PPI_USER_ABORT,
		TPM_E_PPI_BIOS_FAILURE,
		TPM_E_PPI_NOT_SUPPORTED,
		PLA_E_DCS_NOT_FOUND,
		PLA_E_DCS_IN_USE,
		PLA_E_TOO_MANY_FOLDERS,
		PLA_E_NO_MIN_DISK,
		PLA_E_DCS_ALREADY_EXISTS,
		PLA_S_PROPERTY_IGNORED,
		PLA_E_PROPERTY_CONFLICT,
		PLA_E_DCS_SINGLETON_REQUIRED,
		PLA_E_CREDENTIALS_REQUIRED,
		PLA_E_DCS_NOT_RUNNING,
		PLA_E_CONFLICT_INCL_EXCL_API,
		PLA_E_NETWORK_EXE_NOT_VALID,
		PLA_E_EXE_ALREADY_CONFIGURED,
		PLA_E_EXE_PATH_NOT_VALID,
		PLA_E_DC_ALREADY_EXISTS,
		PLA_E_DCS_START_WAIT_TIMEOUT,
		PLA_E_DC_START_WAIT_TIMEOUT,
		PLA_E_REPORT_WAIT_TIMEOUT,
		PLA_E_NO_DUPLICATES,
		PLA_E_EXE_FULL_PATH_REQUIRED,
		PLA_E_INVALID_SESSION_NAME,
		PLA_E_PLA_CHANNEL_NOT_ENABLED,
		PLA_E_TASKSCHED_CHANNEL_NOT_ENABLED,
		PLA_E_RULES_MANAGER_FAILED,
		PLA_E_CABAPI_FAILURE,
		FVE_E_LOCKED_VOLUME,
		FVE_E_NOT_ENCRYPTED,
		FVE_E_NO_TPM_BIOS,
		FVE_E_NO_MBR_METRIC,
		FVE_E_NO_BOOTSECTOR_METRIC,
		FVE_E_NO_BOOTMGR_METRIC,
		FVE_E_WRONG_BOOTMGR,
		FVE_E_SECURE_KEY_REQUIRED,
		FVE_E_NOT_ACTIVATED,
		FVE_E_ACTION_NOT_ALLOWED,
		FVE_E_AD_SCHEMA_NOT_INSTALLED,
		FVE_E_AD_INVALID_DATATYPE,
		FVE_E_AD_INVALID_DATASIZE,
		FVE_E_AD_NO_VALUES,
		FVE_E_AD_ATTR_NOT_SET,
		FVE_E_AD_GUID_NOT_FOUND,
		FVE_E_BAD_INFORMATION,
		FVE_E_TOO_SMALL,
		FVE_E_SYSTEM_VOLUME,
		FVE_E_FAILED_WRONG_FS,
		FVE_E_BAD_PARTITION_SIZE,
		FVE_E_NOT_SUPPORTED,
		FVE_E_BAD_DATA,
		FVE_E_VOLUME_NOT_BOUND,
		FVE_E_TPM_NOT_OWNED,
		FVE_E_NOT_DATA_VOLUME,
		FVE_E_AD_INSUFFICIENT_BUFFER,
		FVE_E_CONV_READ,
		FVE_E_CONV_WRITE,
		FVE_E_KEY_REQUIRED,
		FVE_E_CLUSTERING_NOT_SUPPORTED,
		FVE_E_VOLUME_BOUND_ALREADY,
		FVE_E_OS_NOT_PROTECTED,
		FVE_E_PROTECTION_DISABLED,
		FVE_E_RECOVERY_KEY_REQUIRED,
		FVE_E_FOREIGN_VOLUME,
		FVE_E_OVERLAPPED_UPDATE,
		FVE_E_TPM_SRK_AUTH_NOT_ZERO,
		FVE_E_FAILED_SECTOR_SIZE,
		FVE_E_FAILED_AUTHENTICATION,
		FVE_E_NOT_OS_VOLUME,
		FVE_E_AUTOUNLOCK_ENABLED,
		FVE_E_WRONG_BOOTSECTOR,
		FVE_E_WRONG_SYSTEM_FS,
		FVE_E_POLICY_PASSWORD_REQUIRED,
		FVE_E_CANNOT_SET_FVEK_ENCRYPTED,
		FVE_E_CANNOT_ENCRYPT_NO_KEY,
		FVE_E_BOOTABLE_CDDVD,
		FVE_E_PROTECTOR_EXISTS,
		FVE_E_RELATIVE_PATH,
		FVE_E_PROTECTOR_NOT_FOUND,
		FVE_E_INVALID_KEY_FORMAT,
		FVE_E_INVALID_PASSWORD_FORMAT,
		FVE_E_FIPS_RNG_CHECK_FAILED,
		FVE_E_FIPS_PREVENTS_RECOVERY_PASSWORD,
		FVE_E_FIPS_PREVENTS_EXTERNAL_KEY_EXPORT,
		FVE_E_NOT_DECRYPTED,
		FVE_E_INVALID_PROTECTOR_TYPE,
		FVE_E_NO_PROTECTORS_TO_TEST,
		FVE_E_KEYFILE_NOT_FOUND,
		FVE_E_KEYFILE_INVALID,
		FVE_E_KEYFILE_NO_VMK,
		FVE_E_TPM_DISABLED,
		FVE_E_NOT_ALLOWED_IN_SAFE_MODE,
		FVE_E_TPM_INVALID_PCR,
		FVE_E_TPM_NO_VMK,
		FVE_E_PIN_INVALID,
		FVE_E_AUTH_INVALID_APPLICATION,
		FVE_E_AUTH_INVALID_CONFIG,
		FVE_E_FIPS_DISABLE_PROTECTION_NOT_ALLOWED,
		FVE_E_FS_NOT_EXTENDED,
		FVE_E_FIRMWARE_TYPE_NOT_SUPPORTED,
		FVE_E_NO_LICENSE,
		FVE_E_NOT_ON_STACK,
		FVE_E_FS_MOUNTED,
		FVE_E_TOKEN_NOT_IMPERSONATED,
		FVE_E_DRY_RUN_FAILED,
		FVE_E_REBOOT_REQUIRED,
		FVE_E_DEBUGGER_ENABLED,
		FVE_E_RAW_ACCESS,
		FVE_E_RAW_BLOCKED,
		FVE_E_BCD_APPLICATIONS_PATH_INCORRECT,
		FVE_E_NOT_ALLOWED_IN_VERSION,
		FVE_E_NO_AUTOUNLOCK_MASTER_KEY,
		FVE_E_MOR_FAILED,
		FVE_E_HIDDEN_VOLUME,
		FVE_E_TRANSIENT_STATE,
		FVE_E_PUBKEY_NOT_ALLOWED,
		FVE_E_VOLUME_HANDLE_OPEN,
		FVE_E_NO_FEATURE_LICENSE,
		FVE_E_INVALID_STARTUP_OPTIONS,
		FVE_E_POLICY_RECOVERY_PASSWORD_NOT_ALLOWED,
		FVE_E_POLICY_RECOVERY_PASSWORD_REQUIRED,
		FVE_E_POLICY_RECOVERY_KEY_NOT_ALLOWED,
		FVE_E_POLICY_RECOVERY_KEY_REQUIRED,
		FVE_E_POLICY_STARTUP_PIN_NOT_ALLOWED,
		FVE_E_POLICY_STARTUP_PIN_REQUIRED,
		FVE_E_POLICY_STARTUP_KEY_NOT_ALLOWED,
		FVE_E_POLICY_STARTUP_KEY_REQUIRED,
		FVE_E_POLICY_STARTUP_PIN_KEY_NOT_ALLOWED,
		FVE_E_POLICY_STARTUP_PIN_KEY_REQUIRED,
		FVE_E_POLICY_STARTUP_TPM_NOT_ALLOWED,
		FVE_E_POLICY_STARTUP_TPM_REQUIRED,
		FVE_E_POLICY_INVALID_PIN_LENGTH,
		FVE_E_KEY_PROTECTOR_NOT_SUPPORTED,
		FVE_E_POLICY_PASSPHRASE_NOT_ALLOWED,
		FVE_E_POLICY_PASSPHRASE_REQUIRED,
		FVE_E_FIPS_PREVENTS_PASSPHRASE,
		FVE_E_OS_VOLUME_PASSPHRASE_NOT_ALLOWED,
		FVE_E_INVALID_BITLOCKER_OID,
		FVE_E_VOLUME_TOO_SMALL,
		FVE_E_DV_NOT_SUPPORTED_ON_FS,
		FVE_E_DV_NOT_ALLOWED_BY_GP,
		FVE_E_POLICY_USER_CERTIFICATE_NOT_ALLOWED,
		FVE_E_POLICY_USER_CERTIFICATE_REQUIRED,
		FVE_E_POLICY_USER_CERT_MUST_BE_HW,
		FVE_E_POLICY_USER_CONFIGURE_FDV_AUTOUNLOCK_NOT_ALLOWED,
		FVE_E_POLICY_USER_CONFIGURE_RDV_AUTOUNLOCK_NOT_ALLOWED,
		FVE_E_POLICY_USER_CONFIGURE_RDV_NOT_ALLOWED,
		FVE_E_POLICY_USER_ENABLE_RDV_NOT_ALLOWED,
		FVE_E_POLICY_USER_DISABLE_RDV_NOT_ALLOWED,
		FVE_E_POLICY_INVALID_PASSPHRASE_LENGTH,
		FVE_E_POLICY_PASSPHRASE_TOO_SIMPLE,
		FVE_E_RECOVERY_PARTITION,
		FVE_E_POLICY_CONFLICT_FDV_RK_OFF_AUK_ON,
		FVE_E_POLICY_CONFLICT_RDV_RK_OFF_AUK_ON,
		FVE_E_NON_BITLOCKER_OID,
		FVE_E_POLICY_PROHIBITS_SELFSIGNED,
		FVE_E_POLICY_CONFLICT_RO_AND_STARTUP_KEY_REQUIRED,
		FVE_E_CONV_RECOVERY_FAILED,
		FVE_E_VIRTUALIZED_SPACE_TOO_BIG,
		FVE_E_POLICY_CONFLICT_OSV_RP_OFF_ADB_ON,
		FVE_E_POLICY_CONFLICT_FDV_RP_OFF_ADB_ON,
		FVE_E_POLICY_CONFLICT_RDV_RP_OFF_ADB_ON,
		FVE_E_NON_BITLOCKER_KU,
		FVE_E_PRIVATEKEY_AUTH_FAILED,
		FVE_E_REMOVAL_OF_DRA_FAILED,
		FVE_E_OPERATION_NOT_SUPPORTED_ON_VISTA_VOLUME,
		FVE_E_CANT_LOCK_AUTOUNLOCK_ENABLED_VOLUME,
		FVE_E_FIPS_HASH_KDF_NOT_ALLOWED,
		FVE_E_ENH_PIN_INVALID,
		FVE_E_INVALID_PIN_CHARS,
		FVE_E_INVALID_DATUM_TYPE,
		FWP_E_CALLOUT_NOT_FOUND,
		FWP_E_CONDITION_NOT_FOUND,
		FWP_E_FILTER_NOT_FOUND,
		FWP_E_LAYER_NOT_FOUND,
		FWP_E_PROVIDER_NOT_FOUND,
		FWP_E_PROVIDER_CONTEXT_NOT_FOUND,
		FWP_E_SUBLAYER_NOT_FOUND,
		FWP_E_NOT_FOUND,
		FWP_E_ALREADY_EXISTS,
		FWP_E_IN_USE,
		FWP_E_DYNAMIC_SESSION_IN_PROGRESS,
		FWP_E_WRONG_SESSION,
		FWP_E_NO_TXN_IN_PROGRESS,
		FWP_E_TXN_IN_PROGRESS,
		FWP_E_TXN_ABORTED,
		FWP_E_SESSION_ABORTED,
		FWP_E_INCOMPATIBLE_TXN,
		FWP_E_TIMEOUT,
		FWP_E_NET_EVENTS_DISABLED,
		FWP_E_INCOMPATIBLE_LAYER,
		FWP_E_KM_CLIENTS_ONLY,
		FWP_E_LIFETIME_MISMATCH,
		FWP_E_BUILTIN_OBJECT,
		FWP_E_TOO_MANY_CALLOUTS,
		FWP_E_NOTIFICATION_DROPPED,
		FWP_E_TRAFFIC_MISMATCH,
		FWP_E_INCOMPATIBLE_SA_STATE,
		FWP_E_NULL_POINTER,
		FWP_E_INVALID_ENUMERATOR,
		FWP_E_INVALID_FLAGS,
		FWP_E_INVALID_NET_MASK,
		FWP_E_INVALID_RANGE,
		FWP_E_INVALID_INTERVAL,
		FWP_E_ZERO_LENGTH_ARRAY,
		FWP_E_NULL_DISPLAY_NAME,
		FWP_E_INVALID_ACTION_TYPE,
		FWP_E_INVALID_WEIGHT,
		FWP_E_MATCH_TYPE_MISMATCH,
		FWP_E_TYPE_MISMATCH,
		FWP_E_OUT_OF_BOUNDS,
		FWP_E_RESERVED,
		FWP_E_DUPLICATE_CONDITION,
		FWP_E_DUPLICATE_KEYMOD,
		FWP_E_ACTION_INCOMPATIBLE_WITH_LAYER,
		FWP_E_ACTION_INCOMPATIBLE_WITH_SUBLAYER,
		FWP_E_CONTEXT_INCOMPATIBLE_WITH_LAYER,
		FWP_E_CONTEXT_INCOMPATIBLE_WITH_CALLOUT,
		FWP_E_INCOMPATIBLE_AUTH_METHOD,
		FWP_E_INCOMPATIBLE_DH_GROUP,
		FWP_E_EM_NOT_SUPPORTED,
		FWP_E_NEVER_MATCH,
		FWP_E_PROVIDER_CONTEXT_MISMATCH,
		FWP_E_INVALID_PARAMETER,
		FWP_E_TOO_MANY_SUBLAYERS,
		FWP_E_CALLOUT_NOTIFICATION_FAILED,
		FWP_E_INVALID_AUTH_TRANSFORM,
		FWP_E_INVALID_CIPHER_TRANSFORM,
		FWP_E_DROP_NOICMP,
		FWP_E_INCOMPATIBLE_CIPHER_TRANSFORM,
		FWP_E_INVALID_TRANSFORM_COMBINATION,
		FWP_E_DUPLICATE_AUTH_METHOD,
		WS_S_ASYNC,
		WS_S_END,
		WS_E_INVALID_FORMAT,
		WS_E_OBJECT_FAULTED,
		WS_E_NUMERIC_OVERFLOW,
		WS_E_INVALID_OPERATION,
		WS_E_OPERATION_ABORTED,
		WS_E_ENDPOINT_ACCESS_DENIED,
		WS_E_OPERATION_TIMED_OUT,
		WS_E_OPERATION_ABANDONED,
		WS_E_QUOTA_EXCEEDED,
		WS_E_NO_TRANSLATION_AVAILABLE,
		WS_E_SECURITY_VERIFICATION_FAILURE,
		WS_E_ADDRESS_IN_USE,
		WS_E_ADDRESS_NOT_AVAILABLE,
		WS_E_ENDPOINT_NOT_FOUND,
		WS_E_ENDPOINT_NOT_AVAILABLE,
		WS_E_ENDPOINT_FAILURE,
		WS_E_ENDPOINT_UNREACHABLE,
		WS_E_ENDPOINT_ACTION_NOT_SUPPORTED,
		WS_E_ENDPOINT_TOO_BUSY,
		WS_E_ENDPOINT_FAULT_RECEIVED,
		WS_E_ENDPOINT_DISCONNECTED,
		WS_E_PROXY_FAILURE,
		WS_E_PROXY_ACCESS_DENIED,
		WS_E_NOT_SUPPORTED,
		WS_E_PROXY_REQUIRES_BASIC_AUTH,
		WS_E_PROXY_REQUIRES_DIGEST_AUTH,
		WS_E_PROXY_REQUIRES_NTLM_AUTH,
		WS_E_PROXY_REQUIRES_NEGOTIATE_AUTH,
		WS_E_SERVER_REQUIRES_BASIC_AUTH,
		WS_E_SERVER_REQUIRES_DIGEST_AUTH,
		WS_E_SERVER_REQUIRES_NTLM_AUTH,
		WS_E_SERVER_REQUIRES_NEGOTIATE_AUTH,
		WS_E_INVALID_ENDPOINT_URL,
		WS_E_OTHER,
		WS_E_SECURITY_TOKEN_EXPIRED,
		WS_E_SECURITY_SYSTEM_FAILURE,
		ERROR_NDIS_INTERFACE_CLOSING,
		ERROR_NDIS_BAD_VERSION,
		ERROR_NDIS_BAD_CHARACTERISTICS,
		ERROR_NDIS_ADAPTER_NOT_FOUND,
		ERROR_NDIS_OPEN_FAILED,
		ERROR_NDIS_DEVICE_FAILED,
		ERROR_NDIS_MULTICAST_FULL,
		ERROR_NDIS_MULTICAST_EXISTS,
		ERROR_NDIS_MULTICAST_NOT_FOUND,
		ERROR_NDIS_REQUEST_ABORTED,
		ERROR_NDIS_RESET_IN_PROGRESS,
		ERROR_NDIS_NOT_SUPPORTED,
		ERROR_NDIS_INVALID_PACKET,
		ERROR_NDIS_ADAPTER_NOT_READY,
		ERROR_NDIS_INVALID_LENGTH,
		ERROR_NDIS_INVALID_DATA,
		ERROR_NDIS_BUFFER_TOO_SHORT,
		ERROR_NDIS_INVALID_OID,
		ERROR_NDIS_ADAPTER_REMOVED,
		ERROR_NDIS_UNSUPPORTED_MEDIA,
		ERROR_NDIS_GROUP_ADDRESS_IN_USE,
		ERROR_NDIS_FILE_NOT_FOUND,
		ERROR_NDIS_ERROR_READING_FILE,
		ERROR_NDIS_ALREADY_MAPPED,
		ERROR_NDIS_RESOURCE_CONFLICT,
		ERROR_NDIS_MEDIA_DISCONNECTED,
		ERROR_NDIS_INVALID_ADDRESS,
		ERROR_NDIS_INVALID_DEVICE_REQUEST,
		ERROR_NDIS_PAUSED,
		ERROR_NDIS_INTERFACE_NOT_FOUND,
		ERROR_NDIS_UNSUPPORTED_REVISION,
		ERROR_NDIS_INVALID_PORT,
		ERROR_NDIS_INVALID_PORT_STATE,
		ERROR_NDIS_LOW_POWER_STATE,
		ERROR_NDIS_DOT11_AUTO_CONFIG_ENABLED,
		ERROR_NDIS_DOT11_MEDIA_IN_USE,
		ERROR_NDIS_DOT11_POWER_STATE_INVALID,
		ERROR_NDIS_PM_WOL_PATTERN_LIST_FULL,
		ERROR_NDIS_PM_PROTOCOL_OFFLOAD_LIST_FULL,
		ERROR_NDIS_INDICATION_REQUIRED,
		ERROR_NDIS_OFFLOAD_POLICY,
		ERROR_NDIS_OFFLOAD_CONNECTION_REJECTED,
		ERROR_NDIS_OFFLOAD_PATH_REJECTED,
		ERROR_HV_INVALID_HYPERCALL_CODE,
		ERROR_HV_INVALID_HYPERCALL_INPUT,
		ERROR_HV_INVALID_ALIGNMENT,
		ERROR_HV_INVALID_PARAMETER,
		ERROR_HV_ACCESS_DENIED,
		ERROR_HV_INVALID_PARTITION_STATE,
		ERROR_HV_OPERATION_DENIED,
		ERROR_HV_UNKNOWN_PROPERTY,
		ERROR_HV_PROPERTY_VALUE_OUT_OF_RANGE,
		ERROR_HV_INSUFFICIENT_MEMORY,
		ERROR_HV_PARTITION_TOO_DEEP,
		ERROR_HV_INVALID_PARTITION_ID,
		ERROR_HV_INVALID_VP_INDEX,
		ERROR_HV_INVALID_PORT_ID,
		ERROR_HV_INVALID_CONNECTION_ID,
		ERROR_HV_INSUFFICIENT_BUFFERS,
		ERROR_HV_NOT_ACKNOWLEDGED,
		ERROR_HV_ACKNOWLEDGED,
		ERROR_HV_INVALID_SAVE_RESTORE_STATE,
		ERROR_HV_INVALID_SYNIC_STATE,
		ERROR_HV_OBJECT_IN_USE,
		ERROR_HV_INVALID_PROXIMITY_DOMAIN_INFO,
		ERROR_HV_NO_DATA,
		ERROR_HV_INACTIVE,
		ERROR_HV_NO_RESOURCES,
		ERROR_HV_FEATURE_UNAVAILABLE,
		ERROR_HV_NOT_PRESENT,
		ERROR_VID_DUPLICATE_HANDLER,
		ERROR_VID_TOO_MANY_HANDLERS,
		ERROR_VID_QUEUE_FULL,
		ERROR_VID_HANDLER_NOT_PRESENT,
		ERROR_VID_INVALID_OBJECT_NAME,
		ERROR_VID_PARTITION_NAME_TOO_LONG,
		ERROR_VID_MESSAGE_QUEUE_NAME_TOO_LONG,
		ERROR_VID_PARTITION_ALREADY_EXISTS,
		ERROR_VID_PARTITION_DOES_NOT_EXIST,
		ERROR_VID_PARTITION_NAME_NOT_FOUND,
		ERROR_VID_MESSAGE_QUEUE_ALREADY_EXISTS,
		ERROR_VID_EXCEEDED_MBP_ENTRY_MAP_LIMIT,
		ERROR_VID_MB_STILL_REFERENCED,
		ERROR_VID_CHILD_GPA_PAGE_SET_CORRUPTED,
		ERROR_VID_INVALID_NUMA_SETTINGS,
		ERROR_VID_INVALID_NUMA_NODE_INDEX,
		ERROR_VID_NOTIFICATION_QUEUE_ALREADY_ASSOCIATED,
		ERROR_VID_INVALID_MEMORY_BLOCK_HANDLE,
		ERROR_VID_PAGE_RANGE_OVERFLOW,
		ERROR_VID_INVALID_MESSAGE_QUEUE_HANDLE,
		ERROR_VID_INVALID_GPA_RANGE_HANDLE,
		ERROR_VID_NO_MEMORY_BLOCK_NOTIFICATION_QUEUE,
		ERROR_VID_MEMORY_BLOCK_LOCK_COUNT_EXCEEDED,
		ERROR_VID_INVALID_PPM_HANDLE,
		ERROR_VID_MBPS_ARE_LOCKED,
		ERROR_VID_MESSAGE_QUEUE_CLOSED,
		ERROR_VID_VIRTUAL_PROCESSOR_LIMIT_EXCEEDED,
		ERROR_VID_STOP_PENDING,
		ERROR_VID_INVALID_PROCESSOR_STATE,
		ERROR_VID_EXCEEDED_KM_CONTEXT_COUNT_LIMIT,
		ERROR_VID_KM_INTERFACE_ALREADY_INITIALIZED,
		ERROR_VID_MB_PROPERTY_ALREADY_SET_RESET,
		ERROR_VID_MMIO_RANGE_DESTROYED,
		ERROR_VID_INVALID_CHILD_GPA_PAGE_SET,
		ERROR_VID_RESERVE_PAGE_SET_IS_BEING_USED,
		ERROR_VID_RESERVE_PAGE_SET_TOO_SMALL,
		ERROR_VID_MBP_ALREADY_LOCKED_USING_RESERVED_PAGE,
		ERROR_VID_MBP_COUNT_EXCEEDED_LIMIT,
		ERROR_VID_SAVED_STATE_CORRUPT,
		ERROR_VID_SAVED_STATE_UNRECOGNIZED_ITEM,
		ERROR_VID_SAVED_STATE_INCOMPATIBLE,
		ERROR_VID_REMOTE_NODE_PARENT_GPA_PAGES_USED,
		ERROR_VOLMGR_INCOMPLETE_REGENERATION,
		ERROR_VOLMGR_INCOMPLETE_DISK_MIGRATION,
		ERROR_VOLMGR_DATABASE_FULL,
		ERROR_VOLMGR_DISK_CONFIGURATION_CORRUPTED,
		ERROR_VOLMGR_DISK_CONFIGURATION_NOT_IN_SYNC,
		ERROR_VOLMGR_PACK_CONFIG_UPDATE_FAILED,
		ERROR_VOLMGR_DISK_CONTAINS_NON_SIMPLE_VOLUME,
		ERROR_VOLMGR_DISK_DUPLICATE,
		ERROR_VOLMGR_DISK_DYNAMIC,
		ERROR_VOLMGR_DISK_ID_INVALID,
		ERROR_VOLMGR_DISK_INVALID,
		ERROR_VOLMGR_DISK_LAST_VOTER,
		ERROR_VOLMGR_DISK_LAYOUT_INVALID,
		ERROR_VOLMGR_DISK_LAYOUT_NON_BASIC_BETWEEN_BASIC_PARTITIONS,
		ERROR_VOLMGR_DISK_LAYOUT_NOT_CYLINDER_ALIGNED,
		ERROR_VOLMGR_DISK_LAYOUT_PARTITIONS_TOO_SMALL,
		ERROR_VOLMGR_DISK_LAYOUT_PRIMARY_BETWEEN_LOGICAL_PARTITIONS,
		ERROR_VOLMGR_DISK_LAYOUT_TOO_MANY_PARTITIONS,
		ERROR_VOLMGR_DISK_MISSING,
		ERROR_VOLMGR_DISK_NOT_EMPTY,
		ERROR_VOLMGR_DISK_NOT_ENOUGH_SPACE,
		ERROR_VOLMGR_DISK_REVECTORING_FAILED,
		ERROR_VOLMGR_DISK_SECTOR_SIZE_INVALID,
		ERROR_VOLMGR_DISK_SET_NOT_CONTAINED,
		ERROR_VOLMGR_DISK_USED_BY_MULTIPLE_MEMBERS,
		ERROR_VOLMGR_DISK_USED_BY_MULTIPLE_PLEXES,
		ERROR_VOLMGR_DYNAMIC_DISK_NOT_SUPPORTED,
		ERROR_VOLMGR_EXTENT_ALREADY_USED,
		ERROR_VOLMGR_EXTENT_NOT_CONTIGUOUS,
		ERROR_VOLMGR_EXTENT_NOT_IN_PUBLIC_REGION,
		ERROR_VOLMGR_EXTENT_NOT_SECTOR_ALIGNED,
		ERROR_VOLMGR_EXTENT_OVERLAPS_EBR_PARTITION,
		ERROR_VOLMGR_EXTENT_VOLUME_LENGTHS_DO_NOT_MATCH,
		ERROR_VOLMGR_FAULT_TOLERANT_NOT_SUPPORTED,
		ERROR_VOLMGR_INTERLEAVE_LENGTH_INVALID,
		ERROR_VOLMGR_MAXIMUM_REGISTERED_USERS,
		ERROR_VOLMGR_MEMBER_IN_SYNC,
		ERROR_VOLMGR_MEMBER_INDEX_DUPLICATE,
		ERROR_VOLMGR_MEMBER_INDEX_INVALID,
		ERROR_VOLMGR_MEMBER_MISSING,
		ERROR_VOLMGR_MEMBER_NOT_DETACHED,
		ERROR_VOLMGR_MEMBER_REGENERATING,
		ERROR_VOLMGR_ALL_DISKS_FAILED,
		ERROR_VOLMGR_NO_REGISTERED_USERS,
		ERROR_VOLMGR_NO_SUCH_USER,
		ERROR_VOLMGR_NOTIFICATION_RESET,
		ERROR_VOLMGR_NUMBER_OF_MEMBERS_INVALID,
		ERROR_VOLMGR_NUMBER_OF_PLEXES_INVALID,
		ERROR_VOLMGR_PACK_DUPLICATE,
		ERROR_VOLMGR_PACK_ID_INVALID,
		ERROR_VOLMGR_PACK_INVALID,
		ERROR_VOLMGR_PACK_NAME_INVALID,
		ERROR_VOLMGR_PACK_OFFLINE,
		ERROR_VOLMGR_PACK_HAS_QUORUM,
		ERROR_VOLMGR_PACK_WITHOUT_QUORUM,
		ERROR_VOLMGR_PARTITION_STYLE_INVALID,
		ERROR_VOLMGR_PARTITION_UPDATE_FAILED,
		ERROR_VOLMGR_PLEX_IN_SYNC,
		ERROR_VOLMGR_PLEX_INDEX_DUPLICATE,
		ERROR_VOLMGR_PLEX_INDEX_INVALID,
		ERROR_VOLMGR_PLEX_LAST_ACTIVE,
		ERROR_VOLMGR_PLEX_MISSING,
		ERROR_VOLMGR_PLEX_REGENERATING,
		ERROR_VOLMGR_PLEX_TYPE_INVALID,
		ERROR_VOLMGR_PLEX_NOT_RAID5,
		ERROR_VOLMGR_PLEX_NOT_SIMPLE,
		ERROR_VOLMGR_STRUCTURE_SIZE_INVALID,
		ERROR_VOLMGR_TOO_MANY_NOTIFICATION_REQUESTS,
		ERROR_VOLMGR_TRANSACTION_IN_PROGRESS,
		ERROR_VOLMGR_UNEXPECTED_DISK_LAYOUT_CHANGE,
		ERROR_VOLMGR_VOLUME_CONTAINS_MISSING_DISK,
		ERROR_VOLMGR_VOLUME_ID_INVALID,
		ERROR_VOLMGR_VOLUME_LENGTH_INVALID,
		ERROR_VOLMGR_VOLUME_LENGTH_NOT_SECTOR_SIZE_MULTIPLE,
		ERROR_VOLMGR_VOLUME_NOT_MIRRORED,
		ERROR_VOLMGR_VOLUME_NOT_RETAINED,
		ERROR_VOLMGR_VOLUME_OFFLINE,
		ERROR_VOLMGR_VOLUME_RETAINED,
		ERROR_VOLMGR_NUMBER_OF_EXTENTS_INVALID,
		ERROR_VOLMGR_DIFFERENT_SECTOR_SIZE,
		ERROR_VOLMGR_BAD_BOOT_DISK,
		ERROR_VOLMGR_PACK_CONFIG_OFFLINE,
		ERROR_VOLMGR_PACK_CONFIG_ONLINE,
		ERROR_VOLMGR_NOT_PRIMARY_PACK,
		ERROR_VOLMGR_PACK_LOG_UPDATE_FAILED,
		ERROR_VOLMGR_NUMBER_OF_DISKS_IN_PLEX_INVALID,
		ERROR_VOLMGR_NUMBER_OF_DISKS_IN_MEMBER_INVALID,
		ERROR_VOLMGR_VOLUME_MIRRORED,
		ERROR_VOLMGR_PLEX_NOT_SIMPLE_SPANNED,
		ERROR_VOLMGR_NO_VALID_LOG_COPIES,
		ERROR_VOLMGR_PRIMARY_PACK_PRESENT,
		ERROR_VOLMGR_NUMBER_OF_DISKS_INVALID,
		ERROR_VOLMGR_MIRROR_NOT_SUPPORTED,
		ERROR_VOLMGR_RAID5_NOT_SUPPORTED,
		ERROR_BCD_NOT_ALL_ENTRIES_IMPORTED,
		ERROR_BCD_TOO_MANY_ELEMENTS,
		ERROR_BCD_NOT_ALL_ENTRIES_SYNCHRONIZED,
		ERROR_VHD_DRIVE_FOOTER_MISSING,
		ERROR_VHD_DRIVE_FOOTER_CHECKSUM_MISMATCH,
		ERROR_VHD_DRIVE_FOOTER_CORRUPT,
		ERROR_VHD_FORMAT_UNKNOWN,
		ERROR_VHD_FORMAT_UNSUPPORTED_VERSION,
		ERROR_VHD_SPARSE_HEADER_CHECKSUM_MISMATCH,
		ERROR_VHD_SPARSE_HEADER_UNSUPPORTED_VERSION,
		ERROR_VHD_SPARSE_HEADER_CORRUPT,
		ERROR_VHD_BLOCK_ALLOCATION_FAILURE,
		ERROR_VHD_BLOCK_ALLOCATION_TABLE_CORRUPT,
		ERROR_VHD_INVALID_BLOCK_SIZE,
		ERROR_VHD_BITMAP_MISMATCH,
		ERROR_VHD_PARENT_VHD_NOT_FOUND,
		ERROR_VHD_CHILD_PARENT_ID_MISMATCH,
		ERROR_VHD_CHILD_PARENT_TIMESTAMP_MISMATCH,
		ERROR_VHD_METADATA_READ_FAILURE,
		ERROR_VHD_METADATA_WRITE_FAILURE,
		ERROR_VHD_INVALID_SIZE,
		ERROR_VHD_INVALID_FILE_SIZE,
		ERROR_VIRTDISK_PROVIDER_NOT_FOUND,
		ERROR_VIRTDISK_NOT_VIRTUAL_DISK,
		ERROR_VHD_PARENT_VHD_ACCESS_DENIED,
		ERROR_VHD_CHILD_PARENT_SIZE_MISMATCH,
		ERROR_VHD_DIFFERENCING_CHAIN_CYCLE_DETECTED,
		ERROR_VHD_DIFFERENCING_CHAIN_ERROR_IN_PARENT,
		ERROR_VIRTUAL_DISK_LIMITATION,
		ERROR_VHD_INVALID_TYPE,
		ERROR_VHD_INVALID_STATE,
		ERROR_VIRTDISK_UNSUPPORTED_DISK_SECTOR_SIZE,
		ERROR_QUERY_STORAGE_ERROR,
		SDIAG_E_CANCELLED,
		SDIAG_E_SCRIPT,
		SDIAG_E_POWERSHELL,
		SDIAG_E_MANAGEDHOST,
		SDIAG_E_NOVERIFIER,
		SDIAG_S_CANNOTRUN,
		SDIAG_E_DISABLED,
		SDIAG_E_TRUST,
		SDIAG_E_CANNOTRUN,
		SDIAG_E_VERSION,
		SDIAG_E_RESOURCE,
		SDIAG_E_ROOTCAUSE,
		E_MBN_CONTEXT_NOT_ACTIVATED,
		E_MBN_BAD_SIM,
		E_MBN_DATA_CLASS_NOT_AVAILABLE,
		E_MBN_INVALID_ACCESS_STRING,
		E_MBN_MAX_ACTIVATED_CONTEXTS,
		E_MBN_PACKET_SVC_DETACHED,
		E_MBN_PROVIDER_NOT_VISIBLE,
		E_MBN_RADIO_POWER_OFF,
		E_MBN_SERVICE_NOT_ACTIVATED,
		E_MBN_SIM_NOT_INSERTED,
		E_MBN_VOICE_CALL_IN_PROGRESS,
		E_MBN_INVALID_CACHE,
		E_MBN_NOT_REGISTERED,
		E_MBN_PROVIDERS_NOT_FOUND,
		E_MBN_PIN_NOT_SUPPORTED,
		E_MBN_PIN_REQUIRED,
		E_MBN_PIN_DISABLED,
		E_MBN_FAILURE,
		E_MBN_INVALID_PROFILE,
		E_MBN_DEFAULT_PROFILE_EXIST,
		E_MBN_SMS_ENCODING_NOT_SUPPORTED,
		E_MBN_SMS_FILTER_NOT_SUPPORTED,
		E_MBN_SMS_INVALID_MEMORY_INDEX,
		E_MBN_SMS_LANG_NOT_SUPPORTED,
		E_MBN_SMS_MEMORY_FAILURE,
		E_MBN_SMS_NETWORK_TIMEOUT,
		E_MBN_SMS_UNKNOWN_SMSC_ADDRESS,
		E_MBN_SMS_FORMAT_NOT_SUPPORTED,
		E_MBN_SMS_OPERATION_NOT_ALLOWED,
		E_MBN_SMS_MEMORY_FULL,
		UI_E_CREATE_FAILED,
		UI_E_SHUTDOWN_CALLED,
		UI_E_ILLEGAL_REENTRANCY,
		UI_E_OBJECT_SEALED,
		UI_E_VALUE_NOT_SET,
		UI_E_VALUE_NOT_DETERMINED,
		UI_E_INVALID_OUTPUT,
		UI_E_BOOLEAN_EXPECTED,
		UI_E_DIFFERENT_OWNER,
		UI_E_AMBIGUOUS_MATCH,
		UI_E_FP_OVERFLOW,
		UI_E_WRONG_THREAD,
		UI_E_STORYBOARD_ACTIVE,
		UI_E_STORYBOARD_NOT_PLAYING,
		UI_E_START_KEYFRAME_AFTER_END,
		UI_E_END_KEYFRAME_NOT_DETERMINED,
		UI_E_LOOPS_OVERLAP,
		UI_E_TRANSITION_ALREADY_USED,
		UI_E_TRANSITION_NOT_IN_STORYBOARD,
		UI_E_TRANSITION_ECLIPSED,
		UI_E_TIME_BEFORE_LAST_UPDATE,
		UI_E_TIMER_CLIENT_ALREADY_CONNECTED,
	}
}
